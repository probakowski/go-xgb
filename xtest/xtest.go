// FILE GENERATED AUTOMATICALLY FROM "xtest.xml"
package xtest

import (
	"encoding/binary"
	"errors"
	"fmt"

	"codeberg.org/gruf/go-xgb"
	"codeberg.org/gruf/go-xgb/xproto"
)

const (
	// ExtName is the user-friendly name string of this X extension.
	ExtName = "Test"

	// ExtXName is the name string this extension is known by to the X server.
	ExtXName = "XTEST"
)

var (
	// generated index maps of defined event and error numbers -> unmarshalers.
	eventFuncs = make(map[uint8]xgb.EventUnmarshaler)
	errorFuncs = make(map[uint8]xgb.ErrorUnmarshaler)
)

func registerEvent(n uint8, fn xgb.EventUnmarshaler) {
	if _, ok := eventFuncs[n]; ok {
		panic("BUG: overlapping event unmarshaler")
	}
	eventFuncs[n] = fn
}

func registerError(n uint8, fn xgb.ErrorUnmarshaler) {
	if _, ok := errorFuncs[n]; ok {
		panic("BUG: overlapping error unmarshaler")
	}
	errorFuncs[n] = fn
}

// Register will query the X server for Test extension support, and register relevant extension unmarshalers with the XConn.
func Register(xconn *xgb.XConn) error {
	// Query the X server for this extension
	reply, err := xproto.QueryExtension(xconn, uint16(len(ExtXName)), ExtXName)
	if err != nil {
		return fmt.Errorf("error querying X for \"Test\": %w", err)
	} else if !reply.Present {
		return fmt.Errorf("no extension named \"Test\" is known to the X server: reply=%+v", reply)
	}

	// Clone event funcs map but set our event no. start index
	extEventFuncs := make(map[uint8]xgb.EventUnmarshaler, len(eventFuncs))
	for n, fn := range eventFuncs {
		extEventFuncs[n+reply.FirstEvent] = fn
	}

	// Clone error funcs map but set our error no. start index
	extErrorFuncs := make(map[uint8]xgb.ErrorUnmarshaler, len(errorFuncs))
	for n, fn := range errorFuncs {
		extErrorFuncs[n+reply.FirstError] = fn
	}

	// Register ourselves with the X server connection
	return xconn.Register(xgb.XExtension{
		XName:       ExtXName,
		MajorOpcode: reply.MajorOpcode,
		EventFuncs:  extEventFuncs,
		ErrorFuncs:  extErrorFuncs,
	})
}

const (
	CursorNone    = 0
	CursorCurrent = 1
)

// Skipping definition for base type 'Bool'

// Skipping definition for base type 'Byte'

// Skipping definition for base type 'Card8'

// Skipping definition for base type 'Char'

// Skipping definition for base type 'Void'

// Skipping definition for base type 'Double'

// Skipping definition for base type 'Float'

// Skipping definition for base type 'Int16'

// Skipping definition for base type 'Int32'

// Skipping definition for base type 'Int8'

// Skipping definition for base type 'Card16'

// Skipping definition for base type 'Card32'

// CompareCursor sends a checked request.
func CompareCursor(c *xgb.XConn, Window xproto.Window, Cursor xproto.Cursor) (CompareCursorReply, error) {
	var reply CompareCursorReply
	op, ok := c.Ext("XTEST")
	if !ok {
		return reply, errors.New("cannot issue request \"CompareCursor\" using the uninitialized extension \"XTEST\". xtest.Register(xconn) must be called first.")
	}
	err := c.SendRecv(compareCursorRequest(op, Window, Cursor), &reply)
	return reply, err
}

// CompareCursorUnchecked sends an unchecked request.
func CompareCursorUnchecked(c *xgb.XConn, Window xproto.Window, Cursor xproto.Cursor) error {
	op, ok := c.Ext("XTEST")
	if !ok {
		return errors.New("cannot issue request \"CompareCursor\" using the uninitialized extension \"XTEST\". xtest.Register(xconn) must be called first.")
	}
	return c.Send(compareCursorRequest(op, Window, Cursor))
}

// CompareCursorReply represents the data returned from a CompareCursor request.
type CompareCursorReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	Same     bool
}

// Unmarshal reads a byte slice into a CompareCursorReply value.
func (v *CompareCursorReply) Unmarshal(buf []byte) error {
	if size := 8; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"CompareCursorReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	v.Same = (buf[b] == 1)
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	return nil
}

// Write request to wire for CompareCursor
// compareCursorRequest writes a CompareCursor request to a byte slice.
func compareCursorRequest(opcode uint8, Window xproto.Window, Cursor xproto.Cursor) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 1 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Cursor))
	b += 4

	return buf
}

