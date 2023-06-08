// FILE GENERATED AUTOMATICALLY FROM "xfixes.xml"
package xfixes

import (
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"codeberg.org/gruf/go-xgb"
	"codeberg.org/gruf/go-xgb/internal"
	"codeberg.org/gruf/go-xgb/render"
	"codeberg.org/gruf/go-xgb/shape"
	"codeberg.org/gruf/go-xgb/xproto"
)

const (
	// ExtName is the user-friendly name string of this X extension.
	ExtName = "XFixes"

	// ExtXName is the name string this extension is known by to the X server.
	ExtXName = "XFIXES"
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
		return fmt.Errorf("error querying X for \"XFixes\": %w", err)
	} else if !reply.Present {
		return fmt.Errorf("no extension named \"XFixes\" is known to the X server: reply=%+v", reply)
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

// BadBadRegion is the error number for a BadBadRegion.
const BadBadRegion = 0

type BadRegionError struct {
	Sequence uint16
	NiceName string
}

// UnmarshalBadRegionError constructs a BadRegionError value that implements xgb.Error from a byte slice.
func UnmarshalBadRegionError(buf []byte) (xgb.XError, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"BadRegionError\"", len(buf))
	}

	v := &BadRegionError{}
	v.NiceName = "BadRegion"

	b := 1 // skip error determinant
	b += 1 // don't read error number

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return v, nil
}

// SeqID returns the sequence id attached to the BadBadRegion error.
// This is mostly used internally.
func (err *BadRegionError) SeqID() uint16 {
	return err.Sequence
}

// BadID returns the 'BadValue' number if one exists for the BadBadRegion error. If no bad value exists, 0 is returned.
func (err *BadRegionError) BadID() uint32 {
	return 0
}

// Error returns a rudimentary string representation of the BadBadRegion error.
func (err *BadRegionError) Error() string {
	var buf strings.Builder

	buf.WriteString("BadBadRegion{")
	buf.WriteString("NiceName: " + err.NiceName)
	buf.WriteByte(' ')
	buf.WriteString("Sequence: " + strconv.FormatUint(uint64(err.Sequence), 10))

	buf.WriteByte('}')

	return buf.String()
}

func init() { registerError(0, UnmarshalBadRegionError) }

type Barrier uint32

func NewBarrierID(c *xgb.XConn) Barrier {
	id := c.NewXID()
	return Barrier(id)
}

const (
	BarrierDirectionsPositiveX = 1
	BarrierDirectionsPositiveY = 2
	BarrierDirectionsNegativeX = 4
	BarrierDirectionsNegativeY = 8
)

const (
	ClientDisconnectFlagsDefault   = 0
	ClientDisconnectFlagsTerminate = 1
)

// CursorNotify is the event number for a CursorNotifyEvent.
const CursorNotify = 1

type CursorNotifyEvent struct {
	Sequence     uint16
	Subtype      byte
	Window       xproto.Window
	CursorSerial uint32
	Timestamp    xproto.Timestamp
	Name         xproto.Atom
	// padding: 12 bytes
}

// UnmarshalCursorNotifyEvent constructs a CursorNotifyEvent value that implements xgb.Event from a byte slice.
func UnmarshalCursorNotifyEvent(buf []byte) (xgb.XEvent, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"CursorNotifyEvent\"", len(buf))
	}

	v := &CursorNotifyEvent{}
	b := 1 // don't read event number

	v.Subtype = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Window = xproto.Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.CursorSerial = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.Timestamp = xproto.Timestamp(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Name = xproto.Atom(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	b += 12 // padding

	return v, nil
}

// Bytes writes a CursorNotifyEvent value to a byte slice.
func (v *CursorNotifyEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 1
	b += 1

	buf[b] = v.Subtype
	b += 1

	b += 2 // skip sequence number

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Window))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], v.CursorSerial)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Timestamp))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Name))
	b += 4

	b += 12 // padding

	return buf
}

// SeqID returns the sequence id attached to the CursorNotify event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v *CursorNotifyEvent) SeqID() uint16 {
	return v.Sequence
}

func init() { registerEvent(1, UnmarshalCursorNotifyEvent) }

const (
	CursorNotifyDisplayCursor = 0
)

const (
	CursorNotifyMaskDisplayCursor = 1
)

type Region uint32

func NewRegionID(c *xgb.XConn) Region {
	id := c.NewXID()
	return Region(id)
}

const (
	RegionNone = 0
)

const (
	SaveSetMappingMap   = 0
	SaveSetMappingUnmap = 1
)

const (
	SaveSetModeInsert = 0
	SaveSetModeDelete = 1
)

const (
	SaveSetTargetNearest = 0
	SaveSetTargetRoot    = 1
)

const (
	SelectionEventSetSelectionOwner      = 0
	SelectionEventSelectionWindowDestroy = 1
	SelectionEventSelectionClientClose   = 2
)

const (
	SelectionEventMaskSetSelectionOwner      = 1
	SelectionEventMaskSelectionWindowDestroy = 2
	SelectionEventMaskSelectionClientClose   = 4
)

// SelectionNotify is the event number for a SelectionNotifyEvent.
const SelectionNotify = 0

type SelectionNotifyEvent struct {
	Sequence           uint16
	Subtype            byte
	Window             xproto.Window
	Owner              xproto.Window
	Selection          xproto.Atom
	Timestamp          xproto.Timestamp
	SelectionTimestamp xproto.Timestamp
	// padding: 8 bytes
}

