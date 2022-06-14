package product_handler

import (
	"net/http"
	"strconv"

	"github.com/FauzanAr/clean-and-go/helpers/logger"
	"github.com/FauzanAr/clean-and-go/helpers/response"
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

func (handler *ProductHandler) Product(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST" :
		handler.Post(w, r)
		return
	case "GET" :
		query, ok := r.URL.Query()["id"]

		if ok && len(query[0]) > 0 {
			handler.GetOne(w, r)
		} else {
			handler.GetAll(w, r)
		}

		return
	default:
		response.ResponseErr(w, response.MethodNotAllowed(nil))
	}
}

func(handler *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	err, res := handler.productDomain.GetAll(r.Context());

	if err != nil {
		response.ResponseErr(w, err)
		return
	}

	logger.InfoLogger.Println("Request successfully handled")
	response.Response(w, res, "Success", 200, 200)
	return
}

func(handler *ProductHandler) GetOne(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query()["id"]
	id, err := strconv.Atoi(param[0])

	if err != nil {
		response.ResponseErr(w, response.BadRequest("id must be integer"))
		return
	}

	err, res := handler.productDomain.GetById(r.Context(), id)

	if err != nil {
		response.ResponseErr(w, err)
		return
	}

	response.Response(w, res, "Success", 200, 200)
	return
}

func(handler *ProductHandler) Post(w http.ResponseWriter, r *http.Request) {

}