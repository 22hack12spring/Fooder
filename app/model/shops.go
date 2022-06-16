package model

type ShopsRepository interface {
	GetShops() ([]Shops, error)
	GetShopsBySearchId(id string) ([7]Shops, error)
}

type Shops struct {
	Shop_id			string 	`db:"shop_id"`
	Name			string	`db:"name"`
	Image			string	`db:"image"`
	Genre_code		string 	`db:"genre_code"`
	Subgenre_code	string	`db:"subgenre_code"`
	Price_code 		string	`db:"price_code"`
	CreatedAt		string	`db:"created_at"`
}

// GetShops  Shops の一覧を取得(Limit: 100)
func (repo *SqlxRepository) GetShops() ([]Shops, error) {
	sql := "SELECT * FROM shops ORDER BY created_at LIMIT 100";

	rows, err := repo.db.DB.Query(sql)

	if err != nil {
		panic(err)
	}

	var shops []Shops
	var s Shops

	for rows.Next() {
		err = rows.Scan(&s.Shop_id, &s.Name, &s.Image, &s.Genre_code, &s.Subgenre_code, &s.Price_code, &s.CreatedAt)

		if err != nil {
			return []Shops{}, nil
		}

		shops = append(shops, s)
	}

	return shops, nil
}

// GetShopsBySearchId  Searches の id から使われた Shops の配列を取得する
func (repo *SqlxRepository) GetShopsBySearchId(search_id string) ([7]Shops, error) {
	questions, err := repo.GetQuestionsBySearchId(search_id)

	if err != nil {
		return [7]Shops{}, err
	}

	var shops [7]Shops
	var shop Shops
	sql := "SELECT * FROM shops WHERE shop_id = ?"

	for i, q := range(questions) {
		row := repo.db.DB.QueryRow(sql, q.Shop_id)
		row.Scan(&shop.Shop_id, &shop.Name, &shop.Image, &shop.Genre_code, &shop.Subgenre_code, &shop.Price_code, &shop.CreatedAt)
		shops[i] = shop
	}

	return shops, nil
}