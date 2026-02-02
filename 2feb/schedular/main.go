package main

import (
	"fmt"
	"runtime"
	"time"
)

func cpuHog() {
	for {
	}
}

func syscallBlock() {
	// time.Sleep(time.Second)
	fmt.Println("syscall goroutine ran")
}

func main() {
	runtime.GOMAXPROCS(1)

	go cpuHog()
	go syscallBlock()

	time.Sleep(1 * time.Second) // keep main alive
	fmt.Println("main finished")
}