// UnmarshalSelectionNotifyEvent constructs a SelectionNotifyEvent value that implements xgb.Event from a byte slice.
func UnmarshalSelectionNotifyEvent(buf []byte) (xgb.XEvent, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"SelectionNotifyEvent\"", len(buf))
	}

	v := &SelectionNotifyEvent{}
	b := 1 // don't read event number

	v.Subtype = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Window = xproto.Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Owner = xproto.Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Selection = xproto.Atom(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Timestamp = xproto.Timestamp(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.SelectionTimestamp = xproto.Timestamp(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	b += 8 // padding

	return v, nil
}

// Bytes writes a SelectionNotifyEvent value to a byte slice.
func (v *SelectionNotifyEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 0
	b += 1

	buf[b] = v.Subtype
	b += 1

	b += 2 // skip sequence number

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Window))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Owner))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Selection))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Timestamp))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.SelectionTimestamp))
	b += 4

	b += 8 // padding

	return buf
}

// SeqID returns the sequence id attached to the SelectionNotify event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v *SelectionNotifyEvent) SeqID() uint16 {
	return v.Sequence
}

func init() { registerEvent(0, UnmarshalSelectionNotifyEvent) }

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

// ChangeCursor sends a checked request.
func ChangeCursor(c *xgb.XConn, Source xproto.Cursor, Destination xproto.Cursor) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"ChangeCursor\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.SendRecv(changeCursorRequest(op, Source, Destination), nil)
}

// ChangeCursorUnchecked sends an unchecked request.
func ChangeCursorUnchecked(c *xgb.XConn, Source xproto.Cursor, Destination xproto.Cursor) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"ChangeCursor\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.Send(changeCursorRequest(op, Source, Destination))
}

// Write request to wire for ChangeCursor
// changeCursorRequest writes a ChangeCursor request to a byte slice.
func changeCursorRequest(opcode uint8, Source xproto.Cursor, Destination xproto.Cursor) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 26 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Source))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Destination))
	b += 4

	return buf
}

// ChangeCursorByName sends a checked request.
func ChangeCursorByName(c *xgb.XConn, Src xproto.Cursor, Nbytes uint16, Name string) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"ChangeCursorByName\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.SendRecv(changeCursorByNameRequest(op, Src, Nbytes, Name), nil)
}

// ChangeCursorByNameUnchecked sends an unchecked request.
func ChangeCursorByNameUnchecked(c *xgb.XConn, Src xproto.Cursor, Nbytes uint16, Name string) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"ChangeCursorByName\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.Send(changeCursorByNameRequest(op, Src, Nbytes, Name))
}

// Write request to wire for ChangeCursorByName
// changeCursorByNameRequest writes a ChangeCursorByName request to a byte slice.
func changeCursorByNameRequest(opcode uint8, Src xproto.Cursor, Nbytes uint16, Name string) []byte {
	size := internal.Pad4((12 + internal.Pad4((int(Nbytes) * 1))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 27 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Src))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], Nbytes)
	b += 2

	b += 2 // padding

	copy(buf[b:], Name[:Nbytes])
	b += int(Nbytes)

	return buf
}

// ChangeSaveSet sends a checked request.
func ChangeSaveSet(c *xgb.XConn, Mode byte, Target byte, Map byte, Window xproto.Window) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"ChangeSaveSet\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.SendRecv(changeSaveSetRequest(op, Mode, Target, Map, Window), nil)
}

// ChangeSaveSetUnchecked sends an unchecked request.
func ChangeSaveSetUnchecked(c *xgb.XConn, Mode byte, Target byte, Map byte, Window xproto.Window) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"ChangeSaveSet\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.Send(changeSaveSetRequest(op, Mode, Target, Map, Window))
}

// Write request to wire for ChangeSaveSet
// changeSaveSetRequest writes a ChangeSaveSet request to a byte slice.
func changeSaveSetRequest(opcode uint8, Mode byte, Target byte, Map byte, Window xproto.Window) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 1 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	buf[b] = Mode
	b += 1

	buf[b] = Target
	b += 1

	buf[b] = Map
	b += 1

	b += 1 // padding

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	return buf
}

// CopyRegion sends a checked request.
func CopyRegion(c *xgb.XConn, Source Region, Destination Region) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"CopyRegion\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.SendRecv(copyRegionRequest(op, Source, Destination), nil)
}

// CopyRegionUnchecked sends an unchecked request.
func CopyRegionUnchecked(c *xgb.XConn, Source Region, Destination Region) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"CopyRegion\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.Send(copyRegionRequest(op, Source, Destination))
}

// Write request to wire for CopyRegion
// copyRegionRequest writes a CopyRegion request to a byte slice.
func copyRegionRequest(opcode uint8, Source Region, Destination Region) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 12 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Source))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Destination))
	b += 4

	return buf
}

// CreatePointerBarrier sends a checked request.
func CreatePointerBarrier(c *xgb.XConn, Barrier Barrier, Window xproto.Window, X1 uint16, Y1 uint16, X2 uint16, Y2 uint16, Directions uint32, NumDevices uint16, Devices []uint16) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"CreatePointerBarrier\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.SendRecv(createPointerBarrierRequest(op, Barrier, Window, X1, Y1, X2, Y2, Directions, NumDevices, Devices), nil)
}

