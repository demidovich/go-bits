/*
Custom Map имени 20-го съезда велосипедостроителей
*/
package cmap

import (
	"strconv"
)

const INIT_CMAP_SIZE = 1

type Cmap struct {
	size    int
	buckets [][]item
}

type item struct {
	key   string
	value CmapValue
}

type CmapValue int

func NewCmap() Cmap {
	return Cmap{
		size:    INIT_CMAP_SIZE,
		buckets: make([][]item, INIT_CMAP_SIZE),
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

func (m *Cmap) Set(key string, value CmapValue) {
	// m.rebalance()

	i := m.bucketIndex(m.size, key)
	for k, item := range m.buckets[i] {
		if item.key == key {
			m.buckets[i][k].value = value
			return
		}
	}

	item := item{
		key:   key,
		value: value,
	}

	m.buckets[i] = append(m.buckets[i], item)
}

func (m *Cmap) bucketIndex(size int, key string) int {
	h, _ := strconv.Atoi(key)
	return h % size
}

func (m *Cmap) item(key string) (*item, bool) {
	i := m.bucketIndex(m.size, key)
	b := m.buckets[i]

	for _, item := range b {
		if item.key == key {
			return &item, true
		}
	}

	return nil, false
}

// func (m *Cmap) rebalance() {
// }
