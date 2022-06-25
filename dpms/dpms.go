// FILE GENERATED AUTOMATICALLY FROM "dpms.xml"
package dpms

import (
	"encoding/binary"
	"errors"
	"fmt"

	"codeberg.org/gruf/go-xgb"
	"codeberg.org/gruf/go-xgb/xproto"
)

const (
	// ExtName is the user-friendly name string of this X extension.
	ExtName = "DPMS"

	// ExtXName is the name string this extension is known by to the X server.
	ExtXName = "DPMS"
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
		return fmt.Errorf("error querying X for \"DPMS\": %w", err)
	} else if !reply.Present {
		return fmt.Errorf("no extension named \"DPMS\" is known to the X server: reply=%+v", reply)
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

const (
	DPMSModeOn      = 0
	DPMSModeStandby = 1
	DPMSModeSuspend = 2
	DPMSModeOff     = 3
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

// Capable sends a checked request.
func Capable(c *xgb.XConn) (CapableReply, error) {
	var reply CapableReply
	op, ok := c.Ext("DPMS")
	if !ok {
		return reply, errors.New("cannot issue request \"Capable\" using the uninitialized extension \"DPMS\". dpms.Register(xconn) must be called first.")
	}
	err := c.SendRecv(capableRequest(op), &reply)
	return reply, err
}

// CapableUnchecked sends an unchecked request.
func CapableUnchecked(c *xgb.XConn) error {
	op, ok := c.Ext("DPMS")
	if !ok {
		return errors.New("cannot issue request \"Capable\" using the uninitialized extension \"DPMS\". dpms.Register(xconn) must be called first.")
	}
	return c.Send(capableRequest(op))
}

// CapableReply represents the data returned from a Capable request.
type CapableReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Capable bool
	// padding: 23 bytes
}

// Unmarshal reads a byte slice into a CapableReply value.
func (v *CapableReply) Unmarshal(buf []byte) error {
	if size := 32; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"CapableReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Capable = (buf[b] == 1)
	b += 1

	b += 23 // padding

	return nil
}

// Write request to wire for Capable
// capableRequest writes a Capable request to a byte slice.
func capableRequest(opcode uint8) []byte {
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

// Disable sends a checked request.
func Disable(c *xgb.XConn) error {
	op, ok := c.Ext("DPMS")
	if !ok {
		return errors.New("cannot issue request \"Disable\" using the uninitialized extension \"DPMS\". dpms.Register(xconn) must be called first.")
	}
	return c.SendRecv(disableRequest(op), nil)
}

// DisableUnchecked sends an unchecked request.
func DisableUnchecked(c *xgb.XConn) error {
	op, ok := c.Ext("DPMS")
	if !ok {
		return errors.New("cannot issue request \"Disable\" using the uninitialized extension \"DPMS\". dpms.Register(xconn) must be called first.")
	}
	return c.Send(disableRequest(op))
}

// Write request to wire for Disable
// disableRequest writes a Disable request to a byte slice.
func disableRequest(opcode uint8) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 5 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// Enable sends a checked request.
func Enable(c *xgb.XConn) error {
	op, ok := c.Ext("DPMS")
	if !ok {
		return errors.New("cannot issue request \"Enable\" using the uninitialized extension \"DPMS\". dpms.Register(xconn) must be called first.")
	}
	return c.SendRecv(enableRequest(op), nil)
}

// EnableUnchecked sends an unchecked request.
func EnableUnchecked(c *xgb.XConn) error {
	op, ok := c.Ext("DPMS")
	if !ok {
		return errors.New("cannot issue request \"Enable\" using the uninitialized extension \"DPMS\". dpms.Register(xconn) must be called first.")
	}
	return c.Send(enableRequest(op))
}

// Write request to wire for Enable
// enableRequest writes a Enable request to a byte slice.
func enableRequest(opcode uint8) []byte {
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

// ForceLevel sends a checked request.
func ForceLevel(c *xgb.XConn, PowerLevel uint16) error {
	op, ok := c.Ext("DPMS")
	if !ok {
		return errors.New("cannot issue request \"ForceLevel\" using the uninitialized extension \"DPMS\". dpms.Register(xconn) must be called first.")
	}
	return c.SendRecv(forceLevelRequest(op, PowerLevel), nil)
}

// ForceLevelUnchecked sends an unchecked request.
func ForceLevelUnchecked(c *xgb.XConn, PowerLevel uint16) error {
	op, ok := c.Ext("DPMS")
	if !ok {
		return errors.New("cannot issue request \"ForceLevel\" using the uninitialized extension \"DPMS\". dpms.Register(xconn) must be called first.")
	}
	return c.Send(forceLevelRequest(op, PowerLevel))
}

// Write request to wire for ForceLevel
// forceLevelRequest writes a ForceLevel request to a byte slice.
func forceLevelRequest(opcode uint8, PowerLevel uint16) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 6 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], PowerLevel)
	b += 2

	return buf
}

