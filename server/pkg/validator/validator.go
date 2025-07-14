package validator

import (
	"net/url"
	"regexp"
	"strings"

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
	urlStr := fl.Field().String()
	if urlStr == "" {
		return false
	}

	u, err := url.Parse(urlStr)
	if err != nil {
		return false
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	}

	parts := strings.Split(u.Hostname(), ".")
	if len(parts) < 2 {
		return false
	}
	for _, part := range parts {
		if len(part) == 0 {
			return false
		}
	}

	return true
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
