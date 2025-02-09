package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/oquga/koshelek/pkg/postgres"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Write "Hello, World!" to the response

	connStr := "user=user password=password dbname=mydb host=postgres_db port=5432 sslmode=disable"

	// Проверяем подключение к PostgreSQL
	var err error
	for i := 0; i < 10; i++ {
		err = postgres.CheckConnection(connStr)
		if err == nil {
			break
		}
		log.Printf("Attempt %d: PostgreSQL is not ready yet...", i+1)
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		fmt.Fprintf(w, "Some problems with communication to the server")
		log.Fatalf("Error: %v", err)
		return

	}

	fmt.Fprintf(w, "Connection with postgres is checked")
	return
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
