package model

type Genre struct {
	GenreCode string `db:"genre_code"`
	Name      string `db:"name"`
}

type GenresRepository interface {
	GetGenres() ([]*Genre, error)
}

func (repo *SqlxRepository) GetGenres() ([]*Genre, error) {
	return nil, nil
}
