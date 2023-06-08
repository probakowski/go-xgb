// FILE GENERATED AUTOMATICALLY FROM "shape.xml"
package shape

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
	ExtName = "Shape"

	// ExtXName is the name string this extension is known by to the X server.
	ExtXName = "SHAPE"
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

// Register ...
func Register(xconn *xgb.XConn) error {
	// Query the X server for this extension
	reply, err := xproto.QueryExtension(xconn, uint16(len(ExtXName)), ExtXName)
	if err != nil {
		return fmt.Errorf("error querying X for \"Shape\": %w", err)
	} else if !reply.Present {
		return fmt.Errorf("no extension named \"Shape\" is known to the X server: reply=%+v", reply)
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

type Kind byte

// Notify is the event number for a NotifyEvent.
const Notify = 0

type NotifyEvent struct {
	Sequence       uint16
	ShapeKind      Kind
	AffectedWindow xproto.Window
	ExtentsX       int16
	ExtentsY       int16
	ExtentsWidth   uint16
	ExtentsHeight  uint16
	ServerTime     xproto.Timestamp
	Shaped         bool
	// padding: 11 bytes
}

// UnmarshalNotifyEvent constructs a NotifyEvent value that implements xgb.Event from a byte slice.
func UnmarshalNotifyEvent(buf []byte) (xgb.XEvent, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"NotifyEvent\"", len(buf))
	}

	v := &NotifyEvent{}
	b := 1 // don't read event number

	v.ShapeKind = Kind(buf[b])
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.AffectedWindow = xproto.Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.ExtentsX = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.ExtentsY = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.ExtentsWidth = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.ExtentsHeight = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.ServerTime = xproto.Timestamp(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Shaped = (buf[b] == 1)
	b += 1

	b += 11 // padding

	return v, nil
}

// Bytes writes a NotifyEvent value to a byte slice.
func (v *NotifyEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 0
	b += 1

	buf[b] = uint8(v.ShapeKind)
	b += 1

	b += 2 // skip sequence number

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.AffectedWindow))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.ExtentsX))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.ExtentsY))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.ExtentsWidth)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.ExtentsHeight)
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.ServerTime))
	b += 4

	if v.Shaped {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	b += 11 // padding

	return buf
}

// SeqID returns the sequence id attached to the Notify event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v *NotifyEvent) SeqID() uint16 {
	return v.Sequence
}

func init() { registerEvent(0, UnmarshalNotifyEvent) }

type Op byte

const (
	SkBounding = 0
	SkClip     = 1
	SkInput    = 2
)

const (
	SoSet       = 0
	SoUnion     = 1
	SoIntersect = 2
	SoSubtract  = 3
	SoInvert    = 4
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

// Combine sends a checked request.
func Combine(c *xgb.XConn, Operation Op, DestinationKind Kind, SourceKind Kind, DestinationWindow xproto.Window, XOffset int16, YOffset int16, SourceWindow xproto.Window) error {
	op, ok := c.Ext("SHAPE")
	if !ok {
		return errors.New("cannot issue request \"Combine\" using the uninitialized extension \"SHAPE\". shape.Register(xconn) must be called first.")
	}
	return c.SendRecv(combineRequest(op, Operation, DestinationKind, SourceKind, DestinationWindow, XOffset, YOffset, SourceWindow), nil)
}

// CombineUnchecked sends an unchecked request.
func CombineUnchecked(c *xgb.XConn, Operation Op, DestinationKind Kind, SourceKind Kind, DestinationWindow xproto.Window, XOffset int16, YOffset int16, SourceWindow xproto.Window) error {
	op, ok := c.Ext("SHAPE")
	if !ok {
		return errors.New("cannot issue request \"Combine\" using the uninitialized extension \"SHAPE\". shape.Register(xconn) must be called first.")
	}
	return c.Send(combineRequest(op, Operation, DestinationKind, SourceKind, DestinationWindow, XOffset, YOffset, SourceWindow))
}

// Write request to wire for Combine
// combineRequest writes a Combine request to a byte slice.
func combineRequest(opcode uint8, Operation Op, DestinationKind Kind, SourceKind Kind, DestinationWindow xproto.Window, XOffset int16, YOffset int16, SourceWindow xproto.Window) []byte {
	size := 20
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 3 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	buf[b] = uint8(Operation)
	b += 1

	buf[b] = uint8(DestinationKind)
	b += 1

	buf[b] = uint8(SourceKind)
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint32(buf[b:], uint32(DestinationWindow))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(XOffset))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(YOffset))
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(SourceWindow))
	b += 4

	return buf
}

