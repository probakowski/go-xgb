// FILE GENERATED AUTOMATICALLY FROM "dri2.xml"
package dri2

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
	ExtName = "DRI2"

	// ExtXName is the name string this extension is known by to the X server.
	ExtXName = "DRI2"
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
		return fmt.Errorf("error querying X for \"DRI2\": %w", err)
	} else if !reply.Present {
		return fmt.Errorf("no extension named \"DRI2\" is known to the X server: reply=%+v", reply)
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

type AttachFormat struct {
	Attachment uint32
	Format     uint32
}

// AttachFormatRead reads a byte slice into a AttachFormat value.
func AttachFormatRead(buf []byte, v *AttachFormat) int {
	b := 0

	v.Attachment = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.Format = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	return b
}

// AttachFormatReadList reads a byte slice into a list of AttachFormat values.
func AttachFormatReadList(buf []byte, dest []AttachFormat) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = AttachFormat{}
		b += AttachFormatRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a AttachFormat value to a byte slice.
func (v AttachFormat) Bytes() []byte {
	buf := make([]byte, 8)
	b := 0

	binary.LittleEndian.PutUint32(buf[b:], v.Attachment)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], v.Format)
	b += 4

	return buf[:b]
}

// AttachFormatListBytes writes a list of AttachFormat values to a byte slice.
func AttachFormatListBytes(buf []byte, list []AttachFormat) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

const (
	AttachmentBufferFrontLeft      = 0
	AttachmentBufferBackLeft       = 1
	AttachmentBufferFrontRight     = 2
	AttachmentBufferBackRight      = 3
	AttachmentBufferDepth          = 4
	AttachmentBufferStencil        = 5
	AttachmentBufferAccum          = 6
	AttachmentBufferFakeFrontLeft  = 7
	AttachmentBufferFakeFrontRight = 8
	AttachmentBufferDepthStencil   = 9
	AttachmentBufferHiz            = 10
)

// BufferSwapComplete is the event number for a BufferSwapCompleteEvent.
const BufferSwapComplete = 0

type BufferSwapCompleteEvent struct {
	Sequence uint16
	// padding: 1 bytes
	EventType uint16
	// padding: 2 bytes
	Drawable xproto.Drawable
	UstHi    uint32
	UstLo    uint32
	MscHi    uint32
	MscLo    uint32
	Sbc      uint32
}

// UnmarshalBufferSwapCompleteEvent constructs a BufferSwapCompleteEvent value that implements xgb.Event from a byte slice.
func UnmarshalBufferSwapCompleteEvent(buf []byte) (xgb.XEvent, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"BufferSwapCompleteEvent\"", len(buf))
	}

	v := BufferSwapCompleteEvent{}
	b := 1 // don't read event number

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.EventType = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	b += 2 // padding

	v.Drawable = xproto.Drawable(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.UstHi = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.UstLo = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.MscHi = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.MscLo = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.Sbc = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	return v, nil
}

// Bytes writes a BufferSwapCompleteEvent value to a byte slice.
func (v BufferSwapCompleteEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 0
	b += 1

	b += 1 // padding

	b += 2 // skip sequence number

	binary.LittleEndian.PutUint16(buf[b:], v.EventType)
	b += 2

	b += 2 // padding

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Drawable))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], v.UstHi)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], v.UstLo)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], v.MscHi)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], v.MscLo)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], v.Sbc)
	b += 4

	return buf
}

// SeqID returns the sequence id attached to the BufferSwapComplete event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v BufferSwapCompleteEvent) SeqID() uint16 {
	return v.Sequence
}

func init() { registerEvent(0, UnmarshalBufferSwapCompleteEvent) }

type DRI2Buffer struct {
	Attachment uint32
	Name       uint32
	Pitch      uint32
	Cpp        uint32
	Flags      uint32
}

// DRI2BufferRead reads a byte slice into a DRI2Buffer value.
func DRI2BufferRead(buf []byte, v *DRI2Buffer) int {
	b := 0

	v.Attachment = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.Name = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.Pitch = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.Cpp = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.Flags = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	return b
}

