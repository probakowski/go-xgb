// FILE GENERATED AUTOMATICALLY FROM "xv.xml"
package xv

import (
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"codeberg.org/gruf/go-xgb"
	"codeberg.org/gruf/go-xgb/internal"
	"codeberg.org/gruf/go-xgb/shm"
	"codeberg.org/gruf/go-xgb/xproto"
)

const (
	// ExtName is the user-friendly name string of this X extension.
	ExtName = "Xv"

	// ExtXName is the name string this extension is known by to the X server.
	ExtXName = "XVideo"
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
		return fmt.Errorf("error querying X for \"Xv\": %w", err)
	} else if !reply.Present {
		return fmt.Errorf("no extension named \"Xv\" is known to the X server: reply=%+v", reply)
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

type AdaptorInfo struct {
	BaseId     Port
	NameSize   uint16
	NumPorts   uint16
	NumFormats uint16
	Type       byte
	// padding: 1 bytes
	Name string // size: internal.Pad4((int(NameSize) * 1))
	// padding: 0 bytes
	Formats []Format // size: internal.Pad4((int(NumFormats) * 8))
}

// AdaptorInfoRead reads a byte slice into a AdaptorInfo value.
func AdaptorInfoRead(buf []byte, v *AdaptorInfo) int {
	b := 0

	v.BaseId = Port(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.NameSize = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.NumPorts = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.NumFormats = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Type = buf[b]
	b += 1

	b += 1 // padding

	{
		byteString := make([]byte, v.NameSize)
		copy(byteString[:v.NameSize], buf[b:])
		v.Name = string(byteString)
		b += int(v.NameSize)
	}

	b += 0 // padding

	v.Formats = make([]Format, v.NumFormats)
	b += FormatReadList(buf[b:], v.Formats)

	return b
}

// AdaptorInfoReadList reads a byte slice into a list of AdaptorInfo values.
func AdaptorInfoReadList(buf []byte, dest []AdaptorInfo) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = AdaptorInfo{}
		b += AdaptorInfoRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a AdaptorInfo value to a byte slice.
func (v AdaptorInfo) Bytes() []byte {
	buf := make([]byte, (((12 + internal.Pad4((int(v.NameSize) * 1))) + 0) + internal.Pad4((int(v.NumFormats) * 8))))
	b := 0

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.BaseId))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], v.NameSize)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.NumPorts)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.NumFormats)
	b += 2

	buf[b] = v.Type
	b += 1

	b += 1 // padding

	copy(buf[b:], v.Name[:v.NameSize])
	b += int(v.NameSize)

	b += 0 // padding

	b += FormatListBytes(buf[b:], v.Formats)

	return buf[:b]
}

// AdaptorInfoListBytes writes a list of AdaptorInfo values to a byte slice.
func AdaptorInfoListBytes(buf []byte, list []AdaptorInfo) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

// AdaptorInfoListSize computes the size (bytes) of a list of AdaptorInfo values.
func AdaptorInfoListSize(list []AdaptorInfo) int {
	size := 0
	for _, item := range list {
		size += (((12 + internal.Pad4((int(item.NameSize) * 1))) + 0) + internal.Pad4((int(item.NumFormats) * 8)))
	}
	return size
}

const (
	AttributeFlagGettable = 1
	AttributeFlagSettable = 2
)

type AttributeInfo struct {
	Flags uint32
	Min   int32
	Max   int32
	Size  uint32
	Name  string // size: internal.Pad4((int(Size) * 1))
	// padding: 0 bytes
}

// AttributeInfoRead reads a byte slice into a AttributeInfo value.
func AttributeInfoRead(buf []byte, v *AttributeInfo) int {
	b := 0

	v.Flags = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.Min = int32(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Max = int32(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Size = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	{
		byteString := make([]byte, v.Size)
		copy(byteString[:v.Size], buf[b:])
		v.Name = string(byteString)
		b += int(v.Size)
	}

	b += 0 // padding

	return b
}

// AttributeInfoReadList reads a byte slice into a list of AttributeInfo values.
func AttributeInfoReadList(buf []byte, dest []AttributeInfo) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = AttributeInfo{}
		b += AttributeInfoRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a AttributeInfo value to a byte slice.
func (v AttributeInfo) Bytes() []byte {
	buf := make([]byte, ((16 + internal.Pad4((int(v.Size) * 1))) + 0))
	b := 0

	binary.LittleEndian.PutUint32(buf[b:], v.Flags)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Min))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Max))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], v.Size)
	b += 4

	copy(buf[b:], v.Name[:v.Size])
	b += int(v.Size)

	b += 0 // padding

	return buf[:b]
}

// AttributeInfoListBytes writes a list of AttributeInfo values to a byte slice.
func AttributeInfoListBytes(buf []byte, list []AttributeInfo) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

// AttributeInfoListSize computes the size (bytes) of a list of AttributeInfo values.
func AttributeInfoListSize(list []AttributeInfo) int {
	size := 0
	for _, item := range list {
		size += ((16 + internal.Pad4((int(item.Size) * 1))) + 0)
	}
	return size
}

// BadBadControl is the error number for a BadBadControl.
const BadBadControl = 2

type BadControlError struct {
	Sequence uint16
	NiceName string
}

// UnmarshalBadControlError constructs a BadControlError value that implements xgb.Error from a byte slice.
func UnmarshalBadControlError(buf []byte) (xgb.XError, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"BadControlError\"", len(buf))
	}

	v := &BadControlError{}
	v.NiceName = "BadControl"

	b := 1 // skip error determinant
	b += 1 // don't read error number

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return v, nil
}

// SeqID returns the sequence id attached to the BadBadControl error.
// This is mostly used internally.
func (err *BadControlError) SeqID() uint16 {
	return err.Sequence
}

// BadID returns the 'BadValue' number if one exists for the BadBadControl error. If no bad value exists, 0 is returned.
func (err *BadControlError) BadID() uint32 {
	return 0
}

// Error returns a rudimentary string representation of the BadBadControl error.
func (err *BadControlError) Error() string {
	var buf strings.Builder

	buf.WriteString("BadBadControl{")
	buf.WriteString("NiceName: " + err.NiceName)
	buf.WriteByte(' ')
	buf.WriteString("Sequence: " + strconv.FormatUint(uint64(err.Sequence), 10))

	buf.WriteByte('}')

	return buf.String()
}

func init() { registerError(2, UnmarshalBadControlError) }

// BadBadEncoding is the error number for a BadBadEncoding.
const BadBadEncoding = 1

type BadEncodingError struct {
	Sequence uint16
	NiceName string
}

// UnmarshalBadEncodingError constructs a BadEncodingError value that implements xgb.Error from a byte slice.
func UnmarshalBadEncodingError(buf []byte) (xgb.XError, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"BadEncodingError\"", len(buf))
	}

	v := &BadEncodingError{}
	v.NiceName = "BadEncoding"

	b := 1 // skip error determinant
	b += 1 // don't read error number

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return v, nil
}

// SeqID returns the sequence id attached to the BadBadEncoding error.
// This is mostly used internally.
func (err *BadEncodingError) SeqID() uint16 {
	return err.Sequence
}

// BadID returns the 'BadValue' number if one exists for the BadBadEncoding error. If no bad value exists, 0 is returned.
func (err *BadEncodingError) BadID() uint32 {
	return 0
}

