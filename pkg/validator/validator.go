package validator

import (
	validators "github.com/go-playground/validator/v10"
)

type Validator interface {
	ValidateStruct(inf interface{}) error
}

type validator struct {
	validate *validators.Validate
}

func New() Validator {
	return &validator{
		validate: validators.New(),
	}
}

func (v *validator) ValidateStruct(inf interface{}) error {
	return v.validate.Struct(inf)
}
