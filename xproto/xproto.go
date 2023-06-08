// FILE GENERATED AUTOMATICALLY FROM "xproto.xml"
package xproto

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
	_ "unsafe"

	"codeberg.org/gruf/go-xgb"
	"codeberg.org/gruf/go-xgb/internal"
)

var (
	// generated index maps of defined event and error numbers -> unmarshalers.
	eventFuncs = make(map[uint8]xgb.EventUnmarshaler)
	errorFuncs = make(map[uint8]xgb.ErrorUnmarshaler)
)

// registerEvent will register an event unmarshaler in global map, panics on overlap
func registerEvent(n uint8, fn xgb.EventUnmarshaler) {
	if _, ok := eventFuncs[n]; ok {
		panic("BUG: overlapping event unmarshaler")
	}
	eventFuncs[n] = fn
}

// registerError will register an error unmarshaler in global map, panics on overlap
func registerError(n uint8, fn xgb.ErrorUnmarshaler) {
	if _, ok := errorFuncs[n]; ok {
		panic("BUG: overlapping error unmarshaler")
	}
	errorFuncs[n] = fn
}

// sorcery to give us access to package-private functions.
//
//go:linkname xproto_init codeberg.org/gruf/go-xgb.xproto_init
func xproto_init(*xgb.XConn, map[uint8]xgb.EventUnmarshaler, map[uint8]xgb.ErrorUnmarshaler) error

// Setup ...
func Setup(xconn *xgb.XConn, buf []byte) (*SetupInfo, error) {
	// Register ourselves with the X server connection
	if err := xproto_init(xconn, eventFuncs, errorFuncs); err != nil {
		return nil, err
	}

	info := &SetupInfo{}

	// Read setup information from buf
	_ = SetupInfoRead(buf, info)

	return info, nil
}

// BadAccess is the error number for a BadAccess.
const BadAccess = 10

type AccessError RequestError

// AccessErrorNew constructs a AccessError value that implements xgb.Error from a byte slice.
func UnmarshalAccessError(buf []byte) (xgb.XError, error) {
	return UnmarshalRequestError(buf)
}

// SequenceId returns the sequence id attached to the BadAccess error.
// This is mostly used internally.
func (err AccessError) SeqID() uint16 {
	return err.Sequence
}

// BadId returns the 'BadValue' number if one exists for the BadAccess error. If no bad value exists, 0 is returned.
func (err AccessError) BadID() uint32 {
	return err.BadValue
}

// Error returns a rudimentary string representation of the BadAccess error.
func (err AccessError) Error() string {
	var buf strings.Builder

	buf.WriteString("BadAccess{")
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
	registerError(10, UnmarshalAccessError)
}

const (
	AccessControlDisable = 0
	AccessControlEnable  = 1
)

// BadAlloc is the error number for a BadAlloc.
const BadAlloc = 11

type AllocError RequestError

// AllocErrorNew constructs a AllocError value that implements xgb.Error from a byte slice.
func UnmarshalAllocError(buf []byte) (xgb.XError, error) {
	return UnmarshalRequestError(buf)
}

// SequenceId returns the sequence id attached to the BadAlloc error.
// This is mostly used internally.
func (err AllocError) SeqID() uint16 {
	return err.Sequence
}

// BadId returns the 'BadValue' number if one exists for the BadAlloc error. If no bad value exists, 0 is returned.
func (err AllocError) BadID() uint32 {
	return err.BadValue
}

// Error returns a rudimentary string representation of the BadAlloc error.
func (err AllocError) Error() string {
	var buf strings.Builder

	buf.WriteString("BadAlloc{")
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
	registerError(11, UnmarshalAllocError)
}

const (
	AllowAsyncPointer   = 0
	AllowSyncPointer    = 1
	AllowReplayPointer  = 2
	AllowAsyncKeyboard  = 3
	AllowSyncKeyboard   = 4
	AllowReplayKeyboard = 5
	AllowAsyncBoth      = 6
	AllowSyncBoth       = 7
)

type Arc struct {
	X      int16
	Y      int16
	Width  uint16
	Height uint16
	Angle1 int16
	Angle2 int16
}

// ArcRead reads a byte slice into a Arc value.
func ArcRead(buf []byte, v *Arc) int {
	b := 0

	v.X = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.Y = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.Width = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Height = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Angle1 = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.Angle2 = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	return b
}

// ArcReadList reads a byte slice into a list of Arc values.
func ArcReadList(buf []byte, dest []Arc) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Arc{}
		b += ArcRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a Arc value to a byte slice.
func (v Arc) Bytes() []byte {
	buf := make([]byte, 12)
	b := 0

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.X))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.Y))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Width)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Height)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.Angle1))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.Angle2))
	b += 2

	return buf[:b]
}

// ArcListBytes writes a list of Arc values to a byte slice.
func ArcListBytes(buf []byte, list []Arc) int {
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
	ArcModeChord    = 0
	ArcModePieSlice = 1
)

type Atom uint32

func NewAtomID(c *xgb.XConn) Atom {
	id := c.NewXID()
	return Atom(id)
}

const (
	AtomNone               = 0
	AtomAny                = 0
	AtomPrimary            = 1
	AtomSecondary          = 2
	AtomArc                = 3
	AtomAtom               = 4
	AtomBitmap             = 5
	AtomCardinal           = 6
	AtomColormap           = 7
	AtomCursor             = 8
	AtomCutBuffer0         = 9
	AtomCutBuffer1         = 10
	AtomCutBuffer2         = 11
	AtomCutBuffer3         = 12
	AtomCutBuffer4         = 13
	AtomCutBuffer5         = 14
	AtomCutBuffer6         = 15
	AtomCutBuffer7         = 16
	AtomDrawable           = 17
	AtomFont               = 18
	AtomInteger            = 19
	AtomPixmap             = 20
	AtomPoint              = 21
	AtomRectangle          = 22
	AtomResourceManager    = 23
	AtomRgbColorMap        = 24
	AtomRgbBestMap         = 25
	AtomRgbBlueMap         = 26
	AtomRgbDefaultMap      = 27
	AtomRgbGrayMap         = 28
	AtomRgbGreenMap        = 29
	AtomRgbRedMap          = 30
	AtomString             = 31
	AtomVisualid           = 32
	AtomWindow             = 33
	AtomWmCommand          = 34
	AtomWmHints            = 35
	AtomWmClientMachine    = 36
	AtomWmIconName         = 37
	AtomWmIconSize         = 38
	AtomWmName             = 39
	AtomWmNormalHints      = 40
	AtomWmSizeHints        = 41
	AtomWmZoomHints        = 42
	AtomMinSpace           = 43
	AtomNormSpace          = 44
	AtomMaxSpace           = 45
	AtomEndSpace           = 46
	AtomSuperscriptX       = 47
	AtomSuperscriptY       = 48
	AtomSubscriptX         = 49
	AtomSubscriptY         = 50
	AtomUnderlinePosition  = 51
	AtomUnderlineThickness = 52
	AtomStrikeoutAscent    = 53
	AtomStrikeoutDescent   = 54
	AtomItalicAngle        = 55
	AtomXHeight            = 56
	AtomQuadWidth          = 57
	AtomWeight             = 58
	AtomPointSize          = 59
	AtomResolution         = 60
	AtomCopyright          = 61
	AtomNotice             = 62
	AtomFontName           = 63
	AtomFamilyName         = 64
	AtomFullName           = 65
	AtomCapHeight          = 66
	AtomWmClass            = 67
	AtomWmTransientFor     = 68
)

// BadAtom is the error number for a BadAtom.
const BadAtom = 5

type AtomError ValueError

// AtomErrorNew constructs a AtomError value that implements xgb.Error from a byte slice.
func UnmarshalAtomError(buf []byte) (xgb.XError, error) {
	return UnmarshalValueError(buf)
}

// SequenceId returns the sequence id attached to the BadAtom error.
// This is mostly used internally.
func (err AtomError) SeqID() uint16 {
	return err.Sequence
}

// BadId returns the 'BadValue' number if one exists for the BadAtom error. If no bad value exists, 0 is returned.
func (err AtomError) BadID() uint32 {
	return err.BadValue
}

// Error returns a rudimentary string representation of the BadAtom error.
func (err AtomError) Error() string {
	var buf strings.Builder

	buf.WriteString("BadAtom{")
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
	registerError(5, UnmarshalAtomError)
}

const (
	AutoRepeatModeOff     = 0
	AutoRepeatModeOn      = 1
	AutoRepeatModeDefault = 2
)

const (
	BackPixmapNone           = 0
	BackPixmapParentRelative = 1
)

const (
	BackingStoreNotUseful  = 0
	BackingStoreWhenMapped = 1
	BackingStoreAlways     = 2
)

const (
	BlankingNotPreferred = 0
	BlankingPreferred    = 1
	BlankingDefault      = 2
)

type Bool32 uint32

type Button byte

const (
	ButtonIndexAny = 0
	ButtonIndex1   = 1
	ButtonIndex2   = 2
	ButtonIndex3   = 3
	ButtonIndex4   = 4
	ButtonIndex5   = 5
)

const (
	ButtonMask1   = 256
	ButtonMask2   = 512
	ButtonMask3   = 1024
	ButtonMask4   = 2048
	ButtonMask5   = 4096
	ButtonMaskAny = 32768
)

// ButtonPress is the event number for a ButtonPressEvent.
const ButtonPress = 4

type ButtonPressEvent struct {
	Sequence   uint16
	Detail     Button
	Time       Timestamp
	Root       Window
	Event      Window
	Child      Window
	RootX      int16
	RootY      int16
	EventX     int16
	EventY     int16
	State      uint16
	SameScreen bool
	// padding: 1 bytes
}

// UnmarshalButtonPressEvent constructs a ButtonPressEvent value that implements xgb.Event from a byte slice.
func UnmarshalButtonPressEvent(buf []byte) (xgb.XEvent, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"ButtonPressEvent\"", len(buf))
	}

	v := ButtonPressEvent{}
	b := 1 // don't read event number

	v.Detail = Button(buf[b])
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Time = Timestamp(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Root = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Event = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Child = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.RootX = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.RootY = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.EventX = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.EventY = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.State = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.SameScreen = (buf[b] == 1)
	b += 1

	b += 1 // padding

	return v, nil
}

// Bytes writes a ButtonPressEvent value to a byte slice.
func (v ButtonPressEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 4
	b += 1

	buf[b] = uint8(v.Detail)
	b += 1

	b += 2 // skip sequence number

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Time))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Root))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Event))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Child))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.RootX))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.RootY))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.EventX))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.EventY))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.State)
	b += 2

	if v.SameScreen {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	b += 1 // padding

	return buf
}

// SeqID returns the sequence id attached to the ButtonPress event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v ButtonPressEvent) SeqID() uint16 {
	return v.Sequence
}

func init() { registerEvent(4, UnmarshalButtonPressEvent) }

// ButtonRelease is the event number for a ButtonReleaseEvent.
const ButtonRelease = 5

type ButtonReleaseEvent ButtonPressEvent

// ButtonReleaseEventNew constructs a ButtonReleaseEvent value that implements xgb.Event from a byte slice.
func UnmarshalButtonReleaseEvent(buf []byte) (xgb.XEvent, error) {
	ev, err := UnmarshalButtonPressEvent(buf)
	return ButtonReleaseEvent(ev.(ButtonPressEvent)), err
}

// Bytes writes a ButtonReleaseEvent value to a byte slice.
func (v ButtonReleaseEvent) Bytes() []byte {
	buf := ButtonPressEvent(v).Bytes()
	buf[0] = 5
	return buf
}

// SeqID returns the sequence id attached to the ButtonRelease event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v ButtonReleaseEvent) SeqID() uint16 {
	return v.Sequence
}

func init() {
	registerEvent(5, UnmarshalButtonReleaseEvent)
}

const (
	CapStyleNotLast    = 0
	CapStyleButt       = 1
	CapStyleRound      = 2
	CapStyleProjecting = 3
)

type Char2b struct {
	Byte1 byte
	Byte2 byte
}

// Char2bRead reads a byte slice into a Char2b value.
func Char2bRead(buf []byte, v *Char2b) int {
	b := 0

	v.Byte1 = buf[b]
	b += 1

	v.Byte2 = buf[b]
	b += 1

	return b
}

// Char2bReadList reads a byte slice into a list of Char2b values.
func Char2bReadList(buf []byte, dest []Char2b) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Char2b{}
		b += Char2bRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a Char2b value to a byte slice.
func (v Char2b) Bytes() []byte {
	buf := make([]byte, 2)
	b := 0

	buf[b] = v.Byte1
	b += 1

	buf[b] = v.Byte2
	b += 1

	return buf[:b]
}

// Char2bListBytes writes a list of Char2b values to a byte slice.
func Char2bListBytes(buf []byte, list []Char2b) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

type Charinfo struct {
	LeftSideBearing  int16
	RightSideBearing int16
	CharacterWidth   int16
	Ascent           int16
	Descent          int16
	Attributes       uint16
}

// CharinfoRead reads a byte slice into a Charinfo value.
func CharinfoRead(buf []byte, v *Charinfo) int {
	b := 0

	v.LeftSideBearing = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.RightSideBearing = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.CharacterWidth = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.Ascent = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.Descent = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.Attributes = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return b
}

// CharinfoReadList reads a byte slice into a list of Charinfo values.
func CharinfoReadList(buf []byte, dest []Charinfo) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Charinfo{}
		b += CharinfoRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a Charinfo value to a byte slice.
func (v Charinfo) Bytes() []byte {
	buf := make([]byte, 12)
	b := 0

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.LeftSideBearing))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.RightSideBearing))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.CharacterWidth))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.Ascent))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.Descent))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Attributes)
	b += 2

	return buf[:b]
}

// CharinfoListBytes writes a list of Charinfo values to a byte slice.
func CharinfoListBytes(buf []byte, list []Charinfo) int {
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
	CirculateRaiseLowest  = 0
	CirculateLowerHighest = 1
)

// CirculateNotify is the event number for a CirculateNotifyEvent.
const CirculateNotify = 26

type CirculateNotifyEvent struct {
	Sequence uint16
	// padding: 1 bytes
	Event  Window
	Window Window
	// padding: 4 bytes
	Place byte
	// padding: 3 bytes
}

// UnmarshalCirculateNotifyEvent constructs a CirculateNotifyEvent value that implements xgb.Event from a byte slice.
func UnmarshalCirculateNotifyEvent(buf []byte) (xgb.XEvent, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"CirculateNotifyEvent\"", len(buf))
	}

	v := CirculateNotifyEvent{}
	b := 1 // don't read event number

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Event = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Window = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	b += 4 // padding

	v.Place = buf[b]
	b += 1

	b += 3 // padding

	return v, nil
}

// Bytes writes a CirculateNotifyEvent value to a byte slice.
func (v CirculateNotifyEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 26
	b += 1

	b += 1 // padding

	b += 2 // skip sequence number

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Event))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Window))
	b += 4

	b += 4 // padding

	buf[b] = v.Place
	b += 1

	b += 3 // padding

	return buf
}

// SeqID returns the sequence id attached to the CirculateNotify event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v CirculateNotifyEvent) SeqID() uint16 {
	return v.Sequence
}

func init() { registerEvent(26, UnmarshalCirculateNotifyEvent) }

// CirculateRequest is the event number for a CirculateRequestEvent.
const CirculateRequest = 27

type CirculateRequestEvent CirculateNotifyEvent

// CirculateRequestEventNew constructs a CirculateRequestEvent value that implements xgb.Event from a byte slice.
func UnmarshalCirculateRequestEvent(buf []byte) (xgb.XEvent, error) {
	ev, err := UnmarshalCirculateNotifyEvent(buf)
	return CirculateRequestEvent(ev.(CirculateNotifyEvent)), err
}

// Bytes writes a CirculateRequestEvent value to a byte slice.
func (v CirculateRequestEvent) Bytes() []byte {
	buf := CirculateNotifyEvent(v).Bytes()
	buf[0] = 27
	return buf
}

// SeqID returns the sequence id attached to the CirculateRequest event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v CirculateRequestEvent) SeqID() uint16 {
	return v.Sequence
}

func init() {
	registerEvent(27, UnmarshalCirculateRequestEvent)
}

// ClientMessage is the event number for a ClientMessageEvent.
const ClientMessage = 33

type ClientMessageEvent struct {
	Sequence uint16
	Format   byte
	Window   Window
	Type     Atom
	Data     ClientMessageDataUnion
}

// UnmarshalClientMessageEvent constructs a ClientMessageEvent value that implements xgb.Event from a byte slice.
func UnmarshalClientMessageEvent(buf []byte) (xgb.XEvent, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"ClientMessageEvent\"", len(buf))
	}

	v := ClientMessageEvent{}
	b := 1 // don't read event number

	v.Format = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Window = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Type = Atom(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Data = ClientMessageDataUnion{}
	b += ClientMessageDataUnionRead(buf[b:], &v.Data)

	return v, nil
}

// Bytes writes a ClientMessageEvent value to a byte slice.
func (v ClientMessageEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 33
	b += 1

	buf[b] = v.Format
	b += 1

	b += 2 // skip sequence number

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Window))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Type))
	b += 4

	{
		unionBytes := v.Data.Bytes()
		copy(buf[b:], unionBytes)
		b += len(unionBytes)
	}

	return buf
}

// SeqID returns the sequence id attached to the ClientMessage event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v ClientMessageEvent) SeqID() uint16 {
	return v.Sequence
}

func init() { registerEvent(33, UnmarshalClientMessageEvent) }

// ClientMessageDataUnion is a representation of the ClientMessageDataUnion union type.
// Note that to *create* a Union, you should *never* create
// this struct directly (unless you know what you're doing).
// Instead use one of the following constructors for 'ClientMessageDataUnion':
//     ClientMessageDataUnionData8New(Data8 []byte) ClientMessageDataUnion
//     ClientMessageDataUnionData16New(Data16 []uint16) ClientMessageDataUnion
//     ClientMessageDataUnionData32New(Data32 []uint32) ClientMessageDataUnion
type ClientMessageDataUnion struct {
	Data8  []byte   // size: 20
	Data16 []uint16 // size: 20
	Data32 []uint32 // size: 20
}

// ClientMessageDataUnionData8New constructs a new ClientMessageDataUnion union type with the Data8 field.
func ClientMessageDataUnionData8New(Data8 []byte) ClientMessageDataUnion {
	var b int
	buf := make([]byte, 20)

	copy(buf[b:], Data8[:20])
	b += int(20)

	// Create the Union type
	v := ClientMessageDataUnion{}

	// Now copy buf into all fields

	b = 0 // always read the same bytes
	v.Data8 = make([]byte, 20)
	copy(v.Data8[:20], buf[b:])
	b += int(20)

	b = 0 // always read the same bytes
	v.Data16 = make([]uint16, 10)
	for i := 0; i < int(10); i++ {
		v.Data16[i] = binary.LittleEndian.Uint16(buf[b:])
		b += 2
	}

	b = 0 // always read the same bytes
	v.Data32 = make([]uint32, 5)
	for i := 0; i < int(5); i++ {
		v.Data32[i] = binary.LittleEndian.Uint32(buf[b:])
		b += 4
	}

	return v
}

// ClientMessageDataUnionData16New constructs a new ClientMessageDataUnion union type with the Data16 field.
func ClientMessageDataUnionData16New(Data16 []uint16) ClientMessageDataUnion {
	var b int
	buf := make([]byte, 20)

	for i := 0; i < int(10); i++ {
		binary.LittleEndian.PutUint16(buf[b:], Data16[i])
		b += 2
	}

	// Create the Union type
	v := ClientMessageDataUnion{}

	// Now copy buf into all fields

	b = 0 // always read the same bytes
	v.Data8 = make([]byte, 20)
	copy(v.Data8[:20], buf[b:])
	b += int(20)

	b = 0 // always read the same bytes
	v.Data16 = make([]uint16, 10)
	for i := 0; i < int(10); i++ {
		v.Data16[i] = binary.LittleEndian.Uint16(buf[b:])
		b += 2
	}

	b = 0 // always read the same bytes
	v.Data32 = make([]uint32, 5)
	for i := 0; i < int(5); i++ {
		v.Data32[i] = binary.LittleEndian.Uint32(buf[b:])
		b += 4
	}

	return v
}

// ClientMessageDataUnionData32New constructs a new ClientMessageDataUnion union type with the Data32 field.
func ClientMessageDataUnionData32New(Data32 []uint32) ClientMessageDataUnion {
	var b int
	buf := make([]byte, 20)

	for i := 0; i < int(5); i++ {
		binary.LittleEndian.PutUint32(buf[b:], Data32[i])
		b += 4
	}

	// Create the Union type
	v := ClientMessageDataUnion{}

	// Now copy buf into all fields

	b = 0 // always read the same bytes
	v.Data8 = make([]byte, 20)
	copy(v.Data8[:20], buf[b:])
	b += int(20)

	b = 0 // always read the same bytes
	v.Data16 = make([]uint16, 10)
	for i := 0; i < int(10); i++ {
		v.Data16[i] = binary.LittleEndian.Uint16(buf[b:])
		b += 2
	}

	b = 0 // always read the same bytes
	v.Data32 = make([]uint32, 5)
	for i := 0; i < int(5); i++ {
		v.Data32[i] = binary.LittleEndian.Uint32(buf[b:])
		b += 4
	}

	return v
}

// ClientMessageDataUnionRead reads a byte slice into a ClientMessageDataUnion value.
func ClientMessageDataUnionRead(buf []byte, v *ClientMessageDataUnion) int {
	var b int

	b = 0 // re-read the same bytes
	v.Data8 = make([]byte, 20)
	copy(v.Data8[:20], buf[b:])
	b += int(20)

	b = 0 // re-read the same bytes
	v.Data16 = make([]uint16, 10)
	for i := 0; i < int(10); i++ {
		v.Data16[i] = binary.LittleEndian.Uint16(buf[b:])
		b += 2
	}

	b = 0 // re-read the same bytes
	v.Data32 = make([]uint32, 5)
	for i := 0; i < int(5); i++ {
		v.Data32[i] = binary.LittleEndian.Uint32(buf[b:])
		b += 4
	}

	return 20
}

// ClientMessageDataUnionReadList reads a byte slice into a list of ClientMessageDataUnion values.
func ClientMessageDataUnionReadList(buf []byte, dest []ClientMessageDataUnion) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = ClientMessageDataUnion{}
		b += ClientMessageDataUnionRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a ClientMessageDataUnion value to a byte slice.
// Each field in a union must contain the same data.
// So simply pick the first field and write that to the wire.
func (v ClientMessageDataUnion) Bytes() []byte {
	buf := make([]byte, 20)
	b := 0

	copy(buf[b:], v.Data8[:20])
	b += int(20)
	return buf
}

// ClientMessageDataUnionListBytes writes a list of ClientMessageDataUnion values to a byte slice.
func ClientMessageDataUnionListBytes(buf []byte, list []ClientMessageDataUnion) int {
	b := 0
	var unionBytes []byte
	for _, item := range list {
		unionBytes = item.Bytes()
		copy(buf[b:], unionBytes)
		b += internal.Pad4(len(unionBytes))
	}
	return b
}

const (
	ClipOrderingUnsorted = 0
	ClipOrderingYSorted  = 1
	ClipOrderingYXSorted = 2
	ClipOrderingYXBanded = 3
)

const (
	CloseDownDestroyAll      = 0
	CloseDownRetainPermanent = 1
	CloseDownRetainTemporary = 2
)

const (
	ColorFlagRed   = 1
	ColorFlagGreen = 2
	ColorFlagBlue  = 4
)

type Coloritem struct {
	Pixel uint32
	Red   uint16
	Green uint16
	Blue  uint16
	Flags byte
	// padding: 1 bytes
}

// ColoritemRead reads a byte slice into a Coloritem value.
func ColoritemRead(buf []byte, v *Coloritem) int {
	b := 0

	v.Pixel = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.Red = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Green = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Blue = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Flags = buf[b]
	b += 1

	b += 1 // padding

	return b
}

// ColoritemReadList reads a byte slice into a list of Coloritem values.
func ColoritemReadList(buf []byte, dest []Coloritem) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Coloritem{}
		b += ColoritemRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a Coloritem value to a byte slice.
func (v Coloritem) Bytes() []byte {
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

	buf[b] = v.Flags
	b += 1

	b += 1 // padding

	return buf[:b]
}

// ColoritemListBytes writes a list of Coloritem values to a byte slice.
func ColoritemListBytes(buf []byte, list []Coloritem) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

type Colormap uint32

func NewColormapID(c *xgb.XConn) Colormap {
	id := c.NewXID()
	return Colormap(id)
}

const (
	ColormapNone = 0
)

// BadColormap is the error number for a BadColormap.
const BadColormap = 12

type ColormapError ValueError

// ColormapErrorNew constructs a ColormapError value that implements xgb.Error from a byte slice.
func UnmarshalColormapError(buf []byte) (xgb.XError, error) {
	return UnmarshalValueError(buf)
}

// SequenceId returns the sequence id attached to the BadColormap error.
// This is mostly used internally.
func (err ColormapError) SeqID() uint16 {
	return err.Sequence
}

// BadId returns the 'BadValue' number if one exists for the BadColormap error. If no bad value exists, 0 is returned.
func (err ColormapError) BadID() uint32 {
	return err.BadValue
}

// Error returns a rudimentary string representation of the BadColormap error.
func (err ColormapError) Error() string {
	var buf strings.Builder

	buf.WriteString("BadColormap{")
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
	registerError(12, UnmarshalColormapError)
}

const (
	ColormapAllocNone = 0
	ColormapAllocAll  = 1
)

// ColormapNotify is the event number for a ColormapNotifyEvent.
const ColormapNotify = 32

type ColormapNotifyEvent struct {
	Sequence uint16
	// padding: 1 bytes
	Window   Window
	Colormap Colormap
	New      bool
	State    byte
	// padding: 2 bytes
}

// UnmarshalColormapNotifyEvent constructs a ColormapNotifyEvent value that implements xgb.Event from a byte slice.
func UnmarshalColormapNotifyEvent(buf []byte) (xgb.XEvent, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"ColormapNotifyEvent\"", len(buf))
	}

	v := ColormapNotifyEvent{}
	b := 1 // don't read event number

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Window = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Colormap = Colormap(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.New = (buf[b] == 1)
	b += 1

	v.State = buf[b]
	b += 1

	b += 2 // padding

	return v, nil
}

// Bytes writes a ColormapNotifyEvent value to a byte slice.
func (v ColormapNotifyEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 32
	b += 1

	b += 1 // padding

	b += 2 // skip sequence number

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Window))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Colormap))
	b += 4

	if v.New {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	buf[b] = v.State
	b += 1

	b += 2 // padding

	return buf
}

// SeqID returns the sequence id attached to the ColormapNotify event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v ColormapNotifyEvent) SeqID() uint16 {
	return v.Sequence
}

func init() { registerEvent(32, UnmarshalColormapNotifyEvent) }

const (
	ColormapStateUninstalled = 0
	ColormapStateInstalled   = 1
)

const (
	ConfigWindowX           = 1
	ConfigWindowY           = 2
	ConfigWindowWidth       = 4
	ConfigWindowHeight      = 8
	ConfigWindowBorderWidth = 16
	ConfigWindowSibling     = 32
	ConfigWindowStackMode   = 64
)

// ConfigureNotify is the event number for a ConfigureNotifyEvent.
const ConfigureNotify = 22

type ConfigureNotifyEvent struct {
	Sequence uint16
	// padding: 1 bytes
	Event            Window
	Window           Window
	AboveSibling     Window
	X                int16
	Y                int16
	Width            uint16
	Height           uint16
	BorderWidth      uint16
	OverrideRedirect bool
	// padding: 1 bytes
}

// UnmarshalConfigureNotifyEvent constructs a ConfigureNotifyEvent value that implements xgb.Event from a byte slice.
func UnmarshalConfigureNotifyEvent(buf []byte) (xgb.XEvent, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"ConfigureNotifyEvent\"", len(buf))
	}

	v := ConfigureNotifyEvent{}
	b := 1 // don't read event number

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Event = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Window = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.AboveSibling = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.X = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.Y = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.Width = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Height = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.BorderWidth = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.OverrideRedirect = (buf[b] == 1)
	b += 1

	b += 1 // padding

	return v, nil
}

// Bytes writes a ConfigureNotifyEvent value to a byte slice.
func (v ConfigureNotifyEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 22
	b += 1

	b += 1 // padding

	b += 2 // skip sequence number

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Event))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Window))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.AboveSibling))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.X))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.Y))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Width)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Height)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.BorderWidth)
	b += 2

	if v.OverrideRedirect {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	b += 1 // padding

	return buf
}

// SeqID returns the sequence id attached to the ConfigureNotify event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v ConfigureNotifyEvent) SeqID() uint16 {
	return v.Sequence
}

func init() { registerEvent(22, UnmarshalConfigureNotifyEvent) }

// ConfigureRequest is the event number for a ConfigureRequestEvent.
const ConfigureRequest = 23

type ConfigureRequestEvent struct {
	Sequence    uint16
	StackMode   byte
	Parent      Window
	Window      Window
	Sibling     Window
	X           int16
	Y           int16
	Width       uint16
	Height      uint16
	BorderWidth uint16
	ValueMask   uint16
}

// UnmarshalConfigureRequestEvent constructs a ConfigureRequestEvent value that implements xgb.Event from a byte slice.
func UnmarshalConfigureRequestEvent(buf []byte) (xgb.XEvent, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"ConfigureRequestEvent\"", len(buf))
	}

	v := ConfigureRequestEvent{}
	b := 1 // don't read event number

	v.StackMode = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Parent = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Window = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Sibling = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.X = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.Y = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.Width = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Height = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.BorderWidth = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.ValueMask = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return v, nil
}

// Bytes writes a ConfigureRequestEvent value to a byte slice.
func (v ConfigureRequestEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 23
	b += 1

	buf[b] = v.StackMode
	b += 1

	b += 2 // skip sequence number

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Parent))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Window))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Sibling))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.X))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.Y))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Width)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Height)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.BorderWidth)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.ValueMask)
	b += 2

	return buf
}

// SeqID returns the sequence id attached to the ConfigureRequest event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v ConfigureRequestEvent) SeqID() uint16 {
	return v.Sequence
}

func init() { registerEvent(23, UnmarshalConfigureRequestEvent) }

const (
	CoordModeOrigin   = 0
	CoordModePrevious = 1
)

// CreateNotify is the event number for a CreateNotifyEvent.
const CreateNotify = 16

type CreateNotifyEvent struct {
	Sequence uint16
	// padding: 1 bytes
	Parent           Window
	Window           Window
	X                int16
	Y                int16
	Width            uint16
	Height           uint16
	BorderWidth      uint16
	OverrideRedirect bool
	// padding: 1 bytes
}

// UnmarshalCreateNotifyEvent constructs a CreateNotifyEvent value that implements xgb.Event from a byte slice.
func UnmarshalCreateNotifyEvent(buf []byte) (xgb.XEvent, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"CreateNotifyEvent\"", len(buf))
	}

	v := CreateNotifyEvent{}
	b := 1 // don't read event number

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Parent = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Window = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.X = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.Y = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.Width = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Height = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.BorderWidth = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.OverrideRedirect = (buf[b] == 1)
	b += 1

	b += 1 // padding

	return v, nil
}

// Bytes writes a CreateNotifyEvent value to a byte slice.
func (v CreateNotifyEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 16
	b += 1

	b += 1 // padding

	b += 2 // skip sequence number

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Parent))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Window))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.X))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.Y))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Width)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Height)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.BorderWidth)
	b += 2

	if v.OverrideRedirect {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	b += 1 // padding

	return buf
}

// SeqID returns the sequence id attached to the CreateNotify event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v CreateNotifyEvent) SeqID() uint16 {
	return v.Sequence
}

func init() { registerEvent(16, UnmarshalCreateNotifyEvent) }

type Cursor uint32

func NewCursorID(c *xgb.XConn) Cursor {
	id := c.NewXID()
	return Cursor(id)
}

const (
	CursorNone = 0
)

// BadCursor is the error number for a BadCursor.
const BadCursor = 6

type CursorError ValueError

// CursorErrorNew constructs a CursorError value that implements xgb.Error from a byte slice.
func UnmarshalCursorError(buf []byte) (xgb.XError, error) {
	return UnmarshalValueError(buf)
}

