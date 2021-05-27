package service

import (
	"time"

	"github.com/sdoshi579/go-practice/domain"
	"github.com/sdoshi579/go-practice/dto"
	"github.com/sdoshi579/go-practice/errs"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
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

func NewAccountService(repo domain.AccountRepositoryDb) DefaultAccountService {
	return DefaultAccountService{repo: repo}
} 