package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handlers) PostGourmetStart(c echo.Context) error {
	// TODO: implement
	return c.String(http.StatusOK, "start")
}

func (h *Handlers) PostGourmetAnswer(c echo.Context) error {
	// TODO: implement
	return c.String(http.StatusOK, "answer")
}
