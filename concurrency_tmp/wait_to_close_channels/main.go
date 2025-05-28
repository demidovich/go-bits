package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 2)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		WaitToClose2(ch1, ch2)
		wg.Done()
	}()

	ch1 <- 1
	ch2 <- 2

	close(ch1)
	close(ch2)

	wg.Wait()
}

func WaitToClose1(ch1, ch2 chan int) {
	ch1closed, ch2closed := false, false

	for !ch1closed || !ch2closed {
		select {
		case val, ok := <-ch1:
			fmt.Println("ch1:", val)
			if !ok {
				ch1closed = true
			}
		case val, ok := <-ch2:
			fmt.Println("ch2:", val)
			if !ok {
				ch2closed = true
			}
		}
	}
}

func WaitToClose2(ch1, ch2 chan int) {
	for ch1 != nil || ch2 != nil {
		select {
		case val, ok := <-ch1:
			fmt.Println("ch1:", val)
			if !ok {
				ch1 = nil
			}
		case val, ok := <-ch2:
			fmt.Println("ch2:", val)
			if !ok {
				ch2 = nil
			}
		}
	}
}
