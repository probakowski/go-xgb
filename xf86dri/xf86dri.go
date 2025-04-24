// FILE GENERATED AUTOMATICALLY FROM "xf86dri.xml"
package xf86dri

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/probakowski/go-xgb"
	"github.com/probakowski/go-xgb/internal"
	"github.com/probakowski/go-xgb/xproto"
)

const (
	// ExtName is the user-friendly name string of this X extension.
	ExtName = "XF86Dri"

	// ExtXName is the name string this extension is known by to the X server.
	ExtXName = "XFree86-DRI"
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

// Register will query the X server for XF86Dri extension support, and register relevant extension unmarshalers with the XConn.
func Register(xconn *xgb.XConn) error {
	// Query the X server for this extension
	reply, err := xproto.QueryExtension(xconn, uint16(len(ExtXName)), ExtXName)
	if err != nil {
		return fmt.Errorf("error querying X for \"XF86Dri\": %w", err)
	} else if !reply.Present {
		return fmt.Errorf("no extension named \"XF86Dri\" is known to the X server: reply=%+v", reply)
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

type DrmClipRect struct {
	X1 int16
	Y1 int16
	X2 int16
	X3 int16
}

// DrmClipRectRead reads a byte slice into a DrmClipRect value.
func DrmClipRectRead(buf []byte, v *DrmClipRect) int {
	b := 0

	v.X1 = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.Y1 = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.X2 = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.X3 = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	return b
}

// DrmClipRectReadList reads a byte slice into a list of DrmClipRect values.
func DrmClipRectReadList(buf []byte, dest []DrmClipRect) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = DrmClipRect{}
		b += DrmClipRectRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a DrmClipRect value to a byte slice.
func (v DrmClipRect) Bytes() []byte {
	buf := make([]byte, 8)
	b := 0

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.X1))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.Y1))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.X2))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.X3))
	b += 2

	return buf[:b]
}

// DrmClipRectListBytes writes a list of DrmClipRect values to a byte slice.
func DrmClipRectListBytes(buf []byte, list []DrmClipRect) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
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

// AuthConnection sends a checked request.
func AuthConnection(c *xgb.XConn, Screen uint32, Magic uint32) (AuthConnectionReply, error) {
	var reply AuthConnectionReply
	op, ok := c.Ext("XFree86-DRI")
	if !ok {
		return reply, errors.New("cannot issue request \"AuthConnection\" using the uninitialized extension \"XFree86-DRI\". xf86dri.Register(xconn) must be called first.")
	}
	err := c.SendRecv(authConnectionRequest(op, Screen, Magic), &reply)
	return reply, err
}

// AuthConnectionUnchecked sends an unchecked request.
func AuthConnectionUnchecked(c *xgb.XConn, Screen uint32, Magic uint32) error {
	op, ok := c.Ext("XFree86-DRI")
	if !ok {
		return errors.New("cannot issue request \"AuthConnection\" using the uninitialized extension \"XFree86-DRI\". xf86dri.Register(xconn) must be called first.")
	}
	return c.Send(authConnectionRequest(op, Screen, Magic))
}

// AuthConnectionReply represents the data returned from a AuthConnection request.
type AuthConnectionReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Authenticated uint32
}

// Unmarshal reads a byte slice into a AuthConnectionReply value.
func (v *AuthConnectionReply) Unmarshal(buf []byte) error {
	const size = 12
	if len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"AuthConnectionReply\": have=%d need=%d", len(buf), size)
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

// Write request to wire for AuthConnection
// authConnectionRequest writes a AuthConnection request to a byte slice.
func authConnectionRequest(opcode uint8, Screen uint32, Magic uint32) []byte {
	const size = 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 11 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], Screen)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], Magic)
	b += 4

	return buf
}

// CloseConnection sends a checked request.
func CloseConnection(c *xgb.XConn, Screen uint32) error {
	op, ok := c.Ext("XFree86-DRI")
	if !ok {
		return errors.New("cannot issue request \"CloseConnection\" using the uninitialized extension \"XFree86-DRI\". xf86dri.Register(xconn) must be called first.")
	}
	return c.SendRecv(closeConnectionRequest(op, Screen), nil)
}

