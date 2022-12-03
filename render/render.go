// FILE GENERATED AUTOMATICALLY FROM "render.xml"
package render

import (
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"codeberg.org/gruf/go-xgb"
	"codeberg.org/gruf/go-xgb/internal"
	"codeberg.org/gruf/go-xgb/xproto"
)

const (
	// ExtName is the user-friendly name string of this X extension.
	ExtName = "Render"

	// ExtXName is the name string this extension is known by to the X server.
	ExtXName = "RENDER"
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
		return fmt.Errorf("error querying X for \"Render\": %w", err)
	} else if !reply.Present {
		return fmt.Errorf("no extension named \"Render\" is known to the X server: reply=%+v", reply)
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

type Animcursorelt struct {
	Cursor xproto.Cursor
	Delay  uint32
}

// AnimcursoreltRead reads a byte slice into a Animcursorelt value.
func AnimcursoreltRead(buf []byte, v *Animcursorelt) int {
	b := 0

	v.Cursor = xproto.Cursor(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Delay = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	return b
}

// AnimcursoreltReadList reads a byte slice into a list of Animcursorelt values.
func AnimcursoreltReadList(buf []byte, dest []Animcursorelt) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Animcursorelt{}
		b += AnimcursoreltRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a Animcursorelt value to a byte slice.
func (v Animcursorelt) Bytes() []byte {
	buf := make([]byte, 8)
	b := 0

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Cursor))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], v.Delay)
	b += 4

	return buf[:b]
}

// AnimcursoreltListBytes writes a list of Animcursorelt values to a byte slice.
func AnimcursoreltListBytes(buf []byte, list []Animcursorelt) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

type Color struct {
	Red   uint16
	Green uint16
	Blue  uint16
	Alpha uint16
}

// ColorRead reads a byte slice into a Color value.
func ColorRead(buf []byte, v *Color) int {
	b := 0

	v.Red = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Green = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Blue = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Alpha = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return b
}

// ColorReadList reads a byte slice into a list of Color values.
func ColorReadList(buf []byte, dest []Color) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Color{}
		b += ColorRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a Color value to a byte slice.
func (v Color) Bytes() []byte {
	buf := make([]byte, 8)
	b := 0

	binary.LittleEndian.PutUint16(buf[b:], v.Red)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Green)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Blue)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Alpha)
	b += 2

	return buf[:b]
}

// ColorListBytes writes a list of Color values to a byte slice.
func ColorListBytes(buf []byte, list []Color) int {
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
	CpRepeat           = 1
	CpAlphaMap         = 2
	CpAlphaXOrigin     = 4
	CpAlphaYOrigin     = 8
	CpClipXOrigin      = 16
	CpClipYOrigin      = 32
	CpClipMask         = 64
	CpGraphicsExposure = 128
	CpSubwindowMode    = 256
	CpPolyEdge         = 512
	CpPolyMode         = 1024
	CpDither           = 2048
	CpComponentAlpha   = 4096
)

type Directformat struct {
	RedShift   uint16
	RedMask    uint16
	GreenShift uint16
	GreenMask  uint16
	BlueShift  uint16
	BlueMask   uint16
	AlphaShift uint16
	AlphaMask  uint16
}

// DirectformatRead reads a byte slice into a Directformat value.
func DirectformatRead(buf []byte, v *Directformat) int {
	b := 0

	v.RedShift = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.RedMask = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.GreenShift = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.GreenMask = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.BlueShift = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.BlueMask = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.AlphaShift = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.AlphaMask = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return b
}

// DirectformatReadList reads a byte slice into a list of Directformat values.
func DirectformatReadList(buf []byte, dest []Directformat) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Directformat{}
		b += DirectformatRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a Directformat value to a byte slice.
func (v Directformat) Bytes() []byte {
	buf := make([]byte, 16)
	b := 0

	binary.LittleEndian.PutUint16(buf[b:], v.RedShift)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.RedMask)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.GreenShift)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.GreenMask)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.BlueShift)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.BlueMask)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.AlphaShift)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.AlphaMask)
	b += 2

	return buf[:b]
}

// DirectformatListBytes writes a list of Directformat values to a byte slice.
func DirectformatListBytes(buf []byte, list []Directformat) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

type Fixed int32

type Glyph uint32

// BadGlyph is the error number for a BadGlyph.
const BadGlyph = 4

type GlyphError struct {
	Sequence uint16
	NiceName string
}

// UnmarshalGlyphError constructs a GlyphError value that implements xgb.Error from a byte slice.
func UnmarshalGlyphError(buf []byte) (xgb.XError, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"GlyphError\"", len(buf))
	}

	v := GlyphError{}
	v.NiceName = "Glyph"

	b := 1 // skip error determinant
	b += 1 // don't read error number

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return v, nil
}

// SeqID returns the sequence id attached to the BadGlyph error.
// This is mostly used internally.
func (err GlyphError) SeqID() uint16 {
	return err.Sequence
}

// BadID returns the 'BadValue' number if one exists for the BadGlyph error. If no bad value exists, 0 is returned.
func (err GlyphError) BadID() uint32 {
	return 0
}

// Error returns a rudimentary string representation of the BadGlyph error.
func (err GlyphError) Error() string {
	var buf strings.Builder

	buf.WriteString("BadGlyph{")
	buf.WriteString("NiceName: " + err.NiceName)
	buf.WriteByte(' ')
	buf.WriteString("Sequence: " + strconv.FormatUint(uint64(err.Sequence), 10))

	buf.WriteByte('}')

	return buf.String()
}

func init() { registerError(4, UnmarshalGlyphError) }

// BadGlyphSet is the error number for a BadGlyphSet.
const BadGlyphSet = 3

type GlyphSetError struct {
	Sequence uint16
	NiceName string
}

// UnmarshalGlyphSetError constructs a GlyphSetError value that implements xgb.Error from a byte slice.
func UnmarshalGlyphSetError(buf []byte) (xgb.XError, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"GlyphSetError\"", len(buf))
	}

	v := GlyphSetError{}
	v.NiceName = "GlyphSet"

	b := 1 // skip error determinant
	b += 1 // don't read error number

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return v, nil
}

// SeqID returns the sequence id attached to the BadGlyphSet error.
// This is mostly used internally.
func (err GlyphSetError) SeqID() uint16 {
	return err.Sequence
}

// BadID returns the 'BadValue' number if one exists for the BadGlyphSet error. If no bad value exists, 0 is returned.
func (err GlyphSetError) BadID() uint32 {
	return 0
}

// Error returns a rudimentary string representation of the BadGlyphSet error.
func (err GlyphSetError) Error() string {
	var buf strings.Builder

	buf.WriteString("BadGlyphSet{")
	buf.WriteString("NiceName: " + err.NiceName)
	buf.WriteByte(' ')
	buf.WriteString("Sequence: " + strconv.FormatUint(uint64(err.Sequence), 10))

	buf.WriteByte('}')

	return buf.String()
}

func init() { registerError(3, UnmarshalGlyphSetError) }

type Glyphinfo struct {
	Width  uint16
	Height uint16
	X      int16
	Y      int16
	XOff   int16
	YOff   int16
}

// GlyphinfoRead reads a byte slice into a Glyphinfo value.
func GlyphinfoRead(buf []byte, v *Glyphinfo) int {
	b := 0

	v.Width = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Height = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.X = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.Y = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.XOff = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.YOff = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	return b
}

// GlyphinfoReadList reads a byte slice into a list of Glyphinfo values.
func GlyphinfoReadList(buf []byte, dest []Glyphinfo) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Glyphinfo{}
		b += GlyphinfoRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a Glyphinfo value to a byte slice.
func (v Glyphinfo) Bytes() []byte {
	buf := make([]byte, 12)
	b := 0

	binary.LittleEndian.PutUint16(buf[b:], v.Width)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Height)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.X))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.Y))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.XOff))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.YOff))
	b += 2

	return buf[:b]
}

// GlyphinfoListBytes writes a list of Glyphinfo values to a byte slice.
func GlyphinfoListBytes(buf []byte, list []Glyphinfo) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

type Glyphset uint32

func NewGlyphsetID(c *xgb.XConn) (Glyphset, error) {
	id, err := c.NewXID()
	return Glyphset(id), err
}

type Indexvalue struct {
	Pixel uint32
	Red   uint16
	Green uint16
	Blue  uint16
	Alpha uint16
}

// IndexvalueRead reads a byte slice into a Indexvalue value.
func IndexvalueRead(buf []byte, v *Indexvalue) int {
	b := 0

	v.Pixel = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.Red = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Green = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Blue = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Alpha = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return b
}

// IndexvalueReadList reads a byte slice into a list of Indexvalue values.
func IndexvalueReadList(buf []byte, dest []Indexvalue) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Indexvalue{}
		b += IndexvalueRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a Indexvalue value to a byte slice.
