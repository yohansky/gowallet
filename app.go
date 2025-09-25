package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mod/src/domain"
	"go.mod/src/handlers"
	"go.mod/src/service"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/banking")
	if err != nil {
		panic(`Failed create DB handle`+ err.Error())
	} 

	if err := db.Ping(); err != nil {
		panic("Failed to connect to Database: "+ err.Error())
	}
		fmt.Println("Connected to DB")
	
	db.SetConnMaxLifetime(time.Minute *3)
	db.SetMaxOpenConns(10)
db.SetMaxIdleConns(10)

	return db
}

func Start() {
	// mux := http.NewServeMux()
	router := mux.NewRouter()

	db := InitDB()
	defer db.Close()
	// wiring
	// ch := handlers.CustomerHandlers{Service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	customerRepo := domain.NewCustomerRepositoryDb(db)
	ch := handlers.CustomerHandlers{Service: service.NewCustomerService(customerRepo)}

	// define routes
	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", ch.GetCustomer).Methods(http.MethodGet)



	//  starting server
	log.Fatal(http.ListenAndServe(":8080", router))
}