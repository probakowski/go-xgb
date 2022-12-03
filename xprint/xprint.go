// FILE GENERATED AUTOMATICALLY FROM "xprint.xml"
package xprint

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
	ExtName = "XPrint"

	// ExtXName is the name string this extension is known by to the X server.
	ExtXName = "XpExtension"
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
		return fmt.Errorf("error querying X for \"XPrint\": %w", err)
	} else if !reply.Present {
		return fmt.Errorf("no extension named \"XPrint\" is known to the X server: reply=%+v", reply)
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

const (
	AttrJobAttr     = 1
	AttrDocAttr     = 2
	AttrPageAttr    = 3
	AttrPrinterAttr = 4
	AttrServerAttr  = 5
	AttrMediumAttr  = 6
	AttrSpoolerAttr = 7
)

// AttributNotify is the event number for a AttributNotifyEvent.
const AttributNotify = 1

type AttributNotifyEvent struct {
	Sequence uint16
	Detail   byte
	Context  Pcontext
}

// UnmarshalAttributNotifyEvent constructs a AttributNotifyEvent value that implements xgb.Event from a byte slice.
func UnmarshalAttributNotifyEvent(buf []byte) (xgb.XEvent, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"AttributNotifyEvent\"", len(buf))
	}

	v := AttributNotifyEvent{}
	b := 1 // don't read event number

	v.Detail = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Context = Pcontext(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	return v, nil
}

// Bytes writes a AttributNotifyEvent value to a byte slice.
func (v AttributNotifyEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 1
	b += 1

	buf[b] = v.Detail
	b += 1

	b += 2 // skip sequence number

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Context))
	b += 4

	return buf
}

// SeqID returns the sequence id attached to the AttributNotify event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v AttributNotifyEvent) SeqID() uint16 {
	return v.Sequence
}

func init() { registerEvent(1, UnmarshalAttributNotifyEvent) }

// BadBadContext is the error number for a BadBadContext.
const BadBadContext = 0

type BadContextError struct {
	Sequence uint16
	NiceName string
}

// UnmarshalBadContextError constructs a BadContextError value that implements xgb.Error from a byte slice.
func UnmarshalBadContextError(buf []byte) (xgb.XError, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"BadContextError\"", len(buf))
	}

	v := BadContextError{}
	v.NiceName = "BadContext"

	b := 1 // skip error determinant
	b += 1 // don't read error number

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return v, nil
}

// SeqID returns the sequence id attached to the BadBadContext error.
// This is mostly used internally.
func (err BadContextError) SeqID() uint16 {
	return err.Sequence
}

// BadID returns the 'BadValue' number if one exists for the BadBadContext error. If no bad value exists, 0 is returned.
func (err BadContextError) BadID() uint32 {
	return 0
}

// Error returns a rudimentary string representation of the BadBadContext error.
func (err BadContextError) Error() string {
	var buf strings.Builder

	buf.WriteString("BadBadContext{")
	buf.WriteString("NiceName: " + err.NiceName)
	buf.WriteByte(' ')
	buf.WriteString("Sequence: " + strconv.FormatUint(uint64(err.Sequence), 10))

	buf.WriteByte('}')

	return buf.String()
}

func init() { registerError(0, UnmarshalBadContextError) }

// BadBadSequence is the error number for a BadBadSequence.
const BadBadSequence = 1

type BadSequenceError struct {
	Sequence uint16
	NiceName string
}

// UnmarshalBadSequenceError constructs a BadSequenceError value that implements xgb.Error from a byte slice.
func UnmarshalBadSequenceError(buf []byte) (xgb.XError, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"BadSequenceError\"", len(buf))
	}

	v := BadSequenceError{}
	v.NiceName = "BadSequence"

	b := 1 // skip error determinant
	b += 1 // don't read error number

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return v, nil
}

// SeqID returns the sequence id attached to the BadBadSequence error.
// This is mostly used internally.
func (err BadSequenceError) SeqID() uint16 {
	return err.Sequence
}

// BadID returns the 'BadValue' number if one exists for the BadBadSequence error. If no bad value exists, 0 is returned.
func (err BadSequenceError) BadID() uint32 {
	return 0
}

// Error returns a rudimentary string representation of the BadBadSequence error.
func (err BadSequenceError) Error() string {
	var buf strings.Builder

	buf.WriteString("BadBadSequence{")
	buf.WriteString("NiceName: " + err.NiceName)
	buf.WriteByte(' ')
	buf.WriteString("Sequence: " + strconv.FormatUint(uint64(err.Sequence), 10))

	buf.WriteByte('}')

	return buf.String()
}

func init() { registerError(1, UnmarshalBadSequenceError) }

const (
	DetailStartJobNotify  = 1
	DetailEndJobNotify    = 2
	DetailStartDocNotify  = 3
	DetailEndDocNotify    = 4
	DetailStartPageNotify = 5
	DetailEndPageNotify   = 6
)

const (
	EvMaskNoEventMask   = 0
	EvMaskPrintMask     = 1
	EvMaskAttributeMask = 2
)

const (
	GetDocFinished       = 0
	GetDocSecondConsumer = 1
)

// Notify is the event number for a NotifyEvent.
const Notify = 0

type NotifyEvent struct {
	Sequence uint16
	Detail   byte
	Context  Pcontext
	Cancel   bool
}

// UnmarshalNotifyEvent constructs a NotifyEvent value that implements xgb.Event from a byte slice.
func UnmarshalNotifyEvent(buf []byte) (xgb.XEvent, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"NotifyEvent\"", len(buf))
	}

	v := NotifyEvent{}
	b := 1 // don't read event number

	v.Detail = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Context = Pcontext(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Cancel = (buf[b] == 1)
	b += 1

	return v, nil
}

