package main

import (
	"fmt"
	"time"
)

// Goroutine -> lightweight thread of execution in go language.
// A goroutine is a function that runs concurrently with other functions.
// It is scheduled by the Go runtime and can run in parallel with other goroutines.
// Goroutines are multiplexed onto threads, so if one goroutine blocks, others can still run on the same thread.

func work() {
	for i := 0; i < 5; i++ {
		fmt.Println("worker", i)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	fmt.Println("Hello")

	go work()

	for i := 0; i < 5; i++ {
		fmt.Println("main", i)
		time.Sleep(time.Second * 1)
	}

}
