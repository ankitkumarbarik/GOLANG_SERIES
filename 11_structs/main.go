package main

import (
	"fmt"
	"time"
)

// no inheritance in golang; no super or parent

type Order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
}

// constructor
func NewOrder(id string, amount float32, status string) Order {
	return Order{
		id:     id,
		amount: amount,
		status: status,
	}
}

// methods
// NOTE: use (*) pointer when u want to change the value of the field inside the struct (like status)
func (o *Order) changeStatus(status string) {
	o.status = status
}

func (o Order) getAmount() float32 {
	return o.amount
}

func main() {

	// var myOrder1 Order // struct literal
	// myOrder1.id = "101"
	// myOrder1.amount = 45.80
	// myOrder1.status = "received"

	// var myOrder1 = Order{
	// 	id:     "101",
	// 	amount: 45.80,
	// 	status: "received",
	// }

	// note: if u don't set the value of the field then it will be zero value.
	myOrder1 := Order{
		id:     "101",
		amount: 45.80,
		status: "received",
	}

	// adding new field
	myOrder1.createdAt = time.Now()

	fmt.Println("Order struct:", myOrder1)
	fmt.Printf("Order struct: %+v\n", myOrder1)

	// accessing fields
	fmt.Println("Order status:", myOrder1.status)

	myOrder2 := Order{
		id:        "102",
		amount:    67.98,
		status:    "delivered",
		createdAt: time.Now(),
	}
	myOrder1.status = "paid"
	fmt.Println("Order struct:", myOrder1)
	fmt.Println("Order struct:", myOrder2)

	//
	myOrder3 := Order{
		id:        "103",
		amount:    37.32,
		status:    "arrived",
		createdAt: time.Now(),
	}
	fmt.Println(myOrder3)
	myOrder3.changeStatus("confirmed")
	fmt.Println(myOrder3)

	fmt.Println(myOrder3.getAmount())

	o1 := NewOrder("101", 45.80, "received")
	fmt.Println(o1)

	// inline struct -> usecase... if u want to initialize one struct
	lang := struct {
		Name string
		Age  int
	}{Name: "Go", Age: 10}

	fmt.Println(lang)

	// struct embedding
	type School struct {
		school_id   int
		school_name string
	}
	type Student struct {
		id   int
		name string
		School
	}

	newSchool := School{
		school_id:   1,
		school_name: "Harvard",
	}
	newStud := Student{
		id:     101,
		name:   "John",
		School: newSchool,
	}

	fmt.Println(newStud)

	newStud.School.school_name = "Scalar"
	fmt.Println(newStud)

}