// Bytes writes a NotifyEvent value to a byte slice.
func (v NotifyEvent) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	// write event number
	buf[b] = 0
	b += 1

	buf[b] = v.Detail
	b += 1

	b += 2 // skip sequence number

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Context))
	b += 4

	if v.Cancel {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	return buf
}

// SeqID returns the sequence id attached to the Notify event.
// Events without a sequence number (KeymapNotify) return 0.
// This is mostly used internally.
func (v NotifyEvent) SeqID() uint16 {
	return v.Sequence
}

func init() { registerEvent(0, UnmarshalNotifyEvent) }

type Pcontext uint32

func NewPcontextID(c *xgb.XConn) (Pcontext, error) {
	id, err := c.NewXID()
	return Pcontext(id), err
}

type Printer struct {
	NameLen uint32
	Name    []String8 // size: internal.Pad4((int(NameLen) * 1))
	// padding: 0 bytes
	DescLen     uint32
	Description []String8 // size: internal.Pad4((int(DescLen) * 1))
	// padding: 0 bytes
}

// PrinterRead reads a byte slice into a Printer value.
func PrinterRead(buf []byte, v *Printer) int {
	b := 0

	v.NameLen = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.Name = make([]String8, v.NameLen)
	for i := 0; i < int(v.NameLen); i++ {
		v.Name[i] = String8(buf[b])
		b += 1
	}

	b += 0 // padding

	v.DescLen = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.Description = make([]String8, v.DescLen)
	for i := 0; i < int(v.DescLen); i++ {
		v.Description[i] = String8(buf[b])
		b += 1
	}

	b += 0 // padding

	return b
}

// PrinterReadList reads a byte slice into a list of Printer values.
func PrinterReadList(buf []byte, dest []Printer) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Printer{}
		b += PrinterRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a Printer value to a byte slice.
func (v Printer) Bytes() []byte {
	buf := make([]byte, (((((4 + internal.Pad4((int(v.NameLen) * 1))) + 0) + 4) + internal.Pad4((int(v.DescLen) * 1))) + 0))
	b := 0

	binary.LittleEndian.PutUint32(buf[b:], v.NameLen)
	b += 4

	for i := 0; i < int(v.NameLen); i++ {
		buf[b] = uint8(v.Name[i])
		b += 1
	}

	b += 0 // padding

	binary.LittleEndian.PutUint32(buf[b:], v.DescLen)
	b += 4

	for i := 0; i < int(v.DescLen); i++ {
		buf[b] = uint8(v.Description[i])
		b += 1
	}

	b += 0 // padding

	return buf[:b]
}

// PrinterListBytes writes a list of Printer values to a byte slice.
func PrinterListBytes(buf []byte, list []Printer) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

// PrinterListSize computes the size (bytes) of a list of Printer values.
func PrinterListSize(list []Printer) int {
	size := 0
	for _, item := range list {
		size += (((((4 + internal.Pad4((int(item.NameLen) * 1))) + 0) + 4) + internal.Pad4((int(item.DescLen) * 1))) + 0)
	}
	return size
}

type String8 byte

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
func CreateContext(c *xgb.XConn, ContextId uint32, PrinterNameLen uint32, LocaleLen uint32, PrinterName []String8, Locale []String8) error {
	op, ok := c.Ext("XpExtension")
	if !ok {
		return errors.New("cannot issue request \"CreateContext\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	return c.SendRecv(createContextRequest(op, ContextId, PrinterNameLen, LocaleLen, PrinterName, Locale), nil)
}

// CreateContextUnchecked sends an unchecked request.
func CreateContextUnchecked(c *xgb.XConn, ContextId uint32, PrinterNameLen uint32, LocaleLen uint32, PrinterName []String8, Locale []String8) error {
	op, ok := c.Ext("XpExtension")
	if !ok {
		return errors.New("cannot issue request \"CreateContext\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	return c.Send(createContextRequest(op, ContextId, PrinterNameLen, LocaleLen, PrinterName, Locale))
}

// Write request to wire for CreateContext
// createContextRequest writes a CreateContext request to a byte slice.
func createContextRequest(opcode uint8, ContextId uint32, PrinterNameLen uint32, LocaleLen uint32, PrinterName []String8, Locale []String8) []byte {
	size := internal.Pad4(((16 + internal.Pad4((int(PrinterNameLen) * 1))) + internal.Pad4((int(LocaleLen) * 1))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 2 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], ContextId)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], PrinterNameLen)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], LocaleLen)
	b += 4

	for i := 0; i < int(PrinterNameLen); i++ {
		buf[b] = uint8(PrinterName[i])
		b += 1
	}

	for i := 0; i < int(LocaleLen); i++ {
		buf[b] = uint8(Locale[i])
		b += 1
	}

	return buf
}