// CloseConnectionUnchecked sends an unchecked request.
func CloseConnectionUnchecked(c *xgb.XConn, Screen uint32) error {
	op, ok := c.Ext("XFree86-DRI")
	if !ok {
		return errors.New("cannot issue request \"CloseConnection\" using the uninitialized extension \"XFree86-DRI\". xf86dri.Register(xconn) must be called first.")
	}
	return c.Send(closeConnectionRequest(op, Screen))
}

// Write request to wire for CloseConnection
// closeConnectionRequest writes a CloseConnection request to a byte slice.
func closeConnectionRequest(opcode uint8, Screen uint32) []byte {
	const size = 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 3 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], Screen)
	b += 4

	return buf
}

// CreateContext sends a checked request.
func CreateContext(c *xgb.XConn, Screen uint32, Visual uint32, Context uint32) (CreateContextReply, error) {
	var reply CreateContextReply
	op, ok := c.Ext("XFree86-DRI")
	if !ok {
		return reply, errors.New("cannot issue request \"CreateContext\" using the uninitialized extension \"XFree86-DRI\". xf86dri.Register(xconn) must be called first.")
	}
	err := c.SendRecv(createContextRequest(op, Screen, Visual, Context), &reply)
	return reply, err
}

// CreateContextUnchecked sends an unchecked request.
func CreateContextUnchecked(c *xgb.XConn, Screen uint32, Visual uint32, Context uint32) error {
	op, ok := c.Ext("XFree86-DRI")
	if !ok {
		return errors.New("cannot issue request \"CreateContext\" using the uninitialized extension \"XFree86-DRI\". xf86dri.Register(xconn) must be called first.")
	}
	return c.Send(createContextRequest(op, Screen, Visual, Context))
}

// CreateContextReply represents the data returned from a CreateContext request.
type CreateContextReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	HwContext uint32
}

// Unmarshal reads a byte slice into a CreateContextReply value.
func (v *CreateContextReply) Unmarshal(buf []byte) error {
	const size = 12
	if len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"CreateContextReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.HwContext = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	return nil
}

// Write request to wire for CreateContext
// createContextRequest writes a CreateContext request to a byte slice.
func createContextRequest(opcode uint8, Screen uint32, Visual uint32, Context uint32) []byte {
	const size = 16
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 5 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], Screen)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], Visual)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], Context)
	b += 4

	return buf
}

// CreateDrawable sends a checked request.
func CreateDrawable(c *xgb.XConn, Screen uint32, Drawable uint32) (CreateDrawableReply, error) {
	var reply CreateDrawableReply
	op, ok := c.Ext("XFree86-DRI")
	if !ok {
		return reply, errors.New("cannot issue request \"CreateDrawable\" using the uninitialized extension \"XFree86-DRI\". xf86dri.Register(xconn) must be called first.")
	}
	err := c.SendRecv(createDrawableRequest(op, Screen, Drawable), &reply)
	return reply, err
}

// CreateDrawableUnchecked sends an unchecked request.
func CreateDrawableUnchecked(c *xgb.XConn, Screen uint32, Drawable uint32) error {
	op, ok := c.Ext("XFree86-DRI")
	if !ok {
		return errors.New("cannot issue request \"CreateDrawable\" using the uninitialized extension \"XFree86-DRI\". xf86dri.Register(xconn) must be called first.")
	}
	return c.Send(createDrawableRequest(op, Screen, Drawable))
}

// CreateDrawableReply represents the data returned from a CreateDrawable request.
type CreateDrawableReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	HwDrawableHandle uint32
}

// Unmarshal reads a byte slice into a CreateDrawableReply value.
func (v *CreateDrawableReply) Unmarshal(buf []byte) error {
	const size = 12
	if len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"CreateDrawableReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.HwDrawableHandle = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	return nil
}

