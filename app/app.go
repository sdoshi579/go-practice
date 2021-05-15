package app

import (
	"log"
	"net/http"
)

func Start() {
	
	mux := http.NewServeMux()
	//create routes
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getAllCustomers)

	// starting a server
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}