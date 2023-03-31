package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequest() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my API!")
	})

	http.ListenAndServe(":8080", router)
}

func main() {
	handleRequest()
}