// Write request to wire for CreateDrawable
// createDrawableRequest writes a CreateDrawable request to a byte slice.
func createDrawableRequest(opcode uint8, Screen uint32, Drawable uint32) []byte {
	const size = 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 7 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], Screen)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], Drawable)
	b += 4

	return buf
}

// DestroyContext sends a checked request.
func DestroyContext(c *xgb.XConn, Screen uint32, Context uint32) error {
	op, ok := c.Ext("XFree86-DRI")
	if !ok {
		return errors.New("cannot issue request \"DestroyContext\" using the uninitialized extension \"XFree86-DRI\". xf86dri.Register(xconn) must be called first.")
	}
	return c.SendRecv(destroyContextRequest(op, Screen, Context), nil)
}

// DestroyContextUnchecked sends an unchecked request.
func DestroyContextUnchecked(c *xgb.XConn, Screen uint32, Context uint32) error {
	op, ok := c.Ext("XFree86-DRI")
	if !ok {
		return errors.New("cannot issue request \"DestroyContext\" using the uninitialized extension \"XFree86-DRI\". xf86dri.Register(xconn) must be called first.")
	}
	return c.Send(destroyContextRequest(op, Screen, Context))
}

// Write request to wire for DestroyContext
// destroyContextRequest writes a DestroyContext request to a byte slice.
func destroyContextRequest(opcode uint8, Screen uint32, Context uint32) []byte {
	const size = 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 6 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], Screen)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], Context)
	b += 4

	return buf
}

// DestroyDrawable sends a checked request.
func DestroyDrawable(c *xgb.XConn, Screen uint32, Drawable uint32) error {
	op, ok := c.Ext("XFree86-DRI")
	if !ok {
		return errors.New("cannot issue request \"DestroyDrawable\" using the uninitialized extension \"XFree86-DRI\". xf86dri.Register(xconn) must be called first.")
	}
	return c.SendRecv(destroyDrawableRequest(op, Screen, Drawable), nil)
}

// DestroyDrawableUnchecked sends an unchecked request.
func DestroyDrawableUnchecked(c *xgb.XConn, Screen uint32, Drawable uint32) error {
	op, ok := c.Ext("XFree86-DRI")
	if !ok {
		return errors.New("cannot issue request \"DestroyDrawable\" using the uninitialized extension \"XFree86-DRI\". xf86dri.Register(xconn) must be called first.")
	}
	return c.Send(destroyDrawableRequest(op, Screen, Drawable))
}

// Write request to wire for DestroyDrawable
// destroyDrawableRequest writes a DestroyDrawable request to a byte slice.
func destroyDrawableRequest(opcode uint8, Screen uint32, Drawable uint32) []byte {
	const size = 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 8 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], Screen)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], Drawable)
	b += 4

	return buf
}

// GetClientDriverName sends a checked request.
func GetClientDriverName(c *xgb.XConn, Screen uint32) (GetClientDriverNameReply, error) {
	var reply GetClientDriverNameReply
	op, ok := c.Ext("XFree86-DRI")
	if !ok {
		return reply, errors.New("cannot issue request \"GetClientDriverName\" using the uninitialized extension \"XFree86-DRI\". xf86dri.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getClientDriverNameRequest(op, Screen), &reply)
	return reply, err
}

// GetClientDriverNameUnchecked sends an unchecked request.
func GetClientDriverNameUnchecked(c *xgb.XConn, Screen uint32) error {
	op, ok := c.Ext("XFree86-DRI")
	if !ok {
		return errors.New("cannot issue request \"GetClientDriverName\" using the uninitialized extension \"XFree86-DRI\". xf86dri.Register(xconn) must be called first.")
	}
	return c.Send(getClientDriverNameRequest(op, Screen))
}

// GetClientDriverNameReply represents the data returned from a GetClientDriverName request.
type GetClientDriverNameReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	ClientDriverMajorVersion uint32
	ClientDriverMinorVersion uint32
	ClientDriverPatchVersion uint32
	ClientDriverNameLen      uint32
	// padding: 8 bytes
	ClientDriverName string // size: internal.Pad4((int(ClientDriverNameLen) * 1))
}

// Unmarshal reads a byte slice into a GetClientDriverNameReply value.
func (v *GetClientDriverNameReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.ClientDriverNameLen) * 1))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetClientDriverNameReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.ClientDriverMajorVersion = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.ClientDriverMinorVersion = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.ClientDriverPatchVersion = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.ClientDriverNameLen = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 8 // padding

	{
		byteString := make([]byte, v.ClientDriverNameLen)
		copy(byteString[:v.ClientDriverNameLen], buf[b:])
		v.ClientDriverName = string(byteString)
		b += int(v.ClientDriverNameLen)
	}

	return nil
}

