package domain

import (
	"context"
	"time"

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
	err, product := uc.productRepository.GetAll(ctx)

	if err != nil {
		return err, nil
	}

	return nil, product
}

func (uc *ProductDomain) GetById(ctx context.Context, id int) (error, *product.Entity) {
	err, product := uc.productRepository.GetOne(ctx, id)

	if err != nil {
		return err, nil
	}

	return nil, product
}

func (uc *ProductDomain) Create(ctx context.Context, data *product.Entity) error {
	data.CreatedAt = int(time.Now().Unix())
	data.UpdatedAt = int(time.Now().Unix())

	err := uc.productRepository.Insert(ctx, data)

	if err != nil {
		return err
	}
	
	return nil
}

func (uc *ProductDomain) GetByBrand(ctx context.Context, id int) (error, []*product.Entity) {
	err, products := uc.productRepository.GetByBrand(ctx, id)

	if err != nil {
		return err, nil
	}

	return nil, products
}