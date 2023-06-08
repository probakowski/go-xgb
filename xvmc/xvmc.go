// FILE GENERATED AUTOMATICALLY FROM "xvmc.xml"
package xvmc

import (
	"encoding/binary"
	"errors"
	"fmt"

	"codeberg.org/gruf/go-xgb"
	"codeberg.org/gruf/go-xgb/internal"
	"codeberg.org/gruf/go-xgb/xproto"
	"codeberg.org/gruf/go-xgb/xv"
)

const (
	// ExtName is the user-friendly name string of this X extension.
	ExtName = "XvMC"

	// ExtXName is the name string this extension is known by to the X server.
	ExtXName = "XVideo-MotionCompensation"
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

// Register will query the X server for XvMC extension support, and register relevant extension unmarshalers with the XConn.
func Register(xconn *xgb.XConn) error {
	// Query the X server for this extension
	reply, err := xproto.QueryExtension(xconn, uint16(len(ExtXName)), ExtXName)
	if err != nil {
		return fmt.Errorf("error querying X for \"XvMC\": %w", err)
	} else if !reply.Present {
		return fmt.Errorf("no extension named \"XvMC\" is known to the X server: reply=%+v", reply)
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

type Context uint32

func NewContextID(c *xgb.XConn) Context {
	id := c.NewXID()
	return Context(id)
}

type Subpicture uint32

func NewSubpictureID(c *xgb.XConn) Subpicture {
	id := c.NewXID()
	return Subpicture(id)
}

type Surface uint32

func NewSurfaceID(c *xgb.XConn) Surface {
	id := c.NewXID()
	return Surface(id)
}

type SurfaceInfo struct {
	Id                  Surface
	ChromaFormat        uint16
	Pad0                uint16
	MaxWidth            uint16
	MaxHeight           uint16
	SubpictureMaxWidth  uint16
	SubpictureMaxHeight uint16
	McType              uint32
	Flags               uint32
}

// SurfaceInfoRead reads a byte slice into a SurfaceInfo value.
func SurfaceInfoRead(buf []byte, v *SurfaceInfo) int {
	b := 0

	v.Id = Surface(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.ChromaFormat = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Pad0 = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.MaxWidth = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.MaxHeight = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.SubpictureMaxWidth = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.SubpictureMaxHeight = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.McType = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.Flags = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	return b
}

// SurfaceInfoReadList reads a byte slice into a list of SurfaceInfo values.
func SurfaceInfoReadList(buf []byte, dest []SurfaceInfo) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = SurfaceInfo{}
		b += SurfaceInfoRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a SurfaceInfo value to a byte slice.
func (v SurfaceInfo) Bytes() []byte {
	buf := make([]byte, 24)
	b := 0

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Id))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], v.ChromaFormat)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Pad0)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.MaxWidth)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.MaxHeight)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.SubpictureMaxWidth)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.SubpictureMaxHeight)
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], v.McType)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], v.Flags)
	b += 4

	return buf[:b]
}

// SurfaceInfoListBytes writes a list of SurfaceInfo values to a byte slice.
func SurfaceInfoListBytes(buf []byte, list []SurfaceInfo) int {
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

// CreateContext sends a checked request.
func CreateContext(c *xgb.XConn, ContextId Context, PortId xv.Port, SurfaceId Surface, Width uint16, Height uint16, Flags uint32) (CreateContextReply, error) {
	var reply CreateContextReply
	op, ok := c.Ext("XVideo-MotionCompensation")
	if !ok {
		return reply, errors.New("cannot issue request \"CreateContext\" using the uninitialized extension \"XVideo-MotionCompensation\". xvmc.Register(xconn) must be called first.")
	}
	err := c.SendRecv(createContextRequest(op, ContextId, PortId, SurfaceId, Width, Height, Flags), &reply)
	return reply, err
}

// CreateContextUnchecked sends an unchecked request.
func CreateContextUnchecked(c *xgb.XConn, ContextId Context, PortId xv.Port, SurfaceId Surface, Width uint16, Height uint16, Flags uint32) error {
	op, ok := c.Ext("XVideo-MotionCompensation")
	if !ok {
		return errors.New("cannot issue request \"CreateContext\" using the uninitialized extension \"XVideo-MotionCompensation\". xvmc.Register(xconn) must be called first.")
	}
	return c.Send(createContextRequest(op, ContextId, PortId, SurfaceId, Width, Height, Flags))
}

// CreateContextReply represents the data returned from a CreateContext request.
type CreateContextReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	WidthActual  uint16
	HeightActual uint16
	FlagsReturn  uint32
	// padding: 20 bytes
	PrivData []uint32 // size: internal.Pad4((int(Length) * 4))
}

