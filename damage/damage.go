// FILE GENERATED AUTOMATICALLY FROM "damage.xml"
package damage

import (
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"codeberg.org/gruf/go-xgb"
	"codeberg.org/gruf/go-xgb/xfixes"
	"codeberg.org/gruf/go-xgb/xproto"
)

const (
	// ExtName is the user-friendly name string of this X extension.
	ExtName = "Damage"

	// ExtXName is the name string this extension is known by to the X server.
	ExtXName = "DAMAGE"
)

var (
	// generated index maps of defined event and error numbers -> unmarshalers.
	eventFuncs = map[uint8]xgb.EventUnmarshaler{}
	errorFuncs = map[uint8]xgb.ErrorUnmarshaler{}
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
		return fmt.Errorf("error querying X for \"Damage\": %w", err)
	} else if !reply.Present {
		return fmt.Errorf("no extension named \"Damage\" is known to the X server: reply=%+v", reply)
	}

	// Clone event funcs map but set our event no. start index
	extEventFuncs := map[uint8]xgb.EventUnmarshaler{}
	for n, fn := range eventFuncs {
		extEventFuncs[n+reply.FirstEvent] = fn
	}

	// Clone error funcs map but set our error no. start index
	extErrorFuncs := map[uint8]xgb.ErrorUnmarshaler{}
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

// BadBadDamage is the error number for a BadBadDamage.
const BadBadDamage = 0

type BadDamageError struct {
	Sequence uint16
	NiceName string
}

// UnmarshalBadDamageError constructs a BadDamageError value that implements xgb.Error from a byte slice.
func UnmarshalBadDamageError(buf []byte) (xgb.XError, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"BadDamageError\"", len(buf))
	}

	v := BadDamageError{}
	v.NiceName = "BadDamage"

	b := 1 // skip error determinant
	b += 1 // don't read error number

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return v, nil
}

// SeqID returns the sequence id attached to the BadBadDamage error.
// This is mostly used internally.
func (err BadDamageError) SeqID() uint16 {
	return err.Sequence
}

// BadID returns the 'BadValue' number if one exists for the BadBadDamage error. If no bad value exists, 0 is returned.
func (err BadDamageError) BadID() uint32 {
	return 0
}

// Error returns a rudimentary string representation of the BadBadDamage error.

func (err BadDamageError) Error() string {
	fieldVals := make([]string, 0, 0)
	fieldVals = append(fieldVals, "NiceName: "+err.NiceName)
	fieldVals = append(fieldVals, fmt.Sprintf("Sequence: %d", err.Sequence))
	return "BadBadDamage {" + strings.Join(fieldVals, ", ") + "}"
}

func init() {
	registerError(0, UnmarshalBadDamageError)
}

type Damage uint32

func NewDamageID(c *xgb.XConn) (Damage, error) {
	id, err := c.NewXID()
	return Damage(id), err
}

// Notify is the event number for a NotifyEvent.
const Notify = 0

type NotifyEvent struct {
	Sequence  uint16
	Level     byte
	Drawable  xproto.Drawable
	Damage    Damage
	Timestamp xproto.Timestamp
	Area      xproto.Rectangle
	Geometry  xproto.Rectangle
}

// UnmarshalNotifyEvent constructs a NotifyEvent value that implements xgb.Event from a byte slice.
func UnmarshalNotifyEvent(buf []byte) (xgb.XEvent, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"NotifyEvent\"", len(buf))
	}

	v := NotifyEvent{}
	b := 1 // don't read event number

	v.Level = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Drawable = xproto.Drawable(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Damage = Damage(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Timestamp = xproto.Timestamp(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Area = xproto.Rectangle{}
	b += xproto.RectangleRead(buf[b:], &v.Area)

	v.Geometry = xproto.Rectangle{}
	b += xproto.RectangleRead(buf[b:], &v.Geometry)

	return v, nil
}

// Bytes writes a NotifyEvent value to a byte slice.
func (v NotifyEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 0
	b += 1

	buf[b] = v.Level
	b += 1

	b += 2 // skip sequence number

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Drawable))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Damage))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Timestamp))
	b += 4

	{
		structBytes := v.Area.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	{
		structBytes := v.Geometry.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	return buf
}

// SeqID returns the sequence id attached to the Notify event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v NotifyEvent) SeqID() uint16 {
	return v.Sequence
}

func init() {
	registerEvent(0, UnmarshalNotifyEvent)
}