// SequenceId returns the sequence id attached to the BadCursor error.
// This is mostly used internally.
func (err CursorError) SeqID() uint16 {
	return err.Sequence
}

// BadId returns the 'BadValue' number if one exists for the BadCursor error. If no bad value exists, 0 is returned.
func (err CursorError) BadID() uint32 {
	return err.BadValue
}

// Error returns a rudimentary string representation of the BadCursor error.
func (err CursorError) Error() string {
	var buf strings.Builder

	buf.WriteString("BadCursor{")
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
	registerError(6, UnmarshalCursorError)
}

const (
	CwBackPixmap       = 1
	CwBackPixel        = 2
	CwBorderPixmap     = 4
	CwBorderPixel      = 8
	CwBitGravity       = 16
	CwWinGravity       = 32
	CwBackingStore     = 64
	CwBackingPlanes    = 128
	CwBackingPixel     = 256
	CwOverrideRedirect = 512
	CwSaveUnder        = 1024
	CwEventMask        = 2048
	CwDontPropagate    = 4096
	CwColormap         = 8192
	CwCursor           = 16384
)

type DepthInfo struct {
	Depth byte
	// padding: 1 bytes
	VisualsLen uint16
	// padding: 4 bytes
	Visuals []VisualInfo // size: internal.Pad4((int(VisualsLen) * 24))
}

// DepthInfoRead reads a byte slice into a DepthInfo value.
func DepthInfoRead(buf []byte, v *DepthInfo) int {
	b := 0

	v.Depth = buf[b]
	b += 1

	b += 1 // padding

	v.VisualsLen = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	b += 4 // padding

	v.Visuals = make([]VisualInfo, v.VisualsLen)
	b += VisualInfoReadList(buf[b:], v.Visuals)

	return b
}

// DepthInfoReadList reads a byte slice into a list of DepthInfo values.
func DepthInfoReadList(buf []byte, dest []DepthInfo) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = DepthInfo{}
		b += DepthInfoRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a DepthInfo value to a byte slice.
func (v DepthInfo) Bytes() []byte {
	buf := make([]byte, (8 + internal.Pad4((int(v.VisualsLen) * 24))))
	b := 0

	buf[b] = v.Depth
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], v.VisualsLen)
	b += 2

	b += 4 // padding

	b += VisualInfoListBytes(buf[b:], v.Visuals)

	return buf[:b]
}

// DepthInfoListBytes writes a list of DepthInfo values to a byte slice.
func DepthInfoListBytes(buf []byte, list []DepthInfo) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

// DepthInfoListSize computes the size (bytes) of a list of DepthInfo values.
func DepthInfoListSize(list []DepthInfo) int {
	size := 0
	for _, item := range list {
		size += (8 + internal.Pad4((int(item.VisualsLen) * 24)))
	}
	return size
}

// DestroyNotify is the event number for a DestroyNotifyEvent.
const DestroyNotify = 17

type DestroyNotifyEvent struct {
	Sequence uint16
	// padding: 1 bytes
	Event  Window
	Window Window
}

// UnmarshalDestroyNotifyEvent constructs a DestroyNotifyEvent value that implements xgb.Event from a byte slice.
func UnmarshalDestroyNotifyEvent(buf []byte) (xgb.XEvent, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"DestroyNotifyEvent\"", len(buf))
	}

	v := DestroyNotifyEvent{}
	b := 1 // don't read event number

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Event = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Window = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	return v, nil
}

// Bytes writes a DestroyNotifyEvent value to a byte slice.
func (v DestroyNotifyEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 17
	b += 1

	b += 1 // padding

	b += 2 // skip sequence number

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Event))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Window))
	b += 4

	return buf
}

// SeqID returns the sequence id attached to the DestroyNotify event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v DestroyNotifyEvent) SeqID() uint16 {
	return v.Sequence
}

func init() { registerEvent(17, UnmarshalDestroyNotifyEvent) }

type Drawable uint32

func NewDrawableID(c *xgb.XConn) Drawable {
	id := c.NewXID()
	return Drawable(id)
}

// BadDrawable is the error number for a BadDrawable.
const BadDrawable = 9

type DrawableError ValueError

// DrawableErrorNew constructs a DrawableError value that implements xgb.Error from a byte slice.
func UnmarshalDrawableError(buf []byte) (xgb.XError, error) {
	return UnmarshalValueError(buf)
}

// SequenceId returns the sequence id attached to the BadDrawable error.
// This is mostly used internally.
func (err DrawableError) SeqID() uint16 {
	return err.Sequence
}

// BadId returns the 'BadValue' number if one exists for the BadDrawable error. If no bad value exists, 0 is returned.
func (err DrawableError) BadID() uint32 {
	return err.BadValue
}

// Error returns a rudimentary string representation of the BadDrawable error.
func (err DrawableError) Error() string {
	var buf strings.Builder

	buf.WriteString("BadDrawable{")
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
	registerError(9, UnmarshalDrawableError)
}

// EnterNotify is the event number for a EnterNotifyEvent.
const EnterNotify = 7

type EnterNotifyEvent struct {
	Sequence        uint16
	Detail          byte
	Time            Timestamp
	Root            Window
	Event           Window
	Child           Window
	RootX           int16
	RootY           int16
	EventX          int16
	EventY          int16
	State           uint16
	Mode            byte
	SameScreenFocus byte
}

// UnmarshalEnterNotifyEvent constructs a EnterNotifyEvent value that implements xgb.Event from a byte slice.
func UnmarshalEnterNotifyEvent(buf []byte) (xgb.XEvent, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"EnterNotifyEvent\"", len(buf))
	}

	v := EnterNotifyEvent{}
	b := 1 // don't read event number

	v.Detail = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Time = Timestamp(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Root = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Event = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Child = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.RootX = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.RootY = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.EventX = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.EventY = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.State = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Mode = buf[b]
	b += 1

	v.SameScreenFocus = buf[b]
	b += 1

	return v, nil
}

// Bytes writes a EnterNotifyEvent value to a byte slice.
func (v EnterNotifyEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 7
	b += 1

	buf[b] = v.Detail
	b += 1

	b += 2 // skip sequence number

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Time))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Root))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Event))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Child))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.RootX))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.RootY))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.EventX))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.EventY))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.State)
	b += 2

	buf[b] = v.Mode
	b += 1

	buf[b] = v.SameScreenFocus
	b += 1

	return buf
}

// SeqID returns the sequence id attached to the EnterNotify event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v EnterNotifyEvent) SeqID() uint16 {
	return v.Sequence
}

func init() { registerEvent(7, UnmarshalEnterNotifyEvent) }

const (
	EventMaskNoEvent              = 0
	EventMaskKeyPress             = 1
	EventMaskKeyRelease           = 2
	EventMaskButtonPress          = 4
	EventMaskButtonRelease        = 8
	EventMaskEnterWindow          = 16
	EventMaskLeaveWindow          = 32
	EventMaskPointerMotion        = 64
	EventMaskPointerMotionHint    = 128
	EventMaskButton1Motion        = 256
	EventMaskButton2Motion        = 512
	EventMaskButton3Motion        = 1024
	EventMaskButton4Motion        = 2048
	EventMaskButton5Motion        = 4096
	EventMaskButtonMotion         = 8192
	EventMaskKeymapState          = 16384
	EventMaskExposure             = 32768
	EventMaskVisibilityChange     = 65536
	EventMaskStructureNotify      = 131072
	EventMaskResizeRedirect       = 262144
	EventMaskSubstructureNotify   = 524288
	EventMaskSubstructureRedirect = 1048576
	EventMaskFocusChange          = 2097152
	EventMaskPropertyChange       = 4194304
	EventMaskColorMapChange       = 8388608
	EventMaskOwnerGrabButton      = 16777216
)

// Expose is the event number for a ExposeEvent.
const Expose = 12

type ExposeEvent struct {
	Sequence uint16
	// padding: 1 bytes
	Window Window
	X      uint16
	Y      uint16
	Width  uint16
	Height uint16
	Count  uint16
	// padding: 2 bytes
}

// UnmarshalExposeEvent constructs a ExposeEvent value that implements xgb.Event from a byte slice.
func UnmarshalExposeEvent(buf []byte) (xgb.XEvent, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"ExposeEvent\"", len(buf))
	}

	v := ExposeEvent{}
	b := 1 // don't read event number

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Window = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.X = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Y = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Width = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Height = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Count = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	b += 2 // padding

	return v, nil
}

// Bytes writes a ExposeEvent value to a byte slice.
func (v ExposeEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 12
	b += 1

	b += 1 // padding

	b += 2 // skip sequence number

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Window))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], v.X)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Y)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Width)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Height)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Count)
	b += 2

	b += 2 // padding

	return buf
}

// SeqID returns the sequence id attached to the Expose event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v ExposeEvent) SeqID() uint16 {
	return v.Sequence
}

func init() { registerEvent(12, UnmarshalExposeEvent) }

const (
	ExposuresNotAllowed = 0
	ExposuresAllowed    = 1
	ExposuresDefault    = 2
)

const (
	FamilyInternet          = 0
	FamilyDECnet            = 1
	FamilyChaos             = 2
	FamilyServerInterpreted = 5
	FamilyInternet6         = 6
)

const (
	FillRuleEvenOdd = 0
	FillRuleWinding = 1
)

const (
	FillStyleSolid          = 0
	FillStyleTiled          = 1
	FillStyleStippled       = 2
	FillStyleOpaqueStippled = 3
)

// FocusIn is the event number for a FocusInEvent.
const FocusIn = 9

type FocusInEvent struct {
	Sequence uint16
	Detail   byte
	Event    Window
	Mode     byte
	// padding: 3 bytes
}

// UnmarshalFocusInEvent constructs a FocusInEvent value that implements xgb.Event from a byte slice.
func UnmarshalFocusInEvent(buf []byte) (xgb.XEvent, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"FocusInEvent\"", len(buf))
	}

	v := FocusInEvent{}
	b := 1 // don't read event number

	v.Detail = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Event = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Mode = buf[b]
	b += 1

	b += 3 // padding

	return v, nil
}

// Bytes writes a FocusInEvent value to a byte slice.
func (v FocusInEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 9
	b += 1

	buf[b] = v.Detail
	b += 1

	b += 2 // skip sequence number

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Event))
	b += 4

	buf[b] = v.Mode
	b += 1

	b += 3 // padding

	return buf
}

// SeqID returns the sequence id attached to the FocusIn event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v FocusInEvent) SeqID() uint16 {
	return v.Sequence
}

func init() { registerEvent(9, UnmarshalFocusInEvent) }

// FocusOut is the event number for a FocusOutEvent.
const FocusOut = 10

type FocusOutEvent FocusInEvent

// FocusOutEventNew constructs a FocusOutEvent value that implements xgb.Event from a byte slice.
func UnmarshalFocusOutEvent(buf []byte) (xgb.XEvent, error) {
	ev, err := UnmarshalFocusInEvent(buf)
	return FocusOutEvent(ev.(FocusInEvent)), err
}

// Bytes writes a FocusOutEvent value to a byte slice.
func (v FocusOutEvent) Bytes() []byte {
	buf := FocusInEvent(v).Bytes()
	buf[0] = 10
	return buf
}

// SeqID returns the sequence id attached to the FocusOut event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v FocusOutEvent) SeqID() uint16 {
	return v.Sequence
}

func init() {
	registerEvent(10, UnmarshalFocusOutEvent)
}

type Font uint32

func NewFontID(c *xgb.XConn) Font {
	id := c.NewXID()
	return Font(id)
}

// BadFont is the error number for a BadFont.
const BadFont = 7

type FontError ValueError

// FontErrorNew constructs a FontError value that implements xgb.Error from a byte slice.
func UnmarshalFontError(buf []byte) (xgb.XError, error) {
	return UnmarshalValueError(buf)
}

// SequenceId returns the sequence id attached to the BadFont error.
// This is mostly used internally.
func (err FontError) SeqID() uint16 {
	return err.Sequence
}

// BadId returns the 'BadValue' number if one exists for the BadFont error. If no bad value exists, 0 is returned.
func (err FontError) BadID() uint32 {
	return err.BadValue
}

// Error returns a rudimentary string representation of the BadFont error.
func (err FontError) Error() string {
	var buf strings.Builder

	buf.WriteString("BadFont{")
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
	registerError(7, UnmarshalFontError)
}

const (
	FontNone = 0
)

const (
	FontDrawLeftToRight = 0
	FontDrawRightToLeft = 1
)

type Fontable uint32

func NewFontableID(c *xgb.XConn) Fontable {
	id := c.NewXID()
	return Fontable(id)
}

type Fontprop struct {
	Name  Atom
	Value uint32
}

// FontpropRead reads a byte slice into a Fontprop value.
func FontpropRead(buf []byte, v *Fontprop) int {
	b := 0

	v.Name = Atom(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Value = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	return b
}

// FontpropReadList reads a byte slice into a list of Fontprop values.
func FontpropReadList(buf []byte, dest []Fontprop) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Fontprop{}
		b += FontpropRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a Fontprop value to a byte slice.
func (v Fontprop) Bytes() []byte {
	buf := make([]byte, 8)
	b := 0

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Name))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], v.Value)
	b += 4

	return buf[:b]
}

// FontpropListBytes writes a list of Fontprop values to a byte slice.
func FontpropListBytes(buf []byte, list []Fontprop) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

type Format struct {
	Depth        byte
	BitsPerPixel byte
	ScanlinePad  byte
	// padding: 5 bytes
}

// FormatRead reads a byte slice into a Format value.
func FormatRead(buf []byte, v *Format) int {
	b := 0

	v.Depth = buf[b]
	b += 1

	v.BitsPerPixel = buf[b]
	b += 1

	v.ScanlinePad = buf[b]
	b += 1

	b += 5 // padding

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

	buf[b] = v.Depth
	b += 1

	buf[b] = v.BitsPerPixel
	b += 1

	buf[b] = v.ScanlinePad
	b += 1

	b += 5 // padding

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

// BadGContext is the error number for a BadGContext.
const BadGContext = 13

type GContextError ValueError

// GContextErrorNew constructs a GContextError value that implements xgb.Error from a byte slice.
func UnmarshalGContextError(buf []byte) (xgb.XError, error) {
	return UnmarshalValueError(buf)
}

// SequenceId returns the sequence id attached to the BadGContext error.
// This is mostly used internally.
func (err GContextError) SeqID() uint16 {
	return err.Sequence
}

// BadId returns the 'BadValue' number if one exists for the BadGContext error. If no bad value exists, 0 is returned.
func (err GContextError) BadID() uint32 {
	return err.BadValue
}

// Error returns a rudimentary string representation of the BadGContext error.
func (err GContextError) Error() string {
	var buf strings.Builder

	buf.WriteString("BadGContext{")
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
	registerError(13, UnmarshalGContextError)
}

const (
	GcFunction           = 1
	GcPlaneMask          = 2
	GcForeground         = 4
	GcBackground         = 8
	GcLineWidth          = 16
	GcLineStyle          = 32
	GcCapStyle           = 64
	GcJoinStyle          = 128
	GcFillStyle          = 256
	GcFillRule           = 512
	GcTile               = 1024
	GcStipple            = 2048
	GcTileStippleOriginX = 4096
	GcTileStippleOriginY = 8192
	GcFont               = 16384
	GcSubwindowMode      = 32768
	GcGraphicsExposures  = 65536
	GcClipOriginX        = 131072
	GcClipOriginY        = 262144
	GcClipMask           = 524288
	GcDashOffset         = 1048576
	GcDashList           = 2097152
	GcArcMode            = 4194304
)

type Gcontext uint32

func NewGcontextID(c *xgb.XConn) Gcontext {
	id := c.NewXID()
	return Gcontext(id)
}

// GeGeneric is the event number for a GeGenericEvent.
const GeGeneric = 35

type GeGenericEvent struct {
	Sequence uint16
	// padding: 22 bytes
}

// UnmarshalGeGenericEvent constructs a GeGenericEvent value that implements xgb.Event from a byte slice.
func UnmarshalGeGenericEvent(buf []byte) (xgb.XEvent, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"GeGenericEvent\"", len(buf))
	}

	v := GeGenericEvent{}
	b := 1 // don't read event number

	b += 22 // padding

	return v, nil
}

// Bytes writes a GeGenericEvent value to a byte slice.
func (v GeGenericEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 35
	b += 1

	b += 22 // padding

	return buf
}

// SeqID returns the sequence id attached to the GeGeneric event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v GeGenericEvent) SeqID() uint16 {
	return v.Sequence
}

func init() { registerEvent(35, UnmarshalGeGenericEvent) }

const (
	GetPropertyTypeAny = 0
)

const (
	GrabAny = 0
)

const (
	GrabModeSync  = 0
	GrabModeAsync = 1
)

const (
	GrabStatusSuccess        = 0
	GrabStatusAlreadyGrabbed = 1
	GrabStatusInvalidTime    = 2
	GrabStatusNotViewable    = 3
	GrabStatusFrozen         = 4
)

// GraphicsExposure is the event number for a GraphicsExposureEvent.
const GraphicsExposure = 13

type GraphicsExposureEvent struct {
	Sequence uint16
	// padding: 1 bytes
	Drawable    Drawable
	X           uint16
	Y           uint16
	Width       uint16
	Height      uint16
	MinorOpcode uint16
	Count       uint16
	MajorOpcode byte
	// padding: 3 bytes
}

// UnmarshalGraphicsExposureEvent constructs a GraphicsExposureEvent value that implements xgb.Event from a byte slice.
func UnmarshalGraphicsExposureEvent(buf []byte) (xgb.XEvent, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"GraphicsExposureEvent\"", len(buf))
	}

	v := GraphicsExposureEvent{}
	b := 1 // don't read event number

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Drawable = Drawable(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.X = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Y = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Width = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Height = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.MinorOpcode = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Count = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.MajorOpcode = buf[b]
	b += 1

	b += 3 // padding

	return v, nil
}

// Bytes writes a GraphicsExposureEvent value to a byte slice.
func (v GraphicsExposureEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 13
	b += 1

	b += 1 // padding

	b += 2 // skip sequence number

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Drawable))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], v.X)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Y)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Width)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Height)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.MinorOpcode)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Count)
	b += 2

	buf[b] = v.MajorOpcode
	b += 1

	b += 3 // padding

	return buf
}

// SeqID returns the sequence id attached to the GraphicsExposure event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v GraphicsExposureEvent) SeqID() uint16 {
	return v.Sequence
}

func init() { registerEvent(13, UnmarshalGraphicsExposureEvent) }

const (
	GravityBitForget = 0
	GravityWinUnmap  = 0
	GravityNorthWest = 1
	GravityNorth     = 2
	GravityNorthEast = 3
	GravityWest      = 4
	GravityCenter    = 5
	GravityEast      = 6
	GravitySouthWest = 7
	GravitySouth     = 8
	GravitySouthEast = 9
	GravityStatic    = 10
)

// GravityNotify is the event number for a GravityNotifyEvent.
const GravityNotify = 24

type GravityNotifyEvent struct {
	Sequence uint16
	// padding: 1 bytes
	Event  Window
	Window Window
	X      int16
	Y      int16
}

// UnmarshalGravityNotifyEvent constructs a GravityNotifyEvent value that implements xgb.Event from a byte slice.
func UnmarshalGravityNotifyEvent(buf []byte) (xgb.XEvent, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"GravityNotifyEvent\"", len(buf))
	}

	v := GravityNotifyEvent{}
	b := 1 // don't read event number

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Event = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Window = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.X = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.Y = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	return v, nil
}

// Bytes writes a GravityNotifyEvent value to a byte slice.
func (v GravityNotifyEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 24
	b += 1

	b += 1 // padding

	b += 2 // skip sequence number

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Event))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Window))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.X))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.Y))
	b += 2

	return buf
}

// SeqID returns the sequence id attached to the GravityNotify event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v GravityNotifyEvent) SeqID() uint16 {
	return v.Sequence
}

func init() { registerEvent(24, UnmarshalGravityNotifyEvent) }

const (
	GxClear        = 0
	GxAnd          = 1
	GxAndReverse   = 2
	GxCopy         = 3
	GxAndInverted  = 4
	GxNoop         = 5
	GxXor          = 6
	GxOr           = 7
	GxNor          = 8
	GxEquiv        = 9
	GxInvert       = 10
	GxOrReverse    = 11
	GxCopyInverted = 12
	GxOrInverted   = 13
	GxNand         = 14
	GxSet          = 15
)

type Host struct {
	Family byte
	// padding: 1 bytes
	AddressLen uint16
	Address    []byte // size: internal.Pad4((int(AddressLen) * 1))
	// padding: 0 bytes
}

// HostRead reads a byte slice into a Host value.
func HostRead(buf []byte, v *Host) int {
	b := 0

	v.Family = buf[b]
	b += 1

	b += 1 // padding

	v.AddressLen = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Address = make([]byte, v.AddressLen)
	copy(v.Address[:v.AddressLen], buf[b:])
	b += int(v.AddressLen)

	b += 0 // padding

	return b
}

// HostReadList reads a byte slice into a list of Host values.
func HostReadList(buf []byte, dest []Host) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Host{}
		b += HostRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a Host value to a byte slice.
func (v Host) Bytes() []byte {
	buf := make([]byte, ((4 + internal.Pad4((int(v.AddressLen) * 1))) + 0))
	b := 0

	buf[b] = v.Family
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], v.AddressLen)
	b += 2

	copy(buf[b:], v.Address[:v.AddressLen])
	b += int(v.AddressLen)

	b += 0 // padding

	return buf[:b]
}

// HostListBytes writes a list of Host values to a byte slice.
func HostListBytes(buf []byte, list []Host) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

// HostListSize computes the size (bytes) of a list of Host values.
func HostListSize(list []Host) int {
	size := 0
	for _, item := range list {
		size += ((4 + internal.Pad4((int(item.AddressLen) * 1))) + 0)
	}
	return size
}

const (
	HostModeInsert = 0
	HostModeDelete = 1
)

// BadIDChoice is the error number for a BadIDChoice.
const BadIDChoice = 14

type IDChoiceError ValueError

// IDChoiceErrorNew constructs a IDChoiceError value that implements xgb.Error from a byte slice.
func UnmarshalIDChoiceError(buf []byte) (xgb.XError, error) {
	return UnmarshalValueError(buf)
}

// SequenceId returns the sequence id attached to the BadIDChoice error.
// This is mostly used internally.
func (err IDChoiceError) SeqID() uint16 {
	return err.Sequence
}

// BadId returns the 'BadValue' number if one exists for the BadIDChoice error. If no bad value exists, 0 is returned.
func (err IDChoiceError) BadID() uint32 {
	return err.BadValue
}

// Error returns a rudimentary string representation of the BadIDChoice error.
func (err IDChoiceError) Error() string {
	var buf strings.Builder

	buf.WriteString("BadIDChoice{")
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
	registerError(14, UnmarshalIDChoiceError)
}

const (
	ImageFormatXYBitmap = 0
	ImageFormatXYPixmap = 1
	ImageFormatZPixmap  = 2
)

const (
	ImageOrderLSBFirst = 0
	ImageOrderMSBFirst = 1
)

// BadImplementation is the error number for a BadImplementation.
const BadImplementation = 17

type ImplementationError RequestError

// ImplementationErrorNew constructs a ImplementationError value that implements xgb.Error from a byte slice.
func UnmarshalImplementationError(buf []byte) (xgb.XError, error) {
	return UnmarshalRequestError(buf)
}

// SequenceId returns the sequence id attached to the BadImplementation error.
// This is mostly used internally.
func (err ImplementationError) SeqID() uint16 {
	return err.Sequence
}

// BadId returns the 'BadValue' number if one exists for the BadImplementation error. If no bad value exists, 0 is returned.
func (err ImplementationError) BadID() uint32 {
	return err.BadValue
}

// Error returns a rudimentary string representation of the BadImplementation error.
func (err ImplementationError) Error() string {
	var buf strings.Builder

	buf.WriteString("BadImplementation{")
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
	registerError(17, UnmarshalImplementationError)
}

const (
	InputFocusNone           = 0
	InputFocusPointerRoot    = 1
	InputFocusParent         = 2
	InputFocusFollowKeyboard = 3
)

const (
	JoinStyleMiter = 0
	JoinStyleRound = 1
	JoinStyleBevel = 2
)

const (
	KbKeyClickPercent = 1
	KbBellPercent     = 2
	KbBellPitch       = 4
	KbBellDuration    = 8
	KbLed             = 16
	KbLedMode         = 32
	KbKey             = 64
	KbAutoRepeatMode  = 128
)

const (
	KeyButMaskShift   = 1
	KeyButMaskLock    = 2
	KeyButMaskControl = 4
	KeyButMaskMod1    = 8
	KeyButMaskMod2    = 16
	KeyButMaskMod3    = 32
	KeyButMaskMod4    = 64
	KeyButMaskMod5    = 128
	KeyButMaskButton1 = 256
	KeyButMaskButton2 = 512
	KeyButMaskButton3 = 1024
	KeyButMaskButton4 = 2048
	KeyButMaskButton5 = 4096
)

// KeyPress is the event number for a KeyPressEvent.
const KeyPress = 2

type KeyPressEvent struct {
	Sequence   uint16
	Detail     Keycode
	Time       Timestamp
	Root       Window
	Event      Window
	Child      Window
	RootX      int16
	RootY      int16
	EventX     int16
	EventY     int16
	State      uint16
	SameScreen bool
	// padding: 1 bytes
}

// UnmarshalKeyPressEvent constructs a KeyPressEvent value that implements xgb.Event from a byte slice.
func UnmarshalKeyPressEvent(buf []byte) (xgb.XEvent, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"KeyPressEvent\"", len(buf))
	}

	v := KeyPressEvent{}
	b := 1 // don't read event number

	v.Detail = Keycode(buf[b])
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Time = Timestamp(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Root = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Event = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Child = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.RootX = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.RootY = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.EventX = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.EventY = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.State = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.SameScreen = (buf[b] == 1)
	b += 1

	b += 1 // padding

	return v, nil
}

// Bytes writes a KeyPressEvent value to a byte slice.
func (v KeyPressEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 2
	b += 1

	buf[b] = uint8(v.Detail)
	b += 1

	b += 2 // skip sequence number

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Time))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Root))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Event))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Child))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.RootX))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.RootY))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.EventX))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.EventY))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.State)
	b += 2

	if v.SameScreen {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	b += 1 // padding

	return buf
}

// SeqID returns the sequence id attached to the KeyPress event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v KeyPressEvent) SeqID() uint16 {
	return v.Sequence
}

func init() { registerEvent(2, UnmarshalKeyPressEvent) }

// KeyRelease is the event number for a KeyReleaseEvent.
const KeyRelease = 3

type KeyReleaseEvent KeyPressEvent

// KeyReleaseEventNew constructs a KeyReleaseEvent value that implements xgb.Event from a byte slice.
func UnmarshalKeyReleaseEvent(buf []byte) (xgb.XEvent, error) {
	ev, err := UnmarshalKeyPressEvent(buf)
	return KeyReleaseEvent(ev.(KeyPressEvent)), err
}

// Bytes writes a KeyReleaseEvent value to a byte slice.
func (v KeyReleaseEvent) Bytes() []byte {
	buf := KeyPressEvent(v).Bytes()
	buf[0] = 3
	return buf
}

// SeqID returns the sequence id attached to the KeyRelease event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v KeyReleaseEvent) SeqID() uint16 {
	return v.Sequence
}

func init() {
	registerEvent(3, UnmarshalKeyReleaseEvent)
}

type Keycode byte

type Keycode32 uint32

// KeymapNotify is the event number for a KeymapNotifyEvent.
const KeymapNotify = 11

type KeymapNotifyEvent struct {
	Keys []byte // size: 32
}

// UnmarshalKeymapNotifyEvent constructs a KeymapNotifyEvent value that implements xgb.Event from a byte slice.
func UnmarshalKeymapNotifyEvent(buf []byte) (xgb.XEvent, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"KeymapNotifyEvent\"", len(buf))
	}

	v := KeymapNotifyEvent{}
	b := 1 // don't read event number

	v.Keys = make([]byte, 31)
	copy(v.Keys[:31], buf[b:])
	b += int(31)

	return v, nil
}

// Bytes writes a KeymapNotifyEvent value to a byte slice.
func (v KeymapNotifyEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 11
	b += 1

	copy(buf[b:], v.Keys[:31])
	b += int(31)

	return buf
}

// SeqID returns the sequence id attached to the KeymapNotify event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v KeymapNotifyEvent) SeqID() uint16 {
	return 0
}

func init() { registerEvent(11, UnmarshalKeymapNotifyEvent) }

type Keysym uint32

const (
	KillAllTemporary = 0
)

// LeaveNotify is the event number for a LeaveNotifyEvent.
const LeaveNotify = 8

type LeaveNotifyEvent EnterNotifyEvent

// LeaveNotifyEventNew constructs a LeaveNotifyEvent value that implements xgb.Event from a byte slice.
func UnmarshalLeaveNotifyEvent(buf []byte) (xgb.XEvent, error) {
	ev, err := UnmarshalEnterNotifyEvent(buf)
	return LeaveNotifyEvent(ev.(EnterNotifyEvent)), err
}

// Bytes writes a LeaveNotifyEvent value to a byte slice.
func (v LeaveNotifyEvent) Bytes() []byte {
	buf := EnterNotifyEvent(v).Bytes()
	buf[0] = 8
	return buf
}

// SeqID returns the sequence id attached to the LeaveNotify event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v LeaveNotifyEvent) SeqID() uint16 {
	return v.Sequence
}

func init() {
	registerEvent(8, UnmarshalLeaveNotifyEvent)
}

const (
	LedModeOff = 0
	LedModeOn  = 1
)

// BadLength is the error number for a BadLength.
const BadLength = 16

type LengthError RequestError

// LengthErrorNew constructs a LengthError value that implements xgb.Error from a byte slice.
func UnmarshalLengthError(buf []byte) (xgb.XError, error) {
	return UnmarshalRequestError(buf)
}

// SequenceId returns the sequence id attached to the BadLength error.
// This is mostly used internally.
func (err LengthError) SeqID() uint16 {
	return err.Sequence
}

// BadId returns the 'BadValue' number if one exists for the BadLength error. If no bad value exists, 0 is returned.
func (err LengthError) BadID() uint32 {
	return err.BadValue
}

// Error returns a rudimentary string representation of the BadLength error.
func (err LengthError) Error() string {
	var buf strings.Builder

	buf.WriteString("BadLength{")
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
	registerError(16, UnmarshalLengthError)
}

const (
	LineStyleSolid      = 0
	LineStyleOnOffDash  = 1
	LineStyleDoubleDash = 2
)

const (
	MapIndexShift   = 0
	MapIndexLock    = 1
	MapIndexControl = 2
	MapIndex1       = 3
	MapIndex2       = 4
	MapIndex3       = 5
	MapIndex4       = 6
	MapIndex5       = 7
)

// MapNotify is the event number for a MapNotifyEvent.
const MapNotify = 19

type MapNotifyEvent struct {
	Sequence uint16
	// padding: 1 bytes
	Event            Window
	Window           Window
	OverrideRedirect bool
	// padding: 3 bytes
}

// UnmarshalMapNotifyEvent constructs a MapNotifyEvent value that implements xgb.Event from a byte slice.
func UnmarshalMapNotifyEvent(buf []byte) (xgb.XEvent, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"MapNotifyEvent\"", len(buf))
	}

	v := MapNotifyEvent{}
	b := 1 // don't read event number

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Event = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Window = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.OverrideRedirect = (buf[b] == 1)
	b += 1

	b += 3 // padding

	return v, nil
}

// Bytes writes a MapNotifyEvent value to a byte slice.
func (v MapNotifyEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 19
	b += 1

	b += 1 // padding

	b += 2 // skip sequence number

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Event))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Window))
	b += 4

	if v.OverrideRedirect {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	b += 3 // padding

	return buf
}

// SeqID returns the sequence id attached to the MapNotify event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v MapNotifyEvent) SeqID() uint16 {
	return v.Sequence
}

func init() { registerEvent(19, UnmarshalMapNotifyEvent) }

// MapRequest is the event number for a MapRequestEvent.
const MapRequest = 20

