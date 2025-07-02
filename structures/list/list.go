package list

type list[T any] struct {
	head, tail *Node[T]
}

type Node[T any] struct {
	prev, next *Node[T]
	key        string
	val        T
}

func (n *Node[_]) Key() string {
	return n.key
}

func (n *Node[T]) Value() T {
	return n.val
}

func (n *Node[T]) Prev() (*Node[T], bool) {
	if n.prev.prev == nil {
		return nil, false
	}

	return n.prev, true
}

func (n *Node[T]) Next() (*Node[T], bool) {
	if n.next.next == nil {
		return nil, false
	}

	return n.next, true
}

func New[T any]() *list[T] {
	head, tail := &Node[T]{}, &Node[T]{}
	head.next = tail
	tail.prev = head

	return &list[T]{
		head: head,
		tail: tail,
	}
}

func (l *list[T]) PushHead(val T) *Node[T] {
	node := &Node[T]{val: val}
	curr := l.head.next

	node.prev = l.head
	node.next = curr
	curr.prev = node

	l.head.next = node

	return node
}

func (l *list[T]) PushTail(val T) *Node[T] {
	node := &Node[T]{val: val}
	curr := l.tail.prev

	node.prev = curr
	node.next = l.tail
	curr.next = node

	l.tail.prev = node

	return node
}

func (l *list[T]) MoveToHead(n *Node[T]) {
	if n == nil {
		return
	}

	if n.prev == l.head {
		return
	}

	// Необходимо изменить связи
	// 1 Место, откуда вынимается узел n
	// 2 Связь prev текущего первого элемента с новым узлом
	// 3 Связь head списка с новым узлом
	// 4 Связи нового первого элемента

	n.prev.next, n.next.prev = n.next, n.prev

	curr := l.head.next
	curr.prev = n

	l.head.next = n

	n.prev = l.head
	n.next = curr
}

func (l *list[T]) MoveToTail(n *Node[T]) {
	if n == nil {
		return
	}

	if n.next == l.tail {
		return
	}

	// Необходимо изменить связи
	// 1 Место, откуда вынимается узел n
	// 2 Связь next текущего последнего элемента с новым узлом
	// 3 Связь tail списка с новым узлом
	// 4 Связи нового последнего элемента

	n.prev.next, n.next.prev = n.next, n.prev

	curr := l.tail.prev
	curr.next = n
	l.tail.prev = n

	n.prev = curr
	n.next = l.tail
}

func (l *list[T]) Head() (*Node[T], bool) {
	return l.head.Next()
}

func (l *list[T]) Tail() (*Node[T], bool) {
	return l.tail.Prev()
}