// CreatePointerBarrierUnchecked sends an unchecked request.
func CreatePointerBarrierUnchecked(c *xgb.XConn, Barrier Barrier, Window xproto.Window, X1 uint16, Y1 uint16, X2 uint16, Y2 uint16, Directions uint32, NumDevices uint16, Devices []uint16) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"CreatePointerBarrier\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.Send(createPointerBarrierRequest(op, Barrier, Window, X1, Y1, X2, Y2, Directions, NumDevices, Devices))
}

// Write request to wire for CreatePointerBarrier
// createPointerBarrierRequest writes a CreatePointerBarrier request to a byte slice.
func createPointerBarrierRequest(opcode uint8, Barrier Barrier, Window xproto.Window, X1 uint16, Y1 uint16, X2 uint16, Y2 uint16, Directions uint32, NumDevices uint16, Devices []uint16) []byte {
	size := internal.Pad4((28 + internal.Pad4((int(NumDevices) * 2))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 31 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Barrier))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], X1)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Y1)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], X2)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Y2)
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], Directions)
	b += 4

	b += 2 // padding

	binary.LittleEndian.PutUint16(buf[b:], NumDevices)
	b += 2

	for i := 0; i < int(NumDevices); i++ {
		binary.LittleEndian.PutUint16(buf[b:], Devices[i])
		b += 2
	}

	return buf
}

// CreateRegion sends a checked request.
func CreateRegion(c *xgb.XConn, Region Region, Rectangles []xproto.Rectangle) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"CreateRegion\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.SendRecv(createRegionRequest(op, Region, Rectangles), nil)
}

// CreateRegionUnchecked sends an unchecked request.
func CreateRegionUnchecked(c *xgb.XConn, Region Region, Rectangles []xproto.Rectangle) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"CreateRegion\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.Send(createRegionRequest(op, Region, Rectangles))
}

// Write request to wire for CreateRegion
// createRegionRequest writes a CreateRegion request to a byte slice.
func createRegionRequest(opcode uint8, Region Region, Rectangles []xproto.Rectangle) []byte {
	size := internal.Pad4((8 + internal.Pad4((len(Rectangles) * 8))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 5 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Region))
	b += 4

	b += xproto.RectangleListBytes(buf[b:], Rectangles)

	return buf
}

// CreateRegionFromBitmap sends a checked request.
func CreateRegionFromBitmap(c *xgb.XConn, Region Region, Bitmap xproto.Pixmap) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"CreateRegionFromBitmap\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.SendRecv(createRegionFromBitmapRequest(op, Region, Bitmap), nil)
}

// CreateRegionFromBitmapUnchecked sends an unchecked request.
func CreateRegionFromBitmapUnchecked(c *xgb.XConn, Region Region, Bitmap xproto.Pixmap) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"CreateRegionFromBitmap\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.Send(createRegionFromBitmapRequest(op, Region, Bitmap))
}

// Write request to wire for CreateRegionFromBitmap
// createRegionFromBitmapRequest writes a CreateRegionFromBitmap request to a byte slice.
func createRegionFromBitmapRequest(opcode uint8, Region Region, Bitmap xproto.Pixmap) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 6 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Region))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Bitmap))
	b += 4

	return buf
}

// CreateRegionFromGC sends a checked request.
func CreateRegionFromGC(c *xgb.XConn, Region Region, Gc xproto.Gcontext) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"CreateRegionFromGC\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.SendRecv(createRegionFromGCRequest(op, Region, Gc), nil)
}

// CreateRegionFromGCUnchecked sends an unchecked request.
func CreateRegionFromGCUnchecked(c *xgb.XConn, Region Region, Gc xproto.Gcontext) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"CreateRegionFromGC\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.Send(createRegionFromGCRequest(op, Region, Gc))
}

// Write request to wire for CreateRegionFromGC
// createRegionFromGCRequest writes a CreateRegionFromGC request to a byte slice.
func createRegionFromGCRequest(opcode uint8, Region Region, Gc xproto.Gcontext) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 8 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Region))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Gc))
	b += 4

	return buf
}

// CreateRegionFromPicture sends a checked request.
func CreateRegionFromPicture(c *xgb.XConn, Region Region, Picture render.Picture) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"CreateRegionFromPicture\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.SendRecv(createRegionFromPictureRequest(op, Region, Picture), nil)
}

// CreateRegionFromPictureUnchecked sends an unchecked request.
func CreateRegionFromPictureUnchecked(c *xgb.XConn, Region Region, Picture render.Picture) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"CreateRegionFromPicture\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.Send(createRegionFromPictureRequest(op, Region, Picture))
}

// Write request to wire for CreateRegionFromPicture
// createRegionFromPictureRequest writes a CreateRegionFromPicture request to a byte slice.
func createRegionFromPictureRequest(opcode uint8, Region Region, Picture render.Picture) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 9 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Region))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Picture))
	b += 4

	return buf
}

// CreateRegionFromWindow sends a checked request.
func CreateRegionFromWindow(c *xgb.XConn, Region Region, Window xproto.Window, Kind shape.Kind) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"CreateRegionFromWindow\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.SendRecv(createRegionFromWindowRequest(op, Region, Window, Kind), nil)
}

// CreateRegionFromWindowUnchecked sends an unchecked request.
func CreateRegionFromWindowUnchecked(c *xgb.XConn, Region Region, Window xproto.Window, Kind shape.Kind) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"CreateRegionFromWindow\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.Send(createRegionFromWindowRequest(op, Region, Window, Kind))
}

