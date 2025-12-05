package main

import (
	"fmt"
	"net/url"
)

func main() {
	myUrl := "https://api.github.com/users/ankitkumarbarik?key1=value1&key2=value2&key3=value3"
	fmt.Printf("Type of Url: %T\n", myUrl)

	// convert/parse the string type into url type
	parseUrl, err := url.Parse(myUrl)
	if err != nil {
		fmt.Println("error while parsing url", err)
	}
	fmt.Println(parseUrl)
	fmt.Printf("Type of Url: %T\n", parseUrl)

	fmt.Println(parseUrl.Scheme)
	fmt.Println(parseUrl.Host)
	fmt.Println(parseUrl.Path)
	fmt.Println(parseUrl.RawQuery)

	// modify the url
	parseUrl.Path = "/users/ankitbarik"
	parseUrl.RawQuery = "user=admin"
	fmt.Println(parseUrl)
	fmt.Printf("Type of Url: %T\n", parseUrl)

	// constructing a URL string from a URL object
	newUrl := parseUrl.String()
	fmt.Println(newUrl)
	fmt.Printf("Type of Url: %T\n", newUrl)
}
