package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()

	// function handlers.
	router.HandleFunc("/greet", greet)
	router.HandleFunc("/customers", getAllCustomers)

	// failure to start server.
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
