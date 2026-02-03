package main

import (
	"fmt"
	// "math/rand"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d picked job %d\n", id, job)

		// Random time → breaks any illusion of order
		time.Sleep(400 * time.Millisecond)

		fmt.Printf("Worker %d finished job %d\n", id, job)
		results <- job
	}

	fmt.Printf("Worker %d exiting\n", id)
}

func main() {

	const (
		numJobs    = 7
		numWorkers = 4
	)

	jobs := make(chan int) // UNBUFFERED → forces coordination
	results := make(chan int)

	var wg sync.WaitGroup

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Job producer
	go func() {
		for j := 1; j <= numJobs; j++ {
			fmt.Println("Trying to send job", j)
			jobs <- j
			fmt.Println("Successfully sent job", j)

		}
		close(jobs) // signals "no more work"
	}()

	// Result closer (VERY IMPORTANT pattern)
	go func() {
		wg.Wait()      // wait for all workers to finish
		close(results) // now safe to close
	}()

	// Result consumer
	for r := range results {
		fmt.Println("Result received:", r)
	}

	fmt.Println("All work done")
}
