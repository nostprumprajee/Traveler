package main

import (
	"flag"
	"log"
	"module"
	"net/http"
)

var addr = flag.String("addr", "127.0.0.1:1718", "Rest Service")

func makeHandler(fn func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "{\"code\":500, \"desc\":\""+r.Method+"not allowed.\"}", http.StatusMethodNotAllowed)
		} else {
			fn(w, r)
		}
	}
}

// func handleRequests() {
// 	http.HandleFunc("/", index)
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

func main() {
	flag.Parse()
	log.Println("Start Server...")
	http.HandleFunc("/rest/parser", makeHandler(module.JSONParserHandler))
	http.ListenAndServe(*addr, nil)
}
