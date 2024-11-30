package controllers

import (
	"encoding/json"
	"github.com/aboodmm/receipt-processor/cache"
	"net/http"

	"github.com/aboodmm/receipt-processor/models"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	id := uuid.New().String()
	decoder := json.NewDecoder(r.Body)

	var receipt models.Receipt
	err := decoder.Decode(&receipt)

	// Validation code -> deserialize into receipt struct
	if err != nil {
		log.Error().Err(err).Msgf("Error: Deserialization error")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cacheO := cache.GetCache()
	cacheO[id] = receipt

	response := &models.ReceiptResponse{
		Id: id,
	}
	json.NewEncoder(w).Encode(response)
}

func HandlePoints(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]

	cacheO := cache.GetCache()
	receipt, exists := cacheO[id]
	if exists {
		points := 0
		points += countAlpha(receipt.Retailer)
	} else {
		log.Error().Msgf("Error: Receipt with id %s doesn't exist!", id)
	}

}

func countAlpha(name string) int {
	return 0
}