func (v Indexvalue) Bytes() []byte {
	buf := make([]byte, 12)
	b := 0

	binary.LittleEndian.PutUint32(buf[b:], v.Pixel)
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], v.Red)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Green)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Blue)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Alpha)
	b += 2

	return buf[:b]
}

// IndexvalueListBytes writes a list of Indexvalue values to a byte slice.
func IndexvalueListBytes(buf []byte, list []Indexvalue) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

type Linefix struct {
	P1 Pointfix
	P2 Pointfix
}

// LinefixRead reads a byte slice into a Linefix value.
func LinefixRead(buf []byte, v *Linefix) int {
	b := 0

	v.P1 = Pointfix{}
	b += PointfixRead(buf[b:], &v.P1)

	v.P2 = Pointfix{}
	b += PointfixRead(buf[b:], &v.P2)

	return b
}

// LinefixReadList reads a byte slice into a list of Linefix values.
func LinefixReadList(buf []byte, dest []Linefix) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Linefix{}
		b += LinefixRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a Linefix value to a byte slice.
func (v Linefix) Bytes() []byte {
	buf := make([]byte, 16)
	b := 0

	{
		structBytes := v.P1.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	{
		structBytes := v.P2.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	return buf[:b]
}

// LinefixListBytes writes a list of Linefix values to a byte slice.
func LinefixListBytes(buf []byte, list []Linefix) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

// BadPictFormat is the error number for a BadPictFormat.
const BadPictFormat = 0

type PictFormatError struct {
	Sequence uint16
	NiceName string
}

// UnmarshalPictFormatError constructs a PictFormatError value that implements xgb.Error from a byte slice.
func UnmarshalPictFormatError(buf []byte) (xgb.XError, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"PictFormatError\"", len(buf))
	}

	v := PictFormatError{}
	v.NiceName = "PictFormat"

	b := 1 // skip error determinant
	b += 1 // don't read error number

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return v, nil
}

// SeqID returns the sequence id attached to the BadPictFormat error.
// This is mostly used internally.
func (err PictFormatError) SeqID() uint16 {
	return err.Sequence
}

// BadID returns the 'BadValue' number if one exists for the BadPictFormat error. If no bad value exists, 0 is returned.
func (err PictFormatError) BadID() uint32 {
	return 0
}

// Error returns a rudimentary string representation of the BadPictFormat error.
func (err PictFormatError) Error() string {
	var buf strings.Builder

	buf.WriteString("BadPictFormat{")
	buf.WriteString("NiceName: " + err.NiceName)
	buf.WriteByte(' ')
	buf.WriteString("Sequence: " + strconv.FormatUint(uint64(err.Sequence), 10))

	buf.WriteByte('}')

	return buf.String()
}

func init() { registerError(0, UnmarshalPictFormatError) }

const (
	PictOpClear               = 0
	PictOpSrc                 = 1
	PictOpDst                 = 2
	PictOpOver                = 3
	PictOpOverReverse         = 4
	PictOpIn                  = 5
	PictOpInReverse           = 6
	PictOpOut                 = 7
	PictOpOutReverse          = 8
	PictOpAtop                = 9
	PictOpAtopReverse         = 10
	PictOpXor                 = 11
	PictOpAdd                 = 12
	PictOpSaturate            = 13
	PictOpDisjointClear       = 16
	PictOpDisjointSrc         = 17
	PictOpDisjointDst         = 18
	PictOpDisjointOver        = 19
	PictOpDisjointOverReverse = 20
	PictOpDisjointIn          = 21
	PictOpDisjointInReverse   = 22
	PictOpDisjointOut         = 23
	PictOpDisjointOutReverse  = 24
	PictOpDisjointAtop        = 25
	PictOpDisjointAtopReverse = 26
	PictOpDisjointXor         = 27
	PictOpConjointClear       = 32
	PictOpConjointSrc         = 33
	PictOpConjointDst         = 34
	PictOpConjointOver        = 35
	PictOpConjointOverReverse = 36
	PictOpConjointIn          = 37
	PictOpConjointInReverse   = 38
	PictOpConjointOut         = 39
	PictOpConjointOutReverse  = 40
	PictOpConjointAtop        = 41
	PictOpConjointAtopReverse = 42
	PictOpConjointXor         = 43
	PictOpMultiply            = 48
	PictOpScreen              = 49
	PictOpOverlay             = 50
	PictOpDarken              = 51
	PictOpLighten             = 52
	PictOpColorDodge          = 53
	PictOpColorBurn           = 54
	PictOpHardLight           = 55
	PictOpSoftLight           = 56
	PictOpDifference          = 57
	PictOpExclusion           = 58
	PictOpHSLHue              = 59
	PictOpHSLSaturation       = 60
	PictOpHSLColor            = 61
	PictOpHSLLuminosity       = 62
)

// BadPictOp is the error number for a BadPictOp.
const BadPictOp = 2

type PictOpError struct {
	Sequence uint16
	NiceName string
}

// UnmarshalPictOpError constructs a PictOpError value that implements xgb.Error from a byte slice.
func UnmarshalPictOpError(buf []byte) (xgb.XError, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"PictOpError\"", len(buf))
	}

	v := PictOpError{}
	v.NiceName = "PictOp"

	b := 1 // skip error determinant
	b += 1 // don't read error number

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return v, nil
}

// SeqID returns the sequence id attached to the BadPictOp error.
// This is mostly used internally.
func (err PictOpError) SeqID() uint16 {
	return err.Sequence
}

// BadID returns the 'BadValue' number if one exists for the BadPictOp error. If no bad value exists, 0 is returned.
func (err PictOpError) BadID() uint32 {
	return 0
}

// Error returns a rudimentary string representation of the BadPictOp error.
func (err PictOpError) Error() string {
	var buf strings.Builder

	buf.WriteString("BadPictOp{")
	buf.WriteString("NiceName: " + err.NiceName)
	buf.WriteByte(' ')
	buf.WriteString("Sequence: " + strconv.FormatUint(uint64(err.Sequence), 10))

	buf.WriteByte('}')

	return buf.String()
}

func init() { registerError(2, UnmarshalPictOpError) }

const (
	PictTypeIndexed = 0
	PictTypeDirect  = 1
)

type Pictdepth struct {
	Depth byte
	// padding: 1 bytes
	NumVisuals uint16
	// padding: 4 bytes
	Visuals []Pictvisual // size: internal.Pad4((int(NumVisuals) * 8))
}

// PictdepthRead reads a byte slice into a Pictdepth value.
func PictdepthRead(buf []byte, v *Pictdepth) int {
	b := 0

	v.Depth = buf[b]
	b += 1

	b += 1 // padding

	v.NumVisuals = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	b += 4 // padding

	v.Visuals = make([]Pictvisual, v.NumVisuals)
	b += PictvisualReadList(buf[b:], v.Visuals)

	return b
}

// PictdepthReadList reads a byte slice into a list of Pictdepth values.
func PictdepthReadList(buf []byte, dest []Pictdepth) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Pictdepth{}
		b += PictdepthRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a Pictdepth value to a byte slice.
func (v Pictdepth) Bytes() []byte {
	buf := make([]byte, (8 + internal.Pad4((int(v.NumVisuals) * 8))))
	b := 0

	buf[b] = v.Depth
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], v.NumVisuals)
	b += 2

	b += 4 // padding

	b += PictvisualListBytes(buf[b:], v.Visuals)

	return buf[:b]
}

// PictdepthListBytes writes a list of Pictdepth values to a byte slice.
func PictdepthListBytes(buf []byte, list []Pictdepth) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

// PictdepthListSize computes the size (bytes) of a list of Pictdepth values.
func PictdepthListSize(list []Pictdepth) int {
	size := 0
	for _, item := range list {
		size += (8 + internal.Pad4((int(item.NumVisuals) * 8)))
	}
	return size
}

type Pictformat uint32

func NewPictformatID(c *xgb.XConn) (Pictformat, error) {
	id, err := c.NewXID()
	return Pictformat(id), err
}

type Pictforminfo struct {
	Id    Pictformat
	Type  byte
	Depth byte
	// padding: 2 bytes
	Direct   Directformat
	Colormap xproto.Colormap
}

// PictforminfoRead reads a byte slice into a Pictforminfo value.
func PictforminfoRead(buf []byte, v *Pictforminfo) int {
	b := 0

	v.Id = Pictformat(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Type = buf[b]
	b += 1

	v.Depth = buf[b]
	b += 1

	b += 2 // padding

	v.Direct = Directformat{}
	b += DirectformatRead(buf[b:], &v.Direct)

	v.Colormap = xproto.Colormap(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	return b
}

// PictforminfoReadList reads a byte slice into a list of Pictforminfo values.
func PictforminfoReadList(buf []byte, dest []Pictforminfo) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Pictforminfo{}
		b += PictforminfoRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a Pictforminfo value to a byte slice.
func (v Pictforminfo) Bytes() []byte {
	buf := make([]byte, 28)
	b := 0

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Id))
	b += 4

	buf[b] = v.Type
	b += 1

	buf[b] = v.Depth
	b += 1

	b += 2 // padding

	{
		structBytes := v.Direct.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Colormap))
	b += 4

	return buf[:b]
}

