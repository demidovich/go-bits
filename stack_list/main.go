package main

import (
	"fmt"
	"structures/stack_list/stack"
)

func main() {
	//LifoStackExample()
	FifoStackExample()
}

func LifoStackExample() {
	s := stack.NewLifo()
	s.Put(111)
	s.Put(222)
	s.Put(333)

	for {
		item, ok := s.Pop()
		if !ok {
			break
		}
		fmt.Println(item.Value)
	}
}

func FifoStackExample() {
	s := stack.NewFifo()
	s.Put(111)
	s.Put(222)
	s.Put(333)

	for {
		item, ok := s.Pop()
		if !ok {
			break
		}
		fmt.Println(item.Value)
	}
}
