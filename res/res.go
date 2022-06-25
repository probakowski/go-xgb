// FILE GENERATED AUTOMATICALLY FROM "res.xml"
package res

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
	ExtName = "Res"

	// ExtXName is the name string this extension is known by to the X server.
	ExtXName = "X-Resource"
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
		return fmt.Errorf("error querying X for \"Res\": %w", err)
	} else if !reply.Present {
		return fmt.Errorf("no extension named \"Res\" is known to the X server: reply=%+v", reply)
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

type Client struct {
	ResourceBase uint32
	ResourceMask uint32
}

// ClientRead reads a byte slice into a Client value.
func ClientRead(buf []byte, v *Client) int {
	b := 0

	v.ResourceBase = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.ResourceMask = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	return b
}

// ClientReadList reads a byte slice into a list of Client values.
func ClientReadList(buf []byte, dest []Client) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Client{}
		b += ClientRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a Client value to a byte slice.
func (v Client) Bytes() []byte {
	buf := make([]byte, 8)
	b := 0

	binary.LittleEndian.PutUint32(buf[b:], v.ResourceBase)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], v.ResourceMask)
	b += 4

	return buf[:b]
}

// ClientListBytes writes a list of Client values to a byte slice.
func ClientListBytes(buf []byte, list []Client) int {
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
	ClientIdMaskClientXID      = 1
	ClientIdMaskLocalClientPID = 2
)

type ClientIdSpec struct {
	Client uint32
	Mask   uint32
}

// ClientIdSpecRead reads a byte slice into a ClientIdSpec value.
func ClientIdSpecRead(buf []byte, v *ClientIdSpec) int {
	b := 0

	v.Client = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.Mask = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	return b
}

// ClientIdSpecReadList reads a byte slice into a list of ClientIdSpec values.
func ClientIdSpecReadList(buf []byte, dest []ClientIdSpec) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = ClientIdSpec{}
		b += ClientIdSpecRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a ClientIdSpec value to a byte slice.
func (v ClientIdSpec) Bytes() []byte {
	buf := make([]byte, 8)
	b := 0

	binary.LittleEndian.PutUint32(buf[b:], v.Client)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], v.Mask)
	b += 4

	return buf[:b]
}

// ClientIdSpecListBytes writes a list of ClientIdSpec values to a byte slice.
func ClientIdSpecListBytes(buf []byte, list []ClientIdSpec) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

type ClientIdValue struct {
	Spec   ClientIdSpec
	Length uint32
	Value  []uint32 // size: internal.Pad4(((int(Length) / 4) * 4))
}

// ClientIdValueRead reads a byte slice into a ClientIdValue value.
func ClientIdValueRead(buf []byte, v *ClientIdValue) int {
	b := 0

	v.Spec = ClientIdSpec{}
	b += ClientIdSpecRead(buf[b:], &v.Spec)

	v.Length = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.Value = make([]uint32, (int(v.Length) / 4))
	for i := 0; i < int((int(v.Length) / 4)); i++ {
		v.Value[i] = binary.LittleEndian.Uint32(buf[b:])
		b += 4
	}

	return b
}

// ClientIdValueReadList reads a byte slice into a list of ClientIdValue values.
func ClientIdValueReadList(buf []byte, dest []ClientIdValue) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = ClientIdValue{}
		b += ClientIdValueRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a ClientIdValue value to a byte slice.
