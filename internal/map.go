package internal

import (
	"sync/atomic"
	"unsafe"
)

// Map provides a concurrency safe wrapper round a map[Key]Value
// optimized heavily for read speeds. Writes will reallocate the
// existing contents of the map on each ocurrence.
type Map[Key comparable, Value any] struct{ m unsafe.Pointer }

func (m *Map[Key, Value]) Get(key Key) (v Value, ok bool) {
	if p := atomic.LoadPointer(&m.m); p != nil {
		v, ok = (*(*map[Key]Value)(p))[key]
	}
	return
}

func (m *Map[Key, Value]) Set(key Key, value Value) (ok bool) {
	for {
		var mm map[Key]Value

		// Fetch existing map pointer.
		p := atomic.LoadPointer(&m.m)

		if p != nil {
			// Create a clone of existing map.
			mm = clone(*(*map[Key]Value)(p))
		} else {
			// Map not yet allocated.
			mm = make(map[Key]Value)
		}

		// Append value.
		mm[key] = value

		// Create new map pointer.
		p2 := unsafe.Pointer(&mm)

		// Attempt to replace existing map with new.
		if atomic.CompareAndSwapPointer(&m.m, p, p2) {
			return
		}
	}
}

// clone will create and return a clone of provided map 'm'.
func clone[Key comparable, Value any](m map[Key]Value) map[Key]Value {
	m2 := make(map[Key]Value, len(m))
	for k, v := range m {
		m2[k] = v
	}
	return m2
}

// SmallPtrMap is an array of raw (unsafe) pointers of length
// equal to the number of possible values contained in a uint8.
// All read / write operations are performed atomically.
type SmallPtrMap struct {
	arr [256]unsafe.Pointer
}

func (m *SmallPtrMap) Get(n uint8) unsafe.Pointer {
	return atomic.LoadPointer(&m.arr[n])
}

func (m *SmallPtrMap) Set(n uint8, p unsafe.Pointer) bool {
	return atomic.CompareAndSwapPointer(&m.arr[n], nil, p)
}