// Error returns a rudimentary string representation of the BadBadEncoding error.
func (err *BadEncodingError) Error() string {
	var buf strings.Builder

	buf.WriteString("BadBadEncoding{")
	buf.WriteString("NiceName: " + err.NiceName)
	buf.WriteByte(' ')
	buf.WriteString("Sequence: " + strconv.FormatUint(uint64(err.Sequence), 10))

	buf.WriteByte('}')

	return buf.String()
}

func init() { registerError(1, UnmarshalBadEncodingError) }

// BadBadPort is the error number for a BadBadPort.
const BadBadPort = 0

type BadPortError struct {
	Sequence uint16
	NiceName string
}

// UnmarshalBadPortError constructs a BadPortError value that implements xgb.Error from a byte slice.
func UnmarshalBadPortError(buf []byte) (xgb.XError, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"BadPortError\"", len(buf))
	}

	v := &BadPortError{}
	v.NiceName = "BadPort"

	b := 1 // skip error determinant
	b += 1 // don't read error number

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return v, nil
}

// SeqID returns the sequence id attached to the BadBadPort error.
// This is mostly used internally.
func (err *BadPortError) SeqID() uint16 {
	return err.Sequence
}

// BadID returns the 'BadValue' number if one exists for the BadBadPort error. If no bad value exists, 0 is returned.
func (err *BadPortError) BadID() uint32 {
	return 0
}

// Error returns a rudimentary string representation of the BadBadPort error.
func (err *BadPortError) Error() string {
	var buf strings.Builder

	buf.WriteString("BadBadPort{")
	buf.WriteString("NiceName: " + err.NiceName)
	buf.WriteByte(' ')
	buf.WriteString("Sequence: " + strconv.FormatUint(uint64(err.Sequence), 10))

	buf.WriteByte('}')

	return buf.String()
}

func init() { registerError(0, UnmarshalBadPortError) }

type Encoding uint32

func NewEncodingID(c *xgb.XConn) Encoding {
	id := c.NewXID()
	return Encoding(id)
}

type EncodingInfo struct {
	Encoding Encoding
	NameSize uint16
	Width    uint16
	Height   uint16
	// padding: 2 bytes
	Rate Rational
	Name string // size: internal.Pad4((int(NameSize) * 1))
	// padding: 0 bytes
}

// EncodingInfoRead reads a byte slice into a EncodingInfo value.
func EncodingInfoRead(buf []byte, v *EncodingInfo) int {
	b := 0

	v.Encoding = Encoding(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.NameSize = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Width = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Height = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	b += 2 // padding

	v.Rate = Rational{}
	b += RationalRead(buf[b:], &v.Rate)

	{
		byteString := make([]byte, v.NameSize)
		copy(byteString[:v.NameSize], buf[b:])
		v.Name = string(byteString)
		b += int(v.NameSize)
	}

	b += 0 // padding

	return b
}

// EncodingInfoReadList reads a byte slice into a list of EncodingInfo values.
func EncodingInfoReadList(buf []byte, dest []EncodingInfo) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = EncodingInfo{}
		b += EncodingInfoRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a EncodingInfo value to a byte slice.
func (v EncodingInfo) Bytes() []byte {
	buf := make([]byte, ((20 + internal.Pad4((int(v.NameSize) * 1))) + 0))
	b := 0

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Encoding))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], v.NameSize)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Width)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Height)
	b += 2

	b += 2 // padding

	{
		structBytes := v.Rate.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	copy(buf[b:], v.Name[:v.NameSize])
	b += int(v.NameSize)

	b += 0 // padding

	return buf[:b]
}

// EncodingInfoListBytes writes a list of EncodingInfo values to a byte slice.
func EncodingInfoListBytes(buf []byte, list []EncodingInfo) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

// EncodingInfoListSize computes the size (bytes) of a list of EncodingInfo values.
func EncodingInfoListSize(list []EncodingInfo) int {
	size := 0
	for _, item := range list {
		size += ((20 + internal.Pad4((int(item.NameSize) * 1))) + 0)
	}
	return size
}

type Format struct {
	Visual xproto.Visualid
	Depth  byte
	// padding: 3 bytes
}

// FormatRead reads a byte slice into a Format value.
func FormatRead(buf []byte, v *Format) int {
	b := 0

	v.Visual = xproto.Visualid(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Depth = buf[b]
	b += 1

	b += 3 // padding

	return b
}

// FormatReadList reads a byte slice into a list of Format values.
func FormatReadList(buf []byte, dest []Format) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Format{}
		b += FormatRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a Format value to a byte slice.
func (v Format) Bytes() []byte {
	buf := make([]byte, 8)
	b := 0

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Visual))
	b += 4

	buf[b] = v.Depth
	b += 1

	b += 3 // padding

	return buf[:b]
}

// FormatListBytes writes a list of Format values to a byte slice.
func FormatListBytes(buf []byte, list []Format) int {
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
	GrabPortStatusSuccess        = 0
	GrabPortStatusBadExtension   = 1
	GrabPortStatusAlreadyGrabbed = 2
	GrabPortStatusInvalidTime    = 3
	GrabPortStatusBadReply       = 4
	GrabPortStatusBadAlloc       = 5
)

type Image struct {
	Id        uint32
	Width     uint16
	Height    uint16
	DataSize  uint32
	NumPlanes uint32
	Pitches   []uint32 // size: internal.Pad4((int(NumPlanes) * 4))
	// alignment gap to multiple of 4
	Offsets []uint32 // size: internal.Pad4((int(NumPlanes) * 4))
	Data    []byte   // size: internal.Pad4((int(DataSize) * 1))
}

// ImageRead reads a byte slice into a Image value.
func ImageRead(buf []byte, v *Image) int {
	b := 0

	v.Id = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.Width = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Height = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.DataSize = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.NumPlanes = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.Pitches = make([]uint32, v.NumPlanes)
	for i := 0; i < int(v.NumPlanes); i++ {
		v.Pitches[i] = binary.LittleEndian.Uint32(buf[b:])
		b += 4
	}

	b = (b + 3) & ^3 // alignment gap

	v.Offsets = make([]uint32, v.NumPlanes)
	for i := 0; i < int(v.NumPlanes); i++ {
		v.Offsets[i] = binary.LittleEndian.Uint32(buf[b:])
		b += 4
	}

	v.Data = make([]byte, v.DataSize)
	copy(v.Data[:v.DataSize], buf[b:])
	b += int(v.DataSize)

	return b
}

// ImageReadList reads a byte slice into a list of Image values.
func ImageReadList(buf []byte, dest []Image) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Image{}
		b += ImageRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a Image value to a byte slice.
func (v Image) Bytes() []byte {
	buf := make([]byte, ((((16 + internal.Pad4((int(v.NumPlanes) * 4))) + 4) + internal.Pad4((int(v.NumPlanes) * 4))) + internal.Pad4((int(v.DataSize) * 1))))
	b := 0

	binary.LittleEndian.PutUint32(buf[b:], v.Id)
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], v.Width)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Height)
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], v.DataSize)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], v.NumPlanes)
	b += 4

	for i := 0; i < int(v.NumPlanes); i++ {
		binary.LittleEndian.PutUint32(buf[b:], v.Pitches[i])
		b += 4
	}

	b = (b + 3) & ^3 // alignment gap

	for i := 0; i < int(v.NumPlanes); i++ {
		binary.LittleEndian.PutUint32(buf[b:], v.Offsets[i])
		b += 4
	}

	copy(buf[b:], v.Data[:v.DataSize])
	b += int(v.DataSize)

	return buf[:b]
}

