package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", addHandler)
	http.ListenAndServe(":8080", nil)
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World")
}
