package main

import (
	"log"

	"github.com/22hack12spring/backend/router"

	"github.com/labstack/echo/v4"
)

func main () {
	e := echo.New();

	err := router.SetRouting(e);

	if err != nil {
		log.Fatal(err)
	}

	e.Logger.Fatal(e.Start(":8080"))
}