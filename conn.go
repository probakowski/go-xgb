package xgb

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"sync"

	"codeberg.org/gruf/go-byteutil"
	"codeberg.org/gruf/go-xgb/internal"
)

// verbose tracks whether verbose logging is enabled via $GODEBUG.
var verbose = strings.Contains(os.Getenv("GODEBUG"), "xgbdebug")

// XConn ...
type XConn struct {
	conn net.Conn             // underlying network connection
	logf func(string, ...any) // user provided log output function
	inCh chan any             // inbound event / error channel
	done chan struct{}        // closed on conn closed, used to select against cookie channels

	xidg xidGenerator // generates new XIDs

	seq  uint16     // cookie (event) sequence number
	popd *cookie    // popped cookie from queue
	ckQu []*cookie  // queue of waiting cookies
	mu   sync.Mutex // conn mutex: protects seq,ckQu and writes

	evfn internal.Map[uint8, EventUnmarshaler] // map of registered event no. to event unmarshaler funcs
	erfn internal.Map[uint8, ErrorUnmarshaler] // map of registered error no. to error unmarshaler funcs
	exts internal.Map[string, uint8]           // map of opcodes to extensions
}

// Register ...
func (conn *XConn) Register(ext XExtension) error {
	// Add this extension to our map
	if !conn.exts.Set(ext.XName, ext.MajorOpcode) {
		return fmt.Errorf("already registered ext %q", ext.XName)
	}

	// Iniitialize X type unmarshalers
	return conn.init(ext.EventFuncs, ext.ErrorFuncs)
}

// xproto_init is a package-private function used by xproto (with go linkname hackery) to initialize the xproto extension.
func xproto_init(conn *XConn, eventFuncs map[uint8]EventUnmarshaler, errorFuncs map[uint8]ErrorUnmarshaler) error {
	return conn.init(eventFuncs, errorFuncs)
}

// init ...
func (conn *XConn) init(eventFuncs map[uint8]EventUnmarshaler, errorFuncs map[uint8]ErrorUnmarshaler) error {
	select {
	// Check if closed
	case <-conn.done:
		return net.ErrClosed
	default:
	}

	// Copy over the extension event unmarshalers
	for n, fn := range eventFuncs {
		if !conn.evfn.Set(uint8(n), fn) {
			return fmt.Errorf("overlapping event unmarshaler for %d", n)
		}
	}

	// Copy over the extension error unmarshalers
	for n, fn := range errorFuncs {
		if !conn.erfn.Set(uint8(n), fn) {
			return fmt.Errorf("overlapping error unmarshaler for %d", n)
		}
	}

	return nil
}

// Ext returns the major opcode for extension with given name.
func (conn *XConn) Ext(name string) (op uint8, ok bool) {
	return conn.exts.Get(name)
}

// NextID ...
func (conn *XConn) NewXID() (uint32, error) {
	return conn.xidg.Next()
}

// Send will send given data to the X server.
func (conn *XConn) Send(data []byte) (err error) {
	// Acquire cookie from pool,
	// conn.cookie() will release.
	ck := acquireCookie()

	// Acquire conn lock
	conn.mu.Lock()

	// Write data to X conn, set sequence
	ck.seq, err = conn.write(data)
	if err != nil {
		conn.mu.Unlock()
		return
	}

	// Register cookie in conn queue
	conn.ckQu = append(conn.ckQu, ck)

	// Unlock conn
	conn.mu.Unlock()

	return
}

