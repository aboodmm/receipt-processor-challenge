package controllers

import "net/http"

func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func HandlePoints(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
