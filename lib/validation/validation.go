package validation

import (
	"fmt"
	"regexp"
)

type Validator interface {
	Validate(interface{}) error
}

type StringValidator struct {
	maxLength int
	minLength int
	regexp    *regexp.Regexp
	required  bool
}

func String() StringValidator {
	return StringValidator{}
}

func (v StringValidator) MaxLength(maxLength int) StringValidator {
	v.maxLength = maxLength
	return v
}

func (v StringValidator) MinLength(minLength int) StringValidator {
	v.minLength = minLength
	return v
}

func (v StringValidator) Regexp(regexp *regexp.Regexp) StringValidator {
	v.regexp = regexp
	return v
}

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

type MapValidator struct {
	validators map[string]Validator
}

func Map() MapValidator {
	return MapValidator{
		validators: map[string]Validator{},
	}
}

func (v MapValidator) Add(key string, validator Validator) MapValidator {
	v.validators[key] = validator
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
