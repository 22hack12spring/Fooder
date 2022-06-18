package service

import (
	"context"
	"errors"

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
func (s *Services) GenerateQuestions(ctx context.Context, arg model.SearchArgs, uuid string) ([]*ShopData, [7]string, error) {
	// 質問に使えるお店一覧
	questionShops, err := s.Repo.GetShops(ctx)
	if err != nil {
		return nil, [7]string{}, err
	}
	questionVectors, err := s.ShopsToQuestionSimilarity(ctx, questionShops)
	if err != nil {
		return nil, [7]string{}, err
	}
	if len(questionVectors) < 3 {
		return nil, [7]string{}, errors.New("question vectors is less than 3")
	}

	vecs := RandomRotate()
	res := []*ShopData{}
	// 1個目
	// TODO: 質問を3種類しか生成してないので、もったいないかも?V1アルゴリズムに期待
	// TODO: 質問で埋めるために無理やりループしてるのも実装がびみょい。
	ids := [7]string{}
	q1, err := FindSimilarQuestionVec3(vecs[0], questionVectors)
	if err != nil {
		return nil, [7]string{}, err
	}
	q1data, err := s.ShopToShopData(ctx, q1.Shops, 1)
	if err != nil {
		return nil, [7]string{}, err
	}
	res = append(res, q1data)
	ids[0] = q1.Shops.ShopId

	// 2個目
	q2, err := FindSimilarQuestionVec3(vecs[1], questionVectors)
	if err != nil {
		return nil, [7]string{}, err
	}
	for i := 0; i < 2; i++ {
		q2data, err := s.ShopToShopData(ctx, q2.Shops, i+1)
		if err != nil {
			return nil, [7]string{}, err
		}
		res = append(res, q2data)
		ids[i+1] = q2.Shops.ShopId
	}

	// 3個目
	q3, err := FindSimilarQuestionVec3(vecs[2], questionVectors)
	if err != nil {
		return nil, [7]string{}, err
	}
	for i := 0; i < 4; i++ {
		q3data, err := s.ShopToShopData(ctx, q3.Shops, i+3)
		if err != nil {
			return nil, [7]string{}, err
		}
		res = append(res, q3data)
		ids[i+3] = q3.Shops.ShopId
	}
	return res, ids, nil
}

// model.Shops to ShopData
func (s *Services) ShopToShopData(ctx context.Context, shop *model.Shops, id int) (*ShopData, error) {
	genre, err := s.Repo.GenreCodeToName(ctx, shop.GenreCode)
	if err != nil {
		return nil, err
	}
	var subGenre string
	if shop.SubgenreCode.Valid {
		subGenre, err = s.Repo.GenreCodeToName(ctx, shop.SubgenreCode.String)
		if err != nil {
			return nil, err
		}
	}
	price, err := s.Repo.PriceCodeToName(ctx, shop.PriceCode)
	if err != nil {
		return nil, err
	}

	return &ShopData{
		Id:       id,
		Image:    shop.Image,
		Genre:    genre,
		SubGenre: subGenre,
		Price:    price,
	}, nil
}