// ImageListBytes writes a list of Image values to a byte slice.
func ImageListBytes(buf []byte, list []Image) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

// ImageListSize computes the size (bytes) of a list of Image values.
func ImageListSize(list []Image) int {
	size := 0
	for _, item := range list {
		size += ((((16 + internal.Pad4((int(item.NumPlanes) * 4))) + 4) + internal.Pad4((int(item.NumPlanes) * 4))) + internal.Pad4((int(item.DataSize) * 1)))
	}
	return size
}

type ImageFormatInfo struct {
	Id        uint32
	Type      byte
	ByteOrder byte
	// padding: 2 bytes
	Guid      []byte // size: 16
	Bpp       byte
	NumPlanes byte
	// padding: 2 bytes
	Depth byte
	// padding: 3 bytes
	RedMask   uint32
	GreenMask uint32
	BlueMask  uint32
	Format    byte
	// padding: 3 bytes
	YSampleBits    uint32
	USampleBits    uint32
	VSampleBits    uint32
	VhorzYPeriod   uint32
	VhorzUPeriod   uint32
	VhorzVPeriod   uint32
	VvertYPeriod   uint32
	VvertUPeriod   uint32
	VvertVPeriod   uint32
	VcompOrder     []byte // size: 32
	VscanlineOrder byte
	// padding: 11 bytes
}

// ImageFormatInfoRead reads a byte slice into a ImageFormatInfo value.
func ImageFormatInfoRead(buf []byte, v *ImageFormatInfo) int {
	b := 0

	v.Id = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.Type = buf[b]
	b += 1

	v.ByteOrder = buf[b]
	b += 1

	b += 2 // padding

	v.Guid = make([]byte, 16)
	copy(v.Guid[:16], buf[b:])
	b += int(16)

	v.Bpp = buf[b]
	b += 1

	v.NumPlanes = buf[b]
	b += 1

	b += 2 // padding

	v.Depth = buf[b]
	b += 1

	b += 3 // padding

	v.RedMask = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.GreenMask = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.BlueMask = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.Format = buf[b]
	b += 1

	b += 3 // padding

	v.YSampleBits = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.USampleBits = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.VSampleBits = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.VhorzYPeriod = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.VhorzUPeriod = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.VhorzVPeriod = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.VvertYPeriod = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.VvertUPeriod = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.VvertVPeriod = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.VcompOrder = make([]byte, 32)
	copy(v.VcompOrder[:32], buf[b:])
	b += int(32)

	v.VscanlineOrder = buf[b]
	b += 1

	b += 11 // padding

	return b
}

// ImageFormatInfoReadList reads a byte slice into a list of ImageFormatInfo values.
func ImageFormatInfoReadList(buf []byte, dest []ImageFormatInfo) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = ImageFormatInfo{}
		b += ImageFormatInfoRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a ImageFormatInfo value to a byte slice.
func (v ImageFormatInfo) Bytes() []byte {
	buf := make([]byte, 128)
	b := 0

	binary.LittleEndian.PutUint32(buf[b:], v.Id)
	b += 4

	buf[b] = v.Type
	b += 1

	buf[b] = v.ByteOrder
	b += 1

	b += 2 // padding

	copy(buf[b:], v.Guid[:16])
	b += int(16)

	buf[b] = v.Bpp
	b += 1

	buf[b] = v.NumPlanes
	b += 1

	b += 2 // padding

	buf[b] = v.Depth
	b += 1

	b += 3 // padding

	binary.LittleEndian.PutUint32(buf[b:], v.RedMask)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], v.GreenMask)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], v.BlueMask)
	b += 4

	buf[b] = v.Format
	b += 1

	b += 3 // padding

	binary.LittleEndian.PutUint32(buf[b:], v.YSampleBits)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], v.USampleBits)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], v.VSampleBits)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], v.VhorzYPeriod)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], v.VhorzUPeriod)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], v.VhorzVPeriod)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], v.VvertYPeriod)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], v.VvertUPeriod)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], v.VvertVPeriod)
	b += 4

	copy(buf[b:], v.VcompOrder[:32])
	b += int(32)

	buf[b] = v.VscanlineOrder
	b += 1

	b += 11 // padding

	return buf[:b]
}

// ImageFormatInfoListBytes writes a list of ImageFormatInfo values to a byte slice.
func ImageFormatInfoListBytes(buf []byte, list []ImageFormatInfo) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

// ImageFormatInfoListSize computes the size (bytes) of a list of ImageFormatInfo values.
func ImageFormatInfoListSize(list []ImageFormatInfo) int {
	size := 0
	for _ = range list {
		size += 128
	}
	return size
}

const (
	ImageFormatInfoFormatPacked = 0
	ImageFormatInfoFormatPlanar = 1
)

const (
	ImageFormatInfoTypeRgb = 0
	ImageFormatInfoTypeYuv = 1
)

type Port uint32

func NewPortID(c *xgb.XConn) Port {
	id := c.NewXID()
	return Port(id)
}

// PortNotify is the event number for a PortNotifyEvent.
const PortNotify = 1

type PortNotifyEvent struct {
	Sequence uint16
	// padding: 1 bytes
	Time      xproto.Timestamp
	Port      Port
	Attribute xproto.Atom
	Value     int32
}

// UnmarshalPortNotifyEvent constructs a PortNotifyEvent value that implements xgb.Event from a byte slice.
func UnmarshalPortNotifyEvent(buf []byte) (xgb.XEvent, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"PortNotifyEvent\"", len(buf))
	}

	v := &PortNotifyEvent{}
	b := 1 // don't read event number

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Time = xproto.Timestamp(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Port = Port(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Attribute = xproto.Atom(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Value = int32(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	return v, nil
}

// Bytes writes a PortNotifyEvent value to a byte slice.
func (v *PortNotifyEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 1
	b += 1

	b += 1 // padding

	b += 2 // skip sequence number

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Time))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Port))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Attribute))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Value))
	b += 4

	return buf
}

// SeqID returns the sequence id attached to the PortNotify event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v *PortNotifyEvent) SeqID() uint16 {
	return v.Sequence
}

func init() { registerEvent(1, UnmarshalPortNotifyEvent) }

type Rational struct {
	Numerator   int32
	Denominator int32
}

// RationalRead reads a byte slice into a Rational value.
func RationalRead(buf []byte, v *Rational) int {
	b := 0

	v.Numerator = int32(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Denominator = int32(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	return b
}

// RationalReadList reads a byte slice into a list of Rational values.
func RationalReadList(buf []byte, dest []Rational) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Rational{}
		b += RationalRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a Rational value to a byte slice.
func (v Rational) Bytes() []byte {
	buf := make([]byte, 8)
	b := 0

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Numerator))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Denominator))
	b += 4

	return buf[:b]
}

// RationalListBytes writes a list of Rational values to a byte slice.
func RationalListBytes(buf []byte, list []Rational) int {
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
	ScanlineOrderTopToBottom = 0
	ScanlineOrderBottomToTop = 1
)

const (
	TypeInputMask  = 1
	TypeOutputMask = 2
	TypeVideoMask  = 4
	TypeStillMask  = 8
	TypeImageMask  = 16
)

// VideoNotify is the event number for a VideoNotifyEvent.
const VideoNotify = 0

type VideoNotifyEvent struct {
	Sequence uint16
	Reason   byte
	Time     xproto.Timestamp
	Drawable xproto.Drawable
	Port     Port
}

