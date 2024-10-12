package main

import (
	"fmt"
	"sync"
	"time"
)

// Ответ запроса готовится несколькими горутинами
// При этом они могут ходить в одну и туже ручку API, запрос которой очень дорогой
// Необходимо сделать так, чтобы в рамках обработки одного запроса
// выполнялся только один запрос дорогой ручки
func main() {

	request := singleRequest{
		mu: &sync.Mutex{},
	}

	go func() {
		result := request.Result(1)
		fmt.Println(result)
	}()

	go func() {
		result := request.Result(2)
		fmt.Println(result)
	}()

	go func() {
		result := request.Result(3)
		fmt.Println(result)
	}()

	time.Sleep(time.Millisecond)
}

type singleRequest struct {
	mu       *sync.Mutex
	result   string
	complete bool
}

func (s *singleRequest) Result(num int) string {
	s.mu.Lock()
	if !s.complete {
		s.result = fmt.Sprintf("request %d", num)
		s.complete = true
	}
	s.mu.Unlock()

	return s.result
}