// DRI2BufferReadList reads a byte slice into a list of DRI2Buffer values.
func DRI2BufferReadList(buf []byte, dest []DRI2Buffer) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = DRI2Buffer{}
		b += DRI2BufferRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a DRI2Buffer value to a byte slice.
func (v DRI2Buffer) Bytes() []byte {
	buf := make([]byte, 20)
	b := 0

	binary.LittleEndian.PutUint32(buf[b:], v.Attachment)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], v.Name)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], v.Pitch)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], v.Cpp)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], v.Flags)
	b += 4

	return buf[:b]
}

// DRI2BufferListBytes writes a list of DRI2Buffer values to a byte slice.
func DRI2BufferListBytes(buf []byte, list []DRI2Buffer) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

const (
	DriverTypeDri   = 0
	DriverTypeVdpau = 1
)

const (
	EventTypeExchangeComplete = 1
	EventTypeBlitComplete     = 2
	EventTypeFlipComplete     = 3
)

// InvalidateBuffers is the event number for a InvalidateBuffersEvent.
const InvalidateBuffers = 1

type InvalidateBuffersEvent struct {
	Sequence uint16
	// padding: 1 bytes
	Drawable xproto.Drawable
}

// UnmarshalInvalidateBuffersEvent constructs a InvalidateBuffersEvent value that implements xgb.Event from a byte slice.
func UnmarshalInvalidateBuffersEvent(buf []byte) (xgb.XEvent, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"InvalidateBuffersEvent\"", len(buf))
	}

	v := InvalidateBuffersEvent{}
	b := 1 // don't read event number

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Drawable = xproto.Drawable(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	return v, nil
}

// Bytes writes a InvalidateBuffersEvent value to a byte slice.
func (v InvalidateBuffersEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 1
	b += 1

	b += 1 // padding

	b += 2 // skip sequence number

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Drawable))
	b += 4

	return buf
}

// SeqID returns the sequence id attached to the InvalidateBuffers event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v InvalidateBuffersEvent) SeqID() uint16 {
	return v.Sequence
}

func init() { registerEvent(1, UnmarshalInvalidateBuffersEvent) }

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

// Authenticate sends a checked request.
func Authenticate(c *xgb.XConn, Window xproto.Window, Magic uint32) (AuthenticateReply, error) {
	var reply AuthenticateReply
	op, ok := c.Ext("DRI2")
	if !ok {
		return reply, errors.New("cannot issue request \"Authenticate\" using the uninitialized extension \"DRI2\". dri2.Register(xconn) must be called first.")
	}
	err := c.SendRecv(authenticateRequest(op, Window, Magic), &reply)
	return reply, err
}

// AuthenticateUnchecked sends an unchecked request.
func AuthenticateUnchecked(c *xgb.XConn, Window xproto.Window, Magic uint32) error {
	op, ok := c.Ext("DRI2")
	if !ok {
		return errors.New("cannot issue request \"Authenticate\" using the uninitialized extension \"DRI2\". dri2.Register(xconn) must be called first.")
	}
	return c.Send(authenticateRequest(op, Window, Magic))
}

// AuthenticateReply represents the data returned from a Authenticate request.
type AuthenticateReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Authenticated uint32
}

// Unmarshal reads a byte slice into a AuthenticateReply value.
func (v *AuthenticateReply) Unmarshal(buf []byte) error {
	if size := 12; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"AuthenticateReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Authenticated = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	return nil
}

// Write request to wire for Authenticate
// authenticateRequest writes a Authenticate request to a byte slice.
func authenticateRequest(opcode uint8, Window xproto.Window, Magic uint32) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 2 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], Magic)
	b += 4

	return buf
}

// Connect sends a checked request.
func Connect(c *xgb.XConn, Window xproto.Window, DriverType uint32) (ConnectReply, error) {
	var reply ConnectReply
	op, ok := c.Ext("DRI2")
	if !ok {
		return reply, errors.New("cannot issue request \"Connect\" using the uninitialized extension \"DRI2\". dri2.Register(xconn) must be called first.")
	}
	err := c.SendRecv(connectRequest(op, Window, DriverType), &reply)
	return reply, err
}

