// FILE GENERATED AUTOMATICALLY FROM "xinerama.xml"
package xinerama

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
	ExtName = "Xinerama"

	// ExtXName is the name string this extension is known by to the X server.
	ExtXName = "XINERAMA"
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
		return fmt.Errorf("error querying X for \"Xinerama\": %w", err)
	} else if !reply.Present {
		return fmt.Errorf("no extension named \"Xinerama\" is known to the X server: reply=%+v", reply)
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

type ScreenInfo struct {
	XOrg   int16
	YOrg   int16
	Width  uint16
	Height uint16
}

// ScreenInfoRead reads a byte slice into a ScreenInfo value.
func ScreenInfoRead(buf []byte, v *ScreenInfo) int {
	b := 0

	v.XOrg = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.YOrg = int16(binary.LittleEndian.Uint16(buf[b:]))
	b += 2

	v.Width = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Height = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return b
}

// ScreenInfoReadList reads a byte slice into a list of ScreenInfo values.
func ScreenInfoReadList(buf []byte, dest []ScreenInfo) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = ScreenInfo{}
		b += ScreenInfoRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a ScreenInfo value to a byte slice.
func (v ScreenInfo) Bytes() []byte {
	buf := make([]byte, 8)
	b := 0

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.XOrg))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], uint16(v.YOrg))
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Width)
	b += 2

	binary.LittleEndian.PutUint16(buf[b:], v.Height)
	b += 2

	return buf[:b]
}

// ScreenInfoListBytes writes a list of ScreenInfo values to a byte slice.
func ScreenInfoListBytes(buf []byte, list []ScreenInfo) int {
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

// GetScreenCount sends a checked request.
func GetScreenCount(c *xgb.XConn, Window xproto.Window) (GetScreenCountReply, error) {
	var reply GetScreenCountReply
	op, ok := c.Ext("XINERAMA")
	if !ok {
		return reply, errors.New("cannot issue request \"GetScreenCount\" using the uninitialized extension \"XINERAMA\". xinerama.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getScreenCountRequest(op, Window), &reply)
	return reply, err
}

// GetScreenCountUnchecked sends an unchecked request.
func GetScreenCountUnchecked(c *xgb.XConn, Window xproto.Window) error {
	op, ok := c.Ext("XINERAMA")
	if !ok {
		return errors.New("cannot issue request \"GetScreenCount\" using the uninitialized extension \"XINERAMA\". xinerama.Register(xconn) must be called first.")
	}
	return c.Send(getScreenCountRequest(op, Window))
}

// GetScreenCountReply represents the data returned from a GetScreenCount request.
type GetScreenCountReply struct {
	Sequence    uint16 // sequence number of the request for this reply
	Length      uint32 // number of bytes in this reply
	ScreenCount byte
	Window      xproto.Window
}

// Unmarshal reads a byte slice into a GetScreenCountReply value.
func (v *GetScreenCountReply) Unmarshal(buf []byte) error {
	if size := 12; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetScreenCountReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	v.ScreenCount = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Window = xproto.Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	return nil
}

// Write request to wire for GetScreenCount
// getScreenCountRequest writes a GetScreenCount request to a byte slice.
func getScreenCountRequest(opcode uint8, Window xproto.Window) []byte {
	size := 8
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

	return buf
}

// GetScreenSize sends a checked request.
func GetScreenSize(c *xgb.XConn, Window xproto.Window, Screen uint32) (GetScreenSizeReply, error) {
	var reply GetScreenSizeReply
	op, ok := c.Ext("XINERAMA")
	if !ok {
		return reply, errors.New("cannot issue request \"GetScreenSize\" using the uninitialized extension \"XINERAMA\". xinerama.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getScreenSizeRequest(op, Window, Screen), &reply)
	return reply, err
}

// GetScreenSizeUnchecked sends an unchecked request.
func GetScreenSizeUnchecked(c *xgb.XConn, Window xproto.Window, Screen uint32) error {
	op, ok := c.Ext("XINERAMA")
	if !ok {
		return errors.New("cannot issue request \"GetScreenSize\" using the uninitialized extension \"XINERAMA\". xinerama.Register(xconn) must be called first.")
	}
	return c.Send(getScreenSizeRequest(op, Window, Screen))
}

// GetScreenSizeReply represents the data returned from a GetScreenSize request.
type GetScreenSizeReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Width  uint32
	Height uint32
	Window xproto.Window
	Screen uint32
}

// Unmarshal reads a byte slice into a GetScreenSizeReply value.
func (v *GetScreenSizeReply) Unmarshal(buf []byte) error {
	if size := 24; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetScreenSizeReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Width = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.Height = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.Window = xproto.Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Screen = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	return nil
}

// Write request to wire for GetScreenSize
// getScreenSizeRequest writes a GetScreenSize request to a byte slice.
func getScreenSizeRequest(opcode uint8, Window xproto.Window, Screen uint32) []byte {
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

	binary.LittleEndian.PutUint32(buf[b:], Screen)
	b += 4

	return buf
}

// GetState sends a checked request.
func GetState(c *xgb.XConn, Window xproto.Window) (GetStateReply, error) {
	var reply GetStateReply
	op, ok := c.Ext("XINERAMA")
	if !ok {
		return reply, errors.New("cannot issue request \"GetState\" using the uninitialized extension \"XINERAMA\". xinerama.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getStateRequest(op, Window), &reply)
	return reply, err
}

// GetStateUnchecked sends an unchecked request.
func GetStateUnchecked(c *xgb.XConn, Window xproto.Window) error {
	op, ok := c.Ext("XINERAMA")
	if !ok {
		return errors.New("cannot issue request \"GetState\" using the uninitialized extension \"XINERAMA\". xinerama.Register(xconn) must be called first.")
	}
	return c.Send(getStateRequest(op, Window))
}

// GetStateReply represents the data returned from a GetState request.
type GetStateReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	State    byte
	Window   xproto.Window
}

// Unmarshal reads a byte slice into a GetStateReply value.
func (v *GetStateReply) Unmarshal(buf []byte) error {
	if size := 12; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetStateReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	v.State = buf[b]
	b += 1

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Window = xproto.Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	return nil
}

// Write request to wire for GetState
// getStateRequest writes a GetState request to a byte slice.
func getStateRequest(opcode uint8, Window xproto.Window) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 1 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	return buf
}

// IsActive sends a checked request.
func IsActive(c *xgb.XConn) (IsActiveReply, error) {
	var reply IsActiveReply
	op, ok := c.Ext("XINERAMA")
	if !ok {
		return reply, errors.New("cannot issue request \"IsActive\" using the uninitialized extension \"XINERAMA\". xinerama.Register(xconn) must be called first.")
	}
	err := c.SendRecv(isActiveRequest(op), &reply)
	return reply, err
}

// IsActiveUnchecked sends an unchecked request.
func IsActiveUnchecked(c *xgb.XConn) error {
	op, ok := c.Ext("XINERAMA")
	if !ok {
		return errors.New("cannot issue request \"IsActive\" using the uninitialized extension \"XINERAMA\". xinerama.Register(xconn) must be called first.")
	}
	return c.Send(isActiveRequest(op))
}

// IsActiveReply represents the data returned from a IsActive request.
type IsActiveReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	State uint32
}

// Unmarshal reads a byte slice into a IsActiveReply value.
func (v *IsActiveReply) Unmarshal(buf []byte) error {
	if size := 12; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"IsActiveReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.State = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	return nil
}

// Write request to wire for IsActive
// isActiveRequest writes a IsActive request to a byte slice.
func isActiveRequest(opcode uint8) []byte {
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

// QueryScreens sends a checked request.
func QueryScreens(c *xgb.XConn) (QueryScreensReply, error) {
	var reply QueryScreensReply
	op, ok := c.Ext("XINERAMA")
	if !ok {
		return reply, errors.New("cannot issue request \"QueryScreens\" using the uninitialized extension \"XINERAMA\". xinerama.Register(xconn) must be called first.")
	}
	err := c.SendRecv(queryScreensRequest(op), &reply)
	return reply, err
}

// QueryScreensUnchecked sends an unchecked request.
func QueryScreensUnchecked(c *xgb.XConn) error {
	op, ok := c.Ext("XINERAMA")
	if !ok {
		return errors.New("cannot issue request \"QueryScreens\" using the uninitialized extension \"XINERAMA\". xinerama.Register(xconn) must be called first.")
	}
	return c.Send(queryScreensRequest(op))
}

// QueryScreensReply represents the data returned from a QueryScreens request.
type QueryScreensReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Number uint32
	// padding: 20 bytes
	ScreenInfo []ScreenInfo // size: internal.Pad4((int(Number) * 8))
}

// Unmarshal reads a byte slice into a QueryScreensReply value.
func (v *QueryScreensReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.Number) * 8))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"QueryScreensReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Number = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 20 // padding

	v.ScreenInfo = make([]ScreenInfo, v.Number)
	b += ScreenInfoReadList(buf[b:], v.ScreenInfo)

	return nil
}

