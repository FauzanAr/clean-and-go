package product

import (
	"context"
	"net/http"
)

type Domain interface {
	GetAll(ctx context.Context) (error, []*Entity)
	GetById(ctx context.Context, id int) (error, *Entity)
	GetByBrand(ctx context.Context, id int) (error, []*Entity)
	Create(ctx context.Context, data *Entity) error
}

type Repository interface {
	GetAll(ctx context.Context) (error, []*Entity)
	GetOne(ctx context.Context, id int) (error, *Entity)
	GetByBrand(ctx context.Context, id int) (error, []*Entity)
	Fetch(ctx context.Context, query string, args ...interface{}) (error, []*Entity)
	Insert(ctx context.Context, data *Entity) error
	Update(ctx context.Context, data *Entity) error
}

type Handler interface {
	Product(response http.ResponseWriter, request *http.Request)
	GetAll(response http.ResponseWriter, request *http.Request)
	GetOne(response http.ResponseWriter, request *http.Request)
	Post(response http.ResponseWriter, request *http.Request)
	GetByBrand(response http.ResponseWriter, request *http.Request)
}

type Entity struct {
	ID			int		`json:"id"`
	BrandId		int		`json:"brandId" validate:"required"`
	Name		string	`json:"name" validate:"required"`
	Description	string	`json:"description" validate:"required"`
	Price		int		`json:"price" validate:"required"`
	Stock		int		`json:"stock" validate:"required"`
	CreatedAt	int		`json:"createdAt"`
	UpdatedAt	int		`json:"updatedAt"`
}