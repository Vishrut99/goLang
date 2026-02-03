package main

import (
	"fmt"
)

func main() {

	ch := make(chan int)
	go func() {
		for i := 1; i < 6; i++ {
			ch <- i
		}
		close(ch)
	}()

	for range ch {
		fmt.Println(<-ch) 

	}

	// fmt.Println("done")
}
