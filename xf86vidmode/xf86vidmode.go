// FILE GENERATED AUTOMATICALLY FROM "xf86vidmode.xml"
package xf86vidmode

import (
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/probakowski/go-xgb"
	"github.com/probakowski/go-xgb/internal"
	"github.com/probakowski/go-xgb/xproto"
)

const (
	// ExtName is the user-friendly name string of this X extension.
	ExtName = "XF86VidMode"

	// ExtXName is the name string this extension is known by to the X server.
	ExtXName = "XFree86-VidModeExtension"
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

// Register will query the X server for XF86VidMode extension support, and register relevant extension unmarshalers with the XConn.
func Register(xconn *xgb.XConn) error {
	// Query the X server for this extension
	reply, err := xproto.QueryExtension(xconn, uint16(len(ExtXName)), ExtXName)
	if err != nil {
		return fmt.Errorf("error querying X for \"XF86VidMode\": %w", err)
	} else if !reply.Present {
		return fmt.Errorf("no extension named \"XF86VidMode\" is known to the X server: reply=%+v", reply)
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

// BadBadClock is the error number for a BadBadClock.
const BadBadClock = 0

type BadClockError struct {
	Sequence uint16
	NiceName string
}

// UnmarshalBadClockError constructs a BadClockError value that implements xgb.Error from a byte slice.
func UnmarshalBadClockError(buf []byte) (xgb.XError, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"BadClockError\"", len(buf))
	}

	v := &BadClockError{}
	v.NiceName = "BadClock"

	b := 1 // skip error determinant
	b += 1 // don't read error number

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return v, nil
}

// SeqID returns the sequence id attached to the BadBadClock error.
// This is mostly used internally.
func (err *BadClockError) SeqID() uint16 {
	return err.Sequence
}

// BadID returns the 'BadValue' number if one exists for the BadBadClock error. If no bad value exists, 0 is returned.
func (err *BadClockError) BadID() uint32 {
	return 0
}

// Error returns a rudimentary string representation of the BadBadClock error.
func (err *BadClockError) Error() string {
	var buf strings.Builder

	buf.WriteString("BadBadClock{")
	buf.WriteString("NiceName: " + err.NiceName)
	buf.WriteByte(' ')
	buf.WriteString("Sequence: " + strconv.FormatUint(uint64(err.Sequence), 10))

	buf.WriteByte('}')

	return buf.String()
}

func init() { registerError(0, UnmarshalBadClockError) }

// BadBadHTimings is the error number for a BadBadHTimings.
const BadBadHTimings = 1

type BadHTimingsError struct {
	Sequence uint16
	NiceName string
}

// UnmarshalBadHTimingsError constructs a BadHTimingsError value that implements xgb.Error from a byte slice.
func UnmarshalBadHTimingsError(buf []byte) (xgb.XError, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"BadHTimingsError\"", len(buf))
	}

	v := &BadHTimingsError{}
	v.NiceName = "BadHTimings"

	b := 1 // skip error determinant
	b += 1 // don't read error number

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return v, nil
}

// SeqID returns the sequence id attached to the BadBadHTimings error.
// This is mostly used internally.
func (err *BadHTimingsError) SeqID() uint16 {
	return err.Sequence
}

// BadID returns the 'BadValue' number if one exists for the BadBadHTimings error. If no bad value exists, 0 is returned.
func (err *BadHTimingsError) BadID() uint32 {
	return 0
}

// Error returns a rudimentary string representation of the BadBadHTimings error.
func (err *BadHTimingsError) Error() string {
	var buf strings.Builder

	buf.WriteString("BadBadHTimings{")
	buf.WriteString("NiceName: " + err.NiceName)
	buf.WriteByte(' ')
	buf.WriteString("Sequence: " + strconv.FormatUint(uint64(err.Sequence), 10))

	buf.WriteByte('}')

	return buf.String()
}

func init() { registerError(1, UnmarshalBadHTimingsError) }

// BadBadVTimings is the error number for a BadBadVTimings.
const BadBadVTimings = 2

type BadVTimingsError struct {
	Sequence uint16
	NiceName string
}

// UnmarshalBadVTimingsError constructs a BadVTimingsError value that implements xgb.Error from a byte slice.
func UnmarshalBadVTimingsError(buf []byte) (xgb.XError, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"BadVTimingsError\"", len(buf))
	}

	v := &BadVTimingsError{}
	v.NiceName = "BadVTimings"

	b := 1 // skip error determinant
	b += 1 // don't read error number

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return v, nil
}

// SeqID returns the sequence id attached to the BadBadVTimings error.
// This is mostly used internally.
func (err *BadVTimingsError) SeqID() uint16 {
	return err.Sequence
}

// BadID returns the 'BadValue' number if one exists for the BadBadVTimings error. If no bad value exists, 0 is returned.
func (err *BadVTimingsError) BadID() uint32 {
	return 0
}

// Error returns a rudimentary string representation of the BadBadVTimings error.
func (err *BadVTimingsError) Error() string {
	var buf strings.Builder

	buf.WriteString("BadBadVTimings{")
	buf.WriteString("NiceName: " + err.NiceName)
	buf.WriteByte(' ')
	buf.WriteString("Sequence: " + strconv.FormatUint(uint64(err.Sequence), 10))

	buf.WriteByte('}')

	return buf.String()
}

func init() { registerError(2, UnmarshalBadVTimingsError) }

// BadClientNotLocal is the error number for a BadClientNotLocal.
const BadClientNotLocal = 5

type ClientNotLocalError struct {
	Sequence uint16
	NiceName string
}

// UnmarshalClientNotLocalError constructs a ClientNotLocalError value that implements xgb.Error from a byte slice.
func UnmarshalClientNotLocalError(buf []byte) (xgb.XError, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"ClientNotLocalError\"", len(buf))
	}

	v := &ClientNotLocalError{}
	v.NiceName = "ClientNotLocal"

	b := 1 // skip error determinant
	b += 1 // don't read error number

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return v, nil
}

// SeqID returns the sequence id attached to the BadClientNotLocal error.
// This is mostly used internally.
func (err *ClientNotLocalError) SeqID() uint16 {
	return err.Sequence
}

// BadID returns the 'BadValue' number if one exists for the BadClientNotLocal error. If no bad value exists, 0 is returned.
func (err *ClientNotLocalError) BadID() uint32 {
	return 0
}

// Error returns a rudimentary string representation of the BadClientNotLocal error.
func (err *ClientNotLocalError) Error() string {
	var buf strings.Builder

	buf.WriteString("BadClientNotLocal{")
	buf.WriteString("NiceName: " + err.NiceName)
	buf.WriteByte(' ')
	buf.WriteString("Sequence: " + strconv.FormatUint(uint64(err.Sequence), 10))

	buf.WriteByte('}')

	return buf.String()
}

func init() { registerError(5, UnmarshalClientNotLocalError) }

const (
	ClockFlagProgramable = 1
)

type Dotclock uint32

// BadExtensionDisabled is the error number for a BadExtensionDisabled.
const BadExtensionDisabled = 4

type ExtensionDisabledError struct {
	Sequence uint16
	NiceName string
}

// UnmarshalExtensionDisabledError constructs a ExtensionDisabledError value that implements xgb.Error from a byte slice.
func UnmarshalExtensionDisabledError(buf []byte) (xgb.XError, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"ExtensionDisabledError\"", len(buf))
	}

	v := &ExtensionDisabledError{}
	v.NiceName = "ExtensionDisabled"

	b := 1 // skip error determinant
	b += 1 // don't read error number

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return v, nil
}

// SeqID returns the sequence id attached to the BadExtensionDisabled error.
// This is mostly used internally.
func (err *ExtensionDisabledError) SeqID() uint16 {
	return err.Sequence
}

