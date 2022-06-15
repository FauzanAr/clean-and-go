package transaction

import (
	"context"
	"net/http"
)

type Repository interface {
	Fetch(ctx context.Context, query string, args ...interface{}) (error, []*Entity)
	GetByEmail(ctx context.Context, email string) (error, []*Entity)
	Insert(ctx context.Context, data *Entity) error
}

type Domain interface {
	GetAllByEmail(ctx context.Context, email string) (error, []*Entity)
	Create(ctx context.Context, body *Entity) error
}

type Handler interface {
	Transaction(response http.ResponseWriter, request *http.Request)
	Get(response http.ResponseWriter, request *http.Request)
	GetDetail(response http.ResponseWriter, request *http.Request)
	Post(response http.ResponseWriter, request *http.Request)
}

type Entity struct {
	ID				int		`json:"id"`
	ProductId		int		`json:"productId" validate:"required"`
	Qty				int		`json:"qty" validate:"required"`
	TotalPrice		int		`json:"totalPrice"`
	Email			string	`json:"email" validate:"required,email"`
	CreatedAt		int		`json:"createdAt"`
	UpdatedAt		int		`json:"updateAt"`
}