// PictforminfoListBytes writes a list of Pictforminfo values to a byte slice.
func PictforminfoListBytes(buf []byte, list []Pictforminfo) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

type Pictscreen struct {
	NumDepths uint32
	Fallback  Pictformat
	Depths    []Pictdepth // size: PictdepthListSize(Depths)
}

// PictscreenRead reads a byte slice into a Pictscreen value.
func PictscreenRead(buf []byte, v *Pictscreen) int {
	b := 0

	v.NumDepths = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.Fallback = Pictformat(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Depths = make([]Pictdepth, v.NumDepths)
	b += PictdepthReadList(buf[b:], v.Depths)

	return b
}

// PictscreenReadList reads a byte slice into a list of Pictscreen values.
func PictscreenReadList(buf []byte, dest []Pictscreen) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Pictscreen{}
		b += PictscreenRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a Pictscreen value to a byte slice.
func (v Pictscreen) Bytes() []byte {
	buf := make([]byte, (8 + PictdepthListSize(v.Depths)))
	b := 0

	binary.LittleEndian.PutUint32(buf[b:], v.NumDepths)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Fallback))
	b += 4

	b += PictdepthListBytes(buf[b:], v.Depths)

	return buf[:b]
}

// PictscreenListBytes writes a list of Pictscreen values to a byte slice.
func PictscreenListBytes(buf []byte, list []Pictscreen) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

// PictscreenListSize computes the size (bytes) of a list of Pictscreen values.
func PictscreenListSize(list []Pictscreen) int {
	size := 0
	for _, item := range list {
		size += (8 + PictdepthListSize(item.Depths))
	}
	return size
}

type Picture uint32

func NewPictureID(c *xgb.XConn) (Picture, error) {
	id, err := c.NewXID()
	return Picture(id), err
}

// BadPicture is the error number for a BadPicture.
const BadPicture = 1

type PictureError struct {
	Sequence uint16
	NiceName string
}

// UnmarshalPictureError constructs a PictureError value that implements xgb.Error from a byte slice.
func UnmarshalPictureError(buf []byte) (xgb.XError, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"PictureError\"", len(buf))
	}

	v := PictureError{}
	v.NiceName = "Picture"

	b := 1 // skip error determinant
	b += 1 // don't read error number

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return v, nil
}

// SeqID returns the sequence id attached to the BadPicture error.
// This is mostly used internally.
func (err PictureError) SeqID() uint16 {
	return err.Sequence
}

// BadID returns the 'BadValue' number if one exists for the BadPicture error. If no bad value exists, 0 is returned.
func (err PictureError) BadID() uint32 {
	return 0
}

// Error returns a rudimentary string representation of the BadPicture error.
func (err PictureError) Error() string {
	var buf strings.Builder

	buf.WriteString("BadPicture{")
	buf.WriteString("NiceName: " + err.NiceName)
	buf.WriteByte(' ')
	buf.WriteString("Sequence: " + strconv.FormatUint(uint64(err.Sequence), 10))

	buf.WriteByte('}')

	return buf.String()
}

func init() { registerError(1, UnmarshalPictureError) }

const (
	PictureNone = 0
)

type Pictvisual struct {
	Visual xproto.Visualid
	Format Pictformat
}

// PictvisualRead reads a byte slice into a Pictvisual value.
func PictvisualRead(buf []byte, v *Pictvisual) int {
	b := 0

	v.Visual = xproto.Visualid(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Format = Pictformat(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	return b
}

// PictvisualReadList reads a byte slice into a list of Pictvisual values.
func PictvisualReadList(buf []byte, dest []Pictvisual) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Pictvisual{}
		b += PictvisualRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a Pictvisual value to a byte slice.
func (v Pictvisual) Bytes() []byte {
	buf := make([]byte, 8)
	b := 0

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Visual))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Format))
	b += 4

	return buf[:b]
}

// PictvisualListBytes writes a list of Pictvisual values to a byte slice.
func PictvisualListBytes(buf []byte, list []Pictvisual) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

type Pointfix struct {
	X Fixed
	Y Fixed
}

// PointfixRead reads a byte slice into a Pointfix value.
func PointfixRead(buf []byte, v *Pointfix) int {
	b := 0

	v.X = Fixed(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Y = Fixed(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	return b
}

// PointfixReadList reads a byte slice into a list of Pointfix values.
func PointfixReadList(buf []byte, dest []Pointfix) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Pointfix{}
		b += PointfixRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a Pointfix value to a byte slice.
func (v Pointfix) Bytes() []byte {
	buf := make([]byte, 8)
	b := 0

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.X))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Y))
	b += 4

	return buf[:b]
}

// PointfixListBytes writes a list of Pointfix values to a byte slice.
func PointfixListBytes(buf []byte, list []Pointfix) int {
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
	PolyEdgeSharp  = 0
	PolyEdgeSmooth = 1
)

const (
	PolyModePrecise   = 0
	PolyModeImprecise = 1
)

const (
	RepeatNone    = 0
	RepeatNormal  = 1
	RepeatPad     = 2
	RepeatReflect = 3
)

type Spanfix struct {
	L Fixed
	R Fixed
	Y Fixed
}

// SpanfixRead reads a byte slice into a Spanfix value.
func SpanfixRead(buf []byte, v *Spanfix) int {
	b := 0

	v.L = Fixed(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.R = Fixed(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Y = Fixed(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	return b
}

// SpanfixReadList reads a byte slice into a list of Spanfix values.
func SpanfixReadList(buf []byte, dest []Spanfix) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Spanfix{}
		b += SpanfixRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a Spanfix value to a byte slice.
func (v Spanfix) Bytes() []byte {
	buf := make([]byte, 12)
	b := 0

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.L))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.R))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Y))
	b += 4

	return buf[:b]
}

// SpanfixListBytes writes a list of Spanfix values to a byte slice.
func SpanfixListBytes(buf []byte, list []Spanfix) int {
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
	SubPixelUnknown       = 0
	SubPixelHorizontalRGB = 1
	SubPixelHorizontalBGR = 2
	SubPixelVerticalRGB   = 3
	SubPixelVerticalBGR   = 4
	SubPixelNone          = 5
)

type Transform struct {
	Matrix11 Fixed
	Matrix12 Fixed
	Matrix13 Fixed
	Matrix21 Fixed
	Matrix22 Fixed
	Matrix23 Fixed
	Matrix31 Fixed
	Matrix32 Fixed
	Matrix33 Fixed
}

// TransformRead reads a byte slice into a Transform value.
func TransformRead(buf []byte, v *Transform) int {
	b := 0

	v.Matrix11 = Fixed(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Matrix12 = Fixed(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Matrix13 = Fixed(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Matrix21 = Fixed(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Matrix22 = Fixed(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Matrix23 = Fixed(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Matrix31 = Fixed(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Matrix32 = Fixed(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Matrix33 = Fixed(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	return b
}

// TransformReadList reads a byte slice into a list of Transform values.
func TransformReadList(buf []byte, dest []Transform) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Transform{}
		b += TransformRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a Transform value to a byte slice.
func (v Transform) Bytes() []byte {
	buf := make([]byte, 36)
	b := 0

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Matrix11))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Matrix12))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Matrix13))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Matrix21))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Matrix22))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Matrix23))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Matrix31))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Matrix32))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Matrix33))
	b += 4

	return buf[:b]
}

// TransformListBytes writes a list of Transform values to a byte slice.
func TransformListBytes(buf []byte, list []Transform) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

type Trap struct {
	Top Spanfix
	Bot Spanfix
}

// TrapRead reads a byte slice into a Trap value.
func TrapRead(buf []byte, v *Trap) int {
	b := 0

	v.Top = Spanfix{}
	b += SpanfixRead(buf[b:], &v.Top)

	v.Bot = Spanfix{}
	b += SpanfixRead(buf[b:], &v.Bot)

	return b
}

// TrapReadList reads a byte slice into a list of Trap values.
func TrapReadList(buf []byte, dest []Trap) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Trap{}
		b += TrapRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a Trap value to a byte slice.
