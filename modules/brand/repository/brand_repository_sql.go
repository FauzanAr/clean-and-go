package brand_repository

import (
	"context"
	"database/sql"
	"time"

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

func (r *BrandRepositoryMysql) Insert(ctx context.Context, b *brand.Entity) error {
	b.CreatedAt = int(time.Now().Unix())
	b.UpdatedAt = int(time.Now().Unix())
	loc := "[BrandRepository-Insert]"
	query := `INSERT brands SET name=?, description=?, created_at=?, updated_at=?`
	stmt, err := r.DB.PrepareContext(ctx, query)
	defer stmt.Close()

	if err != nil {
		logger.ErrorLogger.Println(loc + err.Error())
		return response.InternalServerErr(err.Error())
	}

	_, errExec := r.DB.ExecContext(ctx, query, b.Name, b.Description, b.CreatedAt, b.UpdatedAt)

	if errExec != nil {
		logger.ErrorLogger.Println(loc + errExec.Error())
		return response.InternalServerErr(errExec.Error())
	}

	return nil
}

func (r *BrandRepositoryMysql) GetByName(ctx context.Context, name string) (error, *brand.Entity) {
	loc := "[BrandRepository-GetByName]"
	query := `SELECT id, name, description, created_at, updated_at FROM brands WHERE name = ?`
	err, data := r.Fetch(ctx, query, name)

	if err != nil {
		logger.ErrorLogger.Println(loc + err.Error())
		return err, nil
	}

	if len(data) > 0 {
		logger.InfoLogger.Println(loc + "Successfully get data")
		return nil, data[0]
	}

	logger.InfoLogger.Println(loc + "Successfully get data, got nil")
	return nil, nil
}

func (r *BrandRepositoryMysql) GetById(ctx context.Context, id int) (error, *brand.Entity) {
	loc := "[BrandRepository-GetById]"
	query := `SELECT id, name, description, created_at, updated_at FROM brands WHERE id = ?`
	err, data := r.Fetch(ctx, query, id)

	if err != nil {
		logger.ErrorLogger.Println(loc + err.Error())
		return err, nil
	}

	if len(data) > 0 {
		logger.InfoLogger.Println(loc + "Successfully get data")
		return nil, data[0]
	}

	logger.InfoLogger.Println(loc + "Successfully get data, got nil")
	return nil, nil
}