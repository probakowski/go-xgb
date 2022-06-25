// FILE GENERATED AUTOMATICALLY FROM "composite.xml"
package composite

import (
	"encoding/binary"
	"errors"
	"fmt"

	"codeberg.org/gruf/go-xgb"
	"codeberg.org/gruf/go-xgb/xfixes"
	"codeberg.org/gruf/go-xgb/xproto"
)

const (
	// ExtName is the user-friendly name string of this X extension.
	ExtName = "Composite"

	// ExtXName is the name string this extension is known by to the X server.
	ExtXName = "Composite"
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
		return fmt.Errorf("error querying X for \"Composite\": %w", err)
	} else if !reply.Present {
		return fmt.Errorf("no extension named \"Composite\" is known to the X server: reply=%+v", reply)
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
	RedirectAutomatic = 0
	RedirectManual    = 1
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

// CreateRegionFromBorderClip sends a checked request.
func CreateRegionFromBorderClip(c *xgb.XConn, Region xfixes.Region, Window xproto.Window) error {
	op, ok := c.Ext("Composite")
	if !ok {
		return errors.New("cannot issue request \"CreateRegionFromBorderClip\" using the uninitialized extension \"Composite\". composite.Register(xconn) must be called first.")
	}
	return c.SendRecv(createRegionFromBorderClipRequest(op, Region, Window), nil)
}

// CreateRegionFromBorderClipUnchecked sends an unchecked request.
func CreateRegionFromBorderClipUnchecked(c *xgb.XConn, Region xfixes.Region, Window xproto.Window) error {
	op, ok := c.Ext("Composite")
	if !ok {
		return errors.New("cannot issue request \"CreateRegionFromBorderClip\" using the uninitialized extension \"Composite\". composite.Register(xconn) must be called first.")
	}
	return c.Send(createRegionFromBorderClipRequest(op, Region, Window))
}

// Write request to wire for CreateRegionFromBorderClip
// createRegionFromBorderClipRequest writes a CreateRegionFromBorderClip request to a byte slice.
func createRegionFromBorderClipRequest(opcode uint8, Region xfixes.Region, Window xproto.Window) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 5 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Region))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	return buf
}

// GetOverlayWindow sends a checked request.
func GetOverlayWindow(c *xgb.XConn, Window xproto.Window) (GetOverlayWindowReply, error) {
	var reply GetOverlayWindowReply
	op, ok := c.Ext("Composite")
	if !ok {
		return reply, errors.New("cannot issue request \"GetOverlayWindow\" using the uninitialized extension \"Composite\". composite.Register(xconn) must be called first.")
	}
	err := c.SendRecv(getOverlayWindowRequest(op, Window), &reply)
	return reply, err
}

// GetOverlayWindowUnchecked sends an unchecked request.
func GetOverlayWindowUnchecked(c *xgb.XConn, Window xproto.Window) error {
	op, ok := c.Ext("Composite")
	if !ok {
		return errors.New("cannot issue request \"GetOverlayWindow\" using the uninitialized extension \"Composite\". composite.Register(xconn) must be called first.")
	}
	return c.Send(getOverlayWindowRequest(op, Window))
}

// GetOverlayWindowReply represents the data returned from a GetOverlayWindow request.
type GetOverlayWindowReply struct {
	Sequence uint16 // sequence number of the request for this reply
	Length   uint32 // number of bytes in this reply
	// padding: 1 bytes
	OverlayWin xproto.Window
	// padding: 20 bytes
}

// Unmarshal reads a byte slice into a GetOverlayWindowReply value.
func (v *GetOverlayWindowReply) Unmarshal(buf []byte) error {
	if size := 32; len(buf) < size {
		return fmt.Errorf("not enough data to unmarshal \"GetOverlayWindowReply\": have=%d need=%d", len(buf), size)
	}

	b := 1 // skip reply determinant

	b += 1 // padding

	v.Sequence = binary.LittleEndian.Uint16(buf[b:])
	b += 2

	v.Length = binary.LittleEndian.Uint32(buf[b:]) // 4-byte units
	b += 4

	v.OverlayWin = xproto.Window(binary.LittleEndian.Uint32(buf[b:]))
	b += 4

	b += 20 // padding

	return nil
}

