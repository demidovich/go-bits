/*
Custom Map имени 20-го съезда велосипедостроителей
*/
package cmap

import (
	"strconv"
)

const INIT_SIZE = 1
const MAX_LIST_SIZE = 4

type Cmap struct {
	size     int
	items    []*cmapItem
	listSize int
}

type cmapItem struct {
	key   string
	value CmapValue
	next  *cmapItem
}

type CmapValue int

func NewCmap() Cmap {
	return Cmap{
		size:  INIT_SIZE,
		items: make([]*cmapItem, INIT_SIZE),
	}
}

func (m *Cmap) Size() int {
	return m.size
}

func (m *Cmap) Set(key string, value CmapValue) {
	m.rebalance()

	newItem := cmapItem{
		key:   key,
		value: value,
	}

	i := itemIndex(m.size, key)
	item := m.items[i]
	listSize := 0
	for {
		listSize++
		if item == nil {
			m.items[i] = &newItem
			break
		}
		if item.key == key {
			item.value = value
			break
		}
		if item.next == nil {
			item.next = &newItem
			break
		}
		item = item.next
	}

	if listSize > m.listSize {
		m.listSize = listSize
	}
}

func (m *Cmap) Get(key string) (value CmapValue, ok bool) {
	item, ok := m.item(key)
	if ok {
		value = item.value
		ok = true
	}
	return
}

func itemIndex(size int, key string) int {
	h, _ := strconv.Atoi(key)
	return h % size
}

func (m *Cmap) item(key string) (*cmapItem, bool) {
	i := itemIndex(m.size, key)

	item := m.items[i]
	for item != nil {
		if item.key == key {
			return item, true
		}
		item = item.next
	}

	return &cmapItem{}, false
}

func (m *Cmap) rebalance() {
	if m.listSize < MAX_LIST_SIZE {
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
	newListSize := 0
	// for _, item := range m.items {
	// 	nIndex := itemIndex(newSize, item.key)
	// 	// Здесь перебалансировка
	// }

	m.size = newSize
	m.items = newItems
	m.listSize = newListSize
}