type MapRequestEvent struct {
	Sequence uint16
	// padding: 1 bytes
	Parent Window
	Window Window
}

// UnmarshalMapRequestEvent constructs a MapRequestEvent value that implements xgb.Event from a byte slice.
func UnmarshalMapRequestEvent(buf []byte) (xgb.XEvent, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"MapRequestEvent\"", len(buf))
	}

	v := MapRequestEvent{}
	b := 1 // don't read event number

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Parent = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Window = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	return v, nil
}

// Bytes writes a MapRequestEvent value to a byte slice.
func (v MapRequestEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 20
	b += 1

	b += 1 // padding

	b += 2 // skip sequence number

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Parent))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Window))
	b += 4

	return buf
}

// SeqID returns the sequence id attached to the MapRequest event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v MapRequestEvent) SeqID() uint16 {
	return v.Sequence
}

func init() { registerEvent(20, UnmarshalMapRequestEvent) }

const (
	MapStateUnmapped   = 0
	MapStateUnviewable = 1
	MapStateViewable   = 2
)

const (
	MappingModifier = 0
	MappingKeyboard = 1
	MappingPointer  = 2
)

// MappingNotify is the event number for a MappingNotifyEvent.
const MappingNotify = 34

type MappingNotifyEvent struct {
	Sequence uint16
	// padding: 1 bytes
	Request      byte
	FirstKeycode Keycode
	Count        byte
	// padding: 1 bytes
}

// UnmarshalMappingNotifyEvent constructs a MappingNotifyEvent value that implements xgb.Event from a byte slice.
func UnmarshalMappingNotifyEvent(buf []byte) (xgb.XEvent, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"MappingNotifyEvent\"", len(buf))
	}

	v := MappingNotifyEvent{}
	b := 1 // don't read event number

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Request = buf[b]
	b += 1

	v.FirstKeycode = Keycode(buf[b])
	b += 1

	v.Count = buf[b]
	b += 1

	b += 1 // padding

	return v, nil
}

// Bytes writes a MappingNotifyEvent value to a byte slice.
func (v MappingNotifyEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 34
	b += 1

	b += 1 // padding

	b += 2 // skip sequence number

	buf[b] = v.Request
	b += 1

	buf[b] = uint8(v.FirstKeycode)
	b += 1

	buf[b] = v.Count
	b += 1

	b += 1 // padding

	return buf
}

// SeqID returns the sequence id attached to the MappingNotify event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v MappingNotifyEvent) SeqID() uint16 {
	return v.Sequence
}

func init() { registerEvent(34, UnmarshalMappingNotifyEvent) }

const (
	MappingStatusSuccess = 0
	MappingStatusBusy    = 1
	MappingStatusFailure = 2
)

// BadMatch is the error number for a BadMatch.
const BadMatch = 8

type MatchError RequestError

// MatchErrorNew constructs a MatchError value that implements xgb.Error from a byte slice.
func UnmarshalMatchError(buf []byte) (xgb.XError, error) {
	return UnmarshalRequestError(buf)
}

// SequenceId returns the sequence id attached to the BadMatch error.
// This is mostly used internally.
func (err MatchError) SeqID() uint16 {
	return err.Sequence
}

// BadId returns the 'BadValue' number if one exists for the BadMatch error. If no bad value exists, 0 is returned.
func (err MatchError) BadID() uint32 {
	return err.BadValue
}

// Error returns a rudimentary string representation of the BadMatch error.
func (err MatchError) Error() string {
	var buf strings.Builder

	buf.WriteString("BadMatch{")
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
	registerError(8, UnmarshalMatchError)
}

const (
	ModMaskShift   = 1
	ModMaskLock    = 2
	ModMaskControl = 4
	ModMask1       = 8
	ModMask2       = 16
	ModMask3       = 32
	ModMask4       = 64
	ModMask5       = 128
	ModMaskAny     = 32768
)

const (
	MotionNormal = 0
	MotionHint   = 1
)

// MotionNotify is the event number for a MotionNotifyEvent.
const MotionNotify = 6

type MotionNotifyEvent struct {
	Sequence   uint16
	Detail     byte
	Time       Timestamp
	Root       Window
	Event      Window
	Child      Window
	RootX      int16
	RootY      int16
	EventX     int16
	EventY     int16
	State      uint16
	SameScreen bool
	// padding: 1 bytes
}

// UnmarshalMotionNotifyEvent constructs a MotionNotifyEvent value that implements xgb.Event from a byte slice.
func UnmarshalMotionNotifyEvent(buf []byte) (xgb.XEvent, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"MotionNotifyEvent\"", len(buf))
	}

	v := MotionNotifyEvent{}
	b := 1 // don't read event number

	v.Detail = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Time = Timestamp(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Root = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Event = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Child = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.RootX = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.RootY = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.EventX = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.EventY = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.State = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.SameScreen = (buf[b] == 1)
	b += 1

	b += 1 // padding

	return v, nil
}

// Bytes writes a MotionNotifyEvent value to a byte slice.
func (v MotionNotifyEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 6
	b += 1

	buf[b] = v.Detail
	b += 1

	b += 2 // skip sequence number

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Time))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Root))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Event))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Child))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.RootX))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.RootY))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.EventX))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.EventY))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.State)
	b += 2

	if v.SameScreen {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	b += 1 // padding

	return buf
}

// SeqID returns the sequence id attached to the MotionNotify event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v MotionNotifyEvent) SeqID() uint16 {
	return v.Sequence
}

func init() { registerEvent(6, UnmarshalMotionNotifyEvent) }

// BadName is the error number for a BadName.
const BadName = 15

type NameError RequestError

// NameErrorNew constructs a NameError value that implements xgb.Error from a byte slice.
func UnmarshalNameError(buf []byte) (xgb.XError, error) {
	return UnmarshalRequestError(buf)
}

// SequenceId returns the sequence id attached to the BadName error.
// This is mostly used internally.
func (err NameError) SeqID() uint16 {
	return err.Sequence
}

// BadId returns the 'BadValue' number if one exists for the BadName error. If no bad value exists, 0 is returned.
func (err NameError) BadID() uint32 {
	return err.BadValue
}

// Error returns a rudimentary string representation of the BadName error.
func (err NameError) Error() string {
	var buf strings.Builder

	buf.WriteString("BadName{")
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
	registerError(15, UnmarshalNameError)
}

// NoExposure is the event number for a NoExposureEvent.
const NoExposure = 14

type NoExposureEvent struct {
	Sequence uint16
	// padding: 1 bytes
	Drawable    Drawable
	MinorOpcode uint16
	MajorOpcode byte
	// padding: 1 bytes
}

// UnmarshalNoExposureEvent constructs a NoExposureEvent value that implements xgb.Event from a byte slice.
func UnmarshalNoExposureEvent(buf []byte) (xgb.XEvent, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"NoExposureEvent\"", len(buf))
	}

	v := NoExposureEvent{}
	b := 1 // don't read event number

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Drawable = Drawable(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.MinorOpcode = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.MajorOpcode = buf[b]
	b += 1

	b += 1 // padding

	return v, nil
}

// Bytes writes a NoExposureEvent value to a byte slice.
func (v NoExposureEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 14
	b += 1

	b += 1 // padding

	b += 2 // skip sequence number

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Drawable))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], v.MinorOpcode)
	b += 2

	buf[b] = v.MajorOpcode
	b += 1

	b += 1 // padding

	return buf
}

// SeqID returns the sequence id attached to the NoExposure event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v NoExposureEvent) SeqID() uint16 {
	return v.Sequence
}

func init() { registerEvent(14, UnmarshalNoExposureEvent) }

const (
	NotifyDetailAncestor         = 0
	NotifyDetailVirtual          = 1
	NotifyDetailInferior         = 2
	NotifyDetailNonlinear        = 3
	NotifyDetailNonlinearVirtual = 4
	NotifyDetailPointer          = 5
	NotifyDetailPointerRoot      = 6
	NotifyDetailNone             = 7
)

const (
	NotifyModeNormal       = 0
	NotifyModeGrab         = 1
	NotifyModeUngrab       = 2
	NotifyModeWhileGrabbed = 3
)

type Pixmap uint32

func NewPixmapID(c *xgb.XConn) Pixmap {
	id := c.NewXID()
	return Pixmap(id)
}

// BadPixmap is the error number for a BadPixmap.
const BadPixmap = 4

type PixmapError ValueError

// PixmapErrorNew constructs a PixmapError value that implements xgb.Error from a byte slice.
func UnmarshalPixmapError(buf []byte) (xgb.XError, error) {
	return UnmarshalValueError(buf)
}

// SequenceId returns the sequence id attached to the BadPixmap error.
// This is mostly used internally.
func (err PixmapError) SeqID() uint16 {
	return err.Sequence
}

// BadId returns the 'BadValue' number if one exists for the BadPixmap error. If no bad value exists, 0 is returned.
func (err PixmapError) BadID() uint32 {
	return err.BadValue
}

// Error returns a rudimentary string representation of the BadPixmap error.
func (err PixmapError) Error() string {
	var buf strings.Builder

	buf.WriteString("BadPixmap{")
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
	registerError(4, UnmarshalPixmapError)
}

const (
	PixmapNone = 0
)

const (
	PlaceOnTop    = 0
	PlaceOnBottom = 1
)

type Point struct {
	X int16
	Y int16
}

// PointRead reads a byte slice into a Point value.
func PointRead(buf []byte, v *Point) int {
	b := 0

	v.X = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.Y = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	return b
}

// PointReadList reads a byte slice into a list of Point values.
func PointReadList(buf []byte, dest []Point) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Point{}
		b += PointRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a Point value to a byte slice.
func (v Point) Bytes() []byte {
	buf := make([]byte, 4)
	b := 0

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.X))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.Y))
	b += 2

	return buf[:b]
}

// PointListBytes writes a list of Point values to a byte slice.
func PointListBytes(buf []byte, list []Point) int {
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
	PolyShapeComplex   = 0
	PolyShapeNonconvex = 1
	PolyShapeConvex    = 2
)

const (
	PropModeReplace = 0
	PropModePrepend = 1
	PropModeAppend  = 2
)

const (
	PropertyNewValue = 0
	PropertyDelete   = 1
)

// PropertyNotify is the event number for a PropertyNotifyEvent.
const PropertyNotify = 28

type PropertyNotifyEvent struct {
	Sequence uint16
	// padding: 1 bytes
	Window Window
	Atom   Atom
	Time   Timestamp
	State  byte
	// padding: 3 bytes
}

// UnmarshalPropertyNotifyEvent constructs a PropertyNotifyEvent value that implements xgb.Event from a byte slice.
func UnmarshalPropertyNotifyEvent(buf []byte) (xgb.XEvent, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"PropertyNotifyEvent\"", len(buf))
	}

	v := PropertyNotifyEvent{}
	b := 1 // don't read event number

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Window = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Atom = Atom(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Time = Timestamp(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.State = buf[b]
	b += 1

	b += 3 // padding

	return v, nil
}

// Bytes writes a PropertyNotifyEvent value to a byte slice.
func (v PropertyNotifyEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 28
	b += 1

	b += 1 // padding

	b += 2 // skip sequence number

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Window))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Atom))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Time))
	b += 4

	buf[b] = v.State
	b += 1

	b += 3 // padding

	return buf
}

// SeqID returns the sequence id attached to the PropertyNotify event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v PropertyNotifyEvent) SeqID() uint16 {
	return v.Sequence
}

func init() { registerEvent(28, UnmarshalPropertyNotifyEvent) }

const (
	QueryShapeOfLargestCursor  = 0
	QueryShapeOfFastestTile    = 1
	QueryShapeOfFastestStipple = 2
)

type Rectangle struct {
	X      int16
	Y      int16
	Width  uint16
	Height uint16
}

// RectangleRead reads a byte slice into a Rectangle value.
func RectangleRead(buf []byte, v *Rectangle) int {
	b := 0

	v.X = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.Y = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.Width = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Height = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return b
}

// RectangleReadList reads a byte slice into a list of Rectangle values.
func RectangleReadList(buf []byte, dest []Rectangle) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Rectangle{}
		b += RectangleRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a Rectangle value to a byte slice.
func (v Rectangle) Bytes() []byte {
	buf := make([]byte, 8)
	b := 0

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.X))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.Y))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Width)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Height)
	b += 2

	return buf[:b]
}

// RectangleListBytes writes a list of Rectangle values to a byte slice.
func RectangleListBytes(buf []byte, list []Rectangle) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

// ReparentNotify is the event number for a ReparentNotifyEvent.
const ReparentNotify = 21

type ReparentNotifyEvent struct {
	Sequence uint16
	// padding: 1 bytes
	Event            Window
	Window           Window
	Parent           Window
	X                int16
	Y                int16
	OverrideRedirect bool
	// padding: 3 bytes
}

// UnmarshalReparentNotifyEvent constructs a ReparentNotifyEvent value that implements xgb.Event from a byte slice.
func UnmarshalReparentNotifyEvent(buf []byte) (xgb.XEvent, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"ReparentNotifyEvent\"", len(buf))
	}

	v := ReparentNotifyEvent{}
	b := 1 // don't read event number

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Event = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Window = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Parent = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.X = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.Y = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.OverrideRedirect = (buf[b] == 1)
	b += 1

	b += 3 // padding

	return v, nil
}

// Bytes writes a ReparentNotifyEvent value to a byte slice.
func (v ReparentNotifyEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 21
	b += 1

	b += 1 // padding

	b += 2 // skip sequence number

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Event))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Window))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Parent))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.X))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.Y))
	b += 2

	if v.OverrideRedirect {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	b += 3 // padding

	return buf
}

// SeqID returns the sequence id attached to the ReparentNotify event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v ReparentNotifyEvent) SeqID() uint16 {
	return v.Sequence
}

func init() { registerEvent(21, UnmarshalReparentNotifyEvent) }

// BadRequest is the error number for a BadRequest.
const BadRequest = 1

type RequestError struct {
	Sequence    uint16
	NiceName    string
	BadValue    uint32
	MinorOpcode uint16
	MajorOpcode byte
	// padding: 1 bytes
}

// UnmarshalRequestError constructs a RequestError value that implements xgb.Error from a byte slice.
func UnmarshalRequestError(buf []byte) (xgb.XError, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"RequestError\"", len(buf))
	}

	v := RequestError{}
	v.NiceName = "Request"

	b := 1 // skip error determinant
	b += 1 // don't read error number

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.BadValue = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.MinorOpcode = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.MajorOpcode = buf[b]
	b += 1

	b += 1 // padding

	return v, nil
}

// SeqID returns the sequence id attached to the BadRequest error.
// This is mostly used internally.
func (err RequestError) SeqID() uint16 {
	return err.Sequence
}

// BadID returns the 'BadValue' number if one exists for the BadRequest error. If no bad value exists, 0 is returned.
func (err RequestError) BadID() uint32 {
	return err.BadValue
}

// Error returns a rudimentary string representation of the BadRequest error.
func (err RequestError) Error() string {
	var buf strings.Builder

	buf.WriteString("BadRequest{")
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

func init() { registerError(1, UnmarshalRequestError) }

// ResizeRequest is the event number for a ResizeRequestEvent.
const ResizeRequest = 25

type ResizeRequestEvent struct {
	Sequence uint16
	// padding: 1 bytes
	Window Window
	Width  uint16
	Height uint16
}

// UnmarshalResizeRequestEvent constructs a ResizeRequestEvent value that implements xgb.Event from a byte slice.
func UnmarshalResizeRequestEvent(buf []byte) (xgb.XEvent, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"ResizeRequestEvent\"", len(buf))
	}

	v := ResizeRequestEvent{}
	b := 1 // don't read event number

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Window = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Width = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Height = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return v, nil
}

// Bytes writes a ResizeRequestEvent value to a byte slice.
func (v ResizeRequestEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 25
	b += 1

	b += 1 // padding

	b += 2 // skip sequence number

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Window))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], v.Width)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Height)
	b += 2

	return buf
}

// SeqID returns the sequence id attached to the ResizeRequest event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v ResizeRequestEvent) SeqID() uint16 {
	return v.Sequence
}

func init() { registerEvent(25, UnmarshalResizeRequestEvent) }

type Rgb struct {
	Red   uint16
	Green uint16
	Blue  uint16
	// padding: 2 bytes
}

// RgbRead reads a byte slice into a Rgb value.
func RgbRead(buf []byte, v *Rgb) int {
	b := 0

	v.Red = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Green = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Blue = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	b += 2 // padding

	return b
}

// RgbReadList reads a byte slice into a list of Rgb values.
func RgbReadList(buf []byte, dest []Rgb) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Rgb{}
		b += RgbRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a Rgb value to a byte slice.
func (v Rgb) Bytes() []byte {
	buf := make([]byte, 8)
	b := 0

	binary.LittleEndian.PutUint16(buf[b:], v.Red)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Green)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Blue)
	b += 2

	b += 2 // padding

	return buf[:b]
}

// RgbListBytes writes a list of Rgb values to a byte slice.
func RgbListBytes(buf []byte, list []Rgb) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

type ScreenInfo struct {
	Root                Window
	DefaultColormap     Colormap
	WhitePixel          uint32
	BlackPixel          uint32
	CurrentInputMasks   uint32
	WidthInPixels       uint16
	HeightInPixels      uint16
	WidthInMillimeters  uint16
	HeightInMillimeters uint16
	MinInstalledMaps    uint16
	MaxInstalledMaps    uint16
	RootVisual          Visualid
	BackingStores       byte
	SaveUnders          bool
	RootDepth           byte
	AllowedDepthsLen    byte
	AllowedDepths       []DepthInfo // size: DepthInfoListSize(AllowedDepths)
}

// ScreenInfoRead reads a byte slice into a ScreenInfo value.
func ScreenInfoRead(buf []byte, v *ScreenInfo) int {
	b := 0

	v.Root = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.DefaultColormap = Colormap(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.WhitePixel = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.BlackPixel = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.CurrentInputMasks = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.WidthInPixels = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.HeightInPixels = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.WidthInMillimeters = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.HeightInMillimeters = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.MinInstalledMaps = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.MaxInstalledMaps = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.RootVisual = Visualid(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.BackingStores = buf[b]
	b += 1

	v.SaveUnders = (buf[b] == 1)
	b += 1

	v.RootDepth = buf[b]
	b += 1

	v.AllowedDepthsLen = buf[b]
	b += 1

	v.AllowedDepths = make([]DepthInfo, v.AllowedDepthsLen)
	b += DepthInfoReadList(buf[b:], v.AllowedDepths)

	return b
}

// ScreenInfoReadList reads a byte slice into a list of ScreenInfo values.
func ScreenInfoReadList(buf []byte, dest []ScreenInfo) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = ScreenInfo{}
		b += ScreenInfoRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a ScreenInfo value to a byte slice.
func (v ScreenInfo) Bytes() []byte {
	buf := make([]byte, (40 + DepthInfoListSize(v.AllowedDepths)))
	b := 0

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Root))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.DefaultColormap))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], v.WhitePixel)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], v.BlackPixel)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], v.CurrentInputMasks)
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], v.WidthInPixels)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.HeightInPixels)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.WidthInMillimeters)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.HeightInMillimeters)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.MinInstalledMaps)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.MaxInstalledMaps)
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.RootVisual))
	b += 4

	buf[b] = v.BackingStores
	b += 1

	if v.SaveUnders {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	buf[b] = v.RootDepth
	b += 1

	buf[b] = v.AllowedDepthsLen
	b += 1

	b += DepthInfoListBytes(buf[b:], v.AllowedDepths)

	return buf[:b]
}

// ScreenInfoListBytes writes a list of ScreenInfo values to a byte slice.
func ScreenInfoListBytes(buf []byte, list []ScreenInfo) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

// ScreenInfoListSize computes the size (bytes) of a list of ScreenInfo values.
func ScreenInfoListSize(list []ScreenInfo) int {
	size := 0
	for _, item := range list {
		size += (40 + DepthInfoListSize(item.AllowedDepths))
	}
	return size
}

const (
	ScreenSaverReset  = 0
	ScreenSaverActive = 1
)

type Segment struct {
	X1 int16
	Y1 int16
	X2 int16
	Y2 int16
}

// SegmentRead reads a byte slice into a Segment value.
func SegmentRead(buf []byte, v *Segment) int {
	b := 0

	v.X1 = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.Y1 = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.X2 = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.Y2 = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	return b
}

// SegmentReadList reads a byte slice into a list of Segment values.
func SegmentReadList(buf []byte, dest []Segment) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Segment{}
		b += SegmentRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a Segment value to a byte slice.
func (v Segment) Bytes() []byte {
	buf := make([]byte, 8)
	b := 0

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.X1))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.Y1))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.X2))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.Y2))
	b += 2

	return buf[:b]
}

// SegmentListBytes writes a list of Segment values to a byte slice.
func SegmentListBytes(buf []byte, list []Segment) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

// SelectionClear is the event number for a SelectionClearEvent.
const SelectionClear = 29

type SelectionClearEvent struct {
	Sequence uint16
	// padding: 1 bytes
	Time      Timestamp
	Owner     Window
	Selection Atom
}

// UnmarshalSelectionClearEvent constructs a SelectionClearEvent value that implements xgb.Event from a byte slice.
func UnmarshalSelectionClearEvent(buf []byte) (xgb.XEvent, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"SelectionClearEvent\"", len(buf))
	}

	v := SelectionClearEvent{}
	b := 1 // don't read event number

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Time = Timestamp(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Owner = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Selection = Atom(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	return v, nil
}

// Bytes writes a SelectionClearEvent value to a byte slice.
func (v SelectionClearEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 29
	b += 1

	b += 1 // padding

	b += 2 // skip sequence number

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Time))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Owner))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Selection))
	b += 4

	return buf
}

// SeqID returns the sequence id attached to the SelectionClear event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v SelectionClearEvent) SeqID() uint16 {
	return v.Sequence
}

func init() { registerEvent(29, UnmarshalSelectionClearEvent) }

// SelectionNotify is the event number for a SelectionNotifyEvent.
const SelectionNotify = 31

type SelectionNotifyEvent struct {
	Sequence uint16
	// padding: 1 bytes
	Time      Timestamp
	Requestor Window
	Selection Atom
	Target    Atom
	Property  Atom
}

// UnmarshalSelectionNotifyEvent constructs a SelectionNotifyEvent value that implements xgb.Event from a byte slice.
func UnmarshalSelectionNotifyEvent(buf []byte) (xgb.XEvent, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"SelectionNotifyEvent\"", len(buf))
	}

	v := SelectionNotifyEvent{}
	b := 1 // don't read event number

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Time = Timestamp(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Requestor = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Selection = Atom(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Target = Atom(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Property = Atom(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	return v, nil
}

// Bytes writes a SelectionNotifyEvent value to a byte slice.
func (v SelectionNotifyEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 31
	b += 1

	b += 1 // padding

	b += 2 // skip sequence number

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Time))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Requestor))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Selection))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Target))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Property))
	b += 4

	return buf
}

// SeqID returns the sequence id attached to the SelectionNotify event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v SelectionNotifyEvent) SeqID() uint16 {
	return v.Sequence
}

func init() { registerEvent(31, UnmarshalSelectionNotifyEvent) }

// SelectionRequest is the event number for a SelectionRequestEvent.
const SelectionRequest = 30

type SelectionRequestEvent struct {
	Sequence uint16
	// padding: 1 bytes
	Time      Timestamp
	Owner     Window
	Requestor Window
	Selection Atom
	Target    Atom
	Property  Atom
}

// UnmarshalSelectionRequestEvent constructs a SelectionRequestEvent value that implements xgb.Event from a byte slice.
func UnmarshalSelectionRequestEvent(buf []byte) (xgb.XEvent, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"SelectionRequestEvent\"", len(buf))
	}

	v := SelectionRequestEvent{}
	b := 1 // don't read event number

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Time = Timestamp(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Owner = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Requestor = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Selection = Atom(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Target = Atom(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Property = Atom(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	return v, nil
}

// Bytes writes a SelectionRequestEvent value to a byte slice.
func (v SelectionRequestEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 30
	b += 1

	b += 1 // padding

	b += 2 // skip sequence number

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Time))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Owner))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Requestor))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Selection))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Target))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Property))
	b += 4

	return buf
}

// SeqID returns the sequence id attached to the SelectionRequest event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v SelectionRequestEvent) SeqID() uint16 {
	return v.Sequence
}

func init() { registerEvent(30, UnmarshalSelectionRequestEvent) }

const (
	SendEventDestPointerWindow = 0
	SendEventDestItemFocus     = 1
)

const (
	SetModeInsert = 0
	SetModeDelete = 1
)

type SetupAuthenticate struct {
	Status byte
	// padding: 5 bytes
	Length uint16
	Reason string // size: internal.Pad4(((int(Length) * 4) * 1))
}

// SetupAuthenticateRead reads a byte slice into a SetupAuthenticate value.
func SetupAuthenticateRead(buf []byte, v *SetupAuthenticate) int {
	b := 0

	v.Status = buf[b]
	b += 1

	b += 5 // padding

	v.Length = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	{
		byteString := make([]byte, (int(v.Length) * 4))
		copy(byteString[:(int(v.Length)*4)], buf[b:])
		v.Reason = string(byteString)
		b += int((int(v.Length) * 4))
	}

	return b
}

// SetupAuthenticateReadList reads a byte slice into a list of SetupAuthenticate values.
func SetupAuthenticateReadList(buf []byte, dest []SetupAuthenticate) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = SetupAuthenticate{}
		b += SetupAuthenticateRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a SetupAuthenticate value to a byte slice.
func (v SetupAuthenticate) Bytes() []byte {
	buf := make([]byte, (8 + internal.Pad4(((int(v.Length) * 4) * 1))))
	b := 0

	buf[b] = v.Status
	b += 1

	b += 5 // padding

	binary.LittleEndian.PutUint16(buf[b:], v.Length)
	b += 2

	copy(buf[b:], v.Reason[:(int(v.Length)*4)])
	b += int((int(v.Length) * 4))

	return buf[:b]
}

// SetupAuthenticateListBytes writes a list of SetupAuthenticate values to a byte slice.
func SetupAuthenticateListBytes(buf []byte, list []SetupAuthenticate) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

// SetupAuthenticateListSize computes the size (bytes) of a list of SetupAuthenticate values.
func SetupAuthenticateListSize(list []SetupAuthenticate) int {
	size := 0
	for _, item := range list {
		size += (8 + internal.Pad4(((int(item.Length) * 4) * 1)))
	}
	return size
}

type SetupFailed struct {
	Status               byte
	ReasonLen            byte
	ProtocolMajorVersion uint16
	ProtocolMinorVersion uint16
	Length               uint16
	Reason               string // size: internal.Pad4((int(ReasonLen) * 1))
}

// SetupFailedRead reads a byte slice into a SetupFailed value.
func SetupFailedRead(buf []byte, v *SetupFailed) int {
	b := 0

	v.Status = buf[b]
	b += 1

	v.ReasonLen = buf[b]
	b += 1

	v.ProtocolMajorVersion = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.ProtocolMinorVersion = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	{
		byteString := make([]byte, v.ReasonLen)
		copy(byteString[:v.ReasonLen], buf[b:])
		v.Reason = string(byteString)
		b += int(v.ReasonLen)
	}

	return b
}

// SetupFailedReadList reads a byte slice into a list of SetupFailed values.
func SetupFailedReadList(buf []byte, dest []SetupFailed) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = SetupFailed{}
		b += SetupFailedRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a SetupFailed value to a byte slice.
func (v SetupFailed) Bytes() []byte {
	buf := make([]byte, (8 + internal.Pad4((int(v.ReasonLen) * 1))))
	b := 0

	buf[b] = v.Status
	b += 1

	buf[b] = v.ReasonLen
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], v.ProtocolMajorVersion)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.ProtocolMinorVersion)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Length)
	b += 2

	copy(buf[b:], v.Reason[:v.ReasonLen])
	b += int(v.ReasonLen)

	return buf[:b]
}

// SetupFailedListBytes writes a list of SetupFailed values to a byte slice.
func SetupFailedListBytes(buf []byte, list []SetupFailed) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

// SetupFailedListSize computes the size (bytes) of a list of SetupFailed values.
func SetupFailedListSize(list []SetupFailed) int {
	size := 0
	for _, item := range list {
		size += (8 + internal.Pad4((int(item.ReasonLen) * 1)))
	}
	return size
}

type SetupInfo struct {
	Status byte
	// padding: 1 bytes
	ProtocolMajorVersion     uint16
	ProtocolMinorVersion     uint16
	Length                   uint16
	ReleaseNumber            uint32
	ResourceIdBase           uint32
	ResourceIdMask           uint32
	MotionBufferSize         uint32
	VendorLen                uint16
	MaximumRequestLength     uint16
	RootsLen                 byte
	PixmapFormatsLen         byte
	ImageByteOrder           byte
	BitmapFormatBitOrder     byte
	BitmapFormatScanlineUnit byte
	BitmapFormatScanlinePad  byte
	MinKeycode               Keycode
	MaxKeycode               Keycode
	// padding: 4 bytes
	Vendor string // size: internal.Pad4((int(VendorLen) * 1))
	// padding: 0 bytes
	PixmapFormats []Format     // size: internal.Pad4((int(PixmapFormatsLen) * 8))
	Roots         []ScreenInfo // size: ScreenInfoListSize(Roots)
}

// SetupInfoRead reads a byte slice into a SetupInfo value.
func SetupInfoRead(buf []byte, v *SetupInfo) int {
	b := 0

	v.Status = buf[b]
	b += 1

	b += 1 // padding

	v.ProtocolMajorVersion = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.ProtocolMinorVersion = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.ReleaseNumber = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.ResourceIdBase = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.ResourceIdMask = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.MotionBufferSize = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.VendorLen = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.MaximumRequestLength = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.RootsLen = buf[b]
	b += 1

	v.PixmapFormatsLen = buf[b]
	b += 1

	v.ImageByteOrder = buf[b]
	b += 1

	v.BitmapFormatBitOrder = buf[b]
	b += 1

	v.BitmapFormatScanlineUnit = buf[b]
	b += 1

	v.BitmapFormatScanlinePad = buf[b]
	b += 1

	v.MinKeycode = Keycode(buf[b])
	b += 1

	v.MaxKeycode = Keycode(buf[b])
	b += 1

	b += 4 // padding

	{
		byteString := make([]byte, v.VendorLen)
		copy(byteString[:v.VendorLen], buf[b:])
		v.Vendor = string(byteString)
		b += int(v.VendorLen)
	}

	b += 0 // padding

	v.PixmapFormats = make([]Format, v.PixmapFormatsLen)
	b += FormatReadList(buf[b:], v.PixmapFormats)

	v.Roots = make([]ScreenInfo, v.RootsLen)
	b += ScreenInfoReadList(buf[b:], v.Roots)

	return b
}

// SetupInfoReadList reads a byte slice into a list of SetupInfo values.
func SetupInfoReadList(buf []byte, dest []SetupInfo) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = SetupInfo{}
		b += SetupInfoRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a SetupInfo value to a byte slice.
func (v SetupInfo) Bytes() []byte {
	buf := make([]byte, ((((40 + internal.Pad4((int(v.VendorLen) * 1))) + 0) + internal.Pad4((int(v.PixmapFormatsLen) * 8))) + ScreenInfoListSize(v.Roots)))
	b := 0

	buf[b] = v.Status
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], v.ProtocolMajorVersion)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.ProtocolMinorVersion)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Length)
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], v.ReleaseNumber)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], v.ResourceIdBase)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], v.ResourceIdMask)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], v.MotionBufferSize)
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], v.VendorLen)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.MaximumRequestLength)
	b += 2

	buf[b] = v.RootsLen
	b += 1

	buf[b] = v.PixmapFormatsLen
	b += 1

	buf[b] = v.ImageByteOrder
	b += 1

	buf[b] = v.BitmapFormatBitOrder
	b += 1

	buf[b] = v.BitmapFormatScanlineUnit
	b += 1

	buf[b] = v.BitmapFormatScanlinePad
	b += 1

	buf[b] = uint8(v.MinKeycode)
	b += 1

	buf[b] = uint8(v.MaxKeycode)
	b += 1

	b += 4 // padding

	copy(buf[b:], v.Vendor[:v.VendorLen])
	b += int(v.VendorLen)

	b += 0 // padding

	b += FormatListBytes(buf[b:], v.PixmapFormats)

	b += ScreenInfoListBytes(buf[b:], v.Roots)

	return buf[:b]
}