// UnmarshalVideoNotifyEvent constructs a VideoNotifyEvent value that implements xgb.Event from a byte slice.
func UnmarshalVideoNotifyEvent(buf []byte) (xgb.XEvent, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"VideoNotifyEvent\"", len(buf))
	}

	v := &VideoNotifyEvent{}
	b := 1 // don't read event number

	v.Reason = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Time = xproto.Timestamp(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Drawable = xproto.Drawable(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Port = Port(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	return v, nil
}

// Bytes writes a VideoNotifyEvent value to a byte slice.
func (v *VideoNotifyEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 0
	b += 1

	buf[b] = v.Reason
	b += 1

	b += 2 // skip sequence number

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Time))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Drawable))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Port))
	b += 4

	return buf
}

// SeqID returns the sequence id attached to the VideoNotify event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v *VideoNotifyEvent) SeqID() uint16 {
	return v.Sequence
}

func init() { registerEvent(0, UnmarshalVideoNotifyEvent) }

const (
	VideoNotifyReasonStarted   = 0
	VideoNotifyReasonStopped   = 1
	VideoNotifyReasonBusy      = 2
	VideoNotifyReasonPreempted = 3
	VideoNotifyReasonHardError = 4
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

// GetPortAttribute sends a checked request.
func GetPortAttribute(c *xgb.XConn, Port Port, Attribute xproto.Atom) (GetPortAttributeReply, error) {
	var reply GetPortAttributeReply
	op, ok := c.Ext("XVideo")
	if !ok {
		return reply, errors.New("cannot issue request \"GetPortAttribute\" using the uninitialized extension \"XVideo\". xv.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getPortAttributeRequest(op, Port, Attribute), &reply)
	return reply, err
}

// GetPortAttributeUnchecked sends an unchecked request.
func GetPortAttributeUnchecked(c *xgb.XConn, Port Port, Attribute xproto.Atom) error {
	op, ok := c.Ext("XVideo")
	if !ok {
		return errors.New("cannot issue request \"GetPortAttribute\" using the uninitialized extension \"XVideo\". xv.Register(xconn) must be called first.")
	}
	return c.Send(getPortAttributeRequest(op, Port, Attribute))
}

// GetPortAttributeReply represents the data returned from a GetPortAttribute request.
type GetPortAttributeReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Value int32
}

// Unmarshal reads a byte slice into a GetPortAttributeReply value.
func (v *GetPortAttributeReply) Unmarshal(buf []byte) error {
	if size := 12; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetPortAttributeReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Value = int32(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	return nil
}

// Write request to wire for GetPortAttribute
// getPortAttributeRequest writes a GetPortAttribute request to a byte slice.
func getPortAttributeRequest(opcode uint8, Port Port, Attribute xproto.Atom) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 14 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Port))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Attribute))
	b += 4

	return buf
}

// GetStill sends a checked request.
func GetStill(c *xgb.XConn, Port Port, Drawable xproto.Drawable, Gc xproto.Gcontext, VidX int16, VidY int16, VidW uint16, VidH uint16, DrwX int16, DrwY int16, DrwW uint16, DrwH uint16) error {
	op, ok := c.Ext("XVideo")
	if !ok {
		return errors.New("cannot issue request \"GetStill\" using the uninitialized extension \"XVideo\". xv.Register(xconn) must be called first.")
	}
	return c.SendRecv(getStillRequest(op, Port, Drawable, Gc, VidX, VidY, VidW, VidH, DrwX, DrwY, DrwW, DrwH), nil)
}

// GetStillUnchecked sends an unchecked request.
func GetStillUnchecked(c *xgb.XConn, Port Port, Drawable xproto.Drawable, Gc xproto.Gcontext, VidX int16, VidY int16, VidW uint16, VidH uint16, DrwX int16, DrwY int16, DrwW uint16, DrwH uint16) error {
	op, ok := c.Ext("XVideo")
	if !ok {
		return errors.New("cannot issue request \"GetStill\" using the uninitialized extension \"XVideo\". xv.Register(xconn) must be called first.")
	}
	return c.Send(getStillRequest(op, Port, Drawable, Gc, VidX, VidY, VidW, VidH, DrwX, DrwY, DrwW, DrwH))
}

// Write request to wire for GetStill
// getStillRequest writes a GetStill request to a byte slice.
func getStillRequest(opcode uint8, Port Port, Drawable xproto.Drawable, Gc xproto.Gcontext, VidX int16, VidY int16, VidW uint16, VidH uint16, DrwX int16, DrwY int16, DrwW uint16, DrwH uint16) []byte {
	size := 32
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 8 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Port))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Drawable))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Gc))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(VidX))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(VidY))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], VidW)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], VidH)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(DrwX))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(DrwY))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], DrwW)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], DrwH)
	b += 2

	return buf
}

// GetVideo sends a checked request.
func GetVideo(c *xgb.XConn, Port Port, Drawable xproto.Drawable, Gc xproto.Gcontext, VidX int16, VidY int16, VidW uint16, VidH uint16, DrwX int16, DrwY int16, DrwW uint16, DrwH uint16) error {
	op, ok := c.Ext("XVideo")
	if !ok {
		return errors.New("cannot issue request \"GetVideo\" using the uninitialized extension \"XVideo\". xv.Register(xconn) must be called first.")
	}
	return c.SendRecv(getVideoRequest(op, Port, Drawable, Gc, VidX, VidY, VidW, VidH, DrwX, DrwY, DrwW, DrwH), nil)
}

// GetVideoUnchecked sends an unchecked request.
func GetVideoUnchecked(c *xgb.XConn, Port Port, Drawable xproto.Drawable, Gc xproto.Gcontext, VidX int16, VidY int16, VidW uint16, VidH uint16, DrwX int16, DrwY int16, DrwW uint16, DrwH uint16) error {
	op, ok := c.Ext("XVideo")
	if !ok {
		return errors.New("cannot issue request \"GetVideo\" using the uninitialized extension \"XVideo\". xv.Register(xconn) must be called first.")
	}
	return c.Send(getVideoRequest(op, Port, Drawable, Gc, VidX, VidY, VidW, VidH, DrwX, DrwY, DrwW, DrwH))
}

// Write request to wire for GetVideo
// getVideoRequest writes a GetVideo request to a byte slice.
func getVideoRequest(opcode uint8, Port Port, Drawable xproto.Drawable, Gc xproto.Gcontext, VidX int16, VidY int16, VidW uint16, VidH uint16, DrwX int16, DrwY int16, DrwW uint16, DrwH uint16) []byte {
	size := 32
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 7 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Port))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Drawable))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Gc))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(VidX))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(VidY))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], VidW)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], VidH)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(DrwX))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(DrwY))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], DrwW)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], DrwH)
	b += 2

	return buf
}

// GrabPort sends a checked request.
func GrabPort(c *xgb.XConn, Port Port, Time xproto.Timestamp) (GrabPortReply, error) {
	var reply GrabPortReply
	op, ok := c.Ext("XVideo")
	if !ok {
		return reply, errors.New("cannot issue request \"GrabPort\" using the uninitialized extension \"XVideo\". xv.Register(xconn) must be called first.")
	}
	err := c.SendRecv(grabPortRequest(op, Port, Time), &reply)
	return reply, err
}

