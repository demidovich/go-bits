package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetGet(t *testing.T) {
	cache := NewLRU(10)
	cache.Set("1", "a")

	value, ok := cache.Get("1")

	assert.True(t, ok)
	assert.Equal(t, "a", value)
}

func TestUpdate(t *testing.T) {
	cache := NewLRU(10)
	cache.Set("1", "a")
	cache.Set("1", "b")

	value, ok := cache.Get("1")

	assert.True(t, ok)
	assert.Equal(t, "b", value)
}

func TestRemove(t *testing.T) {
	cache := NewLRU(10)
	cache.Set("1", "a")
	cache.Remove("1")

	_, ok := cache.Get("1")
	assert.False(t, ok)
}

func TestCapacityExceeding(t *testing.T) {
	cache := NewLRU(1)
	cache.Set("1", "a")
	cache.Set("2", "b")

	value1, ok1 := cache.Get("1")
	assert.False(t, ok1)
	assert.Equal(t, "", value1)

	value2, ok2 := cache.Get("2")
	assert.True(t, ok2)
	assert.Equal(t, "b", value2)
}

func TestHotElement(t *testing.T) {
	cache := NewLRU(2)
	cache.Set("1", "a")
	cache.Set("2", "b")

	_, _ = cache.Get("1")
	cache.Set("3", "c")

	_, ok1 := cache.Get("1")
	assert.True(t, ok1)

	_, ok3 := cache.Get("3")
	assert.True(t, ok3)

	_, ok2 := cache.Get("2")
	assert.False(t, ok2)
}
