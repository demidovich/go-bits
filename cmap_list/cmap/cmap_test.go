package cmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetGet(t *testing.T) {
	m := NewCmap()
	m.Set("k1", 10)
	value, ok := m.Get("k1")

	assert.True(t, ok)
	assert.Equal(t, 10, int(value))
}

func TestOverwrite(t *testing.T) {
	m := NewCmap()
	m.Set("k1", 10)
	m.Set("k1", 20)
	value, ok := m.Get("k1")

	assert.True(t, ok)
	assert.Equal(t, 20, int(value))
}

func TestGetDeepItem(t *testing.T) {
	m := NewCmap()
	m.Set("k1", 10)
	m.Set("k2", 20)
	m.Set("k3", 30)
	m.Set("k4", 40)
	assert.Equal(t, 1, m.Size())

	value, _ := m.Get("k4")
	assert.Equal(t, 40, int(value))
}

func TestRebalance(t *testing.T) {
	m := NewCmap()

	m.Set("k1", 1)
	m.Set("k2", 2)
	m.Set("k3", 3)
	m.Set("k4", 4)
	m.Set("k5", 5)
	// m.Debug()

	v1, ok1 := m.Get("k1")
	v2, ok2 := m.Get("k2")
	v3, ok3 := m.Get("k3")
	v4, ok4 := m.Get("k4")
	v5, ok5 := m.Get("k5")

	assert.Equal(t, 2, m.Size())

	assert.True(t, ok1)
	assert.Equal(t, 1, int(v1))

	assert.True(t, ok2)
	assert.Equal(t, 2, int(v2))

	assert.True(t, ok3)
	assert.Equal(t, 3, int(v3))

	assert.True(t, ok4)
	assert.Equal(t, 4, int(v4))

	assert.True(t, ok5)
	assert.Equal(t, 5, int(v5))
}