// Write request to wire for GetClientDriverName
// getClientDriverNameRequest writes a GetClientDriverName request to a byte slice.
func getClientDriverNameRequest(opcode uint8, Screen uint32) []byte {
	const size = 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 4 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], Screen)
	b += 4

	return buf
}

// GetDeviceInfo sends a checked request.
func GetDeviceInfo(c *xgb.XConn, Screen uint32) (GetDeviceInfoReply, error) {
	var reply GetDeviceInfoReply
	op, ok := c.Ext("XFree86-DRI")
	if !ok {
		return reply, errors.New("cannot issue request \"GetDeviceInfo\" using the uninitialized extension \"XFree86-DRI\". xf86dri.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getDeviceInfoRequest(op, Screen), &reply)
	return reply, err
}

// GetDeviceInfoUnchecked sends an unchecked request.
func GetDeviceInfoUnchecked(c *xgb.XConn, Screen uint32) error {
	op, ok := c.Ext("XFree86-DRI")
	if !ok {
		return errors.New("cannot issue request \"GetDeviceInfo\" using the uninitialized extension \"XFree86-DRI\". xf86dri.Register(xconn) must be called first.")
	}
	return c.Send(getDeviceInfoRequest(op, Screen))
}

// GetDeviceInfoReply represents the data returned from a GetDeviceInfo request.
type GetDeviceInfoReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	FramebufferHandleLow    uint32
	FramebufferHandleHigh   uint32
	FramebufferOriginOffset uint32
	FramebufferSize         uint32
	FramebufferStride       uint32
	DevicePrivateSize       uint32
	DevicePrivate           []uint32 // size: internal.Pad4((int(DevicePrivateSize) * 4))
}

// Unmarshal reads a byte slice into a GetDeviceInfoReply value.
func (v *GetDeviceInfoReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.DevicePrivateSize) * 4))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetDeviceInfoReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.FramebufferHandleLow = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.FramebufferHandleHigh = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.FramebufferOriginOffset = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.FramebufferSize = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.FramebufferStride = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.DevicePrivateSize = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.DevicePrivate = make([]uint32, v.DevicePrivateSize)
	for i := 0; i < int(v.DevicePrivateSize); i++ {
		v.DevicePrivate[i] = binary.LittleEndian.Uint32(buf[b:])
		b += 4
	}

	return nil
}

// Write request to wire for GetDeviceInfo
// getDeviceInfoRequest writes a GetDeviceInfo request to a byte slice.
func getDeviceInfoRequest(opcode uint8, Screen uint32) []byte {
	const size = 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 10 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], Screen)
	b += 4

	return buf
}

