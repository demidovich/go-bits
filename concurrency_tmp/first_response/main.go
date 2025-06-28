package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	servers := []string{
		"s1",
		"s2",
		"s3",
		"s4",
		"s5",
	}

	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	for _, s := range servers {
		wg.Add(1)
		go func() {
			defer wg.Done()

			fmt.Println("request " + s)
			select {
			case res := <-request(s):
				fmt.Println(res)
				cancel()
				return
			case <-ctx.Done():
				return
			}
		}()
	}

	wg.Wait()
}

func request(server string) <-chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		ch <- server + " response"
	}()

	return ch
}
