package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()

	// function handlers.
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet) // this option will make sure the value is numerical and mux will error handle.

	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost) // getting more post.

	// failure to start server.
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

// ability to print out a value from the url in this example. We can also use this to get a certain customer.
func getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post request recieved.")

}
