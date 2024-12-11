package main

import (
	"fmt"
	"time"
)

func main() {
	for message := range readMessages(time.Second * 5) {
		fmt.Println(message)
	}
}

func readMessages(timeout time.Duration) <-chan string {
	out := make(chan string)

	go func(out chan<- string) {
		for i := 0; i < 1000; i++ {
			time.Sleep(time.Millisecond * 100)
			out <- fmt.Sprintf("какие-то данные %d", i)
		}
	}(out)

	go func(out chan string, timeout time.Duration) {
		<-time.After(timeout)
		out <- "закрытие канала"
		close(out)
	}(out, timeout)

	return out
}
