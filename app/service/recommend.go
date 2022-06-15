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
	result := rand.Intn(len(shops))
	return &shops[result], nil
}
