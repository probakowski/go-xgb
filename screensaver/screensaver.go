// FILE GENERATED AUTOMATICALLY FROM "screensaver.xml"
package screensaver

import (
	"encoding/binary"
	"errors"
	"fmt"

	"codeberg.org/gruf/go-xgb"
	"codeberg.org/gruf/go-xgb/internal"
	"codeberg.org/gruf/go-xgb/xproto"
)

const (
	// ExtName is the user-friendly name string of this X extension.
	ExtName = "ScreenSaver"

	// ExtXName is the name string this extension is known by to the X server.
	ExtXName = "MIT-SCREEN-SAVER"
)

var (
	// generated index maps of defined event and error numbers -> unmarshalers.
	eventFuncs = map[uint8]xgb.EventUnmarshaler{}
	errorFuncs = map[uint8]xgb.ErrorUnmarshaler{}
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

// Register ...
func Register(xconn *xgb.XConn) error {
	// Query the X server for this extension
	reply, err := xproto.QueryExtension(xconn, uint16(len(ExtXName)), ExtXName)
	if err != nil {
		return fmt.Errorf("error querying X for \"ScreenSaver\": %w", err)
	} else if !reply.Present {
		return fmt.Errorf("no extension named \"ScreenSaver\" is known to the X server: reply=%+v", reply)
	}

	// Clone event funcs map but set our event no. start index
	extEventFuncs := map[uint8]xgb.EventUnmarshaler{}
	for n, fn := range eventFuncs {
		extEventFuncs[n+reply.FirstEvent] = fn
	}

	// Clone error funcs map but set our error no. start index
	extErrorFuncs := map[uint8]xgb.ErrorUnmarshaler{}
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
	EventNotifyMask = 1
	EventCycleMask  = 2
)

const (
	KindBlanked  = 0
	KindInternal = 1
	KindExternal = 2
)

// Notify is the event number for a NotifyEvent.
const Notify = 0

type NotifyEvent struct {
	Sequence uint16
	State    byte
	Time     xproto.Timestamp
	Root     xproto.Window
	Window   xproto.Window
	Kind     byte
	Forced   bool
	// padding: 14 bytes
}

// UnmarshalNotifyEvent constructs a NotifyEvent value that implements xgb.Event from a byte slice.
func UnmarshalNotifyEvent(buf []byte) (xgb.XEvent, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"NotifyEvent\"", len(buf))
	}

	v := NotifyEvent{}
	b := 1 // don't read event number

	v.State = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Time = xproto.Timestamp(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Root = xproto.Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Window = xproto.Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Kind = buf[b]
	b += 1

	v.Forced = (buf[b] == 1)
	b += 1

	b += 14 // padding

	return v, nil
}

// Bytes writes a NotifyEvent value to a byte slice.
func (v NotifyEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 0
	b += 1

	buf[b] = v.State
	b += 1

	b += 2 // skip sequence number

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Time))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Root))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Window))
	b += 4

	buf[b] = v.Kind
	b += 1

	if v.Forced {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	b += 14 // padding

	return buf
}

// SeqID returns the sequence id attached to the Notify event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v NotifyEvent) SeqID() uint16 {
	return v.Sequence
}

func init() {
	registerEvent(0, UnmarshalNotifyEvent)
}

const (
	StateOff      = 0
	StateOn       = 1
	StateCycle    = 2
	StateDisabled = 3
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

// QueryInfo sends a checked request.
func QueryInfo(c *xgb.XConn, Drawable xproto.Drawable) (QueryInfoReply, error) {
	var reply QueryInfoReply
	op, ok := c.Ext("MIT-SCREEN-SAVER")
	if !ok {
		return reply, errors.New("cannot issue request \"QueryInfo\" using the uninitialized extension \"MIT-SCREEN-SAVER\". screensaver.Register(xconn) must be called first.")
	}
	err := c.SendRecv(queryInfoRequest(op, Drawable), &reply)
	return reply, err
}

// QueryInfoUnchecked sends an unchecked request.
func QueryInfoUnchecked(c *xgb.XConn, Drawable xproto.Drawable) error {
	op, ok := c.Ext("MIT-SCREEN-SAVER")
	if !ok {
		return errors.New("cannot issue request \"QueryInfo\" using the uninitialized extension \"MIT-SCREEN-SAVER\". screensaver.Register(xconn) must be called first.")
	}
	return c.Send(queryInfoRequest(op, Drawable))
}

// QueryInfoReply represents the data returned from a QueryInfo request.
type QueryInfoReply struct {
	Sequence         uint16 // sequence number of the request for this reply
	Length           uint32 // number of bytes in this reply
	State            byte
	SaverWindow      xproto.Window
	MsUntilServer    uint32
	MsSinceUserInput uint32
	EventMask        uint32
	Kind             byte
	// padding: 7 bytes
}

// Unmarshal reads a byte slice into a QueryInfoReply value.
func (v *QueryInfoReply) Unmarshal(buf []byte) error {
	if size := 32; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"QueryInfoReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	v.State = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.SaverWindow = xproto.Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.MsUntilServer = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.MsSinceUserInput = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.EventMask = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.Kind = buf[b]
	b += 1

	b += 7 // padding

	return nil
}

// Write request to wire for QueryInfo
// queryInfoRequest writes a QueryInfo request to a byte slice.
func queryInfoRequest(opcode uint8, Drawable xproto.Drawable) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 1 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Drawable))
	b += 4

	return buf
}

