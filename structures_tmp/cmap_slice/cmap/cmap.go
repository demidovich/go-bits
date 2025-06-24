/*
Custom Map имени 20-го съезда велосипедостроителей
*/
package cmap

import (
	"hash/fnv"
)

const INIT_BUCKETS = 1
const MAX_BUCKET_SIZE = 4

type Cmap struct {
	size          int
	buckets       [][]item
	maxBucketSize int
}

type item struct {
	key   string
	value CmapValue
}

type CmapValue int

func NewCmap() Cmap {
	return Cmap{
		size:    INIT_BUCKETS,
		buckets: make([][]item, INIT_BUCKETS),
	}
}

func (m *Cmap) Size() int {
	return m.size
}

func (m *Cmap) Set(key string, value CmapValue) {
	m.rebalance()

	i := bucketIndex(m.size, key)
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
	if len(m.buckets[i]) > m.maxBucketSize {
		m.maxBucketSize = len(m.buckets[i])
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

func bucketIndex(size int, key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))

	return int(h.Sum32()) % size
}

func (m *Cmap) item(key string) (*item, bool) {
	i := bucketIndex(m.size, key)
	b := m.buckets[i]

	for _, item := range b {
		if item.key == key {
			return &item, true
		}
	}

	return nil, false
}

func (m *Cmap) rebalance() {
	if m.maxBucketSize < MAX_BUCKET_SIZE {
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

	newBuckets := make([][]item, newSize)
	for _, bucket := range m.buckets {
		for _, item := range bucket {
			nIndex := bucketIndex(newSize, item.key)
			newBuckets[nIndex] = append(newBuckets[nIndex], item)
		}
	}

	var newMaxBucketSize int
	for _, bucket := range newBuckets {
		if len(bucket) > newMaxBucketSize {
			newMaxBucketSize = len(bucket)
		}
	}

	m.size = newSize
	m.buckets = newBuckets
	m.maxBucketSize = newMaxBucketSize
}
