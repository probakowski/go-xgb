// FILE GENERATED AUTOMATICALLY FROM "xc_misc.xml"
package xcmisc

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
	ExtName = "XCMisc"

	// ExtXName is the name string this extension is known by to the X server.
	ExtXName = "XC-MISC"
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
		return fmt.Errorf("error querying X for \"XCMisc\": %w", err)
	} else if !reply.Present {
		return fmt.Errorf("no extension named \"XCMisc\" is known to the X server: reply=%+v", reply)
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

// GetVersion sends a checked request.
func GetVersion(c *xgb.XConn, ClientMajorVersion uint16, ClientMinorVersion uint16) (GetVersionReply, error) {
	var reply GetVersionReply
	op, ok := c.Ext("XC-MISC")
	if !ok {
		return reply, errors.New("cannot issue request \"GetVersion\" using the uninitialized extension \"XC-MISC\". xcmisc.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getVersionRequest(op, ClientMajorVersion, ClientMinorVersion), &reply)
	return reply, err
}

// GetVersionUnchecked sends an unchecked request.
func GetVersionUnchecked(c *xgb.XConn, ClientMajorVersion uint16, ClientMinorVersion uint16) error {
	op, ok := c.Ext("XC-MISC")
	if !ok {
		return errors.New("cannot issue request \"GetVersion\" using the uninitialized extension \"XC-MISC\". xcmisc.Register(xconn) must be called first.")
	}
	return c.Send(getVersionRequest(op, ClientMajorVersion, ClientMinorVersion))
}

// GetVersionReply represents the data returned from a GetVersion request.
type GetVersionReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	ServerMajorVersion uint16
	ServerMinorVersion uint16
}

// Unmarshal reads a byte slice into a GetVersionReply value.
func (v *GetVersionReply) Unmarshal(buf []byte) error {
	if size := 12; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetVersionReply\": have=%d need=%d", len(buf), size)
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

	return nil
}

// Write request to wire for GetVersion
// getVersionRequest writes a GetVersion request to a byte slice.
func getVersionRequest(opcode uint8, ClientMajorVersion uint16, ClientMinorVersion uint16) []byte {
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

// GetXIDList sends a checked request.
func GetXIDList(c *xgb.XConn, Count uint32) (GetXIDListReply, error) {
	var reply GetXIDListReply
	op, ok := c.Ext("XC-MISC")
	if !ok {
		return reply, errors.New("cannot issue request \"GetXIDList\" using the uninitialized extension \"XC-MISC\". xcmisc.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getXIDListRequest(op, Count), &reply)
	return reply, err
}

// GetXIDListUnchecked sends an unchecked request.
func GetXIDListUnchecked(c *xgb.XConn, Count uint32) error {
	op, ok := c.Ext("XC-MISC")
	if !ok {
		return errors.New("cannot issue request \"GetXIDList\" using the uninitialized extension \"XC-MISC\". xcmisc.Register(xconn) must be called first.")
	}
	return c.Send(getXIDListRequest(op, Count))
}

// GetXIDListReply represents the data returned from a GetXIDList request.
type GetXIDListReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	IdsLen uint32
	// padding: 20 bytes
	Ids []uint32 // size: internal.Pad4((int(IdsLen) * 4))
}

// Unmarshal reads a byte slice into a GetXIDListReply value.
func (v *GetXIDListReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.IdsLen) * 4))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetXIDListReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.IdsLen = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 20 // padding

	v.Ids = make([]uint32, v.IdsLen)
	for i := 0; i < int(v.IdsLen); i++ {
		v.Ids[i] = binary.LittleEndian.Uint32(buf[b:])
		b += 4
	}

	return nil
}

// Write request to wire for GetXIDList
// getXIDListRequest writes a GetXIDList request to a byte slice.
func getXIDListRequest(opcode uint8, Count uint32) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 2 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], Count)
	b += 4

	return buf
}

// GetXIDRange sends a checked request.
func GetXIDRange(c *xgb.XConn) (GetXIDRangeReply, error) {
	var reply GetXIDRangeReply
	op, ok := c.Ext("XC-MISC")
	if !ok {
		return reply, errors.New("cannot issue request \"GetXIDRange\" using the uninitialized extension \"XC-MISC\". xcmisc.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getXIDRangeRequest(op), &reply)
	return reply, err
}

// GetXIDRangeUnchecked sends an unchecked request.
func GetXIDRangeUnchecked(c *xgb.XConn) error {
	op, ok := c.Ext("XC-MISC")
	if !ok {
		return errors.New("cannot issue request \"GetXIDRange\" using the uninitialized extension \"XC-MISC\". xcmisc.Register(xconn) must be called first.")
	}
	return c.Send(getXIDRangeRequest(op))
}

// GetXIDRangeReply represents the data returned from a GetXIDRange request.
type GetXIDRangeReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	StartId uint32
	Count   uint32
}

// Unmarshal reads a byte slice into a GetXIDRangeReply value.
func (v *GetXIDRangeReply) Unmarshal(buf []byte) error {
	if size := 16; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetXIDRangeReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.StartId = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.Count = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	return nil
}

// Write request to wire for GetXIDRange
// getXIDRangeRequest writes a GetXIDRange request to a byte slice.
func getXIDRangeRequest(opcode uint8) []byte {
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
