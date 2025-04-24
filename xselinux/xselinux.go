// FILE GENERATED AUTOMATICALLY FROM "xselinux.xml"
package xselinux

import (
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/probakowski/go-xgb"
	"github.com/probakowski/go-xgb/internal"
	"github.com/probakowski/go-xgb/xproto"
)

const (
	// ExtName is the user-friendly name string of this X extension.
	ExtName = "SELinux"

	// ExtXName is the name string this extension is known by to the X server.
	ExtXName = "SELinux"
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

// Register will query the X server for SELinux extension support, and register relevant extension unmarshalers with the XConn.
func Register(xconn *xgb.XConn) error {
	// Query the X server for this extension
	reply, err := xproto.QueryExtension(xconn, uint16(len(ExtXName)), ExtXName)
	if err != nil {
		return fmt.Errorf("error querying X for \"SELinux\": %w", err)
	} else if !reply.Present {
		return fmt.Errorf("no extension named \"SELinux\" is known to the X server: reply=%+v", reply)
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

type ListItem struct {
	Name             xproto.Atom
	ObjectContextLen uint32
	DataContextLen   uint32
	ObjectContext    string // size: internal.Pad4((int(ObjectContextLen) * 1))
	// padding: 0 bytes
	DataContext string // size: internal.Pad4((int(DataContextLen) * 1))
	// padding: 0 bytes
}

// ListItemRead reads a byte slice into a ListItem value.
func ListItemRead(buf []byte, v *ListItem) int {
	b := 0

	v.Name = xproto.Atom(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.ObjectContextLen = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.DataContextLen = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	{
		byteString := make([]byte, v.ObjectContextLen)
		copy(byteString[:v.ObjectContextLen], buf[b:])
		v.ObjectContext = string(byteString)
		b += int(v.ObjectContextLen)
	}

	b += 0 // padding

	{
		byteString := make([]byte, v.DataContextLen)
		copy(byteString[:v.DataContextLen], buf[b:])
		v.DataContext = string(byteString)
		b += int(v.DataContextLen)
	}

	b += 0 // padding

	return b
}

// ListItemReadList reads a byte slice into a list of ListItem values.
func ListItemReadList(buf []byte, dest []ListItem) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = ListItem{}
		b += ListItemRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a ListItem value to a byte slice.
func (v ListItem) Bytes() []byte {
	buf := make([]byte, ((((12 + internal.Pad4((int(v.ObjectContextLen) * 1))) + 0) + internal.Pad4((int(v.DataContextLen) * 1))) + 0))
	b := 0

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.Name))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], v.ObjectContextLen)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], v.DataContextLen)
	b += 4

	copy(buf[b:], v.ObjectContext[:v.ObjectContextLen])
	b += int(v.ObjectContextLen)

	b += 0 // padding

	copy(buf[b:], v.DataContext[:v.DataContextLen])
	b += int(v.DataContextLen)

	b += 0 // padding

	return buf[:b]
}

// ListItemListBytes writes a list of ListItem values to a byte slice.
func ListItemListBytes(buf []byte, list []ListItem) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

// ListItemListSize computes the size (bytes) of a list of ListItem values.
func ListItemListSize(list []ListItem) int {
	size := 0
	for _, item := range list {
		size += ((((12 + internal.Pad4((int(item.ObjectContextLen) * 1))) + 0) + internal.Pad4((int(item.DataContextLen) * 1))) + 0)
	}
	return size
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

// GetClientContext sends a checked request.
func GetClientContext(c *xgb.XConn, Resource uint32) (GetClientContextReply, error) {
	var reply GetClientContextReply
	op, ok := c.Ext("SELinux")
	if !ok {
		return reply, errors.New("cannot issue request \"GetClientContext\" using the uninitialized extension \"SELinux\". xselinux.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getClientContextRequest(op, Resource), &reply)
	return reply, err
}

// GetClientContextUnchecked sends an unchecked request.
func GetClientContextUnchecked(c *xgb.XConn, Resource uint32) error {
	op, ok := c.Ext("SELinux")
	if !ok {
		return errors.New("cannot issue request \"GetClientContext\" using the uninitialized extension \"SELinux\". xselinux.Register(xconn) must be called first.")
	}
	return c.Send(getClientContextRequest(op, Resource))
}

// GetClientContextReply represents the data returned from a GetClientContext request.
type GetClientContextReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	ContextLen uint32
	// padding: 20 bytes
	Context string // size: internal.Pad4((int(ContextLen) * 1))
}

// Unmarshal reads a byte slice into a GetClientContextReply value.
func (v *GetClientContextReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.ContextLen) * 1))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetClientContextReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.ContextLen = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 20 // padding

	{
		byteString := make([]byte, v.ContextLen)
		copy(byteString[:v.ContextLen], buf[b:])
		v.Context = string(byteString)
		b += int(v.ContextLen)
	}

	return nil
}

// Write request to wire for GetClientContext
// getClientContextRequest writes a GetClientContext request to a byte slice.
func getClientContextRequest(opcode uint8, Resource uint32) []byte {
	const size = 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 22 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], Resource)
	b += 4

	return buf
}

