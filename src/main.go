package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Define routes and handler functions
	http.HandleFunc("/receipts/process", processHandler)
	http.HandleFunc("/receipts/{id}/points", pointsHandler)

	port := 8080
	fmt.Printf("Server running at http://localhost:%d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
