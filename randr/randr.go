// FILE GENERATED AUTOMATICALLY FROM "randr.xml"
package randr

import (
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/probakowski/go-xgb"
	"github.com/probakowski/go-xgb/internal"
	"github.com/probakowski/go-xgb/render"
	"github.com/probakowski/go-xgb/xproto"
)

const (
	// ExtName is the user-friendly name string of this X extension.
	ExtName = "RandR"

	// ExtXName is the name string this extension is known by to the X server.
	ExtXName = "RANDR"
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

// Register will query the X server for RandR extension support, and register relevant extension unmarshalers with the XConn.
func Register(xconn *xgb.XConn) error {
	// Query the X server for this extension
	reply, err := xproto.QueryExtension(xconn, uint16(len(ExtXName)), ExtXName)
	if err != nil {
		return fmt.Errorf("error querying X for \"RandR\": %w", err)
	} else if !reply.Present {
		return fmt.Errorf("no extension named \"RandR\" is known to the X server: reply=%+v", reply)
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

// BadBadCrtc is the error number for a BadBadCrtc.
const BadBadCrtc = 1

type BadCrtcError struct {
	Sequence uint16
	NiceName string
}

// UnmarshalBadCrtcError constructs a BadCrtcError value that implements xgb.Error from a byte slice.
func UnmarshalBadCrtcError(buf []byte) (xgb.XError, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"BadCrtcError\"", len(buf))
	}

	v := &BadCrtcError{}
	v.NiceName = "BadCrtc"

	b := 1 // skip error determinant
	b += 1 // don't read error number

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return v, nil
}

// SeqID returns the sequence id attached to the BadBadCrtc error.
// This is mostly used internally.
func (err *BadCrtcError) SeqID() uint16 {
	return err.Sequence
}

// BadID returns the 'BadValue' number if one exists for the BadBadCrtc error. If no bad value exists, 0 is returned.
func (err *BadCrtcError) BadID() uint32 {
	return 0
}

// Error returns a rudimentary string representation of the BadBadCrtc error.
func (err *BadCrtcError) Error() string {
	var buf strings.Builder

	buf.WriteString("BadBadCrtc{")
	buf.WriteString("NiceName: " + err.NiceName)
	buf.WriteByte(' ')
	buf.WriteString("Sequence: " + strconv.FormatUint(uint64(err.Sequence), 10))

	buf.WriteByte('}')

	return buf.String()
}

func init() { registerError(1, UnmarshalBadCrtcError) }

// BadBadMode is the error number for a BadBadMode.
const BadBadMode = 2

type BadModeError struct {
	Sequence uint16
	NiceName string
}

// UnmarshalBadModeError constructs a BadModeError value that implements xgb.Error from a byte slice.
func UnmarshalBadModeError(buf []byte) (xgb.XError, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"BadModeError\"", len(buf))
	}

	v := &BadModeError{}
	v.NiceName = "BadMode"

	b := 1 // skip error determinant
	b += 1 // don't read error number

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return v, nil
}

// SeqID returns the sequence id attached to the BadBadMode error.
// This is mostly used internally.
func (err *BadModeError) SeqID() uint16 {
	return err.Sequence
}

// BadID returns the 'BadValue' number if one exists for the BadBadMode error. If no bad value exists, 0 is returned.
func (err *BadModeError) BadID() uint32 {
	return 0
}

// Error returns a rudimentary string representation of the BadBadMode error.
func (err *BadModeError) Error() string {
	var buf strings.Builder

	buf.WriteString("BadBadMode{")
	buf.WriteString("NiceName: " + err.NiceName)
	buf.WriteByte(' ')
	buf.WriteString("Sequence: " + strconv.FormatUint(uint64(err.Sequence), 10))

	buf.WriteByte('}')

	return buf.String()
}

func init() { registerError(2, UnmarshalBadModeError) }

// BadBadOutput is the error number for a BadBadOutput.
const BadBadOutput = 0

type BadOutputError struct {
	Sequence uint16
	NiceName string
}

// UnmarshalBadOutputError constructs a BadOutputError value that implements xgb.Error from a byte slice.
func UnmarshalBadOutputError(buf []byte) (xgb.XError, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"BadOutputError\"", len(buf))
	}

	v := &BadOutputError{}
	v.NiceName = "BadOutput"

	b := 1 // skip error determinant
	b += 1 // don't read error number

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return v, nil
}

// SeqID returns the sequence id attached to the BadBadOutput error.
// This is mostly used internally.
func (err *BadOutputError) SeqID() uint16 {
	return err.Sequence
}

// BadID returns the 'BadValue' number if one exists for the BadBadOutput error. If no bad value exists, 0 is returned.
func (err *BadOutputError) BadID() uint32 {
	return 0
}

// Error returns a rudimentary string representation of the BadBadOutput error.
func (err *BadOutputError) Error() string {
	var buf strings.Builder

	buf.WriteString("BadBadOutput{")
	buf.WriteString("NiceName: " + err.NiceName)
	buf.WriteByte(' ')
	buf.WriteString("Sequence: " + strconv.FormatUint(uint64(err.Sequence), 10))

	buf.WriteByte('}')

	return buf.String()
}

func init() { registerError(0, UnmarshalBadOutputError) }

// BadBadProvider is the error number for a BadBadProvider.
const BadBadProvider = 3

type BadProviderError struct {
	Sequence uint16
	NiceName string
}

// UnmarshalBadProviderError constructs a BadProviderError value that implements xgb.Error from a byte slice.
func UnmarshalBadProviderError(buf []byte) (xgb.XError, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"BadProviderError\"", len(buf))
	}

	v := &BadProviderError{}
	v.NiceName = "BadProvider"

	b := 1 // skip error determinant
	b += 1 // don't read error number

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return v, nil
}

// SeqID returns the sequence id attached to the BadBadProvider error.
// This is mostly used internally.
func (err *BadProviderError) SeqID() uint16 {
	return err.Sequence
}

// BadID returns the 'BadValue' number if one exists for the BadBadProvider error. If no bad value exists, 0 is returned.
func (err *BadProviderError) BadID() uint32 {
	return 0
}

// Error returns a rudimentary string representation of the BadBadProvider error.
func (err *BadProviderError) Error() string {
	var buf strings.Builder

	buf.WriteString("BadBadProvider{")
	buf.WriteString("NiceName: " + err.NiceName)
	buf.WriteByte(' ')
	buf.WriteString("Sequence: " + strconv.FormatUint(uint64(err.Sequence), 10))

	buf.WriteByte('}')

	return buf.String()
}

func init() { registerError(3, UnmarshalBadProviderError) }

const (
	ConnectionConnected    = 0
	ConnectionDisconnected = 1
	ConnectionUnknown      = 2
)

type Crtc uint32

func NewCrtcID(c *xgb.XConn) Crtc {
	id := c.NewXID()
	return Crtc(id)
}

type CrtcChange struct {
	Timestamp xproto.Timestamp
	Window    xproto.Window
	Crtc      Crtc
	Mode      Mode
	Rotation  uint16
	// padding: 2 bytes
	X      int16
	Y      int16
	Width  uint16
	Height uint16
}

// CrtcChangeRead reads a byte slice into a CrtcChange value.
func CrtcChangeRead(buf []byte, v *CrtcChange) int {
	b := 0

	v.Timestamp = xproto.Timestamp(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Window = xproto.Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Crtc = Crtc(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Mode = Mode(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Rotation = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	b += 2 // padding

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

// CrtcChangeReadList reads a byte slice into a list of CrtcChange values.
func CrtcChangeReadList(buf []byte, dest []CrtcChange) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = CrtcChange{}
		b += CrtcChangeRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a CrtcChange value to a byte slice.
func (v CrtcChange) Bytes() []byte {
	buf := make([]byte, 28)
	b := 0

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Timestamp))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Window))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Crtc))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Mode))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], v.Rotation)
	b += 2

	b += 2 // padding

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

// CrtcChangeListBytes writes a list of CrtcChange values to a byte slice.
func CrtcChangeListBytes(buf []byte, list []CrtcChange) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

type Lease uint32

func NewLeaseID(c *xgb.XConn) Lease {
	id := c.NewXID()
	return Lease(id)
}

type LeaseNotify struct {
	Timestamp xproto.Timestamp
	Window    xproto.Window
	Lease     Lease
	Created   byte
	// padding: 15 bytes
}

// LeaseNotifyRead reads a byte slice into a LeaseNotify value.
func LeaseNotifyRead(buf []byte, v *LeaseNotify) int {
	b := 0

	v.Timestamp = xproto.Timestamp(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Window = xproto.Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Lease = Lease(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Created = buf[b]
	b += 1

	b += 15 // padding

	return b
}

// LeaseNotifyReadList reads a byte slice into a list of LeaseNotify values.
func LeaseNotifyReadList(buf []byte, dest []LeaseNotify) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = LeaseNotify{}
		b += LeaseNotifyRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a LeaseNotify value to a byte slice.
func (v LeaseNotify) Bytes() []byte {
	buf := make([]byte, 28)
	b := 0

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Timestamp))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Window))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Lease))
	b += 4

	buf[b] = v.Created
	b += 1

	b += 15 // padding

	return buf[:b]
}

// LeaseNotifyListBytes writes a list of LeaseNotify values to a byte slice.
func LeaseNotifyListBytes(buf []byte, list []LeaseNotify) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

type Mode uint32

func NewModeID(c *xgb.XConn) Mode {
	id := c.NewXID()
	return Mode(id)
}

const (
	ModeFlagHsyncPositive  = 1
	ModeFlagHsyncNegative  = 2
	ModeFlagVsyncPositive  = 4
	ModeFlagVsyncNegative  = 8
	ModeFlagInterlace      = 16
	ModeFlagDoubleScan     = 32
	ModeFlagCsync          = 64
	ModeFlagCsyncPositive  = 128
	ModeFlagCsyncNegative  = 256
	ModeFlagHskewPresent   = 512
	ModeFlagBcast          = 1024
	ModeFlagPixelMultiplex = 2048
	ModeFlagDoubleClock    = 4096
	ModeFlagHalveClock     = 8192
)

type ModeInfo struct {
	Id         uint32
	Width      uint16
	Height     uint16
	DotClock   uint32
	HsyncStart uint16
	HsyncEnd   uint16
	Htotal     uint16
	Hskew      uint16
	VsyncStart uint16
	VsyncEnd   uint16
	Vtotal     uint16
	NameLen    uint16
	ModeFlags  uint32
}

// ModeInfoRead reads a byte slice into a ModeInfo value.
func ModeInfoRead(buf []byte, v *ModeInfo) int {
	b := 0

	v.Id = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.Width = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Height = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.DotClock = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.HsyncStart = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.HsyncEnd = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Htotal = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Hskew = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.VsyncStart = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.VsyncEnd = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Vtotal = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.NameLen = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.ModeFlags = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	return b
}

// ModeInfoReadList reads a byte slice into a list of ModeInfo values.
func ModeInfoReadList(buf []byte, dest []ModeInfo) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = ModeInfo{}
		b += ModeInfoRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a ModeInfo value to a byte slice.
func (v ModeInfo) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	binary.LittleEndian.PutUint32(buf[b:], v.Id)
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], v.Width)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Height)
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], v.DotClock)
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], v.HsyncStart)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.HsyncEnd)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Htotal)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Hskew)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.VsyncStart)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.VsyncEnd)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Vtotal)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.NameLen)
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], v.ModeFlags)
	b += 4

	return buf[:b]
}

// ModeInfoListBytes writes a list of ModeInfo values to a byte slice.
func ModeInfoListBytes(buf []byte, list []ModeInfo) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

type MonitorInfo struct {
	Name                xproto.Atom
	Primary             bool
	Automatic           bool
	NOutput             uint16
	X                   int16
	Y                   int16
	Width               uint16
	Height              uint16
	WidthInMillimeters  uint32
	HeightInMillimeters uint32
	Outputs             []Output // size: internal.Pad4((int(NOutput) * 4))
}

// MonitorInfoRead reads a byte slice into a MonitorInfo value.
func MonitorInfoRead(buf []byte, v *MonitorInfo) int {
	b := 0

	v.Name = xproto.Atom(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Primary = (buf[b] == 1)
	b += 1

	v.Automatic = (buf[b] == 1)
	b += 1

	v.NOutput = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.X = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.Y = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.Width = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Height = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.WidthInMillimeters = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.HeightInMillimeters = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.Outputs = make([]Output, v.NOutput)
	for i := 0; i < int(v.NOutput); i++ {
		v.Outputs[i] = Output(binary.LittleEndian.Uint32(buf[b:]))
		b += 4
	}

	return b
}

// MonitorInfoReadList reads a byte slice into a list of MonitorInfo values.
func MonitorInfoReadList(buf []byte, dest []MonitorInfo) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = MonitorInfo{}
		b += MonitorInfoRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a MonitorInfo value to a byte slice.
func (v MonitorInfo) Bytes() []byte {
	buf := make([]byte, (24 + internal.Pad4((int(v.NOutput) * 4))))
	b := 0

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Name))
	b += 4

	if v.Primary {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	if v.Automatic {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], v.NOutput)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.X))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.Y))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Width)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Height)
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], v.WidthInMillimeters)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], v.HeightInMillimeters)
	b += 4

	for i := 0; i < int(v.NOutput); i++ {
		binary.LittleEndian.PutUint32(buf[b:], uint32(v.Outputs[i]))
		b += 4
	}

	return buf[:b]
}

// MonitorInfoListBytes writes a list of MonitorInfo values to a byte slice.
func MonitorInfoListBytes(buf []byte, list []MonitorInfo) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

// MonitorInfoListSize computes the size (bytes) of a list of MonitorInfo values.
func MonitorInfoListSize(list []MonitorInfo) int {
	size := 0
	for _, item := range list {
		size += (24 + internal.Pad4((int(item.NOutput) * 4)))
	}
	return size
}

const (
	NotifyCrtcChange       = 0
	NotifyOutputChange     = 1
	NotifyOutputProperty   = 2
	NotifyProviderChange   = 3
	NotifyProviderProperty = 4
	NotifyResourceChange   = 5
	NotifyLease            = 6
)

// Notify is the event number for a NotifyEvent.
const Notify = 1

type NotifyEvent struct {
	Sequence uint16
	SubCode  byte
	U        NotifyDataUnion
}

// UnmarshalNotifyEvent constructs a NotifyEvent value that implements xgb.Event from a byte slice.
func UnmarshalNotifyEvent(buf []byte) (xgb.XEvent, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"NotifyEvent\"", len(buf))
	}

	v := &NotifyEvent{}
	b := 1 // don't read event number

	v.SubCode = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.U = NotifyDataUnion{}
	b += NotifyDataUnionRead(buf[b:], &v.U)

	return v, nil
}

// Bytes writes a NotifyEvent value to a byte slice.
func (v *NotifyEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 1
	b += 1

	buf[b] = v.SubCode
	b += 1

	b += 2 // skip sequence number

	{
		unionBytes := v.U.Bytes()
		copy(buf[b:], unionBytes)
		b += len(unionBytes)
	}

	return buf
}

// SeqID returns the sequence id attached to the Notify event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v *NotifyEvent) SeqID() uint16 {
	return v.Sequence
}

