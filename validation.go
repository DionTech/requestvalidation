package requestvalidation

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type RequestValidatorInterface interface {
	Validate(request interface{}) (ValidationResponse, error)
}

type ValidationResponse struct {
	Status string              `json:"status"`
	Errors map[string][]string `json:"errors"`
}

type RequestValidator struct {
	RequestValidatorInterface
	validate *validator.Validate
}

func NewRequestValidator() *RequestValidator {
	return &RequestValidator{
		validate: validator.New(),
	}
}

func (r *RequestValidator) Validate(request interface{}) (ValidationResponse, error) {
	err := r.validate.Struct(request)
	validationResponse := ValidationResponse{
		Status: "success",
		Errors: make(map[string][]string),
	}

	if err != nil {
		validationResponse.Status = "error"
		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return validationResponse, err
		}

		for _, err := range err.(validator.ValidationErrors) {

			_, ok := validationResponse.Errors[err.Field()]
			if !ok {
				validationResponse.Errors[err.Field()] = []string{}
			}

			validationResponse.Errors[err.Field()] = append(validationResponse.Errors[err.Field()], err.Tag())
			/*
				fmt.Println(err.Namespace())
				fmt.Println(err.Field())
				fmt.Println(err.StructNamespace())
				fmt.Println(err.StructField())
				fmt.Println(err.Tag())
				fmt.Println(err.ActualTag())
				fmt.Println(err.Kind())
				fmt.Println(err.Type())
				fmt.Println(err.Value())
				fmt.Println(err.Param())
				fmt.Println()
			*/
		}

		return validationResponse, err
	}
	return validationResponse, nil
}
