package main

import (
	"fmt"
	"time"
)

// Simple task structure
type Task struct {
	ID   int
	Name string
}

// Worker function that processes tasks
func worker(id int, tasks chan Task, done chan bool) {
	fmt.Printf("[Worker %d] Started and waiting for tasks\n", id)

	for task := range tasks {
		fmt.Printf("[Worker %d] Picked up Task %d: %s\n", id, task.ID, task.Name)

		// Simulate work
		time.Sleep(1 * time.Second)

		fmt.Printf("[Worker %d] Completed Task %d\n", id, task.ID)
	}

	fmt.Printf("[Worker %d] No more tasks, shutting down\n", id)
	done <- true
}

// Demonstration 1: Basic scheduler behavior
func demonstrateBasicScheduling() {
	fmt.Println("\n=== BASIC GOROUTINE SCHEDULING ===")
	fmt.Println("Creating 3 workers on your CPU cores...")

	// Create channels
	tasks := make(chan Task, 5) // Buffered channel for tasks
	done := make(chan bool, 3)  // Buffered channel for completion signals

	// Start 3 workers (goroutines)
	for i := 1; i <= 3; i++ {
		go worker(i, tasks, done)
	}

	// Give workers time to start
	time.Sleep(500 * time.Millisecond)

	// Send 6 tasks
	fmt.Println("\n[Main] Sending 6 tasks...")
	taskList := []Task{
		{ID: 1, Name: "Process invoice"},
		{ID: 2, Name: "Send email"},
		{ID: 3, Name: "Update database"},
		{ID: 4, Name: "Generate report"},
		{ID: 5, Name: "Backup files"},
		{ID: 6, Name: "Clean cache"},
	}

	for _, task := range taskList {
		tasks <- task
		fmt.Printf("[Main] Sent Task %d to queue\n", task.ID)
	}

	// Close tasks channel (no more work)
	close(tasks)
	fmt.Println("[Main] All tasks sent, waiting for workers to finish...")

	// Wait for all 3 workers to complete
	for i := 0; i < 3; i++ {
		<-done
	}

	fmt.Println("[Main] All workers finished!")
}

// Demonstration 2: Goroutine switching behavior
func demonstrateGoroutineSwitching() {
	fmt.Println("\n\n=== GOROUTINE CONTEXT SWITCHING ===")
	fmt.Println("Watch how scheduler switches between goroutines...")

	// Channel to coordinate completion
	finished := make(chan bool)

	// Goroutine 1: Counts numbers
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Printf("[Goroutine A] Count: %d\n", i)
			time.Sleep(300 * time.Millisecond) // Yield to scheduler
		}
		finished <- true
	}()

	// Goroutine 2: Prints letters
	go func() {
		letters := []string{"A", "B", "C", "D", "E"}
		for _, letter := range letters {
			fmt.Printf("[Goroutine B] Letter: %s\n", letter)
			time.Sleep(500 * time.Millisecond) // Yield to scheduler
		}
		finished <- true
	}()

	// Wait for both to complete
	<-finished
	<-finished

	fmt.Println("\nNotice: Goroutines interleaved execution (scheduler switching)")
}

// Demonstration 3: Blocking and unblocking
func demonstrateBlockingBehavior() {
	fmt.Println("\n\n=== GOROUTINE BLOCKING & UNBLOCKING ===")

	ch := make(chan string)

	// Sender goroutine
	go func() {
		fmt.Println("[Sender] About to send (will block until receiver ready)...")
		ch <- "Important Message"
		fmt.Println("[Sender] Message delivered! (was unblocked)")
	}()

	// Simulate receiver delay
	fmt.Println("[Main] Receiver not ready yet, sender is blocked...")
	time.Sleep(2 * time.Second)

	fmt.Println("[Main] Now receiving...")
	msg := <-ch
	fmt.Printf("[Main] Received: %s\n", msg)
}

// Demonstration 4: Work distribution visualization
func demonstrateWorkDistribution() {
	fmt.Println("\n\n=== WORK DISTRIBUTION ACROSS WORKERS ===")

	jobs := make(chan int, 10)
	results := make(chan string, 10)

	// Start 2 workers
	for w := 1; w <= 2; w++ {
		go func(workerID int) {
			for job := range jobs {
				fmt.Printf("  [Worker %d] Processing job %d...\n", workerID, job)
				time.Sleep(800 * time.Millisecond)
				results <- fmt.Sprintf("Job %d done by Worker %d", job, workerID)
			}
		}(w)
	}

	// Send 5 jobs
	fmt.Println("Sending 5 jobs to worker pool...")
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results
	fmt.Println("\nResults:")
	for i := 1; i <= 5; i++ {
		result := <-results
		fmt.Printf("  ✓ %s\n", result)
	}
}

// Demonstration 5: Simple timeout pattern
func demonstrateSimpleTimeout() {
	fmt.Println("\n\n=== SIMPLE TIMEOUT PATTERN ===")

	// Slow operation
	slowTask := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		slowTask <- "Task completed"
	}()

	// Use select for timeout
	fmt.Println("Waiting for task with 2 second timeout...")

	timeout := time.After(2 * time.Second)

	select {
	case result := <-slowTask:
		fmt.Printf("Success: %s\n", result)
	case <-timeout:
		fmt.Println("Timeout! Task took too long")
	}
}

// Demonstration 6: Producer-Consumer (simple version)
func demonstrateProducerConsumer() {
	fmt.Println("\n\n=== SIMPLE PRODUCER-CONSUMER ===")

	data := make(chan int, 3) // Buffer size 3
	done := make(chan bool)

	// Producer
	go func() {
		fmt.Println("[Producer] Producing items...")
		for i := 1; i <= 5; i++ {
			fmt.Printf("[Producer] Produced item %d\n", i)
			data <- i
			time.Sleep(500 * time.Millisecond)
		}
		close(data)
		fmt.Println("[Producer] Done producing")
	}()

	// Consumer
	go func() {
		fmt.Println("[Consumer] Consuming items...")
		for item := range data {
			fmt.Printf("[Consumer] Consumed item %d\n", item)
			time.Sleep(1 * time.Second) // Slower consumer
		}
		done <- true
	}()

	// Wait for consumer to finish
	<-done
	fmt.Println("[Main] Producer-Consumer complete")
}

func main() {
	fmt.Println("GO SCHEDULER: SIMPLE DEMONSTRATIONS")
	fmt.Println("====================================")
	fmt.Println("Runtime: GOMAXPROCS will use all your CPU cores")

	// Run demonstrations
	demonstrateBasicScheduling()

	demonstrateGoroutineSwitching()

	demonstrateBlockingBehavior()

	demonstrateWorkDistribution()

	demonstrateSimpleTimeout()

	demonstrateProducerConsumer()

	fmt.Println("\n\n=== ALL DEMONSTRATIONS COMPLETE ===")
	fmt.Println("\nKey Concepts Demonstrated:")
	fmt.Println("✓ Goroutines are scheduled across CPU cores")
	fmt.Println("✓ Scheduler switches between goroutines automatically")
	fmt.Println("✓ Channels cause blocking/unblocking")
	fmt.Println("✓ Work is distributed among available goroutines")
	fmt.Println("✓ Buffered channels enable async producer-consumer")
	fmt.Println("✓ Select enables timeout patterns")
}
