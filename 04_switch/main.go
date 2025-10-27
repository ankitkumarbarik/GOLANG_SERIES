package main

import (
	"fmt"
	"time"
)

func main() {
	i := 3

	switch i {
	case 1:
		fmt.Println(i)
	case 2:
		fmt.Println(i)
	case 3:
		fmt.Println(i)
	case 4:
		fmt.Println(i)
	default:
		fmt.Println(i)
	}

	// multiple condition switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's weekend")
	default:
		fmt.Println("it's workday")
	}

	// type swwitch
	whoAmI := func(i interface{}) {
		// switch i.(type) {
		switch t := i.(type) {
		case int:
			fmt.Println("it's an integer")
		case string:
			fmt.Println("it's an string")
		case bool:
			fmt.Println("it's an boolean")
		default:
			fmt.Println("other:", t)
		}
	}

	whoAmI("admin")
	whoAmI(81)
	whoAmI(81.5)
	whoAmI(true)

}
