package stack

import "errors"

type Stack struct {
	items []string
}

func NewStack() Stack {
	return Stack{}
}

func (s *Stack) Put(v string) {
	s.items = append(s.items, v)
}

func (s *Stack) Pop() (string, error) {
	if len(s.items) == 0 {
		return "", errors.New("stack is empty")
	}

	v := s.items[0]
	s.items = s.items[1:]

	return v, nil
}

func (s *Stack) Size() int {
	return len(s.items)
}