func init() { registerEvent(1, UnmarshalNotifyEvent) }

// NotifyDataUnion is a representation of the NotifyDataUnion union type.
// Note that to *create* a Union, you should *never* create
// this struct directly (unless you know what you're doing).
// Instead use one of the following constructors for 'NotifyDataUnion':
//
//	NotifyDataUnionCcNew(Cc CrtcChange) NotifyDataUnion
//	NotifyDataUnionOcNew(Oc OutputChange) NotifyDataUnion
//	NotifyDataUnionOpNew(Op OutputProperty) NotifyDataUnion
//	NotifyDataUnionPcNew(Pc ProviderChange) NotifyDataUnion
//	NotifyDataUnionPpNew(Pp ProviderProperty) NotifyDataUnion
//	NotifyDataUnionRcNew(Rc ResourceChange) NotifyDataUnion
//	NotifyDataUnionLcNew(Lc LeaseNotify) NotifyDataUnion
type NotifyDataUnion struct {
	Cc CrtcChange
	Oc OutputChange
	Op OutputProperty
	Pc ProviderChange
	Pp ProviderProperty
	Rc ResourceChange
	Lc LeaseNotify
}

// NotifyDataUnionCcNew constructs a new NotifyDataUnion union type with the Cc field.
func NotifyDataUnionCcNew(Cc CrtcChange) NotifyDataUnion {
	var b int
	buf := make([]byte, 28)

	{
		structBytes := Cc.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	// Create the Union type
	v := NotifyDataUnion{}

	// Now copy buf into all fields

	b = 0 // always read the same bytes
	v.Cc = CrtcChange{}
	b += CrtcChangeRead(buf[b:], &v.Cc)

	b = 0 // always read the same bytes
	v.Oc = OutputChange{}
	b += OutputChangeRead(buf[b:], &v.Oc)

	b = 0 // always read the same bytes
	v.Op = OutputProperty{}
	b += OutputPropertyRead(buf[b:], &v.Op)

	b = 0 // always read the same bytes
	v.Pc = ProviderChange{}
	b += ProviderChangeRead(buf[b:], &v.Pc)

	b = 0 // always read the same bytes
	v.Pp = ProviderProperty{}
	b += ProviderPropertyRead(buf[b:], &v.Pp)

	b = 0 // always read the same bytes
	v.Rc = ResourceChange{}
	b += ResourceChangeRead(buf[b:], &v.Rc)

	b = 0 // always read the same bytes
	v.Lc = LeaseNotify{}
	b += LeaseNotifyRead(buf[b:], &v.Lc)

	return v
}

// NotifyDataUnionOcNew constructs a new NotifyDataUnion union type with the Oc field.
func NotifyDataUnionOcNew(Oc OutputChange) NotifyDataUnion {
	var b int
	buf := make([]byte, 28)

	{
		structBytes := Oc.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	// Create the Union type
	v := NotifyDataUnion{}

	// Now copy buf into all fields

	b = 0 // always read the same bytes
	v.Cc = CrtcChange{}
	b += CrtcChangeRead(buf[b:], &v.Cc)

	b = 0 // always read the same bytes
	v.Oc = OutputChange{}
	b += OutputChangeRead(buf[b:], &v.Oc)

	b = 0 // always read the same bytes
	v.Op = OutputProperty{}
	b += OutputPropertyRead(buf[b:], &v.Op)

	b = 0 // always read the same bytes
	v.Pc = ProviderChange{}
	b += ProviderChangeRead(buf[b:], &v.Pc)

	b = 0 // always read the same bytes
	v.Pp = ProviderProperty{}
	b += ProviderPropertyRead(buf[b:], &v.Pp)

	b = 0 // always read the same bytes
	v.Rc = ResourceChange{}
	b += ResourceChangeRead(buf[b:], &v.Rc)

	b = 0 // always read the same bytes
	v.Lc = LeaseNotify{}
	b += LeaseNotifyRead(buf[b:], &v.Lc)

	return v
}

// NotifyDataUnionOpNew constructs a new NotifyDataUnion union type with the Op field.
func NotifyDataUnionOpNew(Op OutputProperty) NotifyDataUnion {
	var b int
	buf := make([]byte, 28)

	{
		structBytes := Op.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	// Create the Union type
	v := NotifyDataUnion{}

	// Now copy buf into all fields

	b = 0 // always read the same bytes
	v.Cc = CrtcChange{}
	b += CrtcChangeRead(buf[b:], &v.Cc)

	b = 0 // always read the same bytes
	v.Oc = OutputChange{}
	b += OutputChangeRead(buf[b:], &v.Oc)

	b = 0 // always read the same bytes
	v.Op = OutputProperty{}
	b += OutputPropertyRead(buf[b:], &v.Op)

	b = 0 // always read the same bytes
	v.Pc = ProviderChange{}
	b += ProviderChangeRead(buf[b:], &v.Pc)

	b = 0 // always read the same bytes
	v.Pp = ProviderProperty{}
	b += ProviderPropertyRead(buf[b:], &v.Pp)

	b = 0 // always read the same bytes
	v.Rc = ResourceChange{}
	b += ResourceChangeRead(buf[b:], &v.Rc)

	b = 0 // always read the same bytes
	v.Lc = LeaseNotify{}
	b += LeaseNotifyRead(buf[b:], &v.Lc)

	return v
}

// NotifyDataUnionPcNew constructs a new NotifyDataUnion union type with the Pc field.
func NotifyDataUnionPcNew(Pc ProviderChange) NotifyDataUnion {
	var b int
	buf := make([]byte, 28)

	{
		structBytes := Pc.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	// Create the Union type
	v := NotifyDataUnion{}

	// Now copy buf into all fields

	b = 0 // always read the same bytes
	v.Cc = CrtcChange{}
	b += CrtcChangeRead(buf[b:], &v.Cc)

	b = 0 // always read the same bytes
	v.Oc = OutputChange{}
	b += OutputChangeRead(buf[b:], &v.Oc)

	b = 0 // always read the same bytes
	v.Op = OutputProperty{}
	b += OutputPropertyRead(buf[b:], &v.Op)

	b = 0 // always read the same bytes
	v.Pc = ProviderChange{}
	b += ProviderChangeRead(buf[b:], &v.Pc)

	b = 0 // always read the same bytes
	v.Pp = ProviderProperty{}
	b += ProviderPropertyRead(buf[b:], &v.Pp)

	b = 0 // always read the same bytes
	v.Rc = ResourceChange{}
	b += ResourceChangeRead(buf[b:], &v.Rc)

	b = 0 // always read the same bytes
	v.Lc = LeaseNotify{}
	b += LeaseNotifyRead(buf[b:], &v.Lc)

	return v
}

// NotifyDataUnionPpNew constructs a new NotifyDataUnion union type with the Pp field.
func NotifyDataUnionPpNew(Pp ProviderProperty) NotifyDataUnion {
	var b int
	buf := make([]byte, 28)

	{
		structBytes := Pp.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	// Create the Union type
	v := NotifyDataUnion{}

	// Now copy buf into all fields

	b = 0 // always read the same bytes
	v.Cc = CrtcChange{}
	b += CrtcChangeRead(buf[b:], &v.Cc)

	b = 0 // always read the same bytes
	v.Oc = OutputChange{}
	b += OutputChangeRead(buf[b:], &v.Oc)

	b = 0 // always read the same bytes
	v.Op = OutputProperty{}
	b += OutputPropertyRead(buf[b:], &v.Op)

	b = 0 // always read the same bytes
	v.Pc = ProviderChange{}
	b += ProviderChangeRead(buf[b:], &v.Pc)

	b = 0 // always read the same bytes
	v.Pp = ProviderProperty{}
	b += ProviderPropertyRead(buf[b:], &v.Pp)

	b = 0 // always read the same bytes
	v.Rc = ResourceChange{}
	b += ResourceChangeRead(buf[b:], &v.Rc)

	b = 0 // always read the same bytes
	v.Lc = LeaseNotify{}
	b += LeaseNotifyRead(buf[b:], &v.Lc)

	return v
}

// NotifyDataUnionRcNew constructs a new NotifyDataUnion union type with the Rc field.
func NotifyDataUnionRcNew(Rc ResourceChange) NotifyDataUnion {
	var b int
	buf := make([]byte, 28)

	{
		structBytes := Rc.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	// Create the Union type
	v := NotifyDataUnion{}

	// Now copy buf into all fields

	b = 0 // always read the same bytes
	v.Cc = CrtcChange{}
	b += CrtcChangeRead(buf[b:], &v.Cc)

	b = 0 // always read the same bytes
	v.Oc = OutputChange{}
	b += OutputChangeRead(buf[b:], &v.Oc)

	b = 0 // always read the same bytes
	v.Op = OutputProperty{}
	b += OutputPropertyRead(buf[b:], &v.Op)

	b = 0 // always read the same bytes
	v.Pc = ProviderChange{}
	b += ProviderChangeRead(buf[b:], &v.Pc)

	b = 0 // always read the same bytes
	v.Pp = ProviderProperty{}
	b += ProviderPropertyRead(buf[b:], &v.Pp)

	b = 0 // always read the same bytes
	v.Rc = ResourceChange{}
	b += ResourceChangeRead(buf[b:], &v.Rc)

	b = 0 // always read the same bytes
	v.Lc = LeaseNotify{}
	b += LeaseNotifyRead(buf[b:], &v.Lc)

	return v
}

// NotifyDataUnionLcNew constructs a new NotifyDataUnion union type with the Lc field.
func NotifyDataUnionLcNew(Lc LeaseNotify) NotifyDataUnion {
	var b int
	buf := make([]byte, 28)

	{
		structBytes := Lc.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	// Create the Union type
	v := NotifyDataUnion{}

	// Now copy buf into all fields

	b = 0 // always read the same bytes
	v.Cc = CrtcChange{}
	b += CrtcChangeRead(buf[b:], &v.Cc)

	b = 0 // always read the same bytes
	v.Oc = OutputChange{}
	b += OutputChangeRead(buf[b:], &v.Oc)

	b = 0 // always read the same bytes
	v.Op = OutputProperty{}
	b += OutputPropertyRead(buf[b:], &v.Op)

	b = 0 // always read the same bytes
	v.Pc = ProviderChange{}
	b += ProviderChangeRead(buf[b:], &v.Pc)

	b = 0 // always read the same bytes
	v.Pp = ProviderProperty{}
	b += ProviderPropertyRead(buf[b:], &v.Pp)

	b = 0 // always read the same bytes
	v.Rc = ResourceChange{}
	b += ResourceChangeRead(buf[b:], &v.Rc)

	b = 0 // always read the same bytes
	v.Lc = LeaseNotify{}
	b += LeaseNotifyRead(buf[b:], &v.Lc)

	return v
}

// NotifyDataUnionRead reads a byte slice into a NotifyDataUnion value.
func NotifyDataUnionRead(buf []byte, v *NotifyDataUnion) int {
	var b int

	b = 0 // re-read the same bytes
	v.Cc = CrtcChange{}
	b += CrtcChangeRead(buf[b:], &v.Cc)

	b = 0 // re-read the same bytes
	v.Oc = OutputChange{}
	b += OutputChangeRead(buf[b:], &v.Oc)

	b = 0 // re-read the same bytes
	v.Op = OutputProperty{}
	b += OutputPropertyRead(buf[b:], &v.Op)

	b = 0 // re-read the same bytes
	v.Pc = ProviderChange{}
	b += ProviderChangeRead(buf[b:], &v.Pc)

	b = 0 // re-read the same bytes
	v.Pp = ProviderProperty{}
	b += ProviderPropertyRead(buf[b:], &v.Pp)

	b = 0 // re-read the same bytes
	v.Rc = ResourceChange{}
	b += ResourceChangeRead(buf[b:], &v.Rc)

	b = 0 // re-read the same bytes
	v.Lc = LeaseNotify{}
	b += LeaseNotifyRead(buf[b:], &v.Lc)

	return 28
}

// NotifyDataUnionReadList reads a byte slice into a list of NotifyDataUnion values.
func NotifyDataUnionReadList(buf []byte, dest []NotifyDataUnion) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = NotifyDataUnion{}
		b += NotifyDataUnionRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a NotifyDataUnion value to a byte slice.
// Each field in a union must contain the same data.
// So simply pick the first field and write that to the wire.
func (v NotifyDataUnion) Bytes() []byte {
	buf := make([]byte, 28)
	b := 0

	{
		structBytes := v.Cc.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return buf
}

// NotifyDataUnionListBytes writes a list of NotifyDataUnion values to a byte slice.
func NotifyDataUnionListBytes(buf []byte, list []NotifyDataUnion) int {
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
	NotifyMaskScreenChange     = 1
	NotifyMaskCrtcChange       = 2
	NotifyMaskOutputChange     = 4
	NotifyMaskOutputProperty   = 8
	NotifyMaskProviderChange   = 16
	NotifyMaskProviderProperty = 32
	NotifyMaskResourceChange   = 64
	NotifyMaskLease            = 128
)

type Output uint32

func NewOutputID(c *xgb.XConn) Output {
	id := c.NewXID()
	return Output(id)
}

type OutputChange struct {
	Timestamp       xproto.Timestamp
	ConfigTimestamp xproto.Timestamp
	Window          xproto.Window
	Output          Output
	Crtc            Crtc
	Mode            Mode
	Rotation        uint16
	Connection      byte
	SubpixelOrder   byte
}

// OutputChangeRead reads a byte slice into a OutputChange value.
func OutputChangeRead(buf []byte, v *OutputChange) int {
	b := 0

	v.Timestamp = xproto.Timestamp(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.ConfigTimestamp = xproto.Timestamp(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Window = xproto.Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Output = Output(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Crtc = Crtc(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Mode = Mode(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Rotation = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Connection = buf[b]
	b += 1

	v.SubpixelOrder = buf[b]
	b += 1

	return b
}

// OutputChangeReadList reads a byte slice into a list of OutputChange values.
func OutputChangeReadList(buf []byte, dest []OutputChange) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = OutputChange{}
		b += OutputChangeRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a OutputChange value to a byte slice.
func (v OutputChange) Bytes() []byte {
	buf := make([]byte, 28)
	b := 0

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Timestamp))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.ConfigTimestamp))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Window))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Output))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Crtc))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Mode))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], v.Rotation)
	b += 2

	buf[b] = v.Connection
	b += 1

	buf[b] = v.SubpixelOrder
	b += 1

	return buf[:b]
}

// OutputChangeListBytes writes a list of OutputChange values to a byte slice.
func OutputChangeListBytes(buf []byte, list []OutputChange) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

type OutputProperty struct {
	Window    xproto.Window
	Output    Output
	Atom      xproto.Atom
	Timestamp xproto.Timestamp
	Status    byte
	// padding: 11 bytes
}

