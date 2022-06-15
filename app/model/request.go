package model

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type GourmetsRequest interface {
	GetGourmetsRawAPI(args SearchArgs) (string, error)
}

// APIリクエストを送って、グルメ一覧(json)のデータを取得する
func (repo *SqlxRepository) GetGourmetsRawAPI(args SearchArgs) (string, error) {
	baseurl := "https://webservice.recruit.co.jp/hotpepper/gourmet/v1/"

	request, err := http.NewRequest("GET", baseurl, nil)
	if err != nil {
		return "", err
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return "", errors.New(fmt.Sprintf("Error: %s", "no API_KEY"))
	}

	params := request.URL.Query()
	params.Add("key", apiKey)
	// TODO: 現状だとキーワード検索なので、small_area検索に変更したい
	if args.Station != nil {
		params.Add("keyword", *args.Station)
	} else if args.Lat != nil && args.Lng != nil {
		params.Add("lat", fmt.Sprintf("%f", *args.Lat))
		params.Add("lng", fmt.Sprintf("%f", *args.Lng))
		// 1000m以内
		params.Add("range", "3")
	} else {
		// error
		return "", errors.New(fmt.Sprintf("Error: %s", "Invalid args"))
	}
	params.Add("count", "100")
	params.Add("format", "json")

	request.URL.RawQuery = params.Encode()

	timeout := time.Duration(5 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}

	response, err := client.Do(request)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
