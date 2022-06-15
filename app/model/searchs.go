package model

import "database/sql"

type SearchsRepository interface {
	CreateSearch(arg CreateSearchArgs) (error, Searchs)
}

type CreateSearchArgs struct {
	station *string
	lat     *float64
	lng     *float64
}

type Searchs struct {
	ID        string          `db:"id"`
	Station   sql.NullString  `db:"station"`
	Lat       sql.NullFloat64 `db:"lat"`
	Lng       sql.NullFloat64 `db:"lng"`
	CreatedAt string          `db:"created_at"`
}

// CreateSearch　uuidを発行して、Searchsテーブルにデータを追加する
func (repo *SqlxRepository) CreateSearch(arg CreateSearchArgs) (error, Searchs) {
	// implement here
	return nil, Searchs{}
}