// OutputPropertyRead reads a byte slice into a OutputProperty value.
func OutputPropertyRead(buf []byte, v *OutputProperty) int {
	b := 0

	v.Window = xproto.Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Output = Output(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Atom = xproto.Atom(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Timestamp = xproto.Timestamp(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Status = buf[b]
	b += 1

	b += 11 // padding

	return b
}

// OutputPropertyReadList reads a byte slice into a list of OutputProperty values.
func OutputPropertyReadList(buf []byte, dest []OutputProperty) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = OutputProperty{}
		b += OutputPropertyRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a OutputProperty value to a byte slice.
func (v OutputProperty) Bytes() []byte {
	buf := make([]byte, 28)
	b := 0

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Window))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Output))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Atom))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Timestamp))
	b += 4

	buf[b] = v.Status
	b += 1

	b += 11 // padding

	return buf[:b]
}

// OutputPropertyListBytes writes a list of OutputProperty values to a byte slice.
func OutputPropertyListBytes(buf []byte, list []OutputProperty) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

type Provider uint32

func NewProviderID(c *xgb.XConn) Provider {
	id := c.NewXID()
	return Provider(id)
}

const (
	ProviderCapabilitySourceOutput  = 1
	ProviderCapabilitySinkOutput    = 2
	ProviderCapabilitySourceOffload = 4
	ProviderCapabilitySinkOffload   = 8
)

type ProviderChange struct {
	Timestamp xproto.Timestamp
	Window    xproto.Window
	Provider  Provider
	// padding: 16 bytes
}

// ProviderChangeRead reads a byte slice into a ProviderChange value.
func ProviderChangeRead(buf []byte, v *ProviderChange) int {
	b := 0

	v.Timestamp = xproto.Timestamp(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Window = xproto.Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Provider = Provider(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	b += 16 // padding

	return b
}

// ProviderChangeReadList reads a byte slice into a list of ProviderChange values.
func ProviderChangeReadList(buf []byte, dest []ProviderChange) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = ProviderChange{}
		b += ProviderChangeRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a ProviderChange value to a byte slice.
func (v ProviderChange) Bytes() []byte {
	buf := make([]byte, 28)
	b := 0

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Timestamp))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Window))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Provider))
	b += 4

	b += 16 // padding

	return buf[:b]
}

// ProviderChangeListBytes writes a list of ProviderChange values to a byte slice.
func ProviderChangeListBytes(buf []byte, list []ProviderChange) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

type ProviderProperty struct {
	Window    xproto.Window
	Provider  Provider
	Atom      xproto.Atom
	Timestamp xproto.Timestamp
	State     byte
	// padding: 11 bytes
}

// ProviderPropertyRead reads a byte slice into a ProviderProperty value.
func ProviderPropertyRead(buf []byte, v *ProviderProperty) int {
	b := 0

	v.Window = xproto.Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Provider = Provider(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Atom = xproto.Atom(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Timestamp = xproto.Timestamp(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.State = buf[b]
	b += 1

	b += 11 // padding

	return b
}

// ProviderPropertyReadList reads a byte slice into a list of ProviderProperty values.
func ProviderPropertyReadList(buf []byte, dest []ProviderProperty) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = ProviderProperty{}
		b += ProviderPropertyRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a ProviderProperty value to a byte slice.
func (v ProviderProperty) Bytes() []byte {
	buf := make([]byte, 28)
	b := 0

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Window))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Provider))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Atom))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Timestamp))
	b += 4

	buf[b] = v.State
	b += 1

	b += 11 // padding

	return buf[:b]
}

// ProviderPropertyListBytes writes a list of ProviderProperty values to a byte slice.
func ProviderPropertyListBytes(buf []byte, list []ProviderProperty) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

type RefreshRates struct {
	NRates uint16
	Rates  []uint16 // size: internal.Pad4((int(NRates) * 2))
}

// RefreshRatesRead reads a byte slice into a RefreshRates value.
func RefreshRatesRead(buf []byte, v *RefreshRates) int {
	b := 0

	v.NRates = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Rates = make([]uint16, v.NRates)
	for i := 0; i < int(v.NRates); i++ {
		v.Rates[i] = binary.LittleEndian.Uint16(buf[b:])
		b += 2
	}

	return b
}

// RefreshRatesReadList reads a byte slice into a list of RefreshRates values.
func RefreshRatesReadList(buf []byte, dest []RefreshRates) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = RefreshRates{}
		b += RefreshRatesRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a RefreshRates value to a byte slice.
func (v RefreshRates) Bytes() []byte {
	buf := make([]byte, (2 + internal.Pad4((int(v.NRates) * 2))))
	b := 0

	binary.LittleEndian.PutUint16(buf[b:], v.NRates)
	b += 2

	for i := 0; i < int(v.NRates); i++ {
		binary.LittleEndian.PutUint16(buf[b:], v.Rates[i])
		b += 2
	}

	return buf[:b]
}

// RefreshRatesListBytes writes a list of RefreshRates values to a byte slice.
func RefreshRatesListBytes(buf []byte, list []RefreshRates) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

// RefreshRatesListSize computes the size (bytes) of a list of RefreshRates values.
func RefreshRatesListSize(list []RefreshRates) int {
	size := 0
	for _, item := range list {
		size += (2 + internal.Pad4((int(item.NRates) * 2)))
	}
	return size
}

type ResourceChange struct {
	Timestamp xproto.Timestamp
	Window    xproto.Window
	// padding: 20 bytes
}

// ResourceChangeRead reads a byte slice into a ResourceChange value.
func ResourceChangeRead(buf []byte, v *ResourceChange) int {
	b := 0

	v.Timestamp = xproto.Timestamp(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Window = xproto.Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	b += 20 // padding

	return b
}

// ResourceChangeReadList reads a byte slice into a list of ResourceChange values.
func ResourceChangeReadList(buf []byte, dest []ResourceChange) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = ResourceChange{}
		b += ResourceChangeRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a ResourceChange value to a byte slice.
func (v ResourceChange) Bytes() []byte {
	buf := make([]byte, 28)
	b := 0

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Timestamp))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Window))
	b += 4

	b += 20 // padding

	return buf[:b]
}

// ResourceChangeListBytes writes a list of ResourceChange values to a byte slice.
func ResourceChangeListBytes(buf []byte, list []ResourceChange) int {
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
	RotationRotate0   = 1
	RotationRotate90  = 2
	RotationRotate180 = 4
	RotationRotate270 = 8
	RotationReflectX  = 16
	RotationReflectY  = 32
)

// ScreenChangeNotify is the event number for a ScreenChangeNotifyEvent.
const ScreenChangeNotify = 0

type ScreenChangeNotifyEvent struct {
	Sequence        uint16
	Rotation        byte
	Timestamp       xproto.Timestamp
	ConfigTimestamp xproto.Timestamp
	Root            xproto.Window
	RequestWindow   xproto.Window
	SizeID          uint16
	SubpixelOrder   uint16
	Width           uint16
	Height          uint16
	Mwidth          uint16
	Mheight         uint16
}

// UnmarshalScreenChangeNotifyEvent constructs a ScreenChangeNotifyEvent value that implements xgb.Event from a byte slice.
func UnmarshalScreenChangeNotifyEvent(buf []byte) (xgb.XEvent, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"ScreenChangeNotifyEvent\"", len(buf))
	}

	v := &ScreenChangeNotifyEvent{}
	b := 1 // don't read event number

	v.Rotation = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Timestamp = xproto.Timestamp(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.ConfigTimestamp = xproto.Timestamp(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Root = xproto.Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.RequestWindow = xproto.Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.SizeID = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.SubpixelOrder = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Width = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Height = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Mwidth = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Mheight = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return v, nil
}

// Bytes writes a ScreenChangeNotifyEvent value to a byte slice.
func (v *ScreenChangeNotifyEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 0
	b += 1

	buf[b] = v.Rotation
	b += 1

	b += 2 // skip sequence number

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Timestamp))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.ConfigTimestamp))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Root))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.RequestWindow))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], v.SizeID)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.SubpixelOrder)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Width)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Height)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Mwidth)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Mheight)
	b += 2

	return buf
}

// SeqID returns the sequence id attached to the ScreenChangeNotify event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v *ScreenChangeNotifyEvent) SeqID() uint16 {
	return v.Sequence
}

func init() { registerEvent(0, UnmarshalScreenChangeNotifyEvent) }

type ScreenSize struct {
	Width   uint16
	Height  uint16
	Mwidth  uint16
	Mheight uint16
}

// ScreenSizeRead reads a byte slice into a ScreenSize value.
func ScreenSizeRead(buf []byte, v *ScreenSize) int {
	b := 0

	v.Width = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Height = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Mwidth = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Mheight = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return b
}

// ScreenSizeReadList reads a byte slice into a list of ScreenSize values.
func ScreenSizeReadList(buf []byte, dest []ScreenSize) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = ScreenSize{}
		b += ScreenSizeRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a ScreenSize value to a byte slice.
func (v ScreenSize) Bytes() []byte {
	buf := make([]byte, 8)
	b := 0

	binary.LittleEndian.PutUint16(buf[b:], v.Width)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Height)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Mwidth)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Mheight)
	b += 2

	return buf[:b]
}

// ScreenSizeListBytes writes a list of ScreenSize values to a byte slice.
func ScreenSizeListBytes(buf []byte, list []ScreenSize) int {
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
	SetConfigSuccess           = 0
	SetConfigInvalidConfigTime = 1
	SetConfigInvalidTime       = 2
	SetConfigFailed            = 3
)

const (
	TransformUnit       = 1
	TransformScaleUp    = 2
	TransformScaleDown  = 4
	TransformProjective = 8
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

// AddOutputMode sends a checked request.
func AddOutputMode(c *xgb.XConn, Output Output, Mode Mode) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"AddOutputMode\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.SendRecv(addOutputModeRequest(op, Output, Mode), nil)
}

// AddOutputModeUnchecked sends an unchecked request.
func AddOutputModeUnchecked(c *xgb.XConn, Output Output, Mode Mode) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"AddOutputMode\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.Send(addOutputModeRequest(op, Output, Mode))
}

// Write request to wire for AddOutputMode
// addOutputModeRequest writes a AddOutputMode request to a byte slice.
func addOutputModeRequest(opcode uint8, Output Output, Mode Mode) []byte {
	const size = 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 18 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Output))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Mode))
	b += 4

	return buf
}

// ChangeOutputProperty sends a checked request.
func ChangeOutputProperty(c *xgb.XConn, Output Output, Property xproto.Atom, Type xproto.Atom, Format byte, Mode byte, NumUnits uint32, Data []byte) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"ChangeOutputProperty\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.SendRecv(changeOutputPropertyRequest(op, Output, Property, Type, Format, Mode, NumUnits, Data), nil)
}

// ChangeOutputPropertyUnchecked sends an unchecked request.
func ChangeOutputPropertyUnchecked(c *xgb.XConn, Output Output, Property xproto.Atom, Type xproto.Atom, Format byte, Mode byte, NumUnits uint32, Data []byte) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"ChangeOutputProperty\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.Send(changeOutputPropertyRequest(op, Output, Property, Type, Format, Mode, NumUnits, Data))
}

// Write request to wire for ChangeOutputProperty
// changeOutputPropertyRequest writes a ChangeOutputProperty request to a byte slice.
func changeOutputPropertyRequest(opcode uint8, Output Output, Property xproto.Atom, Type xproto.Atom, Format byte, Mode byte, NumUnits uint32, Data []byte) []byte {
	size := internal.Pad4((24 + internal.Pad4((((int(NumUnits) * int(Format)) / 8) * 1))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 13 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Output))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Property))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Type))
	b += 4

	buf[b] = Format
	b += 1

	buf[b] = Mode
	b += 1

	b += 2 // padding

	binary.LittleEndian.PutUint32(buf[b:], NumUnits)
	b += 4

	copy(buf[b:], Data[:((int(NumUnits)*int(Format))/8)])
	b += int(((int(NumUnits) * int(Format)) / 8))

	return buf
}

// ChangeProviderProperty sends a checked request.
func ChangeProviderProperty(c *xgb.XConn, Provider Provider, Property xproto.Atom, Type xproto.Atom, Format byte, Mode byte, NumItems uint32, Data []byte) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"ChangeProviderProperty\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.SendRecv(changeProviderPropertyRequest(op, Provider, Property, Type, Format, Mode, NumItems, Data), nil)
}

// ChangeProviderPropertyUnchecked sends an unchecked request.
func ChangeProviderPropertyUnchecked(c *xgb.XConn, Provider Provider, Property xproto.Atom, Type xproto.Atom, Format byte, Mode byte, NumItems uint32, Data []byte) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"ChangeProviderProperty\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.Send(changeProviderPropertyRequest(op, Provider, Property, Type, Format, Mode, NumItems, Data))
}

// Write request to wire for ChangeProviderProperty
// changeProviderPropertyRequest writes a ChangeProviderProperty request to a byte slice.
func changeProviderPropertyRequest(opcode uint8, Provider Provider, Property xproto.Atom, Type xproto.Atom, Format byte, Mode byte, NumItems uint32, Data []byte) []byte {
	size := internal.Pad4((24 + internal.Pad4(((int(NumItems) * (int(Format) / 8)) * 1))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 39 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Provider))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Property))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Type))
	b += 4

	buf[b] = Format
	b += 1

	buf[b] = Mode
	b += 1

	b += 2 // padding

	binary.LittleEndian.PutUint32(buf[b:], NumItems)
	b += 4

	copy(buf[b:], Data[:(int(NumItems)*(int(Format)/8))])
	b += int((int(NumItems) * (int(Format) / 8)))

	return buf
}

// ConfigureOutputProperty sends a checked request.
func ConfigureOutputProperty(c *xgb.XConn, Output Output, Property xproto.Atom, Pending bool, Range bool, Values []int32) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"ConfigureOutputProperty\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.SendRecv(configureOutputPropertyRequest(op, Output, Property, Pending, Range, Values), nil)
}

// ConfigureOutputPropertyUnchecked sends an unchecked request.
func ConfigureOutputPropertyUnchecked(c *xgb.XConn, Output Output, Property xproto.Atom, Pending bool, Range bool, Values []int32) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"ConfigureOutputProperty\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.Send(configureOutputPropertyRequest(op, Output, Property, Pending, Range, Values))
}

// Write request to wire for ConfigureOutputProperty
// configureOutputPropertyRequest writes a ConfigureOutputProperty request to a byte slice.
func configureOutputPropertyRequest(opcode uint8, Output Output, Property xproto.Atom, Pending bool, Range bool, Values []int32) []byte {
	size := internal.Pad4((16 + internal.Pad4((len(Values) * 4))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 12 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Output))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Property))
	b += 4

	if Pending {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	if Range {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	b += 2 // padding

	for i := 0; i < int(len(Values)); i++ {
		binary.LittleEndian.PutUint32(buf[b:], uint32(Values[i]))
		b += 4
	}

	return buf
}

// ConfigureProviderProperty sends a checked request.
func ConfigureProviderProperty(c *xgb.XConn, Provider Provider, Property xproto.Atom, Pending bool, Range bool, Values []int32) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"ConfigureProviderProperty\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.SendRecv(configureProviderPropertyRequest(op, Provider, Property, Pending, Range, Values), nil)
}

// ConfigureProviderPropertyUnchecked sends an unchecked request.
func ConfigureProviderPropertyUnchecked(c *xgb.XConn, Provider Provider, Property xproto.Atom, Pending bool, Range bool, Values []int32) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"ConfigureProviderProperty\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.Send(configureProviderPropertyRequest(op, Provider, Property, Pending, Range, Values))
}

