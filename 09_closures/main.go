package main

import "fmt"

func counter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}

func main() {

	inc := counter()
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())

}
