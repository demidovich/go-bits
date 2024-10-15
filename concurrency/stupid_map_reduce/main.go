package main

import (
	"fmt"
	"sync"
	"time"
)

// Сложение n-элементов с помощью k-воркеров
func main() {
	startTime := time.Now()

	valuesCount := 1000000
	workers := 100

	values := valuesGenerator(valuesCount)

	workerResult := make(chan int)
	wg := sync.WaitGroup{}

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(values, workerResult)
		}()
	}

	go func() {
		wg.Wait()
		close(workerResult)
	}()

	var total int
	for s := range workerResult {
		total += s
	}

	fmt.Println("Runtime:", time.Since(startTime))
	fmt.Println("Result:", total)
}

func worker(values <-chan int, workerResult chan<- int) {
	var sum int
	for i := range values {
		sum += i
	}
	workerResult <- sum
}

func valuesGenerator(valuesCount int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < valuesCount; i++ {
			ch <- i
		}
	}()
	return ch
}
