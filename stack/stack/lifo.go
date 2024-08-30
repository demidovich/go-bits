package stack

import "sync"

// New LIFO stack
func NewLifo() lifo {
	return lifo{}
}

type lifo struct {
	mu   sync.Mutex
	tail *LifoItem
}

// Put element to stack
func (stack *lifo) Put(value int) {
	stack.mu.Lock()
	defer stack.mu.Unlock()

	item := LifoItem{Value: value}
	if stack.tail != nil {
		item.Prev = stack.tail
	}
	stack.tail = &item
}

// Fetch element from stack
//
// ok - indicates that the result is not empty
func (stack *lifo) Fetch() (item LifoItem, ok bool) {
	stack.mu.Lock()
	defer stack.mu.Unlock()

	if stack.tail == nil {
		return
	}

	item = *stack.tail
	ok = true

	if stack.tail.Prev == nil {
		stack.tail = nil
	} else {
		stack.tail = stack.tail.Prev
	}

	return
}

type LifoItem struct {
	Value int
	Prev  *LifoItem
}