// Unmarshal reads a byte slice into a CreateContextReply value.
func (v *CreateContextReply) Unmarshal(buf []byte) error {
	if size := (36 + internal.Pad4((int(v.Length) * 4))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"CreateContextReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.WidthActual = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.HeightActual = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.FlagsReturn = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 20 // padding

	v.PrivData = make([]uint32, v.Length)
	for i := 0; i < int(v.Length); i++ {
		v.PrivData[i] = binary.LittleEndian.Uint32(buf[b:])
		b += 4
	}

	return nil
}

// Write request to wire for CreateContext
// createContextRequest writes a CreateContext request to a byte slice.
func createContextRequest(opcode uint8, ContextId Context, PortId xv.Port, SurfaceId Surface, Width uint16, Height uint16, Flags uint32) []byte {
	size := 24
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 2 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(ContextId))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(PortId))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(SurfaceId))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], Width)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Height)
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], Flags)
	b += 4

	return buf
}

// CreateSubpicture sends a checked request.
func CreateSubpicture(c *xgb.XConn, SubpictureId Subpicture, Context Context, XvimageId uint32, Width uint16, Height uint16) (CreateSubpictureReply, error) {
	var reply CreateSubpictureReply
	op, ok := c.Ext("XVideo-MotionCompensation")
	if !ok {
		return reply, errors.New("cannot issue request \"CreateSubpicture\" using the uninitialized extension \"XVideo-MotionCompensation\". xvmc.Register(xconn) must be called first.")
	}
	err := c.SendRecv(createSubpictureRequest(op, SubpictureId, Context, XvimageId, Width, Height), &reply)
	return reply, err
}

// CreateSubpictureUnchecked sends an unchecked request.
func CreateSubpictureUnchecked(c *xgb.XConn, SubpictureId Subpicture, Context Context, XvimageId uint32, Width uint16, Height uint16) error {
	op, ok := c.Ext("XVideo-MotionCompensation")
	if !ok {
		return errors.New("cannot issue request \"CreateSubpicture\" using the uninitialized extension \"XVideo-MotionCompensation\". xvmc.Register(xconn) must be called first.")
	}
	return c.Send(createSubpictureRequest(op, SubpictureId, Context, XvimageId, Width, Height))
}

// CreateSubpictureReply represents the data returned from a CreateSubpicture request.
type CreateSubpictureReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	WidthActual       uint16
	HeightActual      uint16
	NumPaletteEntries uint16
	EntryBytes        uint16
	ComponentOrder    []byte // size: 4
	// padding: 12 bytes
	PrivData []uint32 // size: internal.Pad4((int(Length) * 4))
}

// Unmarshal reads a byte slice into a CreateSubpictureReply value.
func (v *CreateSubpictureReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.Length) * 4))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"CreateSubpictureReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.WidthActual = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.HeightActual = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.NumPaletteEntries = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.EntryBytes = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.ComponentOrder = make([]byte, 4)
	copy(v.ComponentOrder[:4], buf[b:])
	b += int(4)

	b += 12 // padding

	v.PrivData = make([]uint32, v.Length)
	for i := 0; i < int(v.Length); i++ {
		v.PrivData[i] = binary.LittleEndian.Uint32(buf[b:])
		b += 4
	}

	return nil
}

