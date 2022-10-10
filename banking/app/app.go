package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	mux := mux.NewRouter()

	// function handlers.
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getAllCustomers)

	// failure to start server.
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
