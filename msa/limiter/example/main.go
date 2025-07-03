package main

import (
	"context"
	"fmt"
	"go-bits/msa/limiter"
	"time"
)

func main() {
	c := limiter.Config{
		Interval: 1 * time.Second,
		Tokens:   2,
	}

	l, _ := limiter.New(context.Background(), c)

	fmt.Println(l.Allow())
	fmt.Println(l.Allow())
	fmt.Println(l.Allow())
}
