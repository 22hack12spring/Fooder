package main

import (
	"log"

	"github.com/22hack12spring/backend/router"
	"github.com/22hack12spring/backend/model"

	"github.com/labstack/echo/v4"
	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
)

func main () {

	db, err := model.EstablishConnection();
	if  err != nil {
		panic(err)
	}

	err = db.DB.Ping()
	if err != nil {
		panic(err)
	}

	e := echo.New();

	err = router.SetRouting(e);

	if err != nil {
		log.Fatal(err)
	}

	e.Logger.Fatal(e.Start(":8080"))
}