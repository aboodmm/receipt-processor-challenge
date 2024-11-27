package models

var ReceiptStore map[string]Receipt

type Receipt struct {
	Retailer     string  `json:"retailer"`
	PurchaseDate string  `json:"purchaseDate"`
	PurchaseTime string  `json:"purchaseTime"`
	Total        string  `json:"total"`
	Items        []items `json:"items"`
}

type items struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}
