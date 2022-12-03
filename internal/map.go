package internal

import (
	"sync"
)

// Map provides a wrapper around a sync.Map with type casting.
type Map[Key comparable, Value any] struct{ m sync.Map }

// Get ...
func (m *Map[Key, Value]) Get(key Key) (v Value, ok bool) {
	i, ok := m.m.Load(key)
	if !ok {
		return v, false
	}
	return i.(Value), true
}

// Set ...
func (m *Map[Key, Value]) Set(key Key, value Value) (ok bool) {
	_, ok = m.m.LoadOrStore(key, value)
	return !ok
}
