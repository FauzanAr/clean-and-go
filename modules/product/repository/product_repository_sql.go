package repository

import (
	"context"
	"database/sql"

	"github.com/FauzanAr/clean-and-go/helpers/logger"
	"github.com/FauzanAr/clean-and-go/helpers/response"
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

func (r *ProductRepositoryMysql) Fetch(ctx context.Context,query string, args ...interface{}) (error, []*product.Entity) {
	loc := "[ProductRepository-Fetch]"
	stmt, err := r.DB.PrepareContext(ctx, query)
	defer stmt.Close()

	if err != nil {
		logger.ErrorLogger.Println(loc + err.Error())
		resErr := response.BadRequest(err.Error())
		return resErr, nil
	}

	products, err := r.DB.QueryContext(ctx, query, args...)
	defer products.Close()

	if err != nil {
		logger.ErrorLogger.Println(loc + err.Error())
		resErr := response.BadRequest(err.Error())
		return resErr, nil
	}

	results := make([]*product.Entity, 0)

	for products.Next() {
		tmp := product.Entity{}
		err := products.Scan(
			&tmp.ID,
			&tmp.BrandId,
			&tmp.Name,
			&tmp.Description,
			&tmp.Price,
			&tmp.Stock,
			&tmp.CreatedAt,
			&tmp.UpdatedAt,
		)

		if err != nil {
			logger.ErrorLogger.Println(loc + err.Error())
			resErr := response.BadRequest(err.Error())
			return resErr, nil
		}

		results = append(results, &tmp)
	}

	logger.InfoLogger.Println(loc + "Successfully get data")
	return nil, results
}

func (r *ProductRepositoryMysql) GetAll(ctx context.Context) (error, []*product.Entity) {
	loc := "[ProductRepository-GetAll]"
	query := `SELECT id, brand_id, name, description, price, stock, created_at, updated_at FROM products`
	err, data := r.Fetch(ctx, query)

	if err != nil {
		return err, nil
	}

	logger.InfoLogger.Println(loc + "Successfully get data")
	return nil, data
}

func (r *ProductRepositoryMysql) GetOne(ctx context.Context, id int) (error, *product.Entity) {
	loc := "[ProductRepository-GetOne]"
	query := `SELECT id, brand_id, name, description, price, stock, created_at, updated_at FROM products WHERE id = ?`
	err, data := r.Fetch(ctx, query, id)

	if err != nil {
		return err, nil
	}

	if len(data) > 0 {
		return nil, data[0]
	}

	logger.InfoLogger.Println(loc + "Successfully get data")
	return response.NotFound(nil), nil
}

func (r *ProductRepositoryMysql) Insert(ctx context.Context) error {
	return nil
}