func (v ClientIdValue) Bytes() []byte {
	buf := make([]byte, (12 + internal.Pad4(((int(v.Length) / 4) * 4))))
	b := 0

	{
		structBytes := v.Spec.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	binary.LittleEndian.PutUint32(buf[b:], v.Length)
	b += 4

	for i := 0; i < int((int(v.Length) / 4)); i++ {
		binary.LittleEndian.PutUint32(buf[b:], v.Value[i])
		b += 4
	}

	return buf[:b]
}

// ClientIdValueListBytes writes a list of ClientIdValue values to a byte slice.
func ClientIdValueListBytes(buf []byte, list []ClientIdValue) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

// ClientIdValueListSize computes the size (bytes) of a list of ClientIdValue values.
func ClientIdValueListSize(list []ClientIdValue) int {
	size := 0
	for _, item := range list {
		size += (12 + internal.Pad4(((int(item.Length) / 4) * 4)))
	}
	return size
}

type ResourceIdSpec struct {
	Resource uint32
	Type     uint32
}

// ResourceIdSpecRead reads a byte slice into a ResourceIdSpec value.
func ResourceIdSpecRead(buf []byte, v *ResourceIdSpec) int {
	b := 0

	v.Resource = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.Type = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	return b
}

// ResourceIdSpecReadList reads a byte slice into a list of ResourceIdSpec values.
func ResourceIdSpecReadList(buf []byte, dest []ResourceIdSpec) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = ResourceIdSpec{}
		b += ResourceIdSpecRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a ResourceIdSpec value to a byte slice.
func (v ResourceIdSpec) Bytes() []byte {
	buf := make([]byte, 8)
	b := 0

	binary.LittleEndian.PutUint32(buf[b:], v.Resource)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], v.Type)
	b += 4

	return buf[:b]
}

// ResourceIdSpecListBytes writes a list of ResourceIdSpec values to a byte slice.
func ResourceIdSpecListBytes(buf []byte, list []ResourceIdSpec) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

type ResourceSizeSpec struct {
	Spec     ResourceIdSpec
	Bytes_   uint32
	RefCount uint32
	UseCount uint32
}

// ResourceSizeSpecRead reads a byte slice into a ResourceSizeSpec value.
func ResourceSizeSpecRead(buf []byte, v *ResourceSizeSpec) int {
	b := 0

	v.Spec = ResourceIdSpec{}
	b += ResourceIdSpecRead(buf[b:], &v.Spec)

	v.Bytes_ = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.RefCount = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.UseCount = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	return b
}

// ResourceSizeSpecReadList reads a byte slice into a list of ResourceSizeSpec values.
func ResourceSizeSpecReadList(buf []byte, dest []ResourceSizeSpec) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = ResourceSizeSpec{}
		b += ResourceSizeSpecRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a ResourceSizeSpec value to a byte slice.
func (v ResourceSizeSpec) Bytes() []byte {
	buf := make([]byte, 20)
	b := 0

	{
		structBytes := v.Spec.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	binary.LittleEndian.PutUint32(buf[b:], v.Bytes_)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], v.RefCount)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], v.UseCount)
	b += 4

	return buf[:b]
}

// ResourceSizeSpecListBytes writes a list of ResourceSizeSpec values to a byte slice.
func ResourceSizeSpecListBytes(buf []byte, list []ResourceSizeSpec) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

type ResourceSizeValue struct {
	Size               ResourceSizeSpec
	NumCrossReferences uint32
	CrossReferences    []ResourceSizeSpec // size: internal.Pad4((int(NumCrossReferences) * 20))
}

// ResourceSizeValueRead reads a byte slice into a ResourceSizeValue value.
func ResourceSizeValueRead(buf []byte, v *ResourceSizeValue) int {
	b := 0

	v.Size = ResourceSizeSpec{}
	b += ResourceSizeSpecRead(buf[b:], &v.Size)

	v.NumCrossReferences = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.CrossReferences = make([]ResourceSizeSpec, v.NumCrossReferences)
	b += ResourceSizeSpecReadList(buf[b:], v.CrossReferences)

	return b
}

// ResourceSizeValueReadList reads a byte slice into a list of ResourceSizeValue values.
func ResourceSizeValueReadList(buf []byte, dest []ResourceSizeValue) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = ResourceSizeValue{}
		b += ResourceSizeValueRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a ResourceSizeValue value to a byte slice.
func (v ResourceSizeValue) Bytes() []byte {
	buf := make([]byte, (24 + internal.Pad4((int(v.NumCrossReferences) * 20))))
	b := 0

	{
		structBytes := v.Size.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}

	binary.LittleEndian.PutUint32(buf[b:], v.NumCrossReferences)
	b += 4

	b += ResourceSizeSpecListBytes(buf[b:], v.CrossReferences)

	return buf[:b]
}

