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

func (uc *ProductDomain) GetAll(ctx context.Context) (error, []*product.Entity) {
	err, product := uc.productRepository.Fetch(ctx)

	if err != nil {
		return err, nil
	}

	return nil, product
}