func (v Trap) Bytes() []byte {
	buf := make([]byte, 24)
	b := 0

	{
		structBytes := v.Top.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	{
		structBytes := v.Bot.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	return buf[:b]
}

// TrapListBytes writes a list of Trap values to a byte slice.
func TrapListBytes(buf []byte, list []Trap) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

type Trapezoid struct {
	Top    Fixed
	Bottom Fixed
	Left   Linefix
	Right  Linefix
}

// TrapezoidRead reads a byte slice into a Trapezoid value.
func TrapezoidRead(buf []byte, v *Trapezoid) int {
	b := 0

	v.Top = Fixed(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Bottom = Fixed(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Left = Linefix{}
	b += LinefixRead(buf[b:], &v.Left)

	v.Right = Linefix{}
	b += LinefixRead(buf[b:], &v.Right)

	return b
}

// TrapezoidReadList reads a byte slice into a list of Trapezoid values.
func TrapezoidReadList(buf []byte, dest []Trapezoid) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Trapezoid{}
		b += TrapezoidRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a Trapezoid value to a byte slice.
func (v Trapezoid) Bytes() []byte {
	buf := make([]byte, 40)
	b := 0

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Top))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Bottom))
	b += 4

	{
		structBytes := v.Left.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	{
		structBytes := v.Right.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	return buf[:b]
}

// TrapezoidListBytes writes a list of Trapezoid values to a byte slice.
func TrapezoidListBytes(buf []byte, list []Trapezoid) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

type Triangle struct {
	P1 Pointfix
	P2 Pointfix
	P3 Pointfix
}

// TriangleRead reads a byte slice into a Triangle value.
func TriangleRead(buf []byte, v *Triangle) int {
	b := 0

	v.P1 = Pointfix{}
	b += PointfixRead(buf[b:], &v.P1)

	v.P2 = Pointfix{}
	b += PointfixRead(buf[b:], &v.P2)

	v.P3 = Pointfix{}
	b += PointfixRead(buf[b:], &v.P3)

	return b
}

// TriangleReadList reads a byte slice into a list of Triangle values.
func TriangleReadList(buf []byte, dest []Triangle) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Triangle{}
		b += TriangleRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a Triangle value to a byte slice.
func (v Triangle) Bytes() []byte {
	buf := make([]byte, 24)
	b := 0

	{
		structBytes := v.P1.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	{
		structBytes := v.P2.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	{
		structBytes := v.P3.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	return buf[:b]
}

// TriangleListBytes writes a list of Triangle values to a byte slice.
func TriangleListBytes(buf []byte, list []Triangle) int {
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

// AddGlyphs sends a checked request.
func AddGlyphs(c *xgb.XConn, Glyphset Glyphset, GlyphsLen uint32, Glyphids []uint32, Glyphs []Glyphinfo, Data []byte) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"AddGlyphs\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.SendRecv(addGlyphsRequest(op, Glyphset, GlyphsLen, Glyphids, Glyphs, Data), nil)
}

// AddGlyphsUnchecked sends an unchecked request.
func AddGlyphsUnchecked(c *xgb.XConn, Glyphset Glyphset, GlyphsLen uint32, Glyphids []uint32, Glyphs []Glyphinfo, Data []byte) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"AddGlyphs\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.Send(addGlyphsRequest(op, Glyphset, GlyphsLen, Glyphids, Glyphs, Data))
}

// Write request to wire for AddGlyphs
// addGlyphsRequest writes a AddGlyphs request to a byte slice.
func addGlyphsRequest(opcode uint8, Glyphset Glyphset, GlyphsLen uint32, Glyphids []uint32, Glyphs []Glyphinfo, Data []byte) []byte {
	size := internal.Pad4(((((12 + internal.Pad4((int(GlyphsLen) * 4))) + 4) + internal.Pad4((int(GlyphsLen) * 12))) + internal.Pad4((len(Data) * 1))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 20 // request opcode
	b += 1

	blen := b
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Glyphset))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], GlyphsLen)
	b += 4

	for i := 0; i < int(GlyphsLen); i++ {
		binary.LittleEndian.PutUint32(buf[b:], Glyphids[i])
		b += 4
	}

	b = (b + 3) & ^3 // alignment gap

	b += GlyphinfoListBytes(buf[b:], Glyphs)

	copy(buf[b:], Data[:len(Data)])
	b += int(len(Data))

	b = internal.Pad4(b)
	binary.LittleEndian.PutUint16(buf[blen:], uint16(b/4)) // write request size in 4-byte units
	return buf[:b]
}

// AddTraps sends a checked request.
func AddTraps(c *xgb.XConn, Picture Picture, XOff int16, YOff int16, Traps []Trap) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"AddTraps\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.SendRecv(addTrapsRequest(op, Picture, XOff, YOff, Traps), nil)
}

// AddTrapsUnchecked sends an unchecked request.
func AddTrapsUnchecked(c *xgb.XConn, Picture Picture, XOff int16, YOff int16, Traps []Trap) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"AddTraps\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.Send(addTrapsRequest(op, Picture, XOff, YOff, Traps))
}

// Write request to wire for AddTraps
// addTrapsRequest writes a AddTraps request to a byte slice.
func addTrapsRequest(opcode uint8, Picture Picture, XOff int16, YOff int16, Traps []Trap) []byte {
	size := internal.Pad4((12 + internal.Pad4((len(Traps) * 24))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 32 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Picture))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(XOff))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(YOff))
	b += 2

	b += TrapListBytes(buf[b:], Traps)

	return buf
}

// ChangePicture sends a checked request.
func ChangePicture(c *xgb.XConn, Picture Picture, ValueMask uint32, ValueList []uint32) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"ChangePicture\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.SendRecv(changePictureRequest(op, Picture, ValueMask, ValueList), nil)
}

// ChangePictureUnchecked sends an unchecked request.
func ChangePictureUnchecked(c *xgb.XConn, Picture Picture, ValueMask uint32, ValueList []uint32) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"ChangePicture\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.Send(changePictureRequest(op, Picture, ValueMask, ValueList))
}

// Write request to wire for ChangePicture
// changePictureRequest writes a ChangePicture request to a byte slice.
func changePictureRequest(opcode uint8, Picture Picture, ValueMask uint32, ValueList []uint32) []byte {
	size := internal.Pad4((12 + internal.Pad4((4 * internal.PopCount(int(ValueMask))))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 5 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Picture))
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

// Composite sends a checked request.
func Composite(c *xgb.XConn, Op byte, Src Picture, Mask Picture, Dst Picture, SrcX int16, SrcY int16, MaskX int16, MaskY int16, DstX int16, DstY int16, Width uint16, Height uint16) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"Composite\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.SendRecv(compositeRequest(op, Op, Src, Mask, Dst, SrcX, SrcY, MaskX, MaskY, DstX, DstY, Width, Height), nil)
}

// CompositeUnchecked sends an unchecked request.
func CompositeUnchecked(c *xgb.XConn, Op byte, Src Picture, Mask Picture, Dst Picture, SrcX int16, SrcY int16, MaskX int16, MaskY int16, DstX int16, DstY int16, Width uint16, Height uint16) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"Composite\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.Send(compositeRequest(op, Op, Src, Mask, Dst, SrcX, SrcY, MaskX, MaskY, DstX, DstY, Width, Height))
}

// Write request to wire for Composite
// compositeRequest writes a Composite request to a byte slice.
func compositeRequest(opcode uint8, Op byte, Src Picture, Mask Picture, Dst Picture, SrcX int16, SrcY int16, MaskX int16, MaskY int16, DstX int16, DstY int16, Width uint16, Height uint16) []byte {
	size := 36
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 8 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	buf[b] = Op
	b += 1

	b += 3 // padding

	binary.LittleEndian.PutUint32(buf[b:], uint32(Src))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Mask))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Dst))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(SrcX))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(SrcY))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(MaskX))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(MaskY))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(DstX))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(DstY))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Width)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Height)
	b += 2

	return buf
}

// CompositeGlyphs16 sends a checked request.
func CompositeGlyphs16(c *xgb.XConn, Op byte, Src Picture, Dst Picture, MaskFormat Pictformat, Glyphset Glyphset, SrcX int16, SrcY int16, Glyphcmds []byte) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"CompositeGlyphs16\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.SendRecv(compositeGlyphs16Request(op, Op, Src, Dst, MaskFormat, Glyphset, SrcX, SrcY, Glyphcmds), nil)
}

// CompositeGlyphs16Unchecked sends an unchecked request.
func CompositeGlyphs16Unchecked(c *xgb.XConn, Op byte, Src Picture, Dst Picture, MaskFormat Pictformat, Glyphset Glyphset, SrcX int16, SrcY int16, Glyphcmds []byte) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"CompositeGlyphs16\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.Send(compositeGlyphs16Request(op, Op, Src, Dst, MaskFormat, Glyphset, SrcX, SrcY, Glyphcmds))
}

// Write request to wire for CompositeGlyphs16
// compositeGlyphs16Request writes a CompositeGlyphs16 request to a byte slice.
func compositeGlyphs16Request(opcode uint8, Op byte, Src Picture, Dst Picture, MaskFormat Pictformat, Glyphset Glyphset, SrcX int16, SrcY int16, Glyphcmds []byte) []byte {
	size := internal.Pad4((28 + internal.Pad4((len(Glyphcmds) * 1))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 24 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	buf[b] = Op
	b += 1

	b += 3 // padding

	binary.LittleEndian.PutUint32(buf[b:], uint32(Src))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Dst))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(MaskFormat))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Glyphset))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(SrcX))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(SrcY))
	b += 2

	copy(buf[b:], Glyphcmds[:len(Glyphcmds)])
	b += int(len(Glyphcmds))

	return buf
}