// Write request to wire for CreateRegionFromWindow
// createRegionFromWindowRequest writes a CreateRegionFromWindow request to a byte slice.
func createRegionFromWindowRequest(opcode uint8, Region Region, Window xproto.Window, Kind shape.Kind) []byte {
	size := 16
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 7 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Region))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	buf[b] = uint8(Kind)
	b += 1

	b += 3 // padding

	return buf
}

// DeletePointerBarrier sends a checked request.
func DeletePointerBarrier(c *xgb.XConn, Barrier Barrier) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"DeletePointerBarrier\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.SendRecv(deletePointerBarrierRequest(op, Barrier), nil)
}

// DeletePointerBarrierUnchecked sends an unchecked request.
func DeletePointerBarrierUnchecked(c *xgb.XConn, Barrier Barrier) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"DeletePointerBarrier\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.Send(deletePointerBarrierRequest(op, Barrier))
}

// Write request to wire for DeletePointerBarrier
// deletePointerBarrierRequest writes a DeletePointerBarrier request to a byte slice.
func deletePointerBarrierRequest(opcode uint8, Barrier Barrier) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 32 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Barrier))
	b += 4

	return buf
}

// DestroyRegion sends a checked request.
func DestroyRegion(c *xgb.XConn, Region Region) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"DestroyRegion\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.SendRecv(destroyRegionRequest(op, Region), nil)
}

// DestroyRegionUnchecked sends an unchecked request.
func DestroyRegionUnchecked(c *xgb.XConn, Region Region) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"DestroyRegion\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.Send(destroyRegionRequest(op, Region))
}

// Write request to wire for DestroyRegion
// destroyRegionRequest writes a DestroyRegion request to a byte slice.
func destroyRegionRequest(opcode uint8, Region Region) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 10 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Region))
	b += 4

	return buf
}

// ExpandRegion sends a checked request.
func ExpandRegion(c *xgb.XConn, Source Region, Destination Region, Left uint16, Right uint16, Top uint16, Bottom uint16) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"ExpandRegion\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.SendRecv(expandRegionRequest(op, Source, Destination, Left, Right, Top, Bottom), nil)
}

// ExpandRegionUnchecked sends an unchecked request.
func ExpandRegionUnchecked(c *xgb.XConn, Source Region, Destination Region, Left uint16, Right uint16, Top uint16, Bottom uint16) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"ExpandRegion\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.Send(expandRegionRequest(op, Source, Destination, Left, Right, Top, Bottom))
}

// Write request to wire for ExpandRegion
// expandRegionRequest writes a ExpandRegion request to a byte slice.
func expandRegionRequest(opcode uint8, Source Region, Destination Region, Left uint16, Right uint16, Top uint16, Bottom uint16) []byte {
	size := 20
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 28 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Source))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Destination))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], Left)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Right)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Top)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], Bottom)
	b += 2

	return buf
}

// FetchRegion sends a checked request.
func FetchRegion(c *xgb.XConn, Region Region) (FetchRegionReply, error) {
	var reply FetchRegionReply
	op, ok := c.Ext("XFIXES")
	if !ok {
		return reply, errors.New("cannot issue request \"FetchRegion\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	err := c.SendRecv(fetchRegionRequest(op, Region), &reply)
	return reply, err
}

// FetchRegionUnchecked sends an unchecked request.
func FetchRegionUnchecked(c *xgb.XConn, Region Region) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"FetchRegion\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.Send(fetchRegionRequest(op, Region))
}

// FetchRegionReply represents the data returned from a FetchRegion request.
type FetchRegionReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Extents xproto.Rectangle
	// padding: 16 bytes
	Rectangles []xproto.Rectangle // size: internal.Pad4(((int(Length) / 2) * 8))
}

// Unmarshal reads a byte slice into a FetchRegionReply value.
func (v *FetchRegionReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4(((int(v.Length) / 2) * 8))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"FetchRegionReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Extents = xproto.Rectangle{}
	b += xproto.RectangleRead(buf[b:], &v.Extents)

	b += 16 // padding

	v.Rectangles = make([]xproto.Rectangle, (int(v.Length) / 2))
	b += xproto.RectangleReadList(buf[b:], v.Rectangles)

	return nil
}

// Write request to wire for FetchRegion
// fetchRegionRequest writes a FetchRegion request to a byte slice.
func fetchRegionRequest(opcode uint8, Region Region) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 19 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Region))
	b += 4

	return buf
}

// GetClientDisconnectMode sends a checked request.
func GetClientDisconnectMode(c *xgb.XConn) (GetClientDisconnectModeReply, error) {
	var reply GetClientDisconnectModeReply
	op, ok := c.Ext("XFIXES")
	if !ok {
		return reply, errors.New("cannot issue request \"GetClientDisconnectMode\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getClientDisconnectModeRequest(op), &reply)
	return reply, err
}

// GetClientDisconnectModeUnchecked sends an unchecked request.
func GetClientDisconnectModeUnchecked(c *xgb.XConn) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"GetClientDisconnectMode\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.Send(getClientDisconnectModeRequest(op))
}

// GetClientDisconnectModeReply represents the data returned from a GetClientDisconnectMode request.
type GetClientDisconnectModeReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	DisconnectMode uint32
	// padding: 20 bytes
}

