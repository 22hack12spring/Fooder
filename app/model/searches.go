package model

import "database/sql"

type SearchesRepository interface {
	CreateSearch(arg SearchArgs) (error, Searches)
}

type SearchArgs struct {
	Station *string
	Lat     *float64
	Lng     *float64
}

type Searches struct {
	ID        string          `db:"id"`
	Station   sql.NullString  `db:"station"`
	Lat       sql.NullFloat64 `db:"lat"`
	Lng       sql.NullFloat64 `db:"lng"`
	CreatedAt string          `db:"created_at"`
}

// CreateSearch　uuidを発行して、Searchesテーブルにデータを追加する
func (repo *SqlxRepository) CreateSearch(arg SearchArgs) (error, Searches) {
	// implement here
	return nil, Searches{}
}
