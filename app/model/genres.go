package model

import (
	"context"
)

type Genre struct {
	GenreCode string `db:"genre_code"`
	Name      string `db:"name"`
}

type GenresRepository interface {
	GetGenres() ([]*Genre, error)
	GenreCodeToName(ctx context.Context, code string) (string, error)
}

func (repo *SqlxRepository) GetGenres() ([]*Genre, error) {
	return nil, nil
}

func (repo *SqlxRepository) GenreCodeToName(ctx context.Context, code string) (string, error) {
	var name string
	if err := repo.db.Get(&name, "SELECT name FROM genres WHERE genre_code=?", code); err != nil {
		return "", err
	}
	return name, nil
}

/*

if err := db.Get(&city, "SELECT * FROM city WHERE Name='Tokyo'"); errors.Is(err, sql.ErrNoRows) {
        log.Printf("no such city Name = %s", "Tokyo")
    } else if err != nil {
        log.Fatalf("DB Error: %s", err)
    }

package model

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

type Genre struct {
	GenreCode string `db:"genre_code"`
	Name      string `db:"name"`
}

func main() {

	e := echo.New()

	e.GET("/luftalian/genres", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World.\n")
	})

	e.Logger.Fatal(e.Start(":8080"))

	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("genres_code"), os.Getenv("name")))
	if err != nil {
		log.Fatalf("Cannot Connect to Database: %s", err)
	}

	fmt.Println("Connected!")

	var name Name
	if err := db.Get(&name, "SELECT genre_code, name;"); errors.Is(err, sql.ErrNoRows) {
		log.Printf("error")
	} else if err != nil {
		log.Fatalf("DB Error: %s", err)
	}

	fmt.Printf("%s\n", name)
}
*/