// SendRecv ...
func (conn *XConn) SendRecv(data []byte, dst XReply) (err error) {
	// Acquire cookie from pool
	ck := acquireCookie()
	defer releaseCookie(ck)

	// SendRecv is synchronous
	ck.syn = true

	// Set unmarshal dst
	ck.dst = dst

	// Acquire conn lock
	conn.mu.Lock()

	// Write data to X conn, set sequence
	ck.seq, err = conn.write(data)
	if err != nil {
		conn.mu.Unlock()
		return
	}

	// Register cookie in conn queue
	conn.ckQu = append(conn.ckQu, ck)

	// Unlock conn
	conn.mu.Unlock()

	if dst == nil {
		// force sync with X
		_ = conn.Send([]byte{
			0x2b, 0x0, 0x1, 0x0,
		})
	}

	// Wait on event/err
	return <-ck.err
}

// Recv ...
func (conn *XConn) Recv() (XEvent, error) {
	// Wait on next event
	v, ok := <-conn.inCh
	if !ok {
		return nil, net.ErrClosed
	}

	switch v := v.(type) {
	case XEvent:
		return v, nil
	case error:
		return nil, v
	default:
		return nil, fmt.Errorf("BUG: unexpected type %T down cookie channel", v)
	}
}

// Sync will force a roundtrip to the X server, by sending a GetInputFocus request and blocking on response.
func (conn *XConn) Sync() error {
	return conn.SendRecv([]byte{0x2b, 0x0, 0x1, 0x0}, IgnoreXReply{})
}

// Close will close the X connection.
func (conn *XConn) Close() error {
	select {
	case <-conn.done:
		// already closed
		return nil

	default:
		// close connection
		close(conn.inCh)
		close(conn.done)
		return conn.conn.Close()
	}
}

// unmarshalEvent will attempt to unmarshal event data 'b' as event type with number 'n'.
func (conn *XConn) unmarshalEvent(n uint8, b []byte) (XEvent, error) {
	if um, ok := conn.evfn.Get(n); ok {
		return um(b)
	}
	return nil, fmt.Errorf("unknown event type %d", n)
}

// unmarshalError will attempt to unmarshal error data 'b' as error type with number 'n'.
func (conn *XConn) unmarshalError(n uint8, b []byte) (XError, error) {
	if um, ok := conn.erfn.Get(n); ok {
		return um(b)
	}
	return nil, fmt.Errorf("unknown error type %d", n)
}

// getCookie will attempt to pop the queued cookie with given sequence number.
func (conn *XConn) getCookie(seq uint16) (*cookie, bool) {
	if ck := conn.popd; ck != nil {
		// Previously popped cookie

		switch {
		// Out of date cookie
		case ck.seq < seq:
			conn.popd = nil
			ck.send(nil) // ping
			releaseCookie(ck)

		// Cookie ahead of this
		case ck.seq > seq:
			return nil, false

		// We found it!
		default:
			conn.popd = nil
			return ck, true
		}
	}

	for len(conn.ckQu) > 0 {
		// Get front of queue
		ck := conn.ckQu[0]

		// Drop front of queue
		nl := len(conn.ckQu) - 1
		copy(conn.ckQu, conn.ckQu[:nl])
		conn.ckQu = conn.ckQu[:nl]

		switch {
		// Out of date cookie
		case ck.seq < seq:
			ck.send(nil) // ping
			releaseCookie(ck)

		// Cookie ahead of this
		case ck.seq > seq:
			conn.popd = ck // store
			return nil, false

		// We found it!
		default:
			return ck, true
		}
	}

	return nil, false
}