// GetDeviceContext sends a checked request.
func GetDeviceContext(c *xgb.XConn, Device uint32) (GetDeviceContextReply, error) {
	var reply GetDeviceContextReply
	op, ok := c.Ext("SELinux")
	if !ok {
		return reply, errors.New("cannot issue request \"GetDeviceContext\" using the uninitialized extension \"SELinux\". xselinux.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getDeviceContextRequest(op, Device), &reply)
	return reply, err
}

// GetDeviceContextUnchecked sends an unchecked request.
func GetDeviceContextUnchecked(c *xgb.XConn, Device uint32) error {
	op, ok := c.Ext("SELinux")
	if !ok {
		return errors.New("cannot issue request \"GetDeviceContext\" using the uninitialized extension \"SELinux\". xselinux.Register(xconn) must be called first.")
	}
	return c.Send(getDeviceContextRequest(op, Device))
}

// GetDeviceContextReply represents the data returned from a GetDeviceContext request.
type GetDeviceContextReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	ContextLen uint32
	// padding: 20 bytes
	Context string // size: internal.Pad4((int(ContextLen) * 1))
}

// Unmarshal reads a byte slice into a GetDeviceContextReply value.
func (v *GetDeviceContextReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.ContextLen) * 1))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetDeviceContextReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.ContextLen = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 20 // padding

	{
		byteString := make([]byte, v.ContextLen)
		copy(byteString[:v.ContextLen], buf[b:])
		v.Context = string(byteString)
		b += int(v.ContextLen)
	}

	return nil
}

// Write request to wire for GetDeviceContext
// getDeviceContextRequest writes a GetDeviceContext request to a byte slice.
func getDeviceContextRequest(opcode uint8, Device uint32) []byte {
	const size = 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 4 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], Device)
	b += 4

	return buf
}

// GetDeviceCreateContext sends a checked request.
func GetDeviceCreateContext(c *xgb.XConn) (GetDeviceCreateContextReply, error) {
	var reply GetDeviceCreateContextReply
	op, ok := c.Ext("SELinux")
	if !ok {
		return reply, errors.New("cannot issue request \"GetDeviceCreateContext\" using the uninitialized extension \"SELinux\". xselinux.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getDeviceCreateContextRequest(op), &reply)
	return reply, err
}

// GetDeviceCreateContextUnchecked sends an unchecked request.
func GetDeviceCreateContextUnchecked(c *xgb.XConn) error {
	op, ok := c.Ext("SELinux")
	if !ok {
		return errors.New("cannot issue request \"GetDeviceCreateContext\" using the uninitialized extension \"SELinux\". xselinux.Register(xconn) must be called first.")
	}
	return c.Send(getDeviceCreateContextRequest(op))
}

// GetDeviceCreateContextReply represents the data returned from a GetDeviceCreateContext request.
type GetDeviceCreateContextReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	ContextLen uint32
	// padding: 20 bytes
	Context string // size: internal.Pad4((int(ContextLen) * 1))
}

// Unmarshal reads a byte slice into a GetDeviceCreateContextReply value.
func (v *GetDeviceCreateContextReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.ContextLen) * 1))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetDeviceCreateContextReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.ContextLen = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 20 // padding

	{
		byteString := make([]byte, v.ContextLen)
		copy(byteString[:v.ContextLen], buf[b:])
		v.Context = string(byteString)
		b += int(v.ContextLen)
	}

	return nil
}

// Write request to wire for GetDeviceCreateContext
// getDeviceCreateContextRequest writes a GetDeviceCreateContext request to a byte slice.
func getDeviceCreateContextRequest(opcode uint8) []byte {
	const size = 4
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

// GetPropertyContext sends a checked request.
func GetPropertyContext(c *xgb.XConn, Window xproto.Window, Property xproto.Atom) (GetPropertyContextReply, error) {
	var reply GetPropertyContextReply
	op, ok := c.Ext("SELinux")
	if !ok {
		return reply, errors.New("cannot issue request \"GetPropertyContext\" using the uninitialized extension \"SELinux\". xselinux.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getPropertyContextRequest(op, Window, Property), &reply)
	return reply, err
}

// GetPropertyContextUnchecked sends an unchecked request.
func GetPropertyContextUnchecked(c *xgb.XConn, Window xproto.Window, Property xproto.Atom) error {
	op, ok := c.Ext("SELinux")
	if !ok {
		return errors.New("cannot issue request \"GetPropertyContext\" using the uninitialized extension \"SELinux\". xselinux.Register(xconn) must be called first.")
	}
	return c.Send(getPropertyContextRequest(op, Window, Property))
}

// GetPropertyContextReply represents the data returned from a GetPropertyContext request.
type GetPropertyContextReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	ContextLen uint32
	// padding: 20 bytes
	Context string // size: internal.Pad4((int(ContextLen) * 1))
}

// Unmarshal reads a byte slice into a GetPropertyContextReply value.
func (v *GetPropertyContextReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.ContextLen) * 1))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetPropertyContextReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.ContextLen = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 20 // padding

	{
		byteString := make([]byte, v.ContextLen)
		copy(byteString[:v.ContextLen], buf[b:])
		v.Context = string(byteString)
		b += int(v.ContextLen)
	}

	return nil
}

// Write request to wire for GetPropertyContext
// getPropertyContextRequest writes a GetPropertyContext request to a byte slice.
func getPropertyContextRequest(opcode uint8, Window xproto.Window, Property xproto.Atom) []byte {
	const size = 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 12 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Property))
	b += 4

	return buf
}