// Write request to wire for ConfigureProviderProperty
// configureProviderPropertyRequest writes a ConfigureProviderProperty request to a byte slice.
func configureProviderPropertyRequest(opcode uint8, Provider Provider, Property xproto.Atom, Pending bool, Range bool, Values []int32) []byte {
	size := internal.Pad4((16 + internal.Pad4((len(Values) * 4))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 38 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Provider))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Property))
	b += 4

	if Pending {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	if Range {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	b += 2 // padding

	for i := 0; i < int(len(Values)); i++ {
		binary.LittleEndian.PutUint32(buf[b:], uint32(Values[i]))
		b += 4
	}

	return buf
}

// CreateLease sends a checked request.
func CreateLease(c *xgb.XConn, Window xproto.Window, Lid Lease, NumCrtcs uint16, NumOutputs uint16, Crtcs []Crtc, Outputs []Output) (CreateLeaseReply, error) {
	var reply CreateLeaseReply
	op, ok := c.Ext("RANDR")
	if !ok {
		return reply, errors.New("cannot issue request \"CreateLease\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	err := c.SendRecv(createLeaseRequest(op, Window, Lid, NumCrtcs, NumOutputs, Crtcs, Outputs), &reply)
	return reply, err
}

// CreateLeaseUnchecked sends an unchecked request.
func CreateLeaseUnchecked(c *xgb.XConn, Window xproto.Window, Lid Lease, NumCrtcs uint16, NumOutputs uint16, Crtcs []Crtc, Outputs []Output) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"CreateLease\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.Send(createLeaseRequest(op, Window, Lid, NumCrtcs, NumOutputs, Crtcs, Outputs))
}

// CreateLeaseReply represents the data returned from a CreateLease request.
type CreateLeaseReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	Nfd      byte
	// padding: 24 bytes
}

// Unmarshal reads a byte slice into a CreateLeaseReply value.
func (v *CreateLeaseReply) Unmarshal(buf []byte) error {
	const size = 32
	if len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"CreateLeaseReply\": have=%d need=%d", len(buf), size)
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

// Write request to wire for CreateLease
// createLeaseRequest writes a CreateLease request to a byte slice.
func createLeaseRequest(opcode uint8, Window xproto.Window, Lid Lease, NumCrtcs uint16, NumOutputs uint16, Crtcs []Crtc, Outputs []Output) []byte {
	size := internal.Pad4((((16 + internal.Pad4((int(NumCrtcs) * 4))) + 4) + internal.Pad4((int(NumOutputs) * 4))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 45 // request opcode
	b += 1

	blen := b
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Lid))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], NumCrtcs)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], NumOutputs)
	b += 2

	for i := 0; i < int(NumCrtcs); i++ {
		binary.LittleEndian.PutUint32(buf[b:], uint32(Crtcs[i]))
		b += 4
	}

	b = (b + 3) & ^3 // alignment gap

	for i := 0; i < int(NumOutputs); i++ {
		binary.LittleEndian.PutUint32(buf[b:], uint32(Outputs[i]))
		b += 4
	}

	b = internal.Pad4(b)
	binary.LittleEndian.PutUint16(buf[blen:], uint16(b/4)) // write request size in 4-byte units
	return buf[:b]
}

// CreateMode sends a checked request.
func CreateMode(c *xgb.XConn, Window xproto.Window, ModeInfo ModeInfo, Name string) (CreateModeReply, error) {
	var reply CreateModeReply
	op, ok := c.Ext("RANDR")
	if !ok {
		return reply, errors.New("cannot issue request \"CreateMode\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	err := c.SendRecv(createModeRequest(op, Window, ModeInfo, Name), &reply)
	return reply, err
}

// CreateModeUnchecked sends an unchecked request.
func CreateModeUnchecked(c *xgb.XConn, Window xproto.Window, ModeInfo ModeInfo, Name string) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"CreateMode\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.Send(createModeRequest(op, Window, ModeInfo, Name))
}

// CreateModeReply represents the data returned from a CreateMode request.
type CreateModeReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Mode Mode
	// padding: 20 bytes
}

// Unmarshal reads a byte slice into a CreateModeReply value.
func (v *CreateModeReply) Unmarshal(buf []byte) error {
	const size = 32
	if len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"CreateModeReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Mode = Mode(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	b += 20 // padding

	return nil
}

// Write request to wire for CreateMode
// createModeRequest writes a CreateMode request to a byte slice.
func createModeRequest(opcode uint8, Window xproto.Window, ModeInfo ModeInfo, Name string) []byte {
	size := internal.Pad4((40 + internal.Pad4((len(Name) * 1))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 16 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	{
		structBytes := ModeInfo.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	copy(buf[b:], Name[:])
	b += int(len(Name))

	return buf
}

// DeleteMonitor sends a checked request.
func DeleteMonitor(c *xgb.XConn, Window xproto.Window, Name xproto.Atom) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"DeleteMonitor\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.SendRecv(deleteMonitorRequest(op, Window, Name), nil)
}

// DeleteMonitorUnchecked sends an unchecked request.
func DeleteMonitorUnchecked(c *xgb.XConn, Window xproto.Window, Name xproto.Atom) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"DeleteMonitor\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.Send(deleteMonitorRequest(op, Window, Name))
}

// Write request to wire for DeleteMonitor
// deleteMonitorRequest writes a DeleteMonitor request to a byte slice.
func deleteMonitorRequest(opcode uint8, Window xproto.Window, Name xproto.Atom) []byte {
	const size = 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 44 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Name))
	b += 4

	return buf
}

// DeleteOutputMode sends a checked request.
func DeleteOutputMode(c *xgb.XConn, Output Output, Mode Mode) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"DeleteOutputMode\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.SendRecv(deleteOutputModeRequest(op, Output, Mode), nil)
}

// DeleteOutputModeUnchecked sends an unchecked request.
func DeleteOutputModeUnchecked(c *xgb.XConn, Output Output, Mode Mode) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"DeleteOutputMode\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.Send(deleteOutputModeRequest(op, Output, Mode))
}

// Write request to wire for DeleteOutputMode
// deleteOutputModeRequest writes a DeleteOutputMode request to a byte slice.
func deleteOutputModeRequest(opcode uint8, Output Output, Mode Mode) []byte {
	const size = 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 19 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Output))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Mode))
	b += 4

	return buf
}

// DeleteOutputProperty sends a checked request.
func DeleteOutputProperty(c *xgb.XConn, Output Output, Property xproto.Atom) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"DeleteOutputProperty\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.SendRecv(deleteOutputPropertyRequest(op, Output, Property), nil)
}

// DeleteOutputPropertyUnchecked sends an unchecked request.
func DeleteOutputPropertyUnchecked(c *xgb.XConn, Output Output, Property xproto.Atom) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"DeleteOutputProperty\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.Send(deleteOutputPropertyRequest(op, Output, Property))
}

// Write request to wire for DeleteOutputProperty
// deleteOutputPropertyRequest writes a DeleteOutputProperty request to a byte slice.
func deleteOutputPropertyRequest(opcode uint8, Output Output, Property xproto.Atom) []byte {
	const size = 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 14 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Output))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Property))
	b += 4

	return buf
}

// DeleteProviderProperty sends a checked request.
func DeleteProviderProperty(c *xgb.XConn, Provider Provider, Property xproto.Atom) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"DeleteProviderProperty\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.SendRecv(deleteProviderPropertyRequest(op, Provider, Property), nil)
}

// DeleteProviderPropertyUnchecked sends an unchecked request.
func DeleteProviderPropertyUnchecked(c *xgb.XConn, Provider Provider, Property xproto.Atom) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"DeleteProviderProperty\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.Send(deleteProviderPropertyRequest(op, Provider, Property))
}

// Write request to wire for DeleteProviderProperty
// deleteProviderPropertyRequest writes a DeleteProviderProperty request to a byte slice.
func deleteProviderPropertyRequest(opcode uint8, Provider Provider, Property xproto.Atom) []byte {
	const size = 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 40 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Provider))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Property))
	b += 4

	return buf
}

// DestroyMode sends a checked request.
func DestroyMode(c *xgb.XConn, Mode Mode) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"DestroyMode\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.SendRecv(destroyModeRequest(op, Mode), nil)
}

// DestroyModeUnchecked sends an unchecked request.
func DestroyModeUnchecked(c *xgb.XConn, Mode Mode) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"DestroyMode\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.Send(destroyModeRequest(op, Mode))
}

// Write request to wire for DestroyMode
// destroyModeRequest writes a DestroyMode request to a byte slice.
func destroyModeRequest(opcode uint8, Mode Mode) []byte {
	const size = 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 17 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Mode))
	b += 4

	return buf
}

// FreeLease sends a checked request.
func FreeLease(c *xgb.XConn, Lid Lease, Terminate byte) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"FreeLease\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.SendRecv(freeLeaseRequest(op, Lid, Terminate), nil)
}

// FreeLeaseUnchecked sends an unchecked request.
func FreeLeaseUnchecked(c *xgb.XConn, Lid Lease, Terminate byte) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"FreeLease\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.Send(freeLeaseRequest(op, Lid, Terminate))
}

// Write request to wire for FreeLease
// freeLeaseRequest writes a FreeLease request to a byte slice.
func freeLeaseRequest(opcode uint8, Lid Lease, Terminate byte) []byte {
	const size = 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 46 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Lid))
	b += 4

	buf[b] = Terminate
	b += 1

	return buf
}

// GetCrtcGamma sends a checked request.
func GetCrtcGamma(c *xgb.XConn, Crtc Crtc) (GetCrtcGammaReply, error) {
	var reply GetCrtcGammaReply
	op, ok := c.Ext("RANDR")
	if !ok {
		return reply, errors.New("cannot issue request \"GetCrtcGamma\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getCrtcGammaRequest(op, Crtc), &reply)
	return reply, err
}

// GetCrtcGammaUnchecked sends an unchecked request.
func GetCrtcGammaUnchecked(c *xgb.XConn, Crtc Crtc) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"GetCrtcGamma\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.Send(getCrtcGammaRequest(op, Crtc))
}

// GetCrtcGammaReply represents the data returned from a GetCrtcGamma request.
type GetCrtcGammaReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Size uint16
	// padding: 22 bytes
	Red []uint16 // size: internal.Pad4((int(Size) * 2))
	// alignment gap to multiple of 2
	Green []uint16 // size: internal.Pad4((int(Size) * 2))
	// alignment gap to multiple of 2
	Blue []uint16 // size: internal.Pad4((int(Size) * 2))
}

// Unmarshal reads a byte slice into a GetCrtcGammaReply value.
func (v *GetCrtcGammaReply) Unmarshal(buf []byte) error {
	if size := (((((32 + internal.Pad4((int(v.Size) * 2))) + 2) + internal.Pad4((int(v.Size) * 2))) + 2) + internal.Pad4((int(v.Size) * 2))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetCrtcGammaReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Size = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	b += 22 // padding

	v.Red = make([]uint16, v.Size)
	for i := 0; i < int(v.Size); i++ {
		v.Red[i] = binary.LittleEndian.Uint16(buf[b:])
		b += 2
	}

	b = (b + 1) & ^1 // alignment gap

	v.Green = make([]uint16, v.Size)
	for i := 0; i < int(v.Size); i++ {
		v.Green[i] = binary.LittleEndian.Uint16(buf[b:])
		b += 2
	}

	b = (b + 1) & ^1 // alignment gap

	v.Blue = make([]uint16, v.Size)
	for i := 0; i < int(v.Size); i++ {
		v.Blue[i] = binary.LittleEndian.Uint16(buf[b:])
		b += 2
	}

	return nil
}

// Write request to wire for GetCrtcGamma
// getCrtcGammaRequest writes a GetCrtcGamma request to a byte slice.
func getCrtcGammaRequest(opcode uint8, Crtc Crtc) []byte {
	const size = 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 23 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Crtc))
	b += 4

	return buf
}

// GetCrtcGammaSize sends a checked request.
func GetCrtcGammaSize(c *xgb.XConn, Crtc Crtc) (GetCrtcGammaSizeReply, error) {
	var reply GetCrtcGammaSizeReply
	op, ok := c.Ext("RANDR")
	if !ok {
		return reply, errors.New("cannot issue request \"GetCrtcGammaSize\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getCrtcGammaSizeRequest(op, Crtc), &reply)
	return reply, err
}

// GetCrtcGammaSizeUnchecked sends an unchecked request.
func GetCrtcGammaSizeUnchecked(c *xgb.XConn, Crtc Crtc) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"GetCrtcGammaSize\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.Send(getCrtcGammaSizeRequest(op, Crtc))
}

// GetCrtcGammaSizeReply represents the data returned from a GetCrtcGammaSize request.
type GetCrtcGammaSizeReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Size uint16
	// padding: 22 bytes
}

// Unmarshal reads a byte slice into a GetCrtcGammaSizeReply value.
func (v *GetCrtcGammaSizeReply) Unmarshal(buf []byte) error {
	const size = 32
	if len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetCrtcGammaSizeReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Size = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	b += 22 // padding

	return nil
}

// Write request to wire for GetCrtcGammaSize
// getCrtcGammaSizeRequest writes a GetCrtcGammaSize request to a byte slice.
func getCrtcGammaSizeRequest(opcode uint8, Crtc Crtc) []byte {
	const size = 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 22 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Crtc))
	b += 4

	return buf
}

