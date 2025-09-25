package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go.mod/src/models"
)

func Greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []models.Customer{
		{Name: "Yohanes", City: "Jakarta", Zipcode: "110075"},
		{Name: "Huber", City: "Jakarta", Zipcode: "110075"},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
	// xml.NewEncoder(w).Encode(customers)
}