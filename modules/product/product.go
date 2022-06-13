package product

import (
	"context"
	"net/http"
)

type Domain interface {
	Fetch(ctx context.Context) error
}

type Repository interface {
	Fetch(ctx context.Context) error
}

type Handler interface {
	Product(response http.ResponseWriter, request *http.Request)
}