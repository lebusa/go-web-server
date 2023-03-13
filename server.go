package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}

func main() {
	http.HandleFunc("/", rootHandler)

	http.HandleFunc("/count", counterHandler)

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

// echo string
func echo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "greetings")
}

// increment counter
func counterHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock() // lock the counter
	counter++
	fmt.Fprintf(w, strconv.Itoa(counter))
	mutex.Unlock() // realease the lock
}