// CompositeGlyphs32 sends a checked request.
func CompositeGlyphs32(c *xgb.XConn, Op byte, Src Picture, Dst Picture, MaskFormat Pictformat, Glyphset Glyphset, SrcX int16, SrcY int16, Glyphcmds []byte) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"CompositeGlyphs32\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.SendRecv(compositeGlyphs32Request(op, Op, Src, Dst, MaskFormat, Glyphset, SrcX, SrcY, Glyphcmds), nil)
}

// CompositeGlyphs32Unchecked sends an unchecked request.
func CompositeGlyphs32Unchecked(c *xgb.XConn, Op byte, Src Picture, Dst Picture, MaskFormat Pictformat, Glyphset Glyphset, SrcX int16, SrcY int16, Glyphcmds []byte) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"CompositeGlyphs32\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.Send(compositeGlyphs32Request(op, Op, Src, Dst, MaskFormat, Glyphset, SrcX, SrcY, Glyphcmds))
}

// Write request to wire for CompositeGlyphs32
// compositeGlyphs32Request writes a CompositeGlyphs32 request to a byte slice.
func compositeGlyphs32Request(opcode uint8, Op byte, Src Picture, Dst Picture, MaskFormat Pictformat, Glyphset Glyphset, SrcX int16, SrcY int16, Glyphcmds []byte) []byte {
	size := internal.Pad4((28 + internal.Pad4((len(Glyphcmds) * 1))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 25 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	buf[b] = Op
	b += 1

	b += 3 // padding

	binary.LittleEndian.PutUint32(buf[b:], uint32(Src))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Dst))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(MaskFormat))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Glyphset))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(SrcX))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(SrcY))
	b += 2

	copy(buf[b:], Glyphcmds[:len(Glyphcmds)])
	b += int(len(Glyphcmds))

	return buf
}

// CompositeGlyphs8 sends a checked request.
func CompositeGlyphs8(c *xgb.XConn, Op byte, Src Picture, Dst Picture, MaskFormat Pictformat, Glyphset Glyphset, SrcX int16, SrcY int16, Glyphcmds []byte) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"CompositeGlyphs8\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.SendRecv(compositeGlyphs8Request(op, Op, Src, Dst, MaskFormat, Glyphset, SrcX, SrcY, Glyphcmds), nil)
}

// CompositeGlyphs8Unchecked sends an unchecked request.
func CompositeGlyphs8Unchecked(c *xgb.XConn, Op byte, Src Picture, Dst Picture, MaskFormat Pictformat, Glyphset Glyphset, SrcX int16, SrcY int16, Glyphcmds []byte) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"CompositeGlyphs8\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.Send(compositeGlyphs8Request(op, Op, Src, Dst, MaskFormat, Glyphset, SrcX, SrcY, Glyphcmds))
}

// Write request to wire for CompositeGlyphs8
// compositeGlyphs8Request writes a CompositeGlyphs8 request to a byte slice.
func compositeGlyphs8Request(opcode uint8, Op byte, Src Picture, Dst Picture, MaskFormat Pictformat, Glyphset Glyphset, SrcX int16, SrcY int16, Glyphcmds []byte) []byte {
	size := internal.Pad4((28 + internal.Pad4((len(Glyphcmds) * 1))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 23 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	buf[b] = Op
	b += 1

	b += 3 // padding

	binary.LittleEndian.PutUint32(buf[b:], uint32(Src))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Dst))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(MaskFormat))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Glyphset))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(SrcX))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(SrcY))
	b += 2

	copy(buf[b:], Glyphcmds[:len(Glyphcmds)])
	b += int(len(Glyphcmds))

	return buf
}

// CreateAnimCursor sends a checked request.
func CreateAnimCursor(c *xgb.XConn, Cid xproto.Cursor, Cursors []Animcursorelt) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"CreateAnimCursor\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.SendRecv(createAnimCursorRequest(op, Cid, Cursors), nil)
}

// CreateAnimCursorUnchecked sends an unchecked request.
func CreateAnimCursorUnchecked(c *xgb.XConn, Cid xproto.Cursor, Cursors []Animcursorelt) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"CreateAnimCursor\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.Send(createAnimCursorRequest(op, Cid, Cursors))
}

// Write request to wire for CreateAnimCursor
// createAnimCursorRequest writes a CreateAnimCursor request to a byte slice.
func createAnimCursorRequest(opcode uint8, Cid xproto.Cursor, Cursors []Animcursorelt) []byte {
	size := internal.Pad4((8 + internal.Pad4((len(Cursors) * 8))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 31 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Cid))
	b += 4

	b += AnimcursoreltListBytes(buf[b:], Cursors)

	return buf
}

// CreateConicalGradient sends a checked request.
func CreateConicalGradient(c *xgb.XConn, Picture Picture, Center Pointfix, Angle Fixed, NumStops uint32, Stops []Fixed, Colors []Color) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"CreateConicalGradient\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.SendRecv(createConicalGradientRequest(op, Picture, Center, Angle, NumStops, Stops, Colors), nil)
}

// CreateConicalGradientUnchecked sends an unchecked request.
func CreateConicalGradientUnchecked(c *xgb.XConn, Picture Picture, Center Pointfix, Angle Fixed, NumStops uint32, Stops []Fixed, Colors []Color) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"CreateConicalGradient\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.Send(createConicalGradientRequest(op, Picture, Center, Angle, NumStops, Stops, Colors))
}

// Write request to wire for CreateConicalGradient
// createConicalGradientRequest writes a CreateConicalGradient request to a byte slice.
func createConicalGradientRequest(opcode uint8, Picture Picture, Center Pointfix, Angle Fixed, NumStops uint32, Stops []Fixed, Colors []Color) []byte {
	size := internal.Pad4((((24 + internal.Pad4((int(NumStops) * 4))) + 4) + internal.Pad4((int(NumStops) * 8))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 36 // request opcode
	b += 1

	blen := b
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Picture))
	b += 4

	{
		structBytes := Center.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	binary.LittleEndian.PutUint32(buf[b:], uint32(Angle))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], NumStops)
	b += 4

	for i := 0; i < int(NumStops); i++ {
		binary.LittleEndian.PutUint32(buf[b:], uint32(Stops[i]))
		b += 4
	}

	b = (b + 3) & ^3 // alignment gap

	b += ColorListBytes(buf[b:], Colors)

	b = internal.Pad4(b)
	binary.LittleEndian.PutUint16(buf[blen:], uint16(b/4)) // write request size in 4-byte units
	return buf[:b]
}

// CreateCursor sends a checked request.
func CreateCursor(c *xgb.XConn, Cid xproto.Cursor, Source Picture, X uint16, Y uint16) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"CreateCursor\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.SendRecv(createCursorRequest(op, Cid, Source, X, Y), nil)
}

// CreateCursorUnchecked sends an unchecked request.
func CreateCursorUnchecked(c *xgb.XConn, Cid xproto.Cursor, Source Picture, X uint16, Y uint16) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"CreateCursor\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.Send(createCursorRequest(op, Cid, Source, X, Y))
}

// Write request to wire for CreateCursor
// createCursorRequest writes a CreateCursor request to a byte slice.
func createCursorRequest(opcode uint8, Cid xproto.Cursor, Source Picture, X uint16, Y uint16) []byte {
	size := 16
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 27 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Cid))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Source))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], X)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Y)
	b += 2

	return buf
}

// CreateGlyphSet sends a checked request.
func CreateGlyphSet(c *xgb.XConn, Gsid Glyphset, Format Pictformat) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"CreateGlyphSet\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.SendRecv(createGlyphSetRequest(op, Gsid, Format), nil)
}

// CreateGlyphSetUnchecked sends an unchecked request.
func CreateGlyphSetUnchecked(c *xgb.XConn, Gsid Glyphset, Format Pictformat) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"CreateGlyphSet\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.Send(createGlyphSetRequest(op, Gsid, Format))
}

// Write request to wire for CreateGlyphSet
// createGlyphSetRequest writes a CreateGlyphSet request to a byte slice.
func createGlyphSetRequest(opcode uint8, Gsid Glyphset, Format Pictformat) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 17 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Gsid))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Format))
	b += 4

	return buf
}

// CreateLinearGradient sends a checked request.
func CreateLinearGradient(c *xgb.XConn, Picture Picture, P1 Pointfix, P2 Pointfix, NumStops uint32, Stops []Fixed, Colors []Color) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"CreateLinearGradient\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.SendRecv(createLinearGradientRequest(op, Picture, P1, P2, NumStops, Stops, Colors), nil)
}

// CreateLinearGradientUnchecked sends an unchecked request.
func CreateLinearGradientUnchecked(c *xgb.XConn, Picture Picture, P1 Pointfix, P2 Pointfix, NumStops uint32, Stops []Fixed, Colors []Color) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"CreateLinearGradient\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.Send(createLinearGradientRequest(op, Picture, P1, P2, NumStops, Stops, Colors))
}

