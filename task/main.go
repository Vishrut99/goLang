package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type event struct {
	userid int
	Value  int
}

func main() {

	const (
		tp = 5
		tc = 5
	)
	producer := make(chan event, tp)
	wg := sync.WaitGroup{}

	ctxp, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Add(1)
	go work(producer, ctxp, cancel, &wg)

	for i := 0; i < tc; i++ {
		wg.Add(1)
		go manager(producer, ctxp, &wg)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	select {
	case <-sigChan:
		cancel()
		wg.Wait()
		fmt.Println("Interrupt signal received, shutting down gracefully...")
		close(producer)

	case <-ctxp.Done():
		wg.Wait()
		fmt.Println("Work completed, shutting down Automatically...")
		close(producer)
	}
}