// Unmarshal reads a byte slice into a GetClientDisconnectModeReply value.
func (v *GetClientDisconnectModeReply) Unmarshal(buf []byte) error {
	if size := 32; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetClientDisconnectModeReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.DisconnectMode = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 20 // padding

	return nil
}

// Write request to wire for GetClientDisconnectMode
// getClientDisconnectModeRequest writes a GetClientDisconnectMode request to a byte slice.
func getClientDisconnectModeRequest(opcode uint8) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 34 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// GetCursorImage sends a checked request.
func GetCursorImage(c *xgb.XConn) (GetCursorImageReply, error) {
	var reply GetCursorImageReply
	op, ok := c.Ext("XFIXES")
	if !ok {
		return reply, errors.New("cannot issue request \"GetCursorImage\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getCursorImageRequest(op), &reply)
	return reply, err
}

// GetCursorImageUnchecked sends an unchecked request.
func GetCursorImageUnchecked(c *xgb.XConn) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"GetCursorImage\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.Send(getCursorImageRequest(op))
}

// GetCursorImageReply represents the data returned from a GetCursorImage request.
type GetCursorImageReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	X            int16
	Y            int16
	Width        uint16
	Height       uint16
	Xhot         uint16
	Yhot         uint16
	CursorSerial uint32
	// padding: 8 bytes
	CursorImage []uint32 // size: internal.Pad4(((int(Width) * int(Height)) * 4))
}

// Unmarshal reads a byte slice into a GetCursorImageReply value.
func (v *GetCursorImageReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4(((int(v.Width) * int(v.Height)) * 4))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetCursorImageReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.X = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.Y = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.Width = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Height = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Xhot = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Yhot = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.CursorSerial = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 8 // padding

	v.CursorImage = make([]uint32, (int(v.Width) * int(v.Height)))
	for i := 0; i < int((int(v.Width) * int(v.Height))); i++ {
		v.CursorImage[i] = binary.LittleEndian.Uint32(buf[b:])
		b += 4
	}

	return nil
}

// Write request to wire for GetCursorImage
// getCursorImageRequest writes a GetCursorImage request to a byte slice.
func getCursorImageRequest(opcode uint8) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 4 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// GetCursorImageAndName sends a checked request.
func GetCursorImageAndName(c *xgb.XConn) (GetCursorImageAndNameReply, error) {
	var reply GetCursorImageAndNameReply
	op, ok := c.Ext("XFIXES")
	if !ok {
		return reply, errors.New("cannot issue request \"GetCursorImageAndName\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getCursorImageAndNameRequest(op), &reply)
	return reply, err
}

// GetCursorImageAndNameUnchecked sends an unchecked request.
func GetCursorImageAndNameUnchecked(c *xgb.XConn) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"GetCursorImageAndName\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.Send(getCursorImageAndNameRequest(op))
}

// GetCursorImageAndNameReply represents the data returned from a GetCursorImageAndName request.
type GetCursorImageAndNameReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	X            int16
	Y            int16
	Width        uint16
	Height       uint16
	Xhot         uint16
	Yhot         uint16
	CursorSerial uint32
	CursorAtom   xproto.Atom
	Nbytes       uint16
	// padding: 2 bytes
	CursorImage []uint32 // size: internal.Pad4(((int(Width) * int(Height)) * 4))
	Name        string   // size: internal.Pad4((int(Nbytes) * 1))
}

// Unmarshal reads a byte slice into a GetCursorImageAndNameReply value.
func (v *GetCursorImageAndNameReply) Unmarshal(buf []byte) error {
	if size := ((32 + internal.Pad4(((int(v.Width) * int(v.Height)) * 4))) + internal.Pad4((int(v.Nbytes) * 1))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetCursorImageAndNameReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.X = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.Y = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.Width = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Height = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Xhot = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Yhot = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.CursorSerial = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.CursorAtom = xproto.Atom(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Nbytes = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	b += 2 // padding

	v.CursorImage = make([]uint32, (int(v.Width) * int(v.Height)))
	for i := 0; i < int((int(v.Width) * int(v.Height))); i++ {
		v.CursorImage[i] = binary.LittleEndian.Uint32(buf[b:])
		b += 4
	}

	{
		byteString := make([]byte, v.Nbytes)
		copy(byteString[:v.Nbytes], buf[b:])
		v.Name = string(byteString)
		b += int(v.Nbytes)
	}

	return nil
}

// Write request to wire for GetCursorImageAndName
// getCursorImageAndNameRequest writes a GetCursorImageAndName request to a byte slice.
func getCursorImageAndNameRequest(opcode uint8) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 25 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// GetCursorName sends a checked request.
func GetCursorName(c *xgb.XConn, Cursor xproto.Cursor) (GetCursorNameReply, error) {
	var reply GetCursorNameReply
	op, ok := c.Ext("XFIXES")
	if !ok {
		return reply, errors.New("cannot issue request \"GetCursorName\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getCursorNameRequest(op, Cursor), &reply)
	return reply, err
}

// GetCursorNameUnchecked sends an unchecked request.
func GetCursorNameUnchecked(c *xgb.XConn, Cursor xproto.Cursor) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"GetCursorName\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.Send(getCursorNameRequest(op, Cursor))
}

// GetCursorNameReply represents the data returned from a GetCursorName request.
type GetCursorNameReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Atom   xproto.Atom
	Nbytes uint16
	// padding: 18 bytes
	Name string // size: internal.Pad4((int(Nbytes) * 1))
}

// Unmarshal reads a byte slice into a GetCursorNameReply value.
func (v *GetCursorNameReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.Nbytes) * 1))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetCursorNameReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Atom = xproto.Atom(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Nbytes = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	b += 18 // padding

	{
		byteString := make([]byte, v.Nbytes)
		copy(byteString[:v.Nbytes], buf[b:])
		v.Name = string(byteString)
		b += int(v.Nbytes)
	}

	return nil
}

