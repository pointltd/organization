package http

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Validator struct {
	Validator *validator.Validate
}

func (cv *Validator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		errors := make(map[string]string)

		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, fieldError := range validationErrors {
				errors[fieldError.Field()] = fmt.Sprintf("Validation failed on '%s' rule", fieldError.Tag())
			}
		}

		return echo.NewHTTPError(http.StatusBadRequest, errors)
	}
	return nil
}
