package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/FauzanAr/clean-and-go/helpers/logger"
	"github.com/FauzanAr/clean-and-go/modules/product/domain"
	"github.com/FauzanAr/clean-and-go/modules/product/handler"
	"github.com/FauzanAr/clean-and-go/modules/product/repository"
	"github.com/FauzanAr/clean-and-go/modules/brand/domain"
	"github.com/FauzanAr/clean-and-go/modules/brand/handler"
	"github.com/FauzanAr/clean-and-go/modules/brand/repository"

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
	brandRepo := brand_repository.NewBrandRepositoryMysql(dbConn)
	brandDomain := brand_domain.NewBrandDomain(brandRepo)
	brandHandler := brand_handler.NewBrandHandler(brandDomain)
	
	productRepo := product_repository.NewProductRepositoryMysql(dbConn)
	productDomain := product_domain.NewProductDomain(productRepo, brandRepo)
	productHandler := product_handler.NewProdutHandler(productDomain)
	
	// Route	
	mux := http.NewServeMux()

	mux.HandleFunc("/product/v1", productHandler.Product)
	mux.HandleFunc("/product/v1/brand", productHandler.GetByBrand)

	mux.HandleFunc("/brand/v1",  brandHandler.Brand)

	errServer := http.ListenAndServe(PORT, mux)
	
	if err != nil {
		logger.ErrorLogger.Fatal(errServer.Error())
	}
}