package response

import "net/http"

// Response code for http response code
// Code for custom code value and custom error
type CommonError struct {
	Data			interface{} `json:"data"`
	ResponseCode	int			`json:"responseCode"`
	Code			int			`json:"code"`
	Message			string		`json:"message"`
}

type BadRequestError struct {
	Data			interface{} `json:"data"`
	ResponseCode	int			`json:"responseCode"`
	Code			int			`json:"code"`
	Message			string		`json:"message"`
}

func (e BadRequestError) Error() string {
	return e.Message
}

func BadRequest(data interface{}) error {
	return BadRequestError{
		Data: data,
		ResponseCode: http.StatusBadRequest,
		Code: http.StatusBadRequest,
		Message: "Bad Request",
	}
}
