// FILE GENERATED AUTOMATICALLY FROM "record.xml"
package record

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
	ExtName = "Record"

	// ExtXName is the name string this extension is known by to the X server.
	ExtXName = "RECORD"
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
		return fmt.Errorf("error querying X for \"Record\": %w", err)
	} else if !reply.Present {
		return fmt.Errorf("no extension named \"Record\" is known to the X server: reply=%+v", reply)
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

// BadBadContext is the error number for a BadBadContext.
const BadBadContext = 0

type BadContextError struct {
	Sequence      uint16
	NiceName      string
	InvalidRecord uint32
}

// UnmarshalBadContextError constructs a BadContextError value that implements xgb.Error from a byte slice.
func UnmarshalBadContextError(buf []byte) (xgb.XError, error) {
	if len(buf) != 32 {
		return nil, fmt.Errorf("invalid data size %d for \"BadContextError\"", len(buf))
	}

	v := &BadContextError{}
	v.NiceName = "BadContext"

	b := 1 // skip error determinant
	b += 1 // don't read error number

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.InvalidRecord = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	return v, nil
}

// SeqID returns the sequence id attached to the BadBadContext error.
// This is mostly used internally.
func (err *BadContextError) SeqID() uint16 {
	return err.Sequence
}

// BadID returns the 'BadValue' number if one exists for the BadBadContext error. If no bad value exists, 0 is returned.
func (err *BadContextError) BadID() uint32 {
	return 0
}

// Error returns a rudimentary string representation of the BadBadContext error.
func (err *BadContextError) Error() string {
	var buf strings.Builder

	buf.WriteString("BadBadContext{")
	buf.WriteString("NiceName: " + err.NiceName)
	buf.WriteByte(' ')
	buf.WriteString("Sequence: " + strconv.FormatUint(uint64(err.Sequence), 10))
	buf.WriteByte(' ')

	fmt.Fprintf(&buf, "InvalidRecord: %d", err.InvalidRecord)
	buf.WriteByte('}')

	return buf.String()
}

func init() { registerError(0, UnmarshalBadContextError) }

type ClientInfo struct {
	ClientResource ClientSpec
	NumRanges      uint32
	Ranges         []Range // size: internal.Pad4((int(NumRanges) * 24))
}

// ClientInfoRead reads a byte slice into a ClientInfo value.
func ClientInfoRead(buf []byte, v *ClientInfo) int {
	b := 0

	v.ClientResource = ClientSpec(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.NumRanges = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.Ranges = make([]Range, v.NumRanges)
	b += RangeReadList(buf[b:], v.Ranges)

	return b
}

// ClientInfoReadList reads a byte slice into a list of ClientInfo values.
func ClientInfoReadList(buf []byte, dest []ClientInfo) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = ClientInfo{}
		b += ClientInfoRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a ClientInfo value to a byte slice.
func (v ClientInfo) Bytes() []byte {
	buf := make([]byte, (8 + internal.Pad4((int(v.NumRanges) * 24))))
	b := 0

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.ClientResource))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], v.NumRanges)
	b += 4

	b += RangeListBytes(buf[b:], v.Ranges)

	return buf[:b]
}

// ClientInfoListBytes writes a list of ClientInfo values to a byte slice.
func ClientInfoListBytes(buf []byte, list []ClientInfo) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

// ClientInfoListSize computes the size (bytes) of a list of ClientInfo values.
func ClientInfoListSize(list []ClientInfo) int {
	size := 0
	for _, item := range list {
		size += (8 + internal.Pad4((int(item.NumRanges) * 24)))
	}
	return size
}

type ClientSpec uint32

type Context uint32

func NewContextID(c *xgb.XConn) Context {
	id := c.NewXID()
	return Context(id)
}

const (
	CsCurrentClients = 1
	CsFutureClients  = 2
	CsAllClients     = 3
)

type ElementHeader byte

type ExtRange struct {
	Major Range8
	Minor Range16
}

// ExtRangeRead reads a byte slice into a ExtRange value.
func ExtRangeRead(buf []byte, v *ExtRange) int {
	b := 0

	v.Major = Range8{}
	b += Range8Read(buf[b:], &v.Major)

	v.Minor = Range16{}
	b += Range16Read(buf[b:], &v.Minor)

	return b
}

// ExtRangeReadList reads a byte slice into a list of ExtRange values.
func ExtRangeReadList(buf []byte, dest []ExtRange) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = ExtRange{}
		b += ExtRangeRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a ExtRange value to a byte slice.
