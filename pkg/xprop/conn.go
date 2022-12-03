package xprop

import (
	"codeberg.org/gruf/go-xgb"
	"codeberg.org/gruf/go-xgb/internal"
	"codeberg.org/gruf/go-xgb/xproto"
)

// XPropConn ...
type XPropConn struct {
	// XConn ...
	XConn *xgb.XConn

	// internal caching
	atoms internal.Map[string, xproto.Atom]
	names internal.Map[xproto.Atom, string]
}
