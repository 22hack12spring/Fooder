package router

import (
	"github.com/22hack12spring/backend/model"
	service "github.com/22hack12spring/backend/services"
	"github.com/labstack/echo/v4"
)

type Handlers struct {
	Repo    model.Repository
	Service service.Service
}

func (h *Handlers) SetRouting(e *echo.Echo) error {
	api := e.Group("/api")
	{
		api.GET("/ping", h.Ping)
	}

	return nil
}