func (v ExtRange) Bytes() []byte {
	buf := make([]byte, 6)
	b := 0

	{
		structBytes := v.Major.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	{
		structBytes := v.Minor.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	return buf[:b]
}

// ExtRangeListBytes writes a list of ExtRange values to a byte slice.
func ExtRangeListBytes(buf []byte, list []ExtRange) int {
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
	HTypeFromServerTime     = 1
	HTypeFromClientTime     = 2
	HTypeFromClientSequence = 4
)

type Range struct {
	CoreRequests    Range8
	CoreReplies     Range8
	ExtRequests     ExtRange
	ExtReplies      ExtRange
	DeliveredEvents Range8
	DeviceEvents    Range8
	Errors          Range8
	ClientStarted   bool
	ClientDied      bool
}

// RangeRead reads a byte slice into a Range value.
func RangeRead(buf []byte, v *Range) int {
	b := 0

	v.CoreRequests = Range8{}
	b += Range8Read(buf[b:], &v.CoreRequests)

	v.CoreReplies = Range8{}
	b += Range8Read(buf[b:], &v.CoreReplies)

	v.ExtRequests = ExtRange{}
	b += ExtRangeRead(buf[b:], &v.ExtRequests)

	v.ExtReplies = ExtRange{}
	b += ExtRangeRead(buf[b:], &v.ExtReplies)

	v.DeliveredEvents = Range8{}
	b += Range8Read(buf[b:], &v.DeliveredEvents)

	v.DeviceEvents = Range8{}
	b += Range8Read(buf[b:], &v.DeviceEvents)

	v.Errors = Range8{}
	b += Range8Read(buf[b:], &v.Errors)

	v.ClientStarted = (buf[b] == 1)
	b += 1

	v.ClientDied = (buf[b] == 1)
	b += 1

	return b
}

// RangeReadList reads a byte slice into a list of Range values.
func RangeReadList(buf []byte, dest []Range) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Range{}
		b += RangeRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a Range value to a byte slice.
func (v Range) Bytes() []byte {
	buf := make([]byte, 24)
	b := 0

	{
		structBytes := v.CoreRequests.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	{
		structBytes := v.CoreReplies.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	{
		structBytes := v.ExtRequests.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	{
		structBytes := v.ExtReplies.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	{
		structBytes := v.DeliveredEvents.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	{
		structBytes := v.DeviceEvents.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	{
		structBytes := v.Errors.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	if v.ClientStarted {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	if v.ClientDied {
		buf[b] = 1
	} else {
		buf[b] = 0
	}
	b += 1

	return buf[:b]
}

// RangeListBytes writes a list of Range values to a byte slice.
func RangeListBytes(buf []byte, list []Range) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

type Range16 struct {
	First uint16
	Last  uint16
}

// Range16Read reads a byte slice into a Range16 value.
func Range16Read(buf []byte, v *Range16) int {
	b := 0

	v.First = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Last = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return b
}

// Range16ReadList reads a byte slice into a list of Range16 values.
func Range16ReadList(buf []byte, dest []Range16) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Range16{}
		b += Range16Read(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a Range16 value to a byte slice.
func (v Range16) Bytes() []byte {
	buf := make([]byte, 4)
	b := 0

	binary.LittleEndian.PutUint16(buf[b:], v.First)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Last)
	b += 2

	return buf[:b]
}

// Range16ListBytes writes a list of Range16 values to a byte slice.
func Range16ListBytes(buf []byte, list []Range16) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

type Range8 struct {
	First byte
	Last  byte
}

// Range8Read reads a byte slice into a Range8 value.
func Range8Read(buf []byte, v *Range8) int {
	b := 0

	v.First = buf[b]
	b += 1

	v.Last = buf[b]
	b += 1

	return b
}

// Range8ReadList reads a byte slice into a list of Range8 values.
func Range8ReadList(buf []byte, dest []Range8) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Range8{}
		b += Range8Read(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a Range8 value to a byte slice.
func (v Range8) Bytes() []byte {
	buf := make([]byte, 2)
	b := 0

	buf[b] = v.First
	b += 1

	buf[b] = v.Last
	b += 1

	return buf[:b]
}

// Range8ListBytes writes a list of Range8 values to a byte slice.
func Range8ListBytes(buf []byte, list []Range8) int {
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
func CreateContext(c *xgb.XConn, Context Context, ElementHeader ElementHeader, NumClientSpecs uint32, NumRanges uint32, ClientSpecs []ClientSpec, Ranges []Range) error {
	op, ok := c.Ext("RECORD")
	if !ok {
		return errors.New("cannot issue request \"CreateContext\" using the uninitialized extension \"RECORD\". record.Register(xconn) must be called first.")
	}
	return c.SendRecv(createContextRequest(op, Context, ElementHeader, NumClientSpecs, NumRanges, ClientSpecs, Ranges), nil)
}

// CreateContextUnchecked sends an unchecked request.
func CreateContextUnchecked(c *xgb.XConn, Context Context, ElementHeader ElementHeader, NumClientSpecs uint32, NumRanges uint32, ClientSpecs []ClientSpec, Ranges []Range) error {
	op, ok := c.Ext("RECORD")
	if !ok {
		return errors.New("cannot issue request \"CreateContext\" using the uninitialized extension \"RECORD\". record.Register(xconn) must be called first.")
	}
	return c.Send(createContextRequest(op, Context, ElementHeader, NumClientSpecs, NumRanges, ClientSpecs, Ranges))
}

// Write request to wire for CreateContext
// createContextRequest writes a CreateContext request to a byte slice.
func createContextRequest(opcode uint8, Context Context, ElementHeader ElementHeader, NumClientSpecs uint32, NumRanges uint32, ClientSpecs []ClientSpec, Ranges []Range) []byte {
	size := internal.Pad4((((20 + internal.Pad4((int(NumClientSpecs) * 4))) + 4) + internal.Pad4((int(NumRanges) * 24))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 1 // request opcode
	b += 1

	blen := b
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Context))
	b += 4

	buf[b] = uint8(ElementHeader)
	b += 1

	b += 3 // padding

	binary.LittleEndian.PutUint32(buf[b:], NumClientSpecs)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], NumRanges)
	b += 4

	for i := 0; i < int(NumClientSpecs); i++ {
		binary.LittleEndian.PutUint32(buf[b:], uint32(ClientSpecs[i]))
		b += 4
	}

	b = (b + 3) & ^3 // alignment gap

	b += RangeListBytes(buf[b:], Ranges)

	b = internal.Pad4(b)
	binary.LittleEndian.PutUint16(buf[blen:], uint16(b/4)) // write request size in 4-byte units
	return buf[:b]
}

// DisableContext sends a checked request.
func DisableContext(c *xgb.XConn, Context Context) error {
	op, ok := c.Ext("RECORD")
	if !ok {
		return errors.New("cannot issue request \"DisableContext\" using the uninitialized extension \"RECORD\". record.Register(xconn) must be called first.")
	}
	return c.SendRecv(disableContextRequest(op, Context), nil)
}

// DisableContextUnchecked sends an unchecked request.
func DisableContextUnchecked(c *xgb.XConn, Context Context) error {
	op, ok := c.Ext("RECORD")
	if !ok {
		return errors.New("cannot issue request \"DisableContext\" using the uninitialized extension \"RECORD\". record.Register(xconn) must be called first.")
	}
	return c.Send(disableContextRequest(op, Context))
}

// Write request to wire for DisableContext
// disableContextRequest writes a DisableContext request to a byte slice.
func disableContextRequest(opcode uint8, Context Context) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 6 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Context))
	b += 4

	return buf
}

// EnableContext sends a checked request.
func EnableContext(c *xgb.XConn, Context Context) (EnableContextReply, error) {
	var reply EnableContextReply
	op, ok := c.Ext("RECORD")
	if !ok {
		return reply, errors.New("cannot issue request \"EnableContext\" using the uninitialized extension \"RECORD\". record.Register(xconn) must be called first.")
	}
	err := c.SendRecv(enableContextRequest(op, Context), &reply)
	return reply, err
}

// EnableContextUnchecked sends an unchecked request.
func EnableContextUnchecked(c *xgb.XConn, Context Context) error {
	op, ok := c.Ext("RECORD")
	if !ok {
		return errors.New("cannot issue request \"EnableContext\" using the uninitialized extension \"RECORD\". record.Register(xconn) must be called first.")
	}
	return c.Send(enableContextRequest(op, Context))
}

// EnableContextReply represents the data returned from a EnableContext request.
type EnableContextReply struct {
	Sequence      uint16 // sequence number of the request for this reply
	Length        uint32 // number of bytes in this reply
	Category      byte
	ElementHeader ElementHeader
	ClientSwapped bool
	// padding: 2 bytes
	XidBase        uint32
	ServerTime     uint32
	RecSequenceNum uint32
	// padding: 8 bytes
	Data []byte // size: internal.Pad4(((int(Length) * 4) * 1))
}

// Unmarshal reads a byte slice into a EnableContextReply value.
func (v *EnableContextReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4(((int(v.Length) * 4) * 1))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"EnableContextReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	v.Category = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.ElementHeader = ElementHeader(buf[b])
	b += 1

	v.ClientSwapped = (buf[b] == 1)
	b += 1

	b += 2 // padding

	v.XidBase = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.ServerTime = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.RecSequenceNum = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 8 // padding

	v.Data = make([]byte, (int(v.Length) * 4))
	copy(v.Data[:(int(v.Length)*4)], buf[b:])
	b += int((int(v.Length) * 4))

	return nil
}

// Write request to wire for EnableContext
// enableContextRequest writes a EnableContext request to a byte slice.
func enableContextRequest(opcode uint8, Context Context) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 5 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Context))
	b += 4

	return buf
}

// FreeContext sends a checked request.
func FreeContext(c *xgb.XConn, Context Context) error {
	op, ok := c.Ext("RECORD")
	if !ok {
		return errors.New("cannot issue request \"FreeContext\" using the uninitialized extension \"RECORD\". record.Register(xconn) must be called first.")
	}
	return c.SendRecv(freeContextRequest(op, Context), nil)
}

// FreeContextUnchecked sends an unchecked request.
func FreeContextUnchecked(c *xgb.XConn, Context Context) error {
	op, ok := c.Ext("RECORD")
	if !ok {
		return errors.New("cannot issue request \"FreeContext\" using the uninitialized extension \"RECORD\". record.Register(xconn) must be called first.")
	}
	return c.Send(freeContextRequest(op, Context))
}

// Write request to wire for FreeContext
// freeContextRequest writes a FreeContext request to a byte slice.
func freeContextRequest(opcode uint8, Context Context) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 7 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Context))
	b += 4

	return buf
}

// GetContext sends a checked request.
func GetContext(c *xgb.XConn, Context Context) (GetContextReply, error) {
	var reply GetContextReply
	op, ok := c.Ext("RECORD")
	if !ok {
		return reply, errors.New("cannot issue request \"GetContext\" using the uninitialized extension \"RECORD\". record.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getContextRequest(op, Context), &reply)
	return reply, err
}

// GetContextUnchecked sends an unchecked request.
func GetContextUnchecked(c *xgb.XConn, Context Context) error {
	op, ok := c.Ext("RECORD")
	if !ok {
		return errors.New("cannot issue request \"GetContext\" using the uninitialized extension \"RECORD\". record.Register(xconn) must be called first.")
	}
	return c.Send(getContextRequest(op, Context))
}

// GetContextReply represents the data returned from a GetContext request.
type GetContextReply struct {
	Sequence      uint16 // sequence number of the request for this reply
	Length        uint32 // number of bytes in this reply
	Enabled       bool
	ElementHeader ElementHeader
	// padding: 3 bytes
	NumInterceptedClients uint32
	// padding: 16 bytes
	InterceptedClients []ClientInfo // size: ClientInfoListSize(InterceptedClients)
}

// Unmarshal reads a byte slice into a GetContextReply value.
func (v *GetContextReply) Unmarshal(buf []byte) error {
	if size := (32 + ClientInfoListSize(v.InterceptedClients)); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetContextReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	v.Enabled = (buf[b] == 1)
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.ElementHeader = ElementHeader(buf[b])
	b += 1

	b += 3 // padding

	v.NumInterceptedClients = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 16 // padding

	v.InterceptedClients = make([]ClientInfo, v.NumInterceptedClients)
	b += ClientInfoReadList(buf[b:], v.InterceptedClients)

	return nil
}

// Write request to wire for GetContext
// getContextRequest writes a GetContext request to a byte slice.
func getContextRequest(opcode uint8, Context Context) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 4 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Context))
	b += 4

	return buf
}