// GrabPortUnchecked sends an unchecked request.
func GrabPortUnchecked(c *xgb.XConn, Port Port, Time xproto.Timestamp) error {
	op, ok := c.Ext("XVideo")
	if !ok {
		return errors.New("cannot issue request \"GrabPort\" using the uninitialized extension \"XVideo\". xv.Register(xconn) must be called first.")
	}
	return c.Send(grabPortRequest(op, Port, Time))
}

// GrabPortReply represents the data returned from a GrabPort request.
type GrabPortReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	Result   byte
}

// Unmarshal reads a byte slice into a GrabPortReply value.
func (v *GrabPortReply) Unmarshal(buf []byte) error {
	if size := 8; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GrabPortReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	v.Result = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	return nil
}

// Write request to wire for GrabPort
// grabPortRequest writes a GrabPort request to a byte slice.
func grabPortRequest(opcode uint8, Port Port, Time xproto.Timestamp) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 3 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Port))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Time))
	b += 4

	return buf
}

// ListImageFormats sends a checked request.
func ListImageFormats(c *xgb.XConn, Port Port) (ListImageFormatsReply, error) {
	var reply ListImageFormatsReply
	op, ok := c.Ext("XVideo")
	if !ok {
		return reply, errors.New("cannot issue request \"ListImageFormats\" using the uninitialized extension \"XVideo\". xv.Register(xconn) must be called first.")
	}
	err := c.SendRecv(listImageFormatsRequest(op, Port), &reply)
	return reply, err
}

// ListImageFormatsUnchecked sends an unchecked request.
func ListImageFormatsUnchecked(c *xgb.XConn, Port Port) error {
	op, ok := c.Ext("XVideo")
	if !ok {
		return errors.New("cannot issue request \"ListImageFormats\" using the uninitialized extension \"XVideo\". xv.Register(xconn) must be called first.")
	}
	return c.Send(listImageFormatsRequest(op, Port))
}

// ListImageFormatsReply represents the data returned from a ListImageFormats request.
type ListImageFormatsReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	NumFormats uint32
	// padding: 20 bytes
	Format []ImageFormatInfo // size: ImageFormatInfoListSize(Format)
}

// Unmarshal reads a byte slice into a ListImageFormatsReply value.
func (v *ListImageFormatsReply) Unmarshal(buf []byte) error {
	if size := (32 + ImageFormatInfoListSize(v.Format)); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"ListImageFormatsReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.NumFormats = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 20 // padding

	v.Format = make([]ImageFormatInfo, v.NumFormats)
	b += ImageFormatInfoReadList(buf[b:], v.Format)

	return nil
}

// Write request to wire for ListImageFormats
// listImageFormatsRequest writes a ListImageFormats request to a byte slice.
func listImageFormatsRequest(opcode uint8, Port Port) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 16 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Port))
	b += 4

	return buf
}

// PutImage sends a checked request.
func PutImage(c *xgb.XConn, Port Port, Drawable xproto.Drawable, Gc xproto.Gcontext, Id uint32, SrcX int16, SrcY int16, SrcW uint16, SrcH uint16, DrwX int16, DrwY int16, DrwW uint16, DrwH uint16, Width uint16, Height uint16, Data []byte) error {
	op, ok := c.Ext("XVideo")
	if !ok {
		return errors.New("cannot issue request \"PutImage\" using the uninitialized extension \"XVideo\". xv.Register(xconn) must be called first.")
	}
	return c.SendRecv(putImageRequest(op, Port, Drawable, Gc, Id, SrcX, SrcY, SrcW, SrcH, DrwX, DrwY, DrwW, DrwH, Width, Height, Data), nil)
}

// PutImageUnchecked sends an unchecked request.
func PutImageUnchecked(c *xgb.XConn, Port Port, Drawable xproto.Drawable, Gc xproto.Gcontext, Id uint32, SrcX int16, SrcY int16, SrcW uint16, SrcH uint16, DrwX int16, DrwY int16, DrwW uint16, DrwH uint16, Width uint16, Height uint16, Data []byte) error {
	op, ok := c.Ext("XVideo")
	if !ok {
		return errors.New("cannot issue request \"PutImage\" using the uninitialized extension \"XVideo\". xv.Register(xconn) must be called first.")
	}
	return c.Send(putImageRequest(op, Port, Drawable, Gc, Id, SrcX, SrcY, SrcW, SrcH, DrwX, DrwY, DrwW, DrwH, Width, Height, Data))
}

// Write request to wire for PutImage
// putImageRequest writes a PutImage request to a byte slice.
func putImageRequest(opcode uint8, Port Port, Drawable xproto.Drawable, Gc xproto.Gcontext, Id uint32, SrcX int16, SrcY int16, SrcW uint16, SrcH uint16, DrwX int16, DrwY int16, DrwW uint16, DrwH uint16, Width uint16, Height uint16, Data []byte) []byte {
	size := internal.Pad4((40 + internal.Pad4((len(Data) * 1))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 18 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Port))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Drawable))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Gc))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], Id)
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(SrcX))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(SrcY))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], SrcW)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], SrcH)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(DrwX))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(DrwY))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], DrwW)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], DrwH)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Width)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Height)
	b += 2

	copy(buf[b:], Data[:len(Data)])
	b += int(len(Data))

	return buf
}

// PutStill sends a checked request.
func PutStill(c *xgb.XConn, Port Port, Drawable xproto.Drawable, Gc xproto.Gcontext, VidX int16, VidY int16, VidW uint16, VidH uint16, DrwX int16, DrwY int16, DrwW uint16, DrwH uint16) error {
	op, ok := c.Ext("XVideo")
	if !ok {
		return errors.New("cannot issue request \"PutStill\" using the uninitialized extension \"XVideo\". xv.Register(xconn) must be called first.")
	}
	return c.SendRecv(putStillRequest(op, Port, Drawable, Gc, VidX, VidY, VidW, VidH, DrwX, DrwY, DrwW, DrwH), nil)
}

// PutStillUnchecked sends an unchecked request.
func PutStillUnchecked(c *xgb.XConn, Port Port, Drawable xproto.Drawable, Gc xproto.Gcontext, VidX int16, VidY int16, VidW uint16, VidH uint16, DrwX int16, DrwY int16, DrwW uint16, DrwH uint16) error {
	op, ok := c.Ext("XVideo")
	if !ok {
		return errors.New("cannot issue request \"PutStill\" using the uninitialized extension \"XVideo\". xv.Register(xconn) must be called first.")
	}
	return c.Send(putStillRequest(op, Port, Drawable, Gc, VidX, VidY, VidW, VidH, DrwX, DrwY, DrwW, DrwH))
}

// Write request to wire for PutStill
// putStillRequest writes a PutStill request to a byte slice.
func putStillRequest(opcode uint8, Port Port, Drawable xproto.Drawable, Gc xproto.Gcontext, VidX int16, VidY int16, VidW uint16, VidH uint16, DrwX int16, DrwY int16, DrwW uint16, DrwH uint16) []byte {
	size := 32
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 6 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Port))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Drawable))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Gc))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(VidX))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(VidY))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], VidW)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], VidH)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(DrwX))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(DrwY))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], DrwW)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], DrwH)
	b += 2

	return buf
}

// PutVideo sends a checked request.
func PutVideo(c *xgb.XConn, Port Port, Drawable xproto.Drawable, Gc xproto.Gcontext, VidX int16, VidY int16, VidW uint16, VidH uint16, DrwX int16, DrwY int16, DrwW uint16, DrwH uint16) error {
	op, ok := c.Ext("XVideo")
	if !ok {
		return errors.New("cannot issue request \"PutVideo\" using the uninitialized extension \"XVideo\". xv.Register(xconn) must be called first.")
	}
	return c.SendRecv(putVideoRequest(op, Port, Drawable, Gc, VidX, VidY, VidW, VidH, DrwX, DrwY, DrwW, DrwH), nil)
}

