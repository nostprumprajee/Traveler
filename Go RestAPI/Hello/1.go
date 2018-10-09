package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/home", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "Hello World")
	})
	http.ListenAndServe(":8000", nil)
}
