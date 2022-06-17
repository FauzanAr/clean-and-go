package brand_repository

import (
	"context"
	"testing"
	"time"

	"github.com/FauzanAr/clean-and-go/modules/brand"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var mockBrands = []brand.Entity {
	{ID: 1, Name: "Brand 1", Description: "Description 1", CreatedAt: int(time.Now().Unix()), UpdatedAt: int(time.Now().Unix())},
	{ID: 2, Name: "Brand 2", Description: "Description 2", CreatedAt: int(time.Now().Unix()), UpdatedAt: int(time.Now().Unix())},
}

var mockBrand = brand.Entity {
	ID: 3,
	Name: "Brand 3",
	Description: "Description 3",
	CreatedAt: int(time.Now().Unix()),
	UpdatedAt: int(time.Now().Unix()),
}

func TestFetch(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("Error occured : %s", err.Error())
	}

	data := sqlmock.NewRows([]string{"id", "name", "description", "created_at", "updated_at"})
	
	for _, b := range mockBrands {
		data.AddRow(b.ID, b.Name, b.Description, b.CreatedAt, b.UpdatedAt)
	}

	query := "SELECT id, name, description, created_at, updated_at FROM brands"
	
	mock.ExpectPrepare(query).ExpectQuery().WillReturnRows(data)
	
	r := NewBrandRepositoryMysql(db)
	err, res := r.Fetch(context.TODO(), query)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestGetAll(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("Error occured : %s", err.Error())
	}

	data := sqlmock.NewRows([]string{"id", "name", "description", "created_at", "updated_at"})
	
	for _, b := range mockBrands {
		data.AddRow(b.ID, b.Name, b.Description, b.CreatedAt, b.UpdatedAt)
	}

	query := "SELECT id, name, description, created_at, updated_at FROM brands"
	
	mock.ExpectPrepare(query).ExpectQuery().WillReturnRows(data)
	
	r := NewBrandRepositoryMysql(db)
	err, res := r.GetAll(context.TODO())
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestGetByName(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("Error occured : %s", err.Error())
	}

	data := sqlmock.NewRows([]string{"id", "name", "description", "created_at", "updated_at"})
	
	for _, b := range mockBrands {
		data.AddRow(b.ID, b.Name, b.Description, b.CreatedAt, b.UpdatedAt)
	}

	query := "SELECT id, name, description, created_at, updated_at FROM brands WHERE name = \\?"
	
	mock.ExpectPrepare(query).ExpectQuery().WithArgs(mockBrands[1].Name).WillReturnRows(data)
	
	r := NewBrandRepositoryMysql(db)
	err, res := r.GetByName(context.TODO(), mockBrands[1].Name)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestGetById(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("Error occured : %s", err.Error())
	}

	data := sqlmock.NewRows([]string{"id", "name", "description", "created_at", "updated_at"})
	
	for _, b := range mockBrands {
		data.AddRow(b.ID, b.Name, b.Description, b.CreatedAt, b.UpdatedAt)
	}

	query := "SELECT id, name, description, created_at, updated_at FROM brands WHERE id = \\?"
	
	mock.ExpectPrepare(query).ExpectQuery().WithArgs(mockBrands[1].ID).WillReturnRows(data)
	
	r := NewBrandRepositoryMysql(db)
	err, res := r.GetById(context.TODO(), mockBrands[1].ID)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestInsert(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("Error occured : %s", err.Error())
	}

	query := "INSERT brands SET name=\\?, description=\\?, created_at=\\?, updated_at=\\?"
	
	mock.ExpectPrepare(query).ExpectExec().WithArgs(mockBrand.Name, mockBrand.Description, mockBrand.CreatedAt,
		mockBrand.UpdatedAt).WillReturnResult(sqlmock.NewResult(int64(mockBrand.ID), 1))
	
	r := NewBrandRepositoryMysql(db)
	errInsert := r.Insert(context.TODO(), &mockBrand)
	assert.NoError(t, err)
	assert.Nil(t, errInsert)
}