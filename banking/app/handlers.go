package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

// These are hints for JSON encoding.
// We can also put them both in the same `` so that we do not need to make duplicate struct.

type Customer struct {
	Name    string `json: "full_name" xml: "full_name"`
	City    string `json: "city" xml: "city"`
	Zipcode string `json: "zip_code" xml: "zip_code"`
}

// handler functions.
func statusCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "John", City: "Test", Zipcode: "12345"},
	}

	// Condition for the different type of headers.
	// This way we can handle both types of data.
	if r.Header.Get("Contet-Type") == "application/json" {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Context-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	}

}
