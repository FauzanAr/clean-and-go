package transaction_handler

import (
	"encoding/json"
	"net/http"

	"github.com/FauzanAr/clean-and-go/helpers/response"
	"github.com/FauzanAr/clean-and-go/helpers/validator"
	"github.com/FauzanAr/clean-and-go/modules/transaction"
)

type TransactionHandler struct {
	td transaction.Domain
}

func NewTransactionHandler(td transaction.Domain) transaction.Handler {
	return &TransactionHandler {
		td: td,
	}
}

func (h *TransactionHandler) Transaction(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.Post(w, r)
		return
	default:
		response.ResponseErr(w, response.MethodNotAllowed(nil))
		return
	}
}

func (h *TransactionHandler) Get(w http.ResponseWriter, r *http.Request) {
	return
}

func (h *TransactionHandler) GetDetail(w http.ResponseWriter, r *http.Request) {
	return
}

func (h *TransactionHandler) Post(w http.ResponseWriter, r *http.Request) {
	var body transaction.Entity
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

	errCreate := h.td.Create(r.Context(), &body)

	if errCreate != nil {
		response.ResponseErr(w, errCreate)
		return
	}

	response.Response(w, nil, "Success created", 201, 201)
	return
}