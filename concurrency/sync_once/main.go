package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Once call")
	}

	done := make(chan bool)
	for range 10 {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}

	for range 10 {
		<-done
	}
}
