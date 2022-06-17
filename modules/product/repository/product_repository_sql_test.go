package product_repository

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/FauzanAr/clean-and-go/modules/product"

	"github.com/stretchr/testify/assert"
)

var mockProducts = []product.Entity {
	{ID: 1, BrandId: 1, Name: "Product 1", Description: "Description", Price: 10000, Stock: 10, CreatedAt: int(time.Now().Unix()), UpdatedAt: int(time.Now().Unix())},
	{ID: 2, BrandId: 2, Name: "Product 2", Description: "Description", Price: 20000, Stock: 20, CreatedAt: int(time.Now().Unix()), UpdatedAt: int(time.Now().Unix())},
}

var mockProduct = product.Entity{
	ID: 3,
	BrandId: 2,
	Name: "Product 2",
	Description: "Description",
	Price: 10000,
	Stock: 2,
	CreatedAt: int(time.Now().Unix()),
	UpdatedAt: int(time.Now().Unix()),
}

func TestFetch(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("Error occured : %s", err.Error())
	}

	data := sqlmock.NewRows([]string{"id", "brand_id", "name", "description", "price", "stock", "created_at", "updated_at"})
	
	for _, p := range mockProducts {
		data.AddRow(p.ID, p.BrandId, p.Name, p.Description, p.Price, p.Stock, p.CreatedAt, p.UpdatedAt)
	}

	query := `SELECT id, brand_id, name, description, price, stock, created_at, updated_at FROM products`
	
	mock.ExpectPrepare(query).ExpectQuery().WillReturnRows(data)
	
	r := NewProductRepositoryMysql(db)
	err, res := r.Fetch(context.TODO(), query)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestGetAll(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("Error occured : %s", err.Error())
	}

	data := sqlmock.NewRows([]string{"id", "brand_id", "name", "description", "price", "stock", "created_at", "updated_at"})
	
	for _, p := range mockProducts {
		data.AddRow(p.ID, p.BrandId, p.Name, p.Description, p.Price, p.Stock, p.CreatedAt, p.UpdatedAt)
	}

	query := `SELECT id, brand_id, name, description, price, stock, created_at, updated_at FROM products`
	
	mock.ExpectPrepare(query).ExpectQuery().WillReturnRows(data)
	
	r := NewProductRepositoryMysql(db)
	err, res := r.GetAll(context.TODO())
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

//TODO make sure WithArgs() return true data
func TestGetOne(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("Error occured : %s", err.Error())
	}

	data := sqlmock.NewRows([]string{"id", "brand_id", "name", "description", "price", "stock", "created_at", "updated_at"})
	
	for _, p := range mockProducts {
		data.AddRow(p.ID, p.BrandId, p.Name, p.Description, p.Price, p.Stock, p.CreatedAt, p.UpdatedAt)
	}

	query := `SELECT id, brand_id, name, description, price, stock, created_at, updated_at FROM products WHERE id = \\?`

	mock.ExpectPrepare(query).ExpectQuery().WithArgs(mockProducts[1].ID).WillReturnRows(data)
	
	r := NewProductRepositoryMysql(db)
	err, res := r.GetOne(context.TODO(), mockProducts[1].ID)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestGetOneNil(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("Error occured : %s", err.Error())
	}

	data := sqlmock.NewRows([]string{"id", "brand_id", "name", "description", "price", "stock", "created_at", "updated_at"})

	query := `SELECT id, brand_id, name, description, price, stock, created_at, updated_at FROM products WHERE id = \\?`
	
	mock.ExpectPrepare(query).ExpectQuery().WithArgs(mockProducts[0].ID).WillReturnRows(data)
	
	r := NewProductRepositoryMysql(db)
	err, res := r.GetOne(context.TODO(), mockProducts[0].ID)
	assert.NoError(t, err)
	assert.Nil(t, res)
}
func TestGetByBrand(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("Error occured : %s", err.Error())
	}

	data := sqlmock.NewRows([]string{"id", "brand_id", "name", "description", "price", "stock", "created_at", "updated_at"})

	for _, p := range mockProducts {
		data.AddRow(p.ID, p.BrandId, p.Name, p.Description, p.Price, p.Stock, p.CreatedAt, p.UpdatedAt)
	}

	query := `SELECT id, brand_id, name, description, price, stock, created_at, updated_at FROM products WHERE brand_id = \\?`
	
	mock.ExpectPrepare(query).ExpectQuery().WithArgs(mockProducts[0].BrandId).WillReturnRows(data)
	
	r := NewProductRepositoryMysql(db)
	err, res := r.GetByBrand(context.TODO(), mockProducts[0].BrandId)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("Error occured : %s", err.Error())
	}

	query := "UPDATE products SET brand_id=\\?, name=\\?, description=\\?, price=\\?, stock=\\?, created_at=\\?, updated_at=\\? WHERE id=\\?"
	
	mockProducts[1].BrandId = 5
	mock.ExpectPrepare(query).ExpectExec().WithArgs(mockProducts[1].BrandId, mockProducts[1].Name, mockProducts[1].Description, mockProducts[1].Price, 
		mockProducts[1].Stock, int(mockProducts[1].CreatedAt), int(mockProducts[1].UpdatedAt), mockProducts[1].ID).WillReturnResult(sqlmock.NewResult(int64(mockProducts[1].ID), 1))
	
	r := NewProductRepositoryMysql(db)
	errUpdate := r.Update(context.TODO(), &mockProducts[1])
	assert.NoError(t, err)
	assert.Nil(t, errUpdate)
}

func TestInsert(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("Error occured : %s", err.Error())
	}

	query := "INSERT products SET brand_id=\\?, name=\\?, description=\\?, price=\\?, stock=\\?, created_at=\\?, updated_at=\\?"
	
	mock.ExpectPrepare(query).ExpectExec().WithArgs(mockProduct.BrandId, mockProduct.Name, mockProduct.Description, mockProduct.Price, 
		mockProduct.Stock, mockProduct.CreatedAt, mockProduct.UpdatedAt).WillReturnResult(sqlmock.NewResult(int64(mockProduct.ID), 1))
	
	r := NewProductRepositoryMysql(db)
	errInsert := r.Insert(context.TODO(), &mockProduct)
	assert.NoError(t, err)
	assert.Nil(t, errInsert)
}