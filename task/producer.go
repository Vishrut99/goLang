package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func produce(producer chan<- event, i int) {
	fmt.Println("produce", i)
	producer <- event{i, i}
}

func work(producer chan<- event, ctx context.Context, c context.CancelFunc, wg *sync.WaitGroup) {
	defer wg.Done()
	defer c()
	for i := range 1000 {
		select {
		case <-ctx.Done():
			fmt.Println("Context done, stopping work")
			return
		default:
			produce(producer, i)
		}
	}

	time.Sleep(time.Second)
}
