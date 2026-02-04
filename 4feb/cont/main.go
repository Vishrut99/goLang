package main

import (
	"context"
	"fmt"
	"time"
)

// Example 1: context.WithCancel - Manual cancellation
func withCancelExample(cxt context.Context) {
	fmt.Println("\n=== WithCancel Example ===")
	
	// Create a cancellable context
	ctx, cancel := context.WithCancel(cxt)
	
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine cancelled:", ctx.Err())
				return
			default:
				fmt.Println("Working...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()
	
	time.Sleep(2 * time.Second)
	cancel() // Manually cancel the context
	time.Sleep(500 * time.Millisecond)
}

// Example 2: context.WithTimeout - Automatic cancellation after timeout
func withTimeoutExample(c context.Context) {
	fmt.Println("\n=== WithTimeout Example ===")
	
	// Context will auto-cancel after 1 second
	ctx, cancel := context.WithTimeout(c, 1*time.Second)
	defer cancel() // Good practice to call cancel even if timeout occurs
	
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Timeout reached:", ctx.Err())
				return
			default:
				fmt.Println("Processing...")
				time.Sleep(800 * time.Millisecond)
			}
		}
	}()
	
	time.Sleep(2*time.Second)
}

// Example 3: context.WithDeadline - Cancel at specific time
func withDeadlineExample(c context.Context) {
	fmt.Println("\n=== WithDeadline Example ===")
	
	// Set deadline to 1.5 seconds from now
	deadline := time.Now().Add(1500 * time.Millisecond)
	ctx, cancel := context.WithDeadline(c, deadline)
	defer cancel()
	
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Operation completed")
	case <-ctx.Done():
		fmt.Println("Deadline exceeded:", ctx.Err())
	}
}

// Example 4: context.WithValue - Passing values through context
func withValueExample(c context.Context) {
	fmt.Println("\n=== WithValue Example ===")
	
	type key string
	userKey := key("userID")
	
	// Create context with a value
	ctx := context.WithValue(c, userKey, "12345")
	
	// Pass context to function
	processRequest(ctx, userKey)
}

func processRequest(ctx context.Context, key any) {
	// Retrieve value from context
	if userID, ok := ctx.Value(key).(string); ok {
		fmt.Printf("Processing request for user: %s\n", userID)
	} else {
		fmt.Println("User ID not found in context")
	}
}

// Example 5: Real-world example - HTTP request simulation
func httpRequestSimulation() {
	fmt.Println("\n=== HTTP Request Simulation ===")
	
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	
	result := make(chan string, 1)
	
	go func() {
		// Simulate a slow API call
		// time.Sleep(3 * time.Second)
		result <- "API Response Data"
	}()
	
	select {
	case res := <-result:
		fmt.Println("Success:", res)
	case <-ctx.Done():
		fmt.Println("Request failed:", ctx.Err())
	}
}

func edgeCase1() {
	fmt.Println("\n=== Edge Case 1: Nil Value ===")
	
	type key string
	ctx := context.WithValue(context.Background(), key("data"), nil)
	
	val := ctx.Value(key("data"))
	fmt.Printf("Value: %v, Type: %T\n", val, val)
}

func edgeCase2() {
	fmt.Println("\n=== Edge Case 2: Multiple Done Reads ===")
	
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	
	<-ctx.Done() // First read
	<-ctx.Done() // Second read - will this block?
	<-ctx.Done() // Third read
	
	fmt.Println("Read Done() three times successfully")
}
func main() {

	cxt:=context.Background()
	withCancelExample(cxt)
	withTimeoutExample(cxt)
	withDeadlineExample(cxt)
	withValueExample(cxt)
	httpRequestSimulation()
	edgeCase1()
	edgeCase2()
	
}