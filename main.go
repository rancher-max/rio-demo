package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func fooHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request to FOO")
	fmt.Fprintln(w, "Hi there, I'm running in Rio:FOO port 8001")
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	// Arbitrary sleep so that we can demonstrate autoscaler
	log.Println("Received request to BAR")
	time.Sleep(101 * time.Millisecond)
	fmt.Fprintln(w, "Hi there, I'm running in Rio:BAR port 8002")
}

func handler(w http.ResponseWriter, r *http.Request) {
	// Arbitrary sleep so that we can demonstrate autoscaler
	log.Println("Basic logging -- request received")
	time.Sleep(501 * time.Millisecond)
	fmt.Fprintln(w, "Hi there, I'm running in Rio:handler func on port 8080")
}

func main() {
	go func() {
		http.HandleFunc("/", fooHandler)
		log.Fatal(http.ListenAndServe(":8001", nil))
	}()

	go func() {
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	//the last call is outside goroutine to avoid that program just exit
	http.HandleFunc("/", fooHandler)
	log.Fatal(http.ListenAndServe(":8002", nil))
}
