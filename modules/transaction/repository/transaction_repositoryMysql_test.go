package transaction_repository

import (
	"time"
	"testing"
	"context"

	"github.com/FauzanAr/clean-and-go/modules/transaction"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var mockTransactions = []transaction.Entity {
	{ID: 1, ProductId: 1, Qty: 1, TotalPrice: 10000, Email: "email@example.com", CreatedAt: int(time.Now().Unix()), UpdatedAt: int(time.Now().Unix())},
	{ID: 2, ProductId: 2, Qty: 1, TotalPrice: 20000, Email: "e-mail@example.com", CreatedAt: int(time.Now().Unix()), UpdatedAt: int(time.Now().Unix())},
}

var mockTransaction = transaction.Entity {
	ID: 3,
	ProductId: 3,
	Qty: 1,
	TotalPrice: 30000,
	Email: "mail@example.com",
	CreatedAt: int(time.Now().Unix()),
	UpdatedAt: int(time.Now().Unix()),
}

func TestFetch(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("Error occured : %s", err.Error())
	}

	data := sqlmock.NewRows([]string{"id", "product_id", "qty", "total_price", "email", "created_at", "updated_at"})
	
	for _, t := range mockTransactions {
		data.AddRow(t.ID, t.ProductId, t.Qty, t.TotalPrice, t.Email, t.CreatedAt, t.UpdatedAt)
	}

	query := "SELECT id, product_id, qty, total_price, email, created_at, updated_at FROM transactions"
	
	mock.ExpectPrepare(query).ExpectQuery().WillReturnRows(data)
	
	r := NewTransactionRepositoryMysql(db)
	err, res := r.Fetch(context.TODO(), query)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestGetByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("Error occured : %s", err.Error())
	}

	data := sqlmock.NewRows([]string{"id", "product_id", "qty", "total_price", "email", "created_at", "updated_at"})
	
	for _, t := range mockTransactions {
		data.AddRow(t.ID, t.ProductId, t.Qty, t.TotalPrice, t.Email, t.CreatedAt, t.UpdatedAt)
	}

	query := "SELECT id, product_id, qty, total_price, email, created_at, updated_at FROM transactions WHERE email = \\?"
	
	mock.ExpectPrepare(query).ExpectQuery().WithArgs(mockTransactions[1].Email).WillReturnRows(data)
	
	r := NewTransactionRepositoryMysql(db)
	err, res := r.GetByEmail(context.TODO(), mockTransactions[1].Email)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestGetById(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("Error occured : %s", err.Error())
	}

	data := sqlmock.NewRows([]string{"id", "product_id", "qty", "total_price", "email", "created_at", "updated_at"})
	
	for _, t := range mockTransactions {
		data.AddRow(t.ID, t.ProductId, t.Qty, t.TotalPrice, t.Email, t.CreatedAt, t.UpdatedAt)
	}

	query := "SELECT id, product_id, qty, total_price, email, created_at, updated_at FROM transactions WHERE id = \\?"
	
	mock.ExpectPrepare(query).ExpectQuery().WithArgs(mockTransactions[1].ID).WillReturnRows(data)
	
	r := NewTransactionRepositoryMysql(db)
	err, res := r.GetById(context.TODO(), mockTransactions[1].ID)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestInsert(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("Error occured : %s", err.Error())
	}

	query := "INSERT transactions SET product_id=\\?, qty=\\?, total_price=\\?, email=\\?, created_at=\\?, updated_at=\\?"
	
	mock.ExpectPrepare(query).ExpectExec().WithArgs(mockTransaction.ProductId, mockTransaction.Qty, mockTransaction.TotalPrice, 
		mockTransaction.Email, mockTransaction.CreatedAt, mockTransaction.UpdatedAt).WillReturnResult(sqlmock.NewResult(int64(mockTransaction.ID), 1))
	
	r := NewTransactionRepositoryMysql(db)
	errInsert := r.Insert(context.TODO(), &mockTransaction)
	assert.NoError(t, errInsert)
	assert.Nil(t, errInsert)
}