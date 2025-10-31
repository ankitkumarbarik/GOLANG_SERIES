package main

import "fmt"

// arrays: number sequence of specific length (fixed size). If u want flexibility then use "slice".

// zeroed values
// int -> 0, string -> "", bool -> false

func main() {

	var nums [5]int

	fmt.Println("Length:", len(nums))
	nums[0] = 12
	nums[1] = 13
	nums[2] = 14
	fmt.Println("Nums:", nums)
	fmt.Println("Nums[1]:", nums[1])

	var vals [4]bool
	vals[2] = true
	fmt.Println(vals)

	var name [3]string
	name[1] = "golang"
	fmt.Println(name)

	// initialize ------------>
	// var fruits = [...]string{"lemon", "apple", "orange"}
	var fruits = [3]string{"lemon", "apple", "orange"}
	fmt.Println(fruits)
	fmt.Println(fruits[1])

	success := [2]bool{true, false}
	fmt.Println(success)
	fmt.Println(success[0])

	// auto length detection ------------->
	colors := [...]string{"black", "white", "blue", "purple", "pink"}
	colors[2] = "brown"
	fmt.Println(colors)
	fmt.Println(len(colors))

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	// 2D arrays --------->

	// var logs = [...][...]int{{1,2,3},{4,5,6}}
	// var logs = [2][3]int{{1,2,3},{4,5,6}}

	// var logs [2][3]int
	// logs[0][0] = 12
	// logs[0][1] = 13
	// fmt.Println(logs[0][1])

	// logs := [...][...]int{{1,2,3},{4,5,6}}
	logs := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(logs[1][2])

	// Matrix Operation -------------------->
	fmt.Printf("Matrix Operation ----------> \n\n")

	var A = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	var B = [2][3]int{{7, 8, 9}, {10, 11, 12}}
	var C [2][3]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d\t", A[i][j])
		}
		fmt.Println()
	}

	fmt.Println()

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d\t", B[i][j])
		}
		fmt.Println()
	}

	fmt.Println()

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			C[i][j] = A[i][j] + B[i][j]
			fmt.Printf("%d\t", C[i][j])
		}
		fmt.Println()
	}

}
