package domain

import "github.com/sdoshi579/go-practice/errs"

type Account struct {
	AccountId 	string
	CustomerId 	string
	OpeningDate string
	Type 		string
	Amount 		float64
	Status 		string
}

type Transact struct {
	Id			string
	AccountId 	string
	CustomerId 	string
	Date 		string
	Type 		string
	Amount 		float64
	NewAmount	float64
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	Transact(Transact) (*Transact, *errs.AppError)
	GetCustomerAccountAmount(string, string) (*float64, *errs.AppError)
}