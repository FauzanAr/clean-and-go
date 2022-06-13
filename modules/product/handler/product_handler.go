package product_handler

import (
	"net/http"

	"github.com/FauzanAr/clean-and-go/modules/product"
)

type ProductHandler struct {
	productDomain product.Domain
}

func NewProdutHandler(domain product.Domain) product.Handler {
	return &ProductHandler{
		productDomain: domain,
	}
}

func (handler *ProductHandler) ShowAll(w http.ResponseWriter, r *http.Request) {

}

func (handler *ProductHandler) Product(w http.ResponseWriter, r *http.Request) {

}