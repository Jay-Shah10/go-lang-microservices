package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json: "full_name"`
	City    string `json: "City"`
	Zipcode string `json: "zip_code"`
}

func main() {

	// function handlers.
	http.HandleFunc("/status/check", statusCheckHandler)
	http.ListenAndServe("localhost:8000", nil)
}

// handler functions.
func statusCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "John", City: "Test", Zipcode: "12345"},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
