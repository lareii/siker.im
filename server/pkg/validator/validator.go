package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func validateSlug(fl validator.FieldLevel) bool {
	slug := fl.Field().String()
	if slug == "" {
		return false
	}

	if len(slug) > 50 {
		return false
	}

	slugRegex := regexp.MustCompile(`^[a-zA-Z0-9-_]+$`)
	return slugRegex.MatchString(slug)
}

func validateURL(fl validator.FieldLevel) bool {
	url := fl.Field().String()
	if url == "" {
		return false
	}

	urlRegex := regexp.MustCompile(`^(https?:\/\/)?((localhost)|(([\w-]+\.)+[\w-]{2,})|(\d{1,3}(\.\d{1,3}){3}))(:\d+)?\/?$`)
	return urlRegex.MatchString(url)
}

func Validate(s any) error {
	if validate == nil {
		return nil
	}

	return validate.Struct(s)
}

func init() {
	validate = validator.New()

	validate.RegisterValidation("url_valid", validateURL)
	validate.RegisterValidation("slug_valid", validateSlug)
}