// GetDrawableInfo sends a checked request.
func GetDrawableInfo(c *xgb.XConn, Screen uint32, Drawable uint32) (GetDrawableInfoReply, error) {
	var reply GetDrawableInfoReply
	op, ok := c.Ext("XFree86-DRI")
	if !ok {
		return reply, errors.New("cannot issue request \"GetDrawableInfo\" using the uninitialized extension \"XFree86-DRI\". xf86dri.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getDrawableInfoRequest(op, Screen, Drawable), &reply)
	return reply, err
}

// GetDrawableInfoUnchecked sends an unchecked request.
func GetDrawableInfoUnchecked(c *xgb.XConn, Screen uint32, Drawable uint32) error {
	op, ok := c.Ext("XFree86-DRI")
	if !ok {
		return errors.New("cannot issue request \"GetDrawableInfo\" using the uninitialized extension \"XFree86-DRI\". xf86dri.Register(xconn) must be called first.")
	}
	return c.Send(getDrawableInfoRequest(op, Screen, Drawable))
}

// GetDrawableInfoReply represents the data returned from a GetDrawableInfo request.
type GetDrawableInfoReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	DrawableTableIndex uint32
	DrawableTableStamp uint32
	DrawableOriginX    int16
	DrawableOriginY    int16
	DrawableSizeW      int16
	DrawableSizeH      int16
	NumClipRects       uint32
	BackX              int16
	BackY              int16
	NumBackClipRects   uint32
	ClipRects          []DrmClipRect // size: internal.Pad4((int(NumClipRects) * 8))
	// alignment gap to multiple of 4
	BackClipRects []DrmClipRect // size: internal.Pad4((int(NumBackClipRects) * 8))
}

// Unmarshal reads a byte slice into a GetDrawableInfoReply value.
func (v *GetDrawableInfoReply) Unmarshal(buf []byte) error {
	if size := (((36 + internal.Pad4((int(v.NumClipRects) * 8))) + 4) + internal.Pad4((int(v.NumBackClipRects) * 8))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetDrawableInfoReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.DrawableTableIndex = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.DrawableTableStamp = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.DrawableOriginX = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.DrawableOriginY = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.DrawableSizeW = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.DrawableSizeH = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.NumClipRects = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.BackX = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.BackY = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.NumBackClipRects = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.ClipRects = make([]DrmClipRect, v.NumClipRects)
	b += DrmClipRectReadList(buf[b:], v.ClipRects)

	b = (b + 3) & ^3 // alignment gap

	v.BackClipRects = make([]DrmClipRect, v.NumBackClipRects)
	b += DrmClipRectReadList(buf[b:], v.BackClipRects)

	return nil
}

// Write request to wire for GetDrawableInfo
// getDrawableInfoRequest writes a GetDrawableInfo request to a byte slice.
func getDrawableInfoRequest(opcode uint8, Screen uint32, Drawable uint32) []byte {
	const size = 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 9 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], Screen)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], Drawable)
	b += 4

	return buf
}

// OpenConnection sends a checked request.
func OpenConnection(c *xgb.XConn, Screen uint32) (OpenConnectionReply, error) {
	var reply OpenConnectionReply
	op, ok := c.Ext("XFree86-DRI")
	if !ok {
		return reply, errors.New("cannot issue request \"OpenConnection\" using the uninitialized extension \"XFree86-DRI\". xf86dri.Register(xconn) must be called first.")
	}
	err := c.SendRecv(openConnectionRequest(op, Screen), &reply)
	return reply, err
}

// OpenConnectionUnchecked sends an unchecked request.
func OpenConnectionUnchecked(c *xgb.XConn, Screen uint32) error {
	op, ok := c.Ext("XFree86-DRI")
	if !ok {
		return errors.New("cannot issue request \"OpenConnection\" using the uninitialized extension \"XFree86-DRI\". xf86dri.Register(xconn) must be called first.")
	}
	return c.Send(openConnectionRequest(op, Screen))
}

// OpenConnectionReply represents the data returned from a OpenConnection request.
type OpenConnectionReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	SareaHandleLow  uint32
	SareaHandleHigh uint32
	BusIdLen        uint32
	// padding: 12 bytes
	BusId string // size: internal.Pad4((int(BusIdLen) * 1))
}

// Unmarshal reads a byte slice into a OpenConnectionReply value.
func (v *OpenConnectionReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.BusIdLen) * 1))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"OpenConnectionReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.SareaHandleLow = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.SareaHandleHigh = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.BusIdLen = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 12 // padding

	{
		byteString := make([]byte, v.BusIdLen)
		copy(byteString[:v.BusIdLen], buf[b:])
		v.BusId = string(byteString)
		b += int(v.BusIdLen)
	}

	return nil
}

