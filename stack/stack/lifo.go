package stack

import "sync"

func NewLifo() lifo {
	return lifo{}
}

type lifo struct {
	mu   sync.Mutex
	tail *LifoItem
}

// Put element to stack
func (l *lifo) Put(value int) {
	l.mu.Lock()
	defer l.mu.Unlock()

	item := LifoItem{Value: value}
	if l.tail != nil {
		item.Prev = l.tail
	}
	l.tail = &item
}

// Fetch element from stack
//
// ok - indicates that the result is not empty
func (l *lifo) Fetch() (item LifoItem, ok bool) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.tail == nil {
		return
	}

	item = *l.tail
	ok = true

	if l.tail.Prev == nil {
		l.tail = nil
	} else {
		l.tail = l.tail.Prev
	}

	return item, true
}

type LifoItem struct {
	Value int
	Prev  *LifoItem
}