// GetPropertyCreateContext sends a checked request.
func GetPropertyCreateContext(c *xgb.XConn) (GetPropertyCreateContextReply, error) {
	var reply GetPropertyCreateContextReply
	op, ok := c.Ext("SELinux")
	if !ok {
		return reply, errors.New("cannot issue request \"GetPropertyCreateContext\" using the uninitialized extension \"SELinux\". xselinux.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getPropertyCreateContextRequest(op), &reply)
	return reply, err
}

// GetPropertyCreateContextUnchecked sends an unchecked request.
func GetPropertyCreateContextUnchecked(c *xgb.XConn) error {
	op, ok := c.Ext("SELinux")
	if !ok {
		return errors.New("cannot issue request \"GetPropertyCreateContext\" using the uninitialized extension \"SELinux\". xselinux.Register(xconn) must be called first.")
	}
	return c.Send(getPropertyCreateContextRequest(op))
}

// GetPropertyCreateContextReply represents the data returned from a GetPropertyCreateContext request.
type GetPropertyCreateContextReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	ContextLen uint32
	// padding: 20 bytes
	Context string // size: internal.Pad4((int(ContextLen) * 1))
}

// Unmarshal reads a byte slice into a GetPropertyCreateContextReply value.
func (v *GetPropertyCreateContextReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.ContextLen) * 1))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetPropertyCreateContextReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.ContextLen = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 20 // padding

	{
		byteString := make([]byte, v.ContextLen)
		copy(byteString[:v.ContextLen], buf[b:])
		v.Context = string(byteString)
		b += int(v.ContextLen)
	}

	return nil
}

// Write request to wire for GetPropertyCreateContext
// getPropertyCreateContextRequest writes a GetPropertyCreateContext request to a byte slice.
func getPropertyCreateContextRequest(opcode uint8) []byte {
	const size = 4
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 9 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// GetPropertyDataContext sends a checked request.
func GetPropertyDataContext(c *xgb.XConn, Window xproto.Window, Property xproto.Atom) (GetPropertyDataContextReply, error) {
	var reply GetPropertyDataContextReply
	op, ok := c.Ext("SELinux")
	if !ok {
		return reply, errors.New("cannot issue request \"GetPropertyDataContext\" using the uninitialized extension \"SELinux\". xselinux.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getPropertyDataContextRequest(op, Window, Property), &reply)
	return reply, err
}

// GetPropertyDataContextUnchecked sends an unchecked request.
func GetPropertyDataContextUnchecked(c *xgb.XConn, Window xproto.Window, Property xproto.Atom) error {
	op, ok := c.Ext("SELinux")
	if !ok {
		return errors.New("cannot issue request \"GetPropertyDataContext\" using the uninitialized extension \"SELinux\". xselinux.Register(xconn) must be called first.")
	}
	return c.Send(getPropertyDataContextRequest(op, Window, Property))
}

// GetPropertyDataContextReply represents the data returned from a GetPropertyDataContext request.
type GetPropertyDataContextReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	ContextLen uint32
	// padding: 20 bytes
	Context string // size: internal.Pad4((int(ContextLen) * 1))
}

// Unmarshal reads a byte slice into a GetPropertyDataContextReply value.
func (v *GetPropertyDataContextReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.ContextLen) * 1))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetPropertyDataContextReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.ContextLen = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 20 // padding

	{
		byteString := make([]byte, v.ContextLen)
		copy(byteString[:v.ContextLen], buf[b:])
		v.Context = string(byteString)
		b += int(v.ContextLen)
	}

	return nil
}

// Write request to wire for GetPropertyDataContext
// getPropertyDataContextRequest writes a GetPropertyDataContext request to a byte slice.
func getPropertyDataContextRequest(opcode uint8, Window xproto.Window, Property xproto.Atom) []byte {
	const size = 12
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

	binary.LittleEndian.PutUint32(buf[b:], uint32(Property))
	b += 4

	return buf
}

// GetPropertyUseContext sends a checked request.
func GetPropertyUseContext(c *xgb.XConn) (GetPropertyUseContextReply, error) {
	var reply GetPropertyUseContextReply
	op, ok := c.Ext("SELinux")
	if !ok {
		return reply, errors.New("cannot issue request \"GetPropertyUseContext\" using the uninitialized extension \"SELinux\". xselinux.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getPropertyUseContextRequest(op), &reply)
	return reply, err
}

// GetPropertyUseContextUnchecked sends an unchecked request.
func GetPropertyUseContextUnchecked(c *xgb.XConn) error {
	op, ok := c.Ext("SELinux")
	if !ok {
		return errors.New("cannot issue request \"GetPropertyUseContext\" using the uninitialized extension \"SELinux\". xselinux.Register(xconn) must be called first.")
	}
	return c.Send(getPropertyUseContextRequest(op))
}

// GetPropertyUseContextReply represents the data returned from a GetPropertyUseContext request.
type GetPropertyUseContextReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	ContextLen uint32
	// padding: 20 bytes
	Context string // size: internal.Pad4((int(ContextLen) * 1))
}

// Unmarshal reads a byte slice into a GetPropertyUseContextReply value.
func (v *GetPropertyUseContextReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.ContextLen) * 1))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetPropertyUseContextReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.ContextLen = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 20 // padding

	{
		byteString := make([]byte, v.ContextLen)
		copy(byteString[:v.ContextLen], buf[b:])
		v.Context = string(byteString)
		b += int(v.ContextLen)
	}

	return nil
}

