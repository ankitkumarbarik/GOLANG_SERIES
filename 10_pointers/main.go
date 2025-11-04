package main

import "fmt"

// call by value
func callVal(num int) {
	num = 5
	fmt.Println("In callVal:", num)
}

// call by reference
func callRef(num *int) {
	*num = 5
	fmt.Println("In callRef:", *num)
}

func main() {

	a := 10
	b := &a
	fmt.Println("address of a:", &a)
	fmt.Println("address of b:", b)
	fmt.Println("value of b:", *b)

	var c int = 20
	var d *int = &c
	fmt.Println("address of c:", &c)
	fmt.Println("address of d:", d)

	*d = 30
	fmt.Println("value of c:", *d)
	fmt.Println("value of d:", c)

	num := 10
	callVal(num)
	println("After callVal:", num)

	no := 10
	callRef(&no)
	println("After callRef:", no)

}
