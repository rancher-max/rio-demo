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
	time.Sleep(101 * time.Millisecond)
	fmt.Fprintln(w, "Hi there, I'm running in Rio:rancher-max:forcaleb")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
