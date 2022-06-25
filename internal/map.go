package internal

import (
	"sync"
)

type Map[Key comparable, Value any] struct {
	kvs map[Key]Value
	mu  sync.Mutex
}

func NewMap[Key comparable, Value any]() *Map[Key, Value] {
	return &Map[Key, Value]{kvs: map[Key]Value{}}
}

func (m *Map[Key, Value]) Get(key Key) (v Value, ok bool) {
	m.mu.Lock()
	v, ok = m.kvs[key]
	m.mu.Unlock()
	return
}

func (m *Map[Key, Value]) Set(key Key, value Value) (ok bool) {
	m.mu.Lock()
	if _, ok = m.kvs[key]; !ok {
		m.kvs[key] = value
	}
	m.mu.Unlock()
	return !ok
}
