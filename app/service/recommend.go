package service

import (
	"context"
	"log"
	"math/rand"

	"github.com/22hack12spring/backend/model"
)

type ShopDetail struct {
	Id        string  `json:"id"`
	Name      string  `json:"name"`
	LogoImage string  `json:"logoImage"`
	Address   string  `json:"address"`
	Lat       float64 `json:"lat"`
	Lng       float64 `json:"lng"`
	Genre     string  `json:"genre"`
	SubGenre  string  `json:"subgenre"`
	Budget    string  `json:"price"`
	Open      string  `json:"open"`
	Close     string  `json:"close"`
	Url       string  `json:"url"`
	Photo     string  `json:"photo"`
	Lunch     string  `json:"lunch"`
}

type Answer struct {
	Id     int    `json:"questionId" validate:"required,min=1,max=7"`
	Answer string `json:"answer" validate:"required,oneof=yes no"`
}

// GenerateRecommend　おすすめのお店を1件返す
func (s *Services) GenerateRecommend(ctx context.Context, uuid string, answers []Answer) (*ShopDetail, error) {
	request, err := s.Repo.GetSearch(ctx, uuid)
	if err != nil {
		return nil, err
	}
	args := model.SearchArgs{}
	if request.Station.Valid {
		args.Station = &request.Station.String
	}
	if request.Lat.Valid && request.Lng.Valid {
		args.Lat = &request.Lat.Float64
		args.Lng = &request.Lng.Float64
	}
	shops, err := s.GetGourmetsData(ctx, args)
	if err != nil {
		return nil, err
	}

	// 予測値を計算する
	prediction := [3]float64{}
	const noWeight = 0.7
	for _, ans := range answers {
		shop, err := s.Repo.GetShopByQuestionId(ctx, ans.Id, uuid)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		pr, err := s.ShopToSimilarityVec3(ctx, shop)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		if ans.Answer == "yes" {
			prediction[0] += pr[0]
			prediction[1] += pr[1]
			prediction[2] += pr[2]
		} else {
			prediction[0] -= noWeight * pr[0]
			prediction[1] -= noWeight * pr[1]
			prediction[2] -= noWeight * pr[2]
		}
	}

	// 類似度の高いものからランダムに返す
	vec3s, err := ShopsToShopParams(shops)
	if err != nil {
		return nil, err
	}

	num := 7
	if len(vec3s) < num {
		num = len(vec3s)
	}
	similarShops := FindSimilarVec3(vec3s, prediction, num)

	result := rand.Intn(len(similarShops))
	return similarShops[result].Shop, nil
}

// 質問結果の値を計算する
func (s *Services) AnswerShopVector(questions []model.Shops, answers []Answer) ([3]float64, error) {
	// ナイーブな実装な気がするのでいい感じに修正してください
	return [3]float64{0.7, 0.7, -0.5}, nil
}
