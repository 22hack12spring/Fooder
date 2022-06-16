package model

import "context"

type Price struct {
	PriceCode string `db:"price_code"`
	Name      string `db:"name"`
}

type PricesRepository interface {
	GetPrices() ([]*Price, error)
	PriceCodeToName(ctx context.Context, code string) (string, error)
}

func (repo *SqlxRepository) GetPrices() ([]*Price, error) {
	return nil, nil
}

func (repo *SqlxRepository) PriceCodeToName(ctx context.Context, code string) (string, error) {
	var name string
	if err := repo.db.Get(&name, "SELECT name FROM prices WHERE price_code=?", code); err != nil {
		return "", err
	}
	return name, nil
}
