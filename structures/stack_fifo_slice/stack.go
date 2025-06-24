package stack

import "errors"

type Stack[T any] struct {
	items []T
}

func New[T any](size int) Stack[T] {
	return Stack[T]{
		items: make([]T, 0, size),
	}
}

func NewDefault[T any]() Stack[T] {
	return New[T](32)
}

func (s *Stack[T]) Put(v T) {
	s.items = append(s.items, v)
}

func (s *Stack[T]) Pop() (T, error) {
	if len(s.items) == 0 {
		return *new(T), errors.New("stack is empty")
	}

	v := s.items[0]
	s.items = s.items[1:]

	return v, nil
}

func (s *Stack[T]) Peek() (T, error) {
	if len(s.items) == 0 {
		return *new(T), errors.New("stack is empty")
	}

	return s.items[0], nil
}

func (s *Stack[_]) Size() int {
	return len(s.items)
}
