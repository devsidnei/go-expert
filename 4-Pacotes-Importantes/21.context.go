package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Create a context with a timeout of 100 milliseconds
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	// Simulate a long-running task
	time.Sleep(101 * time.Millisecond)

	// Check if the context has been canceled
	if ctx.Err() != nil {
		fmt.Println("Task canceled!")
		return
	}

	// Otherwise, the task completed successfully
	fmt.Println("Task completed!")
}
