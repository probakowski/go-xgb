package xprop

import (
	"encoding/binary"
	"fmt"

	"codeberg.org/gruf/go-xgb/xproto"
)

func (conn *XPropConn) GetProp(win xproto.Window, atom xproto.Atom) (xproto.GetPropertyReply, error) {
	atomName, err := conn.AtomName(atom)
	if err != nil {
		return xproto.GetPropertyReply{}, err
	}
	reply, err := xproto.GetProperty(conn.XConn, false, win, atom, xproto.GetPropertyTypeAny, 0, (1<<32)-1)
	if err != nil {
		return xproto.GetPropertyReply{}, fmt.Errorf("error retrieving property %q on window '%d': %w", atomName, win, err)
	} else if reply.Format == 0 {
		return xproto.GetPropertyReply{}, fmt.Errorf("no such property %q on window '%d'", atomName, win)
	}
	return reply, nil
}

func (conn *XPropConn) GetPropName(win xproto.Window, atomName string) (xproto.GetPropertyReply, error) {
	atom, err := conn.Atom(atomName, false)
	if err != nil {
		return xproto.GetPropertyReply{}, err
	}
	reply, err := xproto.GetProperty(conn.XConn, false, win, atom, xproto.GetPropertyTypeAny, 0, (1<<32)-1)
	if err != nil {
		return xproto.GetPropertyReply{}, fmt.Errorf("error retrieving property %q on window '%d': %w", atomName, win, err)
	} else if reply.Format == 0 {
		return xproto.GetPropertyReply{}, fmt.Errorf("no such property %q on window '%d'", atomName, win)
	}
	return reply, nil
}

func (conn *XPropConn) ChangeProp(win xproto.Window, format uint8, prop, typ xproto.Atom, data []byte) error {
	propName, err := conn.AtomName(prop)
	if err != nil {
		return err
	}
	typName, err := conn.AtomName(typ)
	if err != nil {
		return err
	}
	if err := xproto.ChangeProperty(conn.XConn, xproto.PropModeReplace, win, prop, typ, format, uint32(len(data)/(int(format)/8)), data); err != nil {
		return fmt.Errorf("error changing property %q of type %q on window '%d': %w", propName, typName, win, err)
	}
	return nil
}

func (conn *XPropConn) ChangePropName(win xproto.Window, format uint8, propName, typName string, data []byte) error {
	prop, err := conn.Atom(propName, false)
	if err != nil {
		return err
	}
	typ, err := conn.Atom(typName, false)
	if err != nil {
		return err
	}
	if err := xproto.ChangeProperty(conn.XConn, xproto.PropModeReplace, win, prop, typ, format, uint32(len(data)/(int(format)/8)), data); err != nil {
		return fmt.Errorf("error changing property %q of type %q on window '%d': %w", propName, typName, win, err)
	}
	return nil
}

func FromData32(data32 []uint32) []byte {
	data8 := make([]uint8, len(data32)*4)
	for i, datum := range data32 {
		binary.LittleEndian.PutUint32(data8[i*4:], datum)
	}
	return data8
}
