package dto

type TransactionRequest struct {
	CustomerId 	string 	`json:"customer_id"`
	AccountId 	string	`json:"account_id"`
	Amount 		float64	`json:"amount"`
	Type		string 	`json:"type"`
}