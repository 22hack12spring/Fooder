package model

import (
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
	GetGenres() ([]Genre, error)
	GenreCodeToName(code string) (string, error)
}

func (repo *SqlxRepository) GetGenres() ([]Genre, error) {
	sql := "SELECT * FROM genres"

	var genres []Genre

	err := repo.db.Select(&genres, sql)

	if err != nil {
		return []Genre{}, err
	}

	GenreCacheData.Mux.Lock()

	if GenreCacheData.Name == nil {
		GenreCacheData.Name = make(map[string]string)
	}

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