// BadID returns the 'BadValue' number if one exists for the BadExtensionDisabled error. If no bad value exists, 0 is returned.
func (err *ExtensionDisabledError) BadID() uint32 {
	return 0
}

// Error returns a rudimentary string representation of the BadExtensionDisabled error.
func (err *ExtensionDisabledError) Error() string {
	var buf strings.Builder

	buf.WriteString("BadExtensionDisabled{")
	buf.WriteString("NiceName: " + err.NiceName)
	buf.WriteByte(' ')
	buf.WriteString("Sequence: " + strconv.FormatUint(uint64(err.Sequence), 10))

	buf.WriteByte('}')

	return buf.String()
}

func init() { registerError(4, UnmarshalExtensionDisabledError) }

const (
	ModeFlagPositiveHsync = 1
	ModeFlagNegativeHsync = 2
	ModeFlagPositiveVsync = 4
	ModeFlagNegativeVsync = 8
	ModeFlagInterlace     = 16
	ModeFlagCompositeSync = 32
	ModeFlagPositiveCsync = 64
	ModeFlagNegativeCsync = 128
	ModeFlagHSkew         = 256
	ModeFlagBroadcast     = 512
	ModeFlagPixmux        = 1024
	ModeFlagDoubleClock   = 2048
	ModeFlagHalfClock     = 4096
)

type ModeInfo struct {
	Dotclock   Dotclock
	Hdisplay   uint16
	Hsyncstart uint16
	Hsyncend   uint16
	Htotal     uint16
	Hskew      uint32
	Vdisplay   uint16
	Vsyncstart uint16
	Vsyncend   uint16
	Vtotal     uint16
	// padding: 4 bytes
	Flags uint32
	// padding: 12 bytes
	Privsize uint32
}

// ModeInfoRead reads a byte slice into a ModeInfo value.
func ModeInfoRead(buf []byte, v *ModeInfo) int {
	b := 0

	v.Dotclock = Dotclock(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Hdisplay = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Hsyncstart = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Hsyncend = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Htotal = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Hskew = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.Vdisplay = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Vsyncstart = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Vsyncend = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Vtotal = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	b += 4 // padding

	v.Flags = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 12 // padding

	v.Privsize = binary.LittleEndian.Uint32(buf[b:])
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
	buf := make([]byte, 48)
	b := 0

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Dotclock))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], v.Hdisplay)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Hsyncstart)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Hsyncend)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Htotal)
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], v.Hskew)
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], v.Vdisplay)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Vsyncstart)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Vsyncend)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Vtotal)
	b += 2

	b += 4 // padding

	binary.LittleEndian.PutUint32(buf[b:], v.Flags)
	b += 4

	b += 12 // padding

	binary.LittleEndian.PutUint32(buf[b:], v.Privsize)
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

// BadModeUnsuitable is the error number for a BadModeUnsuitable.
const BadModeUnsuitable = 3

type ModeUnsuitableError struct {
	Sequence uint16
	NiceName string
}

// UnmarshalModeUnsuitableError constructs a ModeUnsuitableError value that implements xgb.Error from a byte slice.
func UnmarshalModeUnsuitableError(buf []byte) (xgb.XError, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"ModeUnsuitableError\"", len(buf))
	}

	v := &ModeUnsuitableError{}
	v.NiceName = "ModeUnsuitable"

	b := 1 // skip error determinant
	b += 1 // don't read error number

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return v, nil
}

// SeqID returns the sequence id attached to the BadModeUnsuitable error.
// This is mostly used internally.
func (err *ModeUnsuitableError) SeqID() uint16 {
	return err.Sequence
}

// BadID returns the 'BadValue' number if one exists for the BadModeUnsuitable error. If no bad value exists, 0 is returned.
func (err *ModeUnsuitableError) BadID() uint32 {
	return 0
}

// Error returns a rudimentary string representation of the BadModeUnsuitable error.
func (err *ModeUnsuitableError) Error() string {
	var buf strings.Builder

	buf.WriteString("BadModeUnsuitable{")
	buf.WriteString("NiceName: " + err.NiceName)
	buf.WriteByte(' ')
	buf.WriteString("Sequence: " + strconv.FormatUint(uint64(err.Sequence), 10))

	buf.WriteByte('}')

	return buf.String()
}

func init() { registerError(3, UnmarshalModeUnsuitableError) }

const (
	PermissionRead  = 1
	PermissionWrite = 2
)

type Syncrange uint32

// BadZoomLocked is the error number for a BadZoomLocked.
const BadZoomLocked = 6

type ZoomLockedError struct {
	Sequence uint16
	NiceName string
}

// UnmarshalZoomLockedError constructs a ZoomLockedError value that implements xgb.Error from a byte slice.
func UnmarshalZoomLockedError(buf []byte) (xgb.XError, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"ZoomLockedError\"", len(buf))
	}

	v := &ZoomLockedError{}
	v.NiceName = "ZoomLocked"

	b := 1 // skip error determinant
	b += 1 // don't read error number

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return v, nil
}

// SeqID returns the sequence id attached to the BadZoomLocked error.
// This is mostly used internally.
func (err *ZoomLockedError) SeqID() uint16 {
	return err.Sequence
}

// BadID returns the 'BadValue' number if one exists for the BadZoomLocked error. If no bad value exists, 0 is returned.
func (err *ZoomLockedError) BadID() uint32 {
	return 0
}

// Error returns a rudimentary string representation of the BadZoomLocked error.
func (err *ZoomLockedError) Error() string {
	var buf strings.Builder

	buf.WriteString("BadZoomLocked{")
	buf.WriteString("NiceName: " + err.NiceName)
	buf.WriteByte(' ')
	buf.WriteString("Sequence: " + strconv.FormatUint(uint64(err.Sequence), 10))

	buf.WriteByte('}')

	return buf.String()
}

func init() { registerError(6, UnmarshalZoomLockedError) }

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

// AddModeLine sends a checked request.
func AddModeLine(c *xgb.XConn, Screen uint32, Dotclock Dotclock, Hdisplay uint16, Hsyncstart uint16, Hsyncend uint16, Htotal uint16, Hskew uint16, Vdisplay uint16, Vsyncstart uint16, Vsyncend uint16, Vtotal uint16, Flags uint32, Privsize uint32, AfterDotclock Dotclock, AfterHdisplay uint16, AfterHsyncstart uint16, AfterHsyncend uint16, AfterHtotal uint16, AfterHskew uint16, AfterVdisplay uint16, AfterVsyncstart uint16, AfterVsyncend uint16, AfterVtotal uint16, AfterFlags uint32, Private []byte) error {
	op, ok := c.Ext("XFree86-VidModeExtension")
	if !ok {
		return errors.New("cannot issue request \"AddModeLine\" using the uninitialized extension \"XFree86-VidModeExtension\". xf86vidmode.Register(xconn) must be called first.")
	}
	return c.SendRecv(addModeLineRequest(op, Screen, Dotclock, Hdisplay, Hsyncstart, Hsyncend, Htotal, Hskew, Vdisplay, Vsyncstart, Vsyncend, Vtotal, Flags, Privsize, AfterDotclock, AfterHdisplay, AfterHsyncstart, AfterHsyncend, AfterHtotal, AfterHskew, AfterVdisplay, AfterVsyncstart, AfterVsyncend, AfterVtotal, AfterFlags, Private), nil)
}

