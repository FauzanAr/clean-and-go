package validator

import (
	"fmt"
	"strings"

	"github.com/FauzanAr/clean-and-go/helpers/response"

	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	validate := cv.Validator.Struct(i)
	fmt.Println("masok" + validate.Error())
	if cv.Validator.Struct(i) != nil {
		var errorMsg string
		fmt.Println("Masok")
		errs := cv.Validator.Struct(i).(validator.ValidationErrors)

		// Register custom message for the tags
		switch errs[0].Tag() {
		case "required":
			errorMsg = errs[0].Field() + " is required"
			return response.BadRequest(errorMsg)
		case "email":
			errorMsg = errs[0].Field() + " is not valid email"
			return response.BadRequest(errorMsg)
		default: 
			errorMsg = fmt.Sprintf("\"%s\" is %s", strings.ToLower(errs[0].Field()), errs[0].Tag())
			return response.BadRequest(errorMsg)
		}
	}

	return nil
}

func New() *validator.Validate {
	return validator.New()
}
