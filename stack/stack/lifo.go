package stack

func NewLifo() lifo {
	return lifo{}
}

type lifo struct {
	tail *LifoItem
}

// Put element to stach
func (l *lifo) Put(value int) {
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
