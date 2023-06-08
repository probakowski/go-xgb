// FILE GENERATED AUTOMATICALLY FROM "xevie.xml"
package xevie

import (
	"encoding/binary"
	"errors"
	"fmt"

	"codeberg.org/gruf/go-xgb"
	"codeberg.org/gruf/go-xgb/internal"
	"codeberg.org/gruf/go-xgb/xproto"
)

const (
	// ExtName is the user-friendly name string of this X extension.
	ExtName = "Xevie"

	// ExtXName is the name string this extension is known by to the X server.
	ExtXName = "XEVIE"
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

// Register will query the X server for Xevie extension support, and register relevant extension unmarshalers with the XConn.
func Register(xconn *xgb.XConn) error {
	// Query the X server for this extension
	reply, err := xproto.QueryExtension(xconn, uint16(len(ExtXName)), ExtXName)
	if err != nil {
		return fmt.Errorf("error querying X for \"Xevie\": %w", err)
	} else if !reply.Present {
		return fmt.Errorf("no extension named \"Xevie\" is known to the X server: reply=%+v", reply)
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
	DatatypeUnmodified = 0
	DatatypeModified   = 1
)

type Event struct {
	// padding: 32 bytes
}

// EventRead reads a byte slice into a Event value.
func EventRead(buf []byte, v *Event) int {
	b := 0

	b += 32 // padding

	return b
}

// EventReadList reads a byte slice into a list of Event values.
func EventReadList(buf []byte, dest []Event) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Event{}
		b += EventRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a Event value to a byte slice.
func (v Event) Bytes() []byte {
	buf := make([]byte, 32)
	b := 0

	b += 32 // padding

	return buf[:b]
}

// EventListBytes writes a list of Event values to a byte slice.
func EventListBytes(buf []byte, list []Event) int {
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

// End sends a checked request.
func End(c *xgb.XConn, Cmap uint32) (EndReply, error) {
	var reply EndReply
	op, ok := c.Ext("XEVIE")
	if !ok {
		return reply, errors.New("cannot issue request \"End\" using the uninitialized extension \"XEVIE\". xevie.Register(xconn) must be called first.")
	}
	err := c.SendRecv(endRequest(op, Cmap), &reply)
	return reply, err
}

// EndUnchecked sends an unchecked request.
func EndUnchecked(c *xgb.XConn, Cmap uint32) error {
	op, ok := c.Ext("XEVIE")
	if !ok {
		return errors.New("cannot issue request \"End\" using the uninitialized extension \"XEVIE\". xevie.Register(xconn) must be called first.")
	}
	return c.Send(endRequest(op, Cmap))
}

// EndReply represents the data returned from a End request.
type EndReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	// padding: 24 bytes
}

// Unmarshal reads a byte slice into a EndReply value.
func (v *EndReply) Unmarshal(buf []byte) error {
	if size := 32; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"EndReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	b += 24 // padding

	return nil
}

// Write request to wire for End
// endRequest writes a End request to a byte slice.
func endRequest(opcode uint8, Cmap uint32) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 2 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], Cmap)
	b += 4

	return buf
}

// QueryVersion sends a checked request.
func QueryVersion(c *xgb.XConn, ClientMajorVersion uint16, ClientMinorVersion uint16) (QueryVersionReply, error) {
	var reply QueryVersionReply
	op, ok := c.Ext("XEVIE")
	if !ok {
		return reply, errors.New("cannot issue request \"QueryVersion\" using the uninitialized extension \"XEVIE\". xevie.Register(xconn) must be called first.")
	}
	err := c.SendRecv(queryVersionRequest(op, ClientMajorVersion, ClientMinorVersion), &reply)
	return reply, err
}

// QueryVersionUnchecked sends an unchecked request.
func QueryVersionUnchecked(c *xgb.XConn, ClientMajorVersion uint16, ClientMinorVersion uint16) error {
	op, ok := c.Ext("XEVIE")
	if !ok {
		return errors.New("cannot issue request \"QueryVersion\" using the uninitialized extension \"XEVIE\". xevie.Register(xconn) must be called first.")
	}
	return c.Send(queryVersionRequest(op, ClientMajorVersion, ClientMinorVersion))
}

// QueryVersionReply represents the data returned from a QueryVersion request.
type QueryVersionReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	ServerMajorVersion uint16
	ServerMinorVersion uint16
	// padding: 20 bytes
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

	v.ServerMajorVersion = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.ServerMinorVersion = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	b += 20 // padding

	return nil
}

