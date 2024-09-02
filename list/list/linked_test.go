package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddToHead(t *testing.T) {
	l := NewLinked()
	l.AddToHead(10)
	l.AddToHead(20)
	l.AddToHead(30)

	assert.Equal(t, 30, l.Head.Value)
	assert.Equal(t, 20, l.Head.Next.Value)
	assert.Equal(t, 10, l.Head.Next.Next.Value)
}

func TestAddToTail(t *testing.T) {
	l := NewLinked()
	l.AddToTail(10)
	l.AddToTail(20)
	l.AddToTail(30)

	assert.Equal(t, 10, l.Head.Value)
	assert.Equal(t, 20, l.Head.Next.Value)
	assert.Equal(t, 30, l.Head.Next.Next.Value)
}
