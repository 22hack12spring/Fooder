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
func (s *Services) GenerateQuestions(ctx context.Context, arg model.SearchArgs) ([]*ShopData, error) {
	// 質問に使えるお店一覧
	questionShops, err := s.Repo.GetShops(ctx)
	if err != nil {
		return nil, err
	}
	questionVectors, err := s.ShopsToQuestionSimilarity(ctx, questionShops)
	if err != nil {
		return nil, err
	}
	if len(questionVectors) < 3 {
		return nil, errors.New("question vectors is less than 3")
	}

	vecs := RandomRotate()
	res := []*ShopData{}
	// 1個目
	// TODO: sqlNullStringでValidを確認せずに""として使ってるのがキモい気がするので修正したい
	// TODO: 質問を3種類しか生成してないので、もったいないかも?V1アルゴリズムに期待
	// TODO: 質問で埋めるために無理やりループしてるのも実装がびみょい。
	q1, err := FindSimilarQuestionVec3(vecs[0], questionVectors)
	if err != nil {
		return nil, err
	}
	q1data := ShopData{
		Id:       1,
		Image:    q1.Shops.Image,
		Genre:    q1.Shops.GenreCode,
		SubGenre: q1.Shops.SubgenreCode.String,
		Price:    q1.Shops.PriceCode,
	}
	res = append(res, &q1data)
	// 2個目
	q2, err := FindSimilarQuestionVec3(vecs[1], questionVectors)
	if err != nil {
		return nil, err
	}
	for i := 0; i < 2; i++ {
		q2data := ShopData{
			Id:       i + 2,
			Image:    q2.Shops.Image,
			Genre:    q2.Shops.GenreCode,
			SubGenre: q2.Shops.SubgenreCode.String,
			Price:    q2.Shops.PriceCode,
		}
		res = append(res, &q2data)
	}
	// 3個目
	q3, err := FindSimilarQuestionVec3(vecs[2], questionVectors)
	if err != nil {
		return nil, err
	}
	for i := 0; i < 4; i++ {
		q3data := ShopData{
			Id:       i + 4,
			Image:    q3.Shops.Image,
			Genre:    q3.Shops.GenreCode,
			SubGenre: q3.Shops.SubgenreCode.String,
			Price:    q3.Shops.PriceCode,
		}
		res = append(res, &q3data)
	}
	return res, nil
	// mock
	// for i := range ShopsMock {
	// 	res = append(res, &ShopsMock[i])
	// }
	// return res, nil
}
