package main

import (
	"fmt"
	// "time"
)

// defer means this code will run at the end of the function,
// no matter how the function exits (normal execution, return, or error).

func task(done chan bool) {
	defer func() {
		// sending true into the channel indicates that the task is completed
		done <- true
	}()

	fmt.Println("processing...")
}

func main() {
	// Step-1: create channel
	done := make(chan bool)

	// Step-2: call goroutine
	go task(done)

	// Step:3: wait for response in done variable from goroutine and print it
	fmt.Println("After process", <-done)

	// <-done //bocking
}