// Write request to wire for GetPropertyUseContext
// getPropertyUseContextRequest writes a GetPropertyUseContext request to a byte slice.
func getPropertyUseContextRequest(opcode uint8) []byte {
	const size = 4
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 11 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// GetSelectionContext sends a checked request.
func GetSelectionContext(c *xgb.XConn, Selection xproto.Atom) (GetSelectionContextReply, error) {
	var reply GetSelectionContextReply
	op, ok := c.Ext("SELinux")
	if !ok {
		return reply, errors.New("cannot issue request \"GetSelectionContext\" using the uninitialized extension \"SELinux\". xselinux.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getSelectionContextRequest(op, Selection), &reply)
	return reply, err
}

// GetSelectionContextUnchecked sends an unchecked request.
func GetSelectionContextUnchecked(c *xgb.XConn, Selection xproto.Atom) error {
	op, ok := c.Ext("SELinux")
	if !ok {
		return errors.New("cannot issue request \"GetSelectionContext\" using the uninitialized extension \"SELinux\". xselinux.Register(xconn) must be called first.")
	}
	return c.Send(getSelectionContextRequest(op, Selection))
}

// GetSelectionContextReply represents the data returned from a GetSelectionContext request.
type GetSelectionContextReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	ContextLen uint32
	// padding: 20 bytes
	Context string // size: internal.Pad4((int(ContextLen) * 1))
}

// Unmarshal reads a byte slice into a GetSelectionContextReply value.
func (v *GetSelectionContextReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.ContextLen) * 1))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetSelectionContextReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.ContextLen = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 20 // padding

	{
		byteString := make([]byte, v.ContextLen)
		copy(byteString[:v.ContextLen], buf[b:])
		v.Context = string(byteString)
		b += int(v.ContextLen)
	}

	return nil
}

// Write request to wire for GetSelectionContext
// getSelectionContextRequest writes a GetSelectionContext request to a byte slice.
func getSelectionContextRequest(opcode uint8, Selection xproto.Atom) []byte {
	const size = 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 19 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Selection))
	b += 4

	return buf
}

// GetSelectionCreateContext sends a checked request.
func GetSelectionCreateContext(c *xgb.XConn) (GetSelectionCreateContextReply, error) {
	var reply GetSelectionCreateContextReply
	op, ok := c.Ext("SELinux")
	if !ok {
		return reply, errors.New("cannot issue request \"GetSelectionCreateContext\" using the uninitialized extension \"SELinux\". xselinux.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getSelectionCreateContextRequest(op), &reply)
	return reply, err
}

// GetSelectionCreateContextUnchecked sends an unchecked request.
func GetSelectionCreateContextUnchecked(c *xgb.XConn) error {
	op, ok := c.Ext("SELinux")
	if !ok {
		return errors.New("cannot issue request \"GetSelectionCreateContext\" using the uninitialized extension \"SELinux\". xselinux.Register(xconn) must be called first.")
	}
	return c.Send(getSelectionCreateContextRequest(op))
}

// GetSelectionCreateContextReply represents the data returned from a GetSelectionCreateContext request.
type GetSelectionCreateContextReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	ContextLen uint32
	// padding: 20 bytes
	Context string // size: internal.Pad4((int(ContextLen) * 1))
}

// Unmarshal reads a byte slice into a GetSelectionCreateContextReply value.
func (v *GetSelectionCreateContextReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.ContextLen) * 1))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetSelectionCreateContextReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.ContextLen = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 20 // padding

	{
		byteString := make([]byte, v.ContextLen)
		copy(byteString[:v.ContextLen], buf[b:])
		v.Context = string(byteString)
		b += int(v.ContextLen)
	}

	return nil
}

// Write request to wire for GetSelectionCreateContext
// getSelectionCreateContextRequest writes a GetSelectionCreateContext request to a byte slice.
func getSelectionCreateContextRequest(opcode uint8) []byte {
	const size = 4
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 16 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// GetSelectionDataContext sends a checked request.
func GetSelectionDataContext(c *xgb.XConn, Selection xproto.Atom) (GetSelectionDataContextReply, error) {
	var reply GetSelectionDataContextReply
	op, ok := c.Ext("SELinux")
	if !ok {
		return reply, errors.New("cannot issue request \"GetSelectionDataContext\" using the uninitialized extension \"SELinux\". xselinux.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getSelectionDataContextRequest(op, Selection), &reply)
	return reply, err
}

// GetSelectionDataContextUnchecked sends an unchecked request.
func GetSelectionDataContextUnchecked(c *xgb.XConn, Selection xproto.Atom) error {
	op, ok := c.Ext("SELinux")
	if !ok {
		return errors.New("cannot issue request \"GetSelectionDataContext\" using the uninitialized extension \"SELinux\". xselinux.Register(xconn) must be called first.")
	}
	return c.Send(getSelectionDataContextRequest(op, Selection))
}

// GetSelectionDataContextReply represents the data returned from a GetSelectionDataContext request.
type GetSelectionDataContextReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	ContextLen uint32
	// padding: 20 bytes
	Context string // size: internal.Pad4((int(ContextLen) * 1))
}

// Unmarshal reads a byte slice into a GetSelectionDataContextReply value.
func (v *GetSelectionDataContextReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.ContextLen) * 1))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetSelectionDataContextReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.ContextLen = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 20 // padding

	{
		byteString := make([]byte, v.ContextLen)
		copy(byteString[:v.ContextLen], buf[b:])
		v.Context = string(byteString)
		b += int(v.ContextLen)
	}

	return nil
}