// ConnectUnchecked sends an unchecked request.
func ConnectUnchecked(c *xgb.XConn, Window xproto.Window, DriverType uint32) error {
	op, ok := c.Ext("DRI2")
	if !ok {
		return errors.New("cannot issue request \"Connect\" using the uninitialized extension \"DRI2\". dri2.Register(xconn) must be called first.")
	}
	return c.Send(connectRequest(op, Window, DriverType))
}

// ConnectReply represents the data returned from a Connect request.
type ConnectReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	DriverNameLength uint32
	DeviceNameLength uint32
	// padding: 16 bytes
	DriverName   string // size: internal.Pad4((int(DriverNameLength) * 1))
	AlignmentPad []byte // size: internal.Pad4(((((int(DriverNameLength) + 3) & -4) - int(DriverNameLength)) * 1))
	DeviceName   string // size: internal.Pad4((int(DeviceNameLength) * 1))
}

// Unmarshal reads a byte slice into a ConnectReply value.
func (v *ConnectReply) Unmarshal(buf []byte) error {
	if size := (((32 + internal.Pad4((int(v.DriverNameLength) * 1))) + internal.Pad4(((((int(v.DriverNameLength) + 3) & -4) - int(v.DriverNameLength)) * 1))) + internal.Pad4((int(v.DeviceNameLength) * 1))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"ConnectReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.DriverNameLength = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.DeviceNameLength = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 16 // padding

	{
		byteString := make([]byte, v.DriverNameLength)
		copy(byteString[:v.DriverNameLength], buf[b:])
		v.DriverName = string(byteString)
		b += int(v.DriverNameLength)
	}

	v.AlignmentPad = make([]byte, (((int(v.DriverNameLength) + 3) & -4) - int(v.DriverNameLength)))
	copy(v.AlignmentPad[:(((int(v.DriverNameLength)+3)&-4)-int(v.DriverNameLength))], buf[b:])
	b += int((((int(v.DriverNameLength) + 3) & -4) - int(v.DriverNameLength)))

	{
		byteString := make([]byte, v.DeviceNameLength)
		copy(byteString[:v.DeviceNameLength], buf[b:])
		v.DeviceName = string(byteString)
		b += int(v.DeviceNameLength)
	}

	return nil
}

// Write request to wire for Connect
// connectRequest writes a Connect request to a byte slice.
func connectRequest(opcode uint8, Window xproto.Window, DriverType uint32) []byte {
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

	binary.LittleEndian.PutUint32(buf[b:], DriverType)
	b += 4

	return buf
}

// CopyRegion sends a checked request.
func CopyRegion(c *xgb.XConn, Drawable xproto.Drawable, Region uint32, Dest uint32, Src uint32) (CopyRegionReply, error) {
	var reply CopyRegionReply
	op, ok := c.Ext("DRI2")
	if !ok {
		return reply, errors.New("cannot issue request \"CopyRegion\" using the uninitialized extension \"DRI2\". dri2.Register(xconn) must be called first.")
	}
	err := c.SendRecv(copyRegionRequest(op, Drawable, Region, Dest, Src), &reply)
	return reply, err
}

// CopyRegionUnchecked sends an unchecked request.
func CopyRegionUnchecked(c *xgb.XConn, Drawable xproto.Drawable, Region uint32, Dest uint32, Src uint32) error {
	op, ok := c.Ext("DRI2")
	if !ok {
		return errors.New("cannot issue request \"CopyRegion\" using the uninitialized extension \"DRI2\". dri2.Register(xconn) must be called first.")
	}
	return c.Send(copyRegionRequest(op, Drawable, Region, Dest, Src))
}

// CopyRegionReply represents the data returned from a CopyRegion request.
type CopyRegionReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
}

// Unmarshal reads a byte slice into a CopyRegionReply value.
func (v *CopyRegionReply) Unmarshal(buf []byte) error {
	if size := 8; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"CopyRegionReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	return nil
}

