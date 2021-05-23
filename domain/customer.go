package domain

import "github.com/sdoshi579/go-practice/errs"

type Customer struct {
	Id			string	`db:"customer_id"`
	Name		string
	City 		string
	Zipcode 	string
	DateOfBirth string	`db:"date_of_birth"`
	Status 		string
}

type CustomerRepository interface {
	FindAll(string) ([]Customer, *errs.AppError)
	FindById(string) (*Customer, *errs.AppError)
}
