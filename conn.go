package xgb

import (
	"codeberg.org/gruf/go-byteutil"
	"container/list"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"sync"
	"unsafe"

	"github.com/probakowski/go-xgb/internal"
)

// le is a shorthand for the littleendian binary enc/dec.
var le = binary.LittleEndian

// XConn represents a connection to an X server, handling
// asychronous background receipt and unmarshaling of incoming
// data for use as the more accessible XError and XEvent types.
type XConn struct {
	conn net.Conn      // underlying network connection
	inCh chan any      // inbound event / error channel
	done chan struct{} // closed on conn closed, used to select against cookie channels

	xidg xidGenerator // generates new XIDs

	seq  uint16     // cookie (event) sequence number
	trnc bool       // sequence truncate flag, set on uint16 overflow
	popd *cookie    // popped cookie from queue
	ckQu list.List  // queue of waiting cookies
	mu   sync.Mutex // conn mutex: protects seq,ckQu and writes

	evfn internal.SmallPtrMap        // "map" of registered event no. to event unmarshaler funcs
	erfn internal.SmallPtrMap        // "map" of registered error no. to error unmarshaler funcs
	exts internal.Map[string, uint8] // map of opcodes to extensions
}

// Register querying the X server for support of this extension, and register relevant event / error unmarshalers internally within XConn.
func (conn *XConn) Register(ext XExtension) error {
	// Attempt to store major op code for this ext name.
	if !conn.exts.Add(ext.XName, ext.MajorOpcode) {
		return fmt.Errorf("already registered ext: %s", ext.XName)
	}

	// Iniitialize X type unmarshalers and store them.
	return conn.init(ext.EventFuncs, ext.ErrorFuncs)
}

// xproto_init is a package-private function used by xproto (with go linkname hackery) to initialize the xproto extension.
func xproto_init(conn *XConn, eventFuncs map[uint8]EventUnmarshaler, errorFuncs map[uint8]ErrorUnmarshaler) error {
	return conn.init(eventFuncs, errorFuncs)
}

func (conn *XConn) init(eventFuncs map[uint8]EventUnmarshaler, errorFuncs map[uint8]ErrorUnmarshaler) error {
	select {
	// Check if closed
	case <-conn.done:
		return net.ErrClosed
	default:
	}

	// Copy over the extension event unmarshalers
	for n := range eventFuncs {
		fn := eventFuncs[n]
		ptr := unsafe.Pointer(&fn)
		if !conn.evfn.Set(uint8(n), ptr) {
			return fmt.Errorf("overlapping event unmarshaler for %d", n)
		}
	}

	// Copy over the extension error unmarshalers
	for n := range errorFuncs {
		fn := errorFuncs[n]
		ptr := unsafe.Pointer(&fn)
		if !conn.erfn.Set(uint8(n), ptr) {
			return fmt.Errorf("overlapping error unmarshaler for %d", n)
		}
	}

	return nil
}

// Ext returns the major opcode for extension with given name.
func (conn *XConn) Ext(name string) (op uint8, ok bool) {
	return conn.exts.Get(name)
}

// NextID returns the next available X ID.
func (conn *XConn) NewXID() uint32 {
	return conn.xidg.Next()
}

// Send will send given data to the X server.
func (conn *XConn) Send(data []byte) error {
	conn.mu.Lock()
	_, err := conn.write(data)
	conn.mu.Unlock()
	return err
}

// SendRecv sends the given data, and blocks until receipt of XReply into given container.
func (conn *XConn) SendRecv(data []byte, dst XReply) error {
	// Acquire cookie from pool
	ck := acquireCookie()
	defer releaseCookie(ck)

	// Acquire conn lock
	conn.mu.Lock()

	// Write data to X connection
	seq, err := conn.write(data)
	if err != nil {
		conn.mu.Unlock()
		return err
	}

	// Setup cookie
	ck.prime(seq, dst)
	conn.ckQu.PushBack(ck)

	if dst == nil {
		// Force "sync" with X by sending
		// simple request with known expected
		// response -- GetInputFocus request.
		_, _ = conn.write([]byte{
			0x2b, 0x0, 0x1, 0x0,
		})
	}

	// Unlock conn
	conn.mu.Unlock()

	// Wait on event/err
	return ck.await()
}

