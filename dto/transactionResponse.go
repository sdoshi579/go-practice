package dto

type TransactionResponse struct {
	Id string `json:"transaction_id"`
	Amount float64 `json:"amount"`
}