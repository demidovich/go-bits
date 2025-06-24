package main

import (
	"fmt"
	"go-bits/structures/list/list"
)

func main() {
	l := list.NewLinked()
	l.AddToTail(111)
	l.AddToTail(222)
	l.AddToTail(333)

	elem := l.Head
	for {
		if elem == nil {
			break
		} else {
			fmt.Println(elem.Value)
		}
		elem = elem.Next
	}
}