// PutVideoUnchecked sends an unchecked request.
func PutVideoUnchecked(c *xgb.XConn, Port Port, Drawable xproto.Drawable, Gc xproto.Gcontext, VidX int16, VidY int16, VidW uint16, VidH uint16, DrwX int16, DrwY int16, DrwW uint16, DrwH uint16) error {
	op, ok := c.Ext("XVideo")
	if !ok {
		return errors.New("cannot issue request \"PutVideo\" using the uninitialized extension \"XVideo\". xv.Register(xconn) must be called first.")
	}
	return c.Send(putVideoRequest(op, Port, Drawable, Gc, VidX, VidY, VidW, VidH, DrwX, DrwY, DrwW, DrwH))
}

// Write request to wire for PutVideo
// putVideoRequest writes a PutVideo request to a byte slice.
func putVideoRequest(opcode uint8, Port Port, Drawable xproto.Drawable, Gc xproto.Gcontext, VidX int16, VidY int16, VidW uint16, VidH uint16, DrwX int16, DrwY int16, DrwW uint16, DrwH uint16) []byte {
	size := 32
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 5 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Port))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Drawable))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Gc))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(VidX))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(VidY))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], VidW)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], VidH)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(DrwX))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(DrwY))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], DrwW)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], DrwH)
	b += 2

	return buf
}

// QueryAdaptors sends a checked request.
func QueryAdaptors(c *xgb.XConn, Window xproto.Window) (QueryAdaptorsReply, error) {
	var reply QueryAdaptorsReply
	op, ok := c.Ext("XVideo")
	if !ok {
		return reply, errors.New("cannot issue request \"QueryAdaptors\" using the uninitialized extension \"XVideo\". xv.Register(xconn) must be called first.")
	}
	err := c.SendRecv(queryAdaptorsRequest(op, Window), &reply)
	return reply, err
}

// QueryAdaptorsUnchecked sends an unchecked request.
func QueryAdaptorsUnchecked(c *xgb.XConn, Window xproto.Window) error {
	op, ok := c.Ext("XVideo")
	if !ok {
		return errors.New("cannot issue request \"QueryAdaptors\" using the uninitialized extension \"XVideo\". xv.Register(xconn) must be called first.")
	}
	return c.Send(queryAdaptorsRequest(op, Window))
}

// QueryAdaptorsReply represents the data returned from a QueryAdaptors request.
type QueryAdaptorsReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	NumAdaptors uint16
	// padding: 22 bytes
	Info []AdaptorInfo // size: AdaptorInfoListSize(Info)
}

// Unmarshal reads a byte slice into a QueryAdaptorsReply value.
func (v *QueryAdaptorsReply) Unmarshal(buf []byte) error {
	if size := (32 + AdaptorInfoListSize(v.Info)); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"QueryAdaptorsReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.NumAdaptors = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	b += 22 // padding

	v.Info = make([]AdaptorInfo, v.NumAdaptors)
	b += AdaptorInfoReadList(buf[b:], v.Info)

	return nil
}

// Write request to wire for QueryAdaptors
// queryAdaptorsRequest writes a QueryAdaptors request to a byte slice.
func queryAdaptorsRequest(opcode uint8, Window xproto.Window) []byte {
	size := 8
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

	return buf
}

// QueryBestSize sends a checked request.
func QueryBestSize(c *xgb.XConn, Port Port, VidW uint16, VidH uint16, DrwW uint16, DrwH uint16, Motion bool) (QueryBestSizeReply, error) {
	var reply QueryBestSizeReply
	op, ok := c.Ext("XVideo")
	if !ok {
		return reply, errors.New("cannot issue request \"QueryBestSize\" using the uninitialized extension \"XVideo\". xv.Register(xconn) must be called first.")
	}
	err := c.SendRecv(queryBestSizeRequest(op, Port, VidW, VidH, DrwW, DrwH, Motion), &reply)
	return reply, err
}

// QueryBestSizeUnchecked sends an unchecked request.
func QueryBestSizeUnchecked(c *xgb.XConn, Port Port, VidW uint16, VidH uint16, DrwW uint16, DrwH uint16, Motion bool) error {
	op, ok := c.Ext("XVideo")
	if !ok {
		return errors.New("cannot issue request \"QueryBestSize\" using the uninitialized extension \"XVideo\". xv.Register(xconn) must be called first.")
	}
	return c.Send(queryBestSizeRequest(op, Port, VidW, VidH, DrwW, DrwH, Motion))
}

// QueryBestSizeReply represents the data returned from a QueryBestSize request.
type QueryBestSizeReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	ActualWidth  uint16
	ActualHeight uint16
}

// Unmarshal reads a byte slice into a QueryBestSizeReply value.
func (v *QueryBestSizeReply) Unmarshal(buf []byte) error {
	if size := 12; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"QueryBestSizeReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.ActualWidth = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.ActualHeight = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return nil
}

// Write request to wire for QueryBestSize
// queryBestSizeRequest writes a QueryBestSize request to a byte slice.
func queryBestSizeRequest(opcode uint8, Port Port, VidW uint16, VidH uint16, DrwW uint16, DrwH uint16, Motion bool) []byte {
	size := 20
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 12 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Port))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], VidW)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], VidH)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], DrwW)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], DrwH)
	b += 2

	if Motion {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	b += 3 // padding

	return buf
}

// QueryEncodings sends a checked request.
func QueryEncodings(c *xgb.XConn, Port Port) (QueryEncodingsReply, error) {
	var reply QueryEncodingsReply
	op, ok := c.Ext("XVideo")
	if !ok {
		return reply, errors.New("cannot issue request \"QueryEncodings\" using the uninitialized extension \"XVideo\". xv.Register(xconn) must be called first.")
	}
	err := c.SendRecv(queryEncodingsRequest(op, Port), &reply)
	return reply, err
}

// QueryEncodingsUnchecked sends an unchecked request.
func QueryEncodingsUnchecked(c *xgb.XConn, Port Port) error {
	op, ok := c.Ext("XVideo")
	if !ok {
		return errors.New("cannot issue request \"QueryEncodings\" using the uninitialized extension \"XVideo\". xv.Register(xconn) must be called first.")
	}
	return c.Send(queryEncodingsRequest(op, Port))
}

// QueryEncodingsReply represents the data returned from a QueryEncodings request.
type QueryEncodingsReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	NumEncodings uint16
	// padding: 22 bytes
	Info []EncodingInfo // size: EncodingInfoListSize(Info)
}

// Unmarshal reads a byte slice into a QueryEncodingsReply value.
func (v *QueryEncodingsReply) Unmarshal(buf []byte) error {
	if size := (32 + EncodingInfoListSize(v.Info)); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"QueryEncodingsReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.NumEncodings = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	b += 22 // padding

	v.Info = make([]EncodingInfo, v.NumEncodings)
	b += EncodingInfoReadList(buf[b:], v.Info)

	return nil
}

// Write request to wire for QueryEncodings
// queryEncodingsRequest writes a QueryEncodings request to a byte slice.
func queryEncodingsRequest(opcode uint8, Port Port) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 2 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Port))
	b += 4

	return buf
}

