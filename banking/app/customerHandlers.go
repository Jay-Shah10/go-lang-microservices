package app

import (
	"encoding/json"
	"net/http"

	"github.com/Jay-Shah10/go-lang-microservices/banking/service"
	"github.com/gorilla/mux"
)

// These are hints for JSON encoding.
// We can also put them both in the same `` so that we do not need to make duplicate struct.

// type Customer struct {
// 	Name    string `json: "full_name" xml: "full_name"`
// 	City    string `json: "city" xml: "city"`
// 	Zipcode string `json: "zip_code" xml: "zip_code"`
// }

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

	// customers, _ := ch.service.GetAllCustomer()
	// // Condition for the different type of headers.
	// // This way we can handle both types of data.
	// if r.Header.Get("Content-Type") == "application/json" {
	// 	w.Header().Add("Content-Type", "application/json")
	// 	json.NewEncoder(w).Encode(customers)
	// } else {
	// 	w.Header().Add("Content-Type", "application/xml")
	// 	xml.NewEncoder(w).Encode(customers)
	// }
	status := r.URL.Query().Get("status")
	customers, err := ch.service.GetAllCustomer(status)
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, err.AsMessage())

	} else {
		writeResponse(w, http.StatusOK, customers)
	}

}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	// Getting the value of customer_id from the url.
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}

}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	// should always be in this order.
	// if not in this order - content type will not be applied.
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
