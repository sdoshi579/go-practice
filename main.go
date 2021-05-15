package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	//create routes
	http.HandleFunc("/greet", greet)

	// starting a server
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World !!!")
}