// Write request to wire for GetCursorName
// getCursorNameRequest writes a GetCursorName request to a byte slice.
func getCursorNameRequest(opcode uint8, Cursor xproto.Cursor) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 24 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Cursor))
	b += 4

	return buf
}

// HideCursor sends a checked request.
func HideCursor(c *xgb.XConn, Window xproto.Window) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"HideCursor\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.SendRecv(hideCursorRequest(op, Window), nil)
}

// HideCursorUnchecked sends an unchecked request.
func HideCursorUnchecked(c *xgb.XConn, Window xproto.Window) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"HideCursor\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.Send(hideCursorRequest(op, Window))
}

// Write request to wire for HideCursor
// hideCursorRequest writes a HideCursor request to a byte slice.
func hideCursorRequest(opcode uint8, Window xproto.Window) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 29 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	return buf
}

// IntersectRegion sends a checked request.
func IntersectRegion(c *xgb.XConn, Source1 Region, Source2 Region, Destination Region) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"IntersectRegion\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.SendRecv(intersectRegionRequest(op, Source1, Source2, Destination), nil)
}

// IntersectRegionUnchecked sends an unchecked request.
func IntersectRegionUnchecked(c *xgb.XConn, Source1 Region, Source2 Region, Destination Region) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"IntersectRegion\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.Send(intersectRegionRequest(op, Source1, Source2, Destination))
}

// Write request to wire for IntersectRegion
// intersectRegionRequest writes a IntersectRegion request to a byte slice.
func intersectRegionRequest(opcode uint8, Source1 Region, Source2 Region, Destination Region) []byte {
	size := 16
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 14 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Source1))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Source2))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Destination))
	b += 4

	return buf
}

// InvertRegion sends a checked request.
func InvertRegion(c *xgb.XConn, Source Region, Bounds xproto.Rectangle, Destination Region) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"InvertRegion\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.SendRecv(invertRegionRequest(op, Source, Bounds, Destination), nil)
}

// InvertRegionUnchecked sends an unchecked request.
func InvertRegionUnchecked(c *xgb.XConn, Source Region, Bounds xproto.Rectangle, Destination Region) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"InvertRegion\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.Send(invertRegionRequest(op, Source, Bounds, Destination))
}

// Write request to wire for InvertRegion
// invertRegionRequest writes a InvertRegion request to a byte slice.
func invertRegionRequest(opcode uint8, Source Region, Bounds xproto.Rectangle, Destination Region) []byte {
	size := 20
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 16 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Source))
	b += 4

	{
		structBytes := Bounds.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	binary.LittleEndian.PutUint32(buf[b:], uint32(Destination))
	b += 4

	return buf
}

// QueryVersion sends a checked request.
func QueryVersion(c *xgb.XConn, ClientMajorVersion uint32, ClientMinorVersion uint32) (QueryVersionReply, error) {
	var reply QueryVersionReply
	op, ok := c.Ext("XFIXES")
	if !ok {
		return reply, errors.New("cannot issue request \"QueryVersion\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	err := c.SendRecv(queryVersionRequest(op, ClientMajorVersion, ClientMinorVersion), &reply)
	return reply, err
}

// QueryVersionUnchecked sends an unchecked request.
func QueryVersionUnchecked(c *xgb.XConn, ClientMajorVersion uint32, ClientMinorVersion uint32) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"QueryVersion\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
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

// RegionExtents sends a checked request.
func RegionExtents(c *xgb.XConn, Source Region, Destination Region) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"RegionExtents\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.SendRecv(regionExtentsRequest(op, Source, Destination), nil)
}

// RegionExtentsUnchecked sends an unchecked request.
func RegionExtentsUnchecked(c *xgb.XConn, Source Region, Destination Region) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"RegionExtents\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.Send(regionExtentsRequest(op, Source, Destination))
}

// Write request to wire for RegionExtents
// regionExtentsRequest writes a RegionExtents request to a byte slice.
func regionExtentsRequest(opcode uint8, Source Region, Destination Region) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 18 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Source))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Destination))
	b += 4

	return buf
}

// SelectCursorInput sends a checked request.
func SelectCursorInput(c *xgb.XConn, Window xproto.Window, EventMask uint32) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"SelectCursorInput\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.SendRecv(selectCursorInputRequest(op, Window, EventMask), nil)
}

// SelectCursorInputUnchecked sends an unchecked request.
func SelectCursorInputUnchecked(c *xgb.XConn, Window xproto.Window, EventMask uint32) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"SelectCursorInput\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.Send(selectCursorInputRequest(op, Window, EventMask))
}

// Write request to wire for SelectCursorInput
// selectCursorInputRequest writes a SelectCursorInput request to a byte slice.
func selectCursorInputRequest(opcode uint8, Window xproto.Window, EventMask uint32) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 3 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], EventMask)
	b += 4

	return buf
}