// Write request to wire for GetOverlayWindow
// getOverlayWindowRequest writes a GetOverlayWindow request to a byte slice.
func getOverlayWindowRequest(opcode uint8, Window xproto.Window) []byte {
	size := 8
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

// NameWindowPixmap sends a checked request.
func NameWindowPixmap(c *xgb.XConn, Window xproto.Window, Pixmap xproto.Pixmap) error {
	op, ok := c.Ext("Composite")
	if !ok {
		return errors.New("cannot issue request \"NameWindowPixmap\" using the uninitialized extension \"Composite\". composite.Register(xconn) must be called first.")
	}
	return c.SendRecv(nameWindowPixmapRequest(op, Window, Pixmap), nil)
}

// NameWindowPixmapUnchecked sends an unchecked request.
func NameWindowPixmapUnchecked(c *xgb.XConn, Window xproto.Window, Pixmap xproto.Pixmap) error {
	op, ok := c.Ext("Composite")
	if !ok {
		return errors.New("cannot issue request \"NameWindowPixmap\" using the uninitialized extension \"Composite\". composite.Register(xconn) must be called first.")
	}
	return c.Send(nameWindowPixmapRequest(op, Window, Pixmap))
}

// Write request to wire for NameWindowPixmap
// nameWindowPixmapRequest writes a NameWindowPixmap request to a byte slice.
func nameWindowPixmapRequest(opcode uint8, Window xproto.Window, Pixmap xproto.Pixmap) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 6 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	binary.LittleEndian.PutUint32(buf[b:], uint32(Pixmap))
	b += 4

	return buf
}

// QueryVersion sends a checked request.
func QueryVersion(c *xgb.XConn, ClientMajorVersion uint32, ClientMinorVersion uint32) (QueryVersionReply, error) {
	var reply QueryVersionReply
	op, ok := c.Ext("Composite")
	if !ok {
		return reply, errors.New("cannot issue request \"QueryVersion\" using the uninitialized extension \"Composite\". composite.Register(xconn) must be called first.")
	}
	err := c.SendRecv(queryVersionRequest(op, ClientMajorVersion, ClientMinorVersion), &reply)
	return reply, err
}

