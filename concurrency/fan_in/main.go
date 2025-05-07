// Объединение данных из нескольких каналов в один

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	ch1 := fakeJobsGenerator(1, 10)
	ch2 := fakeJobsGenerator(2, 10)
	ch3 := fakeJobsGenerator(3, 10)

	out := mergeChannels(ch1, ch2, ch3)

	for v := range out {
		fmt.Println(v)
	}
}

func mergeChannels[T any](in ...<-chan T) <-chan T {
	out := make(chan T)
	wg := sync.WaitGroup{}

	for _, channel := range in {
		if channel == nil {
			continue
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			for value := range channel {
				out <- value
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func fakeJobsGenerator(num, countTasks int) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)
		for i := range countTasks {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			out <- fmt.Sprintf("worker %d, job %d", num, i)
		}
	}()

	return out
}