// Recv blocks until receipt of next XEvent / error.
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
		panic(fmt.Sprintf("BUG: unexpected inbound channel type %T", v))
	}
}

// Sync will force a roundtrip to X server, by sending
// a GetInputFocus request and blocking on X response.
func (conn *XConn) Sync() error {
	return conn.SendRecv([]byte{
		0x2b, 0x0, 0x1, 0x0,
	}, IgnoreXReply{})
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
	if ptr := conn.evfn.Get(n); ptr != nil {
		return (*(*EventUnmarshaler)(ptr))(b)
	}
	return nil, fmt.Errorf("unknown event type %d", n)
}

// unmarshalError will attempt to unmarshal error data 'b' as error type with number 'n'.
func (conn *XConn) unmarshalError(n uint8, b []byte) (XError, error) {
	if ptr := conn.erfn.Get(n); ptr != nil {
		return (*(*ErrorUnmarshaler)(ptr))(b)
	}
	return nil, fmt.Errorf("unknown error type %d", n)
}

// getCookie will attempt to pop the queued cookie with given sequence number.
func (conn *XConn) getCookie(seq uint16) (*cookie, bool) {
	// Acquire conn lock
	conn.mu.Lock()

	for {
		// Pop cookie at front of queue
		ck, ok := conn.popCookie()
		if !ok {
			conn.mu.Unlock()
			return nil, false
		}

		switch {
		// This is the cookie!
		case ck.seq == seq:
			conn.mu.Unlock()
			return ck, true

		// Out of date cookie (uint16 overflow)
		case conn.trnc && ck.seq > seq:
			var err error
			if ck.waiting() {
				err = errors.New("received no reply")
			}
			ck.trigger(err)   // ping
			releaseCookie(ck) //
			conn.trnc = false // reset

		// Cookie is ahead (uint16 overflow)
		case conn.trnc && ck.seq < seq:
			conn.popd = ck    // store
			conn.trnc = false // reset
			conn.mu.Unlock()
			return nil, false

		// Out of date cookie
		case ck.seq < seq:
			var err error
			if ck.waiting() {
				err = errors.New("received no reply")
			}
			ck.trigger(err) // ping
			releaseCookie(ck)

		// Cookie is ahead
		case ck.seq > seq:
			conn.popd = ck // store
			conn.mu.Unlock()
			return nil, false
		}
	}
}

func (conn *XConn) popCookie() (*cookie, bool) {
	// Look for already popped cookie
	if ck := conn.popd; ck != nil {
		conn.popd = nil
		return ck, true
	}

	// Grab first cookie in queue
	elem := conn.ckQu.Front()

	if elem == nil { // none queued
		return nil, false
	}

	// Drop front queue elem
	conn.ckQu.Remove(elem)

	return elem.Value.(*cookie), true
}

