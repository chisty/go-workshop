package data

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator"
)

//Validation struct
type Validation struct {
	validate *validator.Validate
}

//NewValidation to init
func NewValidation() *Validation {
	validate := validator.New()
	validate.RegisterValidation("customSku", validateSKU)
	return &Validation{validate}
}

//Validate function is validating method
func (v *Validation) Validate(i interface{}) ValidationErrors {
	err := v.validate.Struct(i)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		if len(errs) == 0 {
			return nil
		}

		var returnErrs []ValidationError
		for _, err := range errs {
			ve := ValidationError{err.(validator.FieldError)}
			returnErrs = append(returnErrs, ve)

		}
		return returnErrs
	}
	return nil
}

//ValidationError Method
type ValidationError struct {
	validator.FieldError
}

func (v ValidationError) Error() string {
	return fmt.Sprintf(
		"Key: '%s' Error: Field validation for '%s' failed on the '%s' tag",
		v.Namespace(),
		v.Field(),
		v.Tag(),
	)
}

//ValidationErrors struct
type ValidationErrors []ValidationError

//Errors method
func (v ValidationErrors) Errors() []string {
	errors := []string{}

	for _, err := range v {
		errors = append(errors, err.Error())
	}
	return errors
}

func validateSKU(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`[a-z]+-[a-z]+`)
	matches := re.FindAllString(fl.Field().String(), -1)

	if len(matches) == 1 {
		return true
	}
	return false
}
