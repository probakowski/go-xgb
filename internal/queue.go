package internal

import (
	"sync"
)

type Queue[Elem any] struct {
	ll []Elem
	mu sync.Mutex
}

func (q *Queue[Elem]) Pop() (e Elem, ok bool) {
	q.mu.Lock()
	if ok = len(q.ll) > 0; ok {
		e = q.ll[0]
		q.ll = q.ll[1:]
	}
	q.mu.Unlock()
	return
}

func (q *Queue[Elem]) Push(elem Elem) {
	q.mu.Lock()
	q.ll = append(q.ll, elem)
	q.mu.Unlock()
}

func (q *Queue[Elem]) Len() int {
	q.mu.Lock()
	l := len(q.ll)
	q.mu.Unlock()
	return l
}
