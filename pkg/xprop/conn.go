package xprop

import (
	"github.com/probakowski/go-xgb"
	"github.com/probakowski/go-xgb/internal"
	"github.com/probakowski/go-xgb/xproto"
)

// XPropConn wraps an XConn to provide
// atom / property name caching and a
// selection of simple utility methods.
type XPropConn struct {
	// XConn is the underlying
	// X connection interface.
	XConn *xgb.XConn

	// internal caching
	atoms internal.Map[string, xproto.Atom]
	names internal.Map[xproto.Atom, string]
}