// GetCrtcInfo sends a checked request.
func GetCrtcInfo(c *xgb.XConn, Crtc Crtc, ConfigTimestamp xproto.Timestamp) (GetCrtcInfoReply, error) {
	var reply GetCrtcInfoReply
	op, ok := c.Ext("RANDR")
	if !ok {
		return reply, errors.New("cannot issue request \"GetCrtcInfo\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getCrtcInfoRequest(op, Crtc, ConfigTimestamp), &reply)
	return reply, err
}

// GetCrtcInfoUnchecked sends an unchecked request.
func GetCrtcInfoUnchecked(c *xgb.XConn, Crtc Crtc, ConfigTimestamp xproto.Timestamp) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"GetCrtcInfo\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.Send(getCrtcInfoRequest(op, Crtc, ConfigTimestamp))
}

// GetCrtcInfoReply represents the data returned from a GetCrtcInfo request.
type GetCrtcInfoReply struct {
	Sequence           uint16 // sequence number of the request for this reply
	Length             uint32 // number of bytes in this reply
	Status             byte
	Timestamp          xproto.Timestamp
	X                  int16
	Y                  int16
	Width              uint16
	Height             uint16
	Mode               Mode
	Rotation           uint16
	Rotations          uint16
	NumOutputs         uint16
	NumPossibleOutputs uint16
	Outputs            []Output // size: internal.Pad4((int(NumOutputs) * 4))
	// alignment gap to multiple of 4
	Possible []Output // size: internal.Pad4((int(NumPossibleOutputs) * 4))
}

// Unmarshal reads a byte slice into a GetCrtcInfoReply value.
func (v *GetCrtcInfoReply) Unmarshal(buf []byte) error {
	if size := (((32 + internal.Pad4((int(v.NumOutputs) * 4))) + 4) + internal.Pad4((int(v.NumPossibleOutputs) * 4))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetCrtcInfoReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	v.Status = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Timestamp = xproto.Timestamp(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.X = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.Y = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.Width = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Height = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Mode = Mode(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Rotation = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Rotations = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.NumOutputs = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.NumPossibleOutputs = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Outputs = make([]Output, v.NumOutputs)
	for i := 0; i < int(v.NumOutputs); i++ {
		v.Outputs[i] = Output(binary.LittleEndian.Uint32(buf[b:]))
		b += 4
	}

	b = (b + 3) & ^3 // alignment gap

	v.Possible = make([]Output, v.NumPossibleOutputs)
	for i := 0; i < int(v.NumPossibleOutputs); i++ {
		v.Possible[i] = Output(binary.LittleEndian.Uint32(buf[b:]))
		b += 4
	}

	return nil
}

// Write request to wire for GetCrtcInfo
// getCrtcInfoRequest writes a GetCrtcInfo request to a byte slice.
func getCrtcInfoRequest(opcode uint8, Crtc Crtc, ConfigTimestamp xproto.Timestamp) []byte {
	const size = 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 20 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Crtc))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(ConfigTimestamp))
	b += 4

	return buf
}

// GetCrtcTransform sends a checked request.
func GetCrtcTransform(c *xgb.XConn, Crtc Crtc) (GetCrtcTransformReply, error) {
	var reply GetCrtcTransformReply
	op, ok := c.Ext("RANDR")
	if !ok {
		return reply, errors.New("cannot issue request \"GetCrtcTransform\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getCrtcTransformRequest(op, Crtc), &reply)
	return reply, err
}

// GetCrtcTransformUnchecked sends an unchecked request.
func GetCrtcTransformUnchecked(c *xgb.XConn, Crtc Crtc) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"GetCrtcTransform\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.Send(getCrtcTransformRequest(op, Crtc))
}

// GetCrtcTransformReply represents the data returned from a GetCrtcTransform request.
type GetCrtcTransformReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	PendingTransform render.Transform
	HasTransforms    bool
	// padding: 3 bytes
	CurrentTransform render.Transform
	// padding: 4 bytes
	PendingLen        uint16
	PendingNparams    uint16
	CurrentLen        uint16
	CurrentNparams    uint16
	PendingFilterName string // size: internal.Pad4((int(PendingLen) * 1))
	// padding: 0 bytes
	PendingParams     []render.Fixed // size: internal.Pad4((int(PendingNparams) * 4))
	CurrentFilterName string         // size: internal.Pad4((int(CurrentLen) * 1))
	// padding: 0 bytes
	CurrentParams []render.Fixed // size: internal.Pad4((int(CurrentNparams) * 4))
}

// Unmarshal reads a byte slice into a GetCrtcTransformReply value.
func (v *GetCrtcTransformReply) Unmarshal(buf []byte) error {
	if size := ((((((96 + internal.Pad4((int(v.PendingLen) * 1))) + 0) + internal.Pad4((int(v.PendingNparams) * 4))) + internal.Pad4((int(v.CurrentLen) * 1))) + 0) + internal.Pad4((int(v.CurrentNparams) * 4))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetCrtcTransformReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.PendingTransform = render.Transform{}
	b += render.TransformRead(buf[b:], &v.PendingTransform)

	v.HasTransforms = (buf[b] == 1)
	b += 1

	b += 3 // padding

	v.CurrentTransform = render.Transform{}
	b += render.TransformRead(buf[b:], &v.CurrentTransform)

	b += 4 // padding

	v.PendingLen = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.PendingNparams = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.CurrentLen = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.CurrentNparams = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	{
		byteString := make([]byte, v.PendingLen)
		copy(byteString[:v.PendingLen], buf[b:])
		v.PendingFilterName = string(byteString)
		b += int(v.PendingLen)
	}

	b += 0 // padding

	v.PendingParams = make([]render.Fixed, v.PendingNparams)
	for i := 0; i < int(v.PendingNparams); i++ {
		v.PendingParams[i] = render.Fixed(binary.LittleEndian.Uint32(buf[b:]))
		b += 4
	}

	{
		byteString := make([]byte, v.CurrentLen)
		copy(byteString[:v.CurrentLen], buf[b:])
		v.CurrentFilterName = string(byteString)
		b += int(v.CurrentLen)
	}

	b += 0 // padding

	v.CurrentParams = make([]render.Fixed, v.CurrentNparams)
	for i := 0; i < int(v.CurrentNparams); i++ {
		v.CurrentParams[i] = render.Fixed(binary.LittleEndian.Uint32(buf[b:]))
		b += 4
	}

	return nil
}

// Write request to wire for GetCrtcTransform
// getCrtcTransformRequest writes a GetCrtcTransform request to a byte slice.
func getCrtcTransformRequest(opcode uint8, Crtc Crtc) []byte {
	const size = 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 27 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Crtc))
	b += 4

	return buf
}

// GetMonitors sends a checked request.
func GetMonitors(c *xgb.XConn, Window xproto.Window, GetActive bool) (GetMonitorsReply, error) {
	var reply GetMonitorsReply
	op, ok := c.Ext("RANDR")
	if !ok {
		return reply, errors.New("cannot issue request \"GetMonitors\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getMonitorsRequest(op, Window, GetActive), &reply)
	return reply, err
}

// GetMonitorsUnchecked sends an unchecked request.
func GetMonitorsUnchecked(c *xgb.XConn, Window xproto.Window, GetActive bool) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"GetMonitors\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.Send(getMonitorsRequest(op, Window, GetActive))
}

// GetMonitorsReply represents the data returned from a GetMonitors request.
type GetMonitorsReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Timestamp xproto.Timestamp
	NMonitors uint32
	NOutputs  uint32
	// padding: 12 bytes
	Monitors []MonitorInfo // size: MonitorInfoListSize(Monitors)
}

// Unmarshal reads a byte slice into a GetMonitorsReply value.
func (v *GetMonitorsReply) Unmarshal(buf []byte) error {
	if size := (32 + MonitorInfoListSize(v.Monitors)); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetMonitorsReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Timestamp = xproto.Timestamp(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.NMonitors = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.NOutputs = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 12 // padding

	v.Monitors = make([]MonitorInfo, v.NMonitors)
	b += MonitorInfoReadList(buf[b:], v.Monitors)

	return nil
}

// Write request to wire for GetMonitors
// getMonitorsRequest writes a GetMonitors request to a byte slice.
func getMonitorsRequest(opcode uint8, Window xproto.Window, GetActive bool) []byte {
	const size = 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 42 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	if GetActive {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	return buf
}

// GetOutputInfo sends a checked request.
func GetOutputInfo(c *xgb.XConn, Output Output, ConfigTimestamp xproto.Timestamp) (GetOutputInfoReply, error) {
	var reply GetOutputInfoReply
	op, ok := c.Ext("RANDR")
	if !ok {
		return reply, errors.New("cannot issue request \"GetOutputInfo\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getOutputInfoRequest(op, Output, ConfigTimestamp), &reply)
	return reply, err
}

// GetOutputInfoUnchecked sends an unchecked request.
func GetOutputInfoUnchecked(c *xgb.XConn, Output Output, ConfigTimestamp xproto.Timestamp) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"GetOutputInfo\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.Send(getOutputInfoRequest(op, Output, ConfigTimestamp))
}

// GetOutputInfoReply represents the data returned from a GetOutputInfo request.
type GetOutputInfoReply struct {
	Sequence      uint16 // sequence number of the request for this reply
	Length        uint32 // number of bytes in this reply
	Status        byte
	Timestamp     xproto.Timestamp
	Crtc          Crtc
	MmWidth       uint32
	MmHeight      uint32
	Connection    byte
	SubpixelOrder byte
	NumCrtcs      uint16
	NumModes      uint16
	NumPreferred  uint16
	NumClones     uint16
	NameLen       uint16
	Crtcs         []Crtc // size: internal.Pad4((int(NumCrtcs) * 4))
	// alignment gap to multiple of 4
	Modes []Mode // size: internal.Pad4((int(NumModes) * 4))
	// alignment gap to multiple of 4
	Clones []Output // size: internal.Pad4((int(NumClones) * 4))
	Name   []byte   // size: internal.Pad4((int(NameLen) * 1))
}

// Unmarshal reads a byte slice into a GetOutputInfoReply value.
func (v *GetOutputInfoReply) Unmarshal(buf []byte) error {
	if size := ((((((36 + internal.Pad4((int(v.NumCrtcs) * 4))) + 4) + internal.Pad4((int(v.NumModes) * 4))) + 4) + internal.Pad4((int(v.NumClones) * 4))) + internal.Pad4((int(v.NameLen) * 1))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetOutputInfoReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	v.Status = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Timestamp = xproto.Timestamp(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Crtc = Crtc(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.MmWidth = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.MmHeight = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.Connection = buf[b]
	b += 1

	v.SubpixelOrder = buf[b]
	b += 1

	v.NumCrtcs = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.NumModes = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.NumPreferred = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.NumClones = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.NameLen = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Crtcs = make([]Crtc, v.NumCrtcs)
	for i := 0; i < int(v.NumCrtcs); i++ {
		v.Crtcs[i] = Crtc(binary.LittleEndian.Uint32(buf[b:]))
		b += 4
	}

	b = (b + 3) & ^3 // alignment gap

	v.Modes = make([]Mode, v.NumModes)
	for i := 0; i < int(v.NumModes); i++ {
		v.Modes[i] = Mode(binary.LittleEndian.Uint32(buf[b:]))
		b += 4
	}

	b = (b + 3) & ^3 // alignment gap

	v.Clones = make([]Output, v.NumClones)
	for i := 0; i < int(v.NumClones); i++ {
		v.Clones[i] = Output(binary.LittleEndian.Uint32(buf[b:]))
		b += 4
	}

	v.Name = make([]byte, v.NameLen)
	copy(v.Name[:v.NameLen], buf[b:])
	b += int(v.NameLen)

	return nil
}

// Write request to wire for GetOutputInfo
// getOutputInfoRequest writes a GetOutputInfo request to a byte slice.
func getOutputInfoRequest(opcode uint8, Output Output, ConfigTimestamp xproto.Timestamp) []byte {
	const size = 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 9 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Output))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(ConfigTimestamp))
	b += 4

	return buf
}

// GetOutputPrimary sends a checked request.
func GetOutputPrimary(c *xgb.XConn, Window xproto.Window) (GetOutputPrimaryReply, error) {
	var reply GetOutputPrimaryReply
	op, ok := c.Ext("RANDR")
	if !ok {
		return reply, errors.New("cannot issue request \"GetOutputPrimary\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getOutputPrimaryRequest(op, Window), &reply)
	return reply, err
}

// GetOutputPrimaryUnchecked sends an unchecked request.
func GetOutputPrimaryUnchecked(c *xgb.XConn, Window xproto.Window) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"GetOutputPrimary\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.Send(getOutputPrimaryRequest(op, Window))
}

// GetOutputPrimaryReply represents the data returned from a GetOutputPrimary request.
type GetOutputPrimaryReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Output Output
}

// Unmarshal reads a byte slice into a GetOutputPrimaryReply value.
func (v *GetOutputPrimaryReply) Unmarshal(buf []byte) error {
	const size = 12
	if len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetOutputPrimaryReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Output = Output(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	return nil
}

// Write request to wire for GetOutputPrimary
// getOutputPrimaryRequest writes a GetOutputPrimary request to a byte slice.
func getOutputPrimaryRequest(opcode uint8, Window xproto.Window) []byte {
	const size = 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 31 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	return buf
}

// GetOutputProperty sends a checked request.
func GetOutputProperty(c *xgb.XConn, Output Output, Property xproto.Atom, Type xproto.Atom, LongOffset uint32, LongLength uint32, Delete bool, Pending bool) (GetOutputPropertyReply, error) {
	var reply GetOutputPropertyReply
	op, ok := c.Ext("RANDR")
	if !ok {
		return reply, errors.New("cannot issue request \"GetOutputProperty\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getOutputPropertyRequest(op, Output, Property, Type, LongOffset, LongLength, Delete, Pending), &reply)
	return reply, err
}

// GetOutputPropertyUnchecked sends an unchecked request.
func GetOutputPropertyUnchecked(c *xgb.XConn, Output Output, Property xproto.Atom, Type xproto.Atom, LongOffset uint32, LongLength uint32, Delete bool, Pending bool) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"GetOutputProperty\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.Send(getOutputPropertyRequest(op, Output, Property, Type, LongOffset, LongLength, Delete, Pending))
}

// GetOutputPropertyReply represents the data returned from a GetOutputProperty request.
type GetOutputPropertyReply struct {
	Sequence   uint16 // sequence number of the request for this reply
	Length     uint32 // number of bytes in this reply
	Format     byte
	Type       xproto.Atom
	BytesAfter uint32
	NumItems   uint32
	// padding: 12 bytes
	Data []byte // size: internal.Pad4(((int(NumItems) * (int(Format) / 8)) * 1))
}

// Unmarshal reads a byte slice into a GetOutputPropertyReply value.
func (v *GetOutputPropertyReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4(((int(v.NumItems) * (int(v.Format) / 8)) * 1))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetOutputPropertyReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	v.Format = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Type = xproto.Atom(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.BytesAfter = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.NumItems = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 12 // padding

	v.Data = make([]byte, (int(v.NumItems) * (int(v.Format) / 8)))
	copy(v.Data[:(int(v.NumItems)*(int(v.Format)/8))], buf[b:])
	b += int((int(v.NumItems) * (int(v.Format) / 8)))

	return nil
}

// Write request to wire for GetOutputProperty
// getOutputPropertyRequest writes a GetOutputProperty request to a byte slice.
func getOutputPropertyRequest(opcode uint8, Output Output, Property xproto.Atom, Type xproto.Atom, LongOffset uint32, LongLength uint32, Delete bool, Pending bool) []byte {
	const size = 28
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 15 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Output))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Property))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Type))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], LongOffset)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], LongLength)
	b += 4

	if Delete {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	if Pending {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	b += 2 // padding

	return buf
}

// GetPanning sends a checked request.
func GetPanning(c *xgb.XConn, Crtc Crtc) (GetPanningReply, error) {
	var reply GetPanningReply
	op, ok := c.Ext("RANDR")
	if !ok {
		return reply, errors.New("cannot issue request \"GetPanning\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getPanningRequest(op, Crtc), &reply)
	return reply, err
}

// GetPanningUnchecked sends an unchecked request.
func GetPanningUnchecked(c *xgb.XConn, Crtc Crtc) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"GetPanning\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.Send(getPanningRequest(op, Crtc))
}

// GetPanningReply represents the data returned from a GetPanning request.
type GetPanningReply struct {
	Sequence     uint16 // sequence number of the request for this reply
	Length       uint32 // number of bytes in this reply
	Status       byte
	Timestamp    xproto.Timestamp
	Left         uint16
	Top          uint16
	Width        uint16
	Height       uint16
	TrackLeft    uint16
	TrackTop     uint16
	TrackWidth   uint16
	TrackHeight  uint16
	BorderLeft   int16
	BorderTop    int16
	BorderRight  int16
	BorderBottom int16
}

// Unmarshal reads a byte slice into a GetPanningReply value.
func (v *GetPanningReply) Unmarshal(buf []byte) error {
	const size = 36
	if len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetPanningReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	v.Status = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Timestamp = xproto.Timestamp(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Left = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Top = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Width = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Height = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.TrackLeft = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.TrackTop = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.TrackWidth = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.TrackHeight = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.BorderLeft = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.BorderTop = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.BorderRight = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.BorderBottom = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	return nil
}

// Write request to wire for GetPanning
// getPanningRequest writes a GetPanning request to a byte slice.
func getPanningRequest(opcode uint8, Crtc Crtc) []byte {
	const size = 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 28 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Crtc))
	b += 4

	return buf
}

// GetProviderInfo sends a checked request.
func GetProviderInfo(c *xgb.XConn, Provider Provider, ConfigTimestamp xproto.Timestamp) (GetProviderInfoReply, error) {
	var reply GetProviderInfoReply
	op, ok := c.Ext("RANDR")
	if !ok {
		return reply, errors.New("cannot issue request \"GetProviderInfo\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getProviderInfoRequest(op, Provider, ConfigTimestamp), &reply)
	return reply, err
}

// GetProviderInfoUnchecked sends an unchecked request.
func GetProviderInfoUnchecked(c *xgb.XConn, Provider Provider, ConfigTimestamp xproto.Timestamp) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"GetProviderInfo\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.Send(getProviderInfoRequest(op, Provider, ConfigTimestamp))
}

// GetProviderInfoReply represents the data returned from a GetProviderInfo request.
type GetProviderInfoReply struct {
	Sequence               uint16 // sequence number of the request for this reply
	Length                 uint32 // number of bytes in this reply
	Status                 byte
	Timestamp              xproto.Timestamp
	Capabilities           uint32
	NumCrtcs               uint16
	NumOutputs             uint16
	NumAssociatedProviders uint16
	NameLen                uint16
	// padding: 8 bytes
	Crtcs []Crtc // size: internal.Pad4((int(NumCrtcs) * 4))
	// alignment gap to multiple of 4
	Outputs []Output // size: internal.Pad4((int(NumOutputs) * 4))
	// alignment gap to multiple of 4
	AssociatedProviders []Provider // size: internal.Pad4((int(NumAssociatedProviders) * 4))
	// alignment gap to multiple of 4
	AssociatedCapability []uint32 // size: internal.Pad4((int(NumAssociatedProviders) * 4))
	Name                 string   // size: internal.Pad4((int(NameLen) * 1))
}

// Unmarshal reads a byte slice into a GetProviderInfoReply value.
func (v *GetProviderInfoReply) Unmarshal(buf []byte) error {
	if size := ((((((((32 + internal.Pad4((int(v.NumCrtcs) * 4))) + 4) + internal.Pad4((int(v.NumOutputs) * 4))) + 4) + internal.Pad4((int(v.NumAssociatedProviders) * 4))) + 4) + internal.Pad4((int(v.NumAssociatedProviders) * 4))) + internal.Pad4((int(v.NameLen) * 1))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetProviderInfoReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	v.Status = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Timestamp = xproto.Timestamp(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Capabilities = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.NumCrtcs = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.NumOutputs = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.NumAssociatedProviders = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.NameLen = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	b += 8 // padding

	v.Crtcs = make([]Crtc, v.NumCrtcs)
	for i := 0; i < int(v.NumCrtcs); i++ {
		v.Crtcs[i] = Crtc(binary.LittleEndian.Uint32(buf[b:]))
		b += 4
	}

	b = (b + 3) & ^3 // alignment gap

	v.Outputs = make([]Output, v.NumOutputs)
	for i := 0; i < int(v.NumOutputs); i++ {
		v.Outputs[i] = Output(binary.LittleEndian.Uint32(buf[b:]))
		b += 4
	}

	b = (b + 3) & ^3 // alignment gap

	v.AssociatedProviders = make([]Provider, v.NumAssociatedProviders)
	for i := 0; i < int(v.NumAssociatedProviders); i++ {
		v.AssociatedProviders[i] = Provider(binary.LittleEndian.Uint32(buf[b:]))
		b += 4
	}

	b = (b + 3) & ^3 // alignment gap

	v.AssociatedCapability = make([]uint32, v.NumAssociatedProviders)
	for i := 0; i < int(v.NumAssociatedProviders); i++ {
		v.AssociatedCapability[i] = binary.LittleEndian.Uint32(buf[b:])
		b += 4
	}

	{
		byteString := make([]byte, v.NameLen)
		copy(byteString[:v.NameLen], buf[b:])
		v.Name = string(byteString)
		b += int(v.NameLen)
	}

	return nil
}

// Write request to wire for GetProviderInfo
// getProviderInfoRequest writes a GetProviderInfo request to a byte slice.
func getProviderInfoRequest(opcode uint8, Provider Provider, ConfigTimestamp xproto.Timestamp) []byte {
	const size = 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 33 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Provider))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(ConfigTimestamp))
	b += 4

	return buf
}

// GetProviderProperty sends a checked request.
func GetProviderProperty(c *xgb.XConn, Provider Provider, Property xproto.Atom, Type xproto.Atom, LongOffset uint32, LongLength uint32, Delete bool, Pending bool) (GetProviderPropertyReply, error) {
	var reply GetProviderPropertyReply
	op, ok := c.Ext("RANDR")
	if !ok {
		return reply, errors.New("cannot issue request \"GetProviderProperty\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getProviderPropertyRequest(op, Provider, Property, Type, LongOffset, LongLength, Delete, Pending), &reply)
	return reply, err
}

// GetProviderPropertyUnchecked sends an unchecked request.
func GetProviderPropertyUnchecked(c *xgb.XConn, Provider Provider, Property xproto.Atom, Type xproto.Atom, LongOffset uint32, LongLength uint32, Delete bool, Pending bool) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"GetProviderProperty\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.Send(getProviderPropertyRequest(op, Provider, Property, Type, LongOffset, LongLength, Delete, Pending))
}

// GetProviderPropertyReply represents the data returned from a GetProviderProperty request.
type GetProviderPropertyReply struct {
	Sequence   uint16 // sequence number of the request for this reply
	Length     uint32 // number of bytes in this reply
	Format     byte
	Type       xproto.Atom
	BytesAfter uint32
	NumItems   uint32
	// padding: 12 bytes
	Data []byte // size: internal.Pad4(((int(NumItems) * (int(Format) / 8)) * 1))
}

// Unmarshal reads a byte slice into a GetProviderPropertyReply value.
func (v *GetProviderPropertyReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4(((int(v.NumItems) * (int(v.Format) / 8)) * 1))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetProviderPropertyReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	v.Format = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Type = xproto.Atom(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.BytesAfter = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.NumItems = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 12 // padding

	v.Data = make([]byte, (int(v.NumItems) * (int(v.Format) / 8)))
	copy(v.Data[:(int(v.NumItems)*(int(v.Format)/8))], buf[b:])
	b += int((int(v.NumItems) * (int(v.Format) / 8)))

	return nil
}

// Write request to wire for GetProviderProperty
// getProviderPropertyRequest writes a GetProviderProperty request to a byte slice.
func getProviderPropertyRequest(opcode uint8, Provider Provider, Property xproto.Atom, Type xproto.Atom, LongOffset uint32, LongLength uint32, Delete bool, Pending bool) []byte {
	const size = 28
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 41 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Provider))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Property))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Type))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], LongOffset)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], LongLength)
	b += 4

	if Delete {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	if Pending {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	b += 2 // padding

	return buf
}

// GetProviders sends a checked request.
func GetProviders(c *xgb.XConn, Window xproto.Window) (GetProvidersReply, error) {
	var reply GetProvidersReply
	op, ok := c.Ext("RANDR")
	if !ok {
		return reply, errors.New("cannot issue request \"GetProviders\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getProvidersRequest(op, Window), &reply)
	return reply, err
}

// GetProvidersUnchecked sends an unchecked request.
func GetProvidersUnchecked(c *xgb.XConn, Window xproto.Window) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"GetProviders\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.Send(getProvidersRequest(op, Window))
}

// GetProvidersReply represents the data returned from a GetProviders request.
type GetProvidersReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Timestamp    xproto.Timestamp
	NumProviders uint16
	// padding: 18 bytes
	Providers []Provider // size: internal.Pad4((int(NumProviders) * 4))
}

// Unmarshal reads a byte slice into a GetProvidersReply value.
func (v *GetProvidersReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.NumProviders) * 4))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetProvidersReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Timestamp = xproto.Timestamp(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.NumProviders = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	b += 18 // padding

	v.Providers = make([]Provider, v.NumProviders)
	for i := 0; i < int(v.NumProviders); i++ {
		v.Providers[i] = Provider(binary.LittleEndian.Uint32(buf[b:]))
		b += 4
	}

	return nil
}

