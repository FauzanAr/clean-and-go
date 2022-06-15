package transaction_domain

import (
	"context"

	"github.com/FauzanAr/clean-and-go/helpers/response"
	"github.com/FauzanAr/clean-and-go/modules/product"
	"github.com/FauzanAr/clean-and-go/modules/transaction"
)

type TransactionDomain struct {
	tr	transaction.Repository
	pr	product.Repository
}

func NewTransactionDomain(tr transaction.Repository, pr product.Repository) transaction.Domain {
	return &TransactionDomain{
		tr: tr,
		pr: pr,
	}
}

//TODO Impl db transaction and concurency
func (d *TransactionDomain) Create(ctx context.Context, data *transaction.Entity) error {
	err, product := d.pr.GetOne(ctx, data.ProductId)
	if err != nil {
		return err
	}

	if product == nil {
		return response.NotFound("Product not found!")
	}

	if data.Qty > product.Stock {
		return response.BadRequest("Not enough stock!")
	}

	data.TotalPrice = product.Price * data.Qty
	errInsert := d.tr.Insert(ctx, data)
	if errInsert != nil {
		return errInsert
	}

	product.Stock = product.Stock - data.Qty
	errUpdate := d.pr.Update(ctx, product)
	if errUpdate != nil {
		return errUpdate
	}

	return nil
}

func (d *TransactionDomain) GetAllByEmail(ctx context.Context, email string) (error, []*transaction.Entity) {
	return nil, nil
}