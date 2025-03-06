package utils

import validator2 "github.com/go-playground/validator/v10"

func Validate(data interface{}) error {
	validator := validator2.New()
	err := validator.Struct(data)

	return err
}
