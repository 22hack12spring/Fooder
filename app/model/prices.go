package model

import (
	"context"
	"sync"
)

type Price struct {
	PriceCode string `db:"price_code"`
	Name      string `db:"name"`
}

type PriceCache struct {
	Name map[string]string
	Mux  sync.RWMutex
}

var PriceCacheData PriceCache

type PricesRepository interface {
	GetPrices() ([]Price, error)
	PriceCodeToName(ctx context.Context, code string) (string, error)
}

func (repo *SqlxRepository) GetPrices() ([]Price, error) {
	sql := "Select * FROM prices"

	var prices []Price

	err := repo.db.Select(&prices, sql)

	if err != nil {
		return []Price{}, err
	}

	PriceCacheData.Mux.Lock()

	if PriceCacheData.Name == nil {
		PriceCacheData.Name = make(map[string]string)
	}

	for _, p := range(prices) {
		PriceCacheData.Name[p.PriceCode] = p.Name
	}
	PriceCacheData.Mux.Unlock()

	return prices, nil
}

func (repo *SqlxRepository) PriceCodeToName(ctx context.Context, code string) (string, error) {
	var name string
	if err := repo.db.Get(&name, "SELECT name FROM prices WHERE price_code=?", code); err != nil {
		return "", err
	}
	return name, nil
}
