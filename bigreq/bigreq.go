// FILE GENERATED AUTOMATICALLY FROM "bigreq.xml"
package bigreq

import (
	"encoding/binary"
	"errors"
	"fmt"

	"codeberg.org/gruf/go-xgb"
	"codeberg.org/gruf/go-xgb/xproto"
)

const (
	// ExtName is the user-friendly name string of this X extension.
	ExtName = "BigRequests"

	// ExtXName is the name string this extension is known by to the X server.
	ExtXName = "BIG-REQUESTS"
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
		return fmt.Errorf("error querying X for \"BigRequests\": %w", err)
	} else if !reply.Present {
		return fmt.Errorf("no extension named \"BigRequests\" is known to the X server: reply=%+v", reply)
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

// Enable sends a checked request.
func Enable(c *xgb.XConn) (EnableReply, error) {
	var reply EnableReply
	op, ok := c.Ext("BIG-REQUESTS")
	if !ok {
		return reply, errors.New("cannot issue request \"Enable\" using the uninitialized extension \"BIG-REQUESTS\". bigreq.Register(xconn) must be called first.")
	}
	err := c.SendRecv(enableRequest(op), &reply)
	return reply, err
}

// EnableUnchecked sends an unchecked request.
func EnableUnchecked(c *xgb.XConn) error {
	op, ok := c.Ext("BIG-REQUESTS")
	if !ok {
		return errors.New("cannot issue request \"Enable\" using the uninitialized extension \"BIG-REQUESTS\". bigreq.Register(xconn) must be called first.")
	}
	return c.Send(enableRequest(op))
}

// EnableReply represents the data returned from a Enable request.
type EnableReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	MaximumRequestLength uint32
}

// Unmarshal reads a byte slice into a EnableReply value.
func (v *EnableReply) Unmarshal(buf []byte) error {
	if size := 12; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"EnableReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.MaximumRequestLength = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	return nil
}

// Write request to wire for Enable
// enableRequest writes a Enable request to a byte slice.
func enableRequest(opcode uint8) []byte {
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
