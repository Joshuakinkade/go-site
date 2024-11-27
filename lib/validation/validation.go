// Package validation provides types for validating data, most likey user input.

package validation

import (
	"fmt"
	"regexp"
)

// Validator defines an interface
type Validator interface {
	Validate(interface{}) error
}

// StringValidator validates that the strings match a regex pattern and are
// within a minimum and maximum length. Any empty requirements will be ignored.
type StringValidator struct {
	maxLength int
	minLength int
	regexp    *regexp.Regexp
	required  bool
}

// String returns an empty string validator. This will ensure that the value is
// a string, but nothing more.
func String() StringValidator {
	return StringValidator{}
}

// MaxLength sets the maximum string length for the value.
func (v StringValidator) MaxLength(maxLength int) StringValidator {
	v.maxLength = maxLength
	return v
}

// MinLength sets the minumum string length for the value.
func (v StringValidator) MinLength(minLength int) StringValidator {
	v.minLength = minLength
	return v
}

// Regexp sets a regular expression pattern that the string must match
func (v StringValidator) Regexp(regexp *regexp.Regexp) StringValidator {
	v.regexp = regexp
	return v
}

// Validate checks that the input string fits the criteria
func (v StringValidator) Validate(input interface{}) error {
	switch input := input.(type) {
	case string:
		if len(input) < v.minLength {
			return fmt.Errorf("value is too short")
		} else if v.maxLength > 0 && len(input) > v.maxLength {
			return fmt.Errorf("value is too long")
		} else if v.regexp != nil && !v.regexp.Match([]byte(input)) {
			return fmt.Errorf("value does not match pattern")
		}
	default:
		return fmt.Errorf("value is not a string")
	}

	return nil
}

// IntegerValidator checks that an integer field is in bounds
type IntegerValidator struct {
	min *int
	max *int
}

// Integer returns and empty IntegerValidator
func Integer() IntegerValidator {
	return IntegerValidator{}
}

// Min sets the minumum value of an integer field
func (v IntegerValidator) Min(min int) IntegerValidator {
	v.min = &min
	return v
}

// Max sets the maximum value of an integer field
func (v IntegerValidator) Max(max int) IntegerValidator {
	v.max = &max
	return v
}

// Max checks if the integer value is in bounds
func (v IntegerValidator) Validate(input interface{}) error {
	switch input := input.(type) {
	case int:
		if (v.min != nil && input < *v.min) || (v.max != nil && input > *v.max) {
			return fmt.Errorf("value must be between %v and %v", v.min, v.max)
		}
	default:
		return fmt.Errorf("value is not an integer")
	}
	return nil
}

type MapValidator struct {
	validators map[string]Validator
	required   []string
}

func Map() MapValidator {
	return MapValidator{
		validators: map[string]Validator{},
	}
}

func (v MapValidator) Add(key string, validator Validator, required bool) MapValidator {
	v.validators[key] = validator
	if required {
		v.required = append(v.required, key)
	}
	return v
}

func (v MapValidator) Validate(input interface{}) error {
	switch input := input.(type) {
	case map[string]interface{}:
		for field, value := range input {
			validator, ok := v.validators[field]
			if !ok {
				return fmt.Errorf("field %s is not known", field)
			}
			err := validator.Validate(value)
			if err != nil {
				return fmt.Errorf("field %s: %s", field, err.Error())
			}
		}
	default:
		return fmt.Errorf("value is not a map")
	}
	return nil
}
