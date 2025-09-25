package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mod/src/domain"
	"go.mod/src/handlers"
	"go.mod/src/service"
)

func Start() {
	// mux := http.NewServeMux()
	router := mux.NewRouter()

	// wiring
	ch := handlers.CustomerHandlers{
		Service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	// define routes

	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)



	//  starting server
	log.Fatal(http.ListenAndServe(":8080", router))
}