package service

type ShopData struct {
	Id       int
	Image    string
	Genre    string
	SubGenre string
	Price    string
}

// 質問の配列を生成する
func (s *Services) GenerateQuestions() ([]*ShopData, error) {
	return nil, nil
}
