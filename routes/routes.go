package routes

import (
	"github.com/aboodmm/receipt-processor/controllers"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/receipts/process", controllers.ProcessReceipt).Methods("POST")
	router.HandleFunc("/receipts/{id}/points", controllers.HandlePoints).Methods("GET")

	return router
}