// Write request to wire for OpenConnection
// openConnectionRequest writes a OpenConnection request to a byte slice.
func openConnectionRequest(opcode uint8, Screen uint32) []byte {
	const size = 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 2 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], Screen)
	b += 4

	return buf
}

// QueryDirectRenderingCapable sends a checked request.
func QueryDirectRenderingCapable(c *xgb.XConn, Screen uint32) (QueryDirectRenderingCapableReply, error) {
	var reply QueryDirectRenderingCapableReply
	op, ok := c.Ext("XFree86-DRI")
	if !ok {
		return reply, errors.New("cannot issue request \"QueryDirectRenderingCapable\" using the uninitialized extension \"XFree86-DRI\". xf86dri.Register(xconn) must be called first.")
	}
	err := c.SendRecv(queryDirectRenderingCapableRequest(op, Screen), &reply)
	return reply, err
}

// QueryDirectRenderingCapableUnchecked sends an unchecked request.
func QueryDirectRenderingCapableUnchecked(c *xgb.XConn, Screen uint32) error {
	op, ok := c.Ext("XFree86-DRI")
	if !ok {
		return errors.New("cannot issue request \"QueryDirectRenderingCapable\" using the uninitialized extension \"XFree86-DRI\". xf86dri.Register(xconn) must be called first.")
	}
	return c.Send(queryDirectRenderingCapableRequest(op, Screen))
}

// QueryDirectRenderingCapableReply represents the data returned from a QueryDirectRenderingCapable request.
type QueryDirectRenderingCapableReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	IsCapable bool
}

// Unmarshal reads a byte slice into a QueryDirectRenderingCapableReply value.
func (v *QueryDirectRenderingCapableReply) Unmarshal(buf []byte) error {
	const size = 9
	if len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"QueryDirectRenderingCapableReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.IsCapable = (buf[b] == 1)
	b += 1

	return nil
}

// Write request to wire for QueryDirectRenderingCapable
// queryDirectRenderingCapableRequest writes a QueryDirectRenderingCapable request to a byte slice.
func queryDirectRenderingCapableRequest(opcode uint8, Screen uint32) []byte {
	const size = 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 1 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], Screen)
	b += 4

	return buf
}

// QueryVersion sends a checked request.
func QueryVersion(c *xgb.XConn) (QueryVersionReply, error) {
	var reply QueryVersionReply
	op, ok := c.Ext("XFree86-DRI")
	if !ok {
		return reply, errors.New("cannot issue request \"QueryVersion\" using the uninitialized extension \"XFree86-DRI\". xf86dri.Register(xconn) must be called first.")
	}
	err := c.SendRecv(queryVersionRequest(op), &reply)
	return reply, err
}

// QueryVersionUnchecked sends an unchecked request.
func QueryVersionUnchecked(c *xgb.XConn) error {
	op, ok := c.Ext("XFree86-DRI")
	if !ok {
		return errors.New("cannot issue request \"QueryVersion\" using the uninitialized extension \"XFree86-DRI\". xf86dri.Register(xconn) must be called first.")
	}
	return c.Send(queryVersionRequest(op))
}

// QueryVersionReply represents the data returned from a QueryVersion request.
type QueryVersionReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	DriMajorVersion uint16
	DriMinorVersion uint16
	DriMinorPatch   uint32
}

// Unmarshal reads a byte slice into a QueryVersionReply value.
func (v *QueryVersionReply) Unmarshal(buf []byte) error {
	const size = 16
	if len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"QueryVersionReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.DriMajorVersion = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.DriMinorVersion = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.DriMinorPatch = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	return nil
}

// Write request to wire for QueryVersion
// queryVersionRequest writes a QueryVersion request to a byte slice.
func queryVersionRequest(opcode uint8) []byte {
	const size = 4
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
