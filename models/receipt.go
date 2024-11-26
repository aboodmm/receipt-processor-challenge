package models

type receipt struct {
	Retailer     string              `json:"retailer"`
	PurchaseDate string              `json:"purchaseDate"`
	PurchaseTime string              `json:"purchaseTime"`
	Total        string              `json:"total"`
	Items        []map[string]string `json:"items"`
}

type ReceiptStore struct {
	Id string `json:"id`
	Receipt [map][string]string `json:"receipt"`
}
