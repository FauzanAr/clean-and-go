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

type NotFoundError struct {
	Data			interface{} `json:"data"`
	ResponseCode	int			`json:"responseCode"`
	Code			int			`json:"code"`
	Message			string		`json:"message"`
}

type InternalServerError struct {
	Data			interface{} `json:"data"`
	ResponseCode	int			`json:"responseCode"`
	Code			int			`json:"code"`
	Message			string		`json:"message"`
}

type MethodNotAllowedError struct {
	Data			interface{} `json:"data"`
	ResponseCode	int			`json:"responseCode"`
	Code			int			`json:"code"`
	Message			string		`json:"message"`
}

type UnauthorizedError struct {
	Data			interface{} `json:"data"`
	ResponseCode	int			`json:"responseCode"`
	Code			int			`json:"code"`
	Message			string		`json:"message"`
}

func (e BadRequestError) Error() string {
	return e.Message
}

func (e NotFoundError) Error() string {
	return e.Message
}

func (e InternalServerError) Error() string {
	return e.Message
}

func (e MethodNotAllowedError) Error() string {
	return e.Message
}

func (e UnauthorizedError) Error() string {
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

func NotFound(data interface{}) error {
	return NotFoundError{
		Data: data,
		ResponseCode: http.StatusNotFound,
		Code: http.StatusNotFound,
		Message: "Not Found",
	}
}

func InternalServerErr(data interface{}) error {
	return InternalServerError{
		Data: data,
		ResponseCode: http.StatusInternalServerError,
		Code: http.StatusInternalServerError,
		Message: "Internal Server Error",
	}
}

func MethodNotAllowed(data interface{}) error {
	return MethodNotAllowedError{
		Data: data,
		ResponseCode: http.StatusMethodNotAllowed,
		Code: http.StatusMethodNotAllowed,
		Message: "Method not allowed",
	}
}

func Unauthorized(data interface{}) error {
	return UnauthorizedError{
		Data: data,
		ResponseCode: http.StatusUnauthorized,
		Code: http.StatusUnauthorized,
		Message: "Unauthorized",
	}
}
