package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGetReq() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("error while fetching response", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("error while getting response", res.Status)
		return
	}
	fmt.Println(res.Body)

	var todo Todo
	errors := json.NewDecoder(res.Body).Decode(&todo)
	if errors != nil {
		fmt.Println("error while decoding", errors)
		return
	}
	fmt.Println(todo)
}

func performPostReq() {
	todo := Todo{
		UserID:    23,
		Id:        101,
		Title:     "innovesta task",
		Completed: true,
	}

	// convert struct to json
	data, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("error while marshal into json", err)
		return
	}

	// convert json data to string
	// fmt.Println(string(data))

	// convert json data to reader-stream
	jsonRead := bytes.NewBuffer(data)

	// send post request
	res, errs := http.Post("https://jsonplaceholder.typicode.com/todos/", "application/json", jsonRead)
	if errs != nil {
		fmt.Println("error while creating todo", errs)
		return
	}
	defer res.Body.Close()

	// convert json stream to struct
	var todoRes Todo
	errors := json.NewDecoder(res.Body).Decode(&todoRes)
	if errors != nil {
		fmt.Println("error while response", errors)
		return
	}
	fmt.Println(todoRes, res.Status)
}

func performPutReq() {
	todo := Todo{
		UserID:    23,
		Id:        101,
		Title:     "innovesta task completed",
		Completed: true,
	}
	data, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("error while updating data")
		return
	}
	// fmt.Println(string(data))

	// convert json data to reader-stream
	jsonRead := bytes.NewBuffer(data)

	req, errs := http.NewRequest(http.MethodPut, "https://jsonplaceholder.typicode.com/todos/1", jsonRead)
	if errs != nil {
		fmt.Println("error while updating data", errs)
		return
	}

	req.Header.Set("Content-type", "application/json")

	// send the request
	client := http.Client{}
	res, errors := client.Do(req)
	if errors != nil {
		fmt.Println("error while request data", errors)
		return
	}
	defer res.Body.Close()

	resp, errorr := io.ReadAll(res.Body)
	if errorr != nil {
		fmt.Println("error while responsing data", errorr)
		return
	}
	fmt.Println(string(resp), res.Status)

}

func performDeleteReq() {
	req, err := http.NewRequest(http.MethodDelete, "https://jsonplaceholder.typicode.com/todos/1", nil)
	if err != nil {
		fmt.Println("error while deleting todo", err)
		return
	}

	req.Header.Set("Content-type", "application/json")

	client := http.Client{}
	res, errs := client.Do(req)
	if errs != nil {
		fmt.Println("error while sending delete request", errs)
		return
	}
	defer res.Body.Close()

	data, errors := io.ReadAll(res.Body)
	if errors != nil {
		fmt.Println("error while reading delete response", errors)
		return
	}
	fmt.Println(string(data), res.Status)
}

func main() {
	fmt.Println("Learning CRUD...")

	performGetReq()
	performPostReq()
	performPutReq()
	performDeleteReq()
}
