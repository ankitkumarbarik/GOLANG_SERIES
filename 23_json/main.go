package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	fmt.Println("JSON")

	person1 := Person{
		Name:    "John",
		Age:     34,
		IsAdult: true,
	}
	fmt.Println(person1)

	// convert into json format (encoding)
	data, err := json.Marshal(person1)
	if err != nil {
		fmt.Println("error while converting json")
		return
	}
	fmt.Println(string(data))

	// convert into struct format (decoding)
	var person2 Person
	err = json.Unmarshal(data, &person2)
	if err != nil {
		fmt.Println("error while converting struct")
		return
	}
	fmt.Println(person2)
}
