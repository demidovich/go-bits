package cache

import (
	"container/list"
	"sync"
)

type lru struct {
	capacity int
	keys     map[string]*list.Element
	list     *list.List
	mu       sync.Mutex
}

type lruElement struct {
	key   string
	value string
}

func NewLRU(capacity int) *lru {
	if capacity < 1 {
		panic("cache capacity must be greater than zero")
	}

	return &lru{
		capacity: capacity,
		keys:     map[string]*list.Element{},
		list:     list.New(),
	}
}

func (c *lru) Get(key string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	listNode, ok := c.keys[key]
	if !ok {
		return "", false
	}

	element, ok := listNode.Value.(lruElement)
	if !ok {
		return "", false
	}

	c.list.MoveToFront(listNode)

	return element.value, true
}

func (c *lru) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.releaseLocked()

	listNode, ok := c.keys[key]
	if !ok {
		element := lruElement{
			key:   key,
			value: value,
		}
		listNode = c.list.PushFront(element)
		c.keys[key] = listNode
		return
	}

	element, ok := listNode.Value.(lruElement)
	if !ok {
		element = lruElement{
			key:   key,
			value: value,
		}
	} else {
		element.value = value
	}

	listNode.Value = element
}

func (c *lru) Remove(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	listNode, ok := c.keys[key]
	if !ok {
		return
	}

	delete(c.keys, key)

	if listNode != nil {
		c.list.Remove(listNode)
	}
}

func (c *lru) Size() int {
	return len(c.keys)
}

func (c *lru) releaseLocked() {
	for len(c.keys) >= c.capacity {
		listTail := c.list.Back()
		if listTail == nil {
			return
		}

		lruElement, ok := listTail.Value.(lruElement)
		if !ok {
			continue
		}

		c.list.Remove(listTail)
		delete(c.keys, lruElement.key)
	}
}
