package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func mult(a, b int) int {
	return a * b
}

func addMult(a int, b int, c int) (int, int) {
	return a + b + c, a * b * c
}

func calFunc(a, b, c int, d bool) (int, int, int, bool) {
	return a, b, c, d
}

func multVal(values ...int) int {
	sum := 0
	for _, value := range values {
		sum += value
	}
	return sum
}

// if u want to pass any type of value
func anyValue(values ...interface{}) interface{} {
	return values
}

// func processIt(fn func(a int) int) {
// 	fn(81)
// }

func main() {

	// anonymous func
	greet := func() {
		fmt.Println("Boom")
	}

	fmt.Println(add(1, 2))

	fmt.Println(mult(1, 2))

	sum, mult := addMult(1, 2, 4)
	fmt.Println("Sum:", sum, "Multiply:", mult)

	greet()

	a, b, _, d := calFunc(10, 20, 30, true)
	fmt.Println(a, b, d)

	calcs := multVal(2, 3, 4, 5)
	fmt.Println(calcs)

	// anyCal := []interface{}{10, 20, "admin", true}
	// fmt.Println(anyValue(anyCal...))
	anyCal := anyValue(10, 20, "admin", true)
	fmt.Println(anyCal)

	// fn := func(a int) int {
	// 	return 2
	// }
	// processIt(fn)

}