// Write request to wire for CreateSubpicture
// createSubpictureRequest writes a CreateSubpicture request to a byte slice.
func createSubpictureRequest(opcode uint8, SubpictureId Subpicture, Context Context, XvimageId uint32, Width uint16, Height uint16) []byte {
	size := 20
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 6 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(SubpictureId))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Context))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], XvimageId)
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], Width)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Height)
	b += 2

	return buf
}

// CreateSurface sends a checked request.
func CreateSurface(c *xgb.XConn, SurfaceId Surface, ContextId Context) (CreateSurfaceReply, error) {
	var reply CreateSurfaceReply
	op, ok := c.Ext("XVideo-MotionCompensation")
	if !ok {
		return reply, errors.New("cannot issue request \"CreateSurface\" using the uninitialized extension \"XVideo-MotionCompensation\". xvmc.Register(xconn) must be called first.")
	}
	err := c.SendRecv(createSurfaceRequest(op, SurfaceId, ContextId), &reply)
	return reply, err
}

// CreateSurfaceUnchecked sends an unchecked request.
func CreateSurfaceUnchecked(c *xgb.XConn, SurfaceId Surface, ContextId Context) error {
	op, ok := c.Ext("XVideo-MotionCompensation")
	if !ok {
		return errors.New("cannot issue request \"CreateSurface\" using the uninitialized extension \"XVideo-MotionCompensation\". xvmc.Register(xconn) must be called first.")
	}
	return c.Send(createSurfaceRequest(op, SurfaceId, ContextId))
}

// CreateSurfaceReply represents the data returned from a CreateSurface request.
type CreateSurfaceReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	// padding: 24 bytes
	PrivData []uint32 // size: internal.Pad4((int(Length) * 4))
}

// Unmarshal reads a byte slice into a CreateSurfaceReply value.
func (v *CreateSurfaceReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.Length) * 4))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"CreateSurfaceReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	b += 24 // padding

	v.PrivData = make([]uint32, v.Length)
	for i := 0; i < int(v.Length); i++ {
		v.PrivData[i] = binary.LittleEndian.Uint32(buf[b:])
		b += 4
	}

	return nil
}

// Write request to wire for CreateSurface
// createSurfaceRequest writes a CreateSurface request to a byte slice.
func createSurfaceRequest(opcode uint8, SurfaceId Surface, ContextId Context) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 4 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(SurfaceId))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(ContextId))
	b += 4

	return buf
}

// DestroyContext sends a checked request.
func DestroyContext(c *xgb.XConn, ContextId Context) error {
	op, ok := c.Ext("XVideo-MotionCompensation")
	if !ok {
		return errors.New("cannot issue request \"DestroyContext\" using the uninitialized extension \"XVideo-MotionCompensation\". xvmc.Register(xconn) must be called first.")
	}
	return c.SendRecv(destroyContextRequest(op, ContextId), nil)
}

// DestroyContextUnchecked sends an unchecked request.
func DestroyContextUnchecked(c *xgb.XConn, ContextId Context) error {
	op, ok := c.Ext("XVideo-MotionCompensation")
	if !ok {
		return errors.New("cannot issue request \"DestroyContext\" using the uninitialized extension \"XVideo-MotionCompensation\". xvmc.Register(xconn) must be called first.")
	}
	return c.Send(destroyContextRequest(op, ContextId))
}

// Write request to wire for DestroyContext
// destroyContextRequest writes a DestroyContext request to a byte slice.
func destroyContextRequest(opcode uint8, ContextId Context) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 3 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(ContextId))
	b += 4

	return buf
}

