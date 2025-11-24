package main

import "fmt"

func main() {
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "pong"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Value := <-chan1:
			fmt.Println("receive data from chan1", chan1Value)
		case chan2Value := <-chan2:
			fmt.Println("receive data from chan2", chan2Value)
		}
	}
}
