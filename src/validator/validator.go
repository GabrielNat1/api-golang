package validator

import (
	"regexp"
	"strconv"
	"strings"
)

type ValidationError struct {
	Field   string
	Message string
}

type Validator struct {
	Errors []ValidationError
}

func New() *Validator {
	return &Validator{Errors: []ValidationError{}}
}

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

func (v *Validator) AddError(field, message string) {
	v.Errors = append(v.Errors, ValidationError{
		Field:   field,
		Message: message,
	})
}

func (v *Validator) Check(ok bool, field, message string) {
	if !ok {
		v.AddError(field, message)
	}
}

func (v *Validator) Required(value string, field string) {
	if strings.TrimSpace(value) == "" {
		v.AddError(field, "This field is required")
	}
}
func (v *Validator) MaxLength(value string, field string, max int) {
	if len(value) > max {
		v.AddError(field, "This field cannot be more than "+strconv.Itoa(max)+" characters long")
	}
}

func (v *Validator) MinLength(value string, field string, min int) {
	if len(value) < min {
		v.AddError(field, "This field must be at least "+strconv.Itoa(min)+" characters long")
	}
}

func (v *Validator) Email(value string, field string) {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(pattern, value)
	if !match {
		v.AddError(field, "Invalid email address")
	}
}
func (v *Validator) Between(value int, field string, min, max int) {
	if value < min || value > max {
		v.AddError(field, "This field must be between "+strconv.Itoa(min)+" and "+strconv.Itoa(max))
	}
}

func (v *Validator) GetErrors() []ValidationError {
	return v.Errors
}