// GetRectangles sends a checked request.
func GetRectangles(c *xgb.XConn, Window xproto.Window, SourceKind Kind) (GetRectanglesReply, error) {
	var reply GetRectanglesReply
	op, ok := c.Ext("SHAPE")
	if !ok {
		return reply, errors.New("cannot issue request \"GetRectangles\" using the uninitialized extension \"SHAPE\". shape.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getRectanglesRequest(op, Window, SourceKind), &reply)
	return reply, err
}

// GetRectanglesUnchecked sends an unchecked request.
func GetRectanglesUnchecked(c *xgb.XConn, Window xproto.Window, SourceKind Kind) error {
	op, ok := c.Ext("SHAPE")
	if !ok {
		return errors.New("cannot issue request \"GetRectangles\" using the uninitialized extension \"SHAPE\". shape.Register(xconn) must be called first.")
	}
	return c.Send(getRectanglesRequest(op, Window, SourceKind))
}

// GetRectanglesReply represents the data returned from a GetRectangles request.
type GetRectanglesReply struct {
	Sequence      uint16 // sequence number of the request for this reply
	Length        uint32 // number of bytes in this reply
	Ordering      byte
	RectanglesLen uint32
	// padding: 20 bytes
	Rectangles []xproto.Rectangle // size: internal.Pad4((int(RectanglesLen) * 8))
}

// Unmarshal reads a byte slice into a GetRectanglesReply value.
func (v *GetRectanglesReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.RectanglesLen) * 8))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetRectanglesReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	v.Ordering = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.RectanglesLen = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 20 // padding

	v.Rectangles = make([]xproto.Rectangle, v.RectanglesLen)
	b += xproto.RectangleReadList(buf[b:], v.Rectangles)

	return nil
}

// Write request to wire for GetRectangles
// getRectanglesRequest writes a GetRectangles request to a byte slice.
func getRectanglesRequest(opcode uint8, Window xproto.Window, SourceKind Kind) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 8 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	buf[b] = uint8(SourceKind)
	b += 1

	b += 3 // padding

	return buf
}

// InputSelected sends a checked request.
func InputSelected(c *xgb.XConn, DestinationWindow xproto.Window) (InputSelectedReply, error) {
	var reply InputSelectedReply
	op, ok := c.Ext("SHAPE")
	if !ok {
		return reply, errors.New("cannot issue request \"InputSelected\" using the uninitialized extension \"SHAPE\". shape.Register(xconn) must be called first.")
	}
	err := c.SendRecv(inputSelectedRequest(op, DestinationWindow), &reply)
	return reply, err
}

// InputSelectedUnchecked sends an unchecked request.
func InputSelectedUnchecked(c *xgb.XConn, DestinationWindow xproto.Window) error {
	op, ok := c.Ext("SHAPE")
	if !ok {
		return errors.New("cannot issue request \"InputSelected\" using the uninitialized extension \"SHAPE\". shape.Register(xconn) must be called first.")
	}
	return c.Send(inputSelectedRequest(op, DestinationWindow))
}

// InputSelectedReply represents the data returned from a InputSelected request.
type InputSelectedReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	Enabled  bool
}

// Unmarshal reads a byte slice into a InputSelectedReply value.
func (v *InputSelectedReply) Unmarshal(buf []byte) error {
	if size := 8; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"InputSelectedReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	v.Enabled = (buf[b] == 1)
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	return nil
}

// Write request to wire for InputSelected
// inputSelectedRequest writes a InputSelected request to a byte slice.
func inputSelectedRequest(opcode uint8, DestinationWindow xproto.Window) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 7 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(DestinationWindow))
	b += 4

	return buf
}

// Mask sends a checked request.
func Mask(c *xgb.XConn, Operation Op, DestinationKind Kind, DestinationWindow xproto.Window, XOffset int16, YOffset int16, SourceBitmap xproto.Pixmap) error {
	op, ok := c.Ext("SHAPE")
	if !ok {
		return errors.New("cannot issue request \"Mask\" using the uninitialized extension \"SHAPE\". shape.Register(xconn) must be called first.")
	}
	return c.SendRecv(maskRequest(op, Operation, DestinationKind, DestinationWindow, XOffset, YOffset, SourceBitmap), nil)
}

// MaskUnchecked sends an unchecked request.
func MaskUnchecked(c *xgb.XConn, Operation Op, DestinationKind Kind, DestinationWindow xproto.Window, XOffset int16, YOffset int16, SourceBitmap xproto.Pixmap) error {
	op, ok := c.Ext("SHAPE")
	if !ok {
		return errors.New("cannot issue request \"Mask\" using the uninitialized extension \"SHAPE\". shape.Register(xconn) must be called first.")
	}
	return c.Send(maskRequest(op, Operation, DestinationKind, DestinationWindow, XOffset, YOffset, SourceBitmap))
}

