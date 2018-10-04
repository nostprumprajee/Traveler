package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(responseWriter, "Welcome to Nost Restful API")
	fmt.Println("Access to index.")
}

func handleRequests() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
