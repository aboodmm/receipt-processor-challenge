package main

import (
	"fmt"
	"net/http"
)

func main() {
	receiptStore := make(map[string]Receipt)

	port := 8080
	fmt.Printf("Server running at http://localhost:%d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
