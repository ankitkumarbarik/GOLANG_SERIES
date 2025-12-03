package main

import (
	"fmt"
	"os"
)

func main() {
	// f, err := os.Open("./19_files/exm.txt")
	// if err != nil {
	// 	// log the error, panic
	// 	panic(err)
	// }

	// fileInfo, err := f.Stat()
	// fmt.Println("filename:",fileInfo.Name())
	// fmt.Println("is folder? :",fileInfo.IsDir())
	// fmt.Println("filesize? :",fileInfo.Size())
	// fmt.Println("file permission :",fileInfo.Mode())
	// fmt.Println("file modified :",fileInfo.ModTime())

	// --------------------->

	// read file
	f, err := os.Open("./19_files/exm.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	buf := make([]byte, 12)
	d, err := f.Read(buf)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(buf); i++ {
		fmt.Println("data", d, string(buf[i]))
	}

	//

	data, err := os.ReadFile("./19_files/exm.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

	//

	// read folders
	// dir, err := os.Open("../")
	// if err != nil {
	// 	panic(err)
	// }

	// defer dir.Close()

	// fileInfo, err := dir.ReadDir(-1)

	// for _, fi := range fileInfo {
	// 	fmt.Println(fi.Name(), fi.IsDir())
	// }

	//

	// create a file
	// f, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// f.WriteString("hi go")
	// f.WriteString("nice language")
	// bytes := []byte("Hello Golang")
	// f.Write(bytes)

	//

	// read and write to another file (streaming fashion)
	// sourceFile, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer sourceFile.Close()

	// destFile, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer destFile.Close()

	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(destFile)

	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil {
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}

	// 		break
	// 	}

	// 	e := writer.WriteByte(b)
	// 	if e != nil {
	// 		panic(e)
	// 	}

	// }

	// writer.Flush()

	// fmt.Println("written to new file succesfully")

	// delete a file

	// sourceFile, err := os.Open("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer sourceFile.Close()

	// err := os.Remove("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(" file deleted successfully")

}
