package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	jobs := fakeJobsGenerator(100)

	jobProcessor := func(v string) {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		fmt.Println(v)
	}

	start := time.Now()
	WorkerPool(jobs, 15, jobProcessor)
	fmt.Printf("%s\n", time.Since(start))
}

func WorkerPool[T any](in <-chan T, size int, processor func(T)) {
	wg := sync.WaitGroup{}

	for range size {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for job := range in {
				processor(job)
			}
		}()
	}

	wg.Wait()
}

func fakeJobsGenerator(num int) <-chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)

		for i := range num {
			ch <- fmt.Sprintf("job %d", i)
		}
	}()

	return ch
}
