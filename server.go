package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

var counter int
var mutex = &sync.Mutex{}

func main() {

	err := godotenv.Load("local.env")

	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	CERT_PATH := os.Getenv("CERT_PATH")
	KEY_PATH := os.Getenv("KEY_PATH")

	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/count", counterHandler)

	http.HandleFunc("/hi", greetingHandler)

	log.Fatal(http.ListenAndServeTLS(":8080", CERT_PATH, KEY_PATH, nil))
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
