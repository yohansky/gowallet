package domain

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb ) FindAll() ([]Customer, error) {
	// findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
	findAllSql := "select * from customers"

	rows, err := d.client.Query(findAllSql)
	if err != nil {
		log.Println("Error while querying customer table " + err.Error())
		return nil, err
	}
	// defer rows.Close()

	customers := []Customer{}

	for rows.Next() {
		var c Customer
		rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
		customers = append(customers, c)
	}
	return customers, nil
}

func (d CustomerRepositoryDb ) ById(id string) (*Customer, error) {
	customerSql := "select * from customers where customer_id = ?"

	row := d.client.QueryRow(customerSql, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
	if err != nil {
		log.Println("Error while scanning customer " + err.Error())
		return nil, err
	}
	return &c, nil
}

func NewCustomerRepositoryDb(dbClient *sql.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{client: dbClient}
}