// Write request to wire for CopyRegion
// copyRegionRequest writes a CopyRegion request to a byte slice.
func copyRegionRequest(opcode uint8, Drawable xproto.Drawable, Region uint32, Dest uint32, Src uint32) []byte {
	size := 20
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 6 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Drawable))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], Region)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], Dest)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], Src)
	b += 4

	return buf
}

// CreateDrawable sends a checked request.
func CreateDrawable(c *xgb.XConn, Drawable xproto.Drawable) error {
	op, ok := c.Ext("DRI2")
	if !ok {
		return errors.New("cannot issue request \"CreateDrawable\" using the uninitialized extension \"DRI2\". dri2.Register(xconn) must be called first.")
	}
	return c.SendRecv(createDrawableRequest(op, Drawable), nil)
}

// CreateDrawableUnchecked sends an unchecked request.
func CreateDrawableUnchecked(c *xgb.XConn, Drawable xproto.Drawable) error {
	op, ok := c.Ext("DRI2")
	if !ok {
		return errors.New("cannot issue request \"CreateDrawable\" using the uninitialized extension \"DRI2\". dri2.Register(xconn) must be called first.")
	}
	return c.Send(createDrawableRequest(op, Drawable))
}

// Write request to wire for CreateDrawable
// createDrawableRequest writes a CreateDrawable request to a byte slice.
func createDrawableRequest(opcode uint8, Drawable xproto.Drawable) []byte {
	size := 8
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

	return buf
}

// DestroyDrawable sends a checked request.
func DestroyDrawable(c *xgb.XConn, Drawable xproto.Drawable) error {
	op, ok := c.Ext("DRI2")
	if !ok {
		return errors.New("cannot issue request \"DestroyDrawable\" using the uninitialized extension \"DRI2\". dri2.Register(xconn) must be called first.")
	}
	return c.SendRecv(destroyDrawableRequest(op, Drawable), nil)
}

// DestroyDrawableUnchecked sends an unchecked request.
func DestroyDrawableUnchecked(c *xgb.XConn, Drawable xproto.Drawable) error {
	op, ok := c.Ext("DRI2")
	if !ok {
		return errors.New("cannot issue request \"DestroyDrawable\" using the uninitialized extension \"DRI2\". dri2.Register(xconn) must be called first.")
	}
	return c.Send(destroyDrawableRequest(op, Drawable))
}

// Write request to wire for DestroyDrawable
// destroyDrawableRequest writes a DestroyDrawable request to a byte slice.
func destroyDrawableRequest(opcode uint8, Drawable xproto.Drawable) []byte {
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

// GetBuffers sends a checked request.
func GetBuffers(c *xgb.XConn, Drawable xproto.Drawable, Count uint32, Attachments []uint32) (GetBuffersReply, error) {
	var reply GetBuffersReply
	op, ok := c.Ext("DRI2")
	if !ok {
		return reply, errors.New("cannot issue request \"GetBuffers\" using the uninitialized extension \"DRI2\". dri2.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getBuffersRequest(op, Drawable, Count, Attachments), &reply)
	return reply, err
}

// GetBuffersUnchecked sends an unchecked request.
func GetBuffersUnchecked(c *xgb.XConn, Drawable xproto.Drawable, Count uint32, Attachments []uint32) error {
	op, ok := c.Ext("DRI2")
	if !ok {
		return errors.New("cannot issue request \"GetBuffers\" using the uninitialized extension \"DRI2\". dri2.Register(xconn) must be called first.")
	}
	return c.Send(getBuffersRequest(op, Drawable, Count, Attachments))
}

// GetBuffersReply represents the data returned from a GetBuffers request.
type GetBuffersReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Width  uint32
	Height uint32
	Count  uint32
	// padding: 12 bytes
	Buffers []DRI2Buffer // size: internal.Pad4((int(Count) * 20))
}

// Unmarshal reads a byte slice into a GetBuffersReply value.
func (v *GetBuffersReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.Count) * 20))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetBuffersReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Width = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.Height = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.Count = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 12 // padding

	v.Buffers = make([]DRI2Buffer, v.Count)
	b += DRI2BufferReadList(buf[b:], v.Buffers)

	return nil
}

