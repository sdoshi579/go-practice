package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sdoshi579/go-practice/service"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	
	customers, err := ch.service.GetAllCustomer()
	
	if err != nil {
		decorateResponse(w, err.Code, err.AsMessage())
		return
	} 

	decorateResponse(w, http.StatusOK, customers)
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	customers, err := ch.service.GetCustomer(vars["customer_id"])

	if err != nil {
		decorateResponse(w, err.Code, err.AsMessage())
	} else {
		decorateResponse(w, http.StatusOK, customers)
	}
}

func decorateResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}