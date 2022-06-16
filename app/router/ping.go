package router

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// /api/ping
func (h *Handlers) Ping(c echo.Context) error {
	log.Println("ping recieved")
	name, err := h.Repo.GenreCodeToName(c.Request().Context(), "G001")
	if err != nil {
		return err
	}
	log.Println(name)
	return c.String(http.StatusOK, "あいいう")

}
