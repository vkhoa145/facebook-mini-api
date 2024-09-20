package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ValidateParams(payload interface{}) map[string][]string {
	validate := validator.New(validator.WithRequiredStructEnabled())
	errorParams := validate.Struct(payload)
	errors := make(map[string][]string)

	for _, err := range errorParams.(validator.ValidationErrors) {
		message := GetMessageByKey(fmt.Sprintf("en.error_validations.%s", err.Tag()))
		errors[err.Field()] = append(errors[err.Field()], message)
	}

	return errors
}
