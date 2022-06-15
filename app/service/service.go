package service

import (
	"context"

	"github.com/22hack12spring/backend/model"
)

type Service interface {
	GenerateRecommend(ctx context.Context, uuid string, answers []Answer) (*ShopDetail, error)
	GenerateQuestions() ([]*ShopData, error)
	GetGourmetsData(ctx context.Context, args model.SearchArgs) ([]ShopDetail, error)
}

type Services struct {
	Repo model.Repository
}

func NewServices(repo model.Repository) (*Services, error) {
	return &Services{
		Repo: repo,
	}, nil
}
