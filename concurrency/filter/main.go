// Фильтр данных канала
// Прикидывается таким же каналом, но с отфильтрованными данными
//
// channel -> filter -> channel

package main

import "fmt"

func main() {
	jobs := fakeJobsGenerator(10)

	// Обработка канала без фильтрации

	// for v := range jobs {
	// 	fmt.Println(v)
	// }

	// Обработка канала с только четными значениями

	var evenFilter = func(v int) bool {
		return v%2 == 0
	}

	for v := range Filtered(jobs, evenFilter) {
		fmt.Println(v)
	}
}

func Filtered[T any](in <-chan T, filter func(T) bool) <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)
		for v := range in {
			if filter(v) {
				out <- v
			}
		}
	}()

	return out
}

func fakeJobsGenerator(count int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for i := range count {
			out <- i
		}
	}()

	return out
}