// Write request to wire for GetSelectionDataContext
// getSelectionDataContextRequest writes a GetSelectionDataContext request to a byte slice.
func getSelectionDataContextRequest(opcode uint8, Selection xproto.Atom) []byte {
	const size = 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 20 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Selection))
	b += 4

	return buf
}

// GetSelectionUseContext sends a checked request.
func GetSelectionUseContext(c *xgb.XConn) (GetSelectionUseContextReply, error) {
	var reply GetSelectionUseContextReply
	op, ok := c.Ext("SELinux")
	if !ok {
		return reply, errors.New("cannot issue request \"GetSelectionUseContext\" using the uninitialized extension \"SELinux\". xselinux.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getSelectionUseContextRequest(op), &reply)
	return reply, err
}

// GetSelectionUseContextUnchecked sends an unchecked request.
func GetSelectionUseContextUnchecked(c *xgb.XConn) error {
	op, ok := c.Ext("SELinux")
	if !ok {
		return errors.New("cannot issue request \"GetSelectionUseContext\" using the uninitialized extension \"SELinux\". xselinux.Register(xconn) must be called first.")
	}
	return c.Send(getSelectionUseContextRequest(op))
}

// GetSelectionUseContextReply represents the data returned from a GetSelectionUseContext request.
type GetSelectionUseContextReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	ContextLen uint32
	// padding: 20 bytes
	Context string // size: internal.Pad4((int(ContextLen) * 1))
}

// Unmarshal reads a byte slice into a GetSelectionUseContextReply value.
func (v *GetSelectionUseContextReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.ContextLen) * 1))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetSelectionUseContextReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.ContextLen = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 20 // padding

	{
		byteString := make([]byte, v.ContextLen)
		copy(byteString[:v.ContextLen], buf[b:])
		v.Context = string(byteString)
		b += int(v.ContextLen)
	}

	return nil
}

// Write request to wire for GetSelectionUseContext
// getSelectionUseContextRequest writes a GetSelectionUseContext request to a byte slice.
func getSelectionUseContextRequest(opcode uint8) []byte {
	const size = 4
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 18 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// GetWindowContext sends a checked request.
func GetWindowContext(c *xgb.XConn, Window xproto.Window) (GetWindowContextReply, error) {
	var reply GetWindowContextReply
	op, ok := c.Ext("SELinux")
	if !ok {
		return reply, errors.New("cannot issue request \"GetWindowContext\" using the uninitialized extension \"SELinux\". xselinux.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getWindowContextRequest(op, Window), &reply)
	return reply, err
}

// GetWindowContextUnchecked sends an unchecked request.
func GetWindowContextUnchecked(c *xgb.XConn, Window xproto.Window) error {
	op, ok := c.Ext("SELinux")
	if !ok {
		return errors.New("cannot issue request \"GetWindowContext\" using the uninitialized extension \"SELinux\". xselinux.Register(xconn) must be called first.")
	}
	return c.Send(getWindowContextRequest(op, Window))
}

// GetWindowContextReply represents the data returned from a GetWindowContext request.
type GetWindowContextReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	ContextLen uint32
	// padding: 20 bytes
	Context string // size: internal.Pad4((int(ContextLen) * 1))
}

// Unmarshal reads a byte slice into a GetWindowContextReply value.
func (v *GetWindowContextReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.ContextLen) * 1))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetWindowContextReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.ContextLen = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 20 // padding

	{
		byteString := make([]byte, v.ContextLen)
		copy(byteString[:v.ContextLen], buf[b:])
		v.Context = string(byteString)
		b += int(v.ContextLen)
	}

	return nil
}

// Write request to wire for GetWindowContext
// getWindowContextRequest writes a GetWindowContext request to a byte slice.
func getWindowContextRequest(opcode uint8, Window xproto.Window) []byte {
	const size = 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 7 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	return buf
}

// GetWindowCreateContext sends a checked request.
func GetWindowCreateContext(c *xgb.XConn) (GetWindowCreateContextReply, error) {
	var reply GetWindowCreateContextReply
	op, ok := c.Ext("SELinux")
	if !ok {
		return reply, errors.New("cannot issue request \"GetWindowCreateContext\" using the uninitialized extension \"SELinux\". xselinux.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getWindowCreateContextRequest(op), &reply)
	return reply, err
}

// GetWindowCreateContextUnchecked sends an unchecked request.
func GetWindowCreateContextUnchecked(c *xgb.XConn) error {
	op, ok := c.Ext("SELinux")
	if !ok {
		return errors.New("cannot issue request \"GetWindowCreateContext\" using the uninitialized extension \"SELinux\". xselinux.Register(xconn) must be called first.")
	}
	return c.Send(getWindowCreateContextRequest(op))
}

// GetWindowCreateContextReply represents the data returned from a GetWindowCreateContext request.
type GetWindowCreateContextReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	ContextLen uint32
	// padding: 20 bytes
	Context string // size: internal.Pad4((int(ContextLen) * 1))
}

// Unmarshal reads a byte slice into a GetWindowCreateContextReply value.
func (v *GetWindowCreateContextReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.ContextLen) * 1))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetWindowCreateContextReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.ContextLen = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 20 // padding

	{
		byteString := make([]byte, v.ContextLen)
		copy(byteString[:v.ContextLen], buf[b:])
		v.Context = string(byteString)
		b += int(v.ContextLen)
	}

	return nil
}

