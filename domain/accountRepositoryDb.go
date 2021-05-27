package domain

import (
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

func NewAccountRepositoryDb(client *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{client: client}
}