// SetupInfoListBytes writes a list of SetupInfo values to a byte slice.
func SetupInfoListBytes(buf []byte, list []SetupInfo) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

// SetupInfoListSize computes the size (bytes) of a list of SetupInfo values.
func SetupInfoListSize(list []SetupInfo) int {
	size := 0
	for _, item := range list {
		size += ((((40 + internal.Pad4((int(item.VendorLen) * 1))) + 0) + internal.Pad4((int(item.PixmapFormatsLen) * 8))) + ScreenInfoListSize(item.Roots))
	}
	return size
}

type SetupRequest struct {
	ByteOrder byte
	// padding: 1 bytes
	ProtocolMajorVersion         uint16
	ProtocolMinorVersion         uint16
	AuthorizationProtocolNameLen uint16
	AuthorizationProtocolDataLen uint16
	// padding: 2 bytes
	AuthorizationProtocolName string // size: internal.Pad4((int(AuthorizationProtocolNameLen) * 1))
	// padding: 0 bytes
	AuthorizationProtocolData string // size: internal.Pad4((int(AuthorizationProtocolDataLen) * 1))
	// padding: 0 bytes
}

// SetupRequestRead reads a byte slice into a SetupRequest value.
func SetupRequestRead(buf []byte, v *SetupRequest) int {
	b := 0

	v.ByteOrder = buf[b]
	b += 1

	b += 1 // padding

	v.ProtocolMajorVersion = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.ProtocolMinorVersion = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.AuthorizationProtocolNameLen = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.AuthorizationProtocolDataLen = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	b += 2 // padding

	{
		byteString := make([]byte, v.AuthorizationProtocolNameLen)
		copy(byteString[:v.AuthorizationProtocolNameLen], buf[b:])
		v.AuthorizationProtocolName = string(byteString)
		b += int(v.AuthorizationProtocolNameLen)
	}

	b += 0 // padding

	{
		byteString := make([]byte, v.AuthorizationProtocolDataLen)
		copy(byteString[:v.AuthorizationProtocolDataLen], buf[b:])
		v.AuthorizationProtocolData = string(byteString)
		b += int(v.AuthorizationProtocolDataLen)
	}

	b += 0 // padding

	return b
}

// SetupRequestReadList reads a byte slice into a list of SetupRequest values.
func SetupRequestReadList(buf []byte, dest []SetupRequest) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = SetupRequest{}
		b += SetupRequestRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a SetupRequest value to a byte slice.
func (v SetupRequest) Bytes() []byte {
	buf := make([]byte, ((((12 + internal.Pad4((int(v.AuthorizationProtocolNameLen) * 1))) + 0) + internal.Pad4((int(v.AuthorizationProtocolDataLen) * 1))) + 0))
	b := 0

	buf[b] = v.ByteOrder
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], v.ProtocolMajorVersion)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.ProtocolMinorVersion)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.AuthorizationProtocolNameLen)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.AuthorizationProtocolDataLen)
	b += 2

	b += 2 // padding

	copy(buf[b:], v.AuthorizationProtocolName[:v.AuthorizationProtocolNameLen])
	b += int(v.AuthorizationProtocolNameLen)

	b += 0 // padding

	copy(buf[b:], v.AuthorizationProtocolData[:v.AuthorizationProtocolDataLen])
	b += int(v.AuthorizationProtocolDataLen)

	b += 0 // padding

	return buf[:b]
}

// SetupRequestListBytes writes a list of SetupRequest values to a byte slice.
func SetupRequestListBytes(buf []byte, list []SetupRequest) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

// SetupRequestListSize computes the size (bytes) of a list of SetupRequest values.
func SetupRequestListSize(list []SetupRequest) int {
	size := 0
	for _, item := range list {
		size += ((((12 + internal.Pad4((int(item.AuthorizationProtocolNameLen) * 1))) + 0) + internal.Pad4((int(item.AuthorizationProtocolDataLen) * 1))) + 0)
	}
	return size
}

const (
	StackModeAbove    = 0
	StackModeBelow    = 1
	StackModeTopIf    = 2
	StackModeBottomIf = 3
	StackModeOpposite = 4
)

type Str struct {
	NameLen byte
	Name    string // size: internal.Pad4((int(NameLen) * 1))
}

// StrRead reads a byte slice into a Str value.
func StrRead(buf []byte, v *Str) int {
	b := 0

	v.NameLen = buf[b]
	b += 1

	{
		byteString := make([]byte, v.NameLen)
		copy(byteString[:v.NameLen], buf[b:])
		v.Name = string(byteString)
		b += int(v.NameLen)
	}

	return b
}

// StrReadList reads a byte slice into a list of Str values.
func StrReadList(buf []byte, dest []Str) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Str{}
		b += StrRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a Str value to a byte slice.
func (v Str) Bytes() []byte {
	buf := make([]byte, (1 + internal.Pad4((int(v.NameLen) * 1))))
	b := 0

	buf[b] = v.NameLen
	b += 1

	copy(buf[b:], v.Name[:v.NameLen])
	b += int(v.NameLen)

	return buf[:b]
}

// StrListBytes writes a list of Str values to a byte slice.
func StrListBytes(buf []byte, list []Str) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

// StrListSize computes the size (bytes) of a list of Str values.
func StrListSize(list []Str) int {
	size := 0
	for _, item := range list {
		size += (1 + internal.Pad4((int(item.NameLen) * 1)))
	}
	return size
}

const (
	SubwindowModeClipByChildren   = 0
	SubwindowModeIncludeInferiors = 1
)

const (
	TimeCurrentTime = 0
)

type Timecoord struct {
	Time Timestamp
	X    int16
	Y    int16
}

// TimecoordRead reads a byte slice into a Timecoord value.
func TimecoordRead(buf []byte, v *Timecoord) int {
	b := 0

	v.Time = Timestamp(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.X = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.Y = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	return b
}

// TimecoordReadList reads a byte slice into a list of Timecoord values.
func TimecoordReadList(buf []byte, dest []Timecoord) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Timecoord{}
		b += TimecoordRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a Timecoord value to a byte slice.
func (v Timecoord) Bytes() []byte {
	buf := make([]byte, 8)
	b := 0

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Time))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.X))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.Y))
	b += 2

	return buf[:b]
}

// TimecoordListBytes writes a list of Timecoord values to a byte slice.
func TimecoordListBytes(buf []byte, list []Timecoord) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

type Timestamp uint32

// UnmapNotify is the event number for a UnmapNotifyEvent.
const UnmapNotify = 18

type UnmapNotifyEvent struct {
	Sequence uint16
	// padding: 1 bytes
	Event         Window
	Window        Window
	FromConfigure bool
	// padding: 3 bytes
}

// UnmarshalUnmapNotifyEvent constructs a UnmapNotifyEvent value that implements xgb.Event from a byte slice.
func UnmarshalUnmapNotifyEvent(buf []byte) (xgb.XEvent, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"UnmapNotifyEvent\"", len(buf))
	}

	v := UnmapNotifyEvent{}
	b := 1 // don't read event number

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Event = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Window = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.FromConfigure = (buf[b] == 1)
	b += 1

	b += 3 // padding

	return v, nil
}

// Bytes writes a UnmapNotifyEvent value to a byte slice.
func (v UnmapNotifyEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 18
	b += 1

	b += 1 // padding

	b += 2 // skip sequence number

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Event))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Window))
	b += 4

	if v.FromConfigure {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	b += 3 // padding

	return buf
}

// SeqID returns the sequence id attached to the UnmapNotify event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v UnmapNotifyEvent) SeqID() uint16 {
	return v.Sequence
}

func init() { registerEvent(18, UnmarshalUnmapNotifyEvent) }

// BadValue is the error number for a BadValue.
const BadValue = 2

type ValueError struct {
	Sequence    uint16
	NiceName    string
	BadValue    uint32
	MinorOpcode uint16
	MajorOpcode byte
	// padding: 1 bytes
}

// UnmarshalValueError constructs a ValueError value that implements xgb.Error from a byte slice.
func UnmarshalValueError(buf []byte) (xgb.XError, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"ValueError\"", len(buf))
	}

	v := ValueError{}
	v.NiceName = "Value"

	b := 1 // skip error determinant
	b += 1 // don't read error number

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.BadValue = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.MinorOpcode = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.MajorOpcode = buf[b]
	b += 1

	b += 1 // padding

	return v, nil
}

// SeqID returns the sequence id attached to the BadValue error.
// This is mostly used internally.
func (err ValueError) SeqID() uint16 {
	return err.Sequence
}

// BadID returns the 'BadValue' number if one exists for the BadValue error. If no bad value exists, 0 is returned.
func (err ValueError) BadID() uint32 {
	return err.BadValue
}

// Error returns a rudimentary string representation of the BadValue error.
func (err ValueError) Error() string {
	var buf strings.Builder

	buf.WriteString("BadValue{")
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

func init() { registerError(2, UnmarshalValueError) }

const (
	VisibilityUnobscured        = 0
	VisibilityPartiallyObscured = 1
	VisibilityFullyObscured     = 2
)

// VisibilityNotify is the event number for a VisibilityNotifyEvent.
const VisibilityNotify = 15

type VisibilityNotifyEvent struct {
	Sequence uint16
	// padding: 1 bytes
	Window Window
	State  byte
	// padding: 3 bytes
}

// UnmarshalVisibilityNotifyEvent constructs a VisibilityNotifyEvent value that implements xgb.Event from a byte slice.
func UnmarshalVisibilityNotifyEvent(buf []byte) (xgb.XEvent, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"VisibilityNotifyEvent\"", len(buf))
	}

	v := VisibilityNotifyEvent{}
	b := 1 // don't read event number

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Window = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.State = buf[b]
	b += 1

	b += 3 // padding

	return v, nil
}

// Bytes writes a VisibilityNotifyEvent value to a byte slice.
func (v VisibilityNotifyEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 15
	b += 1

	b += 1 // padding

	b += 2 // skip sequence number

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Window))
	b += 4

	buf[b] = v.State
	b += 1

	b += 3 // padding

	return buf
}

// SeqID returns the sequence id attached to the VisibilityNotify event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v VisibilityNotifyEvent) SeqID() uint16 {
	return v.Sequence
}

func init() { registerEvent(15, UnmarshalVisibilityNotifyEvent) }

const (
	VisualClassStaticGray  = 0
	VisualClassGrayScale   = 1
	VisualClassStaticColor = 2
	VisualClassPseudoColor = 3
	VisualClassTrueColor   = 4
	VisualClassDirectColor = 5
)

type VisualInfo struct {
	VisualId        Visualid
	Class           byte
	BitsPerRgbValue byte
	ColormapEntries uint16
	RedMask         uint32
	GreenMask       uint32
	BlueMask        uint32
	// padding: 4 bytes
}

// VisualInfoRead reads a byte slice into a VisualInfo value.
func VisualInfoRead(buf []byte, v *VisualInfo) int {
	b := 0

	v.VisualId = Visualid(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Class = buf[b]
	b += 1

	v.BitsPerRgbValue = buf[b]
	b += 1

	v.ColormapEntries = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.RedMask = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.GreenMask = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.BlueMask = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 4 // padding

	return b
}

// VisualInfoReadList reads a byte slice into a list of VisualInfo values.
func VisualInfoReadList(buf []byte, dest []VisualInfo) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = VisualInfo{}
		b += VisualInfoRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a VisualInfo value to a byte slice.
func (v VisualInfo) Bytes() []byte {
	buf := make([]byte, 24)
	b := 0

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.VisualId))
	b += 4

	buf[b] = v.Class
	b += 1

	buf[b] = v.BitsPerRgbValue
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], v.ColormapEntries)
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], v.RedMask)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], v.GreenMask)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], v.BlueMask)
	b += 4

	b += 4 // padding

	return buf[:b]
}

// VisualInfoListBytes writes a list of VisualInfo values to a byte slice.
func VisualInfoListBytes(buf []byte, list []VisualInfo) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

type Visualid uint32

type Window uint32

func NewWindowID(c *xgb.XConn) Window {
	id := c.NewXID()
	return Window(id)
}

// BadWindow is the error number for a BadWindow.
const BadWindow = 3

type WindowError ValueError

// WindowErrorNew constructs a WindowError value that implements xgb.Error from a byte slice.
func UnmarshalWindowError(buf []byte) (xgb.XError, error) {
	return UnmarshalValueError(buf)
}

// SequenceId returns the sequence id attached to the BadWindow error.
// This is mostly used internally.
func (err WindowError) SeqID() uint16 {
	return err.Sequence
}

// BadId returns the 'BadValue' number if one exists for the BadWindow error. If no bad value exists, 0 is returned.
func (err WindowError) BadID() uint32 {
	return err.BadValue
}

// Error returns a rudimentary string representation of the BadWindow error.
func (err WindowError) Error() string {
	var buf strings.Builder

	buf.WriteString("BadWindow{")
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
	registerError(3, UnmarshalWindowError)
}

const (
	WindowNone = 0
)

const (
	WindowClassCopyFromParent = 0
	WindowClassInputOutput    = 1
	WindowClassInputOnly      = 2
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

// AllocColor sends a checked request.
func AllocColor(c *xgb.XConn, Cmap Colormap, Red uint16, Green uint16, Blue uint16) (AllocColorReply, error) {
	var reply AllocColorReply
	var op uint8
	err := c.SendRecv(allocColorRequest(op, Cmap, Red, Green, Blue), &reply)
	return reply, err
}

// AllocColorUnchecked sends an unchecked request.
func AllocColorUnchecked(c *xgb.XConn, Cmap Colormap, Red uint16, Green uint16, Blue uint16) error {
	var op uint8
	return c.Send(allocColorRequest(op, Cmap, Red, Green, Blue))
}

// AllocColorReply represents the data returned from a AllocColor request.
type AllocColorReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Red   uint16
	Green uint16
	Blue  uint16
	// padding: 2 bytes
	Pixel uint32
}

// Unmarshal reads a byte slice into a AllocColorReply value.
func (v *AllocColorReply) Unmarshal(buf []byte) error {
	if size := 20; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"AllocColorReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Red = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Green = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Blue = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	b += 2 // padding

	v.Pixel = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	return nil
}

// Write request to wire for AllocColor
// allocColorRequest writes a AllocColor request to a byte slice.
func allocColorRequest(opcode uint8, Cmap Colormap, Red uint16, Green uint16, Blue uint16) []byte {
	size := 16
	b := 0
	buf := make([]byte, size)

	buf[b] = 84 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Cmap))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], Red)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Green)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Blue)
	b += 2

	b += 2 // padding

	return buf
}

// AllocColorCells sends a checked request.
func AllocColorCells(c *xgb.XConn, Contiguous bool, Cmap Colormap, Colors uint16, Planes uint16) (AllocColorCellsReply, error) {
	var reply AllocColorCellsReply
	var op uint8
	err := c.SendRecv(allocColorCellsRequest(op, Contiguous, Cmap, Colors, Planes), &reply)
	return reply, err
}

// AllocColorCellsUnchecked sends an unchecked request.
func AllocColorCellsUnchecked(c *xgb.XConn, Contiguous bool, Cmap Colormap, Colors uint16, Planes uint16) error {
	var op uint8
	return c.Send(allocColorCellsRequest(op, Contiguous, Cmap, Colors, Planes))
}

// AllocColorCellsReply represents the data returned from a AllocColorCells request.
type AllocColorCellsReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	PixelsLen uint16
	MasksLen  uint16
	// padding: 20 bytes
	Pixels []uint32 // size: internal.Pad4((int(PixelsLen) * 4))
	// alignment gap to multiple of 4
	Masks []uint32 // size: internal.Pad4((int(MasksLen) * 4))
}

// Unmarshal reads a byte slice into a AllocColorCellsReply value.
func (v *AllocColorCellsReply) Unmarshal(buf []byte) error {
	if size := (((32 + internal.Pad4((int(v.PixelsLen) * 4))) + 4) + internal.Pad4((int(v.MasksLen) * 4))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"AllocColorCellsReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.PixelsLen = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.MasksLen = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	b += 20 // padding

	v.Pixels = make([]uint32, v.PixelsLen)
	for i := 0; i < int(v.PixelsLen); i++ {
		v.Pixels[i] = binary.LittleEndian.Uint32(buf[b:])
		b += 4
	}

	b = (b + 3) & ^3 // alignment gap

	v.Masks = make([]uint32, v.MasksLen)
	for i := 0; i < int(v.MasksLen); i++ {
		v.Masks[i] = binary.LittleEndian.Uint32(buf[b:])
		b += 4
	}

	return nil
}

// Write request to wire for AllocColorCells
// allocColorCellsRequest writes a AllocColorCells request to a byte slice.
func allocColorCellsRequest(opcode uint8, Contiguous bool, Cmap Colormap, Colors uint16, Planes uint16) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = 86 // request opcode
	b += 1

	if Contiguous {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Cmap))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], Colors)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Planes)
	b += 2

	return buf
}

// AllocColorPlanes sends a checked request.
func AllocColorPlanes(c *xgb.XConn, Contiguous bool, Cmap Colormap, Colors uint16, Reds uint16, Greens uint16, Blues uint16) (AllocColorPlanesReply, error) {
	var reply AllocColorPlanesReply
	var op uint8
	err := c.SendRecv(allocColorPlanesRequest(op, Contiguous, Cmap, Colors, Reds, Greens, Blues), &reply)
	return reply, err
}

// AllocColorPlanesUnchecked sends an unchecked request.
func AllocColorPlanesUnchecked(c *xgb.XConn, Contiguous bool, Cmap Colormap, Colors uint16, Reds uint16, Greens uint16, Blues uint16) error {
	var op uint8
	return c.Send(allocColorPlanesRequest(op, Contiguous, Cmap, Colors, Reds, Greens, Blues))
}

// AllocColorPlanesReply represents the data returned from a AllocColorPlanes request.
type AllocColorPlanesReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	PixelsLen uint16
	// padding: 2 bytes
	RedMask   uint32
	GreenMask uint32
	BlueMask  uint32
	// padding: 8 bytes
	Pixels []uint32 // size: internal.Pad4((int(PixelsLen) * 4))
}

// Unmarshal reads a byte slice into a AllocColorPlanesReply value.
func (v *AllocColorPlanesReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.PixelsLen) * 4))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"AllocColorPlanesReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.PixelsLen = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	b += 2 // padding

	v.RedMask = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.GreenMask = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.BlueMask = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 8 // padding

	v.Pixels = make([]uint32, v.PixelsLen)
	for i := 0; i < int(v.PixelsLen); i++ {
		v.Pixels[i] = binary.LittleEndian.Uint32(buf[b:])
		b += 4
	}

	return nil
}

// Write request to wire for AllocColorPlanes
// allocColorPlanesRequest writes a AllocColorPlanes request to a byte slice.
func allocColorPlanesRequest(opcode uint8, Contiguous bool, Cmap Colormap, Colors uint16, Reds uint16, Greens uint16, Blues uint16) []byte {
	size := 16
	b := 0
	buf := make([]byte, size)

	buf[b] = 87 // request opcode
	b += 1

	if Contiguous {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Cmap))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], Colors)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Reds)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Greens)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Blues)
	b += 2

	return buf
}

// AllocNamedColor sends a checked request.
func AllocNamedColor(c *xgb.XConn, Cmap Colormap, NameLen uint16, Name string) (AllocNamedColorReply, error) {
	var reply AllocNamedColorReply
	var op uint8
	err := c.SendRecv(allocNamedColorRequest(op, Cmap, NameLen, Name), &reply)
	return reply, err
}

// AllocNamedColorUnchecked sends an unchecked request.
func AllocNamedColorUnchecked(c *xgb.XConn, Cmap Colormap, NameLen uint16, Name string) error {
	var op uint8
	return c.Send(allocNamedColorRequest(op, Cmap, NameLen, Name))
}

// AllocNamedColorReply represents the data returned from a AllocNamedColor request.
type AllocNamedColorReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Pixel       uint32
	ExactRed    uint16
	ExactGreen  uint16
	ExactBlue   uint16
	VisualRed   uint16
	VisualGreen uint16
	VisualBlue  uint16
}

// Unmarshal reads a byte slice into a AllocNamedColorReply value.
func (v *AllocNamedColorReply) Unmarshal(buf []byte) error {
	if size := 24; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"AllocNamedColorReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Pixel = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.ExactRed = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.ExactGreen = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.ExactBlue = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.VisualRed = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.VisualGreen = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.VisualBlue = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return nil
}

// Write request to wire for AllocNamedColor
// allocNamedColorRequest writes a AllocNamedColor request to a byte slice.
func allocNamedColorRequest(opcode uint8, Cmap Colormap, NameLen uint16, Name string) []byte {
	size := internal.Pad4((12 + internal.Pad4((int(NameLen) * 1))))
	b := 0
	buf := make([]byte, size)

	buf[b] = 85 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Cmap))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], NameLen)
	b += 2

	b += 2 // padding

	copy(buf[b:], Name[:NameLen])
	b += int(NameLen)

	return buf
}

// AllowEvents sends a checked request.
func AllowEvents(c *xgb.XConn, Mode byte, Time Timestamp) error {
	var op uint8
	return c.SendRecv(allowEventsRequest(op, Mode, Time), nil)
}

// AllowEventsUnchecked sends an unchecked request.
func AllowEventsUnchecked(c *xgb.XConn, Mode byte, Time Timestamp) error {
	var op uint8
	return c.Send(allowEventsRequest(op, Mode, Time))
}

// Write request to wire for AllowEvents
// allowEventsRequest writes a AllowEvents request to a byte slice.
func allowEventsRequest(opcode uint8, Mode byte, Time Timestamp) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = 35 // request opcode
	b += 1

	buf[b] = Mode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Time))
	b += 4

	return buf
}

// Bell sends a checked request.
func Bell(c *xgb.XConn, Percent int8) error {
	var op uint8
	return c.SendRecv(bellRequest(op, Percent), nil)
}

// BellUnchecked sends an unchecked request.
func BellUnchecked(c *xgb.XConn, Percent int8) error {
	var op uint8
	return c.Send(bellRequest(op, Percent))
}

// Write request to wire for Bell
// bellRequest writes a Bell request to a byte slice.
func bellRequest(opcode uint8, Percent int8) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = 104 // request opcode
	b += 1

	buf[b] = uint8(Percent)
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// ChangeActivePointerGrab sends a checked request.
func ChangeActivePointerGrab(c *xgb.XConn, Cursor Cursor, Time Timestamp, EventMask uint16) error {
	var op uint8
	return c.SendRecv(changeActivePointerGrabRequest(op, Cursor, Time, EventMask), nil)
}

// ChangeActivePointerGrabUnchecked sends an unchecked request.
func ChangeActivePointerGrabUnchecked(c *xgb.XConn, Cursor Cursor, Time Timestamp, EventMask uint16) error {
	var op uint8
	return c.Send(changeActivePointerGrabRequest(op, Cursor, Time, EventMask))
}

// Write request to wire for ChangeActivePointerGrab
// changeActivePointerGrabRequest writes a ChangeActivePointerGrab request to a byte slice.
func changeActivePointerGrabRequest(opcode uint8, Cursor Cursor, Time Timestamp, EventMask uint16) []byte {
	size := 16
	b := 0
	buf := make([]byte, size)

	buf[b] = 30 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Cursor))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Time))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], EventMask)
	b += 2

	b += 2 // padding

	return buf
}

// ChangeGC sends a checked request.
func ChangeGC(c *xgb.XConn, Gc Gcontext, ValueMask uint32, ValueList []uint32) error {
	var op uint8
	return c.SendRecv(changeGCRequest(op, Gc, ValueMask, ValueList), nil)
}

// ChangeGCUnchecked sends an unchecked request.
func ChangeGCUnchecked(c *xgb.XConn, Gc Gcontext, ValueMask uint32, ValueList []uint32) error {
	var op uint8
	return c.Send(changeGCRequest(op, Gc, ValueMask, ValueList))
}

// Write request to wire for ChangeGC
// changeGCRequest writes a ChangeGC request to a byte slice.
func changeGCRequest(opcode uint8, Gc Gcontext, ValueMask uint32, ValueList []uint32) []byte {
	size := internal.Pad4((12 + internal.Pad4((4 * internal.PopCount(int(ValueMask))))))
	b := 0
	buf := make([]byte, size)

	buf[b] = 56 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Gc))
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

// ChangeHosts sends a checked request.
func ChangeHosts(c *xgb.XConn, Mode byte, Family byte, AddressLen uint16, Address []byte) error {
	var op uint8
	return c.SendRecv(changeHostsRequest(op, Mode, Family, AddressLen, Address), nil)
}

// ChangeHostsUnchecked sends an unchecked request.
func ChangeHostsUnchecked(c *xgb.XConn, Mode byte, Family byte, AddressLen uint16, Address []byte) error {
	var op uint8
	return c.Send(changeHostsRequest(op, Mode, Family, AddressLen, Address))
}

// Write request to wire for ChangeHosts
// changeHostsRequest writes a ChangeHosts request to a byte slice.
func changeHostsRequest(opcode uint8, Mode byte, Family byte, AddressLen uint16, Address []byte) []byte {
	size := internal.Pad4((8 + internal.Pad4((int(AddressLen) * 1))))
	b := 0
	buf := make([]byte, size)

	buf[b] = 109 // request opcode
	b += 1

	buf[b] = Mode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	buf[b] = Family
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], AddressLen)
	b += 2

	copy(buf[b:], Address[:AddressLen])
	b += int(AddressLen)

	return buf
}

// ChangeKeyboardControl sends a checked request.
func ChangeKeyboardControl(c *xgb.XConn, ValueMask uint32, ValueList []uint32) error {
	var op uint8
	return c.SendRecv(changeKeyboardControlRequest(op, ValueMask, ValueList), nil)
}

// ChangeKeyboardControlUnchecked sends an unchecked request.
func ChangeKeyboardControlUnchecked(c *xgb.XConn, ValueMask uint32, ValueList []uint32) error {
	var op uint8
	return c.Send(changeKeyboardControlRequest(op, ValueMask, ValueList))
}

// Write request to wire for ChangeKeyboardControl
// changeKeyboardControlRequest writes a ChangeKeyboardControl request to a byte slice.
func changeKeyboardControlRequest(opcode uint8, ValueMask uint32, ValueList []uint32) []byte {
	size := internal.Pad4((8 + internal.Pad4((4 * internal.PopCount(int(ValueMask))))))
	b := 0
	buf := make([]byte, size)

	buf[b] = 102 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], ValueMask)
	b += 4

	for i := 0; i < len(ValueList); i++ {
		binary.LittleEndian.PutUint32(buf[b:], ValueList[i])
		b += 4
	}
	b = internal.Pad4(b)

	return buf
}

// ChangeKeyboardMapping sends a checked request.
func ChangeKeyboardMapping(c *xgb.XConn, KeycodeCount byte, FirstKeycode Keycode, KeysymsPerKeycode byte, Keysyms []Keysym) error {
	var op uint8
	return c.SendRecv(changeKeyboardMappingRequest(op, KeycodeCount, FirstKeycode, KeysymsPerKeycode, Keysyms), nil)
}

// ChangeKeyboardMappingUnchecked sends an unchecked request.
func ChangeKeyboardMappingUnchecked(c *xgb.XConn, KeycodeCount byte, FirstKeycode Keycode, KeysymsPerKeycode byte, Keysyms []Keysym) error {
	var op uint8
	return c.Send(changeKeyboardMappingRequest(op, KeycodeCount, FirstKeycode, KeysymsPerKeycode, Keysyms))
}

// Write request to wire for ChangeKeyboardMapping
// changeKeyboardMappingRequest writes a ChangeKeyboardMapping request to a byte slice.
func changeKeyboardMappingRequest(opcode uint8, KeycodeCount byte, FirstKeycode Keycode, KeysymsPerKeycode byte, Keysyms []Keysym) []byte {
	size := internal.Pad4((8 + internal.Pad4(((int(KeycodeCount) * int(KeysymsPerKeycode)) * 4))))
	b := 0
	buf := make([]byte, size)

	buf[b] = 100 // request opcode
	b += 1

	buf[b] = KeycodeCount
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	buf[b] = uint8(FirstKeycode)
	b += 1

	buf[b] = KeysymsPerKeycode
	b += 1

	b += 2 // padding

	for i := 0; i < int((int(KeycodeCount) * int(KeysymsPerKeycode))); i++ {
		binary.LittleEndian.PutUint32(buf[b:], uint32(Keysyms[i]))
		b += 4
	}

	return buf
}

// ChangePointerControl sends a checked request.
func ChangePointerControl(c *xgb.XConn, AccelerationNumerator int16, AccelerationDenominator int16, Threshold int16, DoAcceleration bool, DoThreshold bool) error {
	var op uint8
	return c.SendRecv(changePointerControlRequest(op, AccelerationNumerator, AccelerationDenominator, Threshold, DoAcceleration, DoThreshold), nil)
}

// ChangePointerControlUnchecked sends an unchecked request.
func ChangePointerControlUnchecked(c *xgb.XConn, AccelerationNumerator int16, AccelerationDenominator int16, Threshold int16, DoAcceleration bool, DoThreshold bool) error {
	var op uint8
	return c.Send(changePointerControlRequest(op, AccelerationNumerator, AccelerationDenominator, Threshold, DoAcceleration, DoThreshold))
}

// Write request to wire for ChangePointerControl
// changePointerControlRequest writes a ChangePointerControl request to a byte slice.
func changePointerControlRequest(opcode uint8, AccelerationNumerator int16, AccelerationDenominator int16, Threshold int16, DoAcceleration bool, DoThreshold bool) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = 105 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(AccelerationNumerator))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(AccelerationDenominator))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(Threshold))
	b += 2

	if DoAcceleration {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	if DoThreshold {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	return buf
}

// ChangeProperty sends a checked request.
func ChangeProperty(c *xgb.XConn, Mode byte, Window Window, Property Atom, Type Atom, Format byte, DataLen uint32, Data []byte) error {
	var op uint8
	return c.SendRecv(changePropertyRequest(op, Mode, Window, Property, Type, Format, DataLen, Data), nil)
}

// ChangePropertyUnchecked sends an unchecked request.
func ChangePropertyUnchecked(c *xgb.XConn, Mode byte, Window Window, Property Atom, Type Atom, Format byte, DataLen uint32, Data []byte) error {
	var op uint8
	return c.Send(changePropertyRequest(op, Mode, Window, Property, Type, Format, DataLen, Data))
}

// Write request to wire for ChangeProperty
// changePropertyRequest writes a ChangeProperty request to a byte slice.
func changePropertyRequest(opcode uint8, Mode byte, Window Window, Property Atom, Type Atom, Format byte, DataLen uint32, Data []byte) []byte {
	size := internal.Pad4((24 + internal.Pad4((((int(DataLen) * int(Format)) / 8) * 1))))
	b := 0
	buf := make([]byte, size)

	buf[b] = 18 // request opcode
	b += 1

	buf[b] = Mode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Property))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Type))
	b += 4

	buf[b] = Format
	b += 1

	b += 3 // padding

	binary.LittleEndian.PutUint32(buf[b:], DataLen)
	b += 4

	copy(buf[b:], Data[:((int(DataLen)*int(Format))/8)])
	b += int(((int(DataLen) * int(Format)) / 8))

	return buf
}

// ChangeSaveSet sends a checked request.
func ChangeSaveSet(c *xgb.XConn, Mode byte, Window Window) error {
	var op uint8
	return c.SendRecv(changeSaveSetRequest(op, Mode, Window), nil)
}

// ChangeSaveSetUnchecked sends an unchecked request.
func ChangeSaveSetUnchecked(c *xgb.XConn, Mode byte, Window Window) error {
	var op uint8
	return c.Send(changeSaveSetRequest(op, Mode, Window))
}

// Write request to wire for ChangeSaveSet
// changeSaveSetRequest writes a ChangeSaveSet request to a byte slice.
func changeSaveSetRequest(opcode uint8, Mode byte, Window Window) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = 6 // request opcode
	b += 1

	buf[b] = Mode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	return buf
}