// ResourceSizeValueListBytes writes a list of ResourceSizeValue values to a byte slice.
func ResourceSizeValueListBytes(buf []byte, list []ResourceSizeValue) int {
	b := 0
	var structBytes []byte
	for _, item := range list {
		structBytes = item.Bytes()
		copy(buf[b:], structBytes)
		b += len(structBytes)
	}
	return internal.Pad4(b)
}

// ResourceSizeValueListSize computes the size (bytes) of a list of ResourceSizeValue values.
func ResourceSizeValueListSize(list []ResourceSizeValue) int {
	size := 0
	for _, item := range list {
		size += (24 + internal.Pad4((int(item.NumCrossReferences) * 20)))
	}
	return size
}

type Type struct {
	ResourceType xproto.Atom
	Count        uint32
}

// TypeRead reads a byte slice into a Type value.
func TypeRead(buf []byte, v *Type) int {
	b := 0

	v.ResourceType = xproto.Atom(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	v.Count = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	return b
}

// TypeReadList reads a byte slice into a list of Type values.
func TypeReadList(buf []byte, dest []Type) int {
	b := 0
	for i := 0; i < len(dest); i++ {
		dest[i] = Type{}
		b += TypeRead(buf[b:], &dest[i])
	}
	return internal.Pad4(b)
}

// Bytes writes a Type value to a byte slice.
func (v Type) Bytes() []byte {
	buf := make([]byte, 8)
	b := 0

	binary.LittleEndian.PutUint32(buf[b:], uint32(v.ResourceType))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], v.Count)
	b += 4

	return buf[:b]
}

// TypeListBytes writes a list of Type values to a byte slice.
func TypeListBytes(buf []byte, list []Type) int {
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

// QueryClientIds sends a checked request.
func QueryClientIds(c *xgb.XConn, NumSpecs uint32, Specs []ClientIdSpec) (QueryClientIdsReply, error) {
	var reply QueryClientIdsReply
	op, ok := c.Ext("X-Resource")
	if !ok {
		return reply, errors.New("cannot issue request \"QueryClientIds\" using the uninitialized extension \"X-Resource\". res.Register(xconn) must be called first.")
	}
	err := c.SendRecv(queryClientIdsRequest(op, NumSpecs, Specs), &reply)
	return reply, err
}

// QueryClientIdsUnchecked sends an unchecked request.
func QueryClientIdsUnchecked(c *xgb.XConn, NumSpecs uint32, Specs []ClientIdSpec) error {
	op, ok := c.Ext("X-Resource")
	if !ok {
		return errors.New("cannot issue request \"QueryClientIds\" using the uninitialized extension \"X-Resource\". res.Register(xconn) must be called first.")
	}
	return c.Send(queryClientIdsRequest(op, NumSpecs, Specs))
}

// QueryClientIdsReply represents the data returned from a QueryClientIds request.
type QueryClientIdsReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	NumIds uint32
	// padding: 20 bytes
	Ids []ClientIdValue // size: ClientIdValueListSize(Ids)
}

// Unmarshal reads a byte slice into a QueryClientIdsReply value.
func (v *QueryClientIdsReply) Unmarshal(buf []byte) error {
	if size := (32 + ClientIdValueListSize(v.Ids)); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"QueryClientIdsReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.NumIds = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 20 // padding

	v.Ids = make([]ClientIdValue, v.NumIds)
	b += ClientIdValueReadList(buf[b:], v.Ids)

	return nil
}

// Write request to wire for QueryClientIds
// queryClientIdsRequest writes a QueryClientIds request to a byte slice.
func queryClientIdsRequest(opcode uint8, NumSpecs uint32, Specs []ClientIdSpec) []byte {
	size := internal.Pad4((8 + internal.Pad4((int(NumSpecs) * 8))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 4 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], NumSpecs)
	b += 4

	b += ClientIdSpecListBytes(buf[b:], Specs)

	return buf
}

// QueryClientPixmapBytes sends a checked request.
func QueryClientPixmapBytes(c *xgb.XConn, Xid uint32) (QueryClientPixmapBytesReply, error) {
	var reply QueryClientPixmapBytesReply
	op, ok := c.Ext("X-Resource")
	if !ok {
		return reply, errors.New("cannot issue request \"QueryClientPixmapBytes\" using the uninitialized extension \"X-Resource\". res.Register(xconn) must be called first.")
	}
	err := c.SendRecv(queryClientPixmapBytesRequest(op, Xid), &reply)
	return reply, err
}

