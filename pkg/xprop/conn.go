package xprop

import (
	"codeberg.org/gruf/go-xgb"
	"codeberg.org/gruf/go-xgb/internal"
	"codeberg.org/gruf/go-xgb/xproto"
)

type XPropConn struct {
	atoms *internal.Map[string, xproto.Atom]
	names *internal.Map[xproto.Atom, string]
	xconn *xgb.XConn
}

func NewConn(xconn *xgb.XConn) *XPropConn {
	return &XPropConn{
		atoms: internal.NewMap[string, xproto.Atom](),
		names: internal.NewMap[xproto.Atom, string](),
		xconn: xconn,
	}
}
