package response

import (
	"encoding/json"
	"net/http"
)

type Model struct {
	Success			bool		`json:"success"`
	Message			string		`json:"message"`
	Data			interface{}	`json:"data"`
	Code			int			`json:"code"`
}

func Response(w http.ResponseWriter, data interface{}, message string, code int, responseCode int) {
	success := false

	if code < http.StatusBadRequest {
		success = true
	}

	result := Model {
		Success: success,
		Message: message,
		Data: data,
		Code: code,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseCode)
	json.NewEncoder(w).Encode(result)

	return 
}

func ResponseErr(w http.ResponseWriter, err interface{}) {
	obj := GetErrorType(err)
	result := Model {
		Success: false,
		Message: obj.Message,
		Data: obj.Data,
		Code: obj.Code,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(obj.ResponseCode)
	json.NewEncoder(w).Encode(result)

	return
}

func GetErrorType(err interface{}) CommonError {
	errData := CommonError{}

	switch obj := err.(type) {
	case BadRequestError:
		errData.Code = obj.Code
		errData.Data = obj.Data
		errData.Message = obj.Message
		errData.ResponseCode = http.StatusBadRequest
		return errData
	default:
		errData.ResponseCode = http.StatusInternalServerError
		return errData
	}
}