// QueryClientPixmapBytesUnchecked sends an unchecked request.
func QueryClientPixmapBytesUnchecked(c *xgb.XConn, Xid uint32) error {
	op, ok := c.Ext("X-Resource")
	if !ok {
		return errors.New("cannot issue request \"QueryClientPixmapBytes\" using the uninitialized extension \"X-Resource\". res.Register(xconn) must be called first.")
	}
	return c.Send(queryClientPixmapBytesRequest(op, Xid))
}

// QueryClientPixmapBytesReply represents the data returned from a QueryClientPixmapBytes request.
type QueryClientPixmapBytesReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	Bytes_        uint32
	BytesOverflow uint32
}

// Unmarshal reads a byte slice into a QueryClientPixmapBytesReply value.
func (v *QueryClientPixmapBytesReply) Unmarshal(buf []byte) error {
	if size := 16; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"QueryClientPixmapBytesReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.Bytes_ = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	v.BytesOverflow = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	return nil
}

// Write request to wire for QueryClientPixmapBytes
// queryClientPixmapBytesRequest writes a QueryClientPixmapBytes request to a byte slice.
func queryClientPixmapBytesRequest(opcode uint8, Xid uint32) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 3 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], Xid)
	b += 4

	return buf
}

// QueryClientResources sends a checked request.
func QueryClientResources(c *xgb.XConn, Xid uint32) (QueryClientResourcesReply, error) {
	var reply QueryClientResourcesReply
	op, ok := c.Ext("X-Resource")
	if !ok {
		return reply, errors.New("cannot issue request \"QueryClientResources\" using the uninitialized extension \"X-Resource\". res.Register(xconn) must be called first.")
	}
	err := c.SendRecv(queryClientResourcesRequest(op, Xid), &reply)
	return reply, err
}

// QueryClientResourcesUnchecked sends an unchecked request.
func QueryClientResourcesUnchecked(c *xgb.XConn, Xid uint32) error {
	op, ok := c.Ext("X-Resource")
	if !ok {
		return errors.New("cannot issue request \"QueryClientResources\" using the uninitialized extension \"X-Resource\". res.Register(xconn) must be called first.")
	}
	return c.Send(queryClientResourcesRequest(op, Xid))
}

// QueryClientResourcesReply represents the data returned from a QueryClientResources request.
type QueryClientResourcesReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	NumTypes uint32
	// padding: 20 bytes
	Types []Type // size: internal.Pad4((int(NumTypes) * 8))
}

// Unmarshal reads a byte slice into a QueryClientResourcesReply value.
func (v *QueryClientResourcesReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.NumTypes) * 8))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"QueryClientResourcesReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.NumTypes = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 20 // padding

	v.Types = make([]Type, v.NumTypes)
	b += TypeReadList(buf[b:], v.Types)

	return nil
}

// Write request to wire for QueryClientResources
// queryClientResourcesRequest writes a QueryClientResources request to a byte slice.
func queryClientResourcesRequest(opcode uint8, Xid uint32) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 2 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], Xid)
	b += 4

	return buf
}

// QueryClients sends a checked request.
func QueryClients(c *xgb.XConn) (QueryClientsReply, error) {
	var reply QueryClientsReply
	op, ok := c.Ext("X-Resource")
	if !ok {
		return reply, errors.New("cannot issue request \"QueryClients\" using the uninitialized extension \"X-Resource\". res.Register(xconn) must be called first.")
	}
	err := c.SendRecv(queryClientsRequest(op), &reply)
	return reply, err
}

// QueryClientsUnchecked sends an unchecked request.
func QueryClientsUnchecked(c *xgb.XConn) error {
	op, ok := c.Ext("X-Resource")
	if !ok {
		return errors.New("cannot issue request \"QueryClients\" using the uninitialized extension \"X-Resource\". res.Register(xconn) must be called first.")
	}
	return c.Send(queryClientsRequest(op))
}

