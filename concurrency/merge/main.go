package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch1 := producer(1, time.Millisecond*50)
	ch2 := producer(2, time.Millisecond*70)
	ch3 := producer(3, time.Millisecond*90)

	out := merge(ch1, ch2, ch3)
	for val := range out {
		fmt.Println(val)
	}
}

func merge(in ...<-chan string) <-chan string {
	out := make(chan string)
	wg := sync.WaitGroup{}

	for _, ch := range in {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for val := range ch {
				out <- val
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func producer(num int, latency time.Duration) <-chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)
		for i := 0; i < 10; i++ {
			time.Sleep(latency)
			ch <- fmt.Sprintf("producer %d, value %d", num, i)
		}
	}()

	return ch
}
