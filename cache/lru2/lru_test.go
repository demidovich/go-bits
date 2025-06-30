package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCache_Get(t *testing.T) {
	c := NewLRU[string](5)

	c.Set("1", "a")
	v, ok := c.Get("1")

	assert.Equal(t, "a", v)
	assert.True(t, ok)
}

func TestCache_Capacity(t *testing.T) {
	c := NewLRU[string](2)

	c.Set("1", "a")
	c.Set("2", "b")
	c.Set("3", "c")

	_, ok1 := c.Get("1")
	_, ok2 := c.Get("2")
	_, ok3 := c.Get("3")

	assert.False(t, ok1)
	assert.True(t, ok2)
	assert.True(t, ok3)
}

func TestCache_Hit(t *testing.T) {
	c := NewLRU[string](2)

	c.Set("1", "a")
	c.Set("2", "b")
	_, _ = c.Get("1")
	c.Set("3", "c")

	_, ok1 := c.Get("1")
	_, ok2 := c.Get("2")
	_, ok3 := c.Get("3")

	assert.True(t, ok1)
	assert.False(t, ok2)
	assert.True(t, ok3)
}
