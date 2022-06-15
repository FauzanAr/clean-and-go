package brand_handler

import (
	"net/http"

	"github.com/FauzanAr/clean-and-go/helpers/logger"
	"github.com/FauzanAr/clean-and-go/helpers/response"
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
		h.Insert(w, r)
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

func(h *BrandHandler) Insert(w http.ResponseWriter, r *http.Request) {

}