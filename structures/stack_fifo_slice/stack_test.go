package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewInstance(t *testing.T) {
	s := NewDefault[int]()
	assert.Equal(t, 0, s.Size())

	_, err := s.Pop()
	assert.NotNil(t, err)

	_, err = s.Peek()
	assert.NotNil(t, err)
}

func Test_PutOnePopOne(t *testing.T) {
	s := NewDefault[int]()

	s.Put(1)
	assert.Equal(t, 1, s.Size())

	v, err := s.Pop()
	assert.Equal(t, 0, s.Size())
	assert.Equal(t, 1, v)
	assert.Nil(t, err)
}

func Test_PutTwoPopOne(t *testing.T) {
	s := NewDefault[int]()

	s.Put(1)
	s.Put(2)
	assert.Equal(t, 2, s.Size())

	v, err := s.Pop()
	assert.Equal(t, 1, s.Size())
	assert.Equal(t, 1, v)
	assert.Nil(t, err)
}

func Test_Peek(t *testing.T) {
	s := NewDefault[int]()

	s.Put(1)
	s.Put(2)
	v, err := s.Peek()
	assert.Equal(t, 1, v)
	assert.Nil(t, err)
}
