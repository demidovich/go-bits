package main

import (
	"context"
	"fmt"
	"go-bits/msa/throttler"
	"time"
)

func main() {
	tHandler := throttler.Throttle(handler, 3, 3, 1*time.Second)

	ctx := context.Background()
	for i := range 25 {
		data, err := tHandler(ctx)
		fmt.Printf("%d %s %v\n", i, data, err)
		time.Sleep(100 * time.Millisecond)
	}
}

func handler(_ context.Context) (string, error) {
	return "", nil
}