// Write request to wire for GetProviders
// getProvidersRequest writes a GetProviders request to a byte slice.
func getProvidersRequest(opcode uint8, Window xproto.Window) []byte {
	const size = 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 32 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	return buf
}

// GetScreenInfo sends a checked request.
func GetScreenInfo(c *xgb.XConn, Window xproto.Window) (GetScreenInfoReply, error) {
	var reply GetScreenInfoReply
	op, ok := c.Ext("RANDR")
	if !ok {
		return reply, errors.New("cannot issue request \"GetScreenInfo\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getScreenInfoRequest(op, Window), &reply)
	return reply, err
}

// GetScreenInfoUnchecked sends an unchecked request.
func GetScreenInfoUnchecked(c *xgb.XConn, Window xproto.Window) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"GetScreenInfo\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.Send(getScreenInfoRequest(op, Window))
}

// GetScreenInfoReply represents the data returned from a GetScreenInfo request.
type GetScreenInfoReply struct {
	Sequence        uint16 // sequence number of the request for this reply
	Length          uint32 // number of bytes in this reply
	Rotations       byte
	Root            xproto.Window
	Timestamp       xproto.Timestamp
	ConfigTimestamp xproto.Timestamp
	NSizes          uint16
	SizeID          uint16
	Rotation        uint16
	Rate            uint16
	NInfo           uint16
	// padding: 2 bytes
	Sizes []ScreenSize // size: internal.Pad4((int(NSizes) * 8))
	// alignment gap to multiple of 2
	Rates []RefreshRates // size: RefreshRatesListSize(Rates)
}

// Unmarshal reads a byte slice into a GetScreenInfoReply value.
func (v *GetScreenInfoReply) Unmarshal(buf []byte) error {
	if size := (((32 + internal.Pad4((int(v.NSizes) * 8))) + 2) + RefreshRatesListSize(v.Rates)); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetScreenInfoReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	v.Rotations = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Root = xproto.Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Timestamp = xproto.Timestamp(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.ConfigTimestamp = xproto.Timestamp(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.NSizes = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.SizeID = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Rotation = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Rate = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.NInfo = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	b += 2 // padding

	v.Sizes = make([]ScreenSize, v.NSizes)
	b += ScreenSizeReadList(buf[b:], v.Sizes)

	b = (b + 1) & ^1 // alignment gap

	v.Rates = make([]RefreshRates, (int(v.NInfo) - int(v.NSizes)))
	b += RefreshRatesReadList(buf[b:], v.Rates)

	return nil
}

// Write request to wire for GetScreenInfo
// getScreenInfoRequest writes a GetScreenInfo request to a byte slice.
func getScreenInfoRequest(opcode uint8, Window xproto.Window) []byte {
	const size = 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 5 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	return buf
}

// GetScreenResources sends a checked request.
func GetScreenResources(c *xgb.XConn, Window xproto.Window) (GetScreenResourcesReply, error) {
	var reply GetScreenResourcesReply
	op, ok := c.Ext("RANDR")
	if !ok {
		return reply, errors.New("cannot issue request \"GetScreenResources\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getScreenResourcesRequest(op, Window), &reply)
	return reply, err
}

// GetScreenResourcesUnchecked sends an unchecked request.
func GetScreenResourcesUnchecked(c *xgb.XConn, Window xproto.Window) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"GetScreenResources\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.Send(getScreenResourcesRequest(op, Window))
}

// GetScreenResourcesReply represents the data returned from a GetScreenResources request.
type GetScreenResourcesReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Timestamp       xproto.Timestamp
	ConfigTimestamp xproto.Timestamp
	NumCrtcs        uint16
	NumOutputs      uint16
	NumModes        uint16
	NamesLen        uint16
	// padding: 8 bytes
	Crtcs []Crtc // size: internal.Pad4((int(NumCrtcs) * 4))
	// alignment gap to multiple of 4
	Outputs []Output // size: internal.Pad4((int(NumOutputs) * 4))
	// alignment gap to multiple of 4
	Modes []ModeInfo // size: internal.Pad4((int(NumModes) * 32))
	Names []byte     // size: internal.Pad4((int(NamesLen) * 1))
}

// Unmarshal reads a byte slice into a GetScreenResourcesReply value.
func (v *GetScreenResourcesReply) Unmarshal(buf []byte) error {
	if size := ((((((32 + internal.Pad4((int(v.NumCrtcs) * 4))) + 4) + internal.Pad4((int(v.NumOutputs) * 4))) + 4) + internal.Pad4((int(v.NumModes) * 32))) + internal.Pad4((int(v.NamesLen) * 1))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetScreenResourcesReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Timestamp = xproto.Timestamp(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.ConfigTimestamp = xproto.Timestamp(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.NumCrtcs = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.NumOutputs = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.NumModes = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.NamesLen = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	b += 8 // padding

	v.Crtcs = make([]Crtc, v.NumCrtcs)
	for i := 0; i < int(v.NumCrtcs); i++ {
		v.Crtcs[i] = Crtc(binary.LittleEndian.Uint32(buf[b:]))
		b += 4
	}

	b = (b + 3) & ^3 // alignment gap

	v.Outputs = make([]Output, v.NumOutputs)
	for i := 0; i < int(v.NumOutputs); i++ {
		v.Outputs[i] = Output(binary.LittleEndian.Uint32(buf[b:]))
		b += 4
	}

	b = (b + 3) & ^3 // alignment gap

	v.Modes = make([]ModeInfo, v.NumModes)
	b += ModeInfoReadList(buf[b:], v.Modes)

	v.Names = make([]byte, v.NamesLen)
	copy(v.Names[:v.NamesLen], buf[b:])
	b += int(v.NamesLen)

	return nil
}

// Write request to wire for GetScreenResources
// getScreenResourcesRequest writes a GetScreenResources request to a byte slice.
func getScreenResourcesRequest(opcode uint8, Window xproto.Window) []byte {
	const size = 8
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

	return buf
}

// GetScreenResourcesCurrent sends a checked request.
func GetScreenResourcesCurrent(c *xgb.XConn, Window xproto.Window) (GetScreenResourcesCurrentReply, error) {
	var reply GetScreenResourcesCurrentReply
	op, ok := c.Ext("RANDR")
	if !ok {
		return reply, errors.New("cannot issue request \"GetScreenResourcesCurrent\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getScreenResourcesCurrentRequest(op, Window), &reply)
	return reply, err
}

// GetScreenResourcesCurrentUnchecked sends an unchecked request.
func GetScreenResourcesCurrentUnchecked(c *xgb.XConn, Window xproto.Window) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"GetScreenResourcesCurrent\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.Send(getScreenResourcesCurrentRequest(op, Window))
}

// GetScreenResourcesCurrentReply represents the data returned from a GetScreenResourcesCurrent request.
type GetScreenResourcesCurrentReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Timestamp       xproto.Timestamp
	ConfigTimestamp xproto.Timestamp
	NumCrtcs        uint16
	NumOutputs      uint16
	NumModes        uint16
	NamesLen        uint16
	// padding: 8 bytes
	Crtcs []Crtc // size: internal.Pad4((int(NumCrtcs) * 4))
	// alignment gap to multiple of 4
	Outputs []Output // size: internal.Pad4((int(NumOutputs) * 4))
	// alignment gap to multiple of 4
	Modes []ModeInfo // size: internal.Pad4((int(NumModes) * 32))
	Names []byte     // size: internal.Pad4((int(NamesLen) * 1))
}

// Unmarshal reads a byte slice into a GetScreenResourcesCurrentReply value.
func (v *GetScreenResourcesCurrentReply) Unmarshal(buf []byte) error {
	if size := ((((((32 + internal.Pad4((int(v.NumCrtcs) * 4))) + 4) + internal.Pad4((int(v.NumOutputs) * 4))) + 4) + internal.Pad4((int(v.NumModes) * 32))) + internal.Pad4((int(v.NamesLen) * 1))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetScreenResourcesCurrentReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Timestamp = xproto.Timestamp(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.ConfigTimestamp = xproto.Timestamp(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.NumCrtcs = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.NumOutputs = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.NumModes = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.NamesLen = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	b += 8 // padding

	v.Crtcs = make([]Crtc, v.NumCrtcs)
	for i := 0; i < int(v.NumCrtcs); i++ {
		v.Crtcs[i] = Crtc(binary.LittleEndian.Uint32(buf[b:]))
		b += 4
	}

	b = (b + 3) & ^3 // alignment gap

	v.Outputs = make([]Output, v.NumOutputs)
	for i := 0; i < int(v.NumOutputs); i++ {
		v.Outputs[i] = Output(binary.LittleEndian.Uint32(buf[b:]))
		b += 4
	}

	b = (b + 3) & ^3 // alignment gap

	v.Modes = make([]ModeInfo, v.NumModes)
	b += ModeInfoReadList(buf[b:], v.Modes)

	v.Names = make([]byte, v.NamesLen)
	copy(v.Names[:v.NamesLen], buf[b:])
	b += int(v.NamesLen)

	return nil
}

// Write request to wire for GetScreenResourcesCurrent
// getScreenResourcesCurrentRequest writes a GetScreenResourcesCurrent request to a byte slice.
func getScreenResourcesCurrentRequest(opcode uint8, Window xproto.Window) []byte {
	const size = 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 25 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	return buf
}

// GetScreenSizeRange sends a checked request.
func GetScreenSizeRange(c *xgb.XConn, Window xproto.Window) (GetScreenSizeRangeReply, error) {
	var reply GetScreenSizeRangeReply
	op, ok := c.Ext("RANDR")
	if !ok {
		return reply, errors.New("cannot issue request \"GetScreenSizeRange\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getScreenSizeRangeRequest(op, Window), &reply)
	return reply, err
}

// GetScreenSizeRangeUnchecked sends an unchecked request.
func GetScreenSizeRangeUnchecked(c *xgb.XConn, Window xproto.Window) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"GetScreenSizeRange\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.Send(getScreenSizeRangeRequest(op, Window))
}

// GetScreenSizeRangeReply represents the data returned from a GetScreenSizeRange request.
type GetScreenSizeRangeReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	MinWidth  uint16
	MinHeight uint16
	MaxWidth  uint16
	MaxHeight uint16
	// padding: 16 bytes
}

// Unmarshal reads a byte slice into a GetScreenSizeRangeReply value.
func (v *GetScreenSizeRangeReply) Unmarshal(buf []byte) error {
	const size = 32
	if len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetScreenSizeRangeReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.MinWidth = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.MinHeight = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.MaxWidth = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.MaxHeight = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	b += 16 // padding

	return nil
}

// Write request to wire for GetScreenSizeRange
// getScreenSizeRangeRequest writes a GetScreenSizeRange request to a byte slice.
func getScreenSizeRangeRequest(opcode uint8, Window xproto.Window) []byte {
	const size = 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 6 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	return buf
}

// ListOutputProperties sends a checked request.
func ListOutputProperties(c *xgb.XConn, Output Output) (ListOutputPropertiesReply, error) {
	var reply ListOutputPropertiesReply
	op, ok := c.Ext("RANDR")
	if !ok {
		return reply, errors.New("cannot issue request \"ListOutputProperties\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	err := c.SendRecv(listOutputPropertiesRequest(op, Output), &reply)
	return reply, err
}

// ListOutputPropertiesUnchecked sends an unchecked request.
func ListOutputPropertiesUnchecked(c *xgb.XConn, Output Output) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"ListOutputProperties\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.Send(listOutputPropertiesRequest(op, Output))
}

// ListOutputPropertiesReply represents the data returned from a ListOutputProperties request.
type ListOutputPropertiesReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	NumAtoms uint16
	// padding: 22 bytes
	Atoms []xproto.Atom // size: internal.Pad4((int(NumAtoms) * 4))
}

// Unmarshal reads a byte slice into a ListOutputPropertiesReply value.
func (v *ListOutputPropertiesReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.NumAtoms) * 4))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"ListOutputPropertiesReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.NumAtoms = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	b += 22 // padding

	v.Atoms = make([]xproto.Atom, v.NumAtoms)
	for i := 0; i < int(v.NumAtoms); i++ {
		v.Atoms[i] = xproto.Atom(binary.LittleEndian.Uint32(buf[b:]))
		b += 4
	}

	return nil
}

// Write request to wire for ListOutputProperties
// listOutputPropertiesRequest writes a ListOutputProperties request to a byte slice.
func listOutputPropertiesRequest(opcode uint8, Output Output) []byte {
	const size = 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 10 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Output))
	b += 4

	return buf
}

// ListProviderProperties sends a checked request.
func ListProviderProperties(c *xgb.XConn, Provider Provider) (ListProviderPropertiesReply, error) {
	var reply ListProviderPropertiesReply
	op, ok := c.Ext("RANDR")
	if !ok {
		return reply, errors.New("cannot issue request \"ListProviderProperties\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	err := c.SendRecv(listProviderPropertiesRequest(op, Provider), &reply)
	return reply, err
}

// ListProviderPropertiesUnchecked sends an unchecked request.
func ListProviderPropertiesUnchecked(c *xgb.XConn, Provider Provider) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"ListProviderProperties\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.Send(listProviderPropertiesRequest(op, Provider))
}

// ListProviderPropertiesReply represents the data returned from a ListProviderProperties request.
type ListProviderPropertiesReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	NumAtoms uint16
	// padding: 22 bytes
	Atoms []xproto.Atom // size: internal.Pad4((int(NumAtoms) * 4))
}

// Unmarshal reads a byte slice into a ListProviderPropertiesReply value.
func (v *ListProviderPropertiesReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.NumAtoms) * 4))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"ListProviderPropertiesReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.NumAtoms = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	b += 22 // padding

	v.Atoms = make([]xproto.Atom, v.NumAtoms)
	for i := 0; i < int(v.NumAtoms); i++ {
		v.Atoms[i] = xproto.Atom(binary.LittleEndian.Uint32(buf[b:]))
		b += 4
	}

	return nil
}

// Write request to wire for ListProviderProperties
// listProviderPropertiesRequest writes a ListProviderProperties request to a byte slice.
func listProviderPropertiesRequest(opcode uint8, Provider Provider) []byte {
	const size = 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 36 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Provider))
	b += 4

	return buf
}

