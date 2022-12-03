package xgb

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"sync"
	"sync/atomic"

	"codeberg.org/gruf/go-byteutil"
	"codeberg.org/gruf/go-xgb/internal"
)

var (
	// ErrConnClosed is returned on attempts to use a closed connection.
	ErrConnClosed = errors.New("x connection closed")

	// verbose tracks whether verbose logging is enabled via $GODEBUG.
	verbose = strings.Contains(os.Getenv("GODEBUG"), "xgbdebug")
)

// XConn ...
type XConn struct {
	conn net.Conn                              // underlying net.Conn
	xidg xidGenerator                          // generates new XIDs
	inCh chan any                              // inbound event / error channel
	popd *cookie                               // popped cookie from queue
	ckQu internal.Queue[*cookie]               // queue of waiting cookies
	evfn internal.Map[uint8, EventUnmarshaler] // map of registered event no. to event unmarshaler funcs
	erfn internal.Map[uint8, ErrorUnmarshaler] // map of registered error no. to error unmarshaler funcs
	exts internal.Map[string, uint8]           // map of opcodes to extensions
	logf func(string, ...any)                  // user provided log output function
	done chan struct{}                         // closed on conn closed, used to select against cookie channels
	seq  uint32                                // current sequence no. (updates atomically)
	cls  uint32                                // tracks if conn closed
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
	// Check if connection already closed
	if atomic.LoadUint32(&conn.cls) != 0 {
		return ErrConnClosed
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
func (conn *XConn) Send(data []byte) error {
	_ = atomic.AddUint32(&conn.seq, 1) // iter sequence
	return conn.write(data)            // write data
}

// SendRecv ...
func (conn *XConn) SendRecv(data []byte, dst XReply) error {
	// Acquire cookie from pool
	ck := acquireCookie()
	defer releaseCookie(ck)

	// Set unmarshal dst
	ck.dst = dst

	// Get next sequence ID and register cookie
	ck.id = uint16(atomic.AddUint32(&conn.seq, 1))
	conn.ckQu.Push(ck)

	// Write request data to X server
	if err := conn.write(data); err != nil {
		return err
	}

	if dst == nil {
		// force sync with X
		_ = conn.Send(inputfocus())
	}

	// Wait on rsp
	return <-ck.err
}

// Recv ...
func (conn *XConn) Recv() (XEvent, error) {
	// Check if connection already closed
	if atomic.LoadUint32(&conn.cls) != 0 {
		return nil, ErrConnClosed
	}

	// Wait on next event
	v, ok := <-conn.inCh
	if !ok {
		return nil, ErrConnClosed
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
	return conn.SendRecv(inputfocus(), IgnoreXReply{})
}

// Close will close the X connection.
func (conn *XConn) Close() error {
	if atomic.CompareAndSwapUint32(&conn.cls, 0, 1) {
		close(conn.inCh)         // close inbound ch
		return conn.conn.Close() // close X
	}
	return nil
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

// cookie will attempt to pop the queued cookie with given sequence number.
func (conn *XConn) cookie(seq uint16) (*cookie, bool) {
	if ck := conn.popd; ck != nil {
		// Previously popped cookie

		switch {
		// Out of date cookie
		case ck.id < seq:
			conn.popd = nil
			releaseCookie(ck)

		// Cookie ahead of this
		case ck.id > seq:
			return nil, false

		// We found it!
		default:
			conn.popd = nil
			return ck, true
		}
	}

	for {
		// Pop next cookie from queue
		ck, ok := conn.ckQu.Pop()
		if !ok {
			return nil, false
		}

		switch {
		// Out of date cookie
		case ck.id < seq:
			releaseCookie(ck)

		// Cookie ahead of this
		case ck.id > seq:
			conn.popd = ck // store
			return nil, false

		// We found it!
		default:
			return ck, true
		}
	}
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

		if ck := conn.popd; ck != nil {
			// Close cookie
			close(ck.err)
		}

		for {
			// Pop all cookies from queue
			ck, ok := conn.ckQu.Pop()
			if !ok {
				break
			}

			// Finished with cookie
			releaseCookie(ck)
		}
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
			if ck, ok := conn.cookie(xerr.SeqID()); ok {
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
			if ck, ok := conn.cookie(seq); ok {
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
func (conn *XConn) write(data []byte) error {
	// Check if connection already closed
	if atomic.LoadUint32(&conn.cls) != 0 {
		return ErrConnClosed
	}

	// Write data to the underlying conn
	_, err := conn.conn.Write(data)
	if err != nil {
		// Write error occurred, wrap with context
		err = fmt.Errorf("fatal xconn write error: %w", err)
		_ = conn.Close()
		return err
	}

	conn.debugf("send data=%v\n", data)
	return nil
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
	id  uint16
	err chan error
	dst XReply
}

// acquireCookie will acquire a fresh cookie from the pool.
func acquireCookie() *cookie {
	return cookiePool.Get().(*cookie)
}

// releaseCookie will reset the cookie, drain it and release to pool.
func releaseCookie(ck *cookie) {
	// Reset fields
	ck.id = 0
	ck.dst = nil

	select {
	case ck.err <- nil:
	case <-ck.err:
	default:
	}

	// Place in pool
	cookiePool.Put(ck)
}

// inputfocus returns a newly prepared input focus request.
func inputfocus() []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = 43 // request opcode
	b += 1

	b += 1                                                 // padding
	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}
