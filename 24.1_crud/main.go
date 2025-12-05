package main

import (
	"encoding/json"
	"fmt"
	// "io"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("Learning CRUD...")

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

	// data, errs := io.ReadAll(res.Body)
	// if errs != nil {
	// 	fmt.Println("error while fetching response", errs)
	// 	return
	// }
	// fmt.Println(string(data))

	var todo Todo
	errors := json.NewDecoder(res.Body).Decode(&todo)
	if errors != nil {
		fmt.Println("error while decoding", errors)
		return
	}
	fmt.Println(todo)

}

// NOTE: ======================= JSON IN GOLANG (NOTES) =======================
//
// JSON is just TEXT. Go works with strictly typed variables (structs).
// Go cannot directly use JSON text, so conversion is required.
//
// There are only TWO real conversions:
// 1. JSON -> Go data
// 2. Go data -> JSON
//
// ----------------------- BASIC RULE -----------------------
// JSON -> Go   = Decode / Unmarshal
// Go   -> JSON = Encode / Marshal
//
// ----------------------- NON-HTTP CASE -----------------------
// Marshal   : Converts Go struct into JSON bytes ([]byte)
// Unmarshal : Converts JSON bytes into Go struct
//
// Use Marshal / Unmarshal when:
// - JSON is already in memory
// - Working with files, logs, or tests
//
// ----------------------- HTTP CASE -----------------------
// Encoder : Writes Go struct as JSON directly to HTTP response
// Decoder : Reads JSON from HTTP request body into Go struct
//
// Use Encoder / Decoder when:
// - Handling HTTP requests and responses
//
// ----------------------- STRUCT TAGS -----------------------
// json:"name" maps JSON field names to Go struct fields
// Go struct fields must be exported (capitalized)
//
// ----------------------- FLOW -----------------------
// Client sends JSON text
// -> Decoder converts JSON to Go struct
// -> Business logic runs
// -> Encoder converts Go struct to JSON
// -> JSON response sent to client
//
// ----------------------- GOLDEN RULE -----------------------
// HTTP -> Encoder / Decoder
// NON-HTTP -> Marshal / Unmarshal
// =====================================================================
