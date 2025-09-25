package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name string `json:"full_name"`
	City string `json:"city"`
	Zipcode string `json:"zip_code"`
}

func main() {
	// define routes
	http.HandleFunc("/greet", greeting)
	http.HandleFunc("/customers", getAllCustomers)

	//  starting server
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func greeting(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w,"Hello World!")
	}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
		customers := []Customer {
			{Name: "Yohanes", City: "Jakarta", Zipcode: "110075"},
			{Name: "Huber", City: "Jakarta", Zipcode: "110075"},
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}