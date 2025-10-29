// Practice : Leap Year

// Write a program to check whether a year is a leap year.

// (Hint: A year is leap if divisible by 400, or divisible by 4 but not by 100.)


package main
import "fmt"

func main() {
	year := 2020
	
	if year%400 ==0 || (year%4 == 0 && year%100 != 0) {
		fmt.Println(year, "is a leap year")
	} else {
		fmt.Println(year, "is not a leap year")
	}

}