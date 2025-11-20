package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Write to the response
	fmt.Fprintf(w, "Hello, Web!")
}

func main() {
	// 1. Register a handler function for the "/" route
	http.HandleFunc("/", helloHandler)

	// 2. Start the server on port 8080
	fmt.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
