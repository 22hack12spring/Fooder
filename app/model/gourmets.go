package model

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
)

type GourmetsRepository interface {
	InsertGourmetsCache(ctx context.Context, args SearchArgs, gourmetsRaw string) error
	GetGourmetsCache(ctx context.Context, args SearchArgs) (Gourmets, error)
	ParseGourmetsJSON(raw string) (GourmetsDetail, error)
	DeleteOldGourmets(ctx context.Context) error
}

type Gourmets struct {
	ID        int64           `db:"id"`
	Station   sql.NullString  `db:"station"`
	Lat       sql.NullFloat64 `db:"lat"`
	Lng       sql.NullFloat64 `db:"lng"`
	Shops     string          `db:"shops"`
	CreatedAt sql.NullTime    `db:"created_at"`
}

func (repo *SqlxRepository) InsertGourmetsCache(ctx context.Context, args SearchArgs, gourmetsRaw string) error {
	if args.Station != nil {
		// 駅名検索で保存
		_, err := repo.db.ExecContext(ctx, "INSERT INTO gourmets (station, shops) VALUES (?, ?)", args.Station, gourmetsRaw)
		if err != nil {
			return err
		}
	} else if args.Lat != nil && args.Lng != nil {
		// 緯度経度検索で保存
		_, err := repo.db.ExecContext(ctx, "INSERT INTO gourmets (lat, lng, shops) VALUES (?, ?, ?)", args.Lat, args.Lng, gourmetsRaw)
		if err != nil {
			return err
		}
	} else {
		// error
		return errors.New(fmt.Sprintf("Error: %s", "Invalid args"))
	}
	return nil
}

// キャッシュからデータを取り出す
func (repo *SqlxRepository) GetGourmetsCache(ctx context.Context, args SearchArgs) (Gourmets, error) {
	var shopsRaw Gourmets
	var err error

	query := "SELECT shops FROM gourmets WHERE"
	if args.Station != nil {
		// 駅名検索
		query += " station = ? LIMIT 1"
		err = repo.db.GetContext(ctx, &shopsRaw, query, args.Station)
	} else if args.Lat != nil && args.Lng != nil {
		// 位置情報検索
		query += " lat = ? AND"
		query += " lng = ? LIMIT 1"
		err = repo.db.GetContext(ctx, &shopsRaw, query, args.Lat, args.Lng)
	} else {
		// error
		return Gourmets{}, errors.New(fmt.Sprintf("Error: %s", "Invalid args"))
	}

	if err != nil {
		return Gourmets{}, err
	}

	return shopsRaw, nil
}

// see API doc -> https://webservice.recruit.co.jp/doc/hotpepper/reference.html
type GourmetsDetail struct {
	ApiVersion       string       `json:"api_version"`
	ResultsAvailable int          `json:"results_available"`
	ResultsReturned  string       `json:"results_returned"`
	ResultsStart     int          `json:"results_start"`
	Shop             []ShopDetail `json:"shop"`
}

type ShopDetail struct {
	ID           string     `json:"id"`
	Name         string     `json:"name"`
	LogoImage    string     `json:"logo_image"`
	NameKana     string     `json:"name_kana"`
	Address      string     `json:"address"`
	StationName  string     `json:"station_name"`
	Lat          float64    `json:"lat"`
	Lng          float64    `json:"lng"`
	Genre        ShopGenre  `json:"genre"`
	SubGenre     ShopGenre  `json:"sub_genre"`
	Budget       ShopBudget `json:"budget"`
	Access       string     `json:"access"`
	MobileAccess string     `json:"mobile_access"`
	Urls         struct {
		Pc string `json:"pc"`
	} `json:"urls"`
	Photo struct {
		Pc struct {
			L string `json:"l"`
			M string `json:"m"`
			S string `json:"s"`
		} `json:"pc"`
	} `json:"photo"`
	Open     string `json:"open"`
	Close    string `json:"close"`
	Lunch    string `json:"lunch"`
	Midnight string `json:"midnight"`
}

type ShopGenre struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type ShopBudget struct {
	Code    string `json:"code"`
	Name    string `json:"name"`
	Average string `json:"average"`
}

type GourmetsResult struct {
	Result GourmetsDetail `json:"results"`
}

// raw APIからデータを取り出す
func (repo *SqlxRepository) ParseGourmetsJSON(raw string) (GourmetsDetail, error) {
	var gourmetsDetail GourmetsResult
	err := json.Unmarshal([]byte(raw), &gourmetsDetail)
	if err != nil {
		return GourmetsDetail{}, err
	}
	return gourmetsDetail.Result, nil
}

func (repo *SqlxRepository) DeleteOldGourmets(ctx context.Context) error {
	_, err := repo.db.ExecContext(ctx, "DELETE FROM gourmets WHERE created_at < NOW() - INTERVAL 1 DAY")
	if err != nil {
		return err
	}
	return nil
}