// Write request to wire for GetBuffers
// getBuffersRequest writes a GetBuffers request to a byte slice.
func getBuffersRequest(opcode uint8, Drawable xproto.Drawable, Count uint32, Attachments []uint32) []byte {
	size := internal.Pad4((12 + internal.Pad4((len(Attachments) * 4))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 5 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Drawable))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], Count)
	b += 4

	for i := 0; i < int(len(Attachments)); i++ {
		binary.LittleEndian.PutUint32(buf[b:], Attachments[i])
		b += 4
	}

	return buf
}

// GetBuffersWithFormat sends a checked request.
func GetBuffersWithFormat(c *xgb.XConn, Drawable xproto.Drawable, Count uint32, Attachments []AttachFormat) (GetBuffersWithFormatReply, error) {
	var reply GetBuffersWithFormatReply
	op, ok := c.Ext("DRI2")
	if !ok {
		return reply, errors.New("cannot issue request \"GetBuffersWithFormat\" using the uninitialized extension \"DRI2\". dri2.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getBuffersWithFormatRequest(op, Drawable, Count, Attachments), &reply)
	return reply, err
}

// GetBuffersWithFormatUnchecked sends an unchecked request.
func GetBuffersWithFormatUnchecked(c *xgb.XConn, Drawable xproto.Drawable, Count uint32, Attachments []AttachFormat) error {
	op, ok := c.Ext("DRI2")
	if !ok {
		return errors.New("cannot issue request \"GetBuffersWithFormat\" using the uninitialized extension \"DRI2\". dri2.Register(xconn) must be called first.")
	}
	return c.Send(getBuffersWithFormatRequest(op, Drawable, Count, Attachments))
}

// GetBuffersWithFormatReply represents the data returned from a GetBuffersWithFormat request.
type GetBuffersWithFormatReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Width  uint32
	Height uint32
	Count  uint32
	// padding: 12 bytes
	Buffers []DRI2Buffer // size: internal.Pad4((int(Count) * 20))
}

// Unmarshal reads a byte slice into a GetBuffersWithFormatReply value.
func (v *GetBuffersWithFormatReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.Count) * 20))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetBuffersWithFormatReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Width = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.Height = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.Count = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 12 // padding

	v.Buffers = make([]DRI2Buffer, v.Count)
	b += DRI2BufferReadList(buf[b:], v.Buffers)

	return nil
}

// Write request to wire for GetBuffersWithFormat
// getBuffersWithFormatRequest writes a GetBuffersWithFormat request to a byte slice.
func getBuffersWithFormatRequest(opcode uint8, Drawable xproto.Drawable, Count uint32, Attachments []AttachFormat) []byte {
	size := internal.Pad4((12 + internal.Pad4((len(Attachments) * 8))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 7 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Drawable))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], Count)
	b += 4

	b += AttachFormatListBytes(buf[b:], Attachments)

	return buf
}

// GetMSC sends a checked request.
func GetMSC(c *xgb.XConn, Drawable xproto.Drawable) (GetMSCReply, error) {
	var reply GetMSCReply
	op, ok := c.Ext("DRI2")
	if !ok {
		return reply, errors.New("cannot issue request \"GetMSC\" using the uninitialized extension \"DRI2\". dri2.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getMSCRequest(op, Drawable), &reply)
	return reply, err
}

// GetMSCUnchecked sends an unchecked request.
func GetMSCUnchecked(c *xgb.XConn, Drawable xproto.Drawable) error {
	op, ok := c.Ext("DRI2")
	if !ok {
		return errors.New("cannot issue request \"GetMSC\" using the uninitialized extension \"DRI2\". dri2.Register(xconn) must be called first.")
	}
	return c.Send(getMSCRequest(op, Drawable))
}

// GetMSCReply represents the data returned from a GetMSC request.
type GetMSCReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	UstHi uint32
	UstLo uint32
	MscHi uint32
	MscLo uint32
	SbcHi uint32
	SbcLo uint32
}

// Unmarshal reads a byte slice into a GetMSCReply value.
func (v *GetMSCReply) Unmarshal(buf []byte) error {
	if size := 32; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetMSCReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.UstHi = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.UstLo = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.MscHi = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.MscLo = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.SbcHi = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.SbcLo = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	return nil
}

// Write request to wire for GetMSC
// getMSCRequest writes a GetMSC request to a byte slice.
func getMSCRequest(opcode uint8, Drawable xproto.Drawable) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 9 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Drawable))
	b += 4

	return buf
}

