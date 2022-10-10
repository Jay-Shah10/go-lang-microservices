package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/Jay-Shah10/go-lang-microservices/banking/service"
)

// These are hints for JSON encoding.
// We can also put them both in the same `` so that we do not need to make duplicate struct.

type Customer struct {
	Name    string `json: "full_name" xml: "full_name"`
	City    string `json: "city" xml: "city"`
	Zipcode string `json: "zip_code" xml: "zip_code"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

// handler functions.
// func greet(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Hello World")
// }

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) { // this is starting to make the connection and then to the rest.
	// customers := []Customer{
	// 	{Name: "John", City: "Test", Zipcode: "12345"},
	// }

	customers, _ := ch.service.GetAllCustomer()

	// Condition for the different type of headers.
	// This way we can handle both types of data.
	if r.Header.Get("Content-Type") == "application/json" {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	}

}
