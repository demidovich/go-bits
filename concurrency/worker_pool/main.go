// Вы разрабатываете сервис, который работает с изображениями
// Каждое изображение проходит дорогостоящую обработку, например наложение водяного знака
// Поскольку обработка каждого изображения занимает продолжительное время,
// необходимо обрабатывать их параллельно, чтобы ускорить процесс
// Однака, чтобы уменьшить нагрузку на систему вы хотите уменьшить количество
// одновременно работающих горутин

// Версия с выводом результатов работы воркеров в канал string

package main

import (
	"fmt"
	"go-bits/concurrency/worker_pool/pool"
	"strconv"

	"github.com/google/uuid"
)

const (
	workersCount = 10
	tasksCount   = 100
)

func main() {

	tasks := make(chan pool.Task)
	results := pool.Start(workersCount, tasks)

	// listen results
	go func() {
		for msg := range results {
			fmt.Println(msg)
		}
	}()

	// fake task generator
	for i := range tasksCount {
		task := pool.Task{
			Id:   uuid.New(),
			File: strconv.Itoa(i),
		}
		tasks <- task
	}

	close(tasks)
}
