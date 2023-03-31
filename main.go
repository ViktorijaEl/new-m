package main

import (
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to my API!")
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/", handleRequest)

    fmt.Println("Server listening on port 8080")
    http.ListenAndServe(":8080", router)
}
