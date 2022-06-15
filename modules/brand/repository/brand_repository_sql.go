package brand_repository

import (
	"context"
	"database/sql"

	"github.com/FauzanAr/clean-and-go/helpers/logger"
	"github.com/FauzanAr/clean-and-go/helpers/response"
	"github.com/FauzanAr/clean-and-go/modules/brand"
)

type BrandRepositoryMysql struct {
	DB *sql.DB
}

func NewBrandRepositoryMysql(db *sql.DB) brand.Repository {
	return &BrandRepositoryMysql{
		DB: db,
	}
}

func (r *BrandRepositoryMysql) Fetch(ctx context.Context, query string, args ...interface{}) (error, []*brand.Entity) {
	loc := "[BrandRepository-Fetch]"
	stmt, err := r.DB.PrepareContext(ctx, query)
	defer stmt.Close()

	if err != nil {
		logger.ErrorLogger.Println(loc + err.Error())
		return response.InternalServerErr(err.Error()), nil
	}

	brands, err := r.DB.QueryContext(ctx, query, args...)
	defer brands.Close()

	if err != nil {
		logger.ErrorLogger.Println(loc + err.Error())
		return response.InternalServerErr(err.Error()), nil
	}

	results := make([]*brand.Entity, 0)

	for brands.Next() {
		tmp := brand.Entity{}
		err := brands.Scan(
			&tmp.ID,
			&tmp.Name,
			&tmp.Description,
			&tmp.CreatedAt,
			&tmp.UpdatedAt,
		)

		if err != nil {
			logger.ErrorLogger.Println(loc + err.Error())
			return response.InternalServerErr(err.Error()), nil
		}

		results = append(results, &tmp)
	}

	logger.InfoLogger.Println(loc + "Successfully get data")
	return nil, results
}

func (r *BrandRepositoryMysql) GetAll(ctx context.Context) (error, []*brand.Entity) {
	loc := "[BrandRepository-GetAll]"
	query := `SELECT id, name, description, created_at, updated_at FROM brands`
	err, data := r.Fetch(ctx, query)

	if err != nil {
		return err, nil
	}

	logger.InfoLogger.Println(loc + "Successfully get data")
	return nil, data
}