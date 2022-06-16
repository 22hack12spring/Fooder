package service

import (
	"context"

	"github.com/22hack12spring/backend/model"
)

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
func (s *Services) GenerateQuestions(ctx context.Context, arg model.SearchArgs) ([]*ShopData, error) {
	// mock
	res := []*ShopData{}
	for i := range ShopsMock {
		res = append(res, &ShopsMock[i])
	}
	return res, nil
}
