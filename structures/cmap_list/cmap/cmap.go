/*
Custom Map имени 21-го съезда велосипедостроителей
*/
package cmap

import (
	"fmt"
	"hash/fnv"
)

type Cmap struct {
	size          int
	listSizeLimit int
	items         []*cmapItem
	listSize      int
}

type cmapItem struct {
	key   string
	value CmapValue
	next  *cmapItem
}

type CmapValue int

type CmapConfig struct {
	InitSize      uint
	ListSizeLimit uint
}

func New() Cmap {
	return Cmap{
		size:          1,
		listSizeLimit: 4,
		items:         make([]*cmapItem, 1),
	}
}

func NewWithConfig(c CmapConfig) Cmap {
	if c.InitSize < 1 {
		panic("cmap InitSize less than 1")
	}
	if c.ListSizeLimit < 1 {
		panic("cmap ListSizeLimit less than 1")
	}

	return Cmap{
		size:          int(c.InitSize),
		listSizeLimit: int(c.ListSizeLimit),
		items:         make([]*cmapItem, int(c.InitSize)),
	}
}

func (m *Cmap) Size() int {
	return m.size
}

func (m *Cmap) Set(key string, value CmapValue) {
	m.rebalance()

	newItem := &cmapItem{
		key:   key,
		value: value,
	}

	i := itemIndex(m.size, key)
	item := m.items[i]
	listSize := 1

	if item == nil {
		m.items[i] = newItem
	} else {
		for {
			if item.key == key {
				item.value = value
				break
			}
			listSize++
			if item.next == nil {
				item.next = newItem
				break
			}
			item = item.next
		}
	}

	if listSize > m.listSize {
		m.listSize = listSize
	}
}

func (m *Cmap) Get(key string) (value CmapValue, ok bool) {
	i := itemIndex(m.size, key)

	item := m.items[i]
	for item != nil {
		if item.key == key {
			return item.value, true
		}
		item = item.next
	}

	return 0, false
}

func itemIndex(size int, key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))

	return int(h.Sum32()) % size
}

// rebalance

func (m *Cmap) rebalance() {
	if m.listSize < m.listSizeLimit {
		return
	}

	var newSize int
	switch true {
	case m.size > 5120:
		newSize = int(float64(m.size) * 1.5)
	case m.size > 10240:
		newSize = int(float64(m.size) * 1.2)
	default:
		newSize = m.size * 2
	}

	newItems := make([]*cmapItem, newSize)
	for _, item := range m.items {
		// Элемент может быть пустым, если для него не было хэша
		// При этом значение количества элементов списка может быть превышено
		if item == nil {
			continue
		}

		// Разматываем список с head
		// При релокации элемента будет удалена его связь со следующим элементом
		// next1 нужен для этого
		next := item.next
		relocateItem(newSize, newItems, item)
		for next != nil {
			next1 := next.next
			relocateItem(newSize, newItems, next)
			next = next1
		}
	}

	m.size = newSize
	m.items = newItems
	m.listSize = listSize(newItems)
}

func relocateItem(size int, newItems []*cmapItem, item *cmapItem) {
	item.next = nil
	index := itemIndex(size, item.key)

	if newItems[index] == nil {
		newItems[index] = item
		return
	}

	tail := newItems[index]
	for tail.next != nil {
		tail = tail.next
	}

	tail.next = item
}

func listSize(items []*cmapItem) int {
	maxSize := 0

	for _, item := range items {
		if item == nil {
			continue
		}

		size := 1
		next := item.next
		for next != nil {
			size++
			next = next.next
		}

		if size > maxSize {
			maxSize = size
		}
	}

	return maxSize
}

func (m *Cmap) Debug() {
	fmt.Println()
	fmt.Printf("List Pointer: %p\n", m.items)
	fmt.Printf("List Size: %d\n", m.listSize)
	fmt.Printf("List Size Real: %d\n", listSize(m.items))
	fmt.Println("Size:", len(m.items))

	fmt.Println("Items:")
	for key, item := range m.items {
		fmt.Println(key)

		if item == nil {
			fmt.Println("nil")
			continue
		}

		fmt.Println("    ", "key:", item.key, ", value:", item.value, ", next:", item.next)
		next := item.next
		for next != nil {
			fmt.Println("    ", "key:", next.key, ", value:", next.value, ", next:", next.next)
			next = next.next
		}
	}
}
