package stack

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	s := NewStack()

	for i := 1; i <= 100; i++ {
		s.Put(strconv.Itoa(i))
	}

	assert.Equal(t, 100, s.Size())

	for i := 1; i <= 100; i++ {
		v, err := s.Pop()
		assert.Equal(t, strconv.Itoa(i), v)
		assert.Nil(t, err)
	}

	assert.Equal(t, 0, s.Size())
}