// QueryExtension sends a checked request.
func QueryExtension(c *xgb.XConn) (QueryExtensionReply, error) {
	var reply QueryExtensionReply
	op, ok := c.Ext("XVideo")
	if !ok {
		return reply, errors.New("cannot issue request \"QueryExtension\" using the uninitialized extension \"XVideo\". xv.Register(xconn) must be called first.")
	}
	err := c.SendRecv(queryExtensionRequest(op), &reply)
	return reply, err
}

// QueryExtensionUnchecked sends an unchecked request.
func QueryExtensionUnchecked(c *xgb.XConn) error {
	op, ok := c.Ext("XVideo")
	if !ok {
		return errors.New("cannot issue request \"QueryExtension\" using the uninitialized extension \"XVideo\". xv.Register(xconn) must be called first.")
	}
	return c.Send(queryExtensionRequest(op))
}

// QueryExtensionReply represents the data returned from a QueryExtension request.
type QueryExtensionReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Major uint16
	Minor uint16
}

// Unmarshal reads a byte slice into a QueryExtensionReply value.
func (v *QueryExtensionReply) Unmarshal(buf []byte) error {
	if size := 12; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"QueryExtensionReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Major = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Minor = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return nil
}

// Write request to wire for QueryExtension
// queryExtensionRequest writes a QueryExtension request to a byte slice.
func queryExtensionRequest(opcode uint8) []byte {
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

// QueryImageAttributes sends a checked request.
func QueryImageAttributes(c *xgb.XConn, Port Port, Id uint32, Width uint16, Height uint16) (QueryImageAttributesReply, error) {
	var reply QueryImageAttributesReply
	op, ok := c.Ext("XVideo")
	if !ok {
		return reply, errors.New("cannot issue request \"QueryImageAttributes\" using the uninitialized extension \"XVideo\". xv.Register(xconn) must be called first.")
	}
	err := c.SendRecv(queryImageAttributesRequest(op, Port, Id, Width, Height), &reply)
	return reply, err
}

// QueryImageAttributesUnchecked sends an unchecked request.
func QueryImageAttributesUnchecked(c *xgb.XConn, Port Port, Id uint32, Width uint16, Height uint16) error {
	op, ok := c.Ext("XVideo")
	if !ok {
		return errors.New("cannot issue request \"QueryImageAttributes\" using the uninitialized extension \"XVideo\". xv.Register(xconn) must be called first.")
	}
	return c.Send(queryImageAttributesRequest(op, Port, Id, Width, Height))
}

// QueryImageAttributesReply represents the data returned from a QueryImageAttributes request.
type QueryImageAttributesReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	NumPlanes uint32
	DataSize  uint32
	Width     uint16
	Height    uint16
	// padding: 12 bytes
	Pitches []uint32 // size: internal.Pad4((int(NumPlanes) * 4))
	// alignment gap to multiple of 4
	Offsets []uint32 // size: internal.Pad4((int(NumPlanes) * 4))
}

// Unmarshal reads a byte slice into a QueryImageAttributesReply value.
func (v *QueryImageAttributesReply) Unmarshal(buf []byte) error {
	if size := (((32 + internal.Pad4((int(v.NumPlanes) * 4))) + 4) + internal.Pad4((int(v.NumPlanes) * 4))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"QueryImageAttributesReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.NumPlanes = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.DataSize = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.Width = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Height = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	b += 12 // padding

	v.Pitches = make([]uint32, v.NumPlanes)
	for i := 0; i < int(v.NumPlanes); i++ {
		v.Pitches[i] = binary.LittleEndian.Uint32(buf[b:])
		b += 4
	}

	b = (b + 3) & ^3 // alignment gap

	v.Offsets = make([]uint32, v.NumPlanes)
	for i := 0; i < int(v.NumPlanes); i++ {
		v.Offsets[i] = binary.LittleEndian.Uint32(buf[b:])
		b += 4
	}

	return nil
}

// Write request to wire for QueryImageAttributes
// queryImageAttributesRequest writes a QueryImageAttributes request to a byte slice.
func queryImageAttributesRequest(opcode uint8, Port Port, Id uint32, Width uint16, Height uint16) []byte {
	size := 16
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 17 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Port))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], Id)
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], Width)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Height)
	b += 2

	return buf
}

// QueryPortAttributes sends a checked request.
func QueryPortAttributes(c *xgb.XConn, Port Port) (QueryPortAttributesReply, error) {
	var reply QueryPortAttributesReply
	op, ok := c.Ext("XVideo")
	if !ok {
		return reply, errors.New("cannot issue request \"QueryPortAttributes\" using the uninitialized extension \"XVideo\". xv.Register(xconn) must be called first.")
	}
	err := c.SendRecv(queryPortAttributesRequest(op, Port), &reply)
	return reply, err
}

// QueryPortAttributesUnchecked sends an unchecked request.
func QueryPortAttributesUnchecked(c *xgb.XConn, Port Port) error {
	op, ok := c.Ext("XVideo")
	if !ok {
		return errors.New("cannot issue request \"QueryPortAttributes\" using the uninitialized extension \"XVideo\". xv.Register(xconn) must be called first.")
	}
	return c.Send(queryPortAttributesRequest(op, Port))
}

// QueryPortAttributesReply represents the data returned from a QueryPortAttributes request.
type QueryPortAttributesReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	NumAttributes uint32
	TextSize      uint32
	// padding: 16 bytes
	Attributes []AttributeInfo // size: AttributeInfoListSize(Attributes)
}

// Unmarshal reads a byte slice into a QueryPortAttributesReply value.
func (v *QueryPortAttributesReply) Unmarshal(buf []byte) error {
	if size := (32 + AttributeInfoListSize(v.Attributes)); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"QueryPortAttributesReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.NumAttributes = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.TextSize = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 16 // padding

	v.Attributes = make([]AttributeInfo, v.NumAttributes)
	b += AttributeInfoReadList(buf[b:], v.Attributes)

	return nil
}

// Write request to wire for QueryPortAttributes
// queryPortAttributesRequest writes a QueryPortAttributes request to a byte slice.
func queryPortAttributesRequest(opcode uint8, Port Port) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 15 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Port))
	b += 4

	return buf
}

// SelectPortNotify sends a checked request.
func SelectPortNotify(c *xgb.XConn, Port Port, Onoff bool) error {
	op, ok := c.Ext("XVideo")
	if !ok {
		return errors.New("cannot issue request \"SelectPortNotify\" using the uninitialized extension \"XVideo\". xv.Register(xconn) must be called first.")
	}
	return c.SendRecv(selectPortNotifyRequest(op, Port, Onoff), nil)
}

// SelectPortNotifyUnchecked sends an unchecked request.
func SelectPortNotifyUnchecked(c *xgb.XConn, Port Port, Onoff bool) error {
	op, ok := c.Ext("XVideo")
	if !ok {
		return errors.New("cannot issue request \"SelectPortNotify\" using the uninitialized extension \"XVideo\". xv.Register(xconn) must be called first.")
	}
	return c.Send(selectPortNotifyRequest(op, Port, Onoff))
}

// Write request to wire for SelectPortNotify
// selectPortNotifyRequest writes a SelectPortNotify request to a byte slice.
func selectPortNotifyRequest(opcode uint8, Port Port, Onoff bool) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 11 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Port))
	b += 4

	if Onoff {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	b += 3 // padding

	return buf
}

