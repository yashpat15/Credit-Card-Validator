package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server starting on port 8080...")

	// Serve static files from the 'static' directory
	fs := http.FileServer(http.Dir("../static"))
	http.Handle("/", fs) // Serve static files

	// Setup route and handler for the API
	http.HandleFunc("/validate", validateCardHandler)

	// Start the server on port 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