// Write request to wire for QueryVersion
// queryVersionRequest writes a QueryVersion request to a byte slice.
func queryVersionRequest(opcode uint8, ClientMajorVersion uint16, ClientMinorVersion uint16) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 0 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], ClientMajorVersion)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], ClientMinorVersion)
	b += 2

	return buf
}

// SelectInput sends a checked request.
func SelectInput(c *xgb.XConn, EventMask uint32) (SelectInputReply, error) {
	var reply SelectInputReply
	op, ok := c.Ext("XEVIE")
	if !ok {
		return reply, errors.New("cannot issue request \"SelectInput\" using the uninitialized extension \"XEVIE\". xevie.Register(xconn) must be called first.")
	}
	err := c.SendRecv(selectInputRequest(op, EventMask), &reply)
	return reply, err
}

// SelectInputUnchecked sends an unchecked request.
func SelectInputUnchecked(c *xgb.XConn, EventMask uint32) error {
	op, ok := c.Ext("XEVIE")
	if !ok {
		return errors.New("cannot issue request \"SelectInput\" using the uninitialized extension \"XEVIE\". xevie.Register(xconn) must be called first.")
	}
	return c.Send(selectInputRequest(op, EventMask))
}

// SelectInputReply represents the data returned from a SelectInput request.
type SelectInputReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	// padding: 24 bytes
}

// Unmarshal reads a byte slice into a SelectInputReply value.
func (v *SelectInputReply) Unmarshal(buf []byte) error {
	if size := 32; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"SelectInputReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	b += 24 // padding

	return nil
}

// Write request to wire for SelectInput
// selectInputRequest writes a SelectInput request to a byte slice.
func selectInputRequest(opcode uint8, EventMask uint32) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 4 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], EventMask)
	b += 4

	return buf
}

// Send sends a checked request.
func Send(c *xgb.XConn, Event Event, DataType uint32) (SendReply, error) {
	var reply SendReply
	op, ok := c.Ext("XEVIE")
	if !ok {
		return reply, errors.New("cannot issue request \"Send\" using the uninitialized extension \"XEVIE\". xevie.Register(xconn) must be called first.")
	}
	err := c.SendRecv(sendRequest(op, Event, DataType), &reply)
	return reply, err
}

// SendUnchecked sends an unchecked request.
func SendUnchecked(c *xgb.XConn, Event Event, DataType uint32) error {
	op, ok := c.Ext("XEVIE")
	if !ok {
		return errors.New("cannot issue request \"Send\" using the uninitialized extension \"XEVIE\". xevie.Register(xconn) must be called first.")
	}
	return c.Send(sendRequest(op, Event, DataType))
}

// SendReply represents the data returned from a Send request.
type SendReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	// padding: 24 bytes
}

// Unmarshal reads a byte slice into a SendReply value.
func (v *SendReply) Unmarshal(buf []byte) error {
	if size := 32; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"SendReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	b += 24 // padding

	return nil
}

// Write request to wire for Send
// sendRequest writes a Send request to a byte slice.
func sendRequest(opcode uint8, Event Event, DataType uint32) []byte {
	size := 104
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 3 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	{
		structBytes := Event.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	binary.LittleEndian.PutUint32(buf[b:], DataType)
	b += 4

	b += 64 // padding

	return buf
}

// Start sends a checked request.
func Start(c *xgb.XConn, Screen uint32) (StartReply, error) {
	var reply StartReply
	op, ok := c.Ext("XEVIE")
	if !ok {
		return reply, errors.New("cannot issue request \"Start\" using the uninitialized extension \"XEVIE\". xevie.Register(xconn) must be called first.")
	}
	err := c.SendRecv(startRequest(op, Screen), &reply)
	return reply, err
}

// StartUnchecked sends an unchecked request.
func StartUnchecked(c *xgb.XConn, Screen uint32) error {
	op, ok := c.Ext("XEVIE")
	if !ok {
		return errors.New("cannot issue request \"Start\" using the uninitialized extension \"XEVIE\". xevie.Register(xconn) must be called first.")
	}
	return c.Send(startRequest(op, Screen))
}

// StartReply represents the data returned from a Start request.
type StartReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	// padding: 24 bytes
}

// Unmarshal reads a byte slice into a StartReply value.
func (v *StartReply) Unmarshal(buf []byte) error {
	if size := 32; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"StartReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	b += 24 // padding

	return nil
}

// Write request to wire for Start
// startRequest writes a Start request to a byte slice.
func startRequest(opcode uint8, Screen uint32) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 1 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], Screen)
	b += 4

	return buf
}