// GetTimeouts sends a checked request.
func GetTimeouts(c *xgb.XConn) (GetTimeoutsReply, error) {
	var reply GetTimeoutsReply
	op, ok := c.Ext("DPMS")
	if !ok {
		return reply, errors.New("cannot issue request \"GetTimeouts\" using the uninitialized extension \"DPMS\". dpms.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getTimeoutsRequest(op), &reply)
	return reply, err
}

// GetTimeoutsUnchecked sends an unchecked request.
func GetTimeoutsUnchecked(c *xgb.XConn) error {
	op, ok := c.Ext("DPMS")
	if !ok {
		return errors.New("cannot issue request \"GetTimeouts\" using the uninitialized extension \"DPMS\". dpms.Register(xconn) must be called first.")
	}
	return c.Send(getTimeoutsRequest(op))
}

// GetTimeoutsReply represents the data returned from a GetTimeouts request.
type GetTimeoutsReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	StandbyTimeout uint16
	SuspendTimeout uint16
	OffTimeout     uint16
	// padding: 18 bytes
}

// Unmarshal reads a byte slice into a GetTimeoutsReply value.
func (v *GetTimeoutsReply) Unmarshal(buf []byte) error {
	if size := 32; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetTimeoutsReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.StandbyTimeout = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.SuspendTimeout = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.OffTimeout = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	b += 18 // padding

	return nil
}

// Write request to wire for GetTimeouts
// getTimeoutsRequest writes a GetTimeouts request to a byte slice.
func getTimeoutsRequest(opcode uint8) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 2 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// GetVersion sends a checked request.
func GetVersion(c *xgb.XConn, ClientMajorVersion uint16, ClientMinorVersion uint16) (GetVersionReply, error) {
	var reply GetVersionReply
	op, ok := c.Ext("DPMS")
	if !ok {
		return reply, errors.New("cannot issue request \"GetVersion\" using the uninitialized extension \"DPMS\". dpms.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getVersionRequest(op, ClientMajorVersion, ClientMinorVersion), &reply)
	return reply, err
}

// GetVersionUnchecked sends an unchecked request.
func GetVersionUnchecked(c *xgb.XConn, ClientMajorVersion uint16, ClientMinorVersion uint16) error {
	op, ok := c.Ext("DPMS")
	if !ok {
		return errors.New("cannot issue request \"GetVersion\" using the uninitialized extension \"DPMS\". dpms.Register(xconn) must be called first.")
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

// Info sends a checked request.
func Info(c *xgb.XConn) (InfoReply, error) {
	var reply InfoReply
	op, ok := c.Ext("DPMS")
	if !ok {
		return reply, errors.New("cannot issue request \"Info\" using the uninitialized extension \"DPMS\". dpms.Register(xconn) must be called first.")
	}
	err := c.SendRecv(infoRequest(op), &reply)
	return reply, err
}

// InfoUnchecked sends an unchecked request.
func InfoUnchecked(c *xgb.XConn) error {
	op, ok := c.Ext("DPMS")
	if !ok {
		return errors.New("cannot issue request \"Info\" using the uninitialized extension \"DPMS\". dpms.Register(xconn) must be called first.")
	}
	return c.Send(infoRequest(op))
}

// InfoReply represents the data returned from a Info request.
type InfoReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	PowerLevel uint16
	State      bool
	// padding: 21 bytes
}

// Unmarshal reads a byte slice into a InfoReply value.
func (v *InfoReply) Unmarshal(buf []byte) error {
	if size := 32; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"InfoReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.PowerLevel = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.State = (buf[b] == 1)
	b += 1

	b += 21 // padding

	return nil
}

// Write request to wire for Info
// infoRequest writes a Info request to a byte slice.
func infoRequest(opcode uint8) []byte {
	size := 4
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 7 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// SetTimeouts sends a checked request.
func SetTimeouts(c *xgb.XConn, StandbyTimeout uint16, SuspendTimeout uint16, OffTimeout uint16) error {
	op, ok := c.Ext("DPMS")
	if !ok {
		return errors.New("cannot issue request \"SetTimeouts\" using the uninitialized extension \"DPMS\". dpms.Register(xconn) must be called first.")
	}
	return c.SendRecv(setTimeoutsRequest(op, StandbyTimeout, SuspendTimeout, OffTimeout), nil)
}

// SetTimeoutsUnchecked sends an unchecked request.
func SetTimeoutsUnchecked(c *xgb.XConn, StandbyTimeout uint16, SuspendTimeout uint16, OffTimeout uint16) error {
	op, ok := c.Ext("DPMS")
	if !ok {
		return errors.New("cannot issue request \"SetTimeouts\" using the uninitialized extension \"DPMS\". dpms.Register(xconn) must be called first.")
	}
	return c.Send(setTimeoutsRequest(op, StandbyTimeout, SuspendTimeout, OffTimeout))
}

// Write request to wire for SetTimeouts
// setTimeoutsRequest writes a SetTimeouts request to a byte slice.
func setTimeoutsRequest(opcode uint8, StandbyTimeout uint16, SuspendTimeout uint16, OffTimeout uint16) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 3 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], StandbyTimeout)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], SuspendTimeout)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], OffTimeout)
	b += 2

	return buf
}
