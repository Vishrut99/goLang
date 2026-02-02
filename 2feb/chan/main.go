package main

import (
	"fmt"
	"time"
)

func unbufferedDemo() {
	ch := make(chan int) // unbuffered

	go func() {
		fmt.Println("Unbuffered: sender waiting...")
		ch <- 10 // BLOCKS until receiver is ready
		fmt.Println("Unbuffered: sender sent value")
	}()

	time.Sleep(500 * time.Millisecond) // receiver not ready yet

	fmt.Println("Unbuffered: receiver ready")
	val := <-ch // unblocks sender
	fmt.Println("Unbuffered: received", val)
}

func bufferedDemo() {
	ch := make(chan int, 2) // buffered capacity = 2

	fmt.Println("Buffered: sending 1")
	ch <- 1 // does NOT block

	fmt.Println("Buffered: sending 2")
	ch <- 2 // does NOT block

	fmt.Println("Buffered: buffer full now")

	go func() {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Buffered: receiver reading")
		fmt.Println("Buffered: received", <-ch)
	}()

	fmt.Println("Buffered: sending 3 (will block)")
	ch <- 3 // BLOCKS until one value is read
	fmt.Println("Buffered: sent 3")

	for num := range ch {
		fmt.Println("Buffered: received", num)
		if len(ch) == 0 {
			break
		}
	}
	close(ch)
}

func main() {
	unbufferedDemo()
	fmt.Println("---------------")
	bufferedDemo()
}