// AddModeLineUnchecked sends an unchecked request.
func AddModeLineUnchecked(c *xgb.XConn, Screen uint32, Dotclock Dotclock, Hdisplay uint16, Hsyncstart uint16, Hsyncend uint16, Htotal uint16, Hskew uint16, Vdisplay uint16, Vsyncstart uint16, Vsyncend uint16, Vtotal uint16, Flags uint32, Privsize uint32, AfterDotclock Dotclock, AfterHdisplay uint16, AfterHsyncstart uint16, AfterHsyncend uint16, AfterHtotal uint16, AfterHskew uint16, AfterVdisplay uint16, AfterVsyncstart uint16, AfterVsyncend uint16, AfterVtotal uint16, AfterFlags uint32, Private []byte) error {
	op, ok := c.Ext("XFree86-VidModeExtension")
	if !ok {
		return errors.New("cannot issue request \"AddModeLine\" using the uninitialized extension \"XFree86-VidModeExtension\". xf86vidmode.Register(xconn) must be called first.")
	}
	return c.Send(addModeLineRequest(op, Screen, Dotclock, Hdisplay, Hsyncstart, Hsyncend, Htotal, Hskew, Vdisplay, Vsyncstart, Vsyncend, Vtotal, Flags, Privsize, AfterDotclock, AfterHdisplay, AfterHsyncstart, AfterHsyncend, AfterHtotal, AfterHskew, AfterVdisplay, AfterVsyncstart, AfterVsyncend, AfterVtotal, AfterFlags, Private))
}

// Write request to wire for AddModeLine
// addModeLineRequest writes a AddModeLine request to a byte slice.
func addModeLineRequest(opcode uint8, Screen uint32, Dotclock Dotclock, Hdisplay uint16, Hsyncstart uint16, Hsyncend uint16, Htotal uint16, Hskew uint16, Vdisplay uint16, Vsyncstart uint16, Vsyncend uint16, Vtotal uint16, Flags uint32, Privsize uint32, AfterDotclock Dotclock, AfterHdisplay uint16, AfterHsyncstart uint16, AfterHsyncend uint16, AfterHtotal uint16, AfterHskew uint16, AfterVdisplay uint16, AfterVsyncstart uint16, AfterVsyncend uint16, AfterVtotal uint16, AfterFlags uint32, Private []byte) []byte {
	size := internal.Pad4((92 + internal.Pad4((int(Privsize) * 1))))
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

	binary.LittleEndian.PutUint32(buf[b:], uint32(Dotclock))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], Hdisplay)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Hsyncstart)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Hsyncend)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Htotal)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Hskew)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Vdisplay)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Vsyncstart)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Vsyncend)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Vtotal)
	b += 2

	b += 2 // padding

	binary.LittleEndian.PutUint32(buf[b:], Flags)
	b += 4

	b += 12 // padding

	binary.LittleEndian.PutUint32(buf[b:], Privsize)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(AfterDotclock))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], AfterHdisplay)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], AfterHsyncstart)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], AfterHsyncend)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], AfterHtotal)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], AfterHskew)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], AfterVdisplay)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], AfterVsyncstart)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], AfterVsyncend)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], AfterVtotal)
	b += 2

	b += 2 // padding

	binary.LittleEndian.PutUint32(buf[b:], AfterFlags)
	b += 4

	b += 12 // padding

	copy(buf[b:], Private[:Privsize])
	b += int(Privsize)

	return buf
}

// DeleteModeLine sends a checked request.
func DeleteModeLine(c *xgb.XConn, Screen uint32, Dotclock Dotclock, Hdisplay uint16, Hsyncstart uint16, Hsyncend uint16, Htotal uint16, Hskew uint16, Vdisplay uint16, Vsyncstart uint16, Vsyncend uint16, Vtotal uint16, Flags uint32, Privsize uint32, Private []byte) error {
	op, ok := c.Ext("XFree86-VidModeExtension")
	if !ok {
		return errors.New("cannot issue request \"DeleteModeLine\" using the uninitialized extension \"XFree86-VidModeExtension\". xf86vidmode.Register(xconn) must be called first.")
	}
	return c.SendRecv(deleteModeLineRequest(op, Screen, Dotclock, Hdisplay, Hsyncstart, Hsyncend, Htotal, Hskew, Vdisplay, Vsyncstart, Vsyncend, Vtotal, Flags, Privsize, Private), nil)
}

// DeleteModeLineUnchecked sends an unchecked request.
func DeleteModeLineUnchecked(c *xgb.XConn, Screen uint32, Dotclock Dotclock, Hdisplay uint16, Hsyncstart uint16, Hsyncend uint16, Htotal uint16, Hskew uint16, Vdisplay uint16, Vsyncstart uint16, Vsyncend uint16, Vtotal uint16, Flags uint32, Privsize uint32, Private []byte) error {
	op, ok := c.Ext("XFree86-VidModeExtension")
	if !ok {
		return errors.New("cannot issue request \"DeleteModeLine\" using the uninitialized extension \"XFree86-VidModeExtension\". xf86vidmode.Register(xconn) must be called first.")
	}
	return c.Send(deleteModeLineRequest(op, Screen, Dotclock, Hdisplay, Hsyncstart, Hsyncend, Htotal, Hskew, Vdisplay, Vsyncstart, Vsyncend, Vtotal, Flags, Privsize, Private))
}

// Write request to wire for DeleteModeLine
// deleteModeLineRequest writes a DeleteModeLine request to a byte slice.
func deleteModeLineRequest(opcode uint8, Screen uint32, Dotclock Dotclock, Hdisplay uint16, Hsyncstart uint16, Hsyncend uint16, Htotal uint16, Hskew uint16, Vdisplay uint16, Vsyncstart uint16, Vsyncend uint16, Vtotal uint16, Flags uint32, Privsize uint32, Private []byte) []byte {
	size := internal.Pad4((52 + internal.Pad4((int(Privsize) * 1))))
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

	binary.LittleEndian.PutUint32(buf[b:], uint32(Dotclock))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], Hdisplay)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Hsyncstart)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Hsyncend)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Htotal)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Hskew)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Vdisplay)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Vsyncstart)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Vsyncend)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Vtotal)
	b += 2

	b += 2 // padding

	binary.LittleEndian.PutUint32(buf[b:], Flags)
	b += 4

	b += 12 // padding

	binary.LittleEndian.PutUint32(buf[b:], Privsize)
	b += 4

	copy(buf[b:], Private[:Privsize])
	b += int(Privsize)

	return buf
}

// GetAllModeLines sends a checked request.
func GetAllModeLines(c *xgb.XConn, Screen uint16) (GetAllModeLinesReply, error) {
	var reply GetAllModeLinesReply
	op, ok := c.Ext("XFree86-VidModeExtension")
	if !ok {
		return reply, errors.New("cannot issue request \"GetAllModeLines\" using the uninitialized extension \"XFree86-VidModeExtension\". xf86vidmode.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getAllModeLinesRequest(op, Screen), &reply)
	return reply, err
}

// GetAllModeLinesUnchecked sends an unchecked request.
func GetAllModeLinesUnchecked(c *xgb.XConn, Screen uint16) error {
	op, ok := c.Ext("XFree86-VidModeExtension")
	if !ok {
		return errors.New("cannot issue request \"GetAllModeLines\" using the uninitialized extension \"XFree86-VidModeExtension\". xf86vidmode.Register(xconn) must be called first.")
	}
	return c.Send(getAllModeLinesRequest(op, Screen))
}

// GetAllModeLinesReply represents the data returned from a GetAllModeLines request.
type GetAllModeLinesReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Modecount uint32
	// padding: 20 bytes
	Modeinfo []ModeInfo // size: internal.Pad4((int(Modecount) * 48))
}

// Unmarshal reads a byte slice into a GetAllModeLinesReply value.
func (v *GetAllModeLinesReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.Modecount) * 48))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetAllModeLinesReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Modecount = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 20 // padding

	v.Modeinfo = make([]ModeInfo, v.Modecount)
	b += ModeInfoReadList(buf[b:], v.Modeinfo)

	return nil
}