// DestroySubpicture sends a checked request.
func DestroySubpicture(c *xgb.XConn, SubpictureId Subpicture) error {
	op, ok := c.Ext("XVideo-MotionCompensation")
	if !ok {
		return errors.New("cannot issue request \"DestroySubpicture\" using the uninitialized extension \"XVideo-MotionCompensation\". xvmc.Register(xconn) must be called first.")
	}
	return c.SendRecv(destroySubpictureRequest(op, SubpictureId), nil)
}

// DestroySubpictureUnchecked sends an unchecked request.
func DestroySubpictureUnchecked(c *xgb.XConn, SubpictureId Subpicture) error {
	op, ok := c.Ext("XVideo-MotionCompensation")
	if !ok {
		return errors.New("cannot issue request \"DestroySubpicture\" using the uninitialized extension \"XVideo-MotionCompensation\". xvmc.Register(xconn) must be called first.")
	}
	return c.Send(destroySubpictureRequest(op, SubpictureId))
}

// Write request to wire for DestroySubpicture
// destroySubpictureRequest writes a DestroySubpicture request to a byte slice.
func destroySubpictureRequest(opcode uint8, SubpictureId Subpicture) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 7 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(SubpictureId))
	b += 4

	return buf
}

// DestroySurface sends a checked request.
func DestroySurface(c *xgb.XConn, SurfaceId Surface) error {
	op, ok := c.Ext("XVideo-MotionCompensation")
	if !ok {
		return errors.New("cannot issue request \"DestroySurface\" using the uninitialized extension \"XVideo-MotionCompensation\". xvmc.Register(xconn) must be called first.")
	}
	return c.SendRecv(destroySurfaceRequest(op, SurfaceId), nil)
}

// DestroySurfaceUnchecked sends an unchecked request.
func DestroySurfaceUnchecked(c *xgb.XConn, SurfaceId Surface) error {
	op, ok := c.Ext("XVideo-MotionCompensation")
	if !ok {
		return errors.New("cannot issue request \"DestroySurface\" using the uninitialized extension \"XVideo-MotionCompensation\". xvmc.Register(xconn) must be called first.")
	}
	return c.Send(destroySurfaceRequest(op, SurfaceId))
}

// Write request to wire for DestroySurface
// destroySurfaceRequest writes a DestroySurface request to a byte slice.
func destroySurfaceRequest(opcode uint8, SurfaceId Surface) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 5 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(SurfaceId))
	b += 4

	return buf
}

// ListSubpictureTypes sends a checked request.
func ListSubpictureTypes(c *xgb.XConn, PortId xv.Port, SurfaceId Surface) (ListSubpictureTypesReply, error) {
	var reply ListSubpictureTypesReply
	op, ok := c.Ext("XVideo-MotionCompensation")
	if !ok {
		return reply, errors.New("cannot issue request \"ListSubpictureTypes\" using the uninitialized extension \"XVideo-MotionCompensation\". xvmc.Register(xconn) must be called first.")
	}
	err := c.SendRecv(listSubpictureTypesRequest(op, PortId, SurfaceId), &reply)
	return reply, err
}

// ListSubpictureTypesUnchecked sends an unchecked request.
func ListSubpictureTypesUnchecked(c *xgb.XConn, PortId xv.Port, SurfaceId Surface) error {
	op, ok := c.Ext("XVideo-MotionCompensation")
	if !ok {
		return errors.New("cannot issue request \"ListSubpictureTypes\" using the uninitialized extension \"XVideo-MotionCompensation\". xvmc.Register(xconn) must be called first.")
	}
	return c.Send(listSubpictureTypesRequest(op, PortId, SurfaceId))
}

// ListSubpictureTypesReply represents the data returned from a ListSubpictureTypes request.
type ListSubpictureTypesReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Num uint32
	// padding: 20 bytes
	Types []xv.ImageFormatInfo // size: xv.ImageFormatInfoListSize(Types)
}

