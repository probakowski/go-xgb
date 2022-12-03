// FILE GENERATED AUTOMATICALLY FROM "shm.xml"
package shm

import (
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"codeberg.org/gruf/go-xgb"
	"codeberg.org/gruf/go-xgb/xproto"
)

const (
	// ExtName is the user-friendly name string of this X extension.
	ExtName = "Shm"

	// ExtXName is the name string this extension is known by to the X server.
	ExtXName = "MIT-SHM"
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
		return fmt.Errorf("error querying X for \"Shm\": %w", err)
	} else if !reply.Present {
		return fmt.Errorf("no extension named \"Shm\" is known to the X server: reply=%+v", reply)
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

// BadBadSeg is the error number for a BadBadSeg.
const BadBadSeg = 0

type BadSegError xproto.ValueError

// BadSegErrorNew constructs a BadSegError value that implements xgb.Error from a byte slice.
func UnmarshalBadSegError(buf []byte) (xgb.XError, error) {
	return xproto.UnmarshalValueError(buf)
}

// SequenceId returns the sequence id attached to the BadBadSeg error.
// This is mostly used internally.
func (err BadSegError) SeqID() uint16 {
	return err.Sequence
}

// BadId returns the 'BadValue' number if one exists for the BadBadSeg error. If no bad value exists, 0 is returned.
func (err BadSegError) BadID() uint32 {
	return 0
}

// Error returns a rudimentary string representation of the BadBadSeg error.
func (err BadSegError) Error() string {
	var buf strings.Builder

	buf.WriteString("BadBadSeg{")
	buf.WriteString("NiceName: " + err.NiceName)
	buf.WriteByte(' ')
	buf.WriteString("Sequence: " + strconv.FormatUint(uint64(err.Sequence), 10))
	buf.WriteByte(' ')

	fmt.Fprintf(&buf, "BadValue: %d", err.BadValue)
	buf.WriteString(", ")
	fmt.Fprintf(&buf, "MinorOpcode: %d", err.MinorOpcode)
	buf.WriteString(", ")
	fmt.Fprintf(&buf, "MajorOpcode: %d", err.MajorOpcode)
	buf.WriteString(", ")
	buf.WriteByte('}')

	return buf.String()
}

func init() {
	registerError(0, UnmarshalBadSegError)
}

// Completion is the event number for a CompletionEvent.
const Completion = 0

type CompletionEvent struct {
	Sequence uint16
	// padding: 1 bytes
	Drawable   xproto.Drawable
	MinorEvent uint16
	MajorEvent byte
	// padding: 1 bytes
	Shmseg Seg
	Offset uint32
}

// UnmarshalCompletionEvent constructs a CompletionEvent value that implements xgb.Event from a byte slice.
func UnmarshalCompletionEvent(buf []byte) (xgb.XEvent, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"CompletionEvent\"", len(buf))
	}

	v := CompletionEvent{}
	b := 1 // don't read event number

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Drawable = xproto.Drawable(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.MinorEvent = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.MajorEvent = buf[b]
	b += 1

	b += 1 // padding

	v.Shmseg = Seg(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Offset = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	return v, nil
}

// Bytes writes a CompletionEvent value to a byte slice.
func (v CompletionEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 0
	b += 1

	b += 1 // padding

	b += 2 // skip sequence number

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Drawable))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], v.MinorEvent)
	b += 2

	buf[b] = v.MajorEvent
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Shmseg))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], v.Offset)
	b += 4

	return buf
}

// SeqID returns the sequence id attached to the Completion event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v CompletionEvent) SeqID() uint16 {
	return v.Sequence
}

func init() { registerEvent(0, UnmarshalCompletionEvent) }

type Seg uint32

func NewSegID(c *xgb.XConn) (Seg, error) {
	id, err := c.NewXID()
	return Seg(id), err
}

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

// Attach sends a checked request.
func Attach(c *xgb.XConn, Shmseg Seg, Shmid uint32, ReadOnly bool) error {
	op, ok := c.Ext("MIT-SHM")
	if !ok {
		return errors.New("cannot issue request \"Attach\" using the uninitialized extension \"MIT-SHM\". shm.Register(xconn) must be called first.")
	}
	return c.SendRecv(attachRequest(op, Shmseg, Shmid, ReadOnly), nil)
}

