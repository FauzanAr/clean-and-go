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
	return nil, nil
}

func (r *TransactionRepositoryMysql) GetByEmail(ctx context.Context, email string) (error, []*transaction.Entity) {
	return nil, nil
}