// QueryOutputProperty sends a checked request.
func QueryOutputProperty(c *xgb.XConn, Output Output, Property xproto.Atom) (QueryOutputPropertyReply, error) {
	var reply QueryOutputPropertyReply
	op, ok := c.Ext("RANDR")
	if !ok {
		return reply, errors.New("cannot issue request \"QueryOutputProperty\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	err := c.SendRecv(queryOutputPropertyRequest(op, Output, Property), &reply)
	return reply, err
}

// QueryOutputPropertyUnchecked sends an unchecked request.
func QueryOutputPropertyUnchecked(c *xgb.XConn, Output Output, Property xproto.Atom) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"QueryOutputProperty\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.Send(queryOutputPropertyRequest(op, Output, Property))
}

// QueryOutputPropertyReply represents the data returned from a QueryOutputProperty request.
type QueryOutputPropertyReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Pending   bool
	Range     bool
	Immutable bool
	// padding: 21 bytes
	ValidValues []int32 // size: internal.Pad4((int(Length) * 4))
}

// Unmarshal reads a byte slice into a QueryOutputPropertyReply value.
func (v *QueryOutputPropertyReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.Length) * 4))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"QueryOutputPropertyReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Pending = (buf[b] == 1)
	b += 1

	v.Range = (buf[b] == 1)
	b += 1

	v.Immutable = (buf[b] == 1)
	b += 1

	b += 21 // padding

	v.ValidValues = make([]int32, v.Length)
	for i := 0; i < int(v.Length); i++ {
		v.ValidValues[i] = int32(binary.LittleEndian.Uint32(buf[b:]))
		b += 4
	}

	return nil
}

// Write request to wire for QueryOutputProperty
// queryOutputPropertyRequest writes a QueryOutputProperty request to a byte slice.
func queryOutputPropertyRequest(opcode uint8, Output Output, Property xproto.Atom) []byte {
	const size = 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 11 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Output))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Property))
	b += 4

	return buf
}

// QueryProviderProperty sends a checked request.
func QueryProviderProperty(c *xgb.XConn, Provider Provider, Property xproto.Atom) (QueryProviderPropertyReply, error) {
	var reply QueryProviderPropertyReply
	op, ok := c.Ext("RANDR")
	if !ok {
		return reply, errors.New("cannot issue request \"QueryProviderProperty\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	err := c.SendRecv(queryProviderPropertyRequest(op, Provider, Property), &reply)
	return reply, err
}

// QueryProviderPropertyUnchecked sends an unchecked request.
func QueryProviderPropertyUnchecked(c *xgb.XConn, Provider Provider, Property xproto.Atom) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"QueryProviderProperty\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.Send(queryProviderPropertyRequest(op, Provider, Property))
}

// QueryProviderPropertyReply represents the data returned from a QueryProviderProperty request.
type QueryProviderPropertyReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Pending   bool
	Range     bool
	Immutable bool
	// padding: 21 bytes
	ValidValues []int32 // size: internal.Pad4((int(Length) * 4))
}

// Unmarshal reads a byte slice into a QueryProviderPropertyReply value.
func (v *QueryProviderPropertyReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.Length) * 4))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"QueryProviderPropertyReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Pending = (buf[b] == 1)
	b += 1

	v.Range = (buf[b] == 1)
	b += 1

	v.Immutable = (buf[b] == 1)
	b += 1

	b += 21 // padding

	v.ValidValues = make([]int32, v.Length)
	for i := 0; i < int(v.Length); i++ {
		v.ValidValues[i] = int32(binary.LittleEndian.Uint32(buf[b:]))
		b += 4
	}

	return nil
}

// Write request to wire for QueryProviderProperty
// queryProviderPropertyRequest writes a QueryProviderProperty request to a byte slice.
func queryProviderPropertyRequest(opcode uint8, Provider Provider, Property xproto.Atom) []byte {
	const size = 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 37 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Provider))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Property))
	b += 4

	return buf
}

// QueryVersion sends a checked request.
func QueryVersion(c *xgb.XConn, MajorVersion uint32, MinorVersion uint32) (QueryVersionReply, error) {
	var reply QueryVersionReply
	op, ok := c.Ext("RANDR")
	if !ok {
		return reply, errors.New("cannot issue request \"QueryVersion\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	err := c.SendRecv(queryVersionRequest(op, MajorVersion, MinorVersion), &reply)
	return reply, err
}

// QueryVersionUnchecked sends an unchecked request.
func QueryVersionUnchecked(c *xgb.XConn, MajorVersion uint32, MinorVersion uint32) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"QueryVersion\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
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
	// padding: 16 bytes
}

// Unmarshal reads a byte slice into a QueryVersionReply value.
func (v *QueryVersionReply) Unmarshal(buf []byte) error {
	const size = 32
	if len(buf) < size {
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
func queryVersionRequest(opcode uint8, MajorVersion uint32, MinorVersion uint32) []byte {
	const size = 12
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

// SelectInput sends a checked request.
func SelectInput(c *xgb.XConn, Window xproto.Window, Enable uint16) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"SelectInput\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.SendRecv(selectInputRequest(op, Window, Enable), nil)
}

// SelectInputUnchecked sends an unchecked request.
func SelectInputUnchecked(c *xgb.XConn, Window xproto.Window, Enable uint16) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"SelectInput\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.Send(selectInputRequest(op, Window, Enable))
}

// Write request to wire for SelectInput
// selectInputRequest writes a SelectInput request to a byte slice.
func selectInputRequest(opcode uint8, Window xproto.Window, Enable uint16) []byte {
	const size = 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 4 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], Enable)
	b += 2

	b += 2 // padding

	return buf
}

// SetCrtcConfig sends a checked request.
func SetCrtcConfig(c *xgb.XConn, Crtc Crtc, Timestamp xproto.Timestamp, ConfigTimestamp xproto.Timestamp, X int16, Y int16, Mode Mode, Rotation uint16, Outputs []Output) (SetCrtcConfigReply, error) {
	var reply SetCrtcConfigReply
	op, ok := c.Ext("RANDR")
	if !ok {
		return reply, errors.New("cannot issue request \"SetCrtcConfig\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	err := c.SendRecv(setCrtcConfigRequest(op, Crtc, Timestamp, ConfigTimestamp, X, Y, Mode, Rotation, Outputs), &reply)
	return reply, err
}

// SetCrtcConfigUnchecked sends an unchecked request.
func SetCrtcConfigUnchecked(c *xgb.XConn, Crtc Crtc, Timestamp xproto.Timestamp, ConfigTimestamp xproto.Timestamp, X int16, Y int16, Mode Mode, Rotation uint16, Outputs []Output) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"SetCrtcConfig\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.Send(setCrtcConfigRequest(op, Crtc, Timestamp, ConfigTimestamp, X, Y, Mode, Rotation, Outputs))
}

// SetCrtcConfigReply represents the data returned from a SetCrtcConfig request.
type SetCrtcConfigReply struct {
	Sequence  uint16 // sequence number of the request for this reply
	Length    uint32 // number of bytes in this reply
	Status    byte
	Timestamp xproto.Timestamp
	// padding: 20 bytes
}

// Unmarshal reads a byte slice into a SetCrtcConfigReply value.
func (v *SetCrtcConfigReply) Unmarshal(buf []byte) error {
	const size = 32
	if len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"SetCrtcConfigReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	v.Status = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Timestamp = xproto.Timestamp(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	b += 20 // padding

	return nil
}

// Write request to wire for SetCrtcConfig
// setCrtcConfigRequest writes a SetCrtcConfig request to a byte slice.
func setCrtcConfigRequest(opcode uint8, Crtc Crtc, Timestamp xproto.Timestamp, ConfigTimestamp xproto.Timestamp, X int16, Y int16, Mode Mode, Rotation uint16, Outputs []Output) []byte {
	size := internal.Pad4((28 + internal.Pad4((len(Outputs) * 4))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 21 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Crtc))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Timestamp))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(ConfigTimestamp))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(X))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(Y))
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Mode))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], Rotation)
	b += 2

	b += 2 // padding

	for i := 0; i < int(len(Outputs)); i++ {
		binary.LittleEndian.PutUint32(buf[b:], uint32(Outputs[i]))
		b += 4
	}

	return buf
}

