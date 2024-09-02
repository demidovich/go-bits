package list

type linked struct {
	Head *LinkedItem
}

type LinkedItem struct {
	Value int
	Next  *LinkedItem
}

func NewLinked() linked {
	return linked{}
}

func (list *linked) AddToHead(value int) {
	item := LinkedItem{Value: value}

	if list.Head != nil {
		item.Next = list.Head
	}

	list.Head = &item
}

func (list *linked) AddToTail(value int) {
	item := LinkedItem{Value: value}

	tail, ok := list.Tail()
	if !ok {
		list.Head = &item
	} else {
		tail.Next = &item
	}
}

func (list *linked) Tail() (item *LinkedItem, ok bool) {
	if list.Head == nil {
		return
	}

	curr := list.Head
	for {
		if curr.Next == nil {
			break
		}
		curr = curr.Next
	}

	item = curr
	ok = true

	return
}
