// Распределение данных из одного канала на несколько
//
//          ┌─ channel 1
//          │
// channel ─┼─ channel 2
//          │
//          └─ channel 3

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	in := fakeJobsGenerator(100)

	fanOut(in, 5, worker)
}

func worker(num int, in <-chan string) {
	for job := range in {
		fmt.Printf("worker %d, %s\n", num, job)
	}
}

func fanOut[T any](in <-chan T, workers int, proc func(num int, in <-chan T)) {
	wg := sync.WaitGroup{}

	for num := range workers {
		wg.Add(1)
		num := num
		go func() {
			defer wg.Done()
			proc(num, in)
		}()
	}

	wg.Wait()
}

func fakeJobsGenerator(countTasks int) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)

		for i := range countTasks {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			out <- fmt.Sprintf("job %d", i)
		}
	}()

	return out
}
