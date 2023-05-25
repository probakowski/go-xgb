package xgb

import (
	"sync"
)

// xidGenerator generates X resource identifiers
// give X server provided base, max and inc.
type xidGenerator struct {
	base uint32
	last uint32
	inc  uint32
	max  uint32
	mu   sync.Mutex
}

// Next returns the next available resource identifier.
func (gen *xidGenerator) Next() uint32 {
	// Acquire lock.
	gen.mu.Lock()
	defer gen.mu.Unlock()

	// check if space left in uint32
	if gen.last > (gen.max - gen.inc) {
		panic("no more available resource identifiers")
	}

	// generate next value
	gen.last += gen.inc
	next := (gen.last | gen.base)

	return next
}