// Write request to wire for GetWindowCreateContext
// getWindowCreateContextRequest writes a GetWindowCreateContext request to a byte slice.
func getWindowCreateContextRequest(opcode uint8) []byte {
	const size = 4
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

// ListProperties sends a checked request.
func ListProperties(c *xgb.XConn, Window xproto.Window) (ListPropertiesReply, error) {
	var reply ListPropertiesReply
	op, ok := c.Ext("SELinux")
	if !ok {
		return reply, errors.New("cannot issue request \"ListProperties\" using the uninitialized extension \"SELinux\". xselinux.Register(xconn) must be called first.")
	}
	err := c.SendRecv(listPropertiesRequest(op, Window), &reply)
	return reply, err
}

// ListPropertiesUnchecked sends an unchecked request.
func ListPropertiesUnchecked(c *xgb.XConn, Window xproto.Window) error {
	op, ok := c.Ext("SELinux")
	if !ok {
		return errors.New("cannot issue request \"ListProperties\" using the uninitialized extension \"SELinux\". xselinux.Register(xconn) must be called first.")
	}
	return c.Send(listPropertiesRequest(op, Window))
}

// ListPropertiesReply represents the data returned from a ListProperties request.
type ListPropertiesReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	PropertiesLen uint32
	// padding: 20 bytes
	Properties []ListItem // size: ListItemListSize(Properties)
}

// Unmarshal reads a byte slice into a ListPropertiesReply value.
func (v *ListPropertiesReply) Unmarshal(buf []byte) error {
	if size := (32 + ListItemListSize(v.Properties)); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"ListPropertiesReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.PropertiesLen = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 20 // padding

	v.Properties = make([]ListItem, v.PropertiesLen)
	b += ListItemReadList(buf[b:], v.Properties)

	return nil
}

// Write request to wire for ListProperties
// listPropertiesRequest writes a ListProperties request to a byte slice.
func listPropertiesRequest(opcode uint8, Window xproto.Window) []byte {
	const size = 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 14 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	return buf
}

// ListSelections sends a checked request.
func ListSelections(c *xgb.XConn) (ListSelectionsReply, error) {
	var reply ListSelectionsReply
	op, ok := c.Ext("SELinux")
	if !ok {
		return reply, errors.New("cannot issue request \"ListSelections\" using the uninitialized extension \"SELinux\". xselinux.Register(xconn) must be called first.")
	}
	err := c.SendRecv(listSelectionsRequest(op), &reply)
	return reply, err
}

// ListSelectionsUnchecked sends an unchecked request.
func ListSelectionsUnchecked(c *xgb.XConn) error {
	op, ok := c.Ext("SELinux")
	if !ok {
		return errors.New("cannot issue request \"ListSelections\" using the uninitialized extension \"SELinux\". xselinux.Register(xconn) must be called first.")
	}
	return c.Send(listSelectionsRequest(op))
}

// ListSelectionsReply represents the data returned from a ListSelections request.
type ListSelectionsReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	SelectionsLen uint32
	// padding: 20 bytes
	Selections []ListItem // size: ListItemListSize(Selections)
}

// Unmarshal reads a byte slice into a ListSelectionsReply value.
func (v *ListSelectionsReply) Unmarshal(buf []byte) error {
	if size := (32 + ListItemListSize(v.Selections)); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"ListSelectionsReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.SelectionsLen = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 20 // padding

	v.Selections = make([]ListItem, v.SelectionsLen)
	b += ListItemReadList(buf[b:], v.Selections)

	return nil
}

// Write request to wire for ListSelections
// listSelectionsRequest writes a ListSelections request to a byte slice.
func listSelectionsRequest(opcode uint8) []byte {
	const size = 4
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 21 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	return buf
}

// QueryVersion sends a checked request.
func QueryVersion(c *xgb.XConn, ClientMajor byte, ClientMinor byte) (QueryVersionReply, error) {
	var reply QueryVersionReply
	op, ok := c.Ext("SELinux")
	if !ok {
		return reply, errors.New("cannot issue request \"QueryVersion\" using the uninitialized extension \"SELinux\". xselinux.Register(xconn) must be called first.")
	}
	err := c.SendRecv(queryVersionRequest(op, ClientMajor, ClientMinor), &reply)
	return reply, err
}

// QueryVersionUnchecked sends an unchecked request.
func QueryVersionUnchecked(c *xgb.XConn, ClientMajor byte, ClientMinor byte) error {
	op, ok := c.Ext("SELinux")
	if !ok {
		return errors.New("cannot issue request \"QueryVersion\" using the uninitialized extension \"SELinux\". xselinux.Register(xconn) must be called first.")
	}
	return c.Send(queryVersionRequest(op, ClientMajor, ClientMinor))
}

