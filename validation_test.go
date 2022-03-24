package requestvalidation

import (
	"testing"
)

func TestRequestValidator_ValidateNoError(t *testing.T) {
	type Request struct {
		Name string `validate:"required"`
	}

	requestValidator := NewRequestValidator()

	validation, err := requestValidator.Validate(Request{
		Name: "daniel",
	})

	if err != nil {
		t.Error("Expected non error")
	}

	if validation.Status != "success" {
		t.Error("Expected validation to be valid")
	}

	_, ok := validation.Errors["Name"]
	if ok {
		t.Error("Expected non error")
	}
}

func TestRequestValidator_ValidateRequired(t *testing.T) {
	type Request struct {
		Name string `validate:"required" message:"Name is required"`
	}

	requestValidator := NewRequestValidator()

	validation, err := requestValidator.Validate(Request{
		Name: "",
	})

	if err == nil {
		t.Error("Expected error")
	}

	if validation.Status != "error" {
		t.Error("Expected validation to be invalid")
	}

	_, ok := validation.Errors["Name"]
	if !ok {
		t.Error("Expected error")
	}

	if len(validation.Errors["Name"]) != 1 {
		t.Error("Expected one error only")
	}

	if validation.Errors["Name"][0] != "Name is required" {
		t.Error("Expected error")
	}
}

func TestRequestValidator_ValidateMultiple(t *testing.T) {
	type Request struct {
		Email string `validate:"required,email" message:"Email is required and must be a valid email"`
	}

	requestValidator := NewRequestValidator()

	validation, err := requestValidator.Validate(Request{
		Email: "daniel",
	})

	if err == nil {
		t.Error("Expected error")
	}

	if validation.Status != "error" {
		t.Error("Expected validation to be invalid")
	}

	_, ok := validation.Errors["Email"]
	if !ok {
		t.Error("Expected error")
	}

	if len(validation.Errors["Email"]) != 1 {
		t.Error("Expected one error only")
	}

	if validation.Errors["Email"][0] != "Email is required and must be a valid email" {
		t.Error("Expected error")
	}
}
