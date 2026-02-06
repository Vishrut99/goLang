package main

import (
	"fmt"
	// "math/rand"
	"sync"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		// fmt.Printf("Worker %d picked job %d\n", id, job)

		// Random time → breaks any illusion of order
		// time.Sleep(40 * time.Millisecond)

		sum := 0
		for i := 0; i < 80000; i++ {
			sum += (i ^ job) & 0xFF
		}
		_ = sum

		// fmt.Printf("Worker %d finished job %d\n", id, job)
		results <- job
	}

	// fmt.Printf("Worker %d exiting\n", id)
}

func RunWorkerPool(numJobs, numWorkers int) []int {
	jobs := make(chan int)
	results := make(chan int)

	var wg sync.WaitGroup

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Producer
	go func() {
		for j := 1; j <= numJobs; j++ {
			jobs <- j
		}
		close(jobs)
	}()

	// Close results after workers finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	var out []int
	for r := range results {
		out = append(out, r)
	}

	return out
}

func main() {
	const (
		numJobs    = 7
		numWorkers = 4
	)

	results := RunWorkerPool(numJobs, numWorkers)

	for _, r := range results {
		fmt.Println("Result received:", r)
	}

	fmt.Println("All work done")
}