// AttachUnchecked sends an unchecked request.
func AttachUnchecked(c *xgb.XConn, Shmseg Seg, Shmid uint32, ReadOnly bool) error {
	op, ok := c.Ext("MIT-SHM")
	if !ok {
		return errors.New("cannot issue request \"Attach\" using the uninitialized extension \"MIT-SHM\". shm.Register(xconn) must be called first.")
	}
	return c.Send(attachRequest(op, Shmseg, Shmid, ReadOnly))
}

// Write request to wire for Attach
// attachRequest writes a Attach request to a byte slice.
func attachRequest(opcode uint8, Shmseg Seg, Shmid uint32, ReadOnly bool) []byte {
	size := 16
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 1 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Shmseg))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], Shmid)
	b += 4

	if ReadOnly {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	b += 3 // padding

	return buf
}

// AttachFd sends a checked request.
func AttachFd(c *xgb.XConn, Shmseg Seg, ReadOnly bool) error {
	op, ok := c.Ext("MIT-SHM")
	if !ok {
		return errors.New("cannot issue request \"AttachFd\" using the uninitialized extension \"MIT-SHM\". shm.Register(xconn) must be called first.")
	}
	return c.SendRecv(attachFdRequest(op, Shmseg, ReadOnly), nil)
}

// AttachFdUnchecked sends an unchecked request.
func AttachFdUnchecked(c *xgb.XConn, Shmseg Seg, ReadOnly bool) error {
	op, ok := c.Ext("MIT-SHM")
	if !ok {
		return errors.New("cannot issue request \"AttachFd\" using the uninitialized extension \"MIT-SHM\". shm.Register(xconn) must be called first.")
	}
	return c.Send(attachFdRequest(op, Shmseg, ReadOnly))
}

// Write request to wire for AttachFd
// attachFdRequest writes a AttachFd request to a byte slice.
func attachFdRequest(opcode uint8, Shmseg Seg, ReadOnly bool) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 6 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Shmseg))
	b += 4

	if ReadOnly {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	b += 3 // padding

	return buf
}

// CreatePixmap sends a checked request.
func CreatePixmap(c *xgb.XConn, Pid xproto.Pixmap, Drawable xproto.Drawable, Width uint16, Height uint16, Depth byte, Shmseg Seg, Offset uint32) error {
	op, ok := c.Ext("MIT-SHM")
	if !ok {
		return errors.New("cannot issue request \"CreatePixmap\" using the uninitialized extension \"MIT-SHM\". shm.Register(xconn) must be called first.")
	}
	return c.SendRecv(createPixmapRequest(op, Pid, Drawable, Width, Height, Depth, Shmseg, Offset), nil)
}

// CreatePixmapUnchecked sends an unchecked request.
func CreatePixmapUnchecked(c *xgb.XConn, Pid xproto.Pixmap, Drawable xproto.Drawable, Width uint16, Height uint16, Depth byte, Shmseg Seg, Offset uint32) error {
	op, ok := c.Ext("MIT-SHM")
	if !ok {
		return errors.New("cannot issue request \"CreatePixmap\" using the uninitialized extension \"MIT-SHM\". shm.Register(xconn) must be called first.")
	}
	return c.Send(createPixmapRequest(op, Pid, Drawable, Width, Height, Depth, Shmseg, Offset))
}

// Write request to wire for CreatePixmap
// createPixmapRequest writes a CreatePixmap request to a byte slice.
func createPixmapRequest(opcode uint8, Pid xproto.Pixmap, Drawable xproto.Drawable, Width uint16, Height uint16, Depth byte, Shmseg Seg, Offset uint32) []byte {
	size := 28
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 5 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Pid))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Drawable))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], Width)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Height)
	b += 2

	buf[b] = Depth
	b += 1

	b += 3 // padding

	binary.LittleEndian.PutUint32(buf[b:], uint32(Shmseg))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], Offset)
	b += 4

	return buf
}

// CreateSegment sends a checked request.
func CreateSegment(c *xgb.XConn, Shmseg Seg, Size uint32, ReadOnly bool) (CreateSegmentReply, error) {
	var reply CreateSegmentReply
	op, ok := c.Ext("MIT-SHM")
	if !ok {
		return reply, errors.New("cannot issue request \"CreateSegment\" using the uninitialized extension \"MIT-SHM\". shm.Register(xconn) must be called first.")
	}
	err := c.SendRecv(createSegmentRequest(op, Shmseg, Size, ReadOnly), &reply)
	return reply, err
}

