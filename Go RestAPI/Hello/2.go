package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Print("Starting the service")

	http.HandleFunc("/home", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