// QueryVersionReply represents the data returned from a QueryVersion request.
type QueryVersionReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	ServerMajor uint16
	ServerMinor uint16
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

	v.ServerMajor = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.ServerMinor = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	return nil
}

// Write request to wire for QueryVersion
// queryVersionRequest writes a QueryVersion request to a byte slice.
func queryVersionRequest(opcode uint8, ClientMajor byte, ClientMinor byte) []byte {
	const size = 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 0 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	buf[b] = ClientMajor
	b += 1

	buf[b] = ClientMinor
	b += 1

	return buf
}

// SetDeviceContext sends a checked request.
func SetDeviceContext(c *xgb.XConn, Device uint32, ContextLen uint32, Context string) error {
	op, ok := c.Ext("SELinux")
	if !ok {
		return errors.New("cannot issue request \"SetDeviceContext\" using the uninitialized extension \"SELinux\". xselinux.Register(xconn) must be called first.")
	}
	return c.SendRecv(setDeviceContextRequest(op, Device, ContextLen, Context), nil)
}

// SetDeviceContextUnchecked sends an unchecked request.
func SetDeviceContextUnchecked(c *xgb.XConn, Device uint32, ContextLen uint32, Context string) error {
	op, ok := c.Ext("SELinux")
	if !ok {
		return errors.New("cannot issue request \"SetDeviceContext\" using the uninitialized extension \"SELinux\". xselinux.Register(xconn) must be called first.")
	}
	return c.Send(setDeviceContextRequest(op, Device, ContextLen, Context))
}

// Write request to wire for SetDeviceContext
// setDeviceContextRequest writes a SetDeviceContext request to a byte slice.
func setDeviceContextRequest(opcode uint8, Device uint32, ContextLen uint32, Context string) []byte {
	size := internal.Pad4((12 + internal.Pad4((int(ContextLen) * 1))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 3 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], Device)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], ContextLen)
	b += 4

	copy(buf[b:], Context[:ContextLen])
	b += int(ContextLen)

	return buf
}

// SetDeviceCreateContext sends a checked request.
func SetDeviceCreateContext(c *xgb.XConn, ContextLen uint32, Context string) error {
	op, ok := c.Ext("SELinux")
	if !ok {
		return errors.New("cannot issue request \"SetDeviceCreateContext\" using the uninitialized extension \"SELinux\". xselinux.Register(xconn) must be called first.")
	}
	return c.SendRecv(setDeviceCreateContextRequest(op, ContextLen, Context), nil)
}

// SetDeviceCreateContextUnchecked sends an unchecked request.
func SetDeviceCreateContextUnchecked(c *xgb.XConn, ContextLen uint32, Context string) error {
	op, ok := c.Ext("SELinux")
	if !ok {
		return errors.New("cannot issue request \"SetDeviceCreateContext\" using the uninitialized extension \"SELinux\". xselinux.Register(xconn) must be called first.")
	}
	return c.Send(setDeviceCreateContextRequest(op, ContextLen, Context))
}

// Write request to wire for SetDeviceCreateContext
// setDeviceCreateContextRequest writes a SetDeviceCreateContext request to a byte slice.
func setDeviceCreateContextRequest(opcode uint8, ContextLen uint32, Context string) []byte {
	size := internal.Pad4((8 + internal.Pad4((int(ContextLen) * 1))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 1 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], ContextLen)
	b += 4

	copy(buf[b:], Context[:ContextLen])
	b += int(ContextLen)

	return buf
}

// SetPropertyCreateContext sends a checked request.
func SetPropertyCreateContext(c *xgb.XConn, ContextLen uint32, Context string) error {
	op, ok := c.Ext("SELinux")
	if !ok {
		return errors.New("cannot issue request \"SetPropertyCreateContext\" using the uninitialized extension \"SELinux\". xselinux.Register(xconn) must be called first.")
	}
	return c.SendRecv(setPropertyCreateContextRequest(op, ContextLen, Context), nil)
}

// SetPropertyCreateContextUnchecked sends an unchecked request.
func SetPropertyCreateContextUnchecked(c *xgb.XConn, ContextLen uint32, Context string) error {
	op, ok := c.Ext("SELinux")
	if !ok {
		return errors.New("cannot issue request \"SetPropertyCreateContext\" using the uninitialized extension \"SELinux\". xselinux.Register(xconn) must be called first.")
	}
	return c.Send(setPropertyCreateContextRequest(op, ContextLen, Context))
}

// Write request to wire for SetPropertyCreateContext
// setPropertyCreateContextRequest writes a SetPropertyCreateContext request to a byte slice.
func setPropertyCreateContextRequest(opcode uint8, ContextLen uint32, Context string) []byte {
	size := internal.Pad4((8 + internal.Pad4((int(ContextLen) * 1))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 8 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], ContextLen)
	b += 4

	copy(buf[b:], Context[:ContextLen])
	b += int(ContextLen)

	return buf
}

// SetPropertyUseContext sends a checked request.
func SetPropertyUseContext(c *xgb.XConn, ContextLen uint32, Context string) error {
	op, ok := c.Ext("SELinux")
	if !ok {
		return errors.New("cannot issue request \"SetPropertyUseContext\" using the uninitialized extension \"SELinux\". xselinux.Register(xconn) must be called first.")
	}
	return c.SendRecv(setPropertyUseContextRequest(op, ContextLen, Context), nil)
}

