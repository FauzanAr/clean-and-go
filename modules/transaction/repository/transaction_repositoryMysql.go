package transaction_repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/FauzanAr/clean-and-go/helpers/logger"
	"github.com/FauzanAr/clean-and-go/helpers/response"
	"github.com/FauzanAr/clean-and-go/modules/transaction"
)

type TransactionRepositoryMysql struct {
	DB *sql.DB
}

func NewTransactionRepositoryMysql(db *sql.DB) transaction.Repository {
	return &TransactionRepositoryMysql {
		DB: db,
	}
}

func (r *TransactionRepositoryMysql) Insert(ctx context.Context, t *transaction.Entity) error {
	t.CreatedAt = int(time.Now().Unix())
	t.UpdatedAt = int(time.Now().Unix())
	loc := "[TransactionRepository - Insert]"
	query := `INSERT transactions SET product_id=?, qty=?, total_price=?, email=?, created_at=?, updated_at=?`
	stmt, err := r.DB.PrepareContext(ctx, query)
	defer stmt.Close()

	if err != nil {
		logger.ErrorLogger.Println(loc + err.Error())
		return response.InternalServerErr(err.Error())
	}

	_, errExec := r.DB.ExecContext(ctx, query, t.ProductId, t.Qty, t.TotalPrice, t.Email, t.CreatedAt, t.UpdatedAt)

	if errExec != nil {
		logger.ErrorLogger.Println(loc + errExec.Error())
		return response.InternalServerErr(errExec.Error())
	}

	return nil
}

func (r *TransactionRepositoryMysql) Fetch(ctx context.Context, query string, args ...interface{}) (error, []*transaction.Entity) {
	loc := "[TransactionRepository-Fetch]"
	stmt, err := r.DB.PrepareContext(ctx, query)
	defer stmt.Close()

	if err != nil {
		logger.ErrorLogger.Println(loc + err.Error())
		return response.InternalServerErr(err.Error()), nil
	}

	transactions, err := r.DB.QueryContext(ctx, query, args...)
	defer transactions.Close()

	if err != nil {
		logger.ErrorLogger.Println(loc + err.Error())
		return response.InternalServerErr(err.Error()), nil
	}

	results := make([]*transaction.Entity, 0)

	for transactions.Next() {
		tmp := transaction.Entity{}
		err := transactions.Scan(
			&tmp.ID,
			&tmp.ProductId,
			&tmp.Qty,
			&tmp.TotalPrice,
			&tmp.Email,
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

func (r *TransactionRepositoryMysql) GetByEmail(ctx context.Context, email string) (error, []*transaction.Entity) {
	loc := "[TransactionRepository-GetByEmail]"
	query := `SELECT id, product_id, qty, total_price, email, created_at, updated_at FROM transactions WHERE email = ?`
	err, data := r.Fetch(ctx, query, email)

	if err != nil {
		return err, nil
	}

	logger.InfoLogger.Println(loc + "Successfully get data")
	return nil, data
}

func (r *TransactionRepositoryMysql) GetById(ctx context.Context, id int) (error, *transaction.Entity) {
	loc := "[TransactionRepository-GetById]"
	query := `SELECT id, product_id, qty, total_price, email, created_at, updated_at FROM transactions WHERE id = ?`
	err, data := r.Fetch(ctx, query, id)

	if err != nil {
		return err, nil
	}

	if len(data) > 0 {
		logger.InfoLogger.Println(loc + "Successfully get data")
		return nil, data[0]
	}

	logger.InfoLogger.Println(loc + "Successfully get data")
	return nil, nil
}