// SelectVideoNotify sends a checked request.
func SelectVideoNotify(c *xgb.XConn, Drawable xproto.Drawable, Onoff bool) error {
	op, ok := c.Ext("XVideo")
	if !ok {
		return errors.New("cannot issue request \"SelectVideoNotify\" using the uninitialized extension \"XVideo\". xv.Register(xconn) must be called first.")
	}
	return c.SendRecv(selectVideoNotifyRequest(op, Drawable, Onoff), nil)
}

// SelectVideoNotifyUnchecked sends an unchecked request.
func SelectVideoNotifyUnchecked(c *xgb.XConn, Drawable xproto.Drawable, Onoff bool) error {
	op, ok := c.Ext("XVideo")
	if !ok {
		return errors.New("cannot issue request \"SelectVideoNotify\" using the uninitialized extension \"XVideo\". xv.Register(xconn) must be called first.")
	}
	return c.Send(selectVideoNotifyRequest(op, Drawable, Onoff))
}

// Write request to wire for SelectVideoNotify
// selectVideoNotifyRequest writes a SelectVideoNotify request to a byte slice.
func selectVideoNotifyRequest(opcode uint8, Drawable xproto.Drawable, Onoff bool) []byte {
	size := 12
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

	if Onoff {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	b += 3 // padding

	return buf
}

// SetPortAttribute sends a checked request.
func SetPortAttribute(c *xgb.XConn, Port Port, Attribute xproto.Atom, Value int32) error {
	op, ok := c.Ext("XVideo")
	if !ok {
		return errors.New("cannot issue request \"SetPortAttribute\" using the uninitialized extension \"XVideo\". xv.Register(xconn) must be called first.")
	}
	return c.SendRecv(setPortAttributeRequest(op, Port, Attribute, Value), nil)
}

// SetPortAttributeUnchecked sends an unchecked request.
func SetPortAttributeUnchecked(c *xgb.XConn, Port Port, Attribute xproto.Atom, Value int32) error {
	op, ok := c.Ext("XVideo")
	if !ok {
		return errors.New("cannot issue request \"SetPortAttribute\" using the uninitialized extension \"XVideo\". xv.Register(xconn) must be called first.")
	}
	return c.Send(setPortAttributeRequest(op, Port, Attribute, Value))
}

// Write request to wire for SetPortAttribute
// setPortAttributeRequest writes a SetPortAttribute request to a byte slice.
func setPortAttributeRequest(opcode uint8, Port Port, Attribute xproto.Atom, Value int32) []byte {
	size := 16
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 13 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Port))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Attribute))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Value))
	b += 4

	return buf
}

// ShmPutImage sends a checked request.
func ShmPutImage(c *xgb.XConn, Port Port, Drawable xproto.Drawable, Gc xproto.Gcontext, Shmseg shm.Seg, Id uint32, Offset uint32, SrcX int16, SrcY int16, SrcW uint16, SrcH uint16, DrwX int16, DrwY int16, DrwW uint16, DrwH uint16, Width uint16, Height uint16, SendEvent byte) error {
	op, ok := c.Ext("XVideo")
	if !ok {
		return errors.New("cannot issue request \"ShmPutImage\" using the uninitialized extension \"XVideo\". xv.Register(xconn) must be called first.")
	}
	return c.SendRecv(shmPutImageRequest(op, Port, Drawable, Gc, Shmseg, Id, Offset, SrcX, SrcY, SrcW, SrcH, DrwX, DrwY, DrwW, DrwH, Width, Height, SendEvent), nil)
}

// ShmPutImageUnchecked sends an unchecked request.
func ShmPutImageUnchecked(c *xgb.XConn, Port Port, Drawable xproto.Drawable, Gc xproto.Gcontext, Shmseg shm.Seg, Id uint32, Offset uint32, SrcX int16, SrcY int16, SrcW uint16, SrcH uint16, DrwX int16, DrwY int16, DrwW uint16, DrwH uint16, Width uint16, Height uint16, SendEvent byte) error {
	op, ok := c.Ext("XVideo")
	if !ok {
		return errors.New("cannot issue request \"ShmPutImage\" using the uninitialized extension \"XVideo\". xv.Register(xconn) must be called first.")
	}
	return c.Send(shmPutImageRequest(op, Port, Drawable, Gc, Shmseg, Id, Offset, SrcX, SrcY, SrcW, SrcH, DrwX, DrwY, DrwW, DrwH, Width, Height, SendEvent))
}

// Write request to wire for ShmPutImage
// shmPutImageRequest writes a ShmPutImage request to a byte slice.
func shmPutImageRequest(opcode uint8, Port Port, Drawable xproto.Drawable, Gc xproto.Gcontext, Shmseg shm.Seg, Id uint32, Offset uint32, SrcX int16, SrcY int16, SrcW uint16, SrcH uint16, DrwX int16, DrwY int16, DrwW uint16, DrwH uint16, Width uint16, Height uint16, SendEvent byte) []byte {
	size := 52
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 19 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Port))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Drawable))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Gc))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Shmseg))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], Id)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], Offset)
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(SrcX))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(SrcY))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], SrcW)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], SrcH)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(DrwX))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(DrwY))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], DrwW)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], DrwH)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Width)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Height)
	b += 2

	buf[b] = SendEvent
	b += 1

	b += 3 // padding

	return buf
}

// StopVideo sends a checked request.
func StopVideo(c *xgb.XConn, Port Port, Drawable xproto.Drawable) error {
	op, ok := c.Ext("XVideo")
	if !ok {
		return errors.New("cannot issue request \"StopVideo\" using the uninitialized extension \"XVideo\". xv.Register(xconn) must be called first.")
	}
	return c.SendRecv(stopVideoRequest(op, Port, Drawable), nil)
}

// StopVideoUnchecked sends an unchecked request.
func StopVideoUnchecked(c *xgb.XConn, Port Port, Drawable xproto.Drawable) error {
	op, ok := c.Ext("XVideo")
	if !ok {
		return errors.New("cannot issue request \"StopVideo\" using the uninitialized extension \"XVideo\". xv.Register(xconn) must be called first.")
	}
	return c.Send(stopVideoRequest(op, Port, Drawable))
}

// Write request to wire for StopVideo
// stopVideoRequest writes a StopVideo request to a byte slice.
func stopVideoRequest(opcode uint8, Port Port, Drawable xproto.Drawable) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 9 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Port))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Drawable))
	b += 4

	return buf
}

// UngrabPort sends a checked request.
func UngrabPort(c *xgb.XConn, Port Port, Time xproto.Timestamp) error {
	op, ok := c.Ext("XVideo")
	if !ok {
		return errors.New("cannot issue request \"UngrabPort\" using the uninitialized extension \"XVideo\". xv.Register(xconn) must be called first.")
	}
	return c.SendRecv(ungrabPortRequest(op, Port, Time), nil)
}

// UngrabPortUnchecked sends an unchecked request.
func UngrabPortUnchecked(c *xgb.XConn, Port Port, Time xproto.Timestamp) error {
	op, ok := c.Ext("XVideo")
	if !ok {
		return errors.New("cannot issue request \"UngrabPort\" using the uninitialized extension \"XVideo\". xv.Register(xconn) must be called first.")
	}
	return c.Send(ungrabPortRequest(op, Port, Time))
}

// Write request to wire for UngrabPort
// ungrabPortRequest writes a UngrabPort request to a byte slice.
func ungrabPortRequest(opcode uint8, Port Port, Time xproto.Timestamp) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 4 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Port))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Time))
	b += 4

	return buf
}
