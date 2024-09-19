package utils

import "github.com/go-playground/validator/v10"

func ValidateParams(payload interface{}) map[string][]string {
	validate := validator.New(validator.WithRequiredStructEnabled())
	errorParams := validate.Struct(payload)

	errors := make(map[string][]string)

	for _, err := range errorParams.(validator.ValidationErrors) {
		errors[err.Field()] = append(errors[err.Field()], err.Tag())
	}

	return errors
}
