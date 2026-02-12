package main

import (
	"fmt"
)

func main() {

	done := make(chan bool)
	// startTime := time.Now()
	// defer fmt.Println(time.Since(startTime))
	go func() {
		fmt.Println(":::")
		// done <- 1 == 1
	}()

	// close(done)
	<-done
}
