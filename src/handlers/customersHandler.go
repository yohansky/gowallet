package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mod/src/service"
)

type CustomerHandlers struct {
	Service service.CustomerServices
}

func (ch *CustomerHandlers) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	// jadikan status sebagai parameter
	customers, err := ch.Service.GetAllCustomer(status)
	
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customers)
	}
}

func (ch *CustomerHandlers ) GetCustomer(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	// mengambil customer_id dai url
	id := vars["customer_id"]
	
	customer, err := ch.Service.GetCustomer(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}