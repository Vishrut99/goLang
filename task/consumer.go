package main

import (
	"context"
	"fmt"
	"sync"
)

func consume(a event, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Consume --", a.userid)
	process(a)
}

func manager(producer <-chan event, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			for {
				select {
				case a := <-producer:
					fmt.Println("Context is done, but Channel still has data : ", a.userid)
					wg.Add(1)
					go consume(a, wg)
				default:
					fmt.Println("Context is done and Channel is closed")
					return
				}
			}
		case a := <-producer:
			// fmt.Println("Context is done, but Channel still has data : ",a.userid)
			wg.Add(1)
			go consume(a, wg)
		}
	}
}
