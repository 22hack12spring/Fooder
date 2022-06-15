package router

import (
	"github.com/22hack12spring/backend/model"
	"github.com/22hack12spring/backend/service"
	"github.com/labstack/echo/v4"
)

type Handlers struct {
	Repo    model.Repository
	Service service.Service
}

func (h *Handlers) SetRouting(e *echo.Echo) error {
	api := e.Group("/api/v1")
	gourmet := api.Group("/gourmet")
	{
		gourmet.POST("/start", h.PostGourmetStart)
		gourmet.POST("/answer", h.PostGourmetAnswer)
	}
	api.GET("/ping", h.Ping)
	return nil
}
