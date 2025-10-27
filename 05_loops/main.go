package main

import "fmt"

// for -> one construct in go for looping
func main() {
	// while loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// infinite loop
	// for {
	// 	fmt.Println("admin")
	// }

	// for loop
	for i := 0; i <= 3; i++ {
		fmt.Println(i)
	}
	// TODO: Here, u can use continue and break

	for i := range 3 {
		fmt.Println(i)
	}

}
