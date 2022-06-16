package router

import (
	"net/http"

	"github.com/22hack12spring/backend/service"
	"github.com/labstack/echo/v4"
)

type GourmetStartRequest struct {
	Lat     float64 `json:"lat"`
	Lng     float64 `json:"lng" validate:"required_with=Lat"`
	Station string  `json:"station" validate:"required_without=Lat Lng"`
}

type GourmetAnswerRequest struct {
	ID      string           `json:"id" validate:"required,uuid"`
	Answers []service.Answer `json:"answers" validate:"dive,required,eq=3"`
}

// API:POST /gourmet/start
func (h *Handlers) PostGourmetStart(c echo.Context) error {
	var param GourmetStartRequest
	err := validatedBind(c, &param)
	if err != nil {
		return err
	}
	// uuidの生成とデータベースへの登録

	// 質問の生成
	questions, err := h.Service.GenerateQuestions(c.Request().Context(), param.Station, param.Lat, param.Lng)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, questions)
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
	return c.JSON(http.StatusOK, shop)
}