// QueryVersionUnchecked sends an unchecked request.
func QueryVersionUnchecked(c *xgb.XConn, ClientMajorVersion uint32, ClientMinorVersion uint32) error {
	op, ok := c.Ext("Composite")
	if !ok {
		return errors.New("cannot issue request \"QueryVersion\" using the uninitialized extension \"Composite\". composite.Register(xconn) must be called first.")
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

// RedirectSubwindows sends a checked request.
func RedirectSubwindows(c *xgb.XConn, Window xproto.Window, Update byte) error {
	op, ok := c.Ext("Composite")
	if !ok {
		return errors.New("cannot issue request \"RedirectSubwindows\" using the uninitialized extension \"Composite\". composite.Register(xconn) must be called first.")
	}
	return c.SendRecv(redirectSubwindowsRequest(op, Window, Update), nil)
}

// RedirectSubwindowsUnchecked sends an unchecked request.
func RedirectSubwindowsUnchecked(c *xgb.XConn, Window xproto.Window, Update byte) error {
	op, ok := c.Ext("Composite")
	if !ok {
		return errors.New("cannot issue request \"RedirectSubwindows\" using the uninitialized extension \"Composite\". composite.Register(xconn) must be called first.")
	}
	return c.Send(redirectSubwindowsRequest(op, Window, Update))
}

// Write request to wire for RedirectSubwindows
// redirectSubwindowsRequest writes a RedirectSubwindows request to a byte slice.
func redirectSubwindowsRequest(opcode uint8, Window xproto.Window, Update byte) []byte {
	size := 12
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

	buf[b] = Update
	b += 1

	b += 3 // padding

	return buf
}

// RedirectWindow sends a checked request.
func RedirectWindow(c *xgb.XConn, Window xproto.Window, Update byte) error {
	op, ok := c.Ext("Composite")
	if !ok {
		return errors.New("cannot issue request \"RedirectWindow\" using the uninitialized extension \"Composite\". composite.Register(xconn) must be called first.")
	}
	return c.SendRecv(redirectWindowRequest(op, Window, Update), nil)
}

// RedirectWindowUnchecked sends an unchecked request.
func RedirectWindowUnchecked(c *xgb.XConn, Window xproto.Window, Update byte) error {
	op, ok := c.Ext("Composite")
	if !ok {
		return errors.New("cannot issue request \"RedirectWindow\" using the uninitialized extension \"Composite\". composite.Register(xconn) must be called first.")
	}
	return c.Send(redirectWindowRequest(op, Window, Update))
}

// Write request to wire for RedirectWindow
// redirectWindowRequest writes a RedirectWindow request to a byte slice.
func redirectWindowRequest(opcode uint8, Window xproto.Window, Update byte) []byte {
	size := 12
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

	buf[b] = Update
	b += 1

	b += 3 // padding

	return buf
}

// ReleaseOverlayWindow sends a checked request.
func ReleaseOverlayWindow(c *xgb.XConn, Window xproto.Window) error {
	op, ok := c.Ext("Composite")
	if !ok {
		return errors.New("cannot issue request \"ReleaseOverlayWindow\" using the uninitialized extension \"Composite\". composite.Register(xconn) must be called first.")
	}
	return c.SendRecv(releaseOverlayWindowRequest(op, Window), nil)
}

// ReleaseOverlayWindowUnchecked sends an unchecked request.
func ReleaseOverlayWindowUnchecked(c *xgb.XConn, Window xproto.Window) error {
	op, ok := c.Ext("Composite")
	if !ok {
		return errors.New("cannot issue request \"ReleaseOverlayWindow\" using the uninitialized extension \"Composite\". composite.Register(xconn) must be called first.")
	}
	return c.Send(releaseOverlayWindowRequest(op, Window))
}

// Write request to wire for ReleaseOverlayWindow
// releaseOverlayWindowRequest writes a ReleaseOverlayWindow request to a byte slice.
func releaseOverlayWindowRequest(opcode uint8, Window xproto.Window) []byte {
	size := 8
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 8 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	return buf
}

// UnredirectSubwindows sends a checked request.
func UnredirectSubwindows(c *xgb.XConn, Window xproto.Window, Update byte) error {
	op, ok := c.Ext("Composite")
	if !ok {
		return errors.New("cannot issue request \"UnredirectSubwindows\" using the uninitialized extension \"Composite\". composite.Register(xconn) must be called first.")
	}
	return c.SendRecv(unredirectSubwindowsRequest(op, Window, Update), nil)
}

// UnredirectSubwindowsUnchecked sends an unchecked request.
func UnredirectSubwindowsUnchecked(c *xgb.XConn, Window xproto.Window, Update byte) error {
	op, ok := c.Ext("Composite")
	if !ok {
		return errors.New("cannot issue request \"UnredirectSubwindows\" using the uninitialized extension \"Composite\". composite.Register(xconn) must be called first.")
	}
	return c.Send(unredirectSubwindowsRequest(op, Window, Update))
}

// Write request to wire for UnredirectSubwindows
// unredirectSubwindowsRequest writes a UnredirectSubwindows request to a byte slice.
func unredirectSubwindowsRequest(opcode uint8, Window xproto.Window, Update byte) []byte {
	size := 12
	b := 0
	buf := make([]byte, size)

	buf[b] = opcode
	b += 1

	buf[b] = 4 // request opcode
	b += 1

	binary.LittleEndian.PutUint16(buf[b:], uint16(size/4)) // write request size in 4-byte units
	b += 2

	binary.LittleEndian.PutUint32(buf[b:], uint32(Window))
	b += 4

	buf[b] = Update
	b += 1

	b += 3 // padding

	return buf
}

// UnredirectWindow sends a checked request.
func UnredirectWindow(c *xgb.XConn, Window xproto.Window, Update byte) error {
	op, ok := c.Ext("Composite")
	if !ok {
		return errors.New("cannot issue request \"UnredirectWindow\" using the uninitialized extension \"Composite\". composite.Register(xconn) must be called first.")
	}
	return c.SendRecv(unredirectWindowRequest(op, Window, Update), nil)
}

// UnredirectWindowUnchecked sends an unchecked request.
func UnredirectWindowUnchecked(c *xgb.XConn, Window xproto.Window, Update byte) error {
	op, ok := c.Ext("Composite")
	if !ok {
		return errors.New("cannot issue request \"UnredirectWindow\" using the uninitialized extension \"Composite\". composite.Register(xconn) must be called first.")
	}
	return c.Send(unredirectWindowRequest(op, Window, Update))
}

// Write request to wire for UnredirectWindow
// unredirectWindowRequest writes a UnredirectWindow request to a byte slice.
func unredirectWindowRequest(opcode uint8, Window xproto.Window, Update byte) []byte {
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

	buf[b] = Update
	b += 1

	b += 3 // padding

	return buf
}
