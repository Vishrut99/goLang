package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1) // Correct place to add to the WaitGroup

		go func(i int) {
			// wg.Add(1) // Added an extra wg.Add(1) here, which is unnecessary and will cause a deadlock.
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}

	wg.Wait()
	fmt.Println("done")
}
