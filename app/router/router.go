package router

import (
	"net/http"

	"github.com/22hack12spring/backend/model"
	"github.com/22hack12spring/backend/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Handlers struct {
	Repo    model.Repository
	Service service.Service
}

func (h *Handlers) SetRouting(e *echo.Echo) error {
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodOptions},
	}))
	api := e.Group("/api/v1")
	gourmet := api.Group("/gourmet")
	{
		gourmet.POST("/start", h.PostGourmetStart)
		gourmet.POST("/answer", h.PostGourmetAnswer)
	}
	api.GET("/ping", h.Ping)
	return nil
}
