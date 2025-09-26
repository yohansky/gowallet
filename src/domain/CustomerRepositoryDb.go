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

func (d CustomerRepositoryDb ) FindAll(status string) ([]Customer, *errs.AppError) {
	var rows *sql.Rows
	var err error

	if status == "" {
		findAllSql := "select * from customers"
		rows, err = d.client.Query(findAllSql)
	} else {
		findAllSql := "select * from customers where status = ?"
		rows, err = d.client.Query(findAllSql, status)
			
	}

	if err != nil {
		log.Println("Error while querying customer table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database errorr")
	}
	// defer rows.Close()

	customers := []Customer{}

	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
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