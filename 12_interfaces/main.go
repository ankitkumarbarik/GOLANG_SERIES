package main

import "fmt"

// problem when we don't use interface
// type Razorpay struct{}
// func (p Razorpay) pay(amount float32) {
// 	fmt.Println("Making payment using razorpay", amount)
// }

// type Stripe struct{}
// func (p Stripe) pay(amount float32) {
// 	fmt.Println("Making payment using stripe", amount)
// }

// type Payment struct{
// 	// another way-- create instance here..
// 	stripePaymentGw Stripe
// }

// func (p Payment) makePayment(amount float32) {
// 	razorpayPaymentGw := Razorpay{} // create instance
// 	razorpayPaymentGw.pay(amount)

// 	// another way
// 	p.stripePaymentGw.pay(amount)
// }

//--------------------------------------------------------->

type Payment interface {
	Pay(amount float32)
}

type Razorpay struct{}

func (r Razorpay) Pay(amount float32) {
	fmt.Println("Make the razorpay payment process", amount)
}

type Stripe struct{}

func (s Stripe) Pay(amount float32) {
	fmt.Println("Make the stripe payment process", amount)
}

func MakePayment(p Payment, amount float32) {
	p.Pay(amount)
}

func main() {
	fmt.Println("INTERFACE")

	// stripePay := Stripe{}
	// newPament := Payment{
	// 	stripePaymentGw: stripePay,
	// }

	// newPament.makePayment(100)

	//----------------------------------

	MakePayment(Stripe{}, 1200)
	newRazorpayPayment := Razorpay{}
	MakePayment(newRazorpayPayment, 1300)

}
