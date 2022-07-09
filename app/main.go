package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc(pattern="status/check", statusCheckHandler)
}

func statusCheckHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("OK")
}