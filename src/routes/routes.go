package routes

import (
	"github.com/gorilla/mux" // One of the popular routing frameworks
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/receipts/process", ProcessReceipt).Methods("GET")
	router.HandleFunc("/receipts/{id}/points", HandlePoints).Methods("GET")

	return router
}
