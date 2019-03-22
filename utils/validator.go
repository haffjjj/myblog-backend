package utils

import "github.com/go-playground/validator"

//Validator represent validator struct
type Validator struct {
	Validator *validator.Validate
}

//Validate is method from validator
func (v *Validator) Validate(i interface{}) error {
	return v.Validator.Struct(i)
}
