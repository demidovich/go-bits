package stack

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	s := NewStack()

	for i := range 100 {
		s.Put(strconv.Itoa(i))
	}

	assert.Equal(t, 100, s.Size())

	for i := range 100 {
		v, err := s.Pop()
		assert.Equal(t, strconv.Itoa(i), v)
		assert.Nil(t, err)
	}

	assert.Equal(t, 0, s.Size())
}