// SelectSelectionInput sends a checked request.
func SelectSelectionInput(c *xgb.XConn, Window xproto.Window, Selection xproto.Atom, EventMask uint32) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"SelectSelectionInput\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.SendRecv(selectSelectionInputRequest(op, Window, Selection, EventMask), nil)
}

// SelectSelectionInputUnchecked sends an unchecked request.
func SelectSelectionInputUnchecked(c *xgb.XConn, Window xproto.Window, Selection xproto.Atom, EventMask uint32) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"SelectSelectionInput\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.Send(selectSelectionInputRequest(op, Window, Selection, EventMask))
}

// Write request to wire for SelectSelectionInput
// selectSelectionInputRequest writes a SelectSelectionInput request to a byte slice.
func selectSelectionInputRequest(opcode uint8, Window xproto.Window, Selection xproto.Atom, EventMask uint32) []byte {
	size := 16
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

	binary.LittleEndian.PutUint32(buf[b:], uint32(Selection))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], EventMask)
	b += 4

	return buf
}

// SetClientDisconnectMode sends a checked request.
func SetClientDisconnectMode(c *xgb.XConn, DisconnectMode uint32) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"SetClientDisconnectMode\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.SendRecv(setClientDisconnectModeRequest(op, DisconnectMode), nil)
}

// SetClientDisconnectModeUnchecked sends an unchecked request.
func SetClientDisconnectModeUnchecked(c *xgb.XConn, DisconnectMode uint32) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"SetClientDisconnectMode\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.Send(setClientDisconnectModeRequest(op, DisconnectMode))
}

// Write request to wire for SetClientDisconnectMode
// setClientDisconnectModeRequest writes a SetClientDisconnectMode request to a byte slice.
func setClientDisconnectModeRequest(opcode uint8, DisconnectMode uint32) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 33 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], DisconnectMode)
	b += 4

	return buf
}

// SetCursorName sends a checked request.
func SetCursorName(c *xgb.XConn, Cursor xproto.Cursor, Nbytes uint16, Name string) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"SetCursorName\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.SendRecv(setCursorNameRequest(op, Cursor, Nbytes, Name), nil)
}

// SetCursorNameUnchecked sends an unchecked request.
func SetCursorNameUnchecked(c *xgb.XConn, Cursor xproto.Cursor, Nbytes uint16, Name string) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"SetCursorName\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.Send(setCursorNameRequest(op, Cursor, Nbytes, Name))
}

// Write request to wire for SetCursorName
// setCursorNameRequest writes a SetCursorName request to a byte slice.
func setCursorNameRequest(opcode uint8, Cursor xproto.Cursor, Nbytes uint16, Name string) []byte {
	size := internal.Pad4((12 + internal.Pad4((int(Nbytes) * 1))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 23 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Cursor))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], Nbytes)
	b += 2

	b += 2 // padding

	copy(buf[b:], Name[:Nbytes])
	b += int(Nbytes)

	return buf
}

// SetGCClipRegion sends a checked request.
func SetGCClipRegion(c *xgb.XConn, Gc xproto.Gcontext, Region Region, XOrigin int16, YOrigin int16) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"SetGCClipRegion\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.SendRecv(setGCClipRegionRequest(op, Gc, Region, XOrigin, YOrigin), nil)
}

// SetGCClipRegionUnchecked sends an unchecked request.
func SetGCClipRegionUnchecked(c *xgb.XConn, Gc xproto.Gcontext, Region Region, XOrigin int16, YOrigin int16) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"SetGCClipRegion\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.Send(setGCClipRegionRequest(op, Gc, Region, XOrigin, YOrigin))
}

// Write request to wire for SetGCClipRegion
// setGCClipRegionRequest writes a SetGCClipRegion request to a byte slice.
func setGCClipRegionRequest(opcode uint8, Gc xproto.Gcontext, Region Region, XOrigin int16, YOrigin int16) []byte {
	size := 16
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 20 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Gc))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Region))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(XOrigin))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(YOrigin))
	b += 2

	return buf
}

// SetPictureClipRegion sends a checked request.
func SetPictureClipRegion(c *xgb.XConn, Picture render.Picture, Region Region, XOrigin int16, YOrigin int16) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"SetPictureClipRegion\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.SendRecv(setPictureClipRegionRequest(op, Picture, Region, XOrigin, YOrigin), nil)
}

// SetPictureClipRegionUnchecked sends an unchecked request.
func SetPictureClipRegionUnchecked(c *xgb.XConn, Picture render.Picture, Region Region, XOrigin int16, YOrigin int16) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"SetPictureClipRegion\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.Send(setPictureClipRegionRequest(op, Picture, Region, XOrigin, YOrigin))
}

// Write request to wire for SetPictureClipRegion
// setPictureClipRegionRequest writes a SetPictureClipRegion request to a byte slice.
func setPictureClipRegionRequest(opcode uint8, Picture render.Picture, Region Region, XOrigin int16, YOrigin int16) []byte {
	size := 16
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 22 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Picture))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Region))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(XOrigin))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(YOrigin))
	b += 2

	return buf
}

// SetRegion sends a checked request.
func SetRegion(c *xgb.XConn, Region Region, Rectangles []xproto.Rectangle) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"SetRegion\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.SendRecv(setRegionRequest(op, Region, Rectangles), nil)
}

