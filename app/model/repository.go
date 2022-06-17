package model

import "github.com/jmoiron/sqlx"

type Repository interface {
	GenresRepository
	PricesRepository
	SearchesRepository
	QuestionsRepository
	GourmetsRepository
	GourmetsRequest
	ShopsRepository

	Initialize() error
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

func (repo *SqlxRepository) Initialize() error {
	_, err := repo.GetGenres()
	
	if err != nil {
		return err
	}

	return nil
}