// Write request to wire for GetAllModeLines
// getAllModeLinesRequest writes a GetAllModeLines request to a byte slice.
func getAllModeLinesRequest(opcode uint8, Screen uint16) []byte {
	const size = 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 6 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Screen)
	b += 2

	b += 2 // padding

	return buf
}

// GetDotClocks sends a checked request.
func GetDotClocks(c *xgb.XConn, Screen uint16) (GetDotClocksReply, error) {
	var reply GetDotClocksReply
	op, ok := c.Ext("XFree86-VidModeExtension")
	if !ok {
		return reply, errors.New("cannot issue request \"GetDotClocks\" using the uninitialized extension \"XFree86-VidModeExtension\". xf86vidmode.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getDotClocksRequest(op, Screen), &reply)
	return reply, err
}

// GetDotClocksUnchecked sends an unchecked request.
func GetDotClocksUnchecked(c *xgb.XConn, Screen uint16) error {
	op, ok := c.Ext("XFree86-VidModeExtension")
	if !ok {
		return errors.New("cannot issue request \"GetDotClocks\" using the uninitialized extension \"XFree86-VidModeExtension\". xf86vidmode.Register(xconn) must be called first.")
	}
	return c.Send(getDotClocksRequest(op, Screen))
}

// GetDotClocksReply represents the data returned from a GetDotClocks request.
type GetDotClocksReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Flags     uint32
	Clocks    uint32
	Maxclocks uint32
	// padding: 12 bytes
	Clock []uint32 // size: internal.Pad4((((1 - (int(Flags) & 1)) * int(Clocks)) * 4))
}

// Unmarshal reads a byte slice into a GetDotClocksReply value.
func (v *GetDotClocksReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((((1 - (int(v.Flags) & 1)) * int(v.Clocks)) * 4))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetDotClocksReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Flags = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.Clocks = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.Maxclocks = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 12 // padding

	v.Clock = make([]uint32, ((1 - (int(v.Flags) & 1)) * int(v.Clocks)))
	for i := 0; i < int(((1 - (int(v.Flags) & 1)) * int(v.Clocks))); i++ {
		v.Clock[i] = binary.LittleEndian.Uint32(buf[b:])
		b += 4
	}

	return nil
}

// Write request to wire for GetDotClocks
// getDotClocksRequest writes a GetDotClocks request to a byte slice.
func getDotClocksRequest(opcode uint8, Screen uint16) []byte {
	const size = 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 13 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Screen)
	b += 2

	b += 2 // padding

	return buf
}

// GetGamma sends a checked request.
func GetGamma(c *xgb.XConn, Screen uint16) (GetGammaReply, error) {
	var reply GetGammaReply
	op, ok := c.Ext("XFree86-VidModeExtension")
	if !ok {
		return reply, errors.New("cannot issue request \"GetGamma\" using the uninitialized extension \"XFree86-VidModeExtension\". xf86vidmode.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getGammaRequest(op, Screen), &reply)
	return reply, err
}

// GetGammaUnchecked sends an unchecked request.
func GetGammaUnchecked(c *xgb.XConn, Screen uint16) error {
	op, ok := c.Ext("XFree86-VidModeExtension")
	if !ok {
		return errors.New("cannot issue request \"GetGamma\" using the uninitialized extension \"XFree86-VidModeExtension\". xf86vidmode.Register(xconn) must be called first.")
	}
	return c.Send(getGammaRequest(op, Screen))
}

// GetGammaReply represents the data returned from a GetGamma request.
type GetGammaReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Red   uint32
	Green uint32
	Blue  uint32
	// padding: 12 bytes
}

// Unmarshal reads a byte slice into a GetGammaReply value.
func (v *GetGammaReply) Unmarshal(buf []byte) error {
	const size = 32
	if len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetGammaReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Red = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.Green = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.Blue = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 12 // padding

	return nil
}

// Write request to wire for GetGamma
// getGammaRequest writes a GetGamma request to a byte slice.
func getGammaRequest(opcode uint8, Screen uint16) []byte {
	const size = 32
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 16 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Screen)
	b += 2

	b += 26 // padding

	return buf
}

// GetGammaRamp sends a checked request.
func GetGammaRamp(c *xgb.XConn, Screen uint16, Size uint16) (GetGammaRampReply, error) {
	var reply GetGammaRampReply
	op, ok := c.Ext("XFree86-VidModeExtension")
	if !ok {
		return reply, errors.New("cannot issue request \"GetGammaRamp\" using the uninitialized extension \"XFree86-VidModeExtension\". xf86vidmode.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getGammaRampRequest(op, Screen, Size), &reply)
	return reply, err
}

// GetGammaRampUnchecked sends an unchecked request.
func GetGammaRampUnchecked(c *xgb.XConn, Screen uint16, Size uint16) error {
	op, ok := c.Ext("XFree86-VidModeExtension")
	if !ok {
		return errors.New("cannot issue request \"GetGammaRamp\" using the uninitialized extension \"XFree86-VidModeExtension\". xf86vidmode.Register(xconn) must be called first.")
	}
	return c.Send(getGammaRampRequest(op, Screen, Size))
}

// GetGammaRampReply represents the data returned from a GetGammaRamp request.
type GetGammaRampReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Size uint16
	// padding: 22 bytes
	Red []uint16 // size: internal.Pad4((((int(Size) + 1) & -2) * 2))
	// alignment gap to multiple of 2
	Green []uint16 // size: internal.Pad4((((int(Size) + 1) & -2) * 2))
	// alignment gap to multiple of 2
	Blue []uint16 // size: internal.Pad4((((int(Size) + 1) & -2) * 2))
}

// Unmarshal reads a byte slice into a GetGammaRampReply value.
func (v *GetGammaRampReply) Unmarshal(buf []byte) error {
	if size := (((((32 + internal.Pad4((((int(v.Size) + 1) & -2) * 2))) + 2) + internal.Pad4((((int(v.Size) + 1) & -2) * 2))) + 2) + internal.Pad4((((int(v.Size) + 1) & -2) * 2))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetGammaRampReply\": have=%d need=%d", len(buf), size)
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

	v.Red = make([]uint16, ((int(v.Size) + 1) & -2))
	for i := 0; i < int(((int(v.Size) + 1) & -2)); i++ {
		v.Red[i] = binary.LittleEndian.Uint16(buf[b:])
		b += 2
	}

	b = (b + 1) & ^1 // alignment gap

	v.Green = make([]uint16, ((int(v.Size) + 1) & -2))
	for i := 0; i < int(((int(v.Size) + 1) & -2)); i++ {
		v.Green[i] = binary.LittleEndian.Uint16(buf[b:])
		b += 2
	}

	b = (b + 1) & ^1 // alignment gap

	v.Blue = make([]uint16, ((int(v.Size) + 1) & -2))
	for i := 0; i < int(((int(v.Size) + 1) & -2)); i++ {
		v.Blue[i] = binary.LittleEndian.Uint16(buf[b:])
		b += 2
	}

	return nil
}

// Write request to wire for GetGammaRamp
// getGammaRampRequest writes a GetGammaRamp request to a byte slice.
func getGammaRampRequest(opcode uint8, Screen uint16, Size uint16) []byte {
	const size = 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 17 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Screen)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Size)
	b += 2

	return buf
}