// SetRegionUnchecked sends an unchecked request.
func SetRegionUnchecked(c *xgb.XConn, Region Region, Rectangles []xproto.Rectangle) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"SetRegion\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.Send(setRegionRequest(op, Region, Rectangles))
}

// Write request to wire for SetRegion
// setRegionRequest writes a SetRegion request to a byte slice.
func setRegionRequest(opcode uint8, Region Region, Rectangles []xproto.Rectangle) []byte {
	size := internal.Pad4((8 + internal.Pad4((len(Rectangles) * 8))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 11 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Region))
	b += 4

	b += xproto.RectangleListBytes(buf[b:], Rectangles)

	return buf
}

// SetWindowShapeRegion sends a checked request.
func SetWindowShapeRegion(c *xgb.XConn, Dest xproto.Window, DestKind shape.Kind, XOffset int16, YOffset int16, Region Region) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"SetWindowShapeRegion\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.SendRecv(setWindowShapeRegionRequest(op, Dest, DestKind, XOffset, YOffset, Region), nil)
}

// SetWindowShapeRegionUnchecked sends an unchecked request.
func SetWindowShapeRegionUnchecked(c *xgb.XConn, Dest xproto.Window, DestKind shape.Kind, XOffset int16, YOffset int16, Region Region) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"SetWindowShapeRegion\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.Send(setWindowShapeRegionRequest(op, Dest, DestKind, XOffset, YOffset, Region))
}

// Write request to wire for SetWindowShapeRegion
// setWindowShapeRegionRequest writes a SetWindowShapeRegion request to a byte slice.
func setWindowShapeRegionRequest(opcode uint8, Dest xproto.Window, DestKind shape.Kind, XOffset int16, YOffset int16, Region Region) []byte {
	size := 20
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 21 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Dest))
	b += 4

	buf[b] = uint8(DestKind)
	b += 1

	b += 3 // padding

	binary.LittleEndian.PutUint16(buf[b:], uint16(XOffset))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(YOffset))
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Region))
	b += 4

	return buf
}

// ShowCursor sends a checked request.
func ShowCursor(c *xgb.XConn, Window xproto.Window) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"ShowCursor\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.SendRecv(showCursorRequest(op, Window), nil)
}

// ShowCursorUnchecked sends an unchecked request.
func ShowCursorUnchecked(c *xgb.XConn, Window xproto.Window) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"ShowCursor\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.Send(showCursorRequest(op, Window))
}

// Write request to wire for ShowCursor
// showCursorRequest writes a ShowCursor request to a byte slice.
func showCursorRequest(opcode uint8, Window xproto.Window) []byte {
	size := 8
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

	return buf
}

// SubtractRegion sends a checked request.
func SubtractRegion(c *xgb.XConn, Source1 Region, Source2 Region, Destination Region) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"SubtractRegion\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.SendRecv(subtractRegionRequest(op, Source1, Source2, Destination), nil)
}

// SubtractRegionUnchecked sends an unchecked request.
func SubtractRegionUnchecked(c *xgb.XConn, Source1 Region, Source2 Region, Destination Region) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"SubtractRegion\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.Send(subtractRegionRequest(op, Source1, Source2, Destination))
}

// Write request to wire for SubtractRegion
// subtractRegionRequest writes a SubtractRegion request to a byte slice.
func subtractRegionRequest(opcode uint8, Source1 Region, Source2 Region, Destination Region) []byte {
	size := 16
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 15 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Source1))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Source2))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Destination))
	b += 4

	return buf
}

// TranslateRegion sends a checked request.
func TranslateRegion(c *xgb.XConn, Region Region, Dx int16, Dy int16) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"TranslateRegion\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.SendRecv(translateRegionRequest(op, Region, Dx, Dy), nil)
}

// TranslateRegionUnchecked sends an unchecked request.
func TranslateRegionUnchecked(c *xgb.XConn, Region Region, Dx int16, Dy int16) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"TranslateRegion\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.Send(translateRegionRequest(op, Region, Dx, Dy))
}

// Write request to wire for TranslateRegion
// translateRegionRequest writes a TranslateRegion request to a byte slice.
func translateRegionRequest(opcode uint8, Region Region, Dx int16, Dy int16) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 17 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Region))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], uint16(Dx))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(Dy))
	b += 2

	return buf
}

// UnionRegion sends a checked request.
func UnionRegion(c *xgb.XConn, Source1 Region, Source2 Region, Destination Region) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"UnionRegion\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.SendRecv(unionRegionRequest(op, Source1, Source2, Destination), nil)
}

// UnionRegionUnchecked sends an unchecked request.
func UnionRegionUnchecked(c *xgb.XConn, Source1 Region, Source2 Region, Destination Region) error {
	op, ok := c.Ext("XFIXES")
	if !ok {
		return errors.New("cannot issue request \"UnionRegion\" using the uninitialized extension \"XFIXES\". xfixes.Register(xconn) must be called first.")
	}
	return c.Send(unionRegionRequest(op, Source1, Source2, Destination))
}

// Write request to wire for UnionRegion
// unionRegionRequest writes a UnionRegion request to a byte slice.
func unionRegionRequest(opcode uint8, Source1 Region, Source2 Region, Destination Region) []byte {
	size := 16
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 13 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Source1))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Source2))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Destination))
	b += 4

	return buf
}
