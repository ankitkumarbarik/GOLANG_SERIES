package main

import "fmt"

// enumerated types

// iota:- value of the constant will be 0 by default and will be incremented by 1 for each subsequent constant declaration in the same block of constants or in the same file (whichever comes first).

type PaymentStatus int

const (
	Success PaymentStatus = iota
	Pending
	Failed
)

// if u want string type then use this way below instead of iota type..
// type PaymentStatus string

// const (
// 	Success PaymentStatus = "success"
// 	Pending               = "pending"
// 	Failed                = "failed"
// )

func changePaymentStatus(status PaymentStatus) {
	fmt.Println("changing payment status", status)
}

func main() {
	fmt.Println("Hello world")

	// if 1 == Pending {
	// 	fmt.Println("right")
	// }

	changePaymentStatus(Pending)
	// changePaymentStatus("success")
}
