package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello File")

	// creating file
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error while creating file", err)
		return
	}
	defer file.Close()

	fmt.Println("file created successfully", file)

	// writing file
	content := "Hey, Welcome to GoLang Series\n"
	bytes, errors := io.WriteString(file, content)
	if errors != nil {
		fmt.Println("Error while writing file", errors)
		return
	}

	fmt.Println("writing file successfully", bytes)

	// read file
	fileOpen, errs := os.Open("example.txt")
	if errs != nil {
		fmt.Println("Error while opening file", errs)
		return
	}
	defer fileOpen.Close()

	// create a buffer to read the file content
	buffer := make([]byte, 1024)

	for {
		data, ers := fileOpen.Read(buffer)
		if ers == io.EOF {
			break
		}
		if ers != nil {
			fmt.Println("Error while reading file", ers)
			return
		}
		fmt.Println(string(buffer[:data]))
	}

	// read file
	n, er := os.ReadFile("example.txt")
	if er != nil {
		fmt.Println("Error while reading file", er)
		return
	}
	fmt.Println(string(n))

	//

}
