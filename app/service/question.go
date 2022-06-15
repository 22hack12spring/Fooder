package service

import "context"

// 質問数
const QUESTION_NUM int = 3

type ShopData struct {
	Id       int
	Image    string
	Genre    string
	SubGenre string
	Price    string
}

// 質問の配列を生成する
func (s *Services) GenerateQuestions(ctx context.Context, station string, lat float64, lng float64) ([]*ShopData, error) {

	return nil, nil
}
