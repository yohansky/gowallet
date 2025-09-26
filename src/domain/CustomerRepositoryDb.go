package domain

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"go.mod/src/errs"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb ) FindAll() ([]Customer, *errs.AppError) {
	// findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
	findAllSql := "select * from customers"

	rows, err := d.client.Query(findAllSql)
	if err != nil {
		log.Println("Error while querying customer table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database errorr")
	}
	// defer rows.Close()

	customers := []Customer{}

	for rows.Next() {
		var c Customer
		rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
		if err != nil {
		log.Println("Error while querying customer table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database errorr")
	}
		customers = append(customers, c)
	}
	return customers, nil
}

func (d CustomerRepositoryDb ) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select * from customers where customer_id = ?"

	row := d.client.QueryRow(customerSql, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows{
			return nil, errs.NewNotFoundError("Customer not found")
		} else {

			log.Println("Error while scanning customer " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDb(dbClient *sql.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{client: dbClient}
}