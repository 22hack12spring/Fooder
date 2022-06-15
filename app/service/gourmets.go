package service

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/22hack12spring/backend/model"
)

func (s *Services) GetGourmetsData(ctx context.Context, args model.SearchArgs) ([]ShopDetail, error) {
	var shops []ShopDetail
	cache, err := s.Repo.GetGourmetsCache(ctx, args)
	if err == sql.ErrNoRows {
		// キャッシュがない場合はAPIを叩く
		rawData, err := s.Repo.GetGourmetsRawAPI(args)
		if err != nil {
			return nil, err
		}
		shops, err = s.trimRawDataToShopDetail(rawData)
		if err != nil {
			return nil, err
		}
		saveData, err := json.Marshal(shops)
		if err != nil {
			return nil, err
		}

		// キャッシュに保存
		err = s.Repo.InsertGourmetsCache(ctx, args, string(saveData))
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	} else {
		// 正常な場合
		// TODO: キャッシュ時間を確認する
		// JSONをパース
		err = json.Unmarshal([]byte(cache.Shops), &shops)
		if err != nil {
			return nil, err
		}
	}

	return shops, nil
}

func (s *Services) trimRawDataToShopDetail(raw string) ([]ShopDetail, error) {
	var gourmets model.GourmetsDetail
	gourmets, err := s.Repo.ParseGourmetsJSON(raw)
	if err != nil {
		return nil, err
	}

	var shops []ShopDetail
	for _, shop := range gourmets.Shop {
		genres := [...]string{shop.Genre.Name, shop.SubGenre.Name}
		shopDetail := ShopDetail{
			Id:        shop.ID,
			Name:      shop.Name,
			LogoImage: shop.LogoImage,
			Address:   shop.Address,
			Lat:       shop.Lat,
			Lng:       shop.Lng,
			Open:      shop.Open,
			Close:     shop.Close,
			Url:       shop.Urls.Pc,
			Photo:     shop.Photo.Pc.M,
			Lunch:     shop.Lunch,
			Genre:     genres[:],
		}
		shops = append(shops, shopDetail)
	}

	return shops, nil
}
