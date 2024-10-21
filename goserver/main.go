// server.go
package main

import (
	"fmt"
	"net/http"
)

func greetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, visitor! You've requested: %s\n", r.URL.Path)
}

func main() {
	http.HandleFunc("/", greetHandler)

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