// ChangeWindowAttributes sends a checked request.
func ChangeWindowAttributes(c *xgb.XConn, Window Window, ValueMask uint32, ValueList []uint32) error {
	var op uint8
	return c.SendRecv(changeWindowAttributesRequest(op, Window, ValueMask, ValueList), nil)
}

// ChangeWindowAttributesUnchecked sends an unchecked request.
func ChangeWindowAttributesUnchecked(c *xgb.XConn, Window Window, ValueMask uint32, ValueList []uint32) error {
	var op uint8
	return c.Send(changeWindowAttributesRequest(op, Window, ValueMask, ValueList))
}

// Write request to wire for ChangeWindowAttributes
// changeWindowAttributesRequest writes a ChangeWindowAttributes request to a byte slice.
func changeWindowAttributesRequest(opcode uint8, Window Window, ValueMask uint32, ValueList []uint32) []byte {
	size := internal.Pad4((12 + internal.Pad4((4 * internal.PopCount(int(ValueMask))))))
	b := 0
	buf := make([]byte, size)

	buf[b] = 2 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
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

// CirculateWindow sends a checked request.
func CirculateWindow(c *xgb.XConn, Direction byte, Window Window) error {
	var op uint8
	return c.SendRecv(circulateWindowRequest(op, Direction, Window), nil)
}

// CirculateWindowUnchecked sends an unchecked request.
func CirculateWindowUnchecked(c *xgb.XConn, Direction byte, Window Window) error {
	var op uint8
	return c.Send(circulateWindowRequest(op, Direction, Window))
}

// Write request to wire for CirculateWindow
// circulateWindowRequest writes a CirculateWindow request to a byte slice.
func circulateWindowRequest(opcode uint8, Direction byte, Window Window) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = 13 // request opcode
	b += 1

	buf[b] = Direction
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	return buf
}

// ClearArea sends a checked request.
func ClearArea(c *xgb.XConn, Exposures bool, Window Window, X int16, Y int16, Width uint16, Height uint16) error {
	var op uint8
	return c.SendRecv(clearAreaRequest(op, Exposures, Window, X, Y, Width, Height), nil)
}

// ClearAreaUnchecked sends an unchecked request.
func ClearAreaUnchecked(c *xgb.XConn, Exposures bool, Window Window, X int16, Y int16, Width uint16, Height uint16) error {
	var op uint8
	return c.Send(clearAreaRequest(op, Exposures, Window, X, Y, Width, Height))
}

// Write request to wire for ClearArea
// clearAreaRequest writes a ClearArea request to a byte slice.
func clearAreaRequest(opcode uint8, Exposures bool, Window Window, X int16, Y int16, Width uint16, Height uint16) []byte {
	size := 16
	b := 0
	buf := make([]byte, size)

	buf[b] = 61 // request opcode
	b += 1

	if Exposures {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(X))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(Y))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Width)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Height)
	b += 2

	return buf
}

// CloseFont sends a checked request.
func CloseFont(c *xgb.XConn, Font Font) error {
	var op uint8
	return c.SendRecv(closeFontRequest(op, Font), nil)
}

// CloseFontUnchecked sends an unchecked request.
func CloseFontUnchecked(c *xgb.XConn, Font Font) error {
	var op uint8
	return c.Send(closeFontRequest(op, Font))
}

// Write request to wire for CloseFont
// closeFontRequest writes a CloseFont request to a byte slice.
func closeFontRequest(opcode uint8, Font Font) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = 46 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Font))
	b += 4

	return buf
}

// ConfigureWindow sends a checked request.
func ConfigureWindow(c *xgb.XConn, Window Window, ValueMask uint16, ValueList []uint32) error {
	var op uint8
	return c.SendRecv(configureWindowRequest(op, Window, ValueMask, ValueList), nil)
}

// ConfigureWindowUnchecked sends an unchecked request.
func ConfigureWindowUnchecked(c *xgb.XConn, Window Window, ValueMask uint16, ValueList []uint32) error {
	var op uint8
	return c.Send(configureWindowRequest(op, Window, ValueMask, ValueList))
}

// Write request to wire for ConfigureWindow
// configureWindowRequest writes a ConfigureWindow request to a byte slice.
func configureWindowRequest(opcode uint8, Window Window, ValueMask uint16, ValueList []uint32) []byte {
	size := internal.Pad4((12 + internal.Pad4((4 * internal.PopCount(int(ValueMask))))))
	b := 0
	buf := make([]byte, size)

	buf[b] = 12 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], ValueMask)
	b += 2

	b += 2 // padding

	for i := 0; i < len(ValueList); i++ {
		binary.LittleEndian.PutUint32(buf[b:], ValueList[i])
		b += 4
	}
	b = internal.Pad4(b)

	return buf
}

// ConvertSelection sends a checked request.
func ConvertSelection(c *xgb.XConn, Requestor Window, Selection Atom, Target Atom, Property Atom, Time Timestamp) error {
	var op uint8
	return c.SendRecv(convertSelectionRequest(op, Requestor, Selection, Target, Property, Time), nil)
}

// ConvertSelectionUnchecked sends an unchecked request.
func ConvertSelectionUnchecked(c *xgb.XConn, Requestor Window, Selection Atom, Target Atom, Property Atom, Time Timestamp) error {
	var op uint8
	return c.Send(convertSelectionRequest(op, Requestor, Selection, Target, Property, Time))
}

// Write request to wire for ConvertSelection
// convertSelectionRequest writes a ConvertSelection request to a byte slice.
func convertSelectionRequest(opcode uint8, Requestor Window, Selection Atom, Target Atom, Property Atom, Time Timestamp) []byte {
	size := 24
	b := 0
	buf := make([]byte, size)

	buf[b] = 24 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Requestor))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Selection))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Target))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Property))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Time))
	b += 4

	return buf
}

// CopyArea sends a checked request.
func CopyArea(c *xgb.XConn, SrcDrawable Drawable, DstDrawable Drawable, Gc Gcontext, SrcX int16, SrcY int16, DstX int16, DstY int16, Width uint16, Height uint16) error {
	var op uint8
	return c.SendRecv(copyAreaRequest(op, SrcDrawable, DstDrawable, Gc, SrcX, SrcY, DstX, DstY, Width, Height), nil)
}

// CopyAreaUnchecked sends an unchecked request.
func CopyAreaUnchecked(c *xgb.XConn, SrcDrawable Drawable, DstDrawable Drawable, Gc Gcontext, SrcX int16, SrcY int16, DstX int16, DstY int16, Width uint16, Height uint16) error {
	var op uint8
	return c.Send(copyAreaRequest(op, SrcDrawable, DstDrawable, Gc, SrcX, SrcY, DstX, DstY, Width, Height))
}

// Write request to wire for CopyArea
// copyAreaRequest writes a CopyArea request to a byte slice.
func copyAreaRequest(opcode uint8, SrcDrawable Drawable, DstDrawable Drawable, Gc Gcontext, SrcX int16, SrcY int16, DstX int16, DstY int16, Width uint16, Height uint16) []byte {
	size := 28
	b := 0
	buf := make([]byte, size)

	buf[b] = 62 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(SrcDrawable))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(DstDrawable))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Gc))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(SrcX))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(SrcY))
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

// CopyColormapAndFree sends a checked request.
func CopyColormapAndFree(c *xgb.XConn, Mid Colormap, SrcCmap Colormap) error {
	var op uint8
	return c.SendRecv(copyColormapAndFreeRequest(op, Mid, SrcCmap), nil)
}

// CopyColormapAndFreeUnchecked sends an unchecked request.
func CopyColormapAndFreeUnchecked(c *xgb.XConn, Mid Colormap, SrcCmap Colormap) error {
	var op uint8
	return c.Send(copyColormapAndFreeRequest(op, Mid, SrcCmap))
}

// Write request to wire for CopyColormapAndFree
// copyColormapAndFreeRequest writes a CopyColormapAndFree request to a byte slice.
func copyColormapAndFreeRequest(opcode uint8, Mid Colormap, SrcCmap Colormap) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = 80 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Mid))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(SrcCmap))
	b += 4

	return buf
}

// CopyGC sends a checked request.
func CopyGC(c *xgb.XConn, SrcGc Gcontext, DstGc Gcontext, ValueMask uint32) error {
	var op uint8
	return c.SendRecv(copyGCRequest(op, SrcGc, DstGc, ValueMask), nil)
}

// CopyGCUnchecked sends an unchecked request.
func CopyGCUnchecked(c *xgb.XConn, SrcGc Gcontext, DstGc Gcontext, ValueMask uint32) error {
	var op uint8
	return c.Send(copyGCRequest(op, SrcGc, DstGc, ValueMask))
}

// Write request to wire for CopyGC
// copyGCRequest writes a CopyGC request to a byte slice.
func copyGCRequest(opcode uint8, SrcGc Gcontext, DstGc Gcontext, ValueMask uint32) []byte {
	size := 16
	b := 0
	buf := make([]byte, size)

	buf[b] = 57 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(SrcGc))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(DstGc))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], ValueMask)
	b += 4

	return buf
}

// CopyPlane sends a checked request.
func CopyPlane(c *xgb.XConn, SrcDrawable Drawable, DstDrawable Drawable, Gc Gcontext, SrcX int16, SrcY int16, DstX int16, DstY int16, Width uint16, Height uint16, BitPlane uint32) error {
	var op uint8
	return c.SendRecv(copyPlaneRequest(op, SrcDrawable, DstDrawable, Gc, SrcX, SrcY, DstX, DstY, Width, Height, BitPlane), nil)
}

// CopyPlaneUnchecked sends an unchecked request.
func CopyPlaneUnchecked(c *xgb.XConn, SrcDrawable Drawable, DstDrawable Drawable, Gc Gcontext, SrcX int16, SrcY int16, DstX int16, DstY int16, Width uint16, Height uint16, BitPlane uint32) error {
	var op uint8
	return c.Send(copyPlaneRequest(op, SrcDrawable, DstDrawable, Gc, SrcX, SrcY, DstX, DstY, Width, Height, BitPlane))
}

// Write request to wire for CopyPlane
// copyPlaneRequest writes a CopyPlane request to a byte slice.
func copyPlaneRequest(opcode uint8, SrcDrawable Drawable, DstDrawable Drawable, Gc Gcontext, SrcX int16, SrcY int16, DstX int16, DstY int16, Width uint16, Height uint16, BitPlane uint32) []byte {
	size := 32
	b := 0
	buf := make([]byte, size)

	buf[b] = 63 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(SrcDrawable))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(DstDrawable))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Gc))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(SrcX))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(SrcY))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(DstX))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(DstY))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Width)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Height)
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], BitPlane)
	b += 4

	return buf
}

// CreateColormap sends a checked request.
func CreateColormap(c *xgb.XConn, Alloc byte, Mid Colormap, Window Window, Visual Visualid) error {
	var op uint8
	return c.SendRecv(createColormapRequest(op, Alloc, Mid, Window, Visual), nil)
}

// CreateColormapUnchecked sends an unchecked request.
func CreateColormapUnchecked(c *xgb.XConn, Alloc byte, Mid Colormap, Window Window, Visual Visualid) error {
	var op uint8
	return c.Send(createColormapRequest(op, Alloc, Mid, Window, Visual))
}

// Write request to wire for CreateColormap
// createColormapRequest writes a CreateColormap request to a byte slice.
func createColormapRequest(opcode uint8, Alloc byte, Mid Colormap, Window Window, Visual Visualid) []byte {
	size := 16
	b := 0
	buf := make([]byte, size)

	buf[b] = 78 // request opcode
	b += 1

	buf[b] = Alloc
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Mid))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Visual))
	b += 4

	return buf
}

// CreateCursor sends a checked request.
func CreateCursor(c *xgb.XConn, Cid Cursor, Source Pixmap, Mask Pixmap, ForeRed uint16, ForeGreen uint16, ForeBlue uint16, BackRed uint16, BackGreen uint16, BackBlue uint16, X uint16, Y uint16) error {
	var op uint8
	return c.SendRecv(createCursorRequest(op, Cid, Source, Mask, ForeRed, ForeGreen, ForeBlue, BackRed, BackGreen, BackBlue, X, Y), nil)
}

// CreateCursorUnchecked sends an unchecked request.
func CreateCursorUnchecked(c *xgb.XConn, Cid Cursor, Source Pixmap, Mask Pixmap, ForeRed uint16, ForeGreen uint16, ForeBlue uint16, BackRed uint16, BackGreen uint16, BackBlue uint16, X uint16, Y uint16) error {
	var op uint8
	return c.Send(createCursorRequest(op, Cid, Source, Mask, ForeRed, ForeGreen, ForeBlue, BackRed, BackGreen, BackBlue, X, Y))
}

// Write request to wire for CreateCursor
// createCursorRequest writes a CreateCursor request to a byte slice.
func createCursorRequest(opcode uint8, Cid Cursor, Source Pixmap, Mask Pixmap, ForeRed uint16, ForeGreen uint16, ForeBlue uint16, BackRed uint16, BackGreen uint16, BackBlue uint16, X uint16, Y uint16) []byte {
	size := 32
	b := 0
	buf := make([]byte, size)

	buf[b] = 93 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Cid))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Source))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Mask))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], ForeRed)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], ForeGreen)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], ForeBlue)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], BackRed)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], BackGreen)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], BackBlue)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], X)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Y)
	b += 2

	return buf
}

// CreateGC sends a checked request.
func CreateGC(c *xgb.XConn, Cid Gcontext, Drawable Drawable, ValueMask uint32, ValueList []uint32) error {
	var op uint8
	return c.SendRecv(createGCRequest(op, Cid, Drawable, ValueMask, ValueList), nil)
}

// CreateGCUnchecked sends an unchecked request.
func CreateGCUnchecked(c *xgb.XConn, Cid Gcontext, Drawable Drawable, ValueMask uint32, ValueList []uint32) error {
	var op uint8
	return c.Send(createGCRequest(op, Cid, Drawable, ValueMask, ValueList))
}

// Write request to wire for CreateGC
// createGCRequest writes a CreateGC request to a byte slice.
func createGCRequest(opcode uint8, Cid Gcontext, Drawable Drawable, ValueMask uint32, ValueList []uint32) []byte {
	size := internal.Pad4((16 + internal.Pad4((4 * internal.PopCount(int(ValueMask))))))
	b := 0
	buf := make([]byte, size)

	buf[b] = 55 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Cid))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Drawable))
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

// CreateGlyphCursor sends a checked request.
func CreateGlyphCursor(c *xgb.XConn, Cid Cursor, SourceFont Font, MaskFont Font, SourceChar uint16, MaskChar uint16, ForeRed uint16, ForeGreen uint16, ForeBlue uint16, BackRed uint16, BackGreen uint16, BackBlue uint16) error {
	var op uint8
	return c.SendRecv(createGlyphCursorRequest(op, Cid, SourceFont, MaskFont, SourceChar, MaskChar, ForeRed, ForeGreen, ForeBlue, BackRed, BackGreen, BackBlue), nil)
}

// CreateGlyphCursorUnchecked sends an unchecked request.
func CreateGlyphCursorUnchecked(c *xgb.XConn, Cid Cursor, SourceFont Font, MaskFont Font, SourceChar uint16, MaskChar uint16, ForeRed uint16, ForeGreen uint16, ForeBlue uint16, BackRed uint16, BackGreen uint16, BackBlue uint16) error {
	var op uint8
	return c.Send(createGlyphCursorRequest(op, Cid, SourceFont, MaskFont, SourceChar, MaskChar, ForeRed, ForeGreen, ForeBlue, BackRed, BackGreen, BackBlue))
}

// Write request to wire for CreateGlyphCursor
// createGlyphCursorRequest writes a CreateGlyphCursor request to a byte slice.
func createGlyphCursorRequest(opcode uint8, Cid Cursor, SourceFont Font, MaskFont Font, SourceChar uint16, MaskChar uint16, ForeRed uint16, ForeGreen uint16, ForeBlue uint16, BackRed uint16, BackGreen uint16, BackBlue uint16) []byte {
	size := 32
	b := 0
	buf := make([]byte, size)

	buf[b] = 94 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Cid))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(SourceFont))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(MaskFont))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], SourceChar)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], MaskChar)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], ForeRed)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], ForeGreen)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], ForeBlue)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], BackRed)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], BackGreen)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], BackBlue)
	b += 2

	return buf
}

// CreatePixmap sends a checked request.
func CreatePixmap(c *xgb.XConn, Depth byte, Pid Pixmap, Drawable Drawable, Width uint16, Height uint16) error {
	var op uint8
	return c.SendRecv(createPixmapRequest(op, Depth, Pid, Drawable, Width, Height), nil)
}

// CreatePixmapUnchecked sends an unchecked request.
func CreatePixmapUnchecked(c *xgb.XConn, Depth byte, Pid Pixmap, Drawable Drawable, Width uint16, Height uint16) error {
	var op uint8
	return c.Send(createPixmapRequest(op, Depth, Pid, Drawable, Width, Height))
}

// Write request to wire for CreatePixmap
// createPixmapRequest writes a CreatePixmap request to a byte slice.
func createPixmapRequest(opcode uint8, Depth byte, Pid Pixmap, Drawable Drawable, Width uint16, Height uint16) []byte {
	size := 16
	b := 0
	buf := make([]byte, size)

	buf[b] = 53 // request opcode
	b += 1

	buf[b] = Depth
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

	return buf
}

// CreateWindow sends a checked request.
func CreateWindow(c *xgb.XConn, Depth byte, Wid Window, Parent Window, X int16, Y int16, Width uint16, Height uint16, BorderWidth uint16, Class uint16, Visual Visualid, ValueMask uint32, ValueList []uint32) error {
	var op uint8
	return c.SendRecv(createWindowRequest(op, Depth, Wid, Parent, X, Y, Width, Height, BorderWidth, Class, Visual, ValueMask, ValueList), nil)
}

// CreateWindowUnchecked sends an unchecked request.
func CreateWindowUnchecked(c *xgb.XConn, Depth byte, Wid Window, Parent Window, X int16, Y int16, Width uint16, Height uint16, BorderWidth uint16, Class uint16, Visual Visualid, ValueMask uint32, ValueList []uint32) error {
	var op uint8
	return c.Send(createWindowRequest(op, Depth, Wid, Parent, X, Y, Width, Height, BorderWidth, Class, Visual, ValueMask, ValueList))
}

// Write request to wire for CreateWindow
// createWindowRequest writes a CreateWindow request to a byte slice.
func createWindowRequest(opcode uint8, Depth byte, Wid Window, Parent Window, X int16, Y int16, Width uint16, Height uint16, BorderWidth uint16, Class uint16, Visual Visualid, ValueMask uint32, ValueList []uint32) []byte {
	size := internal.Pad4((32 + internal.Pad4((4 * internal.PopCount(int(ValueMask))))))
	b := 0
	buf := make([]byte, size)

	buf[b] = 1 // request opcode
	b += 1

	buf[b] = Depth
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Wid))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Parent))
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

	binary.LittleEndian.PutUint16(buf[b:], Class)
	b += 2

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

// DeleteProperty sends a checked request.
func DeleteProperty(c *xgb.XConn, Window Window, Property Atom) error {
	var op uint8
	return c.SendRecv(deletePropertyRequest(op, Window, Property), nil)
}

// DeletePropertyUnchecked sends an unchecked request.
func DeletePropertyUnchecked(c *xgb.XConn, Window Window, Property Atom) error {
	var op uint8
	return c.Send(deletePropertyRequest(op, Window, Property))
}

// Write request to wire for DeleteProperty
// deletePropertyRequest writes a DeleteProperty request to a byte slice.
func deletePropertyRequest(opcode uint8, Window Window, Property Atom) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = 19 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Property))
	b += 4

	return buf
}

// DestroySubwindows sends a checked request.
func DestroySubwindows(c *xgb.XConn, Window Window) error {
	var op uint8
	return c.SendRecv(destroySubwindowsRequest(op, Window), nil)
}

// DestroySubwindowsUnchecked sends an unchecked request.
func DestroySubwindowsUnchecked(c *xgb.XConn, Window Window) error {
	var op uint8
	return c.Send(destroySubwindowsRequest(op, Window))
}

// Write request to wire for DestroySubwindows
// destroySubwindowsRequest writes a DestroySubwindows request to a byte slice.
func destroySubwindowsRequest(opcode uint8, Window Window) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = 5 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	return buf
}

// DestroyWindow sends a checked request.
func DestroyWindow(c *xgb.XConn, Window Window) error {
	var op uint8
	return c.SendRecv(destroyWindowRequest(op, Window), nil)
}

// DestroyWindowUnchecked sends an unchecked request.
func DestroyWindowUnchecked(c *xgb.XConn, Window Window) error {
	var op uint8
	return c.Send(destroyWindowRequest(op, Window))
}

// Write request to wire for DestroyWindow
// destroyWindowRequest writes a DestroyWindow request to a byte slice.
func destroyWindowRequest(opcode uint8, Window Window) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = 4 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	return buf
}

// FillPoly sends a checked request.
func FillPoly(c *xgb.XConn, Drawable Drawable, Gc Gcontext, Shape byte, CoordinateMode byte, Points []Point) error {
	var op uint8
	return c.SendRecv(fillPolyRequest(op, Drawable, Gc, Shape, CoordinateMode, Points), nil)
}

// FillPolyUnchecked sends an unchecked request.
func FillPolyUnchecked(c *xgb.XConn, Drawable Drawable, Gc Gcontext, Shape byte, CoordinateMode byte, Points []Point) error {
	var op uint8
	return c.Send(fillPolyRequest(op, Drawable, Gc, Shape, CoordinateMode, Points))
}

// Write request to wire for FillPoly
// fillPolyRequest writes a FillPoly request to a byte slice.
func fillPolyRequest(opcode uint8, Drawable Drawable, Gc Gcontext, Shape byte, CoordinateMode byte, Points []Point) []byte {
	size := internal.Pad4((16 + internal.Pad4((len(Points) * 4))))
	b := 0
	buf := make([]byte, size)

	buf[b] = 69 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Drawable))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Gc))
	b += 4

	buf[b] = Shape
	b += 1

	buf[b] = CoordinateMode
	b += 1

	b += 2 // padding

	b += PointListBytes(buf[b:], Points)

	return buf
}

// ForceScreenSaver sends a checked request.
func ForceScreenSaver(c *xgb.XConn, Mode byte) error {
	var op uint8
	return c.SendRecv(forceScreenSaverRequest(op, Mode), nil)
}

// ForceScreenSaverUnchecked sends an unchecked request.
func ForceScreenSaverUnchecked(c *xgb.XConn, Mode byte) error {
	var op uint8
	return c.Send(forceScreenSaverRequest(op, Mode))
}

// Write request to wire for ForceScreenSaver
// forceScreenSaverRequest writes a ForceScreenSaver request to a byte slice.
func forceScreenSaverRequest(opcode uint8, Mode byte) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = 115 // request opcode
	b += 1

	buf[b] = Mode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// FreeColormap sends a checked request.
func FreeColormap(c *xgb.XConn, Cmap Colormap) error {
	var op uint8
	return c.SendRecv(freeColormapRequest(op, Cmap), nil)
}

// FreeColormapUnchecked sends an unchecked request.
func FreeColormapUnchecked(c *xgb.XConn, Cmap Colormap) error {
	var op uint8
	return c.Send(freeColormapRequest(op, Cmap))
}

// Write request to wire for FreeColormap
// freeColormapRequest writes a FreeColormap request to a byte slice.
func freeColormapRequest(opcode uint8, Cmap Colormap) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = 79 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Cmap))
	b += 4

	return buf
}

// FreeColors sends a checked request.
func FreeColors(c *xgb.XConn, Cmap Colormap, PlaneMask uint32, Pixels []uint32) error {
	var op uint8
	return c.SendRecv(freeColorsRequest(op, Cmap, PlaneMask, Pixels), nil)
}

// FreeColorsUnchecked sends an unchecked request.
func FreeColorsUnchecked(c *xgb.XConn, Cmap Colormap, PlaneMask uint32, Pixels []uint32) error {
	var op uint8
	return c.Send(freeColorsRequest(op, Cmap, PlaneMask, Pixels))
}

// Write request to wire for FreeColors
// freeColorsRequest writes a FreeColors request to a byte slice.
func freeColorsRequest(opcode uint8, Cmap Colormap, PlaneMask uint32, Pixels []uint32) []byte {
	size := internal.Pad4((12 + internal.Pad4((len(Pixels) * 4))))
	b := 0
	buf := make([]byte, size)

	buf[b] = 88 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Cmap))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], PlaneMask)
	b += 4

	for i := 0; i < int(len(Pixels)); i++ {
		binary.LittleEndian.PutUint32(buf[b:], Pixels[i])
		b += 4
	}

	return buf
}

// FreeCursor sends a checked request.
func FreeCursor(c *xgb.XConn, Cursor Cursor) error {
	var op uint8
	return c.SendRecv(freeCursorRequest(op, Cursor), nil)
}

// FreeCursorUnchecked sends an unchecked request.
func FreeCursorUnchecked(c *xgb.XConn, Cursor Cursor) error {
	var op uint8
	return c.Send(freeCursorRequest(op, Cursor))
}

// Write request to wire for FreeCursor
// freeCursorRequest writes a FreeCursor request to a byte slice.
func freeCursorRequest(opcode uint8, Cursor Cursor) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = 95 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Cursor))
	b += 4

	return buf
}

// FreeGC sends a checked request.
func FreeGC(c *xgb.XConn, Gc Gcontext) error {
	var op uint8
	return c.SendRecv(freeGCRequest(op, Gc), nil)
}

// FreeGCUnchecked sends an unchecked request.
func FreeGCUnchecked(c *xgb.XConn, Gc Gcontext) error {
	var op uint8
	return c.Send(freeGCRequest(op, Gc))
}

// Write request to wire for FreeGC
// freeGCRequest writes a FreeGC request to a byte slice.
func freeGCRequest(opcode uint8, Gc Gcontext) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = 60 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Gc))
	b += 4

	return buf
}

// FreePixmap sends a checked request.
func FreePixmap(c *xgb.XConn, Pixmap Pixmap) error {
	var op uint8
	return c.SendRecv(freePixmapRequest(op, Pixmap), nil)
}

// FreePixmapUnchecked sends an unchecked request.
func FreePixmapUnchecked(c *xgb.XConn, Pixmap Pixmap) error {
	var op uint8
	return c.Send(freePixmapRequest(op, Pixmap))
}

// Write request to wire for FreePixmap
// freePixmapRequest writes a FreePixmap request to a byte slice.
func freePixmapRequest(opcode uint8, Pixmap Pixmap) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = 54 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Pixmap))
	b += 4

	return buf
}

// GetAtomName sends a checked request.
func GetAtomName(c *xgb.XConn, Atom Atom) (GetAtomNameReply, error) {
	var reply GetAtomNameReply
	var op uint8
	err := c.SendRecv(getAtomNameRequest(op, Atom), &reply)
	return reply, err
}

// GetAtomNameUnchecked sends an unchecked request.
func GetAtomNameUnchecked(c *xgb.XConn, Atom Atom) error {
	var op uint8
	return c.Send(getAtomNameRequest(op, Atom))
}

// GetAtomNameReply represents the data returned from a GetAtomName request.
type GetAtomNameReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	NameLen uint16
	// padding: 22 bytes
	Name string // size: internal.Pad4((int(NameLen) * 1))
}

// Unmarshal reads a byte slice into a GetAtomNameReply value.
func (v *GetAtomNameReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.NameLen) * 1))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetAtomNameReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.NameLen = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	b += 22 // padding

	{
		byteString := make([]byte, v.NameLen)
		copy(byteString[:v.NameLen], buf[b:])
		v.Name = string(byteString)
		b += int(v.NameLen)
	}

	return nil
}

// Write request to wire for GetAtomName
// getAtomNameRequest writes a GetAtomName request to a byte slice.
func getAtomNameRequest(opcode uint8, Atom Atom) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = 17 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Atom))
	b += 4

	return buf
}

// GetFontPath sends a checked request.
func GetFontPath(c *xgb.XConn) (GetFontPathReply, error) {
	var reply GetFontPathReply
	var op uint8
	err := c.SendRecv(getFontPathRequest(op), &reply)
	return reply, err
}

// GetFontPathUnchecked sends an unchecked request.
func GetFontPathUnchecked(c *xgb.XConn) error {
	var op uint8
	return c.Send(getFontPathRequest(op))
}

// GetFontPathReply represents the data returned from a GetFontPath request.
type GetFontPathReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	PathLen uint16
	// padding: 22 bytes
	Path []Str // size: StrListSize(Path)
}

// Unmarshal reads a byte slice into a GetFontPathReply value.
func (v *GetFontPathReply) Unmarshal(buf []byte) error {
	if size := (32 + StrListSize(v.Path)); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetFontPathReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.PathLen = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	b += 22 // padding

	v.Path = make([]Str, v.PathLen)
	b += StrReadList(buf[b:], v.Path)

	return nil
}

// Write request to wire for GetFontPath
// getFontPathRequest writes a GetFontPath request to a byte slice.
func getFontPathRequest(opcode uint8) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = 52 // request opcode
	b += 1

	b += 1                                                 // padding
	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// GetGeometry sends a checked request.
func GetGeometry(c *xgb.XConn, Drawable Drawable) (GetGeometryReply, error) {
	var reply GetGeometryReply
	var op uint8
	err := c.SendRecv(getGeometryRequest(op, Drawable), &reply)
	return reply, err
}

// GetGeometryUnchecked sends an unchecked request.
func GetGeometryUnchecked(c *xgb.XConn, Drawable Drawable) error {
	var op uint8
	return c.Send(getGeometryRequest(op, Drawable))
}

// GetGeometryReply represents the data returned from a GetGeometry request.
type GetGeometryReply struct {
	Sequence    uint16 // sequence number of the request for this reply
	Length      uint32 // number of bytes in this reply
	Depth       byte
	Root        Window
	X           int16
	Y           int16
	Width       uint16
	Height      uint16
	BorderWidth uint16
	// padding: 2 bytes
}

// Unmarshal reads a byte slice into a GetGeometryReply value.
func (v *GetGeometryReply) Unmarshal(buf []byte) error {
	if size := 24; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetGeometryReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	v.Depth = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Root = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.X = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.Y = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.Width = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Height = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.BorderWidth = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	b += 2 // padding

	return nil
}

// Write request to wire for GetGeometry
// getGeometryRequest writes a GetGeometry request to a byte slice.
func getGeometryRequest(opcode uint8, Drawable Drawable) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = 14 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Drawable))
	b += 4

	return buf
}

// GetImage sends a checked request.
func GetImage(c *xgb.XConn, Format byte, Drawable Drawable, X int16, Y int16, Width uint16, Height uint16, PlaneMask uint32) (GetImageReply, error) {
	var reply GetImageReply
	var op uint8
	err := c.SendRecv(getImageRequest(op, Format, Drawable, X, Y, Width, Height, PlaneMask), &reply)
	return reply, err
}

// GetImageUnchecked sends an unchecked request.
func GetImageUnchecked(c *xgb.XConn, Format byte, Drawable Drawable, X int16, Y int16, Width uint16, Height uint16, PlaneMask uint32) error {
	var op uint8
	return c.Send(getImageRequest(op, Format, Drawable, X, Y, Width, Height, PlaneMask))
}

// GetImageReply represents the data returned from a GetImage request.
type GetImageReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	Depth    byte
	Visual   Visualid
	// padding: 20 bytes
	Data []byte // size: internal.Pad4(((int(Length) * 4) * 1))
}

// Unmarshal reads a byte slice into a GetImageReply value.
func (v *GetImageReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4(((int(v.Length) * 4) * 1))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetImageReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	v.Depth = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Visual = Visualid(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	b += 20 // padding

	v.Data = make([]byte, (int(v.Length) * 4))
	copy(v.Data[:(int(v.Length)*4)], buf[b:])
	b += int((int(v.Length) * 4))

	return nil
}

// Write request to wire for GetImage
// getImageRequest writes a GetImage request to a byte slice.
func getImageRequest(opcode uint8, Format byte, Drawable Drawable, X int16, Y int16, Width uint16, Height uint16, PlaneMask uint32) []byte {
	size := 20
	b := 0
	buf := make([]byte, size)

	buf[b] = 73 // request opcode
	b += 1

	buf[b] = Format
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

	return buf
}

// GetInputFocus sends a checked request.
func GetInputFocus(c *xgb.XConn) (GetInputFocusReply, error) {
	var reply GetInputFocusReply
	var op uint8
	err := c.SendRecv(getInputFocusRequest(op), &reply)
	return reply, err
}

// GetInputFocusUnchecked sends an unchecked request.
func GetInputFocusUnchecked(c *xgb.XConn) error {
	var op uint8
	return c.Send(getInputFocusRequest(op))
}

// GetInputFocusReply represents the data returned from a GetInputFocus request.
type GetInputFocusReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	RevertTo byte
	Focus    Window
}

