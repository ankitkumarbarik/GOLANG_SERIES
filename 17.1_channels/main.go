package main

import (
	"fmt"
	// "time"
)

func processNum(mssgChan chan int, num1 int, num2 int) {
	// Step-4: wait for mssg in mssgChan variable from main function
	mssg := <-mssgChan
	fmt.Println("Before process (getting mssgs) number", mssg)

	// Step-5: send mssg to channel and wait for response
	mssgChan <- num1 + num2
}

func main() {
	// Step-1: create channel
	mssgChan := make(chan int)

	// Step-2: call goroutine
	go processNum(mssgChan, 4, 3)

	// Step-3: send mssg to channel and wait for response in mssgChan variable from goroutine
	mssgChan <- 200

	// Step-5: receive mssg from channel and print it
	fmt.Println("After process", <-mssgChan)

	// time.Sleep(time.Second * 2)

}