// SetPropertyUseContextUnchecked sends an unchecked request.
func SetPropertyUseContextUnchecked(c *xgb.XConn, ContextLen uint32, Context string) error {
	op, ok := c.Ext("SELinux")
	if !ok {
		return errors.New("cannot issue request \"SetPropertyUseContext\" using the uninitialized extension \"SELinux\". xselinux.Register(xconn) must be called first.")
	}
	return c.Send(setPropertyUseContextRequest(op, ContextLen, Context))
}

// Write request to wire for SetPropertyUseContext
// setPropertyUseContextRequest writes a SetPropertyUseContext request to a byte slice.
func setPropertyUseContextRequest(opcode uint8, ContextLen uint32, Context string) []byte {
	size := internal.Pad4((8 + internal.Pad4((int(ContextLen) * 1))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 10 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], ContextLen)
	b += 4

	copy(buf[b:], Context[:ContextLen])
	b += int(ContextLen)

	return buf
}

// SetSelectionCreateContext sends a checked request.
func SetSelectionCreateContext(c *xgb.XConn, ContextLen uint32, Context string) error {
	op, ok := c.Ext("SELinux")
	if !ok {
		return errors.New("cannot issue request \"SetSelectionCreateContext\" using the uninitialized extension \"SELinux\". xselinux.Register(xconn) must be called first.")
	}
	return c.SendRecv(setSelectionCreateContextRequest(op, ContextLen, Context), nil)
}

// SetSelectionCreateContextUnchecked sends an unchecked request.
func SetSelectionCreateContextUnchecked(c *xgb.XConn, ContextLen uint32, Context string) error {
	op, ok := c.Ext("SELinux")
	if !ok {
		return errors.New("cannot issue request \"SetSelectionCreateContext\" using the uninitialized extension \"SELinux\". xselinux.Register(xconn) must be called first.")
	}
	return c.Send(setSelectionCreateContextRequest(op, ContextLen, Context))
}

// Write request to wire for SetSelectionCreateContext
// setSelectionCreateContextRequest writes a SetSelectionCreateContext request to a byte slice.
func setSelectionCreateContextRequest(opcode uint8, ContextLen uint32, Context string) []byte {
	size := internal.Pad4((8 + internal.Pad4((int(ContextLen) * 1))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 15 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], ContextLen)
	b += 4

	copy(buf[b:], Context[:ContextLen])
	b += int(ContextLen)

	return buf
}

// SetSelectionUseContext sends a checked request.
func SetSelectionUseContext(c *xgb.XConn, ContextLen uint32, Context string) error {
	op, ok := c.Ext("SELinux")
	if !ok {
		return errors.New("cannot issue request \"SetSelectionUseContext\" using the uninitialized extension \"SELinux\". xselinux.Register(xconn) must be called first.")
	}
	return c.SendRecv(setSelectionUseContextRequest(op, ContextLen, Context), nil)
}

// SetSelectionUseContextUnchecked sends an unchecked request.
func SetSelectionUseContextUnchecked(c *xgb.XConn, ContextLen uint32, Context string) error {
	op, ok := c.Ext("SELinux")
	if !ok {
		return errors.New("cannot issue request \"SetSelectionUseContext\" using the uninitialized extension \"SELinux\". xselinux.Register(xconn) must be called first.")
	}
	return c.Send(setSelectionUseContextRequest(op, ContextLen, Context))
}

// Write request to wire for SetSelectionUseContext
// setSelectionUseContextRequest writes a SetSelectionUseContext request to a byte slice.
func setSelectionUseContextRequest(opcode uint8, ContextLen uint32, Context string) []byte {
	size := internal.Pad4((8 + internal.Pad4((int(ContextLen) * 1))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 17 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], ContextLen)
	b += 4

	copy(buf[b:], Context[:ContextLen])
	b += int(ContextLen)

	return buf
}

// SetWindowCreateContext sends a checked request.
func SetWindowCreateContext(c *xgb.XConn, ContextLen uint32, Context string) error {
	op, ok := c.Ext("SELinux")
	if !ok {
		return errors.New("cannot issue request \"SetWindowCreateContext\" using the uninitialized extension \"SELinux\". xselinux.Register(xconn) must be called first.")
	}
	return c.SendRecv(setWindowCreateContextRequest(op, ContextLen, Context), nil)
}

// SetWindowCreateContextUnchecked sends an unchecked request.
func SetWindowCreateContextUnchecked(c *xgb.XConn, ContextLen uint32, Context string) error {
	op, ok := c.Ext("SELinux")
	if !ok {
		return errors.New("cannot issue request \"SetWindowCreateContext\" using the uninitialized extension \"SELinux\". xselinux.Register(xconn) must be called first.")
	}
	return c.Send(setWindowCreateContextRequest(op, ContextLen, Context))
}

// Write request to wire for SetWindowCreateContext
// setWindowCreateContextRequest writes a SetWindowCreateContext request to a byte slice.
func setWindowCreateContextRequest(opcode uint8, ContextLen uint32, Context string) []byte {
	size := internal.Pad4((8 + internal.Pad4((int(ContextLen) * 1))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 5 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], ContextLen)
	b += 4

	copy(buf[b:], Context[:ContextLen])
	b += int(ContextLen)

	return buf
}
