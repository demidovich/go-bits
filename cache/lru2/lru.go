package cache

import (
	"container/list"
	"sync"
)

type lru[T any] struct {
	mu   *sync.RWMutex
	cap  int
	data map[string]dataItem[T]
	list *list.List
}

type dataItem[T any] struct {
	value    T
	listElem *list.Element
}

func NewLRU[T any](capacity int) *lru[T] {
	if capacity < 1 {
		panic("cache capacity must be greater than zero")
	}

	return &lru[T]{
		mu:   &sync.RWMutex{},
		cap:  capacity,
		data: make(map[string]dataItem[T]),
		list: list.New(),
	}
}

func (c *lru[T]) Get(key string) (T, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if item, ok := c.data[key]; ok {
		c.list.MoveToFront(item.listElem)
		return item.value, true
	} else {
		return item.value, false
	}
}

func (c *lru[T]) Set(key string, val T) {
	c.mu.Lock()
	defer c.mu.Unlock()

	item, ok := c.data[key]
	if ok {
		item.value = val
		c.list.MoveToFront(item.listElem)
	} else {
		item = dataItem[T]{
			value:    val,
			listElem: c.list.PushFront(key),
		}
		c.releaseLocked()
	}

	c.data[key] = item
}

func (c *lru[T]) releaseLocked() {
	for len(c.data) >= c.cap {
		tail := c.list.Back()
		key := tail.Value.(string)

		delete(c.data, key)
		c.list.Remove(tail)
	}
}
