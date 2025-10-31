package main

import (
	"fmt"
	"slices"
	"sort"
)

// slice -> dynamically
// most used construct in go
// + useful methods

func main() {

	// uninitialized slice is nill
	var fruits = []string{"apple", "banana", "orange"}
	fmt.Println(fruits)
	fmt.Println(fruits == nil)
	fmt.Println(len(fruits))
	fmt.Printf("Type of fruits variable: %T \n", fruits)

	fruits = append(fruits, "mango", "guava")
	fmt.Println(fruits)
	fmt.Println(fruits[3])

	// slice := fruits[1:]
	// slice := fruits[:4]
	// slice := fruits[:]
	slice := fruits[1:4]
	fmt.Println(slice)

	score := make([]int, 4)

	// lenght of slice -> 4
	fmt.Println(len(score))

	score[0] = 10
	score[1] = 20
	score[2] = 30
	score[3] = 40
	fmt.Println(score)
	fmt.Println(score[2])

	score = append(score, 30, 50)
	fmt.Println(score)

	fmt.Println(sort.IntsAreSorted(score))

	sort.Ints(score)
	fmt.Println(score)

	fmt.Println(sort.IntsAreSorted(score))

	//
	// orm := []int{}
	// orm := make([]int, 0, 5)
	orm := make([]int, 4, 5)

	fmt.Println(len(orm))
	fmt.Println(cap(orm)) // capacity of slice

	orm[0] = 14
	orm[1] = 13
	orm[2] = 12
	orm[3] = 11
	fmt.Println(orm)
	fmt.Println(cap(orm)) // capacity of slice

	orm = append(orm, 10, 9)
	fmt.Println(orm)
	fmt.Println(cap(orm)) // capacity of slice
	fmt.Println(len(orm)) // length of slice

	// copy func
	color := []string{}
	color = append(color, "blue")
	color = append(color, "red")
	fmt.Println(color)

	bucket := make([]string, len(color))
	copy(bucket, color)
	fmt.Println(bucket)

	bucket = append(bucket, "green")
	fmt.Println(bucket)

	// compare
	fmt.Println(slices.Equal(color, bucket))

	// 2D slices
	matrix := [][]int{{100, 200}, {100, 200}}
	fmt.Println(matrix)

	// remove
	courses := []string{"node", "django", "golang", "rust"}
	fmt.Println(courses)

	removed := courses[:3]
	fmt.Println(removed)

}
