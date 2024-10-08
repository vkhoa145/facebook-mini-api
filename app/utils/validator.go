package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ValidateParams(payload interface{}) map[string]string {
	validate := validator.New(validator.WithRequiredStructEnabled())
	errorParams := validate.Struct(payload)
	errors := make(map[string]string)

	if errorParams == nil {
		return nil
	}

	for _, err := range errorParams.(validator.ValidationErrors) {
		arg := generateArgumentForLocale(err.Param())
		message := Locale(fmt.Sprintf("en.go_validators.%s", err.Tag()), arg)
		errors[err.Field()] = message
	}

	return errors
}

func generateArgumentForLocale(err string) map[string]string {
	arg := map[string]string{
		"min":    err,
		"gt":     err,
		"lte":    err,
		"unique": err,
		"day":    err,
	}

	return arg
}
