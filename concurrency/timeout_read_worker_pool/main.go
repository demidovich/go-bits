package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	results := request(time.Millisecond * 100)
	for _, v := range results {
		fmt.Println("Accepted:", v)
	}
}

func request(ttl time.Duration) []string {
	results := []string{}

	ch := make(chan string)
	go func() { ch <- worker(1) }()
	go func() { ch <- worker(2) }()
	go func() { ch <- worker(3) }()
	timeout := time.After(ttl)

loop:
	for i := 0; i < 3; i++ {
		select {
		case v := <-ch:
			results = append(results, v)
		case <-timeout:
			break loop
		}
	}

	return results
}

func worker(num int) string {
	t := time.Millisecond * time.Duration(rand.Intn(150))
	time.Sleep(t)
	return fmt.Sprintf("response worker %d, runtime %s", num, t.String())
}