// QueryVersion sends a checked request.
func QueryVersion(c *xgb.XConn, ClientMajorVersion byte, ClientMinorVersion byte) (QueryVersionReply, error) {
	var reply QueryVersionReply
	op, ok := c.Ext("MIT-SCREEN-SAVER")
	if !ok {
		return reply, errors.New("cannot issue request \"QueryVersion\" using the uninitialized extension \"MIT-SCREEN-SAVER\". screensaver.Register(xconn) must be called first.")
	}
	err := c.SendRecv(queryVersionRequest(op, ClientMajorVersion, ClientMinorVersion), &reply)
	return reply, err
}

// QueryVersionUnchecked sends an unchecked request.
func QueryVersionUnchecked(c *xgb.XConn, ClientMajorVersion byte, ClientMinorVersion byte) error {
	op, ok := c.Ext("MIT-SCREEN-SAVER")
	if !ok {
		return errors.New("cannot issue request \"QueryVersion\" using the uninitialized extension \"MIT-SCREEN-SAVER\". screensaver.Register(xconn) must be called first.")
	}
	return c.Send(queryVersionRequest(op, ClientMajorVersion, ClientMinorVersion))
}

// QueryVersionReply represents the data returned from a QueryVersion request.
type QueryVersionReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	ServerMajorVersion uint16
	ServerMinorVersion uint16
	// padding: 20 bytes
}

// Unmarshal reads a byte slice into a QueryVersionReply value.
func (v *QueryVersionReply) Unmarshal(buf []byte) error {
	if size := 32; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"QueryVersionReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.ServerMajorVersion = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.ServerMinorVersion = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	b += 20 // padding

	return nil
}

// Write request to wire for QueryVersion
// queryVersionRequest writes a QueryVersion request to a byte slice.
func queryVersionRequest(opcode uint8, ClientMajorVersion byte, ClientMinorVersion byte) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 0 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	buf[b] = ClientMajorVersion
	b += 1

	buf[b] = ClientMinorVersion
	b += 1

	b += 2 // padding

	return buf
}

// SelectInput sends a checked request.
func SelectInput(c *xgb.XConn, Drawable xproto.Drawable, EventMask uint32) error {
	op, ok := c.Ext("MIT-SCREEN-SAVER")
	if !ok {
		return errors.New("cannot issue request \"SelectInput\" using the uninitialized extension \"MIT-SCREEN-SAVER\". screensaver.Register(xconn) must be called first.")
	}
	return c.SendRecv(selectInputRequest(op, Drawable, EventMask), nil)
}

// SelectInputUnchecked sends an unchecked request.
func SelectInputUnchecked(c *xgb.XConn, Drawable xproto.Drawable, EventMask uint32) error {
	op, ok := c.Ext("MIT-SCREEN-SAVER")
	if !ok {
		return errors.New("cannot issue request \"SelectInput\" using the uninitialized extension \"MIT-SCREEN-SAVER\". screensaver.Register(xconn) must be called first.")
	}
	return c.Send(selectInputRequest(op, Drawable, EventMask))
}

// Write request to wire for SelectInput
// selectInputRequest writes a SelectInput request to a byte slice.
func selectInputRequest(opcode uint8, Drawable xproto.Drawable, EventMask uint32) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 2 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Drawable))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], EventMask)
	b += 4

	return buf
}