// Write request to wire for CreateLinearGradient
// createLinearGradientRequest writes a CreateLinearGradient request to a byte slice.
func createLinearGradientRequest(opcode uint8, Picture Picture, P1 Pointfix, P2 Pointfix, NumStops uint32, Stops []Fixed, Colors []Color) []byte {
	size := internal.Pad4((((28 + internal.Pad4((int(NumStops) * 4))) + 4) + internal.Pad4((int(NumStops) * 8))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 34 // request opcode
	b += 1

	blen := b
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Picture))
	b += 4

	{
		structBytes := P1.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	{
		structBytes := P2.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	binary.LittleEndian.PutUint32(buf[b:], NumStops)
	b += 4

	for i := 0; i < int(NumStops); i++ {
		binary.LittleEndian.PutUint32(buf[b:], uint32(Stops[i]))
		b += 4
	}

	b = (b + 3) & ^3 // alignment gap

	b += ColorListBytes(buf[b:], Colors)

	b = internal.Pad4(b)
	binary.LittleEndian.PutUint16(buf[blen:], uint16(b/4)) // write request size in 4-byte units
	return buf[:b]
}

// CreatePicture sends a checked request.
func CreatePicture(c *xgb.XConn, Pid Picture, Drawable xproto.Drawable, Format Pictformat, ValueMask uint32, ValueList []uint32) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"CreatePicture\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.SendRecv(createPictureRequest(op, Pid, Drawable, Format, ValueMask, ValueList), nil)
}

// CreatePictureUnchecked sends an unchecked request.
func CreatePictureUnchecked(c *xgb.XConn, Pid Picture, Drawable xproto.Drawable, Format Pictformat, ValueMask uint32, ValueList []uint32) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"CreatePicture\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.Send(createPictureRequest(op, Pid, Drawable, Format, ValueMask, ValueList))
}

// Write request to wire for CreatePicture
// createPictureRequest writes a CreatePicture request to a byte slice.
func createPictureRequest(opcode uint8, Pid Picture, Drawable xproto.Drawable, Format Pictformat, ValueMask uint32, ValueList []uint32) []byte {
	size := internal.Pad4((20 + internal.Pad4((4 * internal.PopCount(int(ValueMask))))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 4 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Pid))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Drawable))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Format))
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

// CreateRadialGradient sends a checked request.
func CreateRadialGradient(c *xgb.XConn, Picture Picture, Inner Pointfix, Outer Pointfix, InnerRadius Fixed, OuterRadius Fixed, NumStops uint32, Stops []Fixed, Colors []Color) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"CreateRadialGradient\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.SendRecv(createRadialGradientRequest(op, Picture, Inner, Outer, InnerRadius, OuterRadius, NumStops, Stops, Colors), nil)
}

// CreateRadialGradientUnchecked sends an unchecked request.
func CreateRadialGradientUnchecked(c *xgb.XConn, Picture Picture, Inner Pointfix, Outer Pointfix, InnerRadius Fixed, OuterRadius Fixed, NumStops uint32, Stops []Fixed, Colors []Color) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"CreateRadialGradient\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.Send(createRadialGradientRequest(op, Picture, Inner, Outer, InnerRadius, OuterRadius, NumStops, Stops, Colors))
}

// Write request to wire for CreateRadialGradient
// createRadialGradientRequest writes a CreateRadialGradient request to a byte slice.
func createRadialGradientRequest(opcode uint8, Picture Picture, Inner Pointfix, Outer Pointfix, InnerRadius Fixed, OuterRadius Fixed, NumStops uint32, Stops []Fixed, Colors []Color) []byte {
	size := internal.Pad4((((36 + internal.Pad4((int(NumStops) * 4))) + 4) + internal.Pad4((int(NumStops) * 8))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 35 // request opcode
	b += 1

	blen := b
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Picture))
	b += 4

	{
		structBytes := Inner.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	{
		structBytes := Outer.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	binary.LittleEndian.PutUint32(buf[b:], uint32(InnerRadius))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(OuterRadius))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], NumStops)
	b += 4

	for i := 0; i < int(NumStops); i++ {
		binary.LittleEndian.PutUint32(buf[b:], uint32(Stops[i]))
		b += 4
	}

	b = (b + 3) & ^3 // alignment gap

	b += ColorListBytes(buf[b:], Colors)

	b = internal.Pad4(b)
	binary.LittleEndian.PutUint16(buf[blen:], uint16(b/4)) // write request size in 4-byte units
	return buf[:b]
}

// CreateSolidFill sends a checked request.
func CreateSolidFill(c *xgb.XConn, Picture Picture, Color Color) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"CreateSolidFill\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.SendRecv(createSolidFillRequest(op, Picture, Color), nil)
}

// CreateSolidFillUnchecked sends an unchecked request.
func CreateSolidFillUnchecked(c *xgb.XConn, Picture Picture, Color Color) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"CreateSolidFill\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.Send(createSolidFillRequest(op, Picture, Color))
}

// Write request to wire for CreateSolidFill
// createSolidFillRequest writes a CreateSolidFill request to a byte slice.
func createSolidFillRequest(opcode uint8, Picture Picture, Color Color) []byte {
	size := 16
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 33 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Picture))
	b += 4

	{
		structBytes := Color.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	return buf
}

// FillRectangles sends a checked request.
func FillRectangles(c *xgb.XConn, Op byte, Dst Picture, Color Color, Rects []xproto.Rectangle) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"FillRectangles\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.SendRecv(fillRectanglesRequest(op, Op, Dst, Color, Rects), nil)
}

// FillRectanglesUnchecked sends an unchecked request.
func FillRectanglesUnchecked(c *xgb.XConn, Op byte, Dst Picture, Color Color, Rects []xproto.Rectangle) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"FillRectangles\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.Send(fillRectanglesRequest(op, Op, Dst, Color, Rects))
}

// Write request to wire for FillRectangles
// fillRectanglesRequest writes a FillRectangles request to a byte slice.
func fillRectanglesRequest(opcode uint8, Op byte, Dst Picture, Color Color, Rects []xproto.Rectangle) []byte {
	size := internal.Pad4((20 + internal.Pad4((len(Rects) * 8))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 26 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	buf[b] = Op
	b += 1

	b += 3 // padding

	binary.LittleEndian.PutUint32(buf[b:], uint32(Dst))
	b += 4

	{
		structBytes := Color.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	b += xproto.RectangleListBytes(buf[b:], Rects)

	return buf
}

// FreeGlyphSet sends a checked request.
func FreeGlyphSet(c *xgb.XConn, Glyphset Glyphset) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"FreeGlyphSet\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.SendRecv(freeGlyphSetRequest(op, Glyphset), nil)
}

// FreeGlyphSetUnchecked sends an unchecked request.
func FreeGlyphSetUnchecked(c *xgb.XConn, Glyphset Glyphset) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"FreeGlyphSet\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.Send(freeGlyphSetRequest(op, Glyphset))
}

// Write request to wire for FreeGlyphSet
// freeGlyphSetRequest writes a FreeGlyphSet request to a byte slice.
func freeGlyphSetRequest(opcode uint8, Glyphset Glyphset) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 19 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Glyphset))
	b += 4

	return buf
}

// FreeGlyphs sends a checked request.
func FreeGlyphs(c *xgb.XConn, Glyphset Glyphset, Glyphs []Glyph) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"FreeGlyphs\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.SendRecv(freeGlyphsRequest(op, Glyphset, Glyphs), nil)
}

// FreeGlyphsUnchecked sends an unchecked request.
func FreeGlyphsUnchecked(c *xgb.XConn, Glyphset Glyphset, Glyphs []Glyph) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"FreeGlyphs\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.Send(freeGlyphsRequest(op, Glyphset, Glyphs))
}

// Write request to wire for FreeGlyphs
// freeGlyphsRequest writes a FreeGlyphs request to a byte slice.
func freeGlyphsRequest(opcode uint8, Glyphset Glyphset, Glyphs []Glyph) []byte {
	size := internal.Pad4((8 + internal.Pad4((len(Glyphs) * 4))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 22 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Glyphset))
	b += 4

	for i := 0; i < int(len(Glyphs)); i++ {
		binary.LittleEndian.PutUint32(buf[b:], uint32(Glyphs[i]))
		b += 4
	}

	return buf
}

// FreePicture sends a checked request.
func FreePicture(c *xgb.XConn, Picture Picture) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"FreePicture\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.SendRecv(freePictureRequest(op, Picture), nil)
}

// FreePictureUnchecked sends an unchecked request.
func FreePictureUnchecked(c *xgb.XConn, Picture Picture) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"FreePicture\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.Send(freePictureRequest(op, Picture))
}

// Write request to wire for FreePicture
// freePictureRequest writes a FreePicture request to a byte slice.
func freePictureRequest(opcode uint8, Picture Picture) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 7 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Picture))
	b += 4

	return buf
}

