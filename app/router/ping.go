package router

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// /api/ping
func (h *Handlers) Ping(c echo.Context) error {
	log.Println("ping recieved")
	return c.String(http.StatusOK, "pong")
}
