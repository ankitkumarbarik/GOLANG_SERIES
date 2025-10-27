package main

import "fmt"

// constant
// NOTE:  Capitalization is the rule: Any constant, variable, function, or type name that begins with a capital letter is exported and can be used by other packages. No special keywords: Unlike other languages, Go does not use keywords like public or export. The naming convention is the sole determinant of visibility.
const RefreshToken string = "jwt_token" // public (global)

func main() {

	// int
	fmt.Println(RefreshToken)
	fmt.Println(81)
	fmt.Println(81 + 2)
	// string
	fmt.Println("admin dev")
	// bool
	fmt.Println(true)
	fmt.Println(false)
	// floats
	fmt.Println(3.24)
	fmt.Println(3.24 / 3.0)

	// Variables

	// var name = "admin"  // implicit type

	var name string = "admin"
	fmt.Println(name)
	fmt.Printf("Type of name variable: %T \n", name)
	fmt.Printf("Name: %s \n", name)

	// shorthand syntax - walrus operator (best for initialize the variables) note:- this syntax can't work in global
	// name := "admin"

	var age int = 81
	fmt.Println(age)
	fmt.Printf("Type of age variable: %T \n", age)
	fmt.Printf("Age: %d \n", age)

	var success bool
	success = true
	fmt.Println(success)
	fmt.Printf("Type of success variable: %T \n", success)
	fmt.Printf("Success: %t \n", success)

	var price float32 = 72.8
	fmt.Println(price)
	fmt.Printf("Type of price variable: %T \n", price)
	fmt.Printf("Price: %g \n", price)

	var temp int // by default: 0 / false(bool)
	fmt.Println(temp)

	// constant
	const token string = "access_token"
	fmt.Println(token)
	// token = "456"  // we can;t change the value of constant

	// multiple constants
	const (
		host = "localhost"
		port = 5000
	)
	fmt.Println(host, port)

}
