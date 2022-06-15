package model

import "database/sql"

type SearchsRepository interface {
	CreateSearch(arg SearchArgs) (error, Searchs)
}

type SearchArgs struct {
	Station *string
	Lat     *float64
	Lng     *float64
}

type Searchs struct {
	ID        string          `db:"id"`
	Station   sql.NullString  `db:"station"`
	Lat       sql.NullFloat64 `db:"lat"`
	Lng       sql.NullFloat64 `db:"lng"`
	CreatedAt string          `db:"created_at"`
}

// CreateSearch　uuidを発行して、Searchsテーブルにデータを追加する
func (repo *SqlxRepository) CreateSearch(arg SearchArgs) (error, Searchs) {
	// implement here
	return nil, Searchs{}
}
