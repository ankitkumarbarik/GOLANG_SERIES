package main

import (
	"fmt"
	"maps"
)

// maps -> hash, object, dict
func main() {
	// creating maps
	lang := make(map[string]string)

	// adding
	lang["js"] = "javascript"
	lang["py"] = "python"
	lang["go"] = "golang"

	fmt.Println("List of all lang:", lang)
	fmt.Println("Go shorts for:", lang["go"])

	// deleting
	delete(lang, "go")
	fmt.Println("List of all lang:", lang)

	// clear
	// clear(lang)

	// looping
	for i := range lang {
		fmt.Println(lang[i])
	}

	for key, value := range lang {
		fmt.Println(key, "->", value)
		// fmt.Printf("For key %v, value is %v\n", key, value)
	}

	// another way to create map
	// m := map[string]int{}
	m := map[string]int{"price": 150, "discount": 30}
	m["originalPrice"] = 100
	fmt.Println(m)

	// check conditions
	v, ok := m["price"]
	// _, ok := m["prices"]

	if ok {
		fmt.Println("all ok Price:", v)
	} else {
		fmt.Println("not ok:", v)
	}

	// equals
	m1 := map[string]int{"price": 150, "discount": 30}
	m2 := map[string]int{"price": 150, "discount": 30}

	fmt.Println(maps.Equal(m1, m2))

}
