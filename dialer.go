package xgb

import (
	"context"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"

	"codeberg.org/gruf/go-xgb/internal"
)

// defaultNetDialer is the default net.Dialer
// instance used when none is supplied.
var defaultNetDialer = net.Dialer{}

// DefaultDialer is the default XDialer instance.
var DefaultDialer = XDialer{
	InboundBuffer: 1000,
}

// XDialer ...
type XDialer struct {
	// InboundBuffer ...
	InboundBuffer int

	// NetDialer allows specifying the underlying net.Dialer to use.
	NetDialer *net.Dialer
}

// Dial ...
func Dial(display string) (*XConn, []byte, error) {
	return DefaultDialer.Dial(display)
}

// Dial ...
func (d *XDialer) Dial(display string) (*XConn, []byte, error) {
	return d.DialContext(context.Background(), display)
}

// DialContext ...
func (d *XDialer) DialContext(ctx context.Context, display string) (*XConn, []byte, error) {
	if display == "" {
		// By default grab from env.
		display = os.Getenv("DISPLAY")
	}

	// Keep original str for errors.
	display0 := display

	// Pull out first colon indicating display no.
	colonIdx := strings.LastIndex(display, ":")
	if colonIdx < 0 {
		return nil, nil, fmt.Errorf("bad display string %q", display0)
	}

	var host string
	var protocol, socket string

	if display[0] == '/' {
		// Filesystem location.
		socket = display[:colonIdx]
	} else {
		slashIdx := strings.LastIndex(display, "/")
		if slashIdx >= 0 {
			// Address with protocol.
			protocol = display[:slashIdx]
			host = display[slashIdx+1 : colonIdx]
		} else {
			// Simply an address.
			host = display[:colonIdx]
		}
	}

	// Reslice to drop past colon.
	display = display[colonIdx+1:]
	if display == "" {
		return nil, nil, fmt.Errorf("bad display string %q", display0)
	}

	//
	dotIdx := strings.LastIndex(display, ".")
	if dotIdx >= 0 {
		display = display[:dotIdx]
	}

	dispNum, err := strconv.Atoi(display)
	if err != nil || dispNum < 0 {
		return nil, nil, fmt.Errorf("bad display string %q", display0)
	}

	var conn net.Conn

	if socket != "" {
		// Dial unix socket address at display number.
		conn, err = net.Dial("unix", socket+":"+display)
	} else if host != "" && host != "unix" {
		// default proto is tcp
		if protocol == "" {
			protocol = "tcp"
		}

		// Dial the determined TCP protocol address at determined display no.
		conn, err = net.Dial(protocol, host+":"+strconv.Itoa(6000+dispNum))
	} else {
		host = ""

		// Dial the default tmp unix X11 generated socket path.
		conn, err = net.Dial("unix", "/tmp/.X11-unix/X"+display)
	}

	if err != nil {
		return nil, nil, fmt.Errorf("cannot connect to %q: %w", display0, err)
	}

	// Attempt to get XAuthority data necessary to start.
	authName, authData, err := ReadAuthority(host, display)
	if err != nil {
		authName = ""
		authData = []byte{}
	} else if authName != "MIT-MAGIC-COOKIE-1" || len(authData) != 16 {
		return nil, nil, fmt.Errorf("unsupported auth protocol %q", authName)
	}

	return d.DialConn(authName, authData, conn)
}

func (d *XDialer) DialConn(authName string, authData []byte, conn net.Conn) (*XConn, []byte, error) {
	// Build the initial authorization request into buffer.
	buf := make([]byte, 12+internal.Pad4(len(authName))+internal.Pad4(len(authData)))

	buf[0] = 0x6c
	buf[1] = 0

	le.PutUint16(buf[2:], 11)
	le.PutUint16(buf[4:], 0)
	le.PutUint16(buf[6:], uint16(len(authName)))
	le.PutUint16(buf[8:], uint16(len(authData)))
	le.PutUint16(buf[10:], 0)

	copy(buf[12:], authName)
	copy(buf[12+internal.Pad4(len(authName)):], authData)

	// Write auth request to connection.
	if _, err := conn.Write(buf); err != nil {
		return nil, nil, err
	}

	// Read response header from server.
	head := make([]byte, 8)
	if _, err := io.ReadFull(conn, head[0:8]); err != nil {
		return nil, nil, err
	}

	// Check that this is expect X server version.
	major := le.Uint16(head[2:])
	minor := le.Uint16(head[4:])
	if major != 11 || minor != 0 {
		return nil, nil, fmt.Errorf("x protocol version mismatch: %d.%d", major, minor)
	}

	// Prepare buffer for remainder of data.
	dataLen := le.Uint16(head[6:])
	data := make([]byte, int(dataLen)*4+8, int(dataLen)*4+8)
	copy(data, head)

	// Read the next group of data into buffer.
	if _, err := io.ReadFull(conn, data[8:]); err != nil {
		return nil, nil, err
	}

	// Check authentication return code.
	if code := head[0]; code == 0 {
		reasonLen := head[1]
		reason := data[8 : 8+reasonLen]
		return nil, nil, fmt.Errorf("x protocol authentication refused: %s", reason)
	}

	// Read necessary X resource ID information.
	resourceIDBase := le.Uint32(data[12:])
	resourceIDMask := le.Uint32(data[16:])

	// wrap the net.Conn in *xconn.
	xconn := &XConn{
		conn: conn,
		xidg: xidGenerator{
			base: resourceIDBase,
			last: 0,
			inc:  resourceIDMask & -resourceIDMask,
			max:  resourceIDMask,
		},
		inCh: make(chan any, d.InboundBuffer),
		done: make(chan struct{}),
		seq:  1,
	}

	// Start conn read-loop
	go xconn.readloop()

	return xconn, data, nil
}
