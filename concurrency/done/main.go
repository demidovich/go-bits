package main

import (
	"fmt"
	"time"
)

// Ожидание сообщения в канале,
// сигнализирующем о завершении операций в горутинах
func main() {
	ch := make(chan string)
	done := make(chan struct{})

	go func() {
		time.Sleep(time.Millisecond * 500)
		ch <- "ping"
	}()

	go func() {
		in := <-ch
		fmt.Println(in)
		done <- struct{}{}
	}()

	<-done
	fmt.Println("Complete")
}
