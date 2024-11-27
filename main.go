package main

import (
	"fmt"
	"net/http"

	"github.com/aboodmm/receipt-processor/models"
	"github.com/aboodmm/receipt-processor/routes"
	"github.com/rs/zerolog/log"
)

func main() {

	for k, v := range models.ReceiptStore {
		fmt.Println(k, v)
	}
	router := routes.SetupRoutes()
	log.Info().Msg("Initializing router")

	port := 8080
	fmt.Printf("Server running at http://localhost:%d\n", port)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), router)

	if err != nil {
		log.Fatal().Err(err).Msgf("Fatal Error: Failed to listen on port %d", port)
	}

}
