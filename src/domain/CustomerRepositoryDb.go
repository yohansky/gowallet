package domain

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.mod/src/errs"
	"go.mod/src/logger"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb ) FindAll(status string) ([]Customer, *errs.AppError) {
	var err error
	customers := []Customer{}

	if status == "" {
		findAllSql := "select * from customers"
		err = d.client.Select(&customers, findAllSql)
		
		// rows, err = d.client.Query(findAllSql)
	} else {
		findAllSql := "select * from customers where status = ?"
		err = d.client.Select(&customers, findAllSql, status)
		// rows, err = d.client.Query(findAllSql, status)
			
	}

	if err != nil {
		logger.Error("Error while querying customer table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database errorr")
	}

	// err = sqlx.StructScan(rows, &customers)
	// if err != nil {
	// 	logger.Error("Error while querying customer table " + err.Error())
	// 	return nil, errs.NewUnexpectedError("Unexpected database errorr")
	// }
	return customers, nil
}

func (d CustomerRepositoryDb ) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "select * from customers where customer_id = ?"

	var c Customer
	err := d.client.Get(&c, customerSql, id)
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

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{client: dbClient}
}