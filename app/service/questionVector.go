package service

import (
	"context"

	"github.com/22hack12spring/backend/model"
)

// TODO: リファクタしたい
// shopVectorとquestionVectorでコードが同じなのに共通化できてないのが最悪
// 取り回すデータ型も同じやつが分裂してるので、命名ともたせる情報を変更する

type QuestionSimilarity struct {
	Shops      *model.Shops
	Vec3       [3]float64
	Similarity float64
}

// ShopsToQuestionSimilarity vec3変換が成功した要素すべて返す
func (s *Services) ShopsToQuestionSimilarity(ctx context.Context, shops []model.Shops) ([]QuestionSimilarity, error) {
	var result []QuestionSimilarity
	for i := range shops {
		vec3, err := s.ShopToSimilarityVec3(ctx, shops[i])
		if err != nil {
			continue
		}
		result = append(result, QuestionSimilarity{
			Shops:      &shops[i],
			Vec3:       vec3,
			Similarity: 0,
		})
	}
	return result, nil
}

func (s *Services) ShopToSimilarityVec3(ctx context.Context, shop model.Shops) ([3]float64, error) {
	genre, err := s.Repo.GenreCodeToName(shop.GenreCode)
	if err != nil {
		return [3]float64{}, err
	}
	subGenre, err := s.Repo.GenreCodeToName(shop.SubgenreCode.String)
	if err != nil {
		// skip
		subGenre = ""
	}
	budget, err := s.Repo.PriceCodeToName(ctx, shop.PriceCode)
	if err != nil {
		return [3]float64{}, err
	}
	param, err := GenreBudgetToVec3(genre, subGenre, budget)
	if err != nil {
		return [3]float64{}, err
	}
	return param, nil
}

func FindSimilarQuestionVec3(vec3 [3]float64, questionVec3s []QuestionSimilarity) (QuestionSimilarity, error) {
	var result QuestionSimilarity
	result.Similarity = 1000.0
	for i := range questionVec3s {
		questionVec3s[i].Similarity = SimilarityVec3(vec3, questionVec3s[i].Vec3)
		if questionVec3s[i].Similarity < result.Similarity {
			result = questionVec3s[i]
		}
	}
	return result, nil
}
