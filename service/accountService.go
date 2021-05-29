package service

import (
	"time"

	"github.com/sdoshi579/go-practice/domain"
	"github.com/sdoshi579/go-practice/dto"
	"github.com/sdoshi579/go-practice/errs"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	NewTransaction(dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepositoryDb
}

func (as DefaultAccountService) NewAccount(r dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	account, err := as.repo.Save(domain.Account{
		CustomerId: r.CustomerId,
		Type: r.AccountType,
		Amount: r.Amount,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		Status: "1",
	})
	if err != nil {
		return nil, err
	}

	return &dto.NewAccountResponse{AccountId: account.AccountId}, nil
}

func (as DefaultAccountService) NewTransaction(r dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	account, err := as.repo.Transact(domain.Transact{
		AccountId: r.AccountId,
		Type: r.Type,
		Amount: r.Amount,
		Date: time.Now().Format("2006-01-02 15:04:05"),
		CustomerId: r.CustomerId,
	})
	if err != nil {
		return nil, err
	}

	return &dto.TransactionResponse{Id: account.Id, Amount: account.Amount}, nil
}

func NewAccountService(repo domain.AccountRepositoryDb) DefaultAccountService {
	return DefaultAccountService{repo: repo}
}

