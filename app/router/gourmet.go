package router

import (
	"net/http"

	"github.com/22hack12spring/backend/model"
	"github.com/22hack12spring/backend/service"
	"github.com/labstack/echo/v4"
)

type GourmetStartRequest struct {
	Lat     float64 `json:"lat" validate:"required_without=Station"`
	Lng     float64 `json:"lng" validate:"required_without=Station"`
	Station string  `json:"station"`
}

type GourmetStartResponse struct {
	ID        string              `json:"id"`
	Questions []*service.ShopData `json:"questions"`
}

type GourmetAnswerRequest struct {
	ID      string           `json:"id" validate:"required,uuid"`
	Answers []service.Answer `json:"answers" validate:"dive,required,eq=3"`
}

type GourmetAnswerResponse struct {
	ID   string              `json:"id"`
	Shop *service.ShopDetail `json:"shop"`
}

// API:POST /gourmet/start
func (h *Handlers) PostGourmetStart(c echo.Context) error {
	var param GourmetStartRequest
	err := validatedBind(c, &param)
	if err != nil {
		return err
	}
	arg := model.ToSearchArgs(param.Lat, param.Lng, param.Station)

	// uuidの生成とデータベースへの登録
	searches, err := h.Repo.CreateSearch(c.Request().Context(), arg)
	if err != nil {
		return err
	}

	// 質問の生成
	questions, ids, err := h.Service.GenerateQuestions(c.Request().Context(), arg, searches.ID)
	if err != nil {
		return err
	}
	// 送った質問を保存
	_, err = h.Repo.CreateQuestions(c.Request().Context(), model.QuestionArgs{
		ShopIds:  ids,
		SearchId: searches.ID,
	})
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, GourmetStartResponse{
		ID:        searches.ID,
		Questions: questions,
	})
}

// API:POST /gourmet/answer
func (h *Handlers) PostGourmetAnswer(c echo.Context) error {
	var param GourmetAnswerRequest
	err := validatedBind(c, &param)
	if err != nil {
		return err
	}
	shop, err := h.Service.GenerateRecommend(c.Request().Context(), param.ID, param.Answers)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, GourmetAnswerResponse{
		ID:   param.ID,
		Shop: shop,
	})
}
