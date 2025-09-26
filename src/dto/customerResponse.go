package dto

type CustomerResponse struct {
	Id          string `json:"customer_id"`
	Name        string `json:"full-name"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	DateofBirth string `json:"date_of_birth"`
	Status      string `json:"status"`
}