// SetCrtcGamma sends a checked request.
func SetCrtcGamma(c *xgb.XConn, Crtc Crtc, Size uint16, Red []uint16, Green []uint16, Blue []uint16) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"SetCrtcGamma\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.SendRecv(setCrtcGammaRequest(op, Crtc, Size, Red, Green, Blue), nil)
}

// SetCrtcGammaUnchecked sends an unchecked request.
func SetCrtcGammaUnchecked(c *xgb.XConn, Crtc Crtc, Size uint16, Red []uint16, Green []uint16, Blue []uint16) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"SetCrtcGamma\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.Send(setCrtcGammaRequest(op, Crtc, Size, Red, Green, Blue))
}

// Write request to wire for SetCrtcGamma
// setCrtcGammaRequest writes a SetCrtcGamma request to a byte slice.
func setCrtcGammaRequest(opcode uint8, Crtc Crtc, Size uint16, Red []uint16, Green []uint16, Blue []uint16) []byte {
	size := internal.Pad4((((((12 + internal.Pad4((int(Size) * 2))) + 2) + internal.Pad4((int(Size) * 2))) + 2) + internal.Pad4((int(Size) * 2))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 24 // request opcode
	b += 1

	blen := b
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Crtc))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], Size)
	b += 2

	b += 2 // padding

	for i := 0; i < int(Size); i++ {
		binary.LittleEndian.PutUint16(buf[b:], Red[i])
		b += 2
	}

	b = (b + 1) & ^1 // alignment gap

	for i := 0; i < int(Size); i++ {
		binary.LittleEndian.PutUint16(buf[b:], Green[i])
		b += 2
	}

	b = (b + 1) & ^1 // alignment gap

	for i := 0; i < int(Size); i++ {
		binary.LittleEndian.PutUint16(buf[b:], Blue[i])
		b += 2
	}

	b = internal.Pad4(b)
	binary.LittleEndian.PutUint16(buf[blen:], uint16(b/4)) // write request size in 4-byte units
	return buf[:b]
}

// SetCrtcTransform sends a checked request.
func SetCrtcTransform(c *xgb.XConn, Crtc Crtc, Transform render.Transform, FilterLen uint16, FilterName string, FilterParams []render.Fixed) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"SetCrtcTransform\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.SendRecv(setCrtcTransformRequest(op, Crtc, Transform, FilterLen, FilterName, FilterParams), nil)
}

// SetCrtcTransformUnchecked sends an unchecked request.
func SetCrtcTransformUnchecked(c *xgb.XConn, Crtc Crtc, Transform render.Transform, FilterLen uint16, FilterName string, FilterParams []render.Fixed) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"SetCrtcTransform\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.Send(setCrtcTransformRequest(op, Crtc, Transform, FilterLen, FilterName, FilterParams))
}

// Write request to wire for SetCrtcTransform
// setCrtcTransformRequest writes a SetCrtcTransform request to a byte slice.
func setCrtcTransformRequest(opcode uint8, Crtc Crtc, Transform render.Transform, FilterLen uint16, FilterName string, FilterParams []render.Fixed) []byte {
	size := internal.Pad4((((48 + internal.Pad4((int(FilterLen) * 1))) + 0) + internal.Pad4((len(FilterParams) * 4))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 26 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Crtc))
	b += 4

	{
		structBytes := Transform.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	binary.LittleEndian.PutUint16(buf[b:], FilterLen)
	b += 2

	b += 2 // padding

	copy(buf[b:], FilterName[:FilterLen])
	b += int(FilterLen)

	b += 0 // padding

	for i := 0; i < int(len(FilterParams)); i++ {
		binary.LittleEndian.PutUint32(buf[b:], uint32(FilterParams[i]))
		b += 4
	}

	return buf
}

// SetMonitor sends a checked request.
func SetMonitor(c *xgb.XConn, Window xproto.Window, Monitorinfo MonitorInfo) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"SetMonitor\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.SendRecv(setMonitorRequest(op, Window, Monitorinfo), nil)
}

// SetMonitorUnchecked sends an unchecked request.
func SetMonitorUnchecked(c *xgb.XConn, Window xproto.Window, Monitorinfo MonitorInfo) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"SetMonitor\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.Send(setMonitorRequest(op, Window, Monitorinfo))
}

// Write request to wire for SetMonitor
// setMonitorRequest writes a SetMonitor request to a byte slice.
func setMonitorRequest(opcode uint8, Window xproto.Window, Monitorinfo MonitorInfo) []byte {
	size := internal.Pad4((8 + (24 + internal.Pad4((int(Monitorinfo.NOutput) * 4)))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 43 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	{
		structBytes := Monitorinfo.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	return buf
}

// SetOutputPrimary sends a checked request.
func SetOutputPrimary(c *xgb.XConn, Window xproto.Window, Output Output) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"SetOutputPrimary\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.SendRecv(setOutputPrimaryRequest(op, Window, Output), nil)
}

// SetOutputPrimaryUnchecked sends an unchecked request.
func SetOutputPrimaryUnchecked(c *xgb.XConn, Window xproto.Window, Output Output) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"SetOutputPrimary\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.Send(setOutputPrimaryRequest(op, Window, Output))
}

// Write request to wire for SetOutputPrimary
// setOutputPrimaryRequest writes a SetOutputPrimary request to a byte slice.
func setOutputPrimaryRequest(opcode uint8, Window xproto.Window, Output Output) []byte {
	const size = 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 30 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Output))
	b += 4

	return buf
}

// SetPanning sends a checked request.
func SetPanning(c *xgb.XConn, Crtc Crtc, Timestamp xproto.Timestamp, Left uint16, Top uint16, Width uint16, Height uint16, TrackLeft uint16, TrackTop uint16, TrackWidth uint16, TrackHeight uint16, BorderLeft int16, BorderTop int16, BorderRight int16, BorderBottom int16) (SetPanningReply, error) {
	var reply SetPanningReply
	op, ok := c.Ext("RANDR")
	if !ok {
		return reply, errors.New("cannot issue request \"SetPanning\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	err := c.SendRecv(setPanningRequest(op, Crtc, Timestamp, Left, Top, Width, Height, TrackLeft, TrackTop, TrackWidth, TrackHeight, BorderLeft, BorderTop, BorderRight, BorderBottom), &reply)
	return reply, err
}

// SetPanningUnchecked sends an unchecked request.
func SetPanningUnchecked(c *xgb.XConn, Crtc Crtc, Timestamp xproto.Timestamp, Left uint16, Top uint16, Width uint16, Height uint16, TrackLeft uint16, TrackTop uint16, TrackWidth uint16, TrackHeight uint16, BorderLeft int16, BorderTop int16, BorderRight int16, BorderBottom int16) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"SetPanning\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.Send(setPanningRequest(op, Crtc, Timestamp, Left, Top, Width, Height, TrackLeft, TrackTop, TrackWidth, TrackHeight, BorderLeft, BorderTop, BorderRight, BorderBottom))
}

// SetPanningReply represents the data returned from a SetPanning request.
type SetPanningReply struct {
	Sequence  uint16 // sequence number of the request for this reply
	Length    uint32 // number of bytes in this reply
	Status    byte
	Timestamp xproto.Timestamp
}

// Unmarshal reads a byte slice into a SetPanningReply value.
func (v *SetPanningReply) Unmarshal(buf []byte) error {
	const size = 12
	if len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"SetPanningReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	v.Status = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Timestamp = xproto.Timestamp(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	return nil
}

// Write request to wire for SetPanning
// setPanningRequest writes a SetPanning request to a byte slice.
func setPanningRequest(opcode uint8, Crtc Crtc, Timestamp xproto.Timestamp, Left uint16, Top uint16, Width uint16, Height uint16, TrackLeft uint16, TrackTop uint16, TrackWidth uint16, TrackHeight uint16, BorderLeft int16, BorderTop int16, BorderRight int16, BorderBottom int16) []byte {
	const size = 36
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 29 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Crtc))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Timestamp))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], Left)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Top)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Width)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Height)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], TrackLeft)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], TrackTop)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], TrackWidth)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], TrackHeight)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(BorderLeft))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(BorderTop))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(BorderRight))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(BorderBottom))
	b += 2

	return buf
}

// SetProviderOffloadSink sends a checked request.
func SetProviderOffloadSink(c *xgb.XConn, Provider Provider, SinkProvider Provider, ConfigTimestamp xproto.Timestamp) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"SetProviderOffloadSink\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.SendRecv(setProviderOffloadSinkRequest(op, Provider, SinkProvider, ConfigTimestamp), nil)
}

// SetProviderOffloadSinkUnchecked sends an unchecked request.
func SetProviderOffloadSinkUnchecked(c *xgb.XConn, Provider Provider, SinkProvider Provider, ConfigTimestamp xproto.Timestamp) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"SetProviderOffloadSink\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.Send(setProviderOffloadSinkRequest(op, Provider, SinkProvider, ConfigTimestamp))
}

// Write request to wire for SetProviderOffloadSink
// setProviderOffloadSinkRequest writes a SetProviderOffloadSink request to a byte slice.
func setProviderOffloadSinkRequest(opcode uint8, Provider Provider, SinkProvider Provider, ConfigTimestamp xproto.Timestamp) []byte {
	const size = 16
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 34 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Provider))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(SinkProvider))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(ConfigTimestamp))
	b += 4

	return buf
}

// SetProviderOutputSource sends a checked request.
func SetProviderOutputSource(c *xgb.XConn, Provider Provider, SourceProvider Provider, ConfigTimestamp xproto.Timestamp) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"SetProviderOutputSource\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.SendRecv(setProviderOutputSourceRequest(op, Provider, SourceProvider, ConfigTimestamp), nil)
}

// SetProviderOutputSourceUnchecked sends an unchecked request.
func SetProviderOutputSourceUnchecked(c *xgb.XConn, Provider Provider, SourceProvider Provider, ConfigTimestamp xproto.Timestamp) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"SetProviderOutputSource\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.Send(setProviderOutputSourceRequest(op, Provider, SourceProvider, ConfigTimestamp))
}

// Write request to wire for SetProviderOutputSource
// setProviderOutputSourceRequest writes a SetProviderOutputSource request to a byte slice.
func setProviderOutputSourceRequest(opcode uint8, Provider Provider, SourceProvider Provider, ConfigTimestamp xproto.Timestamp) []byte {
	const size = 16
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 35 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Provider))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(SourceProvider))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(ConfigTimestamp))
	b += 4

	return buf
}

// SetScreenConfig sends a checked request.
func SetScreenConfig(c *xgb.XConn, Window xproto.Window, Timestamp xproto.Timestamp, ConfigTimestamp xproto.Timestamp, SizeID uint16, Rotation uint16, Rate uint16) (SetScreenConfigReply, error) {
	var reply SetScreenConfigReply
	op, ok := c.Ext("RANDR")
	if !ok {
		return reply, errors.New("cannot issue request \"SetScreenConfig\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	err := c.SendRecv(setScreenConfigRequest(op, Window, Timestamp, ConfigTimestamp, SizeID, Rotation, Rate), &reply)
	return reply, err
}

// SetScreenConfigUnchecked sends an unchecked request.
func SetScreenConfigUnchecked(c *xgb.XConn, Window xproto.Window, Timestamp xproto.Timestamp, ConfigTimestamp xproto.Timestamp, SizeID uint16, Rotation uint16, Rate uint16) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"SetScreenConfig\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.Send(setScreenConfigRequest(op, Window, Timestamp, ConfigTimestamp, SizeID, Rotation, Rate))
}

// SetScreenConfigReply represents the data returned from a SetScreenConfig request.
type SetScreenConfigReply struct {
	Sequence        uint16 // sequence number of the request for this reply
	Length          uint32 // number of bytes in this reply
	Status          byte
	NewTimestamp    xproto.Timestamp
	ConfigTimestamp xproto.Timestamp
	Root            xproto.Window
	SubpixelOrder   uint16
	// padding: 10 bytes
}

// Unmarshal reads a byte slice into a SetScreenConfigReply value.
func (v *SetScreenConfigReply) Unmarshal(buf []byte) error {
	const size = 32
	if len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"SetScreenConfigReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	v.Status = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.NewTimestamp = xproto.Timestamp(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.ConfigTimestamp = xproto.Timestamp(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Root = xproto.Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.SubpixelOrder = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	b += 10 // padding

	return nil
}

// Write request to wire for SetScreenConfig
// setScreenConfigRequest writes a SetScreenConfig request to a byte slice.
func setScreenConfigRequest(opcode uint8, Window xproto.Window, Timestamp xproto.Timestamp, ConfigTimestamp xproto.Timestamp, SizeID uint16, Rotation uint16, Rate uint16) []byte {
	const size = 24
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

	binary.LittleEndian.PutUint32(buf[b:], uint32(Timestamp))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(ConfigTimestamp))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], SizeID)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Rotation)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Rate)
	b += 2

	b += 2 // padding

	return buf
}

// SetScreenSize sends a checked request.
func SetScreenSize(c *xgb.XConn, Window xproto.Window, Width uint16, Height uint16, MmWidth uint32, MmHeight uint32) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"SetScreenSize\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.SendRecv(setScreenSizeRequest(op, Window, Width, Height, MmWidth, MmHeight), nil)
}

// SetScreenSizeUnchecked sends an unchecked request.
func SetScreenSizeUnchecked(c *xgb.XConn, Window xproto.Window, Width uint16, Height uint16, MmWidth uint32, MmHeight uint32) error {
	op, ok := c.Ext("RANDR")
	if !ok {
		return errors.New("cannot issue request \"SetScreenSize\" using the uninitialized extension \"RANDR\". randr.Register(xconn) must be called first.")
	}
	return c.Send(setScreenSizeRequest(op, Window, Width, Height, MmWidth, MmHeight))
}

// Write request to wire for SetScreenSize
// setScreenSizeRequest writes a SetScreenSize request to a byte slice.
func setScreenSizeRequest(opcode uint8, Window xproto.Window, Width uint16, Height uint16, MmWidth uint32, MmHeight uint32) []byte {
	const size = 20
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 7 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], Width)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Height)
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], MmWidth)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], MmHeight)
	b += 4

	return buf
}
