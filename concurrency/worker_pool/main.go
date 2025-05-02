// Вы разрабатываете сервис, который работает с изображениями
// Каждое изображение проходит дорогостоящую обработку, например наложение водяного знака
// Поскольку обработка каждого изображения занимает продолжительное время,
// необходимо обрабатывать их параллельно, чтобы ускорить процесс
// Однака, чтобы уменьшить нагрузку на систему вы хотите уменьшить количество
// одновременно работающих горутин

// Версия с выводом результатов работы воркеров в канал стуктурой Result
// Result может содержать как Success так и Failed

package main

import (
	"fmt"
	"go-bits/concurrency/worker_pool/pool"
	"strconv"
	"time"

	"github.com/google/uuid"
)

const (
	workersCount = 10
	tasksCount   = 100
	taskTTL      = 1 * time.Second
)

func main() {

	tasks := make(chan pool.Task)
	results := pool.Start(workersCount, taskTTL, tasks)

	// listen results
	go func() {
		for result := range results {
			printResult(result)
		}
	}()

	// fake task generator
	for i := range tasksCount {
		task := pool.Task{
			Id:       uuid.New(),
			Filepath: strconv.Itoa(i),
		}
		tasks <- task
	}

	close(tasks)
}

func printResult(result pool.Result) {
	var message string
	if result.Err == nil {
		message = "success"
	} else {
		message = result.Err.Error()
	}

	fmt.Printf(
		"worker: %d, task: %s, file: %s, result: %s\n",
		result.WorkerNum,
		result.TaskId,
		result.Filepath,
		message,
	)
}
