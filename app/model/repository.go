package model

import "github.com/jmoiron/sqlx"

type Repository interface {
	GenresRepository
	SearchsRepository
}

type SqlxRepository struct {
	db *sqlx.DB
}

func NewSqlxRepository(db *sqlx.DB) Repository {
	repo := &SqlxRepository{
		db: db,
	}
	return repo
}
