package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Write "Hello, World!" to the response
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	// Register the helloHandler function to handle requests at the root URL "/"
	http.HandleFunc("/", helloHandler)

	// Start the web server on port 8080
	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