// QueryFilters sends a checked request.
func QueryFilters(c *xgb.XConn, Drawable xproto.Drawable) (QueryFiltersReply, error) {
	var reply QueryFiltersReply
	op, ok := c.Ext("RENDER")
	if !ok {
		return reply, errors.New("cannot issue request \"QueryFilters\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	err := c.SendRecv(queryFiltersRequest(op, Drawable), &reply)
	return reply, err
}

// QueryFiltersUnchecked sends an unchecked request.
func QueryFiltersUnchecked(c *xgb.XConn, Drawable xproto.Drawable) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"QueryFilters\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.Send(queryFiltersRequest(op, Drawable))
}

// QueryFiltersReply represents the data returned from a QueryFilters request.
type QueryFiltersReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	NumAliases uint32
	NumFilters uint32
	// padding: 16 bytes
	Aliases []uint16     // size: internal.Pad4((int(NumAliases) * 2))
	Filters []xproto.Str // size: xproto.StrListSize(Filters)
}

// Unmarshal reads a byte slice into a QueryFiltersReply value.
func (v *QueryFiltersReply) Unmarshal(buf []byte) error {
	if size := ((32 + internal.Pad4((int(v.NumAliases) * 2))) + xproto.StrListSize(v.Filters)); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"QueryFiltersReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.NumAliases = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.NumFilters = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 16 // padding

	v.Aliases = make([]uint16, v.NumAliases)
	for i := 0; i < int(v.NumAliases); i++ {
		v.Aliases[i] = binary.LittleEndian.Uint16(buf[b:])
		b += 2
	}

	v.Filters = make([]xproto.Str, v.NumFilters)
	b += xproto.StrReadList(buf[b:], v.Filters)

	return nil
}

// Write request to wire for QueryFilters
// queryFiltersRequest writes a QueryFilters request to a byte slice.
func queryFiltersRequest(opcode uint8, Drawable xproto.Drawable) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 29 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Drawable))
	b += 4

	return buf
}

// QueryPictFormats sends a checked request.
func QueryPictFormats(c *xgb.XConn) (QueryPictFormatsReply, error) {
	var reply QueryPictFormatsReply
	op, ok := c.Ext("RENDER")
	if !ok {
		return reply, errors.New("cannot issue request \"QueryPictFormats\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	err := c.SendRecv(queryPictFormatsRequest(op), &reply)
	return reply, err
}

// QueryPictFormatsUnchecked sends an unchecked request.
func QueryPictFormatsUnchecked(c *xgb.XConn) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"QueryPictFormats\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.Send(queryPictFormatsRequest(op))
}

// QueryPictFormatsReply represents the data returned from a QueryPictFormats request.
type QueryPictFormatsReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	NumFormats  uint32
	NumScreens  uint32
	NumDepths   uint32
	NumVisuals  uint32
	NumSubpixel uint32
	// padding: 4 bytes
	Formats []Pictforminfo // size: internal.Pad4((int(NumFormats) * 28))
	// alignment gap to multiple of 4
	Screens []Pictscreen // size: PictscreenListSize(Screens)
	// alignment gap to multiple of 4
	Subpixels []uint32 // size: internal.Pad4((int(NumSubpixel) * 4))
}

// Unmarshal reads a byte slice into a QueryPictFormatsReply value.
func (v *QueryPictFormatsReply) Unmarshal(buf []byte) error {
	if size := (((((32 + internal.Pad4((int(v.NumFormats) * 28))) + 4) + PictscreenListSize(v.Screens)) + 4) + internal.Pad4((int(v.NumSubpixel) * 4))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"QueryPictFormatsReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.NumFormats = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.NumScreens = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.NumDepths = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.NumVisuals = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.NumSubpixel = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 4 // padding

	v.Formats = make([]Pictforminfo, v.NumFormats)
	b += PictforminfoReadList(buf[b:], v.Formats)

	b = (b + 3) & ^3 // alignment gap

	v.Screens = make([]Pictscreen, v.NumScreens)
	b += PictscreenReadList(buf[b:], v.Screens)

	b = (b + 3) & ^3 // alignment gap

	v.Subpixels = make([]uint32, v.NumSubpixel)
	for i := 0; i < int(v.NumSubpixel); i++ {
		v.Subpixels[i] = binary.LittleEndian.Uint32(buf[b:])
		b += 4
	}

	return nil
}

// Write request to wire for QueryPictFormats
// queryPictFormatsRequest writes a QueryPictFormats request to a byte slice.
func queryPictFormatsRequest(opcode uint8) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 1 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// QueryPictIndexValues sends a checked request.
func QueryPictIndexValues(c *xgb.XConn, Format Pictformat) (QueryPictIndexValuesReply, error) {
	var reply QueryPictIndexValuesReply
	op, ok := c.Ext("RENDER")
	if !ok {
		return reply, errors.New("cannot issue request \"QueryPictIndexValues\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	err := c.SendRecv(queryPictIndexValuesRequest(op, Format), &reply)
	return reply, err
}

// QueryPictIndexValuesUnchecked sends an unchecked request.
func QueryPictIndexValuesUnchecked(c *xgb.XConn, Format Pictformat) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"QueryPictIndexValues\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.Send(queryPictIndexValuesRequest(op, Format))
}

// QueryPictIndexValuesReply represents the data returned from a QueryPictIndexValues request.
type QueryPictIndexValuesReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	NumValues uint32
	// padding: 20 bytes
	Values []Indexvalue // size: internal.Pad4((int(NumValues) * 12))
}

// Unmarshal reads a byte slice into a QueryPictIndexValuesReply value.
func (v *QueryPictIndexValuesReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.NumValues) * 12))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"QueryPictIndexValuesReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.NumValues = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 20 // padding

	v.Values = make([]Indexvalue, v.NumValues)
	b += IndexvalueReadList(buf[b:], v.Values)

	return nil
}

// Write request to wire for QueryPictIndexValues
// queryPictIndexValuesRequest writes a QueryPictIndexValues request to a byte slice.
func queryPictIndexValuesRequest(opcode uint8, Format Pictformat) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 2 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Format))
	b += 4

	return buf
}

// QueryVersion sends a checked request.
func QueryVersion(c *xgb.XConn, ClientMajorVersion uint32, ClientMinorVersion uint32) (QueryVersionReply, error) {
	var reply QueryVersionReply
	op, ok := c.Ext("RENDER")
	if !ok {
		return reply, errors.New("cannot issue request \"QueryVersion\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	err := c.SendRecv(queryVersionRequest(op, ClientMajorVersion, ClientMinorVersion), &reply)
	return reply, err
}

// QueryVersionUnchecked sends an unchecked request.
func QueryVersionUnchecked(c *xgb.XConn, ClientMajorVersion uint32, ClientMinorVersion uint32) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"QueryVersion\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.Send(queryVersionRequest(op, ClientMajorVersion, ClientMinorVersion))
}

// QueryVersionReply represents the data returned from a QueryVersion request.
type QueryVersionReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	MajorVersion uint32
	MinorVersion uint32
	// padding: 16 bytes
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

	v.MajorVersion = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.MinorVersion = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 16 // padding

	return nil
}

// Write request to wire for QueryVersion
// queryVersionRequest writes a QueryVersion request to a byte slice.
func queryVersionRequest(opcode uint8, ClientMajorVersion uint32, ClientMinorVersion uint32) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 0 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], ClientMajorVersion)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], ClientMinorVersion)
	b += 4

	return buf
}

// ReferenceGlyphSet sends a checked request.
func ReferenceGlyphSet(c *xgb.XConn, Gsid Glyphset, Existing Glyphset) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"ReferenceGlyphSet\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.SendRecv(referenceGlyphSetRequest(op, Gsid, Existing), nil)
}

// ReferenceGlyphSetUnchecked sends an unchecked request.
func ReferenceGlyphSetUnchecked(c *xgb.XConn, Gsid Glyphset, Existing Glyphset) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"ReferenceGlyphSet\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.Send(referenceGlyphSetRequest(op, Gsid, Existing))
}

// Write request to wire for ReferenceGlyphSet
// referenceGlyphSetRequest writes a ReferenceGlyphSet request to a byte slice.
func referenceGlyphSetRequest(opcode uint8, Gsid Glyphset, Existing Glyphset) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 18 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Gsid))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Existing))
	b += 4

	return buf
}

// SetPictureClipRectangles sends a checked request.
func SetPictureClipRectangles(c *xgb.XConn, Picture Picture, ClipXOrigin int16, ClipYOrigin int16, Rectangles []xproto.Rectangle) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"SetPictureClipRectangles\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.SendRecv(setPictureClipRectanglesRequest(op, Picture, ClipXOrigin, ClipYOrigin, Rectangles), nil)
}