// QueryVersion sends a checked request.
func QueryVersion(c *xgb.XConn, MajorVersion uint16, MinorVersion uint16) (QueryVersionReply, error) {
	var reply QueryVersionReply
	op, ok := c.Ext("RECORD")
	if !ok {
		return reply, errors.New("cannot issue request \"QueryVersion\" using the uninitialized extension \"RECORD\". record.Register(xconn) must be called first.")
	}
	err := c.SendRecv(queryVersionRequest(op, MajorVersion, MinorVersion), &reply)
	return reply, err
}

// QueryVersionUnchecked sends an unchecked request.
func QueryVersionUnchecked(c *xgb.XConn, MajorVersion uint16, MinorVersion uint16) error {
	op, ok := c.Ext("RECORD")
	if !ok {
		return errors.New("cannot issue request \"QueryVersion\" using the uninitialized extension \"RECORD\". record.Register(xconn) must be called first.")
	}
	return c.Send(queryVersionRequest(op, MajorVersion, MinorVersion))
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
	if size := 12; len(buf) < size {
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
func queryVersionRequest(opcode uint8, MajorVersion uint16, MinorVersion uint16) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 0 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], MajorVersion)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], MinorVersion)
	b += 2

	return buf
}

// RegisterClients sends a checked request.
func RegisterClients(c *xgb.XConn, Context Context, ElementHeader ElementHeader, NumClientSpecs uint32, NumRanges uint32, ClientSpecs []ClientSpec, Ranges []Range) error {
	op, ok := c.Ext("RECORD")
	if !ok {
		return errors.New("cannot issue request \"RegisterClients\" using the uninitialized extension \"RECORD\". record.Register(xconn) must be called first.")
	}
	return c.SendRecv(registerClientsRequest(op, Context, ElementHeader, NumClientSpecs, NumRanges, ClientSpecs, Ranges), nil)
}

