package main

import (
	"fmt"
)

func greet(a, b int) (int, int, int) {
	return a + b, a - b, a * b
}

func device(num ...int) int {
	sum := 0
	for _, v := range num {
		sum += v
	}
	return sum
}

func garvit() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func puma(num int) int {
	if num > 0 {
		fmt.Println(num)
		num--
	} else {
		return num
	}
	return puma(num)
}

func nano(ok bool) {
	ok = false
	fmt.Println("Inner Func:", ok)
}

func vim(ouk *bool) {
	*ouk = false
	fmt.Println("Inner Func:", *ouk)
}

type School struct {
	SchoolName string
	Year       int
}

func (s School) books(book string) string {
	return s.SchoolName + ": " + book
}

type Student struct {
	name string
	age  int
	School
}

type University struct {
	enrollNo int
	stud     Student
}

func (u *University) Semester(eno int) {
	u.enrollNo = eno
	fmt.Println(u.enrollNo)
}

type Payment interface {
	Pay(amount float32)
}

type Stripe struct{}

func (s Stripe) Pay(amount float32) {
	fmt.Println("Make process Stripe:", amount)
}

type Razorpay struct{}

func (r Razorpay) Pay(amount float32) {
	fmt.Println("Make process Razorpay:", amount)
}

func MakePayment(p Payment, amount float32) {
	p.Pay(amount)
}

type ServerState int

const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func MakeFunc(state ServerState) string {
	return "Server: " + stateName[state]
}

func main() {
	fmt.Println("Hello World")

	// fmt.Println("Hello", "Namaste")

	// var name string
	// name = "admin"
	// fmt.Println("Name:", name)

	// walrus operator
	// student := "John"
	// fmt.Println("Student:", student)

	// // firstName, lastName := "John", "Wick"
	// var firstName ,lastName string = "John","Wick"
	// fmt.Println("FirstName:", firstName)
	// fmt.Println("LastName:", lastName)

	// const enrollNo string = "23466f"
	// fmt.Println(enrollNo)
	// const (
	// 	rollno int = 81
	// 	enrollno string = "e2025"
	// )
	// fmt.Println(rollno)

	// for i := range 4 {
	// 	if i == 2 {
	// 		// break
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// for i := 1; i < 3; i++ {
	// 	fmt.Println(i)
	// }

	// i := 0
	// for i < 3 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// for {
	// 	println("hey")
	// }

	// num:=10
	// if num := 12; num <= 0 {
	// 	fmt.Println("Ops")
	// } else if num < 10 {
	// 	fmt.Println("Fine")
	// } else {
	// 	fmt.Println("Perfect")
	// }

	// switch day := "tuesday"; day {
	// case "monday":
	// 	fmt.Println("Its monday")
	// case "tuesday":
	// 	fmt.Println("Its tuesday")
	// default:
	// 	fmt.Println("Its day")
	// }

	// fruits := [3]string{"apple","banana","mango"}

	// var fruits = [3]string{"apple", "banana", "mango"}

	// var fruits [3]string
	// fruits[0] = "apple"
	// fruits[1] = "banana"
	// fruits[2] = "mango"

	// fruits := [...]int{1,2,3,4}
	// fruits := [2][4]int{{1,2,3,4},{5,6,7,8}}
	// fmt.Println(fruits)
	// fmt.Println(len(fruits))
	// fmt.Println(cap(fruits))

	// courses := []string{"bca", "bba", "btech"}
	// fmt.Println(courses)

	// courses = append(courses, "bcom", "mcom")
	// fmt.Println(courses)
	// fmt.Println(courses[:2])

	// books := make([]string, 0)
	// books = append(books, "cpp", "golang", "js", "python")
	// fmt.Println(books)

	// var books = []string{"cpp","js","golang"}
	// fmt.Println(books)

	// var books []string
	// books = append(books, "bca")
	// fmt.Println(books)

	// info := make(map[string]int)
	// info["monday"] = 1
	// info["tuesday"] = 2
	// info["wednesday"] = 3
	// fmt.Println(info["tuesday"])
	// fmt.Println(info)

	// info := map[string]int{"monday": 1, "tuesday": 2, "wednesday": 3}
	// fmt.Println(info)

	// a, b, _ := greet(2, 3)
	// fmt.Println(a, b)

	// demo := func() {
	// 	fmt.Println("World")
	// }
	// demo()

	// fmt.Println(device(2, 3, 4, 5))

	// fukre := []int{2, 4, 5, 6}
	// fmt.Println(device(fukre...))

	// grv := garvit()
	// fmt.Println(grv())
	// fmt.Println(grv())

	// fmt.Println(puma(5))

	// var use string = "admin"
	// reuse := &use
	// fmt.Println(reuse)
	// fmt.Println(*reuse)
	// *reuse = "panel"
	// fmt.Println(*reuse)
	// fmt.Println(use)

	// ok := true
	// ouk := true

	// nano(ok)
	// fmt.Println(ok)

	// vim(&ouk)
	// fmt.Println(ouk)

	// username := "admin"
	// for i := range username {
	// 	fmt.Printf("%c ", username[i])
	// }
	// fmt.Println()

	// for i := 0; i < len(username); i++ {
	// 	fmt.Printf("%c \n", username[i])
	// }

	// newSchool := School{
	// 	SchoolName: "Harvard",
	// 	Year:       1998,
	// }
	// newStud := Student{
	// 	name:   "admin",
	// 	age:    19,
	// 	School: newSchool,
	// }

	// fmt.Println(newSchool)
	// fmt.Println(newStud)

	// fmt.Println(newStud.books("Golang"))

	// newUniversity := University{
	// 	enrollNo: 22657995,
	// 	stud:     newStud,
	// }
	// fmt.Println(newUniversity)
	// newUniversity.Semester(12345678)
	// fmt.Println(newUniversity)

	// newUser := struct {
	// 	name string
	// 	age  int
	// }{age: 21}

	// fmt.Println(newUser)
	// newUser.name = "admin"
	// fmt.Println(newUser)

	// MakePayment(Stripe{}, 1200)
	// MakePayment(Razorpay{}, 1300)

	
	// fmt.Println(MakeFunc(StateConnected))


}
