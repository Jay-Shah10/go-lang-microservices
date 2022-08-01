package app

import (
	"log"
	"net/http"
)

func Start() {
	// function handlers.
	http.HandleFunc("/status/check", statusCheckHandler)
	http.ListenAndServe("localhost:8000", nil)

	// failure to start server.
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
