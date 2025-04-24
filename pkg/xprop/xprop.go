package xprop

import (
	"codeberg.org/gruf/go-byteutil"
	"encoding/binary"
	"fmt"

	"github.com/probakowski/go-xgb/xproto"
)

// PropValAtom transforms a GetPropertyReply struct into an ATOM name. The property reply must be in 32 bit format.
func PropValAtom(conn *XPropConn, reply xproto.GetPropertyReply) (string, error) {
	if reply.Format != 32 {
		return "", fmt.Errorf("expected reply format 32 but got %d", reply.Format)
	}
	return conn.AtomName(xproto.Atom(binary.LittleEndian.Uint32(reply.Value)))
}

func PropValAtoms(conn *XPropConn, reply xproto.GetPropertyReply) ([]xproto.Atom, error) {
	if reply.Format != 32 {
		return nil, fmt.Errorf("expected reply format 32 but got %d", reply.Format)
	}

	ids := make([]xproto.Atom, reply.ValueLen)
	vals := reply.Value

	for i := 0; len(vals) >= 4; i++ {
		ids[i] = xproto.Atom(binary.LittleEndian.Uint32(vals))
		vals = vals[4:]
	}

	return ids, nil
}

func PropValAtomNames(conn *XPropConn, reply xproto.GetPropertyReply) ([]string, error) {
	if reply.Format != 32 {
		return nil, fmt.Errorf("expected reply format 32 but got %d", reply.Format)
	}

	names := make([]string, reply.ValueLen)
	vals := reply.Value

	for i := 0; len(vals) >= 4; i++ {
		var err error

		names[i], err = conn.AtomName(xproto.Atom(binary.LittleEndian.Uint32(vals)))
		if err != nil {
			return nil, err
		}

		vals = vals[4:]
	}

	return names, nil
}

// PropValWindow transforms a GetPropertyReply struct into an X resource window identifier. The property reply must be in 32 bit format.
func PropValWindow(reply xproto.GetPropertyReply) (xproto.Window, error) {
	if reply.Format != 32 {
		return 0, fmt.Errorf("expected reply format 32 but got %d",
			reply.Format)
	}
	return xproto.Window(binary.LittleEndian.Uint32(reply.Value)), nil
}

// PropValWindows is the same as PropValWindow, except that it returns a slice of identifiers. Also must be 32 bit format.
func PropValWindows(reply xproto.GetPropertyReply) ([]xproto.Window, error) {
	if reply.Format != 32 {
		return nil, fmt.Errorf("expected reply format 32 but got %d",
			reply.Format)
	}
	ids := make([]xproto.Window, reply.ValueLen)
	vals := reply.Value
	for i := 0; len(vals) >= 4; i++ {
		ids[i] = xproto.Window(binary.LittleEndian.Uint32(vals))
		vals = vals[4:]
	}
	return ids, nil
}

// PropValNum transforms a GetPropertyReply struct into an unsigned integer. Useful when the property value is a single integer.
func PropValUint32(reply xproto.GetPropertyReply) (uint32, error) {
	if reply.Format != 32 {
		return 0, fmt.Errorf("expected reply format 32 but got %d", reply.Format)
	}
	return binary.LittleEndian.Uint32(reply.Value), nil
}

// PropValNums is the same as PropValNum, except that it returns a slice of integers. Also must be 32 bit format.
func PropValUint32s(reply xproto.GetPropertyReply) ([]uint32, error) {
	if reply.Format != 32 {
		return nil, fmt.Errorf("expected reply format 32 but got %d", reply.Format)
	}

	nums := make([]uint32, reply.ValueLen)
	vals := reply.Value

	for i := 0; len(vals) >= 4; i++ {
		nums[i] = binary.LittleEndian.Uint32(vals)
		vals = vals[4:]
	}

	return nums, nil
}

// PropValInt32 transforms a GetPropertyReply struct into a 64 bit
// integer. Useful when the property value is a single integer.
func PropValInt32(reply xproto.GetPropertyReply) (int32, error) {
	if reply.Format != 32 {
		return 0, fmt.Errorf("expected reply format 32 but got %d",
			reply.Format)
	}
	return int32(binary.LittleEndian.Uint32(reply.Value)), nil
}

// PropValStr transforms a GetPropertyReply struct into a string.
// Useful when the property value is a null terminated string represented
// by integers. Also must be 8 bit format.
func PropValStr(reply xproto.GetPropertyReply) (string, error) {
	if reply.Format != 8 {
		return "", fmt.Errorf("expected reply format 8 but got %d", reply.Format)
	}
	return byteutil.B2S(reply.Value), nil
}

// PropValStrs is the same as PropValStr, except that it returns a slice
// of strings. The raw byte string is a sequence of null terminated strings,
// which is translated into a slice of strings.
func PropValStrs(reply xproto.GetPropertyReply) ([]string, error) {
	if reply.Format != 8 {
		return nil, fmt.Errorf("expected reply format 8 but got %d", reply.Format)
	}
	var strs []string
	sstart := 0
	for i, c := range reply.Value {
		if c == 0 {
			strs = append(strs, byteutil.B2S(reply.Value[sstart:i]))
			sstart = i + 1
		}
	}
	if sstart < int(reply.ValueLen) {
		strs = append(strs, byteutil.B2S(reply.Value[sstart:]))
	}
	return strs, nil
}