// Unmarshal reads a byte slice into a ListSubpictureTypesReply value.
func (v *ListSubpictureTypesReply) Unmarshal(buf []byte) error {
	if size := (32 + xv.ImageFormatInfoListSize(v.Types)); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"ListSubpictureTypesReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Num = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 20 // padding

	v.Types = make([]xv.ImageFormatInfo, v.Num)
	b += xv.ImageFormatInfoReadList(buf[b:], v.Types)

	return nil
}

// Write request to wire for ListSubpictureTypes
// listSubpictureTypesRequest writes a ListSubpictureTypes request to a byte slice.
func listSubpictureTypesRequest(opcode uint8, PortId xv.Port, SurfaceId Surface) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 8 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(PortId))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(SurfaceId))
	b += 4

	return buf
}

// ListSurfaceTypes sends a checked request.
func ListSurfaceTypes(c *xgb.XConn, PortId xv.Port) (ListSurfaceTypesReply, error) {
	var reply ListSurfaceTypesReply
	op, ok := c.Ext("XVideo-MotionCompensation")
	if !ok {
		return reply, errors.New("cannot issue request \"ListSurfaceTypes\" using the uninitialized extension \"XVideo-MotionCompensation\". xvmc.Register(xconn) must be called first.")
	}
	err := c.SendRecv(listSurfaceTypesRequest(op, PortId), &reply)
	return reply, err
}

// ListSurfaceTypesUnchecked sends an unchecked request.
func ListSurfaceTypesUnchecked(c *xgb.XConn, PortId xv.Port) error {
	op, ok := c.Ext("XVideo-MotionCompensation")
	if !ok {
		return errors.New("cannot issue request \"ListSurfaceTypes\" using the uninitialized extension \"XVideo-MotionCompensation\". xvmc.Register(xconn) must be called first.")
	}
	return c.Send(listSurfaceTypesRequest(op, PortId))
}

// ListSurfaceTypesReply represents the data returned from a ListSurfaceTypes request.
type ListSurfaceTypesReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Num uint32
	// padding: 20 bytes
	Surfaces []SurfaceInfo // size: internal.Pad4((int(Num) * 24))
}

// Unmarshal reads a byte slice into a ListSurfaceTypesReply value.
func (v *ListSurfaceTypesReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.Num) * 24))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"ListSurfaceTypesReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Num = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 20 // padding

	v.Surfaces = make([]SurfaceInfo, v.Num)
	b += SurfaceInfoReadList(buf[b:], v.Surfaces)

	return nil
}

// Write request to wire for ListSurfaceTypes
// listSurfaceTypesRequest writes a ListSurfaceTypes request to a byte slice.
func listSurfaceTypesRequest(opcode uint8, PortId xv.Port) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 1 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(PortId))
	b += 4

	return buf
}

// QueryVersion sends a checked request.
func QueryVersion(c *xgb.XConn) (QueryVersionReply, error) {
	var reply QueryVersionReply
	op, ok := c.Ext("XVideo-MotionCompensation")
	if !ok {
		return reply, errors.New("cannot issue request \"QueryVersion\" using the uninitialized extension \"XVideo-MotionCompensation\". xvmc.Register(xconn) must be called first.")
	}
	err := c.SendRecv(queryVersionRequest(op), &reply)
	return reply, err
}

// QueryVersionUnchecked sends an unchecked request.
func QueryVersionUnchecked(c *xgb.XConn) error {
	op, ok := c.Ext("XVideo-MotionCompensation")
	if !ok {
		return errors.New("cannot issue request \"QueryVersion\" using the uninitialized extension \"XVideo-MotionCompensation\". xvmc.Register(xconn) must be called first.")
	}
	return c.Send(queryVersionRequest(op))
}

// QueryVersionReply represents the data returned from a QueryVersion request.
type QueryVersionReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Major uint32
	Minor uint32
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

	v.Major = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.Minor = binary.LittleEndian.Uint32(buf[b:])
	b += 4

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
