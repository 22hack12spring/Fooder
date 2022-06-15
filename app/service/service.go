package service

import "github.com/22hack12spring/backend/model"

type Service interface {
	GenerateRecommend(uuid string, answers []Answer) (*ShopDetail, error)
	GenerateQuestions() ([]*ShopData, error)
}

type Services struct {
	Repo model.Repository
}

func NewServices(repo model.Repository) (*Services, error) {
	return &Services{
		Repo: repo,
	}, nil
}
