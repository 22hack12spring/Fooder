package model

type ShopsRepository interface {
	GetShops() ([]Shops, error)
	GetShopsOfQuestion(id string) ([7]Shops, error)
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

// GetShopsOfQuestions  Questions の id から使われた Shops の id を返却
func (repo *SqlxRepository) GetShopsOfQuestion(id string) ([7]Shops, error) {
	return [7]Shops{}, nil
}