package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	w1 := worker(1, 10)
	w2 := worker(2, 10)
	w3 := worker(3, 10)

	out := merge(w1, w2, w3)

	for v := range out {
		fmt.Println(v)
	}
}

func merge(in ...<-chan string) <-chan string {
	out := make(chan string)
	wg := sync.WaitGroup{}

	for _, ch := range in {
		ch := ch
		if ch == nil {
			continue
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range ch {
				out <- v
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func worker(num, tasks int) <-chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)
		for i := range tasks {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			ch <- fmt.Sprintf("worker %d, job %d", num, i)
		}
	}()

	return ch
}
