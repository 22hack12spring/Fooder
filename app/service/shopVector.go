package service

import (
	"errors"
	"fmt"
	"sort"
)

// ******* 推薦アルゴリズムV0 ********** //
var genreVec2 = map[string][2]float64{
	"居酒屋":         {0.5, 0.5},
	"ダイニングバー・バル":  {-0.5, -0.5},
	"創作料理":        {0.5, -1},
	"和食":          {1, -1},
	"洋食":          {-1, -0.5},
	"イタリアン・フレンチ":  {-1, -1},
	"中華":          {1, 1},
	"焼肉・ホルモン":     {-1, 1},
	"韓国料理":        {-0.5, 1},
	"アジア・エスニック料理": {-0.5, 0.5},
	"各国料理":        {-1, 0.5},
	"カラオケ・パーティ":   {0.5, 1},
	"バー・カクテル":     {1, -0.5},
	"ラーメン":        {1, 0.5},
	"お好み焼き・もんじゃ":  {0.5, -0.5},
	"カフェ・スイーツ":    {-0.5, -1},
	"その他グルメ":      {-0.5, -1},
}

var budgetVec = map[string]float64{
	"～500円":        -0.8,
	"501～1000円":    -0.7,
	"1001～1500円":   -0.5,
	"1501～2000円":   -0.3,
	"2001～3000円":   -0.1,
	"3001～4000円":   0.1,
	"4001～5000円":   0.5,
	"5001～7000円":   0.8,
	"7001～10000円":  1,
	"10001～15000円": 1.5,
	"15001～20000円": 2,
	"20001～30000円": 2.5,
	"30001円～":      3,
}

const weight float64 = 0.4

type ShopParams struct {
	Shop *ShopDetail
	Vec3 [3]float64
}

// ShopsToShopParams　ShopDetailからまとめてShopベクトル空間を作成する
func ShopsToShopParams(shops []ShopDetail) ([]ShopParams, error) {
	var shopParams []ShopParams
	for i := range shops {
		params, err := ShopToVec3(&shops[i])
		if err != nil {
			return nil, err
		}
		shopParams = append(shopParams, params)
	}
	return shopParams, nil
}

// ShopToVec3 お店情報をベクトルに変換する
func ShopToVec3(shop *ShopDetail) (ShopParams, error) {
	vec3, err := GenreBudgetToVec3(shop.Genre[0], shop.Genre[1], shop.Budget)
	if err != nil {
		return ShopParams{}, nil
	}
	return ShopParams{
		Shop: shop,
		Vec3: vec3,
	}, nil
}

// ジャンルと予算からベクトルに変換する // 推薦アルゴリズムv0
func GenreBudgetToVec3(genre string, subGenre string, budget string) ([3]float64, error) {
	v, ok := genreVec2[genre]
	if !ok {
		return [3]float64{}, errors.New(fmt.Sprintf("Error: %s", "no genre"))
	}
	if subGenre != "" {
		v2, ok := genreVec2[subGenre]
		if ok {
			v[0] = (1-weight)*v[0] + weight*v2[0]
			v[1] = (1-weight)*v[1] + weight*v2[1]
		}
	}

	z, ok := budgetVec[budget]
	if !ok {
		return [3]float64{}, errors.New(fmt.Sprintf("Error: %s", "no budget"))
	}
	return [3]float64{v[0], v[1], z}, nil
}

// AverageVec3 平均を出す
func AverageVec3(vec3s []ShopParams) [3]float64 {
	var sum [3]float64
	for _, v := range vec3s {
		sum[0] += v.Vec3[0]
		sum[1] += v.Vec3[1]
		sum[2] += v.Vec3[2]
	}
	return [3]float64{sum[0] / float64(len(vec3s)), sum[1] / float64(len(vec3s)), sum[2] / float64(len(vec3s))}
}

type ShopParamSimilarity struct {
	Vec3 ShopParams
	Sim  float64
}

// SimilarityVec3 類似度の高い順にN個取り出す
func FindSimilarVec3(vec3s []ShopParams, vec3 [3]float64, num int) []ShopParams {
	var shopSimilarity []ShopParamSimilarity
	for _, v := range vec3s {
		sim := SimilarityVec3(v.Vec3, vec3)
		shopSimilarity = append(shopSimilarity, ShopParamSimilarity{Vec3: v, Sim: sim})
	}
	sort.Slice(shopSimilarity, func(i, j int) bool {
		return shopSimilarity[i].Sim < shopSimilarity[j].Sim
	})

	var res []ShopParams
	for _, v := range shopSimilarity {
		res = append(res, v.Vec3)
	}
	return res[:num]
}
