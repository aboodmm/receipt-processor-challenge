package controllers

import (
	"encoding/json"
	"github.com/aboodmm/receipt-processor/cache"
	"math"
	"net/http"
	"strconv"
	"strings"

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
	points := 0

	cacheO := cache.GetCache()
	receipt, exists := cacheO[id]
	if exists {
		points += countAlpha(receipt.Retailer)
		points += countRound(receipt.Total)
		points += countItems(len(receipt.Items))
		points += countMultandPrice(receipt)
	} else {
		log.Error().Msgf("Error: Receipt with id %s doesn't exist!", id)
	}

	response := &models.PointResponse{
		Points: points,
	}
	json.NewEncoder(w).Encode(response)
}

func countAlpha(name string) int {
	return len(name)
}

func countRound(total string) int {
	a, err := strconv.ParseFloat(total, 32)
	if err != nil {
		log.Error().Err(err).Msgf("Error: Format error")
	}
	if a == math.Trunc(a) {
		return 50
	}
	return 0
}

func countMultiple(total string) int {
	a, err := strconv.ParseFloat(total, 32)
	if err != nil {
		log.Error().Err(err).Msgf("Error: Format error")
	}
	b := a / float64(4)

	if math.Floor(b) == 4 {
		return 25
	}
	return 0
}

func countItems(count int) int {
	a := math.Floor(float64(count) / float64(2))
	b := int(5 * a)
	return b
}

func countMultandPrice(receipt models.Receipt) int {
	c := 0
	for i := 0; i < len(receipt.Items); i++ {
		a := len(strings.TrimSpace(receipt.Items[i].ShortDescription))
		if a%3 == 0 {
			b, _ := strconv.ParseFloat(receipt.Items[i].Price, 32)
			d := math.Ceil(0.2 * b)
			c += int(d)
		}
	}
	return c
}
