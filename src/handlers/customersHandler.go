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
	customers, _ := ch.Service.GetAllCustomer()
	
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (ch *CustomerHandlers ) GetCustomer(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	// mengambil customer_id dai url
	id := vars["customer_id"]
	
	customer, err := ch.Service.GetCustomer(id)
	if err != nil {
		w.WriteHeader(err.Code)
		fmt.Fprint(w, err.Message)
		} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}
}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Post request received")
}