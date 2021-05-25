package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/sdoshi579/go-practice/domain"
	"github.com/sdoshi579/go-practice/service"
)

func Start() {
	
	router := mux.NewRouter()
	
	//ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	// starting a server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}