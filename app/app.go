package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sdoshi579/go-practice/domain"
	"github.com/sdoshi579/go-practice/service"
)

func Start() {
	
	router := mux.NewRouter()
	
	//ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	// starting a server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}