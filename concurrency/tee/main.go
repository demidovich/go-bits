// Разветвитель
// Реплицирование данных из одного канала в несколько
//
//            ┌─ channel A
// channel A ─┤
//            └─ channel A

package main

import (
	"fmt"
	"sync"
)

func main() {
	in := fakeJobsGenerator(10)
	out := Tee(in, 2)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range out[0] {
			fmt.Printf("out 1: %s\n", v)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range out[1] {
			fmt.Printf("out 2: %s\n", v)
		}
	}()

	wg.Wait()
}

func Tee[T any](in <-chan T, num int) []<-chan T {
	out := make([]chan T, num)
	for i := range out {
		out[i] = make(chan T)
	}

	go func() {
		defer func() {
			for _, ch := range out {
				close(ch)
			}
		}()

		for value := range in {
			for _, ch := range out {
				ch <- value // Блокирующий вызов
			}
		}
	}()

	outRO := make([]<-chan T, num)
	for i := range out {
		outRO[i] = out[i]
	}

	return outRO
}

func fakeJobsGenerator(count int) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)
		for i := range count {
			out <- fmt.Sprintf("%d", i)
		}
	}()

	return out
}
