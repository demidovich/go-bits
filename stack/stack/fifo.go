package stack

import (
	"sync"
)

// New FIFO stack
func NewFifo() fifo {
	return fifo{}
}

type fifo struct {
	mu   sync.Mutex
	head *FifoItem
	tail *FifoItem
}

// Put element to stack
func (stack *fifo) Put(value int) {
	stack.mu.Lock()
	defer stack.mu.Unlock()

	item := FifoItem{Value: value}

	if stack.head == nil {
		stack.head = &item
		stack.tail = &item
		return
	}

	stack.tail.Next = &item
	stack.tail = &item
}

// Fetch element from stack
//
// ok - indicates that the result is not empty
func (stack *fifo) Fetch() (item FifoItem, ok bool) {
	stack.mu.Lock()
	defer stack.mu.Unlock()

	if stack.head == nil {
		return
	}

	item = *stack.head
	ok = true

	if item.Next != nil {
		stack.head = item.Next
	} else {
		stack.head = nil
	}

	return
}

type FifoItem struct {
	Value int
	Next  *FifoItem
}
