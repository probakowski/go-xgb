package xprop

/*
xprop/atom.go contains functions related to interning atoms and retrieving
atom names from an atom identifier.

It also manages an atom cache so that once an atom is interned from the X
server, all future atom interns use that value. (So that one and only one
request is sent for interning each atom.)
*/

import (
	"fmt"

	"codeberg.org/gruf/go-xgb/xproto"
)

func (conn *XPropConn) Intern(name string, onlyIfExists bool) error {
	_, err := conn.Atom(name, onlyIfExists)
	return err
}

func (conn *XPropConn) Atom(name string, onlyIfExists bool) (xproto.Atom, error) {
	if id, ok := conn.atoms.Get(name); ok {
		return id, nil
	}

	reply, err := xproto.InternAtom(conn.XConn, onlyIfExists, uint16(len(name)), name)
	if err != nil {
		return 0, fmt.Errorf("error interning atom with name %q: %w", name, err)
	} else if reply.Atom == 0 {
		return 0, fmt.Errorf("atom with name %q returned zero id", name)
	}

	conn.atoms.Add(name, reply.Atom)
	conn.names.Add(reply.Atom, name)

	return reply.Atom, nil
}

func (conn *XPropConn) AtomName(id xproto.Atom) (string, error) {
	if name, ok := conn.names.Get(id); ok {
		return name, nil
	}

	reply, err := xproto.GetAtomName(conn.XConn, id)
	if err != nil {
		return "", fmt.Errorf("error fetching atom name for id '%d': %w", id, err)
	} else if reply.Name == "" {
		return "", fmt.Errorf("atom with id '%d' returned empty name", id)
	}

	conn.atoms.Add(reply.Name, id)
	conn.names.Add(id, reply.Name)

	return reply.Name, nil
}