// Unmarshal reads a byte slice into a GetInputFocusReply value.
func (v *GetInputFocusReply) Unmarshal(buf []byte) error {
	if size := 12; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetInputFocusReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	v.RevertTo = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Focus = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	return nil
}

// Write request to wire for GetInputFocus
// getInputFocusRequest writes a GetInputFocus request to a byte slice.
func getInputFocusRequest(opcode uint8) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = 43 // request opcode
	b += 1

	b += 1                                                 // padding
	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// GetKeyboardControl sends a checked request.
func GetKeyboardControl(c *xgb.XConn) (GetKeyboardControlReply, error) {
	var reply GetKeyboardControlReply
	var op uint8
	err := c.SendRecv(getKeyboardControlRequest(op), &reply)
	return reply, err
}

// GetKeyboardControlUnchecked sends an unchecked request.
func GetKeyboardControlUnchecked(c *xgb.XConn) error {
	var op uint8
	return c.Send(getKeyboardControlRequest(op))
}

// GetKeyboardControlReply represents the data returned from a GetKeyboardControl request.
type GetKeyboardControlReply struct {
	Sequence         uint16 // sequence number of the request for this reply
	Length           uint32 // number of bytes in this reply
	GlobalAutoRepeat byte
	LedMask          uint32
	KeyClickPercent  byte
	BellPercent      byte
	BellPitch        uint16
	BellDuration     uint16
	// padding: 2 bytes
	AutoRepeats []byte // size: 32
}

// Unmarshal reads a byte slice into a GetKeyboardControlReply value.
func (v *GetKeyboardControlReply) Unmarshal(buf []byte) error {
	if size := 52; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetKeyboardControlReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	v.GlobalAutoRepeat = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.LedMask = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.KeyClickPercent = buf[b]
	b += 1

	v.BellPercent = buf[b]
	b += 1

	v.BellPitch = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.BellDuration = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	b += 2 // padding

	v.AutoRepeats = make([]byte, 32)
	copy(v.AutoRepeats[:32], buf[b:])
	b += int(32)

	return nil
}

// Write request to wire for GetKeyboardControl
// getKeyboardControlRequest writes a GetKeyboardControl request to a byte slice.
func getKeyboardControlRequest(opcode uint8) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = 103 // request opcode
	b += 1

	b += 1                                                 // padding
	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// GetKeyboardMapping sends a checked request.
func GetKeyboardMapping(c *xgb.XConn, FirstKeycode Keycode, Count byte) (GetKeyboardMappingReply, error) {
	var reply GetKeyboardMappingReply
	var op uint8
	err := c.SendRecv(getKeyboardMappingRequest(op, FirstKeycode, Count), &reply)
	return reply, err
}

// GetKeyboardMappingUnchecked sends an unchecked request.
func GetKeyboardMappingUnchecked(c *xgb.XConn, FirstKeycode Keycode, Count byte) error {
	var op uint8
	return c.Send(getKeyboardMappingRequest(op, FirstKeycode, Count))
}

// GetKeyboardMappingReply represents the data returned from a GetKeyboardMapping request.
type GetKeyboardMappingReply struct {
	Sequence          uint16 // sequence number of the request for this reply
	Length            uint32 // number of bytes in this reply
	KeysymsPerKeycode byte
	// padding: 24 bytes
	Keysyms []Keysym // size: internal.Pad4((int(Length) * 4))
}

// Unmarshal reads a byte slice into a GetKeyboardMappingReply value.
func (v *GetKeyboardMappingReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.Length) * 4))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetKeyboardMappingReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	v.KeysymsPerKeycode = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	b += 24 // padding

	v.Keysyms = make([]Keysym, v.Length)
	for i := 0; i < int(v.Length); i++ {
		v.Keysyms[i] = Keysym(binary.LittleEndian.Uint32(buf[b:]))
		b += 4
	}

	return nil
}

// Write request to wire for GetKeyboardMapping
// getKeyboardMappingRequest writes a GetKeyboardMapping request to a byte slice.
func getKeyboardMappingRequest(opcode uint8, FirstKeycode Keycode, Count byte) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = 101 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	buf[b] = uint8(FirstKeycode)
	b += 1

	buf[b] = Count
	b += 1

	return buf
}

// GetModifierMapping sends a checked request.
func GetModifierMapping(c *xgb.XConn) (GetModifierMappingReply, error) {
	var reply GetModifierMappingReply
	var op uint8
	err := c.SendRecv(getModifierMappingRequest(op), &reply)
	return reply, err
}

// GetModifierMappingUnchecked sends an unchecked request.
func GetModifierMappingUnchecked(c *xgb.XConn) error {
	var op uint8
	return c.Send(getModifierMappingRequest(op))
}

// GetModifierMappingReply represents the data returned from a GetModifierMapping request.
type GetModifierMappingReply struct {
	Sequence            uint16 // sequence number of the request for this reply
	Length              uint32 // number of bytes in this reply
	KeycodesPerModifier byte
	// padding: 24 bytes
	Keycodes []Keycode // size: internal.Pad4(((int(KeycodesPerModifier) * 8) * 1))
}

// Unmarshal reads a byte slice into a GetModifierMappingReply value.
func (v *GetModifierMappingReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4(((int(v.KeycodesPerModifier) * 8) * 1))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetModifierMappingReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	v.KeycodesPerModifier = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	b += 24 // padding

	v.Keycodes = make([]Keycode, (int(v.KeycodesPerModifier) * 8))
	for i := 0; i < int((int(v.KeycodesPerModifier) * 8)); i++ {
		v.Keycodes[i] = Keycode(buf[b])
		b += 1
	}

	return nil
}

// Write request to wire for GetModifierMapping
// getModifierMappingRequest writes a GetModifierMapping request to a byte slice.
func getModifierMappingRequest(opcode uint8) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = 119 // request opcode
	b += 1

	b += 1                                                 // padding
	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// GetMotionEvents sends a checked request.
func GetMotionEvents(c *xgb.XConn, Window Window, Start Timestamp, Stop Timestamp) (GetMotionEventsReply, error) {
	var reply GetMotionEventsReply
	var op uint8
	err := c.SendRecv(getMotionEventsRequest(op, Window, Start, Stop), &reply)
	return reply, err
}

// GetMotionEventsUnchecked sends an unchecked request.
func GetMotionEventsUnchecked(c *xgb.XConn, Window Window, Start Timestamp, Stop Timestamp) error {
	var op uint8
	return c.Send(getMotionEventsRequest(op, Window, Start, Stop))
}

// GetMotionEventsReply represents the data returned from a GetMotionEvents request.
type GetMotionEventsReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	EventsLen uint32
	// padding: 20 bytes
	Events []Timecoord // size: internal.Pad4((int(EventsLen) * 8))
}

// Unmarshal reads a byte slice into a GetMotionEventsReply value.
func (v *GetMotionEventsReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.EventsLen) * 8))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetMotionEventsReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.EventsLen = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 20 // padding

	v.Events = make([]Timecoord, v.EventsLen)
	b += TimecoordReadList(buf[b:], v.Events)

	return nil
}

// Write request to wire for GetMotionEvents
// getMotionEventsRequest writes a GetMotionEvents request to a byte slice.
func getMotionEventsRequest(opcode uint8, Window Window, Start Timestamp, Stop Timestamp) []byte {
	size := 16
	b := 0
	buf := make([]byte, size)

	buf[b] = 39 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Start))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Stop))
	b += 4

	return buf
}

// GetPointerControl sends a checked request.
func GetPointerControl(c *xgb.XConn) (GetPointerControlReply, error) {
	var reply GetPointerControlReply
	var op uint8
	err := c.SendRecv(getPointerControlRequest(op), &reply)
	return reply, err
}

// GetPointerControlUnchecked sends an unchecked request.
func GetPointerControlUnchecked(c *xgb.XConn) error {
	var op uint8
	return c.Send(getPointerControlRequest(op))
}

// GetPointerControlReply represents the data returned from a GetPointerControl request.
type GetPointerControlReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	AccelerationNumerator   uint16
	AccelerationDenominator uint16
	Threshold               uint16
	// padding: 18 bytes
}

// Unmarshal reads a byte slice into a GetPointerControlReply value.
func (v *GetPointerControlReply) Unmarshal(buf []byte) error {
	if size := 32; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetPointerControlReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.AccelerationNumerator = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.AccelerationDenominator = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Threshold = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	b += 18 // padding

	return nil
}

// Write request to wire for GetPointerControl
// getPointerControlRequest writes a GetPointerControl request to a byte slice.
func getPointerControlRequest(opcode uint8) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = 106 // request opcode
	b += 1

	b += 1                                                 // padding
	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// GetPointerMapping sends a checked request.
func GetPointerMapping(c *xgb.XConn) (GetPointerMappingReply, error) {
	var reply GetPointerMappingReply
	var op uint8
	err := c.SendRecv(getPointerMappingRequest(op), &reply)
	return reply, err
}

// GetPointerMappingUnchecked sends an unchecked request.
func GetPointerMappingUnchecked(c *xgb.XConn) error {
	var op uint8
	return c.Send(getPointerMappingRequest(op))
}

// GetPointerMappingReply represents the data returned from a GetPointerMapping request.
type GetPointerMappingReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	MapLen   byte
	// padding: 24 bytes
	Map []byte // size: internal.Pad4((int(MapLen) * 1))
}

// Unmarshal reads a byte slice into a GetPointerMappingReply value.
func (v *GetPointerMappingReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.MapLen) * 1))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetPointerMappingReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	v.MapLen = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	b += 24 // padding

	v.Map = make([]byte, v.MapLen)
	copy(v.Map[:v.MapLen], buf[b:])
	b += int(v.MapLen)

	return nil
}

// Write request to wire for GetPointerMapping
// getPointerMappingRequest writes a GetPointerMapping request to a byte slice.
func getPointerMappingRequest(opcode uint8) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = 117 // request opcode
	b += 1

	b += 1                                                 // padding
	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// GetProperty sends a checked request.
func GetProperty(c *xgb.XConn, Delete bool, Window Window, Property Atom, Type Atom, LongOffset uint32, LongLength uint32) (GetPropertyReply, error) {
	var reply GetPropertyReply
	var op uint8
	err := c.SendRecv(getPropertyRequest(op, Delete, Window, Property, Type, LongOffset, LongLength), &reply)
	return reply, err
}

// GetPropertyUnchecked sends an unchecked request.
func GetPropertyUnchecked(c *xgb.XConn, Delete bool, Window Window, Property Atom, Type Atom, LongOffset uint32, LongLength uint32) error {
	var op uint8
	return c.Send(getPropertyRequest(op, Delete, Window, Property, Type, LongOffset, LongLength))
}

// GetPropertyReply represents the data returned from a GetProperty request.
type GetPropertyReply struct {
	Sequence   uint16 // sequence number of the request for this reply
	Length     uint32 // number of bytes in this reply
	Format     byte
	Type       Atom
	BytesAfter uint32
	ValueLen   uint32
	// padding: 12 bytes
	Value []byte // size: internal.Pad4(((int(ValueLen) * (int(Format) / 8)) * 1))
}

// Unmarshal reads a byte slice into a GetPropertyReply value.
func (v *GetPropertyReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4(((int(v.ValueLen) * (int(v.Format) / 8)) * 1))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetPropertyReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	v.Format = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Type = Atom(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.BytesAfter = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.ValueLen = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 12 // padding

	v.Value = make([]byte, (int(v.ValueLen) * (int(v.Format) / 8)))
	copy(v.Value[:(int(v.ValueLen)*(int(v.Format)/8))], buf[b:])
	b += int((int(v.ValueLen) * (int(v.Format) / 8)))

	return nil
}

// Write request to wire for GetProperty
// getPropertyRequest writes a GetProperty request to a byte slice.
func getPropertyRequest(opcode uint8, Delete bool, Window Window, Property Atom, Type Atom, LongOffset uint32, LongLength uint32) []byte {
	size := 24
	b := 0
	buf := make([]byte, size)

	buf[b] = 20 // request opcode
	b += 1

	if Delete {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Property))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Type))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], LongOffset)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], LongLength)
	b += 4

	return buf
}

// GetScreenSaver sends a checked request.
func GetScreenSaver(c *xgb.XConn) (GetScreenSaverReply, error) {
	var reply GetScreenSaverReply
	var op uint8
	err := c.SendRecv(getScreenSaverRequest(op), &reply)
	return reply, err
}

// GetScreenSaverUnchecked sends an unchecked request.
func GetScreenSaverUnchecked(c *xgb.XConn) error {
	var op uint8
	return c.Send(getScreenSaverRequest(op))
}

// GetScreenSaverReply represents the data returned from a GetScreenSaver request.
type GetScreenSaverReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Timeout        uint16
	Interval       uint16
	PreferBlanking byte
	AllowExposures byte
	// padding: 18 bytes
}

// Unmarshal reads a byte slice into a GetScreenSaverReply value.
func (v *GetScreenSaverReply) Unmarshal(buf []byte) error {
	if size := 32; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetScreenSaverReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Timeout = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Interval = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.PreferBlanking = buf[b]
	b += 1

	v.AllowExposures = buf[b]
	b += 1

	b += 18 // padding

	return nil
}

// Write request to wire for GetScreenSaver
// getScreenSaverRequest writes a GetScreenSaver request to a byte slice.
func getScreenSaverRequest(opcode uint8) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = 108 // request opcode
	b += 1

	b += 1                                                 // padding
	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// GetSelectionOwner sends a checked request.
func GetSelectionOwner(c *xgb.XConn, Selection Atom) (GetSelectionOwnerReply, error) {
	var reply GetSelectionOwnerReply
	var op uint8
	err := c.SendRecv(getSelectionOwnerRequest(op, Selection), &reply)
	return reply, err
}

// GetSelectionOwnerUnchecked sends an unchecked request.
func GetSelectionOwnerUnchecked(c *xgb.XConn, Selection Atom) error {
	var op uint8
	return c.Send(getSelectionOwnerRequest(op, Selection))
}

// GetSelectionOwnerReply represents the data returned from a GetSelectionOwner request.
type GetSelectionOwnerReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Owner Window
}

// Unmarshal reads a byte slice into a GetSelectionOwnerReply value.
func (v *GetSelectionOwnerReply) Unmarshal(buf []byte) error {
	if size := 12; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetSelectionOwnerReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Owner = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	return nil
}

// Write request to wire for GetSelectionOwner
// getSelectionOwnerRequest writes a GetSelectionOwner request to a byte slice.
func getSelectionOwnerRequest(opcode uint8, Selection Atom) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = 23 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Selection))
	b += 4

	return buf
}

// GetWindowAttributes sends a checked request.
func GetWindowAttributes(c *xgb.XConn, Window Window) (GetWindowAttributesReply, error) {
	var reply GetWindowAttributesReply
	var op uint8
	err := c.SendRecv(getWindowAttributesRequest(op, Window), &reply)
	return reply, err
}

// GetWindowAttributesUnchecked sends an unchecked request.
func GetWindowAttributesUnchecked(c *xgb.XConn, Window Window) error {
	var op uint8
	return c.Send(getWindowAttributesRequest(op, Window))
}

// GetWindowAttributesReply represents the data returned from a GetWindowAttributes request.
type GetWindowAttributesReply struct {
	Sequence           uint16 // sequence number of the request for this reply
	Length             uint32 // number of bytes in this reply
	BackingStore       byte
	Visual             Visualid
	Class              uint16
	BitGravity         byte
	WinGravity         byte
	BackingPlanes      uint32
	BackingPixel       uint32
	SaveUnder          bool
	MapIsInstalled     bool
	MapState           byte
	OverrideRedirect   bool
	Colormap           Colormap
	AllEventMasks      uint32
	YourEventMask      uint32
	DoNotPropagateMask uint16
	// padding: 2 bytes
}

// Unmarshal reads a byte slice into a GetWindowAttributesReply value.
func (v *GetWindowAttributesReply) Unmarshal(buf []byte) error {
	if size := 44; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetWindowAttributesReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	v.BackingStore = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Visual = Visualid(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Class = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.BitGravity = buf[b]
	b += 1

	v.WinGravity = buf[b]
	b += 1

	v.BackingPlanes = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.BackingPixel = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.SaveUnder = (buf[b] == 1)
	b += 1

	v.MapIsInstalled = (buf[b] == 1)
	b += 1

	v.MapState = buf[b]
	b += 1

	v.OverrideRedirect = (buf[b] == 1)
	b += 1

	v.Colormap = Colormap(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.AllEventMasks = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.YourEventMask = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.DoNotPropagateMask = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	b += 2 // padding

	return nil
}

// Write request to wire for GetWindowAttributes
// getWindowAttributesRequest writes a GetWindowAttributes request to a byte slice.
func getWindowAttributesRequest(opcode uint8, Window Window) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = 3 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	return buf
}

// GrabButton sends a checked request.
func GrabButton(c *xgb.XConn, OwnerEvents bool, GrabWindow Window, EventMask uint16, PointerMode byte, KeyboardMode byte, ConfineTo Window, Cursor Cursor, Button byte, Modifiers uint16) error {
	var op uint8
	return c.SendRecv(grabButtonRequest(op, OwnerEvents, GrabWindow, EventMask, PointerMode, KeyboardMode, ConfineTo, Cursor, Button, Modifiers), nil)
}

// GrabButtonUnchecked sends an unchecked request.
func GrabButtonUnchecked(c *xgb.XConn, OwnerEvents bool, GrabWindow Window, EventMask uint16, PointerMode byte, KeyboardMode byte, ConfineTo Window, Cursor Cursor, Button byte, Modifiers uint16) error {
	var op uint8
	return c.Send(grabButtonRequest(op, OwnerEvents, GrabWindow, EventMask, PointerMode, KeyboardMode, ConfineTo, Cursor, Button, Modifiers))
}

// Write request to wire for GrabButton
// grabButtonRequest writes a GrabButton request to a byte slice.
func grabButtonRequest(opcode uint8, OwnerEvents bool, GrabWindow Window, EventMask uint16, PointerMode byte, KeyboardMode byte, ConfineTo Window, Cursor Cursor, Button byte, Modifiers uint16) []byte {
	size := 24
	b := 0
	buf := make([]byte, size)

	buf[b] = 28 // request opcode
	b += 1

	if OwnerEvents {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(GrabWindow))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], EventMask)
	b += 2

	buf[b] = PointerMode
	b += 1

	buf[b] = KeyboardMode
	b += 1

	binary.LittleEndian.PutUint32(buf[b:], uint32(ConfineTo))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Cursor))
	b += 4

	buf[b] = Button
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], Modifiers)
	b += 2

	return buf
}

// GrabKey sends a checked request.
func GrabKey(c *xgb.XConn, OwnerEvents bool, GrabWindow Window, Modifiers uint16, Key Keycode, PointerMode byte, KeyboardMode byte) error {
	var op uint8
	return c.SendRecv(grabKeyRequest(op, OwnerEvents, GrabWindow, Modifiers, Key, PointerMode, KeyboardMode), nil)
}

// GrabKeyUnchecked sends an unchecked request.
func GrabKeyUnchecked(c *xgb.XConn, OwnerEvents bool, GrabWindow Window, Modifiers uint16, Key Keycode, PointerMode byte, KeyboardMode byte) error {
	var op uint8
	return c.Send(grabKeyRequest(op, OwnerEvents, GrabWindow, Modifiers, Key, PointerMode, KeyboardMode))
}

// Write request to wire for GrabKey
// grabKeyRequest writes a GrabKey request to a byte slice.
func grabKeyRequest(opcode uint8, OwnerEvents bool, GrabWindow Window, Modifiers uint16, Key Keycode, PointerMode byte, KeyboardMode byte) []byte {
	size := 16
	b := 0
	buf := make([]byte, size)

	buf[b] = 33 // request opcode
	b += 1

	if OwnerEvents {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(GrabWindow))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], Modifiers)
	b += 2

	buf[b] = uint8(Key)
	b += 1

	buf[b] = PointerMode
	b += 1

	buf[b] = KeyboardMode
	b += 1

	b += 3 // padding

	return buf
}

// GrabKeyboard sends a checked request.
func GrabKeyboard(c *xgb.XConn, OwnerEvents bool, GrabWindow Window, Time Timestamp, PointerMode byte, KeyboardMode byte) (GrabKeyboardReply, error) {
	var reply GrabKeyboardReply
	var op uint8
	err := c.SendRecv(grabKeyboardRequest(op, OwnerEvents, GrabWindow, Time, PointerMode, KeyboardMode), &reply)
	return reply, err
}

// GrabKeyboardUnchecked sends an unchecked request.
func GrabKeyboardUnchecked(c *xgb.XConn, OwnerEvents bool, GrabWindow Window, Time Timestamp, PointerMode byte, KeyboardMode byte) error {
	var op uint8
	return c.Send(grabKeyboardRequest(op, OwnerEvents, GrabWindow, Time, PointerMode, KeyboardMode))
}

// GrabKeyboardReply represents the data returned from a GrabKeyboard request.
type GrabKeyboardReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	Status   byte
}

// Unmarshal reads a byte slice into a GrabKeyboardReply value.
func (v *GrabKeyboardReply) Unmarshal(buf []byte) error {
	if size := 8; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GrabKeyboardReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	v.Status = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	return nil
}

// Write request to wire for GrabKeyboard
// grabKeyboardRequest writes a GrabKeyboard request to a byte slice.
func grabKeyboardRequest(opcode uint8, OwnerEvents bool, GrabWindow Window, Time Timestamp, PointerMode byte, KeyboardMode byte) []byte {
	size := 16
	b := 0
	buf := make([]byte, size)

	buf[b] = 31 // request opcode
	b += 1

	if OwnerEvents {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(GrabWindow))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Time))
	b += 4

	buf[b] = PointerMode
	b += 1

	buf[b] = KeyboardMode
	b += 1

	b += 2 // padding

	return buf
}

// GrabPointer sends a checked request.
func GrabPointer(c *xgb.XConn, OwnerEvents bool, GrabWindow Window, EventMask uint16, PointerMode byte, KeyboardMode byte, ConfineTo Window, Cursor Cursor, Time Timestamp) (GrabPointerReply, error) {
	var reply GrabPointerReply
	var op uint8
	err := c.SendRecv(grabPointerRequest(op, OwnerEvents, GrabWindow, EventMask, PointerMode, KeyboardMode, ConfineTo, Cursor, Time), &reply)
	return reply, err
}

// GrabPointerUnchecked sends an unchecked request.
func GrabPointerUnchecked(c *xgb.XConn, OwnerEvents bool, GrabWindow Window, EventMask uint16, PointerMode byte, KeyboardMode byte, ConfineTo Window, Cursor Cursor, Time Timestamp) error {
	var op uint8
	return c.Send(grabPointerRequest(op, OwnerEvents, GrabWindow, EventMask, PointerMode, KeyboardMode, ConfineTo, Cursor, Time))
}

// GrabPointerReply represents the data returned from a GrabPointer request.
type GrabPointerReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	Status   byte
}

// Unmarshal reads a byte slice into a GrabPointerReply value.
func (v *GrabPointerReply) Unmarshal(buf []byte) error {
	if size := 8; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GrabPointerReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	v.Status = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	return nil
}

// Write request to wire for GrabPointer
// grabPointerRequest writes a GrabPointer request to a byte slice.
func grabPointerRequest(opcode uint8, OwnerEvents bool, GrabWindow Window, EventMask uint16, PointerMode byte, KeyboardMode byte, ConfineTo Window, Cursor Cursor, Time Timestamp) []byte {
	size := 24
	b := 0
	buf := make([]byte, size)

	buf[b] = 26 // request opcode
	b += 1

	if OwnerEvents {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(GrabWindow))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], EventMask)
	b += 2

	buf[b] = PointerMode
	b += 1

	buf[b] = KeyboardMode
	b += 1

	binary.LittleEndian.PutUint32(buf[b:], uint32(ConfineTo))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Cursor))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Time))
	b += 4

	return buf
}

// GrabServer sends a checked request.
func GrabServer(c *xgb.XConn) error {
	var op uint8
	return c.SendRecv(grabServerRequest(op), nil)
}

// GrabServerUnchecked sends an unchecked request.
func GrabServerUnchecked(c *xgb.XConn) error {
	var op uint8
	return c.Send(grabServerRequest(op))
}

// Write request to wire for GrabServer
// grabServerRequest writes a GrabServer request to a byte slice.
func grabServerRequest(opcode uint8) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = 36 // request opcode
	b += 1

	b += 1                                                 // padding
	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// ImageText16 sends a checked request.
func ImageText16(c *xgb.XConn, StringLen byte, Drawable Drawable, Gc Gcontext, X int16, Y int16, String []Char2b) error {
	var op uint8
	return c.SendRecv(imageText16Request(op, StringLen, Drawable, Gc, X, Y, String), nil)
}

// ImageText16Unchecked sends an unchecked request.
func ImageText16Unchecked(c *xgb.XConn, StringLen byte, Drawable Drawable, Gc Gcontext, X int16, Y int16, String []Char2b) error {
	var op uint8
	return c.Send(imageText16Request(op, StringLen, Drawable, Gc, X, Y, String))
}

// Write request to wire for ImageText16
// imageText16Request writes a ImageText16 request to a byte slice.
func imageText16Request(opcode uint8, StringLen byte, Drawable Drawable, Gc Gcontext, X int16, Y int16, String []Char2b) []byte {
	size := internal.Pad4((16 + internal.Pad4((int(StringLen) * 2))))
	b := 0
	buf := make([]byte, size)

	buf[b] = 77 // request opcode
	b += 1

	buf[b] = StringLen
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Drawable))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Gc))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(X))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(Y))
	b += 2

	b += Char2bListBytes(buf[b:], String)

	return buf
}

// ImageText8 sends a checked request.
func ImageText8(c *xgb.XConn, StringLen byte, Drawable Drawable, Gc Gcontext, X int16, Y int16, String string) error {
	var op uint8
	return c.SendRecv(imageText8Request(op, StringLen, Drawable, Gc, X, Y, String), nil)
}

// ImageText8Unchecked sends an unchecked request.
func ImageText8Unchecked(c *xgb.XConn, StringLen byte, Drawable Drawable, Gc Gcontext, X int16, Y int16, String string) error {
	var op uint8
	return c.Send(imageText8Request(op, StringLen, Drawable, Gc, X, Y, String))
}

// Write request to wire for ImageText8
// imageText8Request writes a ImageText8 request to a byte slice.
func imageText8Request(opcode uint8, StringLen byte, Drawable Drawable, Gc Gcontext, X int16, Y int16, String string) []byte {
	size := internal.Pad4((16 + internal.Pad4((int(StringLen) * 1))))
	b := 0
	buf := make([]byte, size)

	buf[b] = 76 // request opcode
	b += 1

	buf[b] = StringLen
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Drawable))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Gc))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(X))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(Y))
	b += 2

	copy(buf[b:], String[:StringLen])
	b += int(StringLen)

	return buf
}

// InstallColormap sends a checked request.
func InstallColormap(c *xgb.XConn, Cmap Colormap) error {
	var op uint8
	return c.SendRecv(installColormapRequest(op, Cmap), nil)
}

// InstallColormapUnchecked sends an unchecked request.
func InstallColormapUnchecked(c *xgb.XConn, Cmap Colormap) error {
	var op uint8
	return c.Send(installColormapRequest(op, Cmap))
}

// Write request to wire for InstallColormap
// installColormapRequest writes a InstallColormap request to a byte slice.
func installColormapRequest(opcode uint8, Cmap Colormap) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = 81 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Cmap))
	b += 4

	return buf
}

// InternAtom sends a checked request.
func InternAtom(c *xgb.XConn, OnlyIfExists bool, NameLen uint16, Name string) (InternAtomReply, error) {
	var reply InternAtomReply
	var op uint8
	err := c.SendRecv(internAtomRequest(op, OnlyIfExists, NameLen, Name), &reply)
	return reply, err
}

// InternAtomUnchecked sends an unchecked request.
func InternAtomUnchecked(c *xgb.XConn, OnlyIfExists bool, NameLen uint16, Name string) error {
	var op uint8
	return c.Send(internAtomRequest(op, OnlyIfExists, NameLen, Name))
}

// InternAtomReply represents the data returned from a InternAtom request.
type InternAtomReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Atom Atom
}

// Unmarshal reads a byte slice into a InternAtomReply value.
func (v *InternAtomReply) Unmarshal(buf []byte) error {
	if size := 12; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"InternAtomReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Atom = Atom(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	return nil
}

// Write request to wire for InternAtom
// internAtomRequest writes a InternAtom request to a byte slice.
func internAtomRequest(opcode uint8, OnlyIfExists bool, NameLen uint16, Name string) []byte {
	size := internal.Pad4((8 + internal.Pad4((int(NameLen) * 1))))
	b := 0
	buf := make([]byte, size)

	buf[b] = 16 // request opcode
	b += 1

	if OnlyIfExists {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], NameLen)
	b += 2

	b += 2 // padding

	copy(buf[b:], Name[:NameLen])
	b += int(NameLen)

	return buf
}

// KillClient sends a checked request.
func KillClient(c *xgb.XConn, Resource uint32) error {
	var op uint8
	return c.SendRecv(killClientRequest(op, Resource), nil)
}

// KillClientUnchecked sends an unchecked request.
func KillClientUnchecked(c *xgb.XConn, Resource uint32) error {
	var op uint8
	return c.Send(killClientRequest(op, Resource))
}

// Write request to wire for KillClient
// killClientRequest writes a KillClient request to a byte slice.
func killClientRequest(opcode uint8, Resource uint32) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = 113 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], Resource)
	b += 4

	return buf
}

// ListExtensions sends a checked request.
func ListExtensions(c *xgb.XConn) (ListExtensionsReply, error) {
	var reply ListExtensionsReply
	var op uint8
	err := c.SendRecv(listExtensionsRequest(op), &reply)
	return reply, err
}

// ListExtensionsUnchecked sends an unchecked request.
func ListExtensionsUnchecked(c *xgb.XConn) error {
	var op uint8
	return c.Send(listExtensionsRequest(op))
}

// ListExtensionsReply represents the data returned from a ListExtensions request.
type ListExtensionsReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	NamesLen byte
	// padding: 24 bytes
	Names []Str // size: StrListSize(Names)
}

// Unmarshal reads a byte slice into a ListExtensionsReply value.
func (v *ListExtensionsReply) Unmarshal(buf []byte) error {
	if size := (32 + StrListSize(v.Names)); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"ListExtensionsReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	v.NamesLen = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	b += 24 // padding

	v.Names = make([]Str, v.NamesLen)
	b += StrReadList(buf[b:], v.Names)

	return nil
}

// Write request to wire for ListExtensions
// listExtensionsRequest writes a ListExtensions request to a byte slice.
func listExtensionsRequest(opcode uint8) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = 99 // request opcode
	b += 1

	b += 1                                                 // padding
	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// ListFonts sends a checked request.
func ListFonts(c *xgb.XConn, MaxNames uint16, PatternLen uint16, Pattern string) (ListFontsReply, error) {
	var reply ListFontsReply
	var op uint8
	err := c.SendRecv(listFontsRequest(op, MaxNames, PatternLen, Pattern), &reply)
	return reply, err
}

// ListFontsUnchecked sends an unchecked request.
func ListFontsUnchecked(c *xgb.XConn, MaxNames uint16, PatternLen uint16, Pattern string) error {
	var op uint8
	return c.Send(listFontsRequest(op, MaxNames, PatternLen, Pattern))
}

// ListFontsReply represents the data returned from a ListFonts request.
type ListFontsReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	NamesLen uint16
	// padding: 22 bytes
	Names []Str // size: StrListSize(Names)
}