// SetAttributes sends a checked request.
func SetAttributes(c *xgb.XConn, Drawable xproto.Drawable, X int16, Y int16, Width uint16, Height uint16, BorderWidth uint16, Class byte, Depth byte, Visual xproto.Visualid, ValueMask uint32, ValueList []uint32) error {
	op, ok := c.Ext("MIT-SCREEN-SAVER")
	if !ok {
		return errors.New("cannot issue request \"SetAttributes\" using the uninitialized extension \"MIT-SCREEN-SAVER\". screensaver.Register(xconn) must be called first.")
	}
	return c.SendRecv(setAttributesRequest(op, Drawable, X, Y, Width, Height, BorderWidth, Class, Depth, Visual, ValueMask, ValueList), nil)
}

// SetAttributesUnchecked sends an unchecked request.
func SetAttributesUnchecked(c *xgb.XConn, Drawable xproto.Drawable, X int16, Y int16, Width uint16, Height uint16, BorderWidth uint16, Class byte, Depth byte, Visual xproto.Visualid, ValueMask uint32, ValueList []uint32) error {
	op, ok := c.Ext("MIT-SCREEN-SAVER")
	if !ok {
		return errors.New("cannot issue request \"SetAttributes\" using the uninitialized extension \"MIT-SCREEN-SAVER\". screensaver.Register(xconn) must be called first.")
	}
	return c.Send(setAttributesRequest(op, Drawable, X, Y, Width, Height, BorderWidth, Class, Depth, Visual, ValueMask, ValueList))
}

// Write request to wire for SetAttributes
// setAttributesRequest writes a SetAttributes request to a byte slice.
func setAttributesRequest(opcode uint8, Drawable xproto.Drawable, X int16, Y int16, Width uint16, Height uint16, BorderWidth uint16, Class byte, Depth byte, Visual xproto.Visualid, ValueMask uint32, ValueList []uint32) []byte {
	size := internal.Pad4((28 + internal.Pad4((4 * internal.PopCount(int(ValueMask))))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 3 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Drawable))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(X))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(Y))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Width)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Height)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], BorderWidth)
	b += 2

	buf[b] = Class
	b += 1

	buf[b] = Depth
	b += 1

	binary.LittleEndian.PutUint32(buf[b:], uint32(Visual))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], ValueMask)
	b += 4

	for i := 0; i < len(ValueList); i++ {
		binary.LittleEndian.PutUint32(buf[b:], ValueList[i])
		b += 4
	}
	b = internal.Pad4(b)

	return buf
}

// Suspend sends a checked request.
func Suspend(c *xgb.XConn, Suspend uint32) error {
	op, ok := c.Ext("MIT-SCREEN-SAVER")
	if !ok {
		return errors.New("cannot issue request \"Suspend\" using the uninitialized extension \"MIT-SCREEN-SAVER\". screensaver.Register(xconn) must be called first.")
	}
	return c.SendRecv(suspendRequest(op, Suspend), nil)
}

// SuspendUnchecked sends an unchecked request.
func SuspendUnchecked(c *xgb.XConn, Suspend uint32) error {
	op, ok := c.Ext("MIT-SCREEN-SAVER")
	if !ok {
		return errors.New("cannot issue request \"Suspend\" using the uninitialized extension \"MIT-SCREEN-SAVER\". screensaver.Register(xconn) must be called first.")
	}
	return c.Send(suspendRequest(op, Suspend))
}

// Write request to wire for Suspend
// suspendRequest writes a Suspend request to a byte slice.
func suspendRequest(opcode uint8, Suspend uint32) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 5 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], Suspend)
	b += 4

	return buf
}

// UnsetAttributes sends a checked request.
func UnsetAttributes(c *xgb.XConn, Drawable xproto.Drawable) error {
	op, ok := c.Ext("MIT-SCREEN-SAVER")
	if !ok {
		return errors.New("cannot issue request \"UnsetAttributes\" using the uninitialized extension \"MIT-SCREEN-SAVER\". screensaver.Register(xconn) must be called first.")
	}
	return c.SendRecv(unsetAttributesRequest(op, Drawable), nil)
}

// UnsetAttributesUnchecked sends an unchecked request.
func UnsetAttributesUnchecked(c *xgb.XConn, Drawable xproto.Drawable) error {
	op, ok := c.Ext("MIT-SCREEN-SAVER")
	if !ok {
		return errors.New("cannot issue request \"UnsetAttributes\" using the uninitialized extension \"MIT-SCREEN-SAVER\". screensaver.Register(xconn) must be called first.")
	}
	return c.Send(unsetAttributesRequest(op, Drawable))
}

// Write request to wire for UnsetAttributes
// unsetAttributesRequest writes a UnsetAttributes request to a byte slice.
func unsetAttributesRequest(opcode uint8, Drawable xproto.Drawable) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 4 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Drawable))
	b += 4

	return buf
}
