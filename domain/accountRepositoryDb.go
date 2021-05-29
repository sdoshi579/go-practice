package domain

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/sdoshi579/go-practice/errs"
	"github.com/sdoshi579/go-practice/logger"
	"go.uber.org/zap"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) Save(account Account) (*Account, *errs.AppError) {
	sqlInsert := "Insert into accounts (customer_id, opening_date, account_type, amount, status) values (?, ?, ?, ?, ?)"

	result, err := d.client.Exec(sqlInsert, account.CustomerId, account.OpeningDate, account.Type, account.Amount, account.Status)

	if err != nil {
		logger.Error("Error in inserting account", zap.Error(err))
		return nil, errs.NewInternalServerError("Error while inserting the account")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error in inserting account", zap.Error(err))
		return nil, errs.NewInternalServerError("Error while inserting the account")
	}
	account.AccountId = strconv.FormatInt(id, 10)
	return &account, nil
}

func (d AccountRepositoryDb) Transact(transaction Transact) (*Transact, *errs.AppError) {
	validateErr := d.validateRequest(transaction)
	if validateErr != nil {
		logger.Error("Validation failed", zap.Any("error",validateErr))
		return nil, validateErr
	}

	tx, _ := d.client.Begin()
	
	sqlInsert := "Insert into transactions (account_id, transaction_date, transaction_type, amount) values (?, ?, ?, ?)"

	result, err := tx.Exec(sqlInsert, transaction.AccountId, transaction.Date, transaction.Type, transaction.Amount)

	if err != nil {
		logger.Error("Error in inserting account", zap.Error(err))
		return nil, errs.NewInternalServerError("Error while inserting the account")
	}

	if transaction.Type == "saving" {
		_, err = tx.Exec(`Update accounts set amount = amount + ? where account_id = ?`, transaction.Amount, transaction.AccountId)
	} else {
		_, err = tx.Exec(`Update accounts set amount = amount - ? where account_id = ?`, transaction.Amount, transaction.AccountId)
	} 

	if err != nil {
		tx.Rollback()
		logger.Error("Error in transacting from account", zap.Error(err))
		return nil, errs.NewInternalServerError("Error while transacting from the account")
	}
	tx.Commit()
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error in inserting account", zap.Error(err))
		return nil, errs.NewInternalServerError("Error while inserting the account")
	}
	transaction.Id = strconv.FormatInt(id, 10)
	return &transaction, nil
}

func (d AccountRepositoryDb) validateRequest(r Transact) *errs.AppError {
	amount, err := d.GetCustomerAccountAmount(r.CustomerId, r.AccountId)

	if err != nil {
		return err
	}

	if r.Amount > *amount  && r.Type == "withdrawal" {
		return errs.NewInternalServerError("Amount greater then the bank amount")
	}
	return nil

}

func (d AccountRepositoryDb) GetCustomerAccountAmount(customerId string, accountId string) (*float64, *errs.AppError) {

	find := fmt.Sprintf("SELECT amount FROM accounts WHERE account_id=%s AND customer_id=%s", accountId, customerId)
	
	var c float64
	err := d.client.Get(&c, find)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		}
		logger.Error("Error in scanning customers for id " + err.Error())
		return nil, errs.NewInternalServerError("unexpected database error")
	}
	return &c, nil
}

func NewAccountRepositoryDb(client *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{client: client}
}