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

// func TestRebalance(t *testing.T) {
// 	m := NewCmap()
// 	m.Set("k1", 1)
// 	assert.Equal(t, 1, m.Size())

// 	m.Set("k2", 1)
// 	m.Set("k3", 1)
// 	m.Set("k4", 1)
// 	m.Set("k5", 1)
// 	assert.Equal(t, 2, m.Size())
// }