const (
	ReportLevelRawRectangles   = 0
	ReportLevelDeltaRectangles = 1
	ReportLevelBoundingBox     = 2
	ReportLevelNonEmpty        = 3
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

// Add sends a checked request.
func Add(c *xgb.XConn, Drawable xproto.Drawable, Region xfixes.Region) error {
	op, ok := c.Ext("DAMAGE")
	if !ok {
		return errors.New("cannot issue request \"Add\" using the uninitialized extension \"DAMAGE\". damage.Register(xconn) must be called first.")
	}
	return c.SendRecv(addRequest(op, Drawable, Region), nil)
}

// AddUnchecked sends an unchecked request.
func AddUnchecked(c *xgb.XConn, Drawable xproto.Drawable, Region xfixes.Region) error {
	op, ok := c.Ext("DAMAGE")
	if !ok {
		return errors.New("cannot issue request \"Add\" using the uninitialized extension \"DAMAGE\". damage.Register(xconn) must be called first.")
	}
	return c.Send(addRequest(op, Drawable, Region))
}

// Write request to wire for Add
// addRequest writes a Add request to a byte slice.
func addRequest(opcode uint8, Drawable xproto.Drawable, Region xfixes.Region) []byte {
	size := 12
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

	binary.LittleEndian.PutUint32(buf[b:], uint32(Region))
	b += 4

	return buf
}

// Create sends a checked request.
func Create(c *xgb.XConn, Damage Damage, Drawable xproto.Drawable, Level byte) error {
	op, ok := c.Ext("DAMAGE")
	if !ok {
		return errors.New("cannot issue request \"Create\" using the uninitialized extension \"DAMAGE\". damage.Register(xconn) must be called first.")
	}
	return c.SendRecv(createRequest(op, Damage, Drawable, Level), nil)
}

// CreateUnchecked sends an unchecked request.
func CreateUnchecked(c *xgb.XConn, Damage Damage, Drawable xproto.Drawable, Level byte) error {
	op, ok := c.Ext("DAMAGE")
	if !ok {
		return errors.New("cannot issue request \"Create\" using the uninitialized extension \"DAMAGE\". damage.Register(xconn) must be called first.")
	}
	return c.Send(createRequest(op, Damage, Drawable, Level))
}

// Write request to wire for Create
// createRequest writes a Create request to a byte slice.
func createRequest(opcode uint8, Damage Damage, Drawable xproto.Drawable, Level byte) []byte {
	size := 16
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 1 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Damage))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Drawable))
	b += 4

	buf[b] = Level
	b += 1

	b += 3 // padding

	return buf
}

// Destroy sends a checked request.
func Destroy(c *xgb.XConn, Damage Damage) error {
	op, ok := c.Ext("DAMAGE")
	if !ok {
		return errors.New("cannot issue request \"Destroy\" using the uninitialized extension \"DAMAGE\". damage.Register(xconn) must be called first.")
	}
	return c.SendRecv(destroyRequest(op, Damage), nil)
}

// DestroyUnchecked sends an unchecked request.
func DestroyUnchecked(c *xgb.XConn, Damage Damage) error {
	op, ok := c.Ext("DAMAGE")
	if !ok {
		return errors.New("cannot issue request \"Destroy\" using the uninitialized extension \"DAMAGE\". damage.Register(xconn) must be called first.")
	}
	return c.Send(destroyRequest(op, Damage))
}

// Write request to wire for Destroy
// destroyRequest writes a Destroy request to a byte slice.
func destroyRequest(opcode uint8, Damage Damage) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 2 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Damage))
	b += 4

	return buf
}

// QueryVersion sends a checked request.
func QueryVersion(c *xgb.XConn, ClientMajorVersion uint32, ClientMinorVersion uint32) (QueryVersionReply, error) {
	var reply QueryVersionReply
	op, ok := c.Ext("DAMAGE")
	if !ok {
		return reply, errors.New("cannot issue request \"QueryVersion\" using the uninitialized extension \"DAMAGE\". damage.Register(xconn) must be called first.")
	}
	err := c.SendRecv(queryVersionRequest(op, ClientMajorVersion, ClientMinorVersion), &reply)
	return reply, err
}

// QueryVersionUnchecked sends an unchecked request.
func QueryVersionUnchecked(c *xgb.XConn, ClientMajorVersion uint32, ClientMinorVersion uint32) error {
	op, ok := c.Ext("DAMAGE")
	if !ok {
		return errors.New("cannot issue request \"QueryVersion\" using the uninitialized extension \"DAMAGE\". damage.Register(xconn) must be called first.")
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

// Subtract sends a checked request.
func Subtract(c *xgb.XConn, Damage Damage, Repair xfixes.Region, Parts xfixes.Region) error {
	op, ok := c.Ext("DAMAGE")
	if !ok {
		return errors.New("cannot issue request \"Subtract\" using the uninitialized extension \"DAMAGE\". damage.Register(xconn) must be called first.")
	}
	return c.SendRecv(subtractRequest(op, Damage, Repair, Parts), nil)
}

// SubtractUnchecked sends an unchecked request.
func SubtractUnchecked(c *xgb.XConn, Damage Damage, Repair xfixes.Region, Parts xfixes.Region) error {
	op, ok := c.Ext("DAMAGE")
	if !ok {
		return errors.New("cannot issue request \"Subtract\" using the uninitialized extension \"DAMAGE\". damage.Register(xconn) must be called first.")
	}
	return c.Send(subtractRequest(op, Damage, Repair, Parts))
}

// Write request to wire for Subtract
// subtractRequest writes a Subtract request to a byte slice.
func subtractRequest(opcode uint8, Damage Damage, Repair xfixes.Region, Parts xfixes.Region) []byte {
	size := 16
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 3 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Damage))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Repair))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Parts))
	b += 4

	return buf
}