// Unmarshal reads a byte slice into a ListFontsReply value.
func (v *ListFontsReply) Unmarshal(buf []byte) error {
	if size := (32 + StrListSize(v.Names)); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"ListFontsReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.NamesLen = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	b += 22 // padding

	v.Names = make([]Str, v.NamesLen)
	b += StrReadList(buf[b:], v.Names)

	return nil
}

// Write request to wire for ListFonts
// listFontsRequest writes a ListFonts request to a byte slice.
func listFontsRequest(opcode uint8, MaxNames uint16, PatternLen uint16, Pattern string) []byte {
	size := internal.Pad4((8 + internal.Pad4((int(PatternLen) * 1))))
	b := 0
	buf := make([]byte, size)

	buf[b] = 49 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], MaxNames)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], PatternLen)
	b += 2

	copy(buf[b:], Pattern[:PatternLen])
	b += int(PatternLen)

	return buf
}

// ListFontsWithInfo sends a checked request.
func ListFontsWithInfo(c *xgb.XConn, MaxNames uint16, PatternLen uint16, Pattern string) (ListFontsWithInfoReply, error) {
	var reply ListFontsWithInfoReply
	var op uint8
	err := c.SendRecv(listFontsWithInfoRequest(op, MaxNames, PatternLen, Pattern), &reply)
	return reply, err
}

// ListFontsWithInfoUnchecked sends an unchecked request.
func ListFontsWithInfoUnchecked(c *xgb.XConn, MaxNames uint16, PatternLen uint16, Pattern string) error {
	var op uint8
	return c.Send(listFontsWithInfoRequest(op, MaxNames, PatternLen, Pattern))
}

// ListFontsWithInfoReply represents the data returned from a ListFontsWithInfo request.
type ListFontsWithInfoReply struct {
	Sequence  uint16 // sequence number of the request for this reply
	Length    uint32 // number of bytes in this reply
	NameLen   byte
	MinBounds Charinfo
	// padding: 4 bytes
	MaxBounds Charinfo
	// padding: 4 bytes
	MinCharOrByte2 uint16
	MaxCharOrByte2 uint16
	DefaultChar    uint16
	PropertiesLen  uint16
	DrawDirection  byte
	MinByte1       byte
	MaxByte1       byte
	AllCharsExist  bool
	FontAscent     int16
	FontDescent    int16
	RepliesHint    uint32
	Properties     []Fontprop // size: internal.Pad4((int(PropertiesLen) * 8))
	Name           string     // size: internal.Pad4((int(NameLen) * 1))
}

// Unmarshal reads a byte slice into a ListFontsWithInfoReply value.
func (v *ListFontsWithInfoReply) Unmarshal(buf []byte) error {
	if size := ((60 + internal.Pad4((int(v.PropertiesLen) * 8))) + internal.Pad4((int(v.NameLen) * 1))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"ListFontsWithInfoReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	v.NameLen = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.MinBounds = Charinfo{}
	b += CharinfoRead(buf[b:], &v.MinBounds)

	b += 4 // padding

	v.MaxBounds = Charinfo{}
	b += CharinfoRead(buf[b:], &v.MaxBounds)

	b += 4 // padding

	v.MinCharOrByte2 = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.MaxCharOrByte2 = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.DefaultChar = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.PropertiesLen = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.DrawDirection = buf[b]
	b += 1

	v.MinByte1 = buf[b]
	b += 1

	v.MaxByte1 = buf[b]
	b += 1

	v.AllCharsExist = (buf[b] == 1)
	b += 1

	v.FontAscent = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.FontDescent = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.RepliesHint = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.Properties = make([]Fontprop, v.PropertiesLen)
	b += FontpropReadList(buf[b:], v.Properties)

	{
		byteString := make([]byte, v.NameLen)
		copy(byteString[:v.NameLen], buf[b:])
		v.Name = string(byteString)
		b += int(v.NameLen)
	}

	return nil
}

// Write request to wire for ListFontsWithInfo
// listFontsWithInfoRequest writes a ListFontsWithInfo request to a byte slice.
func listFontsWithInfoRequest(opcode uint8, MaxNames uint16, PatternLen uint16, Pattern string) []byte {
	size := internal.Pad4((8 + internal.Pad4((int(PatternLen) * 1))))
	b := 0
	buf := make([]byte, size)

	buf[b] = 50 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], MaxNames)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], PatternLen)
	b += 2

	copy(buf[b:], Pattern[:PatternLen])
	b += int(PatternLen)

	return buf
}

// ListHosts sends a checked request.
func ListHosts(c *xgb.XConn) (ListHostsReply, error) {
	var reply ListHostsReply
	var op uint8
	err := c.SendRecv(listHostsRequest(op), &reply)
	return reply, err
}

// ListHostsUnchecked sends an unchecked request.
func ListHostsUnchecked(c *xgb.XConn) error {
	var op uint8
	return c.Send(listHostsRequest(op))
}

// ListHostsReply represents the data returned from a ListHosts request.
type ListHostsReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	Mode     byte
	HostsLen uint16
	// padding: 22 bytes
	Hosts []Host // size: HostListSize(Hosts)
}

// Unmarshal reads a byte slice into a ListHostsReply value.
func (v *ListHostsReply) Unmarshal(buf []byte) error {
	if size := (32 + HostListSize(v.Hosts)); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"ListHostsReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	v.Mode = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.HostsLen = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	b += 22 // padding

	v.Hosts = make([]Host, v.HostsLen)
	b += HostReadList(buf[b:], v.Hosts)

	return nil
}

// Write request to wire for ListHosts
// listHostsRequest writes a ListHosts request to a byte slice.
func listHostsRequest(opcode uint8) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = 110 // request opcode
	b += 1

	b += 1                                                 // padding
	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// ListInstalledColormaps sends a checked request.
func ListInstalledColormaps(c *xgb.XConn, Window Window) (ListInstalledColormapsReply, error) {
	var reply ListInstalledColormapsReply
	var op uint8
	err := c.SendRecv(listInstalledColormapsRequest(op, Window), &reply)
	return reply, err
}

// ListInstalledColormapsUnchecked sends an unchecked request.
func ListInstalledColormapsUnchecked(c *xgb.XConn, Window Window) error {
	var op uint8
	return c.Send(listInstalledColormapsRequest(op, Window))
}

// ListInstalledColormapsReply represents the data returned from a ListInstalledColormaps request.
type ListInstalledColormapsReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	CmapsLen uint16
	// padding: 22 bytes
	Cmaps []Colormap // size: internal.Pad4((int(CmapsLen) * 4))
}

// Unmarshal reads a byte slice into a ListInstalledColormapsReply value.
func (v *ListInstalledColormapsReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.CmapsLen) * 4))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"ListInstalledColormapsReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.CmapsLen = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	b += 22 // padding

	v.Cmaps = make([]Colormap, v.CmapsLen)
	for i := 0; i < int(v.CmapsLen); i++ {
		v.Cmaps[i] = Colormap(binary.LittleEndian.Uint32(buf[b:]))
		b += 4
	}

	return nil
}

// Write request to wire for ListInstalledColormaps
// listInstalledColormapsRequest writes a ListInstalledColormaps request to a byte slice.
func listInstalledColormapsRequest(opcode uint8, Window Window) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = 83 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	return buf
}

// ListProperties sends a checked request.
func ListProperties(c *xgb.XConn, Window Window) (ListPropertiesReply, error) {
	var reply ListPropertiesReply
	var op uint8
	err := c.SendRecv(listPropertiesRequest(op, Window), &reply)
	return reply, err
}

// ListPropertiesUnchecked sends an unchecked request.
func ListPropertiesUnchecked(c *xgb.XConn, Window Window) error {
	var op uint8
	return c.Send(listPropertiesRequest(op, Window))
}

// ListPropertiesReply represents the data returned from a ListProperties request.
type ListPropertiesReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	AtomsLen uint16
	// padding: 22 bytes
	Atoms []Atom // size: internal.Pad4((int(AtomsLen) * 4))
}

// Unmarshal reads a byte slice into a ListPropertiesReply value.
func (v *ListPropertiesReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.AtomsLen) * 4))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"ListPropertiesReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.AtomsLen = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	b += 22 // padding

	v.Atoms = make([]Atom, v.AtomsLen)
	for i := 0; i < int(v.AtomsLen); i++ {
		v.Atoms[i] = Atom(binary.LittleEndian.Uint32(buf[b:]))
		b += 4
	}

	return nil
}

// Write request to wire for ListProperties
// listPropertiesRequest writes a ListProperties request to a byte slice.
func listPropertiesRequest(opcode uint8, Window Window) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = 21 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	return buf
}

// LookupColor sends a checked request.
func LookupColor(c *xgb.XConn, Cmap Colormap, NameLen uint16, Name string) (LookupColorReply, error) {
	var reply LookupColorReply
	var op uint8
	err := c.SendRecv(lookupColorRequest(op, Cmap, NameLen, Name), &reply)
	return reply, err
}

// LookupColorUnchecked sends an unchecked request.
func LookupColorUnchecked(c *xgb.XConn, Cmap Colormap, NameLen uint16, Name string) error {
	var op uint8
	return c.Send(lookupColorRequest(op, Cmap, NameLen, Name))
}

// LookupColorReply represents the data returned from a LookupColor request.
type LookupColorReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	ExactRed    uint16
	ExactGreen  uint16
	ExactBlue   uint16
	VisualRed   uint16
	VisualGreen uint16
	VisualBlue  uint16
}

// Unmarshal reads a byte slice into a LookupColorReply value.
func (v *LookupColorReply) Unmarshal(buf []byte) error {
	if size := 20; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"LookupColorReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.ExactRed = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.ExactGreen = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.ExactBlue = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.VisualRed = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.VisualGreen = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.VisualBlue = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return nil
}

// Write request to wire for LookupColor
// lookupColorRequest writes a LookupColor request to a byte slice.
func lookupColorRequest(opcode uint8, Cmap Colormap, NameLen uint16, Name string) []byte {
	size := internal.Pad4((12 + internal.Pad4((int(NameLen) * 1))))
	b := 0
	buf := make([]byte, size)

	buf[b] = 92 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Cmap))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], NameLen)
	b += 2

	b += 2 // padding

	copy(buf[b:], Name[:NameLen])
	b += int(NameLen)

	return buf
}

// MapSubwindows sends a checked request.
func MapSubwindows(c *xgb.XConn, Window Window) error {
	var op uint8
	return c.SendRecv(mapSubwindowsRequest(op, Window), nil)
}

// MapSubwindowsUnchecked sends an unchecked request.
func MapSubwindowsUnchecked(c *xgb.XConn, Window Window) error {
	var op uint8
	return c.Send(mapSubwindowsRequest(op, Window))
}

// Write request to wire for MapSubwindows
// mapSubwindowsRequest writes a MapSubwindows request to a byte slice.
func mapSubwindowsRequest(opcode uint8, Window Window) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = 9 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	return buf
}

// MapWindow sends a checked request.
func MapWindow(c *xgb.XConn, Window Window) error {
	var op uint8
	return c.SendRecv(mapWindowRequest(op, Window), nil)
}

// MapWindowUnchecked sends an unchecked request.
func MapWindowUnchecked(c *xgb.XConn, Window Window) error {
	var op uint8
	return c.Send(mapWindowRequest(op, Window))
}

// Write request to wire for MapWindow
// mapWindowRequest writes a MapWindow request to a byte slice.
func mapWindowRequest(opcode uint8, Window Window) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = 8 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	return buf
}

// NoOperation sends a checked request.
func NoOperation(c *xgb.XConn) error {
	var op uint8
	return c.SendRecv(noOperationRequest(op), nil)
}

// NoOperationUnchecked sends an unchecked request.
func NoOperationUnchecked(c *xgb.XConn) error {
	var op uint8
	return c.Send(noOperationRequest(op))
}

// Write request to wire for NoOperation
// noOperationRequest writes a NoOperation request to a byte slice.
func noOperationRequest(opcode uint8) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = 127 // request opcode
	b += 1

	b += 1                                                 // padding
	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// OpenFont sends a checked request.
func OpenFont(c *xgb.XConn, Fid Font, NameLen uint16, Name string) error {
	var op uint8
	return c.SendRecv(openFontRequest(op, Fid, NameLen, Name), nil)
}

// OpenFontUnchecked sends an unchecked request.
func OpenFontUnchecked(c *xgb.XConn, Fid Font, NameLen uint16, Name string) error {
	var op uint8
	return c.Send(openFontRequest(op, Fid, NameLen, Name))
}

// Write request to wire for OpenFont
// openFontRequest writes a OpenFont request to a byte slice.
func openFontRequest(opcode uint8, Fid Font, NameLen uint16, Name string) []byte {
	size := internal.Pad4((12 + internal.Pad4((int(NameLen) * 1))))
	b := 0
	buf := make([]byte, size)

	buf[b] = 45 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Fid))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], NameLen)
	b += 2

	b += 2 // padding

	copy(buf[b:], Name[:NameLen])
	b += int(NameLen)

	return buf
}

// PolyArc sends a checked request.
func PolyArc(c *xgb.XConn, Drawable Drawable, Gc Gcontext, Arcs []Arc) error {
	var op uint8
	return c.SendRecv(polyArcRequest(op, Drawable, Gc, Arcs), nil)
}

// PolyArcUnchecked sends an unchecked request.
func PolyArcUnchecked(c *xgb.XConn, Drawable Drawable, Gc Gcontext, Arcs []Arc) error {
	var op uint8
	return c.Send(polyArcRequest(op, Drawable, Gc, Arcs))
}

// Write request to wire for PolyArc
// polyArcRequest writes a PolyArc request to a byte slice.
func polyArcRequest(opcode uint8, Drawable Drawable, Gc Gcontext, Arcs []Arc) []byte {
	size := internal.Pad4((12 + internal.Pad4((len(Arcs) * 12))))
	b := 0
	buf := make([]byte, size)

	buf[b] = 68 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Drawable))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Gc))
	b += 4

	b += ArcListBytes(buf[b:], Arcs)

	return buf
}

// PolyFillArc sends a checked request.
func PolyFillArc(c *xgb.XConn, Drawable Drawable, Gc Gcontext, Arcs []Arc) error {
	var op uint8
	return c.SendRecv(polyFillArcRequest(op, Drawable, Gc, Arcs), nil)
}

// PolyFillArcUnchecked sends an unchecked request.
func PolyFillArcUnchecked(c *xgb.XConn, Drawable Drawable, Gc Gcontext, Arcs []Arc) error {
	var op uint8
	return c.Send(polyFillArcRequest(op, Drawable, Gc, Arcs))
}

// Write request to wire for PolyFillArc
// polyFillArcRequest writes a PolyFillArc request to a byte slice.
func polyFillArcRequest(opcode uint8, Drawable Drawable, Gc Gcontext, Arcs []Arc) []byte {
	size := internal.Pad4((12 + internal.Pad4((len(Arcs) * 12))))
	b := 0
	buf := make([]byte, size)

	buf[b] = 71 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Drawable))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Gc))
	b += 4

	b += ArcListBytes(buf[b:], Arcs)

	return buf
}

// PolyFillRectangle sends a checked request.
func PolyFillRectangle(c *xgb.XConn, Drawable Drawable, Gc Gcontext, Rectangles []Rectangle) error {
	var op uint8
	return c.SendRecv(polyFillRectangleRequest(op, Drawable, Gc, Rectangles), nil)
}

// PolyFillRectangleUnchecked sends an unchecked request.
func PolyFillRectangleUnchecked(c *xgb.XConn, Drawable Drawable, Gc Gcontext, Rectangles []Rectangle) error {
	var op uint8
	return c.Send(polyFillRectangleRequest(op, Drawable, Gc, Rectangles))
}

// Write request to wire for PolyFillRectangle
// polyFillRectangleRequest writes a PolyFillRectangle request to a byte slice.
func polyFillRectangleRequest(opcode uint8, Drawable Drawable, Gc Gcontext, Rectangles []Rectangle) []byte {
	size := internal.Pad4((12 + internal.Pad4((len(Rectangles) * 8))))
	b := 0
	buf := make([]byte, size)

	buf[b] = 70 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Drawable))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Gc))
	b += 4

	b += RectangleListBytes(buf[b:], Rectangles)

	return buf
}

// PolyLine sends a checked request.
func PolyLine(c *xgb.XConn, CoordinateMode byte, Drawable Drawable, Gc Gcontext, Points []Point) error {
	var op uint8
	return c.SendRecv(polyLineRequest(op, CoordinateMode, Drawable, Gc, Points), nil)
}

// PolyLineUnchecked sends an unchecked request.
func PolyLineUnchecked(c *xgb.XConn, CoordinateMode byte, Drawable Drawable, Gc Gcontext, Points []Point) error {
	var op uint8
	return c.Send(polyLineRequest(op, CoordinateMode, Drawable, Gc, Points))
}

// Write request to wire for PolyLine
// polyLineRequest writes a PolyLine request to a byte slice.
func polyLineRequest(opcode uint8, CoordinateMode byte, Drawable Drawable, Gc Gcontext, Points []Point) []byte {
	size := internal.Pad4((12 + internal.Pad4((len(Points) * 4))))
	b := 0
	buf := make([]byte, size)

	buf[b] = 65 // request opcode
	b += 1

	buf[b] = CoordinateMode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Drawable))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Gc))
	b += 4

	b += PointListBytes(buf[b:], Points)

	return buf
}

// PolyPoint sends a checked request.
func PolyPoint(c *xgb.XConn, CoordinateMode byte, Drawable Drawable, Gc Gcontext, Points []Point) error {
	var op uint8
	return c.SendRecv(polyPointRequest(op, CoordinateMode, Drawable, Gc, Points), nil)
}

// PolyPointUnchecked sends an unchecked request.
func PolyPointUnchecked(c *xgb.XConn, CoordinateMode byte, Drawable Drawable, Gc Gcontext, Points []Point) error {
	var op uint8
	return c.Send(polyPointRequest(op, CoordinateMode, Drawable, Gc, Points))
}

// Write request to wire for PolyPoint
// polyPointRequest writes a PolyPoint request to a byte slice.
func polyPointRequest(opcode uint8, CoordinateMode byte, Drawable Drawable, Gc Gcontext, Points []Point) []byte {
	size := internal.Pad4((12 + internal.Pad4((len(Points) * 4))))
	b := 0
	buf := make([]byte, size)

	buf[b] = 64 // request opcode
	b += 1

	buf[b] = CoordinateMode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Drawable))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Gc))
	b += 4

	b += PointListBytes(buf[b:], Points)

	return buf
}

// PolyRectangle sends a checked request.
func PolyRectangle(c *xgb.XConn, Drawable Drawable, Gc Gcontext, Rectangles []Rectangle) error {
	var op uint8
	return c.SendRecv(polyRectangleRequest(op, Drawable, Gc, Rectangles), nil)
}

// PolyRectangleUnchecked sends an unchecked request.
func PolyRectangleUnchecked(c *xgb.XConn, Drawable Drawable, Gc Gcontext, Rectangles []Rectangle) error {
	var op uint8
	return c.Send(polyRectangleRequest(op, Drawable, Gc, Rectangles))
}

// Write request to wire for PolyRectangle
// polyRectangleRequest writes a PolyRectangle request to a byte slice.
func polyRectangleRequest(opcode uint8, Drawable Drawable, Gc Gcontext, Rectangles []Rectangle) []byte {
	size := internal.Pad4((12 + internal.Pad4((len(Rectangles) * 8))))
	b := 0
	buf := make([]byte, size)

	buf[b] = 67 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Drawable))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Gc))
	b += 4

	b += RectangleListBytes(buf[b:], Rectangles)

	return buf
}

// PolySegment sends a checked request.
func PolySegment(c *xgb.XConn, Drawable Drawable, Gc Gcontext, Segments []Segment) error {
	var op uint8
	return c.SendRecv(polySegmentRequest(op, Drawable, Gc, Segments), nil)
}

// PolySegmentUnchecked sends an unchecked request.
func PolySegmentUnchecked(c *xgb.XConn, Drawable Drawable, Gc Gcontext, Segments []Segment) error {
	var op uint8
	return c.Send(polySegmentRequest(op, Drawable, Gc, Segments))
}

// Write request to wire for PolySegment
// polySegmentRequest writes a PolySegment request to a byte slice.
func polySegmentRequest(opcode uint8, Drawable Drawable, Gc Gcontext, Segments []Segment) []byte {
	size := internal.Pad4((12 + internal.Pad4((len(Segments) * 8))))
	b := 0
	buf := make([]byte, size)

	buf[b] = 66 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Drawable))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Gc))
	b += 4

	b += SegmentListBytes(buf[b:], Segments)

	return buf
}

// PolyText16 sends a checked request.
func PolyText16(c *xgb.XConn, Drawable Drawable, Gc Gcontext, X int16, Y int16, Items []byte) error {
	var op uint8
	return c.SendRecv(polyText16Request(op, Drawable, Gc, X, Y, Items), nil)
}

// PolyText16Unchecked sends an unchecked request.
func PolyText16Unchecked(c *xgb.XConn, Drawable Drawable, Gc Gcontext, X int16, Y int16, Items []byte) error {
	var op uint8
	return c.Send(polyText16Request(op, Drawable, Gc, X, Y, Items))
}

// Write request to wire for PolyText16
// polyText16Request writes a PolyText16 request to a byte slice.
func polyText16Request(opcode uint8, Drawable Drawable, Gc Gcontext, X int16, Y int16, Items []byte) []byte {
	size := internal.Pad4((16 + internal.Pad4((len(Items) * 1))))
	b := 0
	buf := make([]byte, size)

	buf[b] = 75 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Drawable))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Gc))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(X))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(Y))
	b += 2

	copy(buf[b:], Items[:len(Items)])
	b += int(len(Items))

	return buf
}

// PolyText8 sends a checked request.
func PolyText8(c *xgb.XConn, Drawable Drawable, Gc Gcontext, X int16, Y int16, Items []byte) error {
	var op uint8
	return c.SendRecv(polyText8Request(op, Drawable, Gc, X, Y, Items), nil)
}

// PolyText8Unchecked sends an unchecked request.
func PolyText8Unchecked(c *xgb.XConn, Drawable Drawable, Gc Gcontext, X int16, Y int16, Items []byte) error {
	var op uint8
	return c.Send(polyText8Request(op, Drawable, Gc, X, Y, Items))
}

// Write request to wire for PolyText8
// polyText8Request writes a PolyText8 request to a byte slice.
func polyText8Request(opcode uint8, Drawable Drawable, Gc Gcontext, X int16, Y int16, Items []byte) []byte {
	size := internal.Pad4((16 + internal.Pad4((len(Items) * 1))))
	b := 0
	buf := make([]byte, size)

	buf[b] = 74 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Drawable))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Gc))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(X))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(Y))
	b += 2

	copy(buf[b:], Items[:len(Items)])
	b += int(len(Items))

	return buf
}

// PutImage sends a checked request.
func PutImage(c *xgb.XConn, Format byte, Drawable Drawable, Gc Gcontext, Width uint16, Height uint16, DstX int16, DstY int16, LeftPad byte, Depth byte, Data []byte) error {
	var op uint8
	return c.SendRecv(putImageRequest(op, Format, Drawable, Gc, Width, Height, DstX, DstY, LeftPad, Depth, Data), nil)
}

// PutImageUnchecked sends an unchecked request.
func PutImageUnchecked(c *xgb.XConn, Format byte, Drawable Drawable, Gc Gcontext, Width uint16, Height uint16, DstX int16, DstY int16, LeftPad byte, Depth byte, Data []byte) error {
	var op uint8
	return c.Send(putImageRequest(op, Format, Drawable, Gc, Width, Height, DstX, DstY, LeftPad, Depth, Data))
}

// Write request to wire for PutImage
// putImageRequest writes a PutImage request to a byte slice.
func putImageRequest(opcode uint8, Format byte, Drawable Drawable, Gc Gcontext, Width uint16, Height uint16, DstX int16, DstY int16, LeftPad byte, Depth byte, Data []byte) []byte {
	size := internal.Pad4((24 + internal.Pad4((len(Data) * 1))))
	b := 0
	buf := make([]byte, size)

	buf[b] = 72 // request opcode
	b += 1

	buf[b] = Format
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Drawable))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Gc))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], Width)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Height)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(DstX))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(DstY))
	b += 2

	buf[b] = LeftPad
	b += 1

	buf[b] = Depth
	b += 1

	b += 2 // padding

	copy(buf[b:], Data[:len(Data)])
	b += int(len(Data))

	return buf
}

// QueryBestSize sends a checked request.
func QueryBestSize(c *xgb.XConn, Class byte, Drawable Drawable, Width uint16, Height uint16) (QueryBestSizeReply, error) {
	var reply QueryBestSizeReply
	var op uint8
	err := c.SendRecv(queryBestSizeRequest(op, Class, Drawable, Width, Height), &reply)
	return reply, err
}

// QueryBestSizeUnchecked sends an unchecked request.
func QueryBestSizeUnchecked(c *xgb.XConn, Class byte, Drawable Drawable, Width uint16, Height uint16) error {
	var op uint8
	return c.Send(queryBestSizeRequest(op, Class, Drawable, Width, Height))
}

// QueryBestSizeReply represents the data returned from a QueryBestSize request.
type QueryBestSizeReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Width  uint16
	Height uint16
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

	v.Width = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Height = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return nil
}

// Write request to wire for QueryBestSize
// queryBestSizeRequest writes a QueryBestSize request to a byte slice.
func queryBestSizeRequest(opcode uint8, Class byte, Drawable Drawable, Width uint16, Height uint16) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = 97 // request opcode
	b += 1

	buf[b] = Class
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Drawable))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], Width)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Height)
	b += 2

	return buf
}

// QueryColors sends a checked request.
func QueryColors(c *xgb.XConn, Cmap Colormap, Pixels []uint32) (QueryColorsReply, error) {
	var reply QueryColorsReply
	var op uint8
	err := c.SendRecv(queryColorsRequest(op, Cmap, Pixels), &reply)
	return reply, err
}

// QueryColorsUnchecked sends an unchecked request.
func QueryColorsUnchecked(c *xgb.XConn, Cmap Colormap, Pixels []uint32) error {
	var op uint8
	return c.Send(queryColorsRequest(op, Cmap, Pixels))
}

// QueryColorsReply represents the data returned from a QueryColors request.
type QueryColorsReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	ColorsLen uint16
	// padding: 22 bytes
	Colors []Rgb // size: internal.Pad4((int(ColorsLen) * 8))
}

// Unmarshal reads a byte slice into a QueryColorsReply value.
func (v *QueryColorsReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.ColorsLen) * 8))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"QueryColorsReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.ColorsLen = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	b += 22 // padding

	v.Colors = make([]Rgb, v.ColorsLen)
	b += RgbReadList(buf[b:], v.Colors)

	return nil
}

// Write request to wire for QueryColors
// queryColorsRequest writes a QueryColors request to a byte slice.
func queryColorsRequest(opcode uint8, Cmap Colormap, Pixels []uint32) []byte {
	size := internal.Pad4((8 + internal.Pad4((len(Pixels) * 4))))
	b := 0
	buf := make([]byte, size)

	buf[b] = 91 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Cmap))
	b += 4

	for i := 0; i < int(len(Pixels)); i++ {
		binary.LittleEndian.PutUint32(buf[b:], Pixels[i])
		b += 4
	}

	return buf
}

// QueryExtension sends a checked request.
func QueryExtension(c *xgb.XConn, NameLen uint16, Name string) (QueryExtensionReply, error) {
	var reply QueryExtensionReply
	var op uint8
	err := c.SendRecv(queryExtensionRequest(op, NameLen, Name), &reply)
	return reply, err
}

// QueryExtensionUnchecked sends an unchecked request.
func QueryExtensionUnchecked(c *xgb.XConn, NameLen uint16, Name string) error {
	var op uint8
	return c.Send(queryExtensionRequest(op, NameLen, Name))
}

// QueryExtensionReply represents the data returned from a QueryExtension request.
type QueryExtensionReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Present     bool
	MajorOpcode byte
	FirstEvent  byte
	FirstError  byte
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

	v.Present = (buf[b] == 1)
	b += 1

	v.MajorOpcode = buf[b]
	b += 1

	v.FirstEvent = buf[b]
	b += 1

	v.FirstError = buf[b]
	b += 1

	return nil
}

// Write request to wire for QueryExtension
// queryExtensionRequest writes a QueryExtension request to a byte slice.
func queryExtensionRequest(opcode uint8, NameLen uint16, Name string) []byte {
	size := internal.Pad4((8 + internal.Pad4((int(NameLen) * 1))))
	b := 0
	buf := make([]byte, size)

	buf[b] = 98 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], NameLen)
	b += 2

	b += 2 // padding

	copy(buf[b:], Name[:NameLen])
	b += int(NameLen)

	return buf
}

// QueryFont sends a checked request.
func QueryFont(c *xgb.XConn, Font Fontable) (QueryFontReply, error) {
	var reply QueryFontReply
	var op uint8
	err := c.SendRecv(queryFontRequest(op, Font), &reply)
	return reply, err
}

// QueryFontUnchecked sends an unchecked request.
func QueryFontUnchecked(c *xgb.XConn, Font Fontable) error {
	var op uint8
	return c.Send(queryFontRequest(op, Font))
}

// QueryFontReply represents the data returned from a QueryFont request.
type QueryFontReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	MinBounds Charinfo
	// padding: 4 bytes
	MaxBounds Charinfo
	// padding: 4 bytes
	MinCharOrByte2 uint16
	MaxCharOrByte2 uint16
	DefaultChar    uint16
	PropertiesLen  uint16
	DrawDirection  byte
	MinByte1       byte
	MaxByte1       byte
	AllCharsExist  bool
	FontAscent     int16
	FontDescent    int16
	CharInfosLen   uint32
	Properties     []Fontprop // size: internal.Pad4((int(PropertiesLen) * 8))
	// alignment gap to multiple of 4
	CharInfos []Charinfo // size: internal.Pad4((int(CharInfosLen) * 12))
}

// Unmarshal reads a byte slice into a QueryFontReply value.
func (v *QueryFontReply) Unmarshal(buf []byte) error {
	if size := (((60 + internal.Pad4((int(v.PropertiesLen) * 8))) + 4) + internal.Pad4((int(v.CharInfosLen) * 12))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"QueryFontReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.MinBounds = Charinfo{}
	b += CharinfoRead(buf[b:], &v.MinBounds)

	b += 4 // padding

	v.MaxBounds = Charinfo{}
	b += CharinfoRead(buf[b:], &v.MaxBounds)

	b += 4 // padding

	v.MinCharOrByte2 = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.MaxCharOrByte2 = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.DefaultChar = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.PropertiesLen = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.DrawDirection = buf[b]
	b += 1

	v.MinByte1 = buf[b]
	b += 1

	v.MaxByte1 = buf[b]
	b += 1

	v.AllCharsExist = (buf[b] == 1)
	b += 1

	v.FontAscent = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.FontDescent = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.CharInfosLen = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.Properties = make([]Fontprop, v.PropertiesLen)
	b += FontpropReadList(buf[b:], v.Properties)

	b = (b + 3) & ^3 // alignment gap

	v.CharInfos = make([]Charinfo, v.CharInfosLen)
	b += CharinfoReadList(buf[b:], v.CharInfos)

	return nil
}

// Write request to wire for QueryFont
// queryFontRequest writes a QueryFont request to a byte slice.
func queryFontRequest(opcode uint8, Font Fontable) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = 47 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Font))
	b += 4

	return buf
}

// QueryKeymap sends a checked request.
func QueryKeymap(c *xgb.XConn) (QueryKeymapReply, error) {
	var reply QueryKeymapReply
	var op uint8
	err := c.SendRecv(queryKeymapRequest(op), &reply)
	return reply, err
}

// QueryKeymapUnchecked sends an unchecked request.
func QueryKeymapUnchecked(c *xgb.XConn) error {
	var op uint8
	return c.Send(queryKeymapRequest(op))
}

// QueryKeymapReply represents the data returned from a QueryKeymap request.
type QueryKeymapReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Keys []byte // size: 32
}

