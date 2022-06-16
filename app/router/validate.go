package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type GourmetValidator struct {
	validator *validator.Validate
}

func (sv *GourmetValidator) Validate(i interface{}) error {
	return sv.validator.Struct(i)
}

func GetValidator() *GourmetValidator {
	v := validator.New()

	return &GourmetValidator{v}
}

// validatedBind return HTTP error
func validatedBind(c echo.Context, i interface{}) error {
	if err := c.Bind(i); err != nil {
		c.Logger().Error(err)
		return errorBind
	}
	if err := c.Validate(i); err != nil {
		c.Logger().Error(err)
		return errorValidate
	}
	return nil
}