// SetPictureClipRectanglesUnchecked sends an unchecked request.
func SetPictureClipRectanglesUnchecked(c *xgb.XConn, Picture Picture, ClipXOrigin int16, ClipYOrigin int16, Rectangles []xproto.Rectangle) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"SetPictureClipRectangles\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.Send(setPictureClipRectanglesRequest(op, Picture, ClipXOrigin, ClipYOrigin, Rectangles))
}

// Write request to wire for SetPictureClipRectangles
// setPictureClipRectanglesRequest writes a SetPictureClipRectangles request to a byte slice.
func setPictureClipRectanglesRequest(opcode uint8, Picture Picture, ClipXOrigin int16, ClipYOrigin int16, Rectangles []xproto.Rectangle) []byte {
	size := internal.Pad4((12 + internal.Pad4((len(Rectangles) * 8))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 6 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Picture))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(ClipXOrigin))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(ClipYOrigin))
	b += 2

	b += xproto.RectangleListBytes(buf[b:], Rectangles)

	return buf
}

// SetPictureFilter sends a checked request.
func SetPictureFilter(c *xgb.XConn, Picture Picture, FilterLen uint16, Filter string, Values []Fixed) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"SetPictureFilter\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.SendRecv(setPictureFilterRequest(op, Picture, FilterLen, Filter, Values), nil)
}

// SetPictureFilterUnchecked sends an unchecked request.
func SetPictureFilterUnchecked(c *xgb.XConn, Picture Picture, FilterLen uint16, Filter string, Values []Fixed) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"SetPictureFilter\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.Send(setPictureFilterRequest(op, Picture, FilterLen, Filter, Values))
}

// Write request to wire for SetPictureFilter
// setPictureFilterRequest writes a SetPictureFilter request to a byte slice.
func setPictureFilterRequest(opcode uint8, Picture Picture, FilterLen uint16, Filter string, Values []Fixed) []byte {
	size := internal.Pad4((((12 + internal.Pad4((int(FilterLen) * 1))) + 0) + internal.Pad4((len(Values) * 4))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 30 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Picture))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], FilterLen)
	b += 2

	b += 2 // padding

	copy(buf[b:], Filter[:FilterLen])
	b += int(FilterLen)

	b += 0 // padding

	for i := 0; i < int(len(Values)); i++ {
		binary.LittleEndian.PutUint32(buf[b:], uint32(Values[i]))
		b += 4
	}

	return buf
}

// SetPictureTransform sends a checked request.
func SetPictureTransform(c *xgb.XConn, Picture Picture, Transform Transform) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"SetPictureTransform\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.SendRecv(setPictureTransformRequest(op, Picture, Transform), nil)
}

// SetPictureTransformUnchecked sends an unchecked request.
func SetPictureTransformUnchecked(c *xgb.XConn, Picture Picture, Transform Transform) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"SetPictureTransform\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.Send(setPictureTransformRequest(op, Picture, Transform))
}

// Write request to wire for SetPictureTransform
// setPictureTransformRequest writes a SetPictureTransform request to a byte slice.
func setPictureTransformRequest(opcode uint8, Picture Picture, Transform Transform) []byte {
	size := 44
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 28 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Picture))
	b += 4

	{
		structBytes := Transform.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	return buf
}

// Trapezoids sends a checked request.
func Trapezoids(c *xgb.XConn, Op byte, Src Picture, Dst Picture, MaskFormat Pictformat, SrcX int16, SrcY int16, Traps []Trapezoid) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"Trapezoids\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.SendRecv(trapezoidsRequest(op, Op, Src, Dst, MaskFormat, SrcX, SrcY, Traps), nil)
}

// TrapezoidsUnchecked sends an unchecked request.
func TrapezoidsUnchecked(c *xgb.XConn, Op byte, Src Picture, Dst Picture, MaskFormat Pictformat, SrcX int16, SrcY int16, Traps []Trapezoid) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"Trapezoids\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.Send(trapezoidsRequest(op, Op, Src, Dst, MaskFormat, SrcX, SrcY, Traps))
}

// Write request to wire for Trapezoids
// trapezoidsRequest writes a Trapezoids request to a byte slice.
func trapezoidsRequest(opcode uint8, Op byte, Src Picture, Dst Picture, MaskFormat Pictformat, SrcX int16, SrcY int16, Traps []Trapezoid) []byte {
	size := internal.Pad4((24 + internal.Pad4((len(Traps) * 40))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 10 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	buf[b] = Op
	b += 1

	b += 3 // padding

	binary.LittleEndian.PutUint32(buf[b:], uint32(Src))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Dst))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(MaskFormat))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(SrcX))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(SrcY))
	b += 2

	b += TrapezoidListBytes(buf[b:], Traps)

	return buf
}

// TriFan sends a checked request.
func TriFan(c *xgb.XConn, Op byte, Src Picture, Dst Picture, MaskFormat Pictformat, SrcX int16, SrcY int16, Points []Pointfix) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"TriFan\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.SendRecv(triFanRequest(op, Op, Src, Dst, MaskFormat, SrcX, SrcY, Points), nil)
}

// TriFanUnchecked sends an unchecked request.
func TriFanUnchecked(c *xgb.XConn, Op byte, Src Picture, Dst Picture, MaskFormat Pictformat, SrcX int16, SrcY int16, Points []Pointfix) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"TriFan\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.Send(triFanRequest(op, Op, Src, Dst, MaskFormat, SrcX, SrcY, Points))
}

// Write request to wire for TriFan
// triFanRequest writes a TriFan request to a byte slice.
func triFanRequest(opcode uint8, Op byte, Src Picture, Dst Picture, MaskFormat Pictformat, SrcX int16, SrcY int16, Points []Pointfix) []byte {
	size := internal.Pad4((24 + internal.Pad4((len(Points) * 8))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 13 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	buf[b] = Op
	b += 1

	b += 3 // padding

	binary.LittleEndian.PutUint32(buf[b:], uint32(Src))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Dst))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(MaskFormat))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(SrcX))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(SrcY))
	b += 2

	b += PointfixListBytes(buf[b:], Points)

	return buf
}

// TriStrip sends a checked request.
func TriStrip(c *xgb.XConn, Op byte, Src Picture, Dst Picture, MaskFormat Pictformat, SrcX int16, SrcY int16, Points []Pointfix) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"TriStrip\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.SendRecv(triStripRequest(op, Op, Src, Dst, MaskFormat, SrcX, SrcY, Points), nil)
}

// TriStripUnchecked sends an unchecked request.
func TriStripUnchecked(c *xgb.XConn, Op byte, Src Picture, Dst Picture, MaskFormat Pictformat, SrcX int16, SrcY int16, Points []Pointfix) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"TriStrip\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.Send(triStripRequest(op, Op, Src, Dst, MaskFormat, SrcX, SrcY, Points))
}

// Write request to wire for TriStrip
// triStripRequest writes a TriStrip request to a byte slice.
func triStripRequest(opcode uint8, Op byte, Src Picture, Dst Picture, MaskFormat Pictformat, SrcX int16, SrcY int16, Points []Pointfix) []byte {
	size := internal.Pad4((24 + internal.Pad4((len(Points) * 8))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 12 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	buf[b] = Op
	b += 1

	b += 3 // padding

	binary.LittleEndian.PutUint32(buf[b:], uint32(Src))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Dst))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(MaskFormat))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(SrcX))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(SrcY))
	b += 2

	b += PointfixListBytes(buf[b:], Points)

	return buf
}

// Triangles sends a checked request.
func Triangles(c *xgb.XConn, Op byte, Src Picture, Dst Picture, MaskFormat Pictformat, SrcX int16, SrcY int16, Triangles []Triangle) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"Triangles\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.SendRecv(trianglesRequest(op, Op, Src, Dst, MaskFormat, SrcX, SrcY, Triangles), nil)
}

// TrianglesUnchecked sends an unchecked request.
func TrianglesUnchecked(c *xgb.XConn, Op byte, Src Picture, Dst Picture, MaskFormat Pictformat, SrcX int16, SrcY int16, Triangles []Triangle) error {
	op, ok := c.Ext("RENDER")
	if !ok {
		return errors.New("cannot issue request \"Triangles\" using the uninitialized extension \"RENDER\". render.Register(xconn) must be called first.")
	}
	return c.Send(trianglesRequest(op, Op, Src, Dst, MaskFormat, SrcX, SrcY, Triangles))
}

// Write request to wire for Triangles
// trianglesRequest writes a Triangles request to a byte slice.
func trianglesRequest(opcode uint8, Op byte, Src Picture, Dst Picture, MaskFormat Pictformat, SrcX int16, SrcY int16, Triangles []Triangle) []byte {
	size := internal.Pad4((24 + internal.Pad4((len(Triangles) * 24))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 11 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	buf[b] = Op
	b += 1

	b += 3 // padding

	binary.LittleEndian.PutUint32(buf[b:], uint32(Src))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Dst))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(MaskFormat))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(SrcX))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(SrcY))
	b += 2

	b += TriangleListBytes(buf[b:], Triangles)

	return buf
}
