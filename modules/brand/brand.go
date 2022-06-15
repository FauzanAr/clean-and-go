package brand

import (
	"context"
	"net/http"
)


type Repository interface {
	Fetch(ctx context.Context, query string, args ...interface{}) (error, []*Entity)
	GetAll(ctx context.Context) (error, []*Entity)
}

type Domain interface {
	GetAll(ctx context.Context) (error, []*Entity)
}

type Handler interface {
	Brand(response http.ResponseWriter, request *http.Request)
	GetAll(response http.ResponseWriter, request *http.Request)
}

type Entity struct {
	ID				int		`json:"id"`
	Name			string	`json:"name" validate:"required"`
	Description		string	`json:"description"`
	CreatedAt		int		`json:"createdAt"`
	UpdatedAt		int		`json:"updatedAt"`
}