// FakeInput sends a checked request.
func FakeInput(c *xgb.XConn, Type byte, Detail byte, Time uint32, Root xproto.Window, RootX int16, RootY int16, Deviceid byte) error {
	op, ok := c.Ext("XTEST")
	if !ok {
		return errors.New("cannot issue request \"FakeInput\" using the uninitialized extension \"XTEST\". xtest.Register(xconn) must be called first.")
	}
	return c.SendRecv(fakeInputRequest(op, Type, Detail, Time, Root, RootX, RootY, Deviceid), nil)
}

// FakeInputUnchecked sends an unchecked request.
func FakeInputUnchecked(c *xgb.XConn, Type byte, Detail byte, Time uint32, Root xproto.Window, RootX int16, RootY int16, Deviceid byte) error {
	op, ok := c.Ext("XTEST")
	if !ok {
		return errors.New("cannot issue request \"FakeInput\" using the uninitialized extension \"XTEST\". xtest.Register(xconn) must be called first.")
	}
	return c.Send(fakeInputRequest(op, Type, Detail, Time, Root, RootX, RootY, Deviceid))
}

// Write request to wire for FakeInput
// fakeInputRequest writes a FakeInput request to a byte slice.
func fakeInputRequest(opcode uint8, Type byte, Detail byte, Time uint32, Root xproto.Window, RootX int16, RootY int16, Deviceid byte) []byte {
	size := 36
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 2 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	buf[b] = Type
	b += 1

	buf[b] = Detail
	b += 1

	b += 2 // padding

	binary.LittleEndian.PutUint32(buf[b:], Time)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Root))
	b += 4

	b += 8 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(RootX))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(RootY))
	b += 2

	b += 7 // padding

	buf[b] = Deviceid
	b += 1

	return buf
}

// GetVersion sends a checked request.
func GetVersion(c *xgb.XConn, MajorVersion byte, MinorVersion uint16) (GetVersionReply, error) {
	var reply GetVersionReply
	op, ok := c.Ext("XTEST")
	if !ok {
		return reply, errors.New("cannot issue request \"GetVersion\" using the uninitialized extension \"XTEST\". xtest.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getVersionRequest(op, MajorVersion, MinorVersion), &reply)
	return reply, err
}

// GetVersionUnchecked sends an unchecked request.
func GetVersionUnchecked(c *xgb.XConn, MajorVersion byte, MinorVersion uint16) error {
	op, ok := c.Ext("XTEST")
	if !ok {
		return errors.New("cannot issue request \"GetVersion\" using the uninitialized extension \"XTEST\". xtest.Register(xconn) must be called first.")
	}
	return c.Send(getVersionRequest(op, MajorVersion, MinorVersion))
}

// GetVersionReply represents the data returned from a GetVersion request.
type GetVersionReply struct {
	Sequence     uint16 // sequence number of the request for this reply
	Length       uint32 // number of bytes in this reply
	MajorVersion byte
	MinorVersion uint16
}

// Unmarshal reads a byte slice into a GetVersionReply value.
func (v *GetVersionReply) Unmarshal(buf []byte) error {
	if size := 10; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetVersionReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	v.MajorVersion = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.MinorVersion = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return nil
}

// Write request to wire for GetVersion
// getVersionRequest writes a GetVersion request to a byte slice.
func getVersionRequest(opcode uint8, MajorVersion byte, MinorVersion uint16) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 0 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	buf[b] = MajorVersion
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], MinorVersion)
	b += 2

	return buf
}

// GrabControl sends a checked request.
func GrabControl(c *xgb.XConn, Impervious bool) error {
	op, ok := c.Ext("XTEST")
	if !ok {
		return errors.New("cannot issue request \"GrabControl\" using the uninitialized extension \"XTEST\". xtest.Register(xconn) must be called first.")
	}
	return c.SendRecv(grabControlRequest(op, Impervious), nil)
}

// GrabControlUnchecked sends an unchecked request.
func GrabControlUnchecked(c *xgb.XConn, Impervious bool) error {
	op, ok := c.Ext("XTEST")
	if !ok {
		return errors.New("cannot issue request \"GrabControl\" using the uninitialized extension \"XTEST\". xtest.Register(xconn) must be called first.")
	}
	return c.Send(grabControlRequest(op, Impervious))
}

// Write request to wire for GrabControl
// grabControlRequest writes a GrabControl request to a byte slice.
func grabControlRequest(opcode uint8, Impervious bool) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 3 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	if Impervious {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	b += 3 // padding

	return buf
}
