package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Arbitrary sleep so that we can demonstrate autoscaler
	log.Println("Basic logging -- request received")
	time.Sleep(2001 * time.Millisecond)
	fmt.Fprintln(w, "Hi there, I'm running in Rio:rancher-max:v3")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