// GetGammaRampSize sends a checked request.
func GetGammaRampSize(c *xgb.XConn, Screen uint16) (GetGammaRampSizeReply, error) {
	var reply GetGammaRampSizeReply
	op, ok := c.Ext("XFree86-VidModeExtension")
	if !ok {
		return reply, errors.New("cannot issue request \"GetGammaRampSize\" using the uninitialized extension \"XFree86-VidModeExtension\". xf86vidmode.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getGammaRampSizeRequest(op, Screen), &reply)
	return reply, err
}

// GetGammaRampSizeUnchecked sends an unchecked request.
func GetGammaRampSizeUnchecked(c *xgb.XConn, Screen uint16) error {
	op, ok := c.Ext("XFree86-VidModeExtension")
	if !ok {
		return errors.New("cannot issue request \"GetGammaRampSize\" using the uninitialized extension \"XFree86-VidModeExtension\". xf86vidmode.Register(xconn) must be called first.")
	}
	return c.Send(getGammaRampSizeRequest(op, Screen))
}

// GetGammaRampSizeReply represents the data returned from a GetGammaRampSize request.
type GetGammaRampSizeReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Size uint16
	// padding: 22 bytes
}

// Unmarshal reads a byte slice into a GetGammaRampSizeReply value.
func (v *GetGammaRampSizeReply) Unmarshal(buf []byte) error {
	const size = 32
	if len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetGammaRampSizeReply\": have=%d need=%d", len(buf), size)
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

// Write request to wire for GetGammaRampSize
// getGammaRampSizeRequest writes a GetGammaRampSize request to a byte slice.
func getGammaRampSizeRequest(opcode uint8, Screen uint16) []byte {
	const size = 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 19 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Screen)
	b += 2

	b += 2 // padding

	return buf
}