// QueryClientsReply represents the data returned from a QueryClients request.
type QueryClientsReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	NumClients uint32
	// padding: 20 bytes
	Clients []Client // size: internal.Pad4((int(NumClients) * 8))
}

// Unmarshal reads a byte slice into a QueryClientsReply value.
func (v *QueryClientsReply) Unmarshal(buf []byte) error {
	if size := (32 + internal.Pad4((int(v.NumClients) * 8))); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"QueryClientsReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.NumClients = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 20 // padding

	v.Clients = make([]Client, v.NumClients)
	b += ClientReadList(buf[b:], v.Clients)

	return nil
}

// Write request to wire for QueryClients
// queryClientsRequest writes a QueryClients request to a byte slice.
func queryClientsRequest(opcode uint8) []byte {
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

// QueryResourceBytes sends a checked request.
func QueryResourceBytes(c *xgb.XConn, Client uint32, NumSpecs uint32, Specs []ResourceIdSpec) (QueryResourceBytesReply, error) {
	var reply QueryResourceBytesReply
	op, ok := c.Ext("X-Resource")
	if !ok {
		return reply, errors.New("cannot issue request \"QueryResourceBytes\" using the uninitialized extension \"X-Resource\". res.Register(xconn) must be called first.")
	}
	err := c.SendRecv(queryResourceBytesRequest(op, Client, NumSpecs, Specs), &reply)
	return reply, err
}

// QueryResourceBytesUnchecked sends an unchecked request.
func QueryResourceBytesUnchecked(c *xgb.XConn, Client uint32, NumSpecs uint32, Specs []ResourceIdSpec) error {
	op, ok := c.Ext("X-Resource")
	if !ok {
		return errors.New("cannot issue request \"QueryResourceBytes\" using the uninitialized extension \"X-Resource\". res.Register(xconn) must be called first.")
	}
	return c.Send(queryResourceBytesRequest(op, Client, NumSpecs, Specs))
}

// QueryResourceBytesReply represents the data returned from a QueryResourceBytes request.
type QueryResourceBytesReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	NumSizes uint32
	// padding: 20 bytes
	Sizes []ResourceSizeValue // size: ResourceSizeValueListSize(Sizes)
}

// Unmarshal reads a byte slice into a QueryResourceBytesReply value.
func (v *QueryResourceBytesReply) Unmarshal(buf []byte) error {
	if size := (32 + ResourceSizeValueListSize(v.Sizes)); len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"QueryResourceBytesReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.NumSizes = binary.LittleEndian.Uint32(buf[b:])
	b += 4

	b += 20 // padding

	v.Sizes = make([]ResourceSizeValue, v.NumSizes)
	b += ResourceSizeValueReadList(buf[b:], v.Sizes)

	return nil
}

// Write request to wire for QueryResourceBytes
// queryResourceBytesRequest writes a QueryResourceBytes request to a byte slice.
func queryResourceBytesRequest(opcode uint8, Client uint32, NumSpecs uint32, Specs []ResourceIdSpec) []byte {
	size := internal.Pad4((12 + internal.Pad4((int(NumSpecs) * 8))))
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 5 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], Client)
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], NumSpecs)
	b += 4

	b += ResourceIdSpecListBytes(buf[b:], Specs)

	return buf
}

// QueryVersion sends a checked request.
func QueryVersion(c *xgb.XConn, ClientMajor byte, ClientMinor byte) (QueryVersionReply, error) {
	var reply QueryVersionReply
	op, ok := c.Ext("X-Resource")
	if !ok {
		return reply, errors.New("cannot issue request \"QueryVersion\" using the uninitialized extension \"X-Resource\". res.Register(xconn) must be called first.")
	}
	err := c.SendRecv(queryVersionRequest(op, ClientMajor, ClientMinor), &reply)
	return reply, err
}

// QueryVersionUnchecked sends an unchecked request.
func QueryVersionUnchecked(c *xgb.XConn, ClientMajor byte, ClientMinor byte) error {
	op, ok := c.Ext("X-Resource")
	if !ok {
		return errors.New("cannot issue request \"QueryVersion\" using the uninitialized extension \"X-Resource\". res.Register(xconn) must be called first.")
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
	if size := 12; len(buf) < size {
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
	size := 8
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
