package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Семафор
// Для обработки N количества задач потребуется N количество горутин
// При большом количестве одновременно запущенных горутин может быть перерасход ресурсов
// и регресс производительности
// Паттерн ограничивает количество одновременно работающих горутин

func main() {
	jobs := fakeJobsGenerator(100)

	jobProcessor := func(job string) {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		fmt.Println(job)
	}

	Semaphore(jobs, 25, jobProcessor)
}

func Semaphore[T any](in <-chan T, limit int, processor func(T)) {
	sem := make(chan struct{}, limit)
	defer close(sem)

	for job := range in {
		sem <- struct{}{}
		go func() {
			defer func() { <-sem }()
			processor(job)
		}()
	}
}

func fakeJobsGenerator(count int) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)
		for i := range count {
			out <- fmt.Sprintf("job %d", i)
		}
	}()

	return out
}
