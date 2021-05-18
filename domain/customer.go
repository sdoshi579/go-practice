package domain

import "github.com/sdoshi579/go-practice/errs"

type Customer struct {
	Id			string
	Name		string
	City 		string
	Zipcode 	string
	DateOfBirth string
	Status 		string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	FindById(string) (*Customer, *errs.AppError)
}
