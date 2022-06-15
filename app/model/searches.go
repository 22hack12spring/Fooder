package model

import (
	"database/sql"

	"github.com/google/uuid"
)

type SearchesRepository interface {
	CreateSearch(arg SearchArgs) (error, Searches)
	GetSearch(id string) (error, Searches)
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
	u, err := uuid.NewRandom()

	if err != nil {
		return err, Searches{}
	}

	var station sql.NullString
	var lat sql.NullFloat64
	var lng sql.NullFloat64

	if arg.Station == nil {
		station.String, station.Valid = "", false
	} else {
		station.String, station.Valid = *arg.Station, true
	}
	if arg.Lat == nil {
		lat.Float64, lat.Valid = 0, false
	} else {
		lat.Float64, lat.Valid = *arg.Lat, true
	}
	if arg.Lng == nil {
		lng.Float64, lng.Valid = 0, false
	} else {
		lng.Float64, lng.Valid = *arg.Lat, true
	}

	search := &Searches{
		ID: u.String(),
		Station: station,
		Lat: lat,
		Lng: lng,
	}

	sql := "INSERT INTO searches (id, station, lat, lng) VALUE (?, ?, ?, ?)"
	_, err = repo.db.Exec(sql, search.ID, search.Station, search.Lat, search.Lng)

	if err != nil {
		return err, Searches{}
	}

	return nil, *search
}

// GetSearch  該当する id の Searches を検索
func (repo *SqlxRepository) GetSearch(id string) (error, Searches) {
	sql := "SELECT * FROM searches WHERE id = ?"
	row := repo.db.QueryRow(sql, id)  
	
	var s Searches

	err := row.Scan(&s.ID, &s.Station, &s.Lat, &s.Lng, &s.CreatedAt)

	if err != nil {
		return err, Searches{}
	}

	return nil, s
}