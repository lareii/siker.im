package utils

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func ValidateStruct(s any) error {
	return validate.Struct(s)
}

func urlValidator(fl validator.FieldLevel) bool {
	regex := regexp.MustCompile(`^https?://[^\s/$.?#].[^\s]*$`)
	return regex.MatchString(fl.Field().String())
}

func slugValidator(fl validator.FieldLevel) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)
	return regex.MatchString(fl.Field().String())
}

func init() {
	validate = validator.New()
	validate.RegisterValidation("slug_valid", slugValidator)
	validate.RegisterValidation("url_valid", urlValidator)
}
