package domain

import (
	"github.com/sdoshi579/go-practice/dto"
	"github.com/sdoshi579/go-practice/errs"
)

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

func (c Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id: 			c.Id,
		Name: 			c.Name,
		Status: 		c.getStatusAsText(),
		DateOfBirth: 	c.DateOfBirth,
		Zipcode: 		c.Zipcode,
		City: 			c.City,
	}
}

func (c Customer) getStatusAsText() string {
	if c.Status == "0" {
		return "inactive"
	}
	return "active"
}