// GetModeLine sends a checked request.
func GetModeLine(c *xgb.XConn, Screen uint16) (GetModeLineReply, error) {
	var reply GetModeLineReply
	op, ok := c.Ext("XFree86-VidModeExtension")
	if !ok {
		return reply, errors.New("cannot issue request \"GetModeLine\" using the uninitialized extension \"XFree86-VidModeExtension\". xf86vidmode.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getModeLineRequest(op, Screen), &reply)
	return reply, err
}

// GetModeLineUnchecked sends an unchecked request.
func GetModeLineUnchecked(c *xgb.XConn, Screen uint16) error {
	op, ok := c.Ext("XFree86-VidModeExtension")
	if !ok {
		return errors.New("cannot issue request \"GetModeLine\" using the uninitialized extension \"XFree86-VidModeExtension\". xf86vidmode.Register(xconn) must be called first.")
	}
	return c.Send(getModeLineRequest(op, Screen))
}

// GetModeLineReply represents the data returned from a GetModeLine request.
type GetModeLineReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Dotclock   Dotclock
	Hdisplay   uint16
	Hsyncstart uint16
	Hsyncend   uint16
	Htotal     uint16
	Hskew      uint16
	Vdisplay   uint16
	Vsyncstart uint16
	Vsyncend   uint16
	Vtotal     uint16
	// padding: 2 bytes
	Flags uint32
	// padding: 12 bytes
	Privsize uint32
	Private  []byte // size: internal.Pad4((int(Privsize) * 1))
}

// Unmarshal reads a byte slice into a GetModeLineReply value.
func (v *GetModeLineReply) Unmarshal(buf []byte) error {
	if size := (52 + internal.Pad4((int(v.Privsize) * 1))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetModeLineReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Dotclock = Dotclock(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Hdisplay = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Hsyncstart = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Hsyncend = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Htotal = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Hskew = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Vdisplay = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Vsyncstart = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Vsyncend = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Vtotal = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	b += 2 // padding

	v.Flags = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 12 // padding

	v.Privsize = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.Private = make([]byte, v.Privsize)
	copy(v.Private[:v.Privsize], buf[b:])
	b += int(v.Privsize)

	return nil
}

// Write request to wire for GetModeLine
// getModeLineRequest writes a GetModeLine request to a byte slice.
func getModeLineRequest(opcode uint8, Screen uint16) []byte {
	const size = 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 1 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Screen)
	b += 2

	b += 2 // padding

	return buf
}

// GetMonitor sends a checked request.
func GetMonitor(c *xgb.XConn, Screen uint16) (GetMonitorReply, error) {
	var reply GetMonitorReply
	op, ok := c.Ext("XFree86-VidModeExtension")
	if !ok {
		return reply, errors.New("cannot issue request \"GetMonitor\" using the uninitialized extension \"XFree86-VidModeExtension\". xf86vidmode.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getMonitorRequest(op, Screen), &reply)
	return reply, err
}

// GetMonitorUnchecked sends an unchecked request.
func GetMonitorUnchecked(c *xgb.XConn, Screen uint16) error {
	op, ok := c.Ext("XFree86-VidModeExtension")
	if !ok {
		return errors.New("cannot issue request \"GetMonitor\" using the uninitialized extension \"XFree86-VidModeExtension\". xf86vidmode.Register(xconn) must be called first.")
	}
	return c.Send(getMonitorRequest(op, Screen))
}

// GetMonitorReply represents the data returned from a GetMonitor request.
type GetMonitorReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	VendorLength byte
	ModelLength  byte
	NumHsync     byte
	NumVsync     byte
	// padding: 20 bytes
	Hsync []Syncrange // size: internal.Pad4((int(NumHsync) * 4))
	// alignment gap to multiple of 4
	Vsync        []Syncrange // size: internal.Pad4((int(NumVsync) * 4))
	Vendor       string      // size: internal.Pad4((int(VendorLength) * 1))
	AlignmentPad []byte      // size: internal.Pad4(((((int(VendorLength) + 3) & -4) - int(VendorLength)) * 1))
	Model        string      // size: internal.Pad4((int(ModelLength) * 1))
}

// Unmarshal reads a byte slice into a GetMonitorReply value.
func (v *GetMonitorReply) Unmarshal(buf []byte) error {
	if size := ((((((32 + internal.Pad4((int(v.NumHsync) * 4))) + 4) + internal.Pad4((int(v.NumVsync) * 4))) + internal.Pad4((int(v.VendorLength) * 1))) + internal.Pad4(((((int(v.VendorLength) + 3) & -4) - int(v.VendorLength)) * 1))) + internal.Pad4((int(v.ModelLength) * 1))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetMonitorReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.VendorLength = buf[b]
	b += 1

	v.ModelLength = buf[b]
	b += 1

	v.NumHsync = buf[b]
	b += 1

	v.NumVsync = buf[b]
	b += 1

	b += 20 // padding

	v.Hsync = make([]Syncrange, v.NumHsync)
	for i := 0; i < int(v.NumHsync); i++ {
		v.Hsync[i] = Syncrange(binary.LittleEndian.Uint32(buf[b:]))
		b += 4
	}

	b = (b + 3) & ^3 // alignment gap

	v.Vsync = make([]Syncrange, v.NumVsync)
	for i := 0; i < int(v.NumVsync); i++ {
		v.Vsync[i] = Syncrange(binary.LittleEndian.Uint32(buf[b:]))
		b += 4
	}

	{
		byteString := make([]byte, v.VendorLength)
		copy(byteString[:v.VendorLength], buf[b:])
		v.Vendor = string(byteString)
		b += int(v.VendorLength)
	}

	v.AlignmentPad = make([]byte, (((int(v.VendorLength) + 3) & -4) - int(v.VendorLength)))
	copy(v.AlignmentPad[:(((int(v.VendorLength)+3)&-4)-int(v.VendorLength))], buf[b:])
	b += int((((int(v.VendorLength) + 3) & -4) - int(v.VendorLength)))

	{
		byteString := make([]byte, v.ModelLength)
		copy(byteString[:v.ModelLength], buf[b:])
		v.Model = string(byteString)
		b += int(v.ModelLength)
	}

	return nil
}

// Write request to wire for GetMonitor
// getMonitorRequest writes a GetMonitor request to a byte slice.
func getMonitorRequest(opcode uint8, Screen uint16) []byte {
	const size = 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 4 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Screen)
	b += 2

	b += 2 // padding

	return buf
}

// GetPermissions sends a checked request.
func GetPermissions(c *xgb.XConn, Screen uint16) (GetPermissionsReply, error) {
	var reply GetPermissionsReply
	op, ok := c.Ext("XFree86-VidModeExtension")
	if !ok {
		return reply, errors.New("cannot issue request \"GetPermissions\" using the uninitialized extension \"XFree86-VidModeExtension\". xf86vidmode.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getPermissionsRequest(op, Screen), &reply)
	return reply, err
}

// GetPermissionsUnchecked sends an unchecked request.
func GetPermissionsUnchecked(c *xgb.XConn, Screen uint16) error {
	op, ok := c.Ext("XFree86-VidModeExtension")
	if !ok {
		return errors.New("cannot issue request \"GetPermissions\" using the uninitialized extension \"XFree86-VidModeExtension\". xf86vidmode.Register(xconn) must be called first.")
	}
	return c.Send(getPermissionsRequest(op, Screen))
}

// GetPermissionsReply represents the data returned from a GetPermissions request.
type GetPermissionsReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Permissions uint32
	// padding: 20 bytes
}

// Unmarshal reads a byte slice into a GetPermissionsReply value.
func (v *GetPermissionsReply) Unmarshal(buf []byte) error {
	const size = 32
	if len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetPermissionsReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Permissions = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 20 // padding

	return nil
}

// Write request to wire for GetPermissions
// getPermissionsRequest writes a GetPermissions request to a byte slice.
func getPermissionsRequest(opcode uint8, Screen uint16) []byte {
	const size = 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 20 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Screen)
	b += 2

	b += 2 // padding

	return buf
}

// GetViewPort sends a checked request.
func GetViewPort(c *xgb.XConn, Screen uint16) (GetViewPortReply, error) {
	var reply GetViewPortReply
	op, ok := c.Ext("XFree86-VidModeExtension")
	if !ok {
		return reply, errors.New("cannot issue request \"GetViewPort\" using the uninitialized extension \"XFree86-VidModeExtension\". xf86vidmode.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getViewPortRequest(op, Screen), &reply)
	return reply, err
}

// GetViewPortUnchecked sends an unchecked request.
func GetViewPortUnchecked(c *xgb.XConn, Screen uint16) error {
	op, ok := c.Ext("XFree86-VidModeExtension")
	if !ok {
		return errors.New("cannot issue request \"GetViewPort\" using the uninitialized extension \"XFree86-VidModeExtension\". xf86vidmode.Register(xconn) must be called first.")
	}
	return c.Send(getViewPortRequest(op, Screen))
}

// GetViewPortReply represents the data returned from a GetViewPort request.
type GetViewPortReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	X uint32
	Y uint32
	// padding: 16 bytes
}

// Unmarshal reads a byte slice into a GetViewPortReply value.
func (v *GetViewPortReply) Unmarshal(buf []byte) error {
	const size = 32
	if len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetViewPortReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.X = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.Y = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 16 // padding

	return nil
}

// Write request to wire for GetViewPort
// getViewPortRequest writes a GetViewPort request to a byte slice.
func getViewPortRequest(opcode uint8, Screen uint16) []byte {
	const size = 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 11 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Screen)
	b += 2

	b += 2 // padding

	return buf
}

// LockModeSwitch sends a checked request.
func LockModeSwitch(c *xgb.XConn, Screen uint16, Lock uint16) error {
	op, ok := c.Ext("XFree86-VidModeExtension")
	if !ok {
		return errors.New("cannot issue request \"LockModeSwitch\" using the uninitialized extension \"XFree86-VidModeExtension\". xf86vidmode.Register(xconn) must be called first.")
	}
	return c.SendRecv(lockModeSwitchRequest(op, Screen, Lock), nil)
}

// LockModeSwitchUnchecked sends an unchecked request.
func LockModeSwitchUnchecked(c *xgb.XConn, Screen uint16, Lock uint16) error {
	op, ok := c.Ext("XFree86-VidModeExtension")
	if !ok {
		return errors.New("cannot issue request \"LockModeSwitch\" using the uninitialized extension \"XFree86-VidModeExtension\". xf86vidmode.Register(xconn) must be called first.")
	}
	return c.Send(lockModeSwitchRequest(op, Screen, Lock))
}

// Write request to wire for LockModeSwitch
// lockModeSwitchRequest writes a LockModeSwitch request to a byte slice.
func lockModeSwitchRequest(opcode uint8, Screen uint16, Lock uint16) []byte {
	const size = 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 5 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Screen)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Lock)
	b += 2

	return buf
}

// ModModeLine sends a checked request.
func ModModeLine(c *xgb.XConn, Screen uint32, Hdisplay uint16, Hsyncstart uint16, Hsyncend uint16, Htotal uint16, Hskew uint16, Vdisplay uint16, Vsyncstart uint16, Vsyncend uint16, Vtotal uint16, Flags uint32, Privsize uint32, Private []byte) error {
	op, ok := c.Ext("XFree86-VidModeExtension")
	if !ok {
		return errors.New("cannot issue request \"ModModeLine\" using the uninitialized extension \"XFree86-VidModeExtension\". xf86vidmode.Register(xconn) must be called first.")
	}
	return c.SendRecv(modModeLineRequest(op, Screen, Hdisplay, Hsyncstart, Hsyncend, Htotal, Hskew, Vdisplay, Vsyncstart, Vsyncend, Vtotal, Flags, Privsize, Private), nil)
}

// ModModeLineUnchecked sends an unchecked request.
func ModModeLineUnchecked(c *xgb.XConn, Screen uint32, Hdisplay uint16, Hsyncstart uint16, Hsyncend uint16, Htotal uint16, Hskew uint16, Vdisplay uint16, Vsyncstart uint16, Vsyncend uint16, Vtotal uint16, Flags uint32, Privsize uint32, Private []byte) error {
	op, ok := c.Ext("XFree86-VidModeExtension")
	if !ok {
		return errors.New("cannot issue request \"ModModeLine\" using the uninitialized extension \"XFree86-VidModeExtension\". xf86vidmode.Register(xconn) must be called first.")
	}
	return c.Send(modModeLineRequest(op, Screen, Hdisplay, Hsyncstart, Hsyncend, Htotal, Hskew, Vdisplay, Vsyncstart, Vsyncend, Vtotal, Flags, Privsize, Private))
}

// Write request to wire for ModModeLine
// modModeLineRequest writes a ModModeLine request to a byte slice.
func modModeLineRequest(opcode uint8, Screen uint32, Hdisplay uint16, Hsyncstart uint16, Hsyncend uint16, Htotal uint16, Hskew uint16, Vdisplay uint16, Vsyncstart uint16, Vsyncend uint16, Vtotal uint16, Flags uint32, Privsize uint32, Private []byte) []byte {
	size := internal.Pad4((48 + internal.Pad4((int(Privsize) * 1))))
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

	binary.LittleEndian.PutUint16(buf[b:], Hdisplay)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Hsyncstart)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Hsyncend)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Htotal)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Hskew)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Vdisplay)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Vsyncstart)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Vsyncend)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Vtotal)
	b += 2

	b += 2 // padding

	binary.LittleEndian.PutUint32(buf[b:], Flags)
	b += 4

	b += 12 // padding

	binary.LittleEndian.PutUint32(buf[b:], Privsize)
	b += 4

	copy(buf[b:], Private[:Privsize])
	b += int(Privsize)

	return buf
}

