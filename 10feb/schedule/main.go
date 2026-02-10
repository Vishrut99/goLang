package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	go func() {
		fmt.Println("locking 1")
		mu.Lock()
		fmt.Println("1")
		time.Sleep(5 * time.Second)
		mu.Unlock()
		fmt.Println("unclocking 1")
	}()

	go func() {
		fmt.Println("2")
		mu.Unlock()
		fmt.Println("unclocking 2")
		mu.Lock()
	}()

	select {}
}
