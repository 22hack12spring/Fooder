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
	return []Shops{}, nil
}

// GetShopsOfQuestions  Questions の id から使われた Shops の id を返却
func (repo *SqlxRepository) GetShopsOfQuestion(id string) ([7]Shops, error) {
	return [7]Shops{}, nil
}