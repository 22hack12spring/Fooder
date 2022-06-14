package router

import (
	"github.com/22hack12spring/backend/model"
	"github.com/labstack/echo/v4"
)

type Handlers struct {
	Repo model.Repository
}

func (h *Handlers) SetRouting(e *echo.Echo) error {
	api := e.Group("/api")
	{
		api.GET("/ping", h.Ping)
	}

	return nil
}