// Unmarshal reads a byte slice into a QueryKeymapReply value.
func (v *QueryKeymapReply) Unmarshal(buf []byte) error {
	if size := 40; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"QueryKeymapReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Keys = make([]byte, 32)
	copy(v.Keys[:32], buf[b:])
	b += int(32)

	return nil
}

// Write request to wire for QueryKeymap
// queryKeymapRequest writes a QueryKeymap request to a byte slice.
func queryKeymapRequest(opcode uint8) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = 44 // request opcode
	b += 1

	b += 1                                                 // padding
	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// QueryPointer sends a checked request.
func QueryPointer(c *xgb.XConn, Window Window) (QueryPointerReply, error) {
	var reply QueryPointerReply
	var op uint8
	err := c.SendRecv(queryPointerRequest(op, Window), &reply)
	return reply, err
}

// QueryPointerUnchecked sends an unchecked request.
func QueryPointerUnchecked(c *xgb.XConn, Window Window) error {
	var op uint8
	return c.Send(queryPointerRequest(op, Window))
}

// QueryPointerReply represents the data returned from a QueryPointer request.
type QueryPointerReply struct {
	Sequence   uint16 // sequence number of the request for this reply
	Length     uint32 // number of bytes in this reply
	SameScreen bool
	Root       Window
	Child      Window
	RootX      int16
	RootY      int16
	WinX       int16
	WinY       int16
	Mask       uint16
	// padding: 2 bytes
}

// Unmarshal reads a byte slice into a QueryPointerReply value.
func (v *QueryPointerReply) Unmarshal(buf []byte) error {
	if size := 28; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"QueryPointerReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	v.SameScreen = (buf[b] == 1)
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Root = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Child = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.RootX = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.RootY = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.WinX = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.WinY = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.Mask = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	b += 2 // padding

	return nil
}

// Write request to wire for QueryPointer
// queryPointerRequest writes a QueryPointer request to a byte slice.
func queryPointerRequest(opcode uint8, Window Window) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = 38 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	return buf
}

// QueryTextExtents sends a checked request.
func QueryTextExtents(c *xgb.XConn, Font Fontable, String []Char2b, StringLen uint16) (QueryTextExtentsReply, error) {
	var reply QueryTextExtentsReply
	var op uint8
	err := c.SendRecv(queryTextExtentsRequest(op, Font, String, StringLen), &reply)
	return reply, err
}

// QueryTextExtentsUnchecked sends an unchecked request.
func QueryTextExtentsUnchecked(c *xgb.XConn, Font Fontable, String []Char2b, StringLen uint16) error {
	var op uint8
	return c.Send(queryTextExtentsRequest(op, Font, String, StringLen))
}

// QueryTextExtentsReply represents the data returned from a QueryTextExtents request.
type QueryTextExtentsReply struct {
	Sequence       uint16 // sequence number of the request for this reply
	Length         uint32 // number of bytes in this reply
	DrawDirection  byte
	FontAscent     int16
	FontDescent    int16
	OverallAscent  int16
	OverallDescent int16
	OverallWidth   int32
	OverallLeft    int32
	OverallRight   int32
}

// Unmarshal reads a byte slice into a QueryTextExtentsReply value.
func (v *QueryTextExtentsReply) Unmarshal(buf []byte) error {
	if size := 28; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"QueryTextExtentsReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	v.DrawDirection = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.FontAscent = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.FontDescent = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.OverallAscent = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.OverallDescent = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.OverallWidth = int32(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.OverallLeft = int32(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.OverallRight = int32(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	return nil
}

// Write request to wire for QueryTextExtents
// queryTextExtentsRequest writes a QueryTextExtents request to a byte slice.
func queryTextExtentsRequest(opcode uint8, Font Fontable, String []Char2b, StringLen uint16) []byte {
	size := internal.Pad4((8 + internal.Pad4((len(String) * 2))))
	b := 0
	buf := make([]byte, size)

	buf[b] = 48 // request opcode
	b += 1

	buf[b] = byte((int(StringLen) & 1))
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Font))
	b += 4

	b += Char2bListBytes(buf[b:], String)

	// skip writing local field: StringLen (2) :: uint16

	return buf
}

// QueryTree sends a checked request.
func QueryTree(c *xgb.XConn, Window Window) (QueryTreeReply, error) {
	var reply QueryTreeReply
	var op uint8
	err := c.SendRecv(queryTreeRequest(op, Window), &reply)
	return reply, err
}

// QueryTreeUnchecked sends an unchecked request.
func QueryTreeUnchecked(c *xgb.XConn, Window Window) error {
	var op uint8
	return c.Send(queryTreeRequest(op, Window))
}

// QueryTreeReply represents the data returned from a QueryTree request.
type QueryTreeReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Root        Window
	Parent      Window
	ChildrenLen uint16
	// padding: 14 bytes
	Children []Window // size: internal.Pad4((int(ChildrenLen) * 4))
}

// Unmarshal reads a byte slice into a QueryTreeReply value.
func (v *QueryTreeReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.ChildrenLen) * 4))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"QueryTreeReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Root = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Parent = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.ChildrenLen = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	b += 14 // padding

	v.Children = make([]Window, v.ChildrenLen)
	for i := 0; i < int(v.ChildrenLen); i++ {
		v.Children[i] = Window(binary.LittleEndian.Uint32(buf[b:]))
		b += 4
	}

	return nil
}

// Write request to wire for QueryTree
// queryTreeRequest writes a QueryTree request to a byte slice.
func queryTreeRequest(opcode uint8, Window Window) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = 15 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	return buf
}

// RecolorCursor sends a checked request.
func RecolorCursor(c *xgb.XConn, Cursor Cursor, ForeRed uint16, ForeGreen uint16, ForeBlue uint16, BackRed uint16, BackGreen uint16, BackBlue uint16) error {
	var op uint8
	return c.SendRecv(recolorCursorRequest(op, Cursor, ForeRed, ForeGreen, ForeBlue, BackRed, BackGreen, BackBlue), nil)
}

// RecolorCursorUnchecked sends an unchecked request.
func RecolorCursorUnchecked(c *xgb.XConn, Cursor Cursor, ForeRed uint16, ForeGreen uint16, ForeBlue uint16, BackRed uint16, BackGreen uint16, BackBlue uint16) error {
	var op uint8
	return c.Send(recolorCursorRequest(op, Cursor, ForeRed, ForeGreen, ForeBlue, BackRed, BackGreen, BackBlue))
}

// Write request to wire for RecolorCursor
// recolorCursorRequest writes a RecolorCursor request to a byte slice.
func recolorCursorRequest(opcode uint8, Cursor Cursor, ForeRed uint16, ForeGreen uint16, ForeBlue uint16, BackRed uint16, BackGreen uint16, BackBlue uint16) []byte {
	size := 20
	b := 0
	buf := make([]byte, size)

	buf[b] = 96 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Cursor))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], ForeRed)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], ForeGreen)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], ForeBlue)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], BackRed)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], BackGreen)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], BackBlue)
	b += 2

	return buf
}

// ReparentWindow sends a checked request.
func ReparentWindow(c *xgb.XConn, Window Window, Parent Window, X int16, Y int16) error {
	var op uint8
	return c.SendRecv(reparentWindowRequest(op, Window, Parent, X, Y), nil)
}

// ReparentWindowUnchecked sends an unchecked request.
func ReparentWindowUnchecked(c *xgb.XConn, Window Window, Parent Window, X int16, Y int16) error {
	var op uint8
	return c.Send(reparentWindowRequest(op, Window, Parent, X, Y))
}

// Write request to wire for ReparentWindow
// reparentWindowRequest writes a ReparentWindow request to a byte slice.
func reparentWindowRequest(opcode uint8, Window Window, Parent Window, X int16, Y int16) []byte {
	size := 16
	b := 0
	buf := make([]byte, size)

	buf[b] = 7 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Parent))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(X))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(Y))
	b += 2

	return buf
}

// RotateProperties sends a checked request.
func RotateProperties(c *xgb.XConn, Window Window, AtomsLen uint16, Delta int16, Atoms []Atom) error {
	var op uint8
	return c.SendRecv(rotatePropertiesRequest(op, Window, AtomsLen, Delta, Atoms), nil)
}

// RotatePropertiesUnchecked sends an unchecked request.
func RotatePropertiesUnchecked(c *xgb.XConn, Window Window, AtomsLen uint16, Delta int16, Atoms []Atom) error {
	var op uint8
	return c.Send(rotatePropertiesRequest(op, Window, AtomsLen, Delta, Atoms))
}

// Write request to wire for RotateProperties
// rotatePropertiesRequest writes a RotateProperties request to a byte slice.
func rotatePropertiesRequest(opcode uint8, Window Window, AtomsLen uint16, Delta int16, Atoms []Atom) []byte {
	size := internal.Pad4((12 + internal.Pad4((int(AtomsLen) * 4))))
	b := 0
	buf := make([]byte, size)

	buf[b] = 114 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], AtomsLen)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(Delta))
	b += 2

	for i := 0; i < int(AtomsLen); i++ {
		binary.LittleEndian.PutUint32(buf[b:], uint32(Atoms[i]))
		b += 4
	}

	return buf
}

// SendEvent sends a checked request.
func SendEvent(c *xgb.XConn, Propagate bool, Destination Window, EventMask uint32, Event string) error {
	var op uint8
	return c.SendRecv(sendEventRequest(op, Propagate, Destination, EventMask, Event), nil)
}

// SendEventUnchecked sends an unchecked request.
func SendEventUnchecked(c *xgb.XConn, Propagate bool, Destination Window, EventMask uint32, Event string) error {
	var op uint8
	return c.Send(sendEventRequest(op, Propagate, Destination, EventMask, Event))
}

// Write request to wire for SendEvent
// sendEventRequest writes a SendEvent request to a byte slice.
func sendEventRequest(opcode uint8, Propagate bool, Destination Window, EventMask uint32, Event string) []byte {
	size := 44
	b := 0
	buf := make([]byte, size)

	buf[b] = 25 // request opcode
	b += 1

	if Propagate {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Destination))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], EventMask)
	b += 4

	copy(buf[b:], Event[:32])
	b += int(32)

	return buf
}

// SetAccessControl sends a checked request.
func SetAccessControl(c *xgb.XConn, Mode byte) error {
	var op uint8
	return c.SendRecv(setAccessControlRequest(op, Mode), nil)
}

// SetAccessControlUnchecked sends an unchecked request.
func SetAccessControlUnchecked(c *xgb.XConn, Mode byte) error {
	var op uint8
	return c.Send(setAccessControlRequest(op, Mode))
}

// Write request to wire for SetAccessControl
// setAccessControlRequest writes a SetAccessControl request to a byte slice.
func setAccessControlRequest(opcode uint8, Mode byte) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = 111 // request opcode
	b += 1

	buf[b] = Mode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// SetClipRectangles sends a checked request.
func SetClipRectangles(c *xgb.XConn, Ordering byte, Gc Gcontext, ClipXOrigin int16, ClipYOrigin int16, Rectangles []Rectangle) error {
	var op uint8
	return c.SendRecv(setClipRectanglesRequest(op, Ordering, Gc, ClipXOrigin, ClipYOrigin, Rectangles), nil)
}

// SetClipRectanglesUnchecked sends an unchecked request.
func SetClipRectanglesUnchecked(c *xgb.XConn, Ordering byte, Gc Gcontext, ClipXOrigin int16, ClipYOrigin int16, Rectangles []Rectangle) error {
	var op uint8
	return c.Send(setClipRectanglesRequest(op, Ordering, Gc, ClipXOrigin, ClipYOrigin, Rectangles))
}

// Write request to wire for SetClipRectangles
// setClipRectanglesRequest writes a SetClipRectangles request to a byte slice.
func setClipRectanglesRequest(opcode uint8, Ordering byte, Gc Gcontext, ClipXOrigin int16, ClipYOrigin int16, Rectangles []Rectangle) []byte {
	size := internal.Pad4((12 + internal.Pad4((len(Rectangles) * 8))))
	b := 0
	buf := make([]byte, size)

	buf[b] = 59 // request opcode
	b += 1

	buf[b] = Ordering
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Gc))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(ClipXOrigin))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(ClipYOrigin))
	b += 2

	b += RectangleListBytes(buf[b:], Rectangles)

	return buf
}

// SetCloseDownMode sends a checked request.
func SetCloseDownMode(c *xgb.XConn, Mode byte) error {
	var op uint8
	return c.SendRecv(setCloseDownModeRequest(op, Mode), nil)
}

// SetCloseDownModeUnchecked sends an unchecked request.
func SetCloseDownModeUnchecked(c *xgb.XConn, Mode byte) error {
	var op uint8
	return c.Send(setCloseDownModeRequest(op, Mode))
}

// Write request to wire for SetCloseDownMode
// setCloseDownModeRequest writes a SetCloseDownMode request to a byte slice.
func setCloseDownModeRequest(opcode uint8, Mode byte) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = 112 // request opcode
	b += 1

	buf[b] = Mode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// SetDashes sends a checked request.
func SetDashes(c *xgb.XConn, Gc Gcontext, DashOffset uint16, DashesLen uint16, Dashes []byte) error {
	var op uint8
	return c.SendRecv(setDashesRequest(op, Gc, DashOffset, DashesLen, Dashes), nil)
}

// SetDashesUnchecked sends an unchecked request.
func SetDashesUnchecked(c *xgb.XConn, Gc Gcontext, DashOffset uint16, DashesLen uint16, Dashes []byte) error {
	var op uint8
	return c.Send(setDashesRequest(op, Gc, DashOffset, DashesLen, Dashes))
}

// Write request to wire for SetDashes
// setDashesRequest writes a SetDashes request to a byte slice.
func setDashesRequest(opcode uint8, Gc Gcontext, DashOffset uint16, DashesLen uint16, Dashes []byte) []byte {
	size := internal.Pad4((12 + internal.Pad4((int(DashesLen) * 1))))
	b := 0
	buf := make([]byte, size)

	buf[b] = 58 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Gc))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], DashOffset)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], DashesLen)
	b += 2

	copy(buf[b:], Dashes[:DashesLen])
	b += int(DashesLen)

	return buf
}

// SetFontPath sends a checked request.
func SetFontPath(c *xgb.XConn, FontQty uint16, Font []Str) error {
	var op uint8
	return c.SendRecv(setFontPathRequest(op, FontQty, Font), nil)
}

// SetFontPathUnchecked sends an unchecked request.
func SetFontPathUnchecked(c *xgb.XConn, FontQty uint16, Font []Str) error {
	var op uint8
	return c.Send(setFontPathRequest(op, FontQty, Font))
}

// Write request to wire for SetFontPath
// setFontPathRequest writes a SetFontPath request to a byte slice.
func setFontPathRequest(opcode uint8, FontQty uint16, Font []Str) []byte {
	size := internal.Pad4((8 + StrListSize(Font)))
	b := 0
	buf := make([]byte, size)

	buf[b] = 51 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], FontQty)
	b += 2

	b += 2 // padding

	b += StrListBytes(buf[b:], Font)

	return buf
}

// SetInputFocus sends a checked request.
func SetInputFocus(c *xgb.XConn, RevertTo byte, Focus Window, Time Timestamp) error {
	var op uint8
	return c.SendRecv(setInputFocusRequest(op, RevertTo, Focus, Time), nil)
}

// SetInputFocusUnchecked sends an unchecked request.
func SetInputFocusUnchecked(c *xgb.XConn, RevertTo byte, Focus Window, Time Timestamp) error {
	var op uint8
	return c.Send(setInputFocusRequest(op, RevertTo, Focus, Time))
}

// Write request to wire for SetInputFocus
// setInputFocusRequest writes a SetInputFocus request to a byte slice.
func setInputFocusRequest(opcode uint8, RevertTo byte, Focus Window, Time Timestamp) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = 42 // request opcode
	b += 1

	buf[b] = RevertTo
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Focus))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Time))
	b += 4

	return buf
}

// SetModifierMapping sends a checked request.
func SetModifierMapping(c *xgb.XConn, KeycodesPerModifier byte, Keycodes []Keycode) (SetModifierMappingReply, error) {
	var reply SetModifierMappingReply
	var op uint8
	err := c.SendRecv(setModifierMappingRequest(op, KeycodesPerModifier, Keycodes), &reply)
	return reply, err
}

// SetModifierMappingUnchecked sends an unchecked request.
func SetModifierMappingUnchecked(c *xgb.XConn, KeycodesPerModifier byte, Keycodes []Keycode) error {
	var op uint8
	return c.Send(setModifierMappingRequest(op, KeycodesPerModifier, Keycodes))
}

// SetModifierMappingReply represents the data returned from a SetModifierMapping request.
type SetModifierMappingReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	Status   byte
}

// Unmarshal reads a byte slice into a SetModifierMappingReply value.
func (v *SetModifierMappingReply) Unmarshal(buf []byte) error {
	if size := 8; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"SetModifierMappingReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	v.Status = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	return nil
}

// Write request to wire for SetModifierMapping
// setModifierMappingRequest writes a SetModifierMapping request to a byte slice.
func setModifierMappingRequest(opcode uint8, KeycodesPerModifier byte, Keycodes []Keycode) []byte {
	size := internal.Pad4((4 + internal.Pad4(((int(KeycodesPerModifier) * 8) * 1))))
	b := 0
	buf := make([]byte, size)

	buf[b] = 118 // request opcode
	b += 1

	buf[b] = KeycodesPerModifier
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	for i := 0; i < int((int(KeycodesPerModifier) * 8)); i++ {
		buf[b] = uint8(Keycodes[i])
		b += 1
	}

	return buf
}

// SetPointerMapping sends a checked request.
func SetPointerMapping(c *xgb.XConn, MapLen byte, Map []byte) (SetPointerMappingReply, error) {
	var reply SetPointerMappingReply
	var op uint8
	err := c.SendRecv(setPointerMappingRequest(op, MapLen, Map), &reply)
	return reply, err
}

// SetPointerMappingUnchecked sends an unchecked request.
func SetPointerMappingUnchecked(c *xgb.XConn, MapLen byte, Map []byte) error {
	var op uint8
	return c.Send(setPointerMappingRequest(op, MapLen, Map))
}

// SetPointerMappingReply represents the data returned from a SetPointerMapping request.
type SetPointerMappingReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	Status   byte
}

// Unmarshal reads a byte slice into a SetPointerMappingReply value.
func (v *SetPointerMappingReply) Unmarshal(buf []byte) error {
	if size := 8; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"SetPointerMappingReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	v.Status = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	return nil
}

// Write request to wire for SetPointerMapping
// setPointerMappingRequest writes a SetPointerMapping request to a byte slice.
func setPointerMappingRequest(opcode uint8, MapLen byte, Map []byte) []byte {
	size := internal.Pad4((4 + internal.Pad4((int(MapLen) * 1))))
	b := 0
	buf := make([]byte, size)

	buf[b] = 116 // request opcode
	b += 1

	buf[b] = MapLen
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	copy(buf[b:], Map[:MapLen])
	b += int(MapLen)

	return buf
}

// SetScreenSaver sends a checked request.
func SetScreenSaver(c *xgb.XConn, Timeout int16, Interval int16, PreferBlanking byte, AllowExposures byte) error {
	var op uint8
	return c.SendRecv(setScreenSaverRequest(op, Timeout, Interval, PreferBlanking, AllowExposures), nil)
}

// SetScreenSaverUnchecked sends an unchecked request.
func SetScreenSaverUnchecked(c *xgb.XConn, Timeout int16, Interval int16, PreferBlanking byte, AllowExposures byte) error {
	var op uint8
	return c.Send(setScreenSaverRequest(op, Timeout, Interval, PreferBlanking, AllowExposures))
}

// Write request to wire for SetScreenSaver
// setScreenSaverRequest writes a SetScreenSaver request to a byte slice.
func setScreenSaverRequest(opcode uint8, Timeout int16, Interval int16, PreferBlanking byte, AllowExposures byte) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = 107 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(Timeout))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(Interval))
	b += 2

	buf[b] = PreferBlanking
	b += 1

	buf[b] = AllowExposures
	b += 1

	return buf
}

// SetSelectionOwner sends a checked request.
func SetSelectionOwner(c *xgb.XConn, Owner Window, Selection Atom, Time Timestamp) error {
	var op uint8
	return c.SendRecv(setSelectionOwnerRequest(op, Owner, Selection, Time), nil)
}

// SetSelectionOwnerUnchecked sends an unchecked request.
func SetSelectionOwnerUnchecked(c *xgb.XConn, Owner Window, Selection Atom, Time Timestamp) error {
	var op uint8
	return c.Send(setSelectionOwnerRequest(op, Owner, Selection, Time))
}

// Write request to wire for SetSelectionOwner
// setSelectionOwnerRequest writes a SetSelectionOwner request to a byte slice.
func setSelectionOwnerRequest(opcode uint8, Owner Window, Selection Atom, Time Timestamp) []byte {
	size := 16
	b := 0
	buf := make([]byte, size)

	buf[b] = 22 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Owner))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Selection))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Time))
	b += 4

	return buf
}

// StoreColors sends a checked request.
func StoreColors(c *xgb.XConn, Cmap Colormap, Items []Coloritem) error {
	var op uint8
	return c.SendRecv(storeColorsRequest(op, Cmap, Items), nil)
}

// StoreColorsUnchecked sends an unchecked request.
func StoreColorsUnchecked(c *xgb.XConn, Cmap Colormap, Items []Coloritem) error {
	var op uint8
	return c.Send(storeColorsRequest(op, Cmap, Items))
}

// Write request to wire for StoreColors
// storeColorsRequest writes a StoreColors request to a byte slice.
func storeColorsRequest(opcode uint8, Cmap Colormap, Items []Coloritem) []byte {
	size := internal.Pad4((8 + internal.Pad4((len(Items) * 12))))
	b := 0
	buf := make([]byte, size)

	buf[b] = 89 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Cmap))
	b += 4

	b += ColoritemListBytes(buf[b:], Items)

	return buf
}

// StoreNamedColor sends a checked request.
func StoreNamedColor(c *xgb.XConn, Flags byte, Cmap Colormap, Pixel uint32, NameLen uint16, Name string) error {
	var op uint8
	return c.SendRecv(storeNamedColorRequest(op, Flags, Cmap, Pixel, NameLen, Name), nil)
}

// StoreNamedColorUnchecked sends an unchecked request.
func StoreNamedColorUnchecked(c *xgb.XConn, Flags byte, Cmap Colormap, Pixel uint32, NameLen uint16, Name string) error {
	var op uint8
	return c.Send(storeNamedColorRequest(op, Flags, Cmap, Pixel, NameLen, Name))
}

// Write request to wire for StoreNamedColor
// storeNamedColorRequest writes a StoreNamedColor request to a byte slice.
func storeNamedColorRequest(opcode uint8, Flags byte, Cmap Colormap, Pixel uint32, NameLen uint16, Name string) []byte {
	size := internal.Pad4((16 + internal.Pad4((int(NameLen) * 1))))
	b := 0
	buf := make([]byte, size)

	buf[b] = 90 // request opcode
	b += 1

	buf[b] = Flags
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Cmap))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], Pixel)
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], NameLen)
	b += 2

	b += 2 // padding

	copy(buf[b:], Name[:NameLen])
	b += int(NameLen)

	return buf
}

// TranslateCoordinates sends a checked request.
func TranslateCoordinates(c *xgb.XConn, SrcWindow Window, DstWindow Window, SrcX int16, SrcY int16) (TranslateCoordinatesReply, error) {
	var reply TranslateCoordinatesReply
	var op uint8
	err := c.SendRecv(translateCoordinatesRequest(op, SrcWindow, DstWindow, SrcX, SrcY), &reply)
	return reply, err
}

// TranslateCoordinatesUnchecked sends an unchecked request.
func TranslateCoordinatesUnchecked(c *xgb.XConn, SrcWindow Window, DstWindow Window, SrcX int16, SrcY int16) error {
	var op uint8
	return c.Send(translateCoordinatesRequest(op, SrcWindow, DstWindow, SrcX, SrcY))
}

// TranslateCoordinatesReply represents the data returned from a TranslateCoordinates request.
type TranslateCoordinatesReply struct {
	Sequence   uint16 // sequence number of the request for this reply
	Length     uint32 // number of bytes in this reply
	SameScreen bool
	Child      Window
	DstX       int16
	DstY       int16
}

// Unmarshal reads a byte slice into a TranslateCoordinatesReply value.
func (v *TranslateCoordinatesReply) Unmarshal(buf []byte) error {
	if size := 16; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"TranslateCoordinatesReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	v.SameScreen = (buf[b] == 1)
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Child = Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.DstX = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.DstY = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	return nil
}

// Write request to wire for TranslateCoordinates
// translateCoordinatesRequest writes a TranslateCoordinates request to a byte slice.
func translateCoordinatesRequest(opcode uint8, SrcWindow Window, DstWindow Window, SrcX int16, SrcY int16) []byte {
	size := 16
	b := 0
	buf := make([]byte, size)

	buf[b] = 40 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(SrcWindow))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(DstWindow))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(SrcX))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(SrcY))
	b += 2

	return buf
}

// UngrabButton sends a checked request.
func UngrabButton(c *xgb.XConn, Button byte, GrabWindow Window, Modifiers uint16) error {
	var op uint8
	return c.SendRecv(ungrabButtonRequest(op, Button, GrabWindow, Modifiers), nil)
}

// UngrabButtonUnchecked sends an unchecked request.
func UngrabButtonUnchecked(c *xgb.XConn, Button byte, GrabWindow Window, Modifiers uint16) error {
	var op uint8
	return c.Send(ungrabButtonRequest(op, Button, GrabWindow, Modifiers))
}

// Write request to wire for UngrabButton
// ungrabButtonRequest writes a UngrabButton request to a byte slice.
func ungrabButtonRequest(opcode uint8, Button byte, GrabWindow Window, Modifiers uint16) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = 29 // request opcode
	b += 1

	buf[b] = Button
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(GrabWindow))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], Modifiers)
	b += 2

	b += 2 // padding

	return buf
}

// UngrabKey sends a checked request.
func UngrabKey(c *xgb.XConn, Key Keycode, GrabWindow Window, Modifiers uint16) error {
	var op uint8
	return c.SendRecv(ungrabKeyRequest(op, Key, GrabWindow, Modifiers), nil)
}

// UngrabKeyUnchecked sends an unchecked request.
func UngrabKeyUnchecked(c *xgb.XConn, Key Keycode, GrabWindow Window, Modifiers uint16) error {
	var op uint8
	return c.Send(ungrabKeyRequest(op, Key, GrabWindow, Modifiers))
}

// Write request to wire for UngrabKey
// ungrabKeyRequest writes a UngrabKey request to a byte slice.
func ungrabKeyRequest(opcode uint8, Key Keycode, GrabWindow Window, Modifiers uint16) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = 34 // request opcode
	b += 1

	buf[b] = uint8(Key)
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(GrabWindow))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], Modifiers)
	b += 2

	b += 2 // padding

	return buf
}

// UngrabKeyboard sends a checked request.
func UngrabKeyboard(c *xgb.XConn, Time Timestamp) error {
	var op uint8
	return c.SendRecv(ungrabKeyboardRequest(op, Time), nil)
}

// UngrabKeyboardUnchecked sends an unchecked request.
func UngrabKeyboardUnchecked(c *xgb.XConn, Time Timestamp) error {
	var op uint8
	return c.Send(ungrabKeyboardRequest(op, Time))
}

// Write request to wire for UngrabKeyboard
// ungrabKeyboardRequest writes a UngrabKeyboard request to a byte slice.
func ungrabKeyboardRequest(opcode uint8, Time Timestamp) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = 32 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Time))
	b += 4

	return buf
}

// UngrabPointer sends a checked request.
func UngrabPointer(c *xgb.XConn, Time Timestamp) error {
	var op uint8
	return c.SendRecv(ungrabPointerRequest(op, Time), nil)
}

// UngrabPointerUnchecked sends an unchecked request.
func UngrabPointerUnchecked(c *xgb.XConn, Time Timestamp) error {
	var op uint8
	return c.Send(ungrabPointerRequest(op, Time))
}

// Write request to wire for UngrabPointer
// ungrabPointerRequest writes a UngrabPointer request to a byte slice.
func ungrabPointerRequest(opcode uint8, Time Timestamp) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = 27 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Time))
	b += 4

	return buf
}

// UngrabServer sends a checked request.
func UngrabServer(c *xgb.XConn) error {
	var op uint8
	return c.SendRecv(ungrabServerRequest(op), nil)
}

// UngrabServerUnchecked sends an unchecked request.
func UngrabServerUnchecked(c *xgb.XConn) error {
	var op uint8
	return c.Send(ungrabServerRequest(op))
}

// Write request to wire for UngrabServer
// ungrabServerRequest writes a UngrabServer request to a byte slice.
func ungrabServerRequest(opcode uint8) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = 37 // request opcode
	b += 1

	b += 1                                                 // padding
	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// UninstallColormap sends a checked request.
func UninstallColormap(c *xgb.XConn, Cmap Colormap) error {
	var op uint8
	return c.SendRecv(uninstallColormapRequest(op, Cmap), nil)
}

// UninstallColormapUnchecked sends an unchecked request.
func UninstallColormapUnchecked(c *xgb.XConn, Cmap Colormap) error {
	var op uint8
	return c.Send(uninstallColormapRequest(op, Cmap))
}

// Write request to wire for UninstallColormap
// uninstallColormapRequest writes a UninstallColormap request to a byte slice.
func uninstallColormapRequest(opcode uint8, Cmap Colormap) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = 82 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Cmap))
	b += 4

	return buf
}

// UnmapSubwindows sends a checked request.
func UnmapSubwindows(c *xgb.XConn, Window Window) error {
	var op uint8
	return c.SendRecv(unmapSubwindowsRequest(op, Window), nil)
}

// UnmapSubwindowsUnchecked sends an unchecked request.
func UnmapSubwindowsUnchecked(c *xgb.XConn, Window Window) error {
	var op uint8
	return c.Send(unmapSubwindowsRequest(op, Window))
}

// Write request to wire for UnmapSubwindows
// unmapSubwindowsRequest writes a UnmapSubwindows request to a byte slice.
func unmapSubwindowsRequest(opcode uint8, Window Window) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = 11 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	return buf
}

// UnmapWindow sends a checked request.
func UnmapWindow(c *xgb.XConn, Window Window) error {
	var op uint8
	return c.SendRecv(unmapWindowRequest(op, Window), nil)
}

// UnmapWindowUnchecked sends an unchecked request.
func UnmapWindowUnchecked(c *xgb.XConn, Window Window) error {
	var op uint8
	return c.Send(unmapWindowRequest(op, Window))
}

// Write request to wire for UnmapWindow
// unmapWindowRequest writes a UnmapWindow request to a byte slice.
func unmapWindowRequest(opcode uint8, Window Window) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = 10 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	return buf
}

// WarpPointer sends a checked request.
func WarpPointer(c *xgb.XConn, SrcWindow Window, DstWindow Window, SrcX int16, SrcY int16, SrcWidth uint16, SrcHeight uint16, DstX int16, DstY int16) error {
	var op uint8
	return c.SendRecv(warpPointerRequest(op, SrcWindow, DstWindow, SrcX, SrcY, SrcWidth, SrcHeight, DstX, DstY), nil)
}

// WarpPointerUnchecked sends an unchecked request.
func WarpPointerUnchecked(c *xgb.XConn, SrcWindow Window, DstWindow Window, SrcX int16, SrcY int16, SrcWidth uint16, SrcHeight uint16, DstX int16, DstY int16) error {
	var op uint8
	return c.Send(warpPointerRequest(op, SrcWindow, DstWindow, SrcX, SrcY, SrcWidth, SrcHeight, DstX, DstY))
}

// Write request to wire for WarpPointer
// warpPointerRequest writes a WarpPointer request to a byte slice.
func warpPointerRequest(opcode uint8, SrcWindow Window, DstWindow Window, SrcX int16, SrcY int16, SrcWidth uint16, SrcHeight uint16, DstX int16, DstY int16) []byte {
	size := 24
	b := 0
	buf := make([]byte, size)

	buf[b] = 41 // request opcode
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(SrcWindow))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(DstWindow))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(SrcX))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(SrcY))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], SrcWidth)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], SrcHeight)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(DstX))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(DstY))
	b += 2

	return buf
}