// Write request to wire for Mask
// maskRequest writes a Mask request to a byte slice.
func maskRequest(opcode uint8, Operation Op, DestinationKind Kind, DestinationWindow xproto.Window, XOffset int16, YOffset int16, SourceBitmap xproto.Pixmap) []byte {
	size := 20
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 2 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	buf[b] = uint8(Operation)
	b += 1

	buf[b] = uint8(DestinationKind)
	b += 1

	b += 2 // padding

	binary.LittleEndian.PutUint32(buf[b:], uint32(DestinationWindow))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(XOffset))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(YOffset))
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(SourceBitmap))
	b += 4

	return buf
}

// Offset sends a checked request.
func Offset(c *xgb.XConn, DestinationKind Kind, DestinationWindow xproto.Window, XOffset int16, YOffset int16) error {
	op, ok := c.Ext("SHAPE")
	if !ok {
		return errors.New("cannot issue request \"Offset\" using the uninitialized extension \"SHAPE\". shape.Register(xconn) must be called first.")
	}
	return c.SendRecv(offsetRequest(op, DestinationKind, DestinationWindow, XOffset, YOffset), nil)
}

// OffsetUnchecked sends an unchecked request.
func OffsetUnchecked(c *xgb.XConn, DestinationKind Kind, DestinationWindow xproto.Window, XOffset int16, YOffset int16) error {
	op, ok := c.Ext("SHAPE")
	if !ok {
		return errors.New("cannot issue request \"Offset\" using the uninitialized extension \"SHAPE\". shape.Register(xconn) must be called first.")
	}
	return c.Send(offsetRequest(op, DestinationKind, DestinationWindow, XOffset, YOffset))
}

// Write request to wire for Offset
// offsetRequest writes a Offset request to a byte slice.
func offsetRequest(opcode uint8, DestinationKind Kind, DestinationWindow xproto.Window, XOffset int16, YOffset int16) []byte {
	size := 16
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 4 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	buf[b] = uint8(DestinationKind)
	b += 1

	b += 3 // padding

	binary.LittleEndian.PutUint32(buf[b:], uint32(DestinationWindow))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(XOffset))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(YOffset))
	b += 2

	return buf
}

// QueryExtents sends a checked request.
func QueryExtents(c *xgb.XConn, DestinationWindow xproto.Window) (QueryExtentsReply, error) {
	var reply QueryExtentsReply
	op, ok := c.Ext("SHAPE")
	if !ok {
		return reply, errors.New("cannot issue request \"QueryExtents\" using the uninitialized extension \"SHAPE\". shape.Register(xconn) must be called first.")
	}
	err := c.SendRecv(queryExtentsRequest(op, DestinationWindow), &reply)
	return reply, err
}

// QueryExtentsUnchecked sends an unchecked request.
func QueryExtentsUnchecked(c *xgb.XConn, DestinationWindow xproto.Window) error {
	op, ok := c.Ext("SHAPE")
	if !ok {
		return errors.New("cannot issue request \"QueryExtents\" using the uninitialized extension \"SHAPE\". shape.Register(xconn) must be called first.")
	}
	return c.Send(queryExtentsRequest(op, DestinationWindow))
}

// QueryExtentsReply represents the data returned from a QueryExtents request.
type QueryExtentsReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	BoundingShaped bool
	ClipShaped     bool
	// padding: 2 bytes
	BoundingShapeExtentsX      int16
	BoundingShapeExtentsY      int16
	BoundingShapeExtentsWidth  uint16
	BoundingShapeExtentsHeight uint16
	ClipShapeExtentsX          int16
	ClipShapeExtentsY          int16
	ClipShapeExtentsWidth      uint16
	ClipShapeExtentsHeight     uint16
}

// Unmarshal reads a byte slice into a QueryExtentsReply value.
func (v *QueryExtentsReply) Unmarshal(buf []byte) error {
	if size := 28; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"QueryExtentsReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.BoundingShaped = (buf[b] == 1)
	b += 1

	v.ClipShaped = (buf[b] == 1)
	b += 1

	b += 2 // padding

	v.BoundingShapeExtentsX = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.BoundingShapeExtentsY = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.BoundingShapeExtentsWidth = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.BoundingShapeExtentsHeight = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.ClipShapeExtentsX = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.ClipShapeExtentsY = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.ClipShapeExtentsWidth = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.ClipShapeExtentsHeight = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return nil
}

// Write request to wire for QueryExtents
// queryExtentsRequest writes a QueryExtents request to a byte slice.
func queryExtentsRequest(opcode uint8, DestinationWindow xproto.Window) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 5 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(DestinationWindow))
	b += 4

	return buf
}

