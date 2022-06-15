package brand_handler

import (
	"encoding/json"
	"net/http"

	"github.com/FauzanAr/clean-and-go/helpers/logger"
	"github.com/FauzanAr/clean-and-go/helpers/response"
	"github.com/FauzanAr/clean-and-go/helpers/validator"
	"github.com/FauzanAr/clean-and-go/modules/brand"
)

type BrandHandler struct {
	brandDomain brand.Domain
}

func NewBrandHandler(d brand.Domain) brand.Handler {
	return &BrandHandler{
		brandDomain: d,
	}
}

func(h *BrandHandler) Brand(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.Post(w, r)
		return
	case "GET":
		h.GetAll(w, r)
	default: 
		response.ResponseErr(w, response.MethodNotAllowed(nil))
		return
	}
}

func(h *BrandHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	err, res := h.brandDomain.GetAll(r.Context())

	if err != nil {
		response.ResponseErr(w, err)
		return
	}

	logger.InfoLogger.Println("Request successfully handled")
	response.Response(w, res, "Success", 200, 200)
	return
}

func(h *BrandHandler) Post(w http.ResponseWriter, r *http.Request) {
	var body brand.Entity
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		response.ResponseErr(w, response.BadRequest("Error in decode the request: " + err.Error()))
		return
	}

	validate := validator.New()
	errValidate := validate.Struct(body)

	if errValidate != nil {
		response.ResponseErr(w, response.BadRequest(errValidate.Error()))
		return
	}

	errCreate := h.brandDomain.Create(r.Context(), &body)

	if errCreate != nil {
		response.ResponseErr(w, errCreate)
		return
	}

	response.Response(w, nil, "Success created", 201, 201)
	return
}