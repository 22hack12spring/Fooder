package model

type Price struct {
	GenreCode string `db:"genre_code"`
	Name      string `db:"name"`
}

type PricesRepository interface {
}