// Write request to wire for QueryScreens
// queryScreensRequest writes a QueryScreens request to a byte slice.
func queryScreensRequest(opcode uint8) []byte {
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

// QueryVersion sends a checked request.
func QueryVersion(c *xgb.XConn, Major byte, Minor byte) (QueryVersionReply, error) {
	var reply QueryVersionReply
	op, ok := c.Ext("XINERAMA")
	if !ok {
		return reply, errors.New("cannot issue request \"QueryVersion\" using the uninitialized extension \"XINERAMA\". xinerama.Register(xconn) must be called first.")
	}
	err := c.SendRecv(queryVersionRequest(op, Major, Minor), &reply)
	return reply, err
}

// QueryVersionUnchecked sends an unchecked request.
func QueryVersionUnchecked(c *xgb.XConn, Major byte, Minor byte) error {
	op, ok := c.Ext("XINERAMA")
	if !ok {
		return errors.New("cannot issue request \"QueryVersion\" using the uninitialized extension \"XINERAMA\". xinerama.Register(xconn) must be called first.")
	}
	return c.Send(queryVersionRequest(op, Major, Minor))
}

// QueryVersionReply represents the data returned from a QueryVersion request.
type QueryVersionReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Major uint16
	Minor uint16
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

	v.Major = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Minor = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return nil
}

// Write request to wire for QueryVersion
// queryVersionRequest writes a QueryVersion request to a byte slice.
func queryVersionRequest(opcode uint8, Major byte, Minor byte) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 0 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	buf[b] = Major
	b += 1

	buf[b] = Minor
	b += 1

	return buf
}