// PrintDestroyContext sends a checked request.
func PrintDestroyContext(c *xgb.XConn, Context uint32) error {
	op, ok := c.Ext("XpExtension")
	if !ok {
		return errors.New("cannot issue request \"PrintDestroyContext\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	return c.SendRecv(printDestroyContextRequest(op, Context), nil)
}

// PrintDestroyContextUnchecked sends an unchecked request.
func PrintDestroyContextUnchecked(c *xgb.XConn, Context uint32) error {
	op, ok := c.Ext("XpExtension")
	if !ok {
		return errors.New("cannot issue request \"PrintDestroyContext\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	return c.Send(printDestroyContextRequest(op, Context))
}

// Write request to wire for PrintDestroyContext
// printDestroyContextRequest writes a PrintDestroyContext request to a byte slice.
func printDestroyContextRequest(opcode uint8, Context uint32) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 5 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], Context)
	b += 4

	return buf
}

// PrintEndDoc sends a checked request.
func PrintEndDoc(c *xgb.XConn, Cancel bool) error {
	op, ok := c.Ext("XpExtension")
	if !ok {
		return errors.New("cannot issue request \"PrintEndDoc\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	return c.SendRecv(printEndDocRequest(op, Cancel), nil)
}

// PrintEndDocUnchecked sends an unchecked request.
func PrintEndDocUnchecked(c *xgb.XConn, Cancel bool) error {
	op, ok := c.Ext("XpExtension")
	if !ok {
		return errors.New("cannot issue request \"PrintEndDoc\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	return c.Send(printEndDocRequest(op, Cancel))
}

// Write request to wire for PrintEndDoc
// printEndDocRequest writes a PrintEndDoc request to a byte slice.
func printEndDocRequest(opcode uint8, Cancel bool) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 10 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	if Cancel {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	return buf
}

// PrintEndJob sends a checked request.
func PrintEndJob(c *xgb.XConn, Cancel bool) error {
	op, ok := c.Ext("XpExtension")
	if !ok {
		return errors.New("cannot issue request \"PrintEndJob\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	return c.SendRecv(printEndJobRequest(op, Cancel), nil)
}

// PrintEndJobUnchecked sends an unchecked request.
func PrintEndJobUnchecked(c *xgb.XConn, Cancel bool) error {
	op, ok := c.Ext("XpExtension")
	if !ok {
		return errors.New("cannot issue request \"PrintEndJob\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	return c.Send(printEndJobRequest(op, Cancel))
}

// Write request to wire for PrintEndJob
// printEndJobRequest writes a PrintEndJob request to a byte slice.
func printEndJobRequest(opcode uint8, Cancel bool) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 8 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	if Cancel {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	return buf
}

// PrintEndPage sends a checked request.
func PrintEndPage(c *xgb.XConn, Cancel bool) error {
	op, ok := c.Ext("XpExtension")
	if !ok {
		return errors.New("cannot issue request \"PrintEndPage\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	return c.SendRecv(printEndPageRequest(op, Cancel), nil)
}

// PrintEndPageUnchecked sends an unchecked request.
func PrintEndPageUnchecked(c *xgb.XConn, Cancel bool) error {
	op, ok := c.Ext("XpExtension")
	if !ok {
		return errors.New("cannot issue request \"PrintEndPage\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	return c.Send(printEndPageRequest(op, Cancel))
}

// Write request to wire for PrintEndPage
// printEndPageRequest writes a PrintEndPage request to a byte slice.
func printEndPageRequest(opcode uint8, Cancel bool) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 14 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	if Cancel {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	b += 3 // padding

	return buf
}

// PrintGetAttributes sends a checked request.
func PrintGetAttributes(c *xgb.XConn, Context Pcontext, Pool byte) (PrintGetAttributesReply, error) {
	var reply PrintGetAttributesReply
	op, ok := c.Ext("XpExtension")
	if !ok {
		return reply, errors.New("cannot issue request \"PrintGetAttributes\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	err := c.SendRecv(printGetAttributesRequest(op, Context, Pool), &reply)
	return reply, err
}

// PrintGetAttributesUnchecked sends an unchecked request.
func PrintGetAttributesUnchecked(c *xgb.XConn, Context Pcontext, Pool byte) error {
	op, ok := c.Ext("XpExtension")
	if !ok {
		return errors.New("cannot issue request \"PrintGetAttributes\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	return c.Send(printGetAttributesRequest(op, Context, Pool))
}

// PrintGetAttributesReply represents the data returned from a PrintGetAttributes request.
type PrintGetAttributesReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	StringLen uint32
	// padding: 20 bytes
	Attributes []String8 // size: internal.Pad4((int(StringLen) * 1))
}

// Unmarshal reads a byte slice into a PrintGetAttributesReply value.
func (v *PrintGetAttributesReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.StringLen) * 1))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"PrintGetAttributesReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.StringLen = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 20 // padding

	v.Attributes = make([]String8, v.StringLen)
	for i := 0; i < int(v.StringLen); i++ {
		v.Attributes[i] = String8(buf[b])
		b += 1
	}

	return nil
}

// Write request to wire for PrintGetAttributes
// printGetAttributesRequest writes a PrintGetAttributes request to a byte slice.
func printGetAttributesRequest(opcode uint8, Context Pcontext, Pool byte) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 17 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Context))
	b += 4

	buf[b] = Pool
	b += 1

	b += 3 // padding

	return buf
}

// PrintGetContext sends a checked request.
func PrintGetContext(c *xgb.XConn) (PrintGetContextReply, error) {
	var reply PrintGetContextReply
	op, ok := c.Ext("XpExtension")
	if !ok {
		return reply, errors.New("cannot issue request \"PrintGetContext\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	err := c.SendRecv(printGetContextRequest(op), &reply)
	return reply, err
}

// PrintGetContextUnchecked sends an unchecked request.
func PrintGetContextUnchecked(c *xgb.XConn) error {
	op, ok := c.Ext("XpExtension")
	if !ok {
		return errors.New("cannot issue request \"PrintGetContext\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	return c.Send(printGetContextRequest(op))
}

// PrintGetContextReply represents the data returned from a PrintGetContext request.
type PrintGetContextReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Context uint32
}

// Unmarshal reads a byte slice into a PrintGetContextReply value.
func (v *PrintGetContextReply) Unmarshal(buf []byte) error {
	if size := 12; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"PrintGetContextReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Context = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	return nil
}

// Write request to wire for PrintGetContext
// printGetContextRequest writes a PrintGetContext request to a byte slice.
func printGetContextRequest(opcode uint8) []byte {
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

// PrintGetDocumentData sends a checked request.
func PrintGetDocumentData(c *xgb.XConn, Context Pcontext, MaxBytes uint32) (PrintGetDocumentDataReply, error) {
	var reply PrintGetDocumentDataReply
	op, ok := c.Ext("XpExtension")
	if !ok {
		return reply, errors.New("cannot issue request \"PrintGetDocumentData\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	err := c.SendRecv(printGetDocumentDataRequest(op, Context, MaxBytes), &reply)
	return reply, err
}

// PrintGetDocumentDataUnchecked sends an unchecked request.
func PrintGetDocumentDataUnchecked(c *xgb.XConn, Context Pcontext, MaxBytes uint32) error {
	op, ok := c.Ext("XpExtension")
	if !ok {
		return errors.New("cannot issue request \"PrintGetDocumentData\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	return c.Send(printGetDocumentDataRequest(op, Context, MaxBytes))
}

// PrintGetDocumentDataReply represents the data returned from a PrintGetDocumentData request.
type PrintGetDocumentDataReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	StatusCode   uint32
	FinishedFlag uint32
	DataLen      uint32
	// padding: 12 bytes
	Data []byte // size: internal.Pad4((int(DataLen) * 1))
}

// Unmarshal reads a byte slice into a PrintGetDocumentDataReply value.
func (v *PrintGetDocumentDataReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.DataLen) * 1))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"PrintGetDocumentDataReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.StatusCode = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.FinishedFlag = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.DataLen = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 12 // padding

	v.Data = make([]byte, v.DataLen)
	copy(v.Data[:v.DataLen], buf[b:])
	b += int(v.DataLen)

	return nil
}

// Write request to wire for PrintGetDocumentData
// printGetDocumentDataRequest writes a PrintGetDocumentData request to a byte slice.
func printGetDocumentDataRequest(opcode uint8, Context Pcontext, MaxBytes uint32) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 12 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Context))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], MaxBytes)
	b += 4

	return buf
}

// PrintGetImageResolution sends a checked request.
func PrintGetImageResolution(c *xgb.XConn, Context Pcontext) (PrintGetImageResolutionReply, error) {
	var reply PrintGetImageResolutionReply
	op, ok := c.Ext("XpExtension")
	if !ok {
		return reply, errors.New("cannot issue request \"PrintGetImageResolution\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	err := c.SendRecv(printGetImageResolutionRequest(op, Context), &reply)
	return reply, err
}

// PrintGetImageResolutionUnchecked sends an unchecked request.
func PrintGetImageResolutionUnchecked(c *xgb.XConn, Context Pcontext) error {
	op, ok := c.Ext("XpExtension")
	if !ok {
		return errors.New("cannot issue request \"PrintGetImageResolution\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	return c.Send(printGetImageResolutionRequest(op, Context))
}

// PrintGetImageResolutionReply represents the data returned from a PrintGetImageResolution request.
type PrintGetImageResolutionReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	ImageResolution uint16
}

// Unmarshal reads a byte slice into a PrintGetImageResolutionReply value.
func (v *PrintGetImageResolutionReply) Unmarshal(buf []byte) error {
	if size := 10; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"PrintGetImageResolutionReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.ImageResolution = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return nil
}

// Write request to wire for PrintGetImageResolution
// printGetImageResolutionRequest writes a PrintGetImageResolution request to a byte slice.
func printGetImageResolutionRequest(opcode uint8, Context Pcontext) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 24 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Context))
	b += 4

	return buf
}

// PrintGetOneAttributes sends a checked request.
func PrintGetOneAttributes(c *xgb.XConn, Context Pcontext, NameLen uint32, Pool byte, Name []String8) (PrintGetOneAttributesReply, error) {
	var reply PrintGetOneAttributesReply
	op, ok := c.Ext("XpExtension")
	if !ok {
		return reply, errors.New("cannot issue request \"PrintGetOneAttributes\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	err := c.SendRecv(printGetOneAttributesRequest(op, Context, NameLen, Pool, Name), &reply)
	return reply, err
}

// PrintGetOneAttributesUnchecked sends an unchecked request.
func PrintGetOneAttributesUnchecked(c *xgb.XConn, Context Pcontext, NameLen uint32, Pool byte, Name []String8) error {
	op, ok := c.Ext("XpExtension")
	if !ok {
		return errors.New("cannot issue request \"PrintGetOneAttributes\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	return c.Send(printGetOneAttributesRequest(op, Context, NameLen, Pool, Name))
}

// PrintGetOneAttributesReply represents the data returned from a PrintGetOneAttributes request.
type PrintGetOneAttributesReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	ValueLen uint32
	// padding: 20 bytes
	Value []String8 // size: internal.Pad4((int(ValueLen) * 1))
}

// Unmarshal reads a byte slice into a PrintGetOneAttributesReply value.
func (v *PrintGetOneAttributesReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.ValueLen) * 1))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"PrintGetOneAttributesReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.ValueLen = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 20 // padding

	v.Value = make([]String8, v.ValueLen)
	for i := 0; i < int(v.ValueLen); i++ {
		v.Value[i] = String8(buf[b])
		b += 1
	}

	return nil
}

// Write request to wire for PrintGetOneAttributes
// printGetOneAttributesRequest writes a PrintGetOneAttributes request to a byte slice.
func printGetOneAttributesRequest(opcode uint8, Context Pcontext, NameLen uint32, Pool byte, Name []String8) []byte {
	size := internal.Pad4((16 + internal.Pad4((int(NameLen) * 1))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 19 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Context))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], NameLen)
	b += 4

	buf[b] = Pool
	b += 1

	b += 3 // padding

	for i := 0; i < int(NameLen); i++ {
		buf[b] = uint8(Name[i])
		b += 1
	}

	return buf
}

// PrintGetPageDimensions sends a checked request.
func PrintGetPageDimensions(c *xgb.XConn, Context Pcontext) (PrintGetPageDimensionsReply, error) {
	var reply PrintGetPageDimensionsReply
	op, ok := c.Ext("XpExtension")
	if !ok {
		return reply, errors.New("cannot issue request \"PrintGetPageDimensions\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	err := c.SendRecv(printGetPageDimensionsRequest(op, Context), &reply)
	return reply, err
}

// PrintGetPageDimensionsUnchecked sends an unchecked request.
func PrintGetPageDimensionsUnchecked(c *xgb.XConn, Context Pcontext) error {
	op, ok := c.Ext("XpExtension")
	if !ok {
		return errors.New("cannot issue request \"PrintGetPageDimensions\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	return c.Send(printGetPageDimensionsRequest(op, Context))
}

// PrintGetPageDimensionsReply represents the data returned from a PrintGetPageDimensions request.
type PrintGetPageDimensionsReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Width              uint16
	Height             uint16
	OffsetX            uint16
	OffsetY            uint16
	ReproducibleWidth  uint16
	ReproducibleHeight uint16
}

// Unmarshal reads a byte slice into a PrintGetPageDimensionsReply value.
func (v *PrintGetPageDimensionsReply) Unmarshal(buf []byte) error {
	if size := 20; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"PrintGetPageDimensionsReply\": have=%d need=%d", len(buf), size)
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

	v.OffsetX = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.OffsetY = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.ReproducibleWidth = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.ReproducibleHeight = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return nil
}

// Write request to wire for PrintGetPageDimensions
// printGetPageDimensionsRequest writes a PrintGetPageDimensions request to a byte slice.
func printGetPageDimensionsRequest(opcode uint8, Context Pcontext) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 21 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Context))
	b += 4

	return buf
}

// PrintGetPrinterList sends a checked request.
func PrintGetPrinterList(c *xgb.XConn, PrinterNameLen uint32, LocaleLen uint32, PrinterName []String8, Locale []String8) (PrintGetPrinterListReply, error) {
	var reply PrintGetPrinterListReply
	op, ok := c.Ext("XpExtension")
	if !ok {
		return reply, errors.New("cannot issue request \"PrintGetPrinterList\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	err := c.SendRecv(printGetPrinterListRequest(op, PrinterNameLen, LocaleLen, PrinterName, Locale), &reply)
	return reply, err
}

// PrintGetPrinterListUnchecked sends an unchecked request.
func PrintGetPrinterListUnchecked(c *xgb.XConn, PrinterNameLen uint32, LocaleLen uint32, PrinterName []String8, Locale []String8) error {
	op, ok := c.Ext("XpExtension")
	if !ok {
		return errors.New("cannot issue request \"PrintGetPrinterList\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	return c.Send(printGetPrinterListRequest(op, PrinterNameLen, LocaleLen, PrinterName, Locale))
}

// PrintGetPrinterListReply represents the data returned from a PrintGetPrinterList request.
type PrintGetPrinterListReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	ListCount uint32
	// padding: 20 bytes
	Printers []Printer // size: PrinterListSize(Printers)
}

// Unmarshal reads a byte slice into a PrintGetPrinterListReply value.
func (v *PrintGetPrinterListReply) Unmarshal(buf []byte) error {
	if size := (32 + PrinterListSize(v.Printers)); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"PrintGetPrinterListReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.ListCount = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 20 // padding

	v.Printers = make([]Printer, v.ListCount)
	b += PrinterReadList(buf[b:], v.Printers)

	return nil
}

// Write request to wire for PrintGetPrinterList
// printGetPrinterListRequest writes a PrintGetPrinterList request to a byte slice.
func printGetPrinterListRequest(opcode uint8, PrinterNameLen uint32, LocaleLen uint32, PrinterName []String8, Locale []String8) []byte {
	size := internal.Pad4(((12 + internal.Pad4((int(PrinterNameLen) * 1))) + internal.Pad4((int(LocaleLen) * 1))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 1 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], PrinterNameLen)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], LocaleLen)
	b += 4

	for i := 0; i < int(PrinterNameLen); i++ {
		buf[b] = uint8(PrinterName[i])
		b += 1
	}

	for i := 0; i < int(LocaleLen); i++ {
		buf[b] = uint8(Locale[i])
		b += 1
	}

	return buf
}

// PrintGetScreenOfContext sends a checked request.
func PrintGetScreenOfContext(c *xgb.XConn) (PrintGetScreenOfContextReply, error) {
	var reply PrintGetScreenOfContextReply
	op, ok := c.Ext("XpExtension")
	if !ok {
		return reply, errors.New("cannot issue request \"PrintGetScreenOfContext\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	err := c.SendRecv(printGetScreenOfContextRequest(op), &reply)
	return reply, err
}

// PrintGetScreenOfContextUnchecked sends an unchecked request.
func PrintGetScreenOfContextUnchecked(c *xgb.XConn) error {
	op, ok := c.Ext("XpExtension")
	if !ok {
		return errors.New("cannot issue request \"PrintGetScreenOfContext\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	return c.Send(printGetScreenOfContextRequest(op))
}

// PrintGetScreenOfContextReply represents the data returned from a PrintGetScreenOfContext request.
type PrintGetScreenOfContextReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Root xproto.Window
}

// Unmarshal reads a byte slice into a PrintGetScreenOfContextReply value.
func (v *PrintGetScreenOfContextReply) Unmarshal(buf []byte) error {
	if size := 12; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"PrintGetScreenOfContextReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Root = xproto.Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	return nil
}

// Write request to wire for PrintGetScreenOfContext
// printGetScreenOfContextRequest writes a PrintGetScreenOfContext request to a byte slice.
func printGetScreenOfContextRequest(opcode uint8) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 6 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// PrintInputSelected sends a checked request.
func PrintInputSelected(c *xgb.XConn, Context Pcontext) (PrintInputSelectedReply, error) {
	var reply PrintInputSelectedReply
	op, ok := c.Ext("XpExtension")
	if !ok {
		return reply, errors.New("cannot issue request \"PrintInputSelected\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	err := c.SendRecv(printInputSelectedRequest(op, Context), &reply)
	return reply, err
}

// PrintInputSelectedUnchecked sends an unchecked request.
func PrintInputSelectedUnchecked(c *xgb.XConn, Context Pcontext) error {
	op, ok := c.Ext("XpExtension")
	if !ok {
		return errors.New("cannot issue request \"PrintInputSelected\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	return c.Send(printInputSelectedRequest(op, Context))
}

// PrintInputSelectedReply represents the data returned from a PrintInputSelected request.
type PrintInputSelectedReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	EventMask     uint32
	AllEventsMask uint32
}

// Unmarshal reads a byte slice into a PrintInputSelectedReply value.
func (v *PrintInputSelectedReply) Unmarshal(buf []byte) error {
	if size := 16; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"PrintInputSelectedReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.EventMask = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.AllEventsMask = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	return nil
}

// Write request to wire for PrintInputSelected
// printInputSelectedRequest writes a PrintInputSelected request to a byte slice.
func printInputSelectedRequest(opcode uint8, Context Pcontext) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 16 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Context))
	b += 4

	return buf
}

// PrintPutDocumentData sends a checked request.
func PrintPutDocumentData(c *xgb.XConn, Drawable xproto.Drawable, LenData uint32, LenFmt uint16, LenOptions uint16, Data []byte, DocFormat []String8, Options []String8) error {
	op, ok := c.Ext("XpExtension")
	if !ok {
		return errors.New("cannot issue request \"PrintPutDocumentData\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	return c.SendRecv(printPutDocumentDataRequest(op, Drawable, LenData, LenFmt, LenOptions, Data, DocFormat, Options), nil)
}

// PrintPutDocumentDataUnchecked sends an unchecked request.
func PrintPutDocumentDataUnchecked(c *xgb.XConn, Drawable xproto.Drawable, LenData uint32, LenFmt uint16, LenOptions uint16, Data []byte, DocFormat []String8, Options []String8) error {
	op, ok := c.Ext("XpExtension")
	if !ok {
		return errors.New("cannot issue request \"PrintPutDocumentData\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	return c.Send(printPutDocumentDataRequest(op, Drawable, LenData, LenFmt, LenOptions, Data, DocFormat, Options))
}

// Write request to wire for PrintPutDocumentData
// printPutDocumentDataRequest writes a PrintPutDocumentData request to a byte slice.
func printPutDocumentDataRequest(opcode uint8, Drawable xproto.Drawable, LenData uint32, LenFmt uint16, LenOptions uint16, Data []byte, DocFormat []String8, Options []String8) []byte {
	size := internal.Pad4((((16 + internal.Pad4((int(LenData) * 1))) + internal.Pad4((int(LenFmt) * 1))) + internal.Pad4((int(LenOptions) * 1))))
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

	binary.LittleEndian.PutUint32(buf[b:], LenData)
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], LenFmt)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], LenOptions)
	b += 2

	copy(buf[b:], Data[:LenData])
	b += int(LenData)

	for i := 0; i < int(LenFmt); i++ {
		buf[b] = uint8(DocFormat[i])
		b += 1
	}

	for i := 0; i < int(LenOptions); i++ {
		buf[b] = uint8(Options[i])
		b += 1
	}

	return buf
}

// PrintQueryScreens sends a checked request.
func PrintQueryScreens(c *xgb.XConn) (PrintQueryScreensReply, error) {
	var reply PrintQueryScreensReply
	op, ok := c.Ext("XpExtension")
	if !ok {
		return reply, errors.New("cannot issue request \"PrintQueryScreens\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	err := c.SendRecv(printQueryScreensRequest(op), &reply)
	return reply, err
}

// PrintQueryScreensUnchecked sends an unchecked request.
func PrintQueryScreensUnchecked(c *xgb.XConn) error {
	op, ok := c.Ext("XpExtension")
	if !ok {
		return errors.New("cannot issue request \"PrintQueryScreens\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	return c.Send(printQueryScreensRequest(op))
}

// PrintQueryScreensReply represents the data returned from a PrintQueryScreens request.
type PrintQueryScreensReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	ListCount uint32
	// padding: 20 bytes
	Roots []xproto.Window // size: internal.Pad4((int(ListCount) * 4))
}

// Unmarshal reads a byte slice into a PrintQueryScreensReply value.
func (v *PrintQueryScreensReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.ListCount) * 4))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"PrintQueryScreensReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.ListCount = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 20 // padding

	v.Roots = make([]xproto.Window, v.ListCount)
	for i := 0; i < int(v.ListCount); i++ {
		v.Roots[i] = xproto.Window(binary.LittleEndian.Uint32(buf[b:]))
		b += 4
	}

	return nil
}

// Write request to wire for PrintQueryScreens
// printQueryScreensRequest writes a PrintQueryScreens request to a byte slice.
func printQueryScreensRequest(opcode uint8) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 22 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// PrintQueryVersion sends a checked request.
func PrintQueryVersion(c *xgb.XConn) (PrintQueryVersionReply, error) {
	var reply PrintQueryVersionReply
	op, ok := c.Ext("XpExtension")
	if !ok {
		return reply, errors.New("cannot issue request \"PrintQueryVersion\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	err := c.SendRecv(printQueryVersionRequest(op), &reply)
	return reply, err
}

// PrintQueryVersionUnchecked sends an unchecked request.
func PrintQueryVersionUnchecked(c *xgb.XConn) error {
	op, ok := c.Ext("XpExtension")
	if !ok {
		return errors.New("cannot issue request \"PrintQueryVersion\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	return c.Send(printQueryVersionRequest(op))
}

// PrintQueryVersionReply represents the data returned from a PrintQueryVersion request.
type PrintQueryVersionReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	MajorVersion uint16
	MinorVersion uint16
}

// Unmarshal reads a byte slice into a PrintQueryVersionReply value.
func (v *PrintQueryVersionReply) Unmarshal(buf []byte) error {
	if size := 12; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"PrintQueryVersionReply\": have=%d need=%d", len(buf), size)
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

// Write request to wire for PrintQueryVersion
// printQueryVersionRequest writes a PrintQueryVersion request to a byte slice.
func printQueryVersionRequest(opcode uint8) []byte {
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

// PrintRehashPrinterList sends a checked request.
func PrintRehashPrinterList(c *xgb.XConn) error {
	op, ok := c.Ext("XpExtension")
	if !ok {
		return errors.New("cannot issue request \"PrintRehashPrinterList\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	return c.SendRecv(printRehashPrinterListRequest(op), nil)
}

// PrintRehashPrinterListUnchecked sends an unchecked request.
func PrintRehashPrinterListUnchecked(c *xgb.XConn) error {
	op, ok := c.Ext("XpExtension")
	if !ok {
		return errors.New("cannot issue request \"PrintRehashPrinterList\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	return c.Send(printRehashPrinterListRequest(op))
}

// Write request to wire for PrintRehashPrinterList
// printRehashPrinterListRequest writes a PrintRehashPrinterList request to a byte slice.
func printRehashPrinterListRequest(opcode uint8) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 20 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// PrintSelectInput sends a checked request.
func PrintSelectInput(c *xgb.XConn, Context Pcontext, EventMask uint32) error {
	op, ok := c.Ext("XpExtension")
	if !ok {
		return errors.New("cannot issue request \"PrintSelectInput\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	return c.SendRecv(printSelectInputRequest(op, Context, EventMask), nil)
}

// PrintSelectInputUnchecked sends an unchecked request.
func PrintSelectInputUnchecked(c *xgb.XConn, Context Pcontext, EventMask uint32) error {
	op, ok := c.Ext("XpExtension")
	if !ok {
		return errors.New("cannot issue request \"PrintSelectInput\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	return c.Send(printSelectInputRequest(op, Context, EventMask))
}

// Write request to wire for PrintSelectInput
// printSelectInputRequest writes a PrintSelectInput request to a byte slice.
func printSelectInputRequest(opcode uint8, Context Pcontext, EventMask uint32) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 15 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Context))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], EventMask)
	b += 4

	return buf
}

// PrintSetAttributes sends a checked request.
func PrintSetAttributes(c *xgb.XConn, Context Pcontext, StringLen uint32, Pool byte, Rule byte, Attributes []String8) error {
	op, ok := c.Ext("XpExtension")
	if !ok {
		return errors.New("cannot issue request \"PrintSetAttributes\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	return c.SendRecv(printSetAttributesRequest(op, Context, StringLen, Pool, Rule, Attributes), nil)
}

// PrintSetAttributesUnchecked sends an unchecked request.
func PrintSetAttributesUnchecked(c *xgb.XConn, Context Pcontext, StringLen uint32, Pool byte, Rule byte, Attributes []String8) error {
	op, ok := c.Ext("XpExtension")
	if !ok {
		return errors.New("cannot issue request \"PrintSetAttributes\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	return c.Send(printSetAttributesRequest(op, Context, StringLen, Pool, Rule, Attributes))
}

// Write request to wire for PrintSetAttributes
// printSetAttributesRequest writes a PrintSetAttributes request to a byte slice.
func printSetAttributesRequest(opcode uint8, Context Pcontext, StringLen uint32, Pool byte, Rule byte, Attributes []String8) []byte {
	size := internal.Pad4((16 + internal.Pad4((len(Attributes) * 1))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 18 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Context))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], StringLen)
	b += 4

	buf[b] = Pool
	b += 1

	buf[b] = Rule
	b += 1

	b += 2 // padding

	for i := 0; i < int(len(Attributes)); i++ {
		buf[b] = uint8(Attributes[i])
		b += 1
	}

	return buf
}

// PrintSetContext sends a checked request.
func PrintSetContext(c *xgb.XConn, Context uint32) error {
	op, ok := c.Ext("XpExtension")
	if !ok {
		return errors.New("cannot issue request \"PrintSetContext\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	return c.SendRecv(printSetContextRequest(op, Context), nil)
}

// PrintSetContextUnchecked sends an unchecked request.
func PrintSetContextUnchecked(c *xgb.XConn, Context uint32) error {
	op, ok := c.Ext("XpExtension")
	if !ok {
		return errors.New("cannot issue request \"PrintSetContext\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	return c.Send(printSetContextRequest(op, Context))
}

// Write request to wire for PrintSetContext
// printSetContextRequest writes a PrintSetContext request to a byte slice.
func printSetContextRequest(opcode uint8, Context uint32) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 3 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], Context)
	b += 4

	return buf
}

// PrintSetImageResolution sends a checked request.
func PrintSetImageResolution(c *xgb.XConn, Context Pcontext, ImageResolution uint16) (PrintSetImageResolutionReply, error) {
	var reply PrintSetImageResolutionReply
	op, ok := c.Ext("XpExtension")
	if !ok {
		return reply, errors.New("cannot issue request \"PrintSetImageResolution\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	err := c.SendRecv(printSetImageResolutionRequest(op, Context, ImageResolution), &reply)
	return reply, err
}

// PrintSetImageResolutionUnchecked sends an unchecked request.
func PrintSetImageResolutionUnchecked(c *xgb.XConn, Context Pcontext, ImageResolution uint16) error {
	op, ok := c.Ext("XpExtension")
	if !ok {
		return errors.New("cannot issue request \"PrintSetImageResolution\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	return c.Send(printSetImageResolutionRequest(op, Context, ImageResolution))
}

// PrintSetImageResolutionReply represents the data returned from a PrintSetImageResolution request.
type PrintSetImageResolutionReply struct {
	Sequence            uint16 // sequence number of the request for this reply
	Length              uint32 // number of bytes in this reply
	Status              bool
	PreviousResolutions uint16
}

// Unmarshal reads a byte slice into a PrintSetImageResolutionReply value.
func (v *PrintSetImageResolutionReply) Unmarshal(buf []byte) error {
	if size := 10; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"PrintSetImageResolutionReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	v.Status = (buf[b] == 1)
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.PreviousResolutions = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return nil
}

// Write request to wire for PrintSetImageResolution
// printSetImageResolutionRequest writes a PrintSetImageResolution request to a byte slice.
func printSetImageResolutionRequest(opcode uint8, Context Pcontext, ImageResolution uint16) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 23 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Context))
	b += 4

	binary.LittleEndian.PutUint16(buf[b:], ImageResolution)
	b += 2

	return buf
}

// PrintStartDoc sends a checked request.
func PrintStartDoc(c *xgb.XConn, DriverMode byte) error {
	op, ok := c.Ext("XpExtension")
	if !ok {
		return errors.New("cannot issue request \"PrintStartDoc\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	return c.SendRecv(printStartDocRequest(op, DriverMode), nil)
}

// PrintStartDocUnchecked sends an unchecked request.
func PrintStartDocUnchecked(c *xgb.XConn, DriverMode byte) error {
	op, ok := c.Ext("XpExtension")
	if !ok {
		return errors.New("cannot issue request \"PrintStartDoc\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	return c.Send(printStartDocRequest(op, DriverMode))
}

// Write request to wire for PrintStartDoc
// printStartDocRequest writes a PrintStartDoc request to a byte slice.
func printStartDocRequest(opcode uint8, DriverMode byte) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 9 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	buf[b] = DriverMode
	b += 1

	return buf
}

// PrintStartJob sends a checked request.
func PrintStartJob(c *xgb.XConn, OutputMode byte) error {
	op, ok := c.Ext("XpExtension")
	if !ok {
		return errors.New("cannot issue request \"PrintStartJob\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	return c.SendRecv(printStartJobRequest(op, OutputMode), nil)
}

// PrintStartJobUnchecked sends an unchecked request.
func PrintStartJobUnchecked(c *xgb.XConn, OutputMode byte) error {
	op, ok := c.Ext("XpExtension")
	if !ok {
		return errors.New("cannot issue request \"PrintStartJob\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	return c.Send(printStartJobRequest(op, OutputMode))
}

// Write request to wire for PrintStartJob
// printStartJobRequest writes a PrintStartJob request to a byte slice.
func printStartJobRequest(opcode uint8, OutputMode byte) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 7 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	buf[b] = OutputMode
	b += 1

	return buf
}

// PrintStartPage sends a checked request.
func PrintStartPage(c *xgb.XConn, Window xproto.Window) error {
	op, ok := c.Ext("XpExtension")
	if !ok {
		return errors.New("cannot issue request \"PrintStartPage\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	return c.SendRecv(printStartPageRequest(op, Window), nil)
}

// PrintStartPageUnchecked sends an unchecked request.
func PrintStartPageUnchecked(c *xgb.XConn, Window xproto.Window) error {
	op, ok := c.Ext("XpExtension")
	if !ok {
		return errors.New("cannot issue request \"PrintStartPage\" using the uninitialized extension \"XpExtension\". xprint.Register(xconn) must be called first.")
	}
	return c.Send(printStartPageRequest(op, Window))
}

// Write request to wire for PrintStartPage
// printStartPageRequest writes a PrintStartPage request to a byte slice.
func printStartPageRequest(opcode uint8, Window xproto.Window) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 13 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	return buf
}
