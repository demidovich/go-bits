package stack

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	s := NewLifo()
	s.Put(10)
	s.Put(20)

	item, ok := s.Fetch()
	require.Equal(t, 20, item.Value)
	require.Equal(t, true, ok)

	item, ok = s.Fetch()
	require.Equal(t, 10, item.Value)
	require.Equal(t, true, ok)

	_, ok = s.Fetch()
	require.Equal(t, false, ok)
}
