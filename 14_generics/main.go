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

}
