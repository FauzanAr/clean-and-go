package repository

import (
	"context"
	"database/sql"

	"github.com/FauzanAr/clean-and-go/modules/product"
)

type ProductRepositoryMysql struct {
	DB *sql.DB
}

func NewProductRepositoryMysql(db *sql.DB) product.Repository {
	return &ProductRepositoryMysql{
		DB: db,
	}
}

func (r *ProductRepositoryMysql) Fetch(ctx context.Context) error {

	return nil
}
