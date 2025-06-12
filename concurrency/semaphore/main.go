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
	in := fakeJobsGenerator(100)

	sem := make(chan struct{}, 10) // Пропускаем 10 задач
	for job := range in {
		sem <- struct{}{}
		go func(j string) {
			defer func() { <-sem }()
			// Обработка задачи
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println(j)
		}(job)
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
