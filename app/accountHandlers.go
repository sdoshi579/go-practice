package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sdoshi579/go-practice/dto"
	"github.com/sdoshi579/go-practice/service"
)

type AccountHandlers struct {
	service service.AccountService
}

func (ah *AccountHandlers) createAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var accountRequest dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&accountRequest)
	if err != nil {
		decorateResponse(w, http.StatusBadRequest, err)
		return
	}
	accountRequest.CustomerId = vars["customer_id"]
	accountResponse, appError := ah.service.NewAccount(accountRequest)
	if appError != nil {
		decorateResponse(w, appError.Code, appError.Message)
		return
	}
	decorateResponse(w, http.StatusCreated, accountResponse)
}

func (ah *AccountHandlers) createTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var transactionRequest dto.TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&transactionRequest)
	if err != nil {
		decorateResponse(w, http.StatusBadRequest, err)
		return
	}
	transactionRequest.CustomerId = vars["customer_id"]
	transactionRequest.AccountId = vars["account_id"]
	accountResponse, appError := ah.service.NewTransaction(transactionRequest)
	if appError != nil {
		decorateResponse(w, appError.Code, appError.Message)
		return
	}
	decorateResponse(w, http.StatusCreated, accountResponse)
}

