package router

import (
	"github.com/labstack/echo/v4"
)

func SetRouting(e *echo.Echo) error {
	api := e.Group("/api")
	{
		api.GET("/ping", Ping)
	}

	return nil;
}