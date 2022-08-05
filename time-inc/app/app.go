package app

import (
	"log"
	"net/http"
)

func Start() {

	// endpoint handlers:
	http.HandleFunc("/api/time", timeHandler)
	http.ListenAndServe("localhost:8080", nil)

	// Failure to start the server.
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

}
