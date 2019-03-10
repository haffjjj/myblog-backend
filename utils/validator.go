package utils

import "github.com/go-playground/validator"

//Validator ...
type Validator struct {
	Validator *validator.Validate
}

//Validate ...
func (v *Validator) Validate(i interface{}) error {
	return v.Validator.Struct(i)
}
