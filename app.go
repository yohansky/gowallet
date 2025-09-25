package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mod/src/handlers"
)

func Start() {
	// mux := http.NewServeMux()
	router := mux.NewRouter()

	// define routes
	router.HandleFunc("/greet", handlers.Greeting).Methods(http.MethodGet)
	router.HandleFunc("/customers", handlers.GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", handlers.CreateCustomer).Methods(http.MethodPost)

	router.HandleFunc("/customers/{customer_id:[0-9]+}", handlers.GetCustomer).Methods(http.MethodGet)

	//  starting server
	log.Fatal(http.ListenAndServe(":8080", router))
}