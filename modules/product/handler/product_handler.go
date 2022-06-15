package product_handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/FauzanAr/clean-and-go/helpers/logger"
	"github.com/FauzanAr/clean-and-go/helpers/response"
	"github.com/FauzanAr/clean-and-go/helpers/validator"
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
	var body product.Entity
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		response.ResponseErr(w, response.BadRequest("Error in decode the request: " + err.Error()))
		return
	}

	validate := validator.New()
	errValidate := validate.Struct(body)

	if errValidate != nil {
		err := response.BadRequest(errValidate.Error())
		response.ResponseErr(w, err)
		return
	}

	errCreate := handler.productDomain.Create(r.Context(), &body)

	if errCreate != nil {
		response.ResponseErr(w, errCreate)
		return
	}

	response.Response(w, nil, "Success created", 201, 201)
	return
}

func (handler *ProductHandler) GetByBrand(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		response.ResponseErr(w, response.MethodNotAllowed(nil))
		return
	}

	query, ok := r.URL.Query()["id"]
	if !ok || len(query)== 0 {
		response.ResponseErr(w, response.BadRequest("id cannot be null"))
		return
	}

	id, err := strconv.Atoi(query[0])
	if err != nil {
		response.ResponseErr(w, response.BadRequest("id must be integer"))
		return
	}

	err, res := handler.productDomain.GetByBrand(r.Context(), id)

	if err != nil {
		response.ResponseErr(w, err)
		return
	}
	response.Response(w, res, "Success", 200, 200)
	return
}