func (conn *XConn) readloop() {
	var (
		// buf is the main read buffer.
		buf = make([]byte, 32)

		// lbuf is a large buf used when an x
		// reply is received len > 32 bytes.
		lbuf byteutil.Buffer
	)

	defer func() {
		// Close connection.
		_ = conn.Close()

		// Forcibly outdate to drop all cookies.
		_, _ = conn.getCookie(^uint16(0))
	}()

	for {
		// Read next set of data from X into buf
		_, err := io.ReadFull(conn.conn, buf)
		if err != nil {
			logf("fatal xconn read error: %v\n", err)
			return
		}

		switch buf[0] {
		case 0 /* error */ :
			// Attempt to unmarshal X error type
			xerr, err := conn.unmarshalError(buf[1], buf)
			if err != nil {
				logf("unable to unmarshal x error: %v\n", err)
				continue
			}

			debugf("recv xerror=%v\n", buf[:])

			// Look for a cookie waiting for a response
			if ck, ok := conn.getCookie(xerr.SeqID()); ok {
				ck.trigger(xerr)
				continue
			}

			select {
			// Connection closed
			case <-conn.done:
				return

			// Pass error to inbound
			case conn.inCh <- xerr:
			}

		case 1 /* reply */ :
			// Decode reply data.
			seq := le.Uint16(buf[2:])
			size := le.Uint32(buf[4:])
			reply := buf[:]

			if size > 0 {
				// Prepare large buffer for
				// expected 32+size bytes.
				lbuf.Reset() // ensure empty
				lbuf.Grow(32 + int(size)*4)
				_ = copy(lbuf.B[:32], buf)

				// Read expected bytes into larger buffer.
				_, err := io.ReadFull(conn.conn, lbuf.B[32:])
				if err != nil {
					logf("fatal xconn read error: %v\n", err)
					conn.inCh <- err
					return
				}

				// Set new data.
				reply = lbuf.B
			}

			debugf("recv xreply=%v\n", reply)

			// Try pop a cookie expecting this reply.
			if ck, ok := conn.getCookie(seq); ok {
				var err error

				if ck.dst != nil {
					// Unmarshal reply into cookie destination obj.
					if err = ck.dst.Unmarshal(reply); err != nil {
						logf("unable to unmarshal x reply: %v\n", err)
					}
				}

				// Send err (if any).
				ck.trigger(err)
			}

		default /* xevent */ :
			// Attempt to unmarshal received X event type.
			xev, err := conn.unmarshalEvent(buf[0]&127, buf)
			if err != nil {
				logf("unable to unmarshal x event: %v\n", err)
				continue
			}

			// Drop any stale cookies waiting on
			// replies / errors up-to this event.
			ck, ok := conn.getCookie(xev.SeqID())
			if ok { // release
				ck.trigger(nil)
				releaseCookie(ck)
			}

			debugf("recv xevent=%v\n", buf)

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

// write will write the given data to underlying connection, update sequence
// and return the updated sequence number. must be calling within mutex lock.
func (conn *XConn) write(data []byte) (seq uint16, err error) {
	// Write data to underlying X connection.
	if _, err = conn.conn.Write(data); err != nil {
		logf("fatal xconn write error: %v\n")
		_ = conn.Close()
		return
	}

	debugf("send data=%v\n", data)

	// Iter sequence.
	seq = conn.seq
	if conn.seq += 1; conn.seq < seq {
		conn.trnc = true // uint16 overflow
	}

	return
}

// cookiePool is a memory pool of X
// cookies for use in SendRecv() requests.
var cookiePool sync.Pool

type cookie struct {
	dst XReply     // unmarshal dest
	err error      // error (if any)
	mtx sync.Mutex // underlying lock / block
	seq uint16     // sequence
}

func (ck *cookie) prime(seq uint16, dst XReply) {
	ck.mtx.Lock()
	ck.seq = seq
	ck.dst = dst
}

// trigger triggers the cookie to be released.
func (ck *cookie) trigger(err error) {
	ck.err = err
	ck.mtx.Unlock()
}

// await blocks on received reply / error.
func (ck *cookie) await() error {
	ck.mtx.Lock()
	err := ck.err
	ck.mtx.Unlock()
	return err
}

// waiting returns whether cookie is
// expecting a reply but still waiting.
func (ck *cookie) waiting() bool {
	return ck.dst != nil
}

// acquireCookie will acquire a fresh cookie from the pool.
func acquireCookie() *cookie {
	return new(cookie)
}

// releaseCookie will reset the cookie, drain it and release to pool.
func releaseCookie(ck *cookie) {
}

func logf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "[go-xgb] "+format, args...)
}