// QueryVersion sends a checked request.
func QueryVersion(c *xgb.XConn) (QueryVersionReply, error) {
	var reply QueryVersionReply
	op, ok := c.Ext("XFree86-VidModeExtension")
	if !ok {
		return reply, errors.New("cannot issue request \"QueryVersion\" using the uninitialized extension \"XFree86-VidModeExtension\". xf86vidmode.Register(xconn) must be called first.")
	}
	err := c.SendRecv(queryVersionRequest(op), &reply)
	return reply, err
}

// QueryVersionUnchecked sends an unchecked request.
func QueryVersionUnchecked(c *xgb.XConn) error {
	op, ok := c.Ext("XFree86-VidModeExtension")
	if !ok {
		return errors.New("cannot issue request \"QueryVersion\" using the uninitialized extension \"XFree86-VidModeExtension\". xf86vidmode.Register(xconn) must be called first.")
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
	const size = 12
	if len(buf) < size {
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

// SetClientVersion sends a checked request.
func SetClientVersion(c *xgb.XConn, Major uint16, Minor uint16) error {
	op, ok := c.Ext("XFree86-VidModeExtension")
	if !ok {
		return errors.New("cannot issue request \"SetClientVersion\" using the uninitialized extension \"XFree86-VidModeExtension\". xf86vidmode.Register(xconn) must be called first.")
	}
	return c.SendRecv(setClientVersionRequest(op, Major, Minor), nil)
}

// SetClientVersionUnchecked sends an unchecked request.
func SetClientVersionUnchecked(c *xgb.XConn, Major uint16, Minor uint16) error {
	op, ok := c.Ext("XFree86-VidModeExtension")
	if !ok {
		return errors.New("cannot issue request \"SetClientVersion\" using the uninitialized extension \"XFree86-VidModeExtension\". xf86vidmode.Register(xconn) must be called first.")
	}
	return c.Send(setClientVersionRequest(op, Major, Minor))
}

// Write request to wire for SetClientVersion
// setClientVersionRequest writes a SetClientVersion request to a byte slice.
func setClientVersionRequest(opcode uint8, Major uint16, Minor uint16) []byte {
	const size = 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 14 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Major)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Minor)
	b += 2

	return buf
}

// SetGamma sends a checked request.
func SetGamma(c *xgb.XConn, Screen uint16, Red uint32, Green uint32, Blue uint32) error {
	op, ok := c.Ext("XFree86-VidModeExtension")
	if !ok {
		return errors.New("cannot issue request \"SetGamma\" using the uninitialized extension \"XFree86-VidModeExtension\". xf86vidmode.Register(xconn) must be called first.")
	}
	return c.SendRecv(setGammaRequest(op, Screen, Red, Green, Blue), nil)
}

// SetGammaUnchecked sends an unchecked request.
func SetGammaUnchecked(c *xgb.XConn, Screen uint16, Red uint32, Green uint32, Blue uint32) error {
	op, ok := c.Ext("XFree86-VidModeExtension")
	if !ok {
		return errors.New("cannot issue request \"SetGamma\" using the uninitialized extension \"XFree86-VidModeExtension\". xf86vidmode.Register(xconn) must be called first.")
	}
	return c.Send(setGammaRequest(op, Screen, Red, Green, Blue))
}

// Write request to wire for SetGamma
// setGammaRequest writes a SetGamma request to a byte slice.
func setGammaRequest(opcode uint8, Screen uint16, Red uint32, Green uint32, Blue uint32) []byte {
	const size = 32
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 15 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Screen)
	b += 2

	b += 2 // padding

	binary.LittleEndian.PutUint32(buf[b:], Red)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], Green)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], Blue)
	b += 4

	b += 12 // padding

	return buf
}

// SetGammaRamp sends a checked request.
func SetGammaRamp(c *xgb.XConn, Screen uint16, Size uint16, Red []uint16, Green []uint16, Blue []uint16) error {
	op, ok := c.Ext("XFree86-VidModeExtension")
	if !ok {
		return errors.New("cannot issue request \"SetGammaRamp\" using the uninitialized extension \"XFree86-VidModeExtension\". xf86vidmode.Register(xconn) must be called first.")
	}
	return c.SendRecv(setGammaRampRequest(op, Screen, Size, Red, Green, Blue), nil)
}

// SetGammaRampUnchecked sends an unchecked request.
func SetGammaRampUnchecked(c *xgb.XConn, Screen uint16, Size uint16, Red []uint16, Green []uint16, Blue []uint16) error {
	op, ok := c.Ext("XFree86-VidModeExtension")
	if !ok {
		return errors.New("cannot issue request \"SetGammaRamp\" using the uninitialized extension \"XFree86-VidModeExtension\". xf86vidmode.Register(xconn) must be called first.")
	}
	return c.Send(setGammaRampRequest(op, Screen, Size, Red, Green, Blue))
}

// Write request to wire for SetGammaRamp
// setGammaRampRequest writes a SetGammaRamp request to a byte slice.
func setGammaRampRequest(opcode uint8, Screen uint16, Size uint16, Red []uint16, Green []uint16, Blue []uint16) []byte {
	size := internal.Pad4((((((8 + internal.Pad4((((int(Size) + 1) & -2) * 2))) + 2) + internal.Pad4((((int(Size) + 1) & -2) * 2))) + 2) + internal.Pad4((((int(Size) + 1) & -2) * 2))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 18 // request opcode
	b += 1

	blen := b
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Screen)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Size)
	b += 2

	for i := 0; i < int(((int(Size) + 1) & -2)); i++ {
		binary.LittleEndian.PutUint16(buf[b:], Red[i])
		b += 2
	}

	b = (b + 1) & ^1 // alignment gap

	for i := 0; i < int(((int(Size) + 1) & -2)); i++ {
		binary.LittleEndian.PutUint16(buf[b:], Green[i])
		b += 2
	}

	b = (b + 1) & ^1 // alignment gap

	for i := 0; i < int(((int(Size) + 1) & -2)); i++ {
		binary.LittleEndian.PutUint16(buf[b:], Blue[i])
		b += 2
	}

	b = internal.Pad4(b)
	binary.LittleEndian.PutUint16(buf[blen:], uint16(b/4)) // write request size in 4-byte units
	return buf[:b]
}

// SetViewPort sends a checked request.
func SetViewPort(c *xgb.XConn, Screen uint16, X uint32, Y uint32) error {
	op, ok := c.Ext("XFree86-VidModeExtension")
	if !ok {
		return errors.New("cannot issue request \"SetViewPort\" using the uninitialized extension \"XFree86-VidModeExtension\". xf86vidmode.Register(xconn) must be called first.")
	}
	return c.SendRecv(setViewPortRequest(op, Screen, X, Y), nil)
}

// SetViewPortUnchecked sends an unchecked request.
func SetViewPortUnchecked(c *xgb.XConn, Screen uint16, X uint32, Y uint32) error {
	op, ok := c.Ext("XFree86-VidModeExtension")
	if !ok {
		return errors.New("cannot issue request \"SetViewPort\" using the uninitialized extension \"XFree86-VidModeExtension\". xf86vidmode.Register(xconn) must be called first.")
	}
	return c.Send(setViewPortRequest(op, Screen, X, Y))
}

// Write request to wire for SetViewPort
// setViewPortRequest writes a SetViewPort request to a byte slice.
func setViewPortRequest(opcode uint8, Screen uint16, X uint32, Y uint32) []byte {
	const size = 16
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 12 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Screen)
	b += 2

	b += 2 // padding

	binary.LittleEndian.PutUint32(buf[b:], X)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], Y)
	b += 4

	return buf
}