// GetParam sends a checked request.
func GetParam(c *xgb.XConn, Drawable xproto.Drawable, Param uint32) (GetParamReply, error) {
	var reply GetParamReply
	op, ok := c.Ext("DRI2")
	if !ok {
		return reply, errors.New("cannot issue request \"GetParam\" using the uninitialized extension \"DRI2\". dri2.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getParamRequest(op, Drawable, Param), &reply)
	return reply, err
}

// GetParamUnchecked sends an unchecked request.
func GetParamUnchecked(c *xgb.XConn, Drawable xproto.Drawable, Param uint32) error {
	op, ok := c.Ext("DRI2")
	if !ok {
		return errors.New("cannot issue request \"GetParam\" using the uninitialized extension \"DRI2\". dri2.Register(xconn) must be called first.")
	}
	return c.Send(getParamRequest(op, Drawable, Param))
}

// GetParamReply represents the data returned from a GetParam request.
type GetParamReply struct {
	Sequence          uint16 // sequence number of the request for this reply
	Length            uint32 // number of bytes in this reply
	IsParamRecognized bool
	ValueHi           uint32
	ValueLo           uint32
}

// Unmarshal reads a byte slice into a GetParamReply value.
func (v *GetParamReply) Unmarshal(buf []byte) error {
	if size := 16; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetParamReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	v.IsParamRecognized = (buf[b] == 1)
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.ValueHi = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.ValueLo = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	return nil
}

// Write request to wire for GetParam
// getParamRequest writes a GetParam request to a byte slice.
func getParamRequest(opcode uint8, Drawable xproto.Drawable, Param uint32) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 13 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Drawable))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], Param)
	b += 4

	return buf
}

// QueryVersion sends a checked request.
func QueryVersion(c *xgb.XConn, MajorVersion uint32, MinorVersion uint32) (QueryVersionReply, error) {
	var reply QueryVersionReply
	op, ok := c.Ext("DRI2")
	if !ok {
		return reply, errors.New("cannot issue request \"QueryVersion\" using the uninitialized extension \"DRI2\". dri2.Register(xconn) must be called first.")
	}
	err := c.SendRecv(queryVersionRequest(op, MajorVersion, MinorVersion), &reply)
	return reply, err
}

// QueryVersionUnchecked sends an unchecked request.
func QueryVersionUnchecked(c *xgb.XConn, MajorVersion uint32, MinorVersion uint32) error {
	op, ok := c.Ext("DRI2")
	if !ok {
		return errors.New("cannot issue request \"QueryVersion\" using the uninitialized extension \"DRI2\". dri2.Register(xconn) must be called first.")
	}
	return c.Send(queryVersionRequest(op, MajorVersion, MinorVersion))
}

// QueryVersionReply represents the data returned from a QueryVersion request.
type QueryVersionReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	MajorVersion uint32
	MinorVersion uint32
}

// Unmarshal reads a byte slice into a QueryVersionReply value.
func (v *QueryVersionReply) Unmarshal(buf []byte) error {
	if size := 16; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"QueryVersionReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.MajorVersion = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.MinorVersion = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	return nil
}

// Write request to wire for QueryVersion
// queryVersionRequest writes a QueryVersion request to a byte slice.
func queryVersionRequest(opcode uint8, MajorVersion uint32, MinorVersion uint32) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 0 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], MajorVersion)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], MinorVersion)
	b += 4

	return buf
}

