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
