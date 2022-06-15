package service

import (
	"context"
	"math/rand"

	"github.com/22hack12spring/backend/model"
)

// TODO: たぶん、jsonのやつも書いたほうが良い↓
type ShopDetail struct {
	Id        string `json:"id"`
	Name      string
	LogoImage string
	Address   string
	Lat       float64
	Lng       float64
	Genre     []string
	Budget    string
	Open      string
	Close     string
	Url       string
	Photo     string
	Lunch     string
}

type Answer struct {
	Id     int    `json:"questionId"`
	Answer string `json:"answer"`
}

// GenerateRecommend　おすすめのお店を1件返す
func (s *Services) GenerateRecommend(ctx context.Context, uuid string, answers []Answer) (*ShopDetail, error) {
	// mock とりあえず大岡山の店を返す
	station := "大岡山"
	args := model.SearchArgs{Station: &station}
	shops, err := s.GetGourmetsData(ctx, args)
	if err != nil {
		return nil, err
	}
	// 類似度の高いものからランダムに返す
	vec3s, err := ShopsToShopParams(shops)
	if err != nil {
		return nil, err
	}
	// 中華が食べたい、お金のない人
	query := [3]float64{0.7, 0.7, -0.5}
	num := 7
	if len(vec3s) < num {
		num = len(vec3s)
	}
	similarShops := FindSimilarVec3(vec3s, query, num)

	result := rand.Intn(len(similarShops))
	return similarShops[result].Shop, nil
}