// QueryVersion sends a checked request.
func QueryVersion(c *xgb.XConn) (QueryVersionReply, error) {
	var reply QueryVersionReply
	op, ok := c.Ext("SHAPE")
	if !ok {
		return reply, errors.New("cannot issue request \"QueryVersion\" using the uninitialized extension \"SHAPE\". shape.Register(xconn) must be called first.")
	}
	err := c.SendRecv(queryVersionRequest(op), &reply)
	return reply, err
}

// QueryVersionUnchecked sends an unchecked request.
func QueryVersionUnchecked(c *xgb.XConn) error {
	op, ok := c.Ext("SHAPE")
	if !ok {
		return errors.New("cannot issue request \"QueryVersion\" using the uninitialized extension \"SHAPE\". shape.Register(xconn) must be called first.")
	}
	return c.Send(queryVersionRequest(op))
}

// QueryVersionReply represents the data returned from a QueryVersion request.
type QueryVersionReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	MajorVersion uint16
	MinorVersion uint16
}

// Unmarshal reads a byte slice into a QueryVersionReply value.
func (v *QueryVersionReply) Unmarshal(buf []byte) error {
	if size := 12; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"QueryVersionReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.MajorVersion = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.MinorVersion = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return nil
}

// Write request to wire for QueryVersion
// queryVersionRequest writes a QueryVersion request to a byte slice.
func queryVersionRequest(opcode uint8) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 0 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// Rectangles sends a checked request.
func Rectangles(c *xgb.XConn, Operation Op, DestinationKind Kind, Ordering byte, DestinationWindow xproto.Window, XOffset int16, YOffset int16, Rectangles []xproto.Rectangle) error {
	op, ok := c.Ext("SHAPE")
	if !ok {
		return errors.New("cannot issue request \"Rectangles\" using the uninitialized extension \"SHAPE\". shape.Register(xconn) must be called first.")
	}
	return c.SendRecv(rectanglesRequest(op, Operation, DestinationKind, Ordering, DestinationWindow, XOffset, YOffset, Rectangles), nil)
}

// RectanglesUnchecked sends an unchecked request.
func RectanglesUnchecked(c *xgb.XConn, Operation Op, DestinationKind Kind, Ordering byte, DestinationWindow xproto.Window, XOffset int16, YOffset int16, Rectangles []xproto.Rectangle) error {
	op, ok := c.Ext("SHAPE")
	if !ok {
		return errors.New("cannot issue request \"Rectangles\" using the uninitialized extension \"SHAPE\". shape.Register(xconn) must be called first.")
	}
	return c.Send(rectanglesRequest(op, Operation, DestinationKind, Ordering, DestinationWindow, XOffset, YOffset, Rectangles))
}

// Write request to wire for Rectangles
// rectanglesRequest writes a Rectangles request to a byte slice.
func rectanglesRequest(opcode uint8, Operation Op, DestinationKind Kind, Ordering byte, DestinationWindow xproto.Window, XOffset int16, YOffset int16, Rectangles []xproto.Rectangle) []byte {
	size := internal.Pad4((16 + internal.Pad4((len(Rectangles) * 8))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 1 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	buf[b] = uint8(Operation)
	b += 1

	buf[b] = uint8(DestinationKind)
	b += 1

	buf[b] = Ordering
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint32(buf[b:], uint32(DestinationWindow))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(XOffset))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(YOffset))
	b += 2

	b += xproto.RectangleListBytes(buf[b:], Rectangles)

	return buf
}

// SelectInput sends a checked request.
func SelectInput(c *xgb.XConn, DestinationWindow xproto.Window, Enable bool) error {
	op, ok := c.Ext("SHAPE")
	if !ok {
		return errors.New("cannot issue request \"SelectInput\" using the uninitialized extension \"SHAPE\". shape.Register(xconn) must be called first.")
	}
	return c.SendRecv(selectInputRequest(op, DestinationWindow, Enable), nil)
}

// SelectInputUnchecked sends an unchecked request.
func SelectInputUnchecked(c *xgb.XConn, DestinationWindow xproto.Window, Enable bool) error {
	op, ok := c.Ext("SHAPE")
	if !ok {
		return errors.New("cannot issue request \"SelectInput\" using the uninitialized extension \"SHAPE\". shape.Register(xconn) must be called first.")
	}
	return c.Send(selectInputRequest(op, DestinationWindow, Enable))
}

// Write request to wire for SelectInput
// selectInputRequest writes a SelectInput request to a byte slice.
func selectInputRequest(opcode uint8, DestinationWindow xproto.Window, Enable bool) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 6 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(DestinationWindow))
	b += 4

	if Enable {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	b += 3 // padding

	return buf
}