// SwapBuffers sends a checked request.
func SwapBuffers(c *xgb.XConn, Drawable xproto.Drawable, TargetMscHi uint32, TargetMscLo uint32, DivisorHi uint32, DivisorLo uint32, RemainderHi uint32, RemainderLo uint32) (SwapBuffersReply, error) {
	var reply SwapBuffersReply
	op, ok := c.Ext("DRI2")
	if !ok {
		return reply, errors.New("cannot issue request \"SwapBuffers\" using the uninitialized extension \"DRI2\". dri2.Register(xconn) must be called first.")
	}
	err := c.SendRecv(swapBuffersRequest(op, Drawable, TargetMscHi, TargetMscLo, DivisorHi, DivisorLo, RemainderHi, RemainderLo), &reply)
	return reply, err
}

// SwapBuffersUnchecked sends an unchecked request.
func SwapBuffersUnchecked(c *xgb.XConn, Drawable xproto.Drawable, TargetMscHi uint32, TargetMscLo uint32, DivisorHi uint32, DivisorLo uint32, RemainderHi uint32, RemainderLo uint32) error {
	op, ok := c.Ext("DRI2")
	if !ok {
		return errors.New("cannot issue request \"SwapBuffers\" using the uninitialized extension \"DRI2\". dri2.Register(xconn) must be called first.")
	}
	return c.Send(swapBuffersRequest(op, Drawable, TargetMscHi, TargetMscLo, DivisorHi, DivisorLo, RemainderHi, RemainderLo))
}

// SwapBuffersReply represents the data returned from a SwapBuffers request.
type SwapBuffersReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	SwapHi uint32
	SwapLo uint32
}

// Unmarshal reads a byte slice into a SwapBuffersReply value.
func (v *SwapBuffersReply) Unmarshal(buf []byte) error {
	if size := 16; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"SwapBuffersReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.SwapHi = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.SwapLo = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	return nil
}

// Write request to wire for SwapBuffers
// swapBuffersRequest writes a SwapBuffers request to a byte slice.
func swapBuffersRequest(opcode uint8, Drawable xproto.Drawable, TargetMscHi uint32, TargetMscLo uint32, DivisorHi uint32, DivisorLo uint32, RemainderHi uint32, RemainderLo uint32) []byte {
	size := 32
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 8 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Drawable))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], TargetMscHi)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], TargetMscLo)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], DivisorHi)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], DivisorLo)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], RemainderHi)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], RemainderLo)
	b += 4

	return buf
}

// SwapInterval sends a checked request.
func SwapInterval(c *xgb.XConn, Drawable xproto.Drawable, Interval uint32) error {
	op, ok := c.Ext("DRI2")
	if !ok {
		return errors.New("cannot issue request \"SwapInterval\" using the uninitialized extension \"DRI2\". dri2.Register(xconn) must be called first.")
	}
	return c.SendRecv(swapIntervalRequest(op, Drawable, Interval), nil)
}

// SwapIntervalUnchecked sends an unchecked request.
func SwapIntervalUnchecked(c *xgb.XConn, Drawable xproto.Drawable, Interval uint32) error {
	op, ok := c.Ext("DRI2")
	if !ok {
		return errors.New("cannot issue request \"SwapInterval\" using the uninitialized extension \"DRI2\". dri2.Register(xconn) must be called first.")
	}
	return c.Send(swapIntervalRequest(op, Drawable, Interval))
}

// Write request to wire for SwapInterval
// swapIntervalRequest writes a SwapInterval request to a byte slice.
func swapIntervalRequest(opcode uint8, Drawable xproto.Drawable, Interval uint32) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 12 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Drawable))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], Interval)
	b += 4

	return buf
}

// WaitMSC sends a checked request.
func WaitMSC(c *xgb.XConn, Drawable xproto.Drawable, TargetMscHi uint32, TargetMscLo uint32, DivisorHi uint32, DivisorLo uint32, RemainderHi uint32, RemainderLo uint32) (WaitMSCReply, error) {
	var reply WaitMSCReply
	op, ok := c.Ext("DRI2")
	if !ok {
		return reply, errors.New("cannot issue request \"WaitMSC\" using the uninitialized extension \"DRI2\". dri2.Register(xconn) must be called first.")
	}
	err := c.SendRecv(waitMSCRequest(op, Drawable, TargetMscHi, TargetMscLo, DivisorHi, DivisorLo, RemainderHi, RemainderLo), &reply)
	return reply, err
}

