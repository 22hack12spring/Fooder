package model

import (
	"context"
	"database/sql"
	"log"
)

type ShopsRepository interface {
	GetShops(ctx context.Context) ([]Shops, error)
}

type Shops struct {
	ShopId       string         `db:"shop_id"`
	Name         string         `db:"name"`
	Image        string         `db:"image"`
	GenreCode    string         `db:"genre_code"`
	SubgenreCode sql.NullString `db:"subgenre_code"`
	PriceCode    string         `db:"price_code"`
	CreatedAt    string         `db:"created_at"`
}

func (repo *SqlxRepository) GetShops(ctx context.Context) ([]Shops, error) {
	var shops []Shops
	err := repo.db.SelectContext(ctx, &shops, "SELECT * FROM shops")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return shops, nil
}
