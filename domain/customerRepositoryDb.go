package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sdoshi579/go-practice/errs"
	"github.com/sdoshi579/go-practice/logger"
)


type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	

	findAllSql := "select * from customers"

	if status != "" {
		findAllSql += " where status = " + status
	}
	rows, err := d.client.Query(findAllSql)
	if err != nil {
		logger.Error("Error in fetching customers " + err.Error())
		return nil, errs.NewInternalServerError("uexpected error")
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.DateOfBirth, &c.City, &c.Zipcode, &c.Status)
		if err != nil {
			logger.Error("Error in scanning customers " + err.Error())
			return nil, errs.NewInternalServerError("unexpected error")
		}
		customers = append(customers, c)
	}

	return customers, nil
}

func (d CustomerRepositoryDb) FindById(id string) (*Customer, *errs.AppError) {

	findAllSql := "select * from customers where customer_id = " + id

	rows := d.client.QueryRow(findAllSql)
	
	var c Customer
	err := rows.Scan(&c.Id, &c.Name, &c.DateOfBirth, &c.City, &c.Zipcode, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		}
		logger.Error("Error in scanning customers for id " + err.Error())
		return nil, errs.NewInternalServerError("unexpected database error")
	}
	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client: client}
}