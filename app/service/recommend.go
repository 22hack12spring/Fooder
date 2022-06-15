package service

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
func (s *Services) GenerateRecommend(uuid string, answers []Answer) (*ShopDetail, error) {
	return nil, nil
}
