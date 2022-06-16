package service

import "context"

// 質問数
const QUESTION_NUM int = 3

type ShopData struct {
	Id       int    `json:"id"`
	Image    string `json:"image"`
	Genre    string `json:"genre"`
	SubGenre string `json:"sub_genre"`
	Price    string `json:"price"`
}

// 質問の配列を生成する
func (s *Services) GenerateQuestions(ctx context.Context, station string, lat float64, lng float64) ([]*ShopData, error) {

	return nil, nil
}
