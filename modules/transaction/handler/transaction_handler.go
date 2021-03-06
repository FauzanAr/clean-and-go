package transaction_handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

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
	case "GET" :
		query := r.URL.Query()
		if query["id"] != nil && query["email"] != nil {
			h.GetDetail(w, r)
		} else if query["email"] != nil {
			h.Get(w, r)
		} else {
			response.ResponseErr(w, response.Unauthorized("Required params 'emai' for authentication"))
			break
		}

		return
	default:
		response.ResponseErr(w, response.MethodNotAllowed(nil))
		return
	}
}

func (h *TransactionHandler) Get(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query()["email"]
	email := param[0]

	if strings.TrimSpace(email) == "" {
		response.ResponseErr(w, response.BadRequest("email cannot be empty"))
		return
	}

	err, res := h.td.GetAllByEmail(r.Context(), email)
	if err != nil {
		response.ResponseErr(w, err)
		return
	}

	response.Response(w, res, "Success", 200, 200)
	return
}

func (h *TransactionHandler) GetDetail(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query()
	email := strings.TrimSpace(param["email"][0])
	id, err := strconv.Atoi(param["id"][0])

	if email == "" {
		response.ResponseErr(w, response.BadRequest("email cannot be empty"))
		return
	}

	if err != nil {
		response.ResponseErr(w, response.BadRequest("id cannot be empty and must be integer"))
		return
	}

	err, res := h.td.GetById(r.Context(), id, email)
	if err != nil {
		response.ResponseErr(w, err)
		return
	}

	response.Response(w, res, "Success", 200, 200)
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