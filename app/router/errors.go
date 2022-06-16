package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	errorValidate = echo.NewHTTPError(http.StatusBadRequest, "Validation error")
	errorBind     = echo.NewHTTPError(http.StatusBadRequest, "Bind error")
)
