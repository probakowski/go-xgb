package xgb

import (
	"errors"
	"sync"
)

type xidGenerator struct {
	base uint32
	last uint32
	inc  uint32
	max  uint32
	mu   sync.Mutex
}

func (gen *xidGenerator) Next() (uint32, error) {
	gen.mu.Lock()

	// check if space left in uint32
	if gen.last > (gen.max - gen.inc) {
		gen.mu.Unlock()
		return 0, errors.New("no more available resource identifiers")
	}

	// generate next value
	gen.last += gen.inc
	next := (gen.last | gen.base)

	// unlock and return
	gen.mu.Unlock()
	return next, nil
}
