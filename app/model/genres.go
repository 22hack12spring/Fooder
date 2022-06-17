package model

import (
	"context"
	"fmt"
	"sync"
)

type Genre struct {
	GenreCode string `db:"genre_code"`
	Name      string `db:"name"`
}

type GenreCache struct {
	Name map[string]string
	Mux  sync.RWMutex
}

var GenreCacheData GenreCache

type GenresRepository interface {
	GetGenres(ctx context.Context) ([]Genre, error)
	GenreCodeToName(code string) (string, error)
}

func (repo *SqlxRepository) GetGenres(ctx context.Context) ([]Genre, error) {
	sql := "SELECT * FROM genres"

	var genres []Genre

	err := repo.db.SelectContext(ctx, &genres, sql)

	if err != nil {
		return []Genre{}, err
	}

	GenreCacheData.Mux.Lock()
	for _, g := range genres {
		GenreCacheData.Name[g.GenreCode] = g.Name
	}
	GenreCacheData.Mux.Unlock()

	return genres, nil
}

func (repo *SqlxRepository) GenreCodeToName(code string) (string, error) {
	GenreCacheData.Mux.RLock()
	name, ok := GenreCacheData.Name[code]
	GenreCacheData.Mux.RUnlock()

	if !ok {
		return "", fmt.Errorf("backend: No such genre code exists")
	}

	return name, nil
}
