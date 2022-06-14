package product

import (
	"context"
	"net/http"
)

type Domain interface {
	GetAll(ctx context.Context) (error, []*Entity)
	GetById(ctx context.Context, id int) (error, *Entity)
}

type Repository interface {
	GetAll(ctx context.Context) (error, []*Entity)
	GetOne(ctx context.Context, id int) (error, *Entity)
	Fetch(ctx context.Context, args ...interface{}) (error, []*Entity)
	Insert(ctx context.Context) error
}

type Handler interface {
	Product(response http.ResponseWriter, request *http.Request)
	GetAll(response http.ResponseWriter, request *http.Request)
	GetOne(response http.ResponseWriter, request *http.Request)
	Post(response http.ResponseWriter, request *http.Request)
}

type Entity struct {
	ID			int		`json:"id"`
	BrandId		int		`json:"brandId"`
	Name		string	`json:"name"`
	Description	string	`json:"description"`
	Price		int		`json:"price"`
	Stock		int		`json:"stock"`
	CreatedAt	int		`json:"createdAt"`
	UpdatedAt	int		`json:"updatedAt"`
}