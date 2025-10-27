package main

import "fmt"

func main() {
	age := 10

	if age >= 18 {
		fmt.Println("U can vote")
	} else if age >= 12 {
		fmt.Println("U can't vote")
	} else {
		fmt.Println("Go school")
	}

	role := "admin"
	hasPermission := true

	if role == "admin" && hasPermission {
		fmt.Println("YES")
	}

	// if role == "admin" || hasPermission {
	// 	fmt.Println("YES")
	// }

	// we can declare a variable inside if construct
	if age := 15; age >= 18 {
		fmt.Println("Adult:", age)
	} else if age >= 12 {
		fmt.Println("Men:", age)
	} else {
		fmt.Println("Child:", age)
	}

	// if err != nil{
	// 	// something went wrong
	// }

	// NOTE: GO doesn't have ternary operator
}
