package main

import (
	"fmt"
	"structures/stack/stack"
)

func main() {
	s := stack.NewLifo()
	s.Put(111)
	s.Put(222)
	s.Put(333)

	for {
		item, ok := s.Fetch()
		if !ok {
			break
		}
		fmt.Println(item.Value)
	}
}
