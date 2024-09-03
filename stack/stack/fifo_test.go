package stack

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStackFifo(t *testing.T) {
	s := NewFifo()
	s.Put(10)
	s.Put(20)

	item, ok := s.Pop()
	require.Equal(t, 10, item.Value)
	require.Equal(t, true, ok)

	item, ok = s.Pop()
	require.Equal(t, 20, item.Value)
	require.Equal(t, true, ok)

	_, ok = s.Pop()
	require.Equal(t, false, ok)
}