// CreateSegmentUnchecked sends an unchecked request.
func CreateSegmentUnchecked(c *xgb.XConn, Shmseg Seg, Size uint32, ReadOnly bool) error {
	op, ok := c.Ext("MIT-SHM")
	if !ok {
		return errors.New("cannot issue request \"CreateSegment\" using the uninitialized extension \"MIT-SHM\". shm.Register(xconn) must be called first.")
	}
	return c.Send(createSegmentRequest(op, Shmseg, Size, ReadOnly))
}

// CreateSegmentReply represents the data returned from a CreateSegment request.
type CreateSegmentReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	Nfd      byte
	// padding: 24 bytes
}

// Unmarshal reads a byte slice into a CreateSegmentReply value.
func (v *CreateSegmentReply) Unmarshal(buf []byte) error {
	if size := 32; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"CreateSegmentReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	v.Nfd = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	b += 24 // padding

	return nil
}

// Write request to wire for CreateSegment
// createSegmentRequest writes a CreateSegment request to a byte slice.
func createSegmentRequest(opcode uint8, Shmseg Seg, Size uint32, ReadOnly bool) []byte {
	size := 16
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 7 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Shmseg))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], Size)
	b += 4

	if ReadOnly {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	b += 3 // padding

	return buf
}

// Detach sends a checked request.
func Detach(c *xgb.XConn, Shmseg Seg) error {
	op, ok := c.Ext("MIT-SHM")
	if !ok {
		return errors.New("cannot issue request \"Detach\" using the uninitialized extension \"MIT-SHM\". shm.Register(xconn) must be called first.")
	}
	return c.SendRecv(detachRequest(op, Shmseg), nil)
}

// DetachUnchecked sends an unchecked request.
func DetachUnchecked(c *xgb.XConn, Shmseg Seg) error {
	op, ok := c.Ext("MIT-SHM")
	if !ok {
		return errors.New("cannot issue request \"Detach\" using the uninitialized extension \"MIT-SHM\". shm.Register(xconn) must be called first.")
	}
	return c.Send(detachRequest(op, Shmseg))
}

// Write request to wire for Detach
// detachRequest writes a Detach request to a byte slice.
func detachRequest(opcode uint8, Shmseg Seg) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 2 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Shmseg))
	b += 4

	return buf
}

// GetImage sends a checked request.
func GetImage(c *xgb.XConn, Drawable xproto.Drawable, X int16, Y int16, Width uint16, Height uint16, PlaneMask uint32, Format byte, Shmseg Seg, Offset uint32) (GetImageReply, error) {
	var reply GetImageReply
	op, ok := c.Ext("MIT-SHM")
	if !ok {
		return reply, errors.New("cannot issue request \"GetImage\" using the uninitialized extension \"MIT-SHM\". shm.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getImageRequest(op, Drawable, X, Y, Width, Height, PlaneMask, Format, Shmseg, Offset), &reply)
	return reply, err
}

// GetImageUnchecked sends an unchecked request.
func GetImageUnchecked(c *xgb.XConn, Drawable xproto.Drawable, X int16, Y int16, Width uint16, Height uint16, PlaneMask uint32, Format byte, Shmseg Seg, Offset uint32) error {
	op, ok := c.Ext("MIT-SHM")
	if !ok {
		return errors.New("cannot issue request \"GetImage\" using the uninitialized extension \"MIT-SHM\". shm.Register(xconn) must be called first.")
	}
	return c.Send(getImageRequest(op, Drawable, X, Y, Width, Height, PlaneMask, Format, Shmseg, Offset))
}

// GetImageReply represents the data returned from a GetImage request.
type GetImageReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	Depth    byte
	Visual   xproto.Visualid
	Size     uint32
}

// Unmarshal reads a byte slice into a GetImageReply value.
func (v *GetImageReply) Unmarshal(buf []byte) error {
	if size := 16; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetImageReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	v.Depth = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Visual = xproto.Visualid(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Size = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	return nil
}

// Write request to wire for GetImage
// getImageRequest writes a GetImage request to a byte slice.
func getImageRequest(opcode uint8, Drawable xproto.Drawable, X int16, Y int16, Width uint16, Height uint16, PlaneMask uint32, Format byte, Shmseg Seg, Offset uint32) []byte {
	size := 32
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

	binary.LittleEndian.PutUint16(buf[b:], uint16(X))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(Y))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Width)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Height)
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], PlaneMask)
	b += 4

	buf[b] = Format
	b += 1

	b += 3 // padding

	binary.LittleEndian.PutUint32(buf[b:], uint32(Shmseg))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], Offset)
	b += 4

	return buf
}

