package main

import (
	"log"
	"net/http"

	"go.mod/src/handlers"
)

func Start() {
	mux := http.NewServeMux()
	// define routes
	mux.HandleFunc("/greet", handlers.Greeting)
	mux.HandleFunc("/customers", handlers.GetAllCustomers)

	//  starting server
	log.Fatal(http.ListenAndServe(":8080", mux))
}