// WaitMSCUnchecked sends an unchecked request.
func WaitMSCUnchecked(c *xgb.XConn, Drawable xproto.Drawable, TargetMscHi uint32, TargetMscLo uint32, DivisorHi uint32, DivisorLo uint32, RemainderHi uint32, RemainderLo uint32) error {
	op, ok := c.Ext("DRI2")
	if !ok {
		return errors.New("cannot issue request \"WaitMSC\" using the uninitialized extension \"DRI2\". dri2.Register(xconn) must be called first.")
	}
	return c.Send(waitMSCRequest(op, Drawable, TargetMscHi, TargetMscLo, DivisorHi, DivisorLo, RemainderHi, RemainderLo))
}

// WaitMSCReply represents the data returned from a WaitMSC request.
type WaitMSCReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	UstHi uint32
	UstLo uint32
	MscHi uint32
	MscLo uint32
	SbcHi uint32
	SbcLo uint32
}

// Unmarshal reads a byte slice into a WaitMSCReply value.
func (v *WaitMSCReply) Unmarshal(buf []byte) error {
	if size := 32; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"WaitMSCReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.UstHi = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.UstLo = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.MscHi = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.MscLo = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.SbcHi = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.SbcLo = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	return nil
}

// Write request to wire for WaitMSC
// waitMSCRequest writes a WaitMSC request to a byte slice.
func waitMSCRequest(opcode uint8, Drawable xproto.Drawable, TargetMscHi uint32, TargetMscLo uint32, DivisorHi uint32, DivisorLo uint32, RemainderHi uint32, RemainderLo uint32) []byte {
	size := 32
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 10 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Drawable))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], TargetMscHi)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], TargetMscLo)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], DivisorHi)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], DivisorLo)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], RemainderHi)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], RemainderLo)
	b += 4

	return buf
}

// WaitSBC sends a checked request.
func WaitSBC(c *xgb.XConn, Drawable xproto.Drawable, TargetSbcHi uint32, TargetSbcLo uint32) (WaitSBCReply, error) {
	var reply WaitSBCReply
	op, ok := c.Ext("DRI2")
	if !ok {
		return reply, errors.New("cannot issue request \"WaitSBC\" using the uninitialized extension \"DRI2\". dri2.Register(xconn) must be called first.")
	}
	err := c.SendRecv(waitSBCRequest(op, Drawable, TargetSbcHi, TargetSbcLo), &reply)
	return reply, err
}

// WaitSBCUnchecked sends an unchecked request.
func WaitSBCUnchecked(c *xgb.XConn, Drawable xproto.Drawable, TargetSbcHi uint32, TargetSbcLo uint32) error {
	op, ok := c.Ext("DRI2")
	if !ok {
		return errors.New("cannot issue request \"WaitSBC\" using the uninitialized extension \"DRI2\". dri2.Register(xconn) must be called first.")
	}
	return c.Send(waitSBCRequest(op, Drawable, TargetSbcHi, TargetSbcLo))
}

// WaitSBCReply represents the data returned from a WaitSBC request.
type WaitSBCReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	UstHi uint32
	UstLo uint32
	MscHi uint32
	MscLo uint32
	SbcHi uint32
	SbcLo uint32
}

// Unmarshal reads a byte slice into a WaitSBCReply value.
func (v *WaitSBCReply) Unmarshal(buf []byte) error {
	if size := 32; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"WaitSBCReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.UstHi = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.UstLo = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.MscHi = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.MscLo = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.SbcHi = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.SbcLo = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	return nil
}

// Write request to wire for WaitSBC
// waitSBCRequest writes a WaitSBC request to a byte slice.
func waitSBCRequest(opcode uint8, Drawable xproto.Drawable, TargetSbcHi uint32, TargetSbcLo uint32) []byte {
	size := 16
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 11 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Drawable))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], TargetSbcHi)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], TargetSbcLo)
	b += 4

	return buf
}
