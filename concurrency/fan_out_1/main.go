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
	splitted := splitChannel(in, 2)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for v := range splitted[0] {
			fmt.Println("goroutine 1:", v)
		}
	}()

	go func() {
		defer wg.Done()
		for v := range splitted[1] {
			fmt.Println("goroutine 2:", v)
		}
	}()

	wg.Wait()
}

func splitChannel[T any](in <-chan T, num int) []<-chan T {
	out := make([]chan T, num)
	for i := range num {
		out[i] = make(chan T)
	}

	go func() {
		defer func() {
			for _, channel := range out {
				close(channel)
			}
		}()

		i := 0
		for value := range in {
			out[i] <- value // Блокирующий вызов
			i = (i + 1) % num
		}
	}()

	outRO := make([]<-chan T, num)
	for i := range out {
		outRO[i] = out[i]
	}

	return outRO
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
