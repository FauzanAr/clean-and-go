package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/FauzanAr/clean-and-go/modules/product/domain"
	product_handler "github.com/FauzanAr/clean-and-go/modules/product/handler"
	"github.com/FauzanAr/clean-and-go/modules/product/repository"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	PORT := ":9000"

	// Database
	dbUser := "fauzan"
	dbPass := "passwordlocal123"
	dbHost := "localhost"
	dbPort := "3306"
	dbName := "watchDB"
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")

	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)

	if err != nil {
		log.Fatal(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	
	// Declaration Modules
	productRepo := repository.NewProductRepositoryMysql(dbConn)
	productDomain := domain.NewProductDomain(productRepo)
	productHandler := product_handler.NewProdutHandler(productDomain)
	
	// Route	
	mux := http.NewServeMux()
	mux.HandleFunc("/product/v1", productHandler.Product)

	http.ListenAndServe(PORT, mux)
}