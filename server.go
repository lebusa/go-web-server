package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", rootHandler)

	http.HandleFunc("/hi", greetingHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// root path handler
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

// handler to request directed to /lumela path
func greetingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi back at you!")
}
