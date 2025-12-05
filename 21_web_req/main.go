package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Web fetch")

	res, err := http.Get("https://api.github.com/users/ankitkumarbarik")
	if err != nil {
		fmt.Println("Error while fetching data", err)
		return
	}
	fmt.Printf("Type of response %T\n", res)
	defer res.Body.Close()

	data, errs := io.ReadAll(res.Body)
	if errs != nil {
		fmt.Println("Error while fetching data", errs)
	}

	fmt.Println(string(data))
}
