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

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}