// readloop ...
func (conn *XConn) readloop() {
	var (
		// buf is the main read buffer.
		buf [32]byte

		// lbuf is a larger buffer used when an x reply
		// is received of size > 32 bytes.
		lbuf byteutil.Buffer

		// prepLargeBuf will prepare the larger buffer for
		// a received x reply of given size, also copying
		// over the initial 32byte contents of small buf.
		prepLargeBuf = func(sz uint32) {
			lbuf.Reset() // ensure empty
			lbuf.Grow(32 + int(sz)*4)
			_ = copy(lbuf.B[:32], buf[:])
		}
	)

	defer func() {
		// Close connection
		_ = conn.Close()

		// Forcibly outdate to drop all cookies
		_, _ = conn.getCookie(^uint16(0))
	}()

	for {
		// Read next set of data from X into buf
		_, err := io.ReadFull(conn.conn, buf[:])
		if err != nil {
			conn.logf("fatal xconn read error: %v\n", err)
			return
		}

		switch buf[0] {
		// Error
		case 0:
			// Attempt to unmarshal X error type
			xerr, err := conn.unmarshalError(buf[1], buf[:])
			if err != nil {
				conn.logf("unable to unmarshal x error: %v\n", err)
				continue
			}

			// Debug log raw error bytes
			conn.debugf("recv xerror=%v\n", buf[:])

			// Look for cookie waiting for response
			if ck, ok := conn.getCookie(xerr.SeqID()); ok {
				ck.err <- xerr
				continue
			}

			select {
			// Conn closed
			case <-conn.done:
				return

			// Pass error to inbound
			case conn.inCh <- xerr:
			}

		// Reply
		case 1:
			var reply []byte

			// Get sequence ID + reply size
			seq := binary.LittleEndian.Uint16(buf[2:])
			size := binary.LittleEndian.Uint32(buf[4:])

			if size > 0 {
				// More bytes to read
				prepLargeBuf(size)

				// Read next bytes into the larger buffer
				_, err := io.ReadFull(conn.conn, lbuf.B[32:])
				if err != nil {
					conn.logf("fatal xconn read error: %v\n", err)
					conn.inCh <- err
					return
				}

				// Set new reply data
				reply = lbuf.B
			} else {
				// Use existing data
				reply = buf[:]
			}

			// Debug log raw reply bytes
			conn.debugf("recv xreply=%v\n", reply)

			// Look for cookie waiting for response
			if ck, ok := conn.getCookie(seq); ok {
				var err error

				if ck.dst != nil {
					// Unmarshal reply into destination obj
					if err = ck.dst.Unmarshal(reply); err != nil {
						conn.logf("unable to unmarshal x reply: %v\n", err)
					}
				}

				// Trigger return
				ck.err <- err
			}

		default:
			// Attempt to unmarshal X event type
			xev, err := conn.unmarshalEvent(buf[0]&127, buf[:])
			if err != nil {
				conn.logf("unable to unmarshal x event: %v\n", err)
				continue
			}

			// Debug log raw event bytes
			conn.debugf("recv xevent=%v\n", buf[:])

			select {
			// Conn closed
			case <-conn.done:
				return

			// Pass event to inbound
			case conn.inCh <- xev:
			}
		}
	}
}

// write ...
func (conn *XConn) write(data []byte) (seq uint16, err error) {
	// Write data to underlying connection
	if _, err = conn.conn.Write(data); err != nil {
		_ = conn.Close()
		return
	}

	// Debug log sent data
	conn.debugf("send data=%v\n", data)

	// Iter sequence
	conn.seq++
	seq = conn.seq

	return
}

// debugf will log given format string and args only if debugging is enabled.
func (conn *XConn) debugf(format string, args ...any) {
	if verbose {
		conn.logf(format, args...)
	}
}

// cookiePool is a memory pool of X cookies for use in SendRecv() requests.
var cookiePool = sync.Pool{
	New: func() any {
		return &cookie{err: make(chan error)}
	},
}

// cookie ...
type cookie struct {
	seq uint16
	err chan error
	dst XReply
	syn bool
}

func (ck *cookie) send(err error) {
	if !ck.syn {
		return
	}
	ck.err <- err
}

// acquireCookie will acquire a fresh cookie from the pool.
func acquireCookie() *cookie {
	return cookiePool.Get().(*cookie)
}

// releaseCookie will reset the cookie, drain it and release to pool.
func releaseCookie(ck *cookie) {
	// Reset fields
	ck.seq = 0
	ck.dst = nil
	ck.syn = false

	// Place in pool
	cookiePool.Put(ck)
}
