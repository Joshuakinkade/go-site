package validation

import (
	"regexp"
)

type Example struct {
	FirstName string `validate:"required"`
}

type ValidationError struct {
	Field   string
	Message string
}

type StringValidator struct {
	regex     *regexp.Regexp
	maxLength int
	minLength int
}

func NewStringValidator(regex string, minLength, maxLength int) StringValidator {
	return StringValidator{
		regex:     regexp.MustCompile(regex),
		maxLength: maxLength,
		minLength: minLength,
	}
}

func ValidateStruct(val interface{}) []ValidationError {
	return nil
}
