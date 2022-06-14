package main

import (
	"log"
	"os"

	"github.com/22hack12spring/backend/model"
	"github.com/22hack12spring/backend/router"
	service "github.com/22hack12spring/backend/services"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

var (
	db *sqlx.DB
)

func main() {
	db, err := model.EstablishConnection()
	if err != nil {
		panic(err)
	}

	err = db.DB.Ping()
	if err != nil {
		panic(err)
	}

	e := echo.New()

	repo := model.NewSqlxRepository(db)
	services, err := service.NewServices(repo)
	if err != nil {
		panic(err)
	}
	handlers := router.Handlers{
		Repo:    repo,
		Service: services,
	}

	err = handlers.SetRouting(e)

	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
