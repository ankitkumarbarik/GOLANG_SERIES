package main

import "fmt"

// func printSlice[T interface{}](items []T) {
// func printSlice[T any](items []T) {
// func printSlice[T int | string | bool](items []T) {
// func printSlice[T int, V string](items []T, name V) {
func printSlice[T comparable](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// func printStrSlice(items []string) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

// LIFO
type Stack[T any] struct {
	elements []T
}

// ~int | ~float32 -> constraints... it used to specify the type of the value inside the interface..
type Number interface {
	~int | ~float32
}

func Pay[T Number](amount []T) T {
	sum := T(0)
	for _, v := range amount {
		sum += v
	}
	return sum
}

func main() {
	nums := []int{1, 2, 3}
	printSlice(nums)

	names := []string{"golang", "ts", "js"}
	// printStrSlice(names)
	printSlice(names)
	printSlice([]bool{true, false, true})

	//
	myStack := Stack[string]{
		elements: []string{"golang", "rust"},
	}
	fmt.Println(myStack)

	//
	amount := Pay([]float32{1.1, 2.2, 3.3})
	fmt.Println(amount)

}
