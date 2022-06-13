package domain

import (
	"context"

	"github.com/FauzanAr/clean-and-go/modules/product"
)
 
type ProductDomain struct {
	productRepository	product.Repository
}

func NewProductDomain(productRepository product.Repository) product.Domain {
	return &ProductDomain{
		productRepository: productRepository,
	}
}

func (uc *ProductDomain) Fetch(ctx context.Context) error {

	return nil
}