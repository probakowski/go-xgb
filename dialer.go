package xgb

import (
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"log"
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

type XDialer struct {
	// InboundBuffer ...
	InboundBuffer int

	// Log allows specifying a log output method, default is simply log.Printf().
	Log func(format string, args ...any)

	// NetDialer allows specifying the underlying net.Dialer to use.
	NetDialer *net.Dialer
}

func Dial(display string) (*XConn, []byte, error) {
	return DefaultDialer.Dial(display)
}

func DialContext(ctx context.Context, display string) (*XConn, []byte, error) {
	return DefaultDialer.DialContext(ctx, display)
}

// Dial ...
func (d *XDialer) Dial(display string) (*XConn, []byte, error) {
	return d.DialContext(context.Background(), display)
}

// DialContext ...
func (d *XDialer) DialContext(ctx context.Context, display string) (*XConn, []byte, error) {
	if len(display) == 0 {
		display = os.Getenv("DISPLAY")
	}

	//
	display0 := display

	colonIdx := strings.LastIndex(display, ":")
	if colonIdx < 0 {
		return nil, nil, fmt.Errorf("bad display string %q", display0)
	}

	var host string
	var protocol, socket string

	if display[0] == '/' {
		socket = display[0:colonIdx]
	} else {
		slashIdx := strings.LastIndex(display, "/")
		if slashIdx >= 0 {
			protocol = display[0:slashIdx]
			host = display[slashIdx+1 : colonIdx]
		} else {
			host = display[0:colonIdx]
		}
	}

	display = display[colonIdx+1:]
	if len(display) == 0 {
		return nil, nil, fmt.Errorf("bad display string %q", display0)
	}

	dotIdx := strings.LastIndex(display, ".")
	if dotIdx < 0 {
		display = display[0:]
	} else {
		display = display[0:dotIdx]
	}

	dispNum, err := strconv.Atoi(display)
	if err != nil || dispNum < 0 {
		return nil, nil, fmt.Errorf("bad display string %q", display0)
	}

	var conn net.Conn

	if len(socket) != 0 {
		// Dial unix socket address at display number
		conn, err = net.Dial("unix", socket+":"+display)
	} else if len(host) != 0 && host != "unix" {
		// default proto is tcp
		if protocol == "" {
			protocol = "tcp"
		}

		conn, err = net.Dial(protocol, host+":"+strconv.Itoa(6000+dispNum))
	} else {
		host = ""

		// Dial the default tmp unix X11 generated socket path
		conn, err = net.Dial("unix", "/tmp/.X11-unix/X"+display)
	}

	if err != nil {
		return nil, nil, fmt.Errorf("cannot connect to %q: %w", display0, err)
	}

	// Attempt to get XAuthority data necessary to start
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
	// Build the initial authorization request
	buf := make([]byte, 12+internal.Pad4(len(authName))+internal.Pad4(len(authData)))

	buf[0] = 0x6c
	buf[1] = 0

	binary.LittleEndian.PutUint16(buf[2:], 11)
	binary.LittleEndian.PutUint16(buf[4:], 0)
	binary.LittleEndian.PutUint16(buf[6:], uint16(len(authName)))
	binary.LittleEndian.PutUint16(buf[8:], uint16(len(authData)))
	binary.LittleEndian.PutUint16(buf[10:], 0)

	copy(buf[12:], authName)
	copy(buf[12+internal.Pad4(len(authName)):], authData)

	// Write auth request to connection
	if _, err := conn.Write(buf); err != nil {
		return nil, nil, err
	}

	head := make([]byte, 8)

	// Read response from server
	if _, err := io.ReadFull(conn, head[0:8]); err != nil {
		return nil, nil, err
	}

	// Check that this is expect X server version
	major := binary.LittleEndian.Uint16(head[2:])
	minor := binary.LittleEndian.Uint16(head[4:])
	if major != 11 || minor != 0 {
		return nil, nil, fmt.Errorf("x protocol version mismatch: %d.%d", major, minor)
	}

	// Prepare buffer for next group of data
	dataLen := binary.LittleEndian.Uint16(head[6:])
	data := make([]byte, int(dataLen)*4+8, int(dataLen)*4+8)
	copy(data, head)

	// Read the next group of data into buffer
	if _, err := io.ReadFull(conn, data[8:]); err != nil {
		return nil, nil, err
	}

	// Check authentication return code
	if code := head[0]; code == 0 {
		reasonLen := head[1]
		reason := data[8 : 8+reasonLen]
		return nil, nil, fmt.Errorf("x protocol authentication refused: %s", reason)
	}

	// read necessary XID gen information
	resourceIDBase := binary.LittleEndian.Uint32(data[12:])
	resourceIDMask := binary.LittleEndian.Uint32(data[16:])

	var logf func(string, ...any)

	if logf = d.Log; logf == nil {
		// Set default logger
		logf = log.Printf
	}

	// wrap the net.Conn in *xconn
	xconn := &XConn{
		conn: conn,
		xidg: xidGenerator{
			base: resourceIDBase,
			last: 0,
			inc:  resourceIDMask & -resourceIDMask,
			max:  resourceIDMask,
		},
		inCh: make(chan any, d.InboundBuffer),
		ckQu: internal.NewQueue[*cookie](),
		evfn: internal.NewMap[uint8, EventUnmarshaler](),
		erfn: internal.NewMap[uint8, ErrorUnmarshaler](),
		exts: internal.NewMap[string, XExtension](),
		logf: logf,
		done: make(chan struct{}),
	}

	// Start conn read-loop
	go xconn.readloop()

	return xconn, data, nil
}
