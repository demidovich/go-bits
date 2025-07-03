package roundrobin

import (
	"sync/atomic"
)

type rr[T any] struct {
	items  []T
	iCount int64
	last   atomic.Int64
}

func New[T any](items []T) *rr[T] {
	return &rr[T]{
		items:  items,
		iCount: int64(len(items)),
		last:   atomic.Int64{},
	}
}

func (r *rr[T]) Next() T {
	last := r.last.Add(1)
	return r.items[last%r.iCount]
}