// SwitchMode sends a checked request.
func SwitchMode(c *xgb.XConn, Screen uint16, Zoom uint16) error {
	op, ok := c.Ext("XFree86-VidModeExtension")
	if !ok {
		return errors.New("cannot issue request \"SwitchMode\" using the uninitialized extension \"XFree86-VidModeExtension\". xf86vidmode.Register(xconn) must be called first.")
	}
	return c.SendRecv(switchModeRequest(op, Screen, Zoom), nil)
}

// SwitchModeUnchecked sends an unchecked request.
func SwitchModeUnchecked(c *xgb.XConn, Screen uint16, Zoom uint16) error {
	op, ok := c.Ext("XFree86-VidModeExtension")
	if !ok {
		return errors.New("cannot issue request \"SwitchMode\" using the uninitialized extension \"XFree86-VidModeExtension\". xf86vidmode.Register(xconn) must be called first.")
	}
	return c.Send(switchModeRequest(op, Screen, Zoom))
}

// Write request to wire for SwitchMode
// switchModeRequest writes a SwitchMode request to a byte slice.
func switchModeRequest(opcode uint8, Screen uint16, Zoom uint16) []byte {
	const size = 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 3 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Screen)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Zoom)
	b += 2

	return buf
}

// SwitchToMode sends a checked request.
func SwitchToMode(c *xgb.XConn, Screen uint32, Dotclock Dotclock, Hdisplay uint16, Hsyncstart uint16, Hsyncend uint16, Htotal uint16, Hskew uint16, Vdisplay uint16, Vsyncstart uint16, Vsyncend uint16, Vtotal uint16, Flags uint32, Privsize uint32, Private []byte) error {
	op, ok := c.Ext("XFree86-VidModeExtension")
	if !ok {
		return errors.New("cannot issue request \"SwitchToMode\" using the uninitialized extension \"XFree86-VidModeExtension\". xf86vidmode.Register(xconn) must be called first.")
	}
	return c.SendRecv(switchToModeRequest(op, Screen, Dotclock, Hdisplay, Hsyncstart, Hsyncend, Htotal, Hskew, Vdisplay, Vsyncstart, Vsyncend, Vtotal, Flags, Privsize, Private), nil)
}

// SwitchToModeUnchecked sends an unchecked request.
func SwitchToModeUnchecked(c *xgb.XConn, Screen uint32, Dotclock Dotclock, Hdisplay uint16, Hsyncstart uint16, Hsyncend uint16, Htotal uint16, Hskew uint16, Vdisplay uint16, Vsyncstart uint16, Vsyncend uint16, Vtotal uint16, Flags uint32, Privsize uint32, Private []byte) error {
	op, ok := c.Ext("XFree86-VidModeExtension")
	if !ok {
		return errors.New("cannot issue request \"SwitchToMode\" using the uninitialized extension \"XFree86-VidModeExtension\". xf86vidmode.Register(xconn) must be called first.")
	}
	return c.Send(switchToModeRequest(op, Screen, Dotclock, Hdisplay, Hsyncstart, Hsyncend, Htotal, Hskew, Vdisplay, Vsyncstart, Vsyncend, Vtotal, Flags, Privsize, Private))
}

// Write request to wire for SwitchToMode
// switchToModeRequest writes a SwitchToMode request to a byte slice.
func switchToModeRequest(opcode uint8, Screen uint32, Dotclock Dotclock, Hdisplay uint16, Hsyncstart uint16, Hsyncend uint16, Htotal uint16, Hskew uint16, Vdisplay uint16, Vsyncstart uint16, Vsyncend uint16, Vtotal uint16, Flags uint32, Privsize uint32, Private []byte) []byte {
	size := internal.Pad4((52 + internal.Pad4((int(Privsize) * 1))))
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

	binary.LittleEndian.PutUint32(buf[b:], uint32(Dotclock))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], Hdisplay)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Hsyncstart)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Hsyncend)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Htotal)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Hskew)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Vdisplay)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Vsyncstart)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Vsyncend)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Vtotal)
	b += 2

	b += 2 // padding

	binary.LittleEndian.PutUint32(buf[b:], Flags)
	b += 4

	b += 12 // padding

	binary.LittleEndian.PutUint32(buf[b:], Privsize)
	b += 4

	copy(buf[b:], Private[:Privsize])
	b += int(Privsize)

	return buf
}

// ValidateModeLine sends a checked request.
func ValidateModeLine(c *xgb.XConn, Screen uint32, Dotclock Dotclock, Hdisplay uint16, Hsyncstart uint16, Hsyncend uint16, Htotal uint16, Hskew uint16, Vdisplay uint16, Vsyncstart uint16, Vsyncend uint16, Vtotal uint16, Flags uint32, Privsize uint32, Private []byte) (ValidateModeLineReply, error) {
	var reply ValidateModeLineReply
	op, ok := c.Ext("XFree86-VidModeExtension")
	if !ok {
		return reply, errors.New("cannot issue request \"ValidateModeLine\" using the uninitialized extension \"XFree86-VidModeExtension\". xf86vidmode.Register(xconn) must be called first.")
	}
	err := c.SendRecv(validateModeLineRequest(op, Screen, Dotclock, Hdisplay, Hsyncstart, Hsyncend, Htotal, Hskew, Vdisplay, Vsyncstart, Vsyncend, Vtotal, Flags, Privsize, Private), &reply)
	return reply, err
}

// ValidateModeLineUnchecked sends an unchecked request.
func ValidateModeLineUnchecked(c *xgb.XConn, Screen uint32, Dotclock Dotclock, Hdisplay uint16, Hsyncstart uint16, Hsyncend uint16, Htotal uint16, Hskew uint16, Vdisplay uint16, Vsyncstart uint16, Vsyncend uint16, Vtotal uint16, Flags uint32, Privsize uint32, Private []byte) error {
	op, ok := c.Ext("XFree86-VidModeExtension")
	if !ok {
		return errors.New("cannot issue request \"ValidateModeLine\" using the uninitialized extension \"XFree86-VidModeExtension\". xf86vidmode.Register(xconn) must be called first.")
	}
	return c.Send(validateModeLineRequest(op, Screen, Dotclock, Hdisplay, Hsyncstart, Hsyncend, Htotal, Hskew, Vdisplay, Vsyncstart, Vsyncend, Vtotal, Flags, Privsize, Private))
}

// ValidateModeLineReply represents the data returned from a ValidateModeLine request.
type ValidateModeLineReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Status uint32
	// padding: 20 bytes
}

// Unmarshal reads a byte slice into a ValidateModeLineReply value.
func (v *ValidateModeLineReply) Unmarshal(buf []byte) error {
	const size = 32
	if len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"ValidateModeLineReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Status = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 20 // padding

	return nil
}

// Write request to wire for ValidateModeLine
// validateModeLineRequest writes a ValidateModeLine request to a byte slice.
func validateModeLineRequest(opcode uint8, Screen uint32, Dotclock Dotclock, Hdisplay uint16, Hsyncstart uint16, Hsyncend uint16, Htotal uint16, Hskew uint16, Vdisplay uint16, Vsyncstart uint16, Vsyncend uint16, Vtotal uint16, Flags uint32, Privsize uint32, Private []byte) []byte {
	size := internal.Pad4((52 + internal.Pad4((int(Privsize) * 1))))
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

	binary.LittleEndian.PutUint32(buf[b:], uint32(Dotclock))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], Hdisplay)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Hsyncstart)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Hsyncend)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Htotal)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Hskew)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Vdisplay)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Vsyncstart)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Vsyncend)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Vtotal)
	b += 2

	b += 2 // padding

	binary.LittleEndian.PutUint32(buf[b:], Flags)
	b += 4

	b += 12 // padding

	binary.LittleEndian.PutUint32(buf[b:], Privsize)
	b += 4

	copy(buf[b:], Private[:Privsize])
	b += int(Privsize)

	return buf
}