// PutImage sends a checked request.
func PutImage(c *xgb.XConn, Drawable xproto.Drawable, Gc xproto.Gcontext, TotalWidth uint16, TotalHeight uint16, SrcX uint16, SrcY uint16, SrcWidth uint16, SrcHeight uint16, DstX int16, DstY int16, Depth byte, Format byte, SendEvent bool, Shmseg Seg, Offset uint32) error {
	op, ok := c.Ext("MIT-SHM")
	if !ok {
		return errors.New("cannot issue request \"PutImage\" using the uninitialized extension \"MIT-SHM\". shm.Register(xconn) must be called first.")
	}
	return c.SendRecv(putImageRequest(op, Drawable, Gc, TotalWidth, TotalHeight, SrcX, SrcY, SrcWidth, SrcHeight, DstX, DstY, Depth, Format, SendEvent, Shmseg, Offset), nil)
}

// PutImageUnchecked sends an unchecked request.
func PutImageUnchecked(c *xgb.XConn, Drawable xproto.Drawable, Gc xproto.Gcontext, TotalWidth uint16, TotalHeight uint16, SrcX uint16, SrcY uint16, SrcWidth uint16, SrcHeight uint16, DstX int16, DstY int16, Depth byte, Format byte, SendEvent bool, Shmseg Seg, Offset uint32) error {
	op, ok := c.Ext("MIT-SHM")
	if !ok {
		return errors.New("cannot issue request \"PutImage\" using the uninitialized extension \"MIT-SHM\". shm.Register(xconn) must be called first.")
	}
	return c.Send(putImageRequest(op, Drawable, Gc, TotalWidth, TotalHeight, SrcX, SrcY, SrcWidth, SrcHeight, DstX, DstY, Depth, Format, SendEvent, Shmseg, Offset))
}

// Write request to wire for PutImage
// putImageRequest writes a PutImage request to a byte slice.
func putImageRequest(opcode uint8, Drawable xproto.Drawable, Gc xproto.Gcontext, TotalWidth uint16, TotalHeight uint16, SrcX uint16, SrcY uint16, SrcWidth uint16, SrcHeight uint16, DstX int16, DstY int16, Depth byte, Format byte, SendEvent bool, Shmseg Seg, Offset uint32) []byte {
	size := 40
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

	binary.LittleEndian.PutUint32(buf[b:], uint32(Gc))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], TotalWidth)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], TotalHeight)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], SrcX)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], SrcY)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], SrcWidth)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], SrcHeight)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(DstX))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(DstY))
	b += 2

	buf[b] = Depth
	b += 1

	buf[b] = Format
	b += 1

	if SendEvent {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint32(buf[b:], uint32(Shmseg))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], Offset)
	b += 4

	return buf
}

// QueryVersion sends a checked request.
func QueryVersion(c *xgb.XConn) (QueryVersionReply, error) {
	var reply QueryVersionReply
	op, ok := c.Ext("MIT-SHM")
	if !ok {
		return reply, errors.New("cannot issue request \"QueryVersion\" using the uninitialized extension \"MIT-SHM\". shm.Register(xconn) must be called first.")
	}
	err := c.SendRecv(queryVersionRequest(op), &reply)
	return reply, err
}

// QueryVersionUnchecked sends an unchecked request.
func QueryVersionUnchecked(c *xgb.XConn) error {
	op, ok := c.Ext("MIT-SHM")
	if !ok {
		return errors.New("cannot issue request \"QueryVersion\" using the uninitialized extension \"MIT-SHM\". shm.Register(xconn) must be called first.")
	}
	return c.Send(queryVersionRequest(op))
}

// QueryVersionReply represents the data returned from a QueryVersion request.
type QueryVersionReply struct {
	Sequence      uint16 // sequence number of the request for this reply
	Length        uint32 // number of bytes in this reply
	SharedPixmaps bool
	MajorVersion  uint16
	MinorVersion  uint16
	Uid           uint16
	Gid           uint16
	PixmapFormat  byte
	// padding: 15 bytes
}

// Unmarshal reads a byte slice into a QueryVersionReply value.
func (v *QueryVersionReply) Unmarshal(buf []byte) error {
	if size := 32; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"QueryVersionReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	v.SharedPixmaps = (buf[b] == 1)
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.MajorVersion = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.MinorVersion = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Uid = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Gid = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.PixmapFormat = buf[b]
	b += 1

	b += 15 // padding

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
