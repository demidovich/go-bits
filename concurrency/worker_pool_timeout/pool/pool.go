package pool

import (
	"math/rand"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id       uuid.UUID
	Filepath string
}

func Start(workersCount int, taskTimeout time.Duration, tasks <-chan Task) <-chan Result {
	results := make(chan Result)

	wg := sync.WaitGroup{}
	for num := range workersCount {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			startWorker(
				num,
				taskTimeout,
				tasks,
				results,
			)
		}(num)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	return results
}

type worker struct {
	num     int
	results chan<- Result
}

func startWorker(num int, taskTimeout time.Duration, tasks <-chan Task, results chan<- Result) {
	w := worker{
		num:     num,
		results: results,
	}

	for task := range tasks {
		w.processImage(task)
	}
}

func (w *worker) processImage(task Task) {
	runtime := rand.Intn(3000-500) + 500
	time.Sleep(time.Millisecond * time.Duration(runtime))

	w.results <- Result{
		WorkerNum: w.num,
		TaskId:    task.Id.String(),
		Filepath:  task.Filepath,
	}
}

type Result struct {
	WorkerNum int
	TaskId    string
	Filepath  string
	Err       error
}