// RegisterClientsUnchecked sends an unchecked request.
func RegisterClientsUnchecked(c *xgb.XConn, Context Context, ElementHeader ElementHeader, NumClientSpecs uint32, NumRanges uint32, ClientSpecs []ClientSpec, Ranges []Range) error {
	op, ok := c.Ext("RECORD")
	if !ok {
		return errors.New("cannot issue request \"RegisterClients\" using the uninitialized extension \"RECORD\". record.Register(xconn) must be called first.")
	}
	return c.Send(registerClientsRequest(op, Context, ElementHeader, NumClientSpecs, NumRanges, ClientSpecs, Ranges))
}

// Write request to wire for RegisterClients
// registerClientsRequest writes a RegisterClients request to a byte slice.
func registerClientsRequest(opcode uint8, Context Context, ElementHeader ElementHeader, NumClientSpecs uint32, NumRanges uint32, ClientSpecs []ClientSpec, Ranges []Range) []byte {
	size := internal.Pad4((((20 + internal.Pad4((int(NumClientSpecs) * 4))) + 4) + internal.Pad4((int(NumRanges) * 24))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 2 // request opcode
	b += 1

	blen := b
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Context))
	b += 4

	buf[b] = uint8(ElementHeader)
	b += 1

	b += 3 // padding

	binary.LittleEndian.PutUint32(buf[b:], NumClientSpecs)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], NumRanges)
	b += 4

	for i := 0; i < int(NumClientSpecs); i++ {
		binary.LittleEndian.PutUint32(buf[b:], uint32(ClientSpecs[i]))
		b += 4
	}

	b = (b + 3) & ^3 // alignment gap

	b += RangeListBytes(buf[b:], Ranges)

	b = internal.Pad4(b)
	binary.LittleEndian.PutUint16(buf[blen:], uint16(b/4)) // write request size in 4-byte units
	return buf[:b]
}

