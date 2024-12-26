package validation

import (
	"github.com/go-playground/validator/v10"
)

// CustomValidator struct
type BodyValidator struct {
	validator *validator.Validate
}

// Validate method
func (cv *BodyValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
