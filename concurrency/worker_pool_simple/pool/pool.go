package pool

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id   uuid.UUID
	File string
}

func Start(workersCount int, tasks <-chan Task) <-chan string {
	results := make(chan string)

	wg := sync.WaitGroup{}
	for num := range workersCount {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			startWorker(num, tasks, results)
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
	results chan<- string
}

func startWorker(num int, tasks <-chan Task, results chan<- string) {
	w := worker{
		num:     num,
		results: results,
	}

	for task := range tasks {
		w.processImage(task)
	}
}

func (w *worker) processImage(task Task) {
	time.Sleep(time.Millisecond * 100)
	w.results <- fmt.Sprintf("worker %d, task %s, file %s", w.num, task.Id, task.File)
}