// UnregisterClients sends a checked request.
func UnregisterClients(c *xgb.XConn, Context Context, NumClientSpecs uint32, ClientSpecs []ClientSpec) error {
	op, ok := c.Ext("RECORD")
	if !ok {
		return errors.New("cannot issue request \"UnregisterClients\" using the uninitialized extension \"RECORD\". record.Register(xconn) must be called first.")
	}
	return c.SendRecv(unregisterClientsRequest(op, Context, NumClientSpecs, ClientSpecs), nil)
}

// UnregisterClientsUnchecked sends an unchecked request.
func UnregisterClientsUnchecked(c *xgb.XConn, Context Context, NumClientSpecs uint32, ClientSpecs []ClientSpec) error {
	op, ok := c.Ext("RECORD")
	if !ok {
		return errors.New("cannot issue request \"UnregisterClients\" using the uninitialized extension \"RECORD\". record.Register(xconn) must be called first.")
	}
	return c.Send(unregisterClientsRequest(op, Context, NumClientSpecs, ClientSpecs))
}

// Write request to wire for UnregisterClients
// unregisterClientsRequest writes a UnregisterClients request to a byte slice.
func unregisterClientsRequest(opcode uint8, Context Context, NumClientSpecs uint32, ClientSpecs []ClientSpec) []byte {
	size := internal.Pad4((12 + internal.Pad4((int(NumClientSpecs) * 4))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 3 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Context))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], NumClientSpecs)
	b += 4

	for i := 0; i < int(NumClientSpecs); i++ {
		binary.LittleEndian.PutUint32(buf[b:], uint32(ClientSpecs[i]))
		b += 4
	}

	return buf
}
