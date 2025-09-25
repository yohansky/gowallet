package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.mod/src/service"
)

type CustomerHandlers struct {
	Service service.CustomerServices
}

func (ch *CustomerHandlers) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	// customers := []models.Customer{
	// 	{Name: "Yohanes", City: "Jakarta", Zipcode: "110075"},
	// 	{Name: "Huber", City: "Jakarta", Zipcode: "110075"},
	// }

	customers, _ := ch.Service.GetAllCustomer()
	
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
	// xml.NewEncoder(w).Encode(customers)
}

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])
}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post request received")
}