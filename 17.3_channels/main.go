// ------------------------------------------------> buffer channel - non-blocking

package main

import (
	"fmt"
	"time"
)

// NOTE: channel direction specification is optional
// emailChannel <-chan string: emailChannel is a read-only channel
// done chan<- bool: done is a write-only channel

// func emailSender(emailChannel chan string, done chan bool) {
func emailSender(emailChannel <-chan string, done chan<- bool) {
	defer func() {
		// sending true into the channel indicates that the task is completed
		done <- true
	}()

	for email := range emailChannel {
		// fmt.Println("processing...")
		fmt.Println("sending email to", email)
		time.Sleep(time.Second * 2)
	}

}

func main() {
	// Step-1: create buffered channel
	emailChannel := make(chan string, 100)
	done := make(chan bool)

	// Step-2: call goroutine
	go emailSender(emailChannel, done)

	// fmt.Println(<-emailChannel)
	// fmt.Println(<-emailChannel)
	// fmt.Println(<-emailChannel)

	// Step-3: send mssg to channel - non-blocking
	for i := 1; i <= 3; i++ {
		emailChannel <- fmt.Sprintf("admin%d@gmail.com", i)
	}

	// Step-4: close the channel
	close(emailChannel) // !important

	// Step-5: wait for response in done variable from goroutine and print it
	// fmt.Println("done processing...", <-done)
	<-done

}

/*
NOTE ABOUT CLOSING CHANNELS:

1) CLOSE the channel when:
	- The channel is used to send MULTIPLE values (like a data stream),
	- And the receiver is using `for range channel` to read the values.
	- Closing the channel tells the receiver that NO MORE values will come, so the loop can stop safely.
	Example: `emailChannel` should be closed after sending all emails.

2) DO NOT close the channel when:
	- The channel is used only for sending ONE single signal/notification,
	- And the receiver reads it only once (not inside a loop).
	- In this case, closing is unnecessary and doesn't provide any benefit.
	Example: A `done` channel used only to send a single "task finished" signal does NOT need to be closed.
*/
