package router

import (
	"net/http"

	"github.com/22hack12spring/backend/service"
	"github.com/labstack/echo/v4"
)

type GourmetStartRequest struct {
	Lat     float64 `json:"lat"`
	Lng     float64 `json:"lng"`
	Station string  `json:"station"`
}

type GourmetAnswerRequest struct {
	ID      string           `json:"id"`
	Answers []service.Answer `json:"answers"`
}

func (h *Handlers) PostGourmetStart(c echo.Context) error {
	var param GourmetStartRequest
	if err := c.Bind(&param); err != nil {
		return err
	}
	// check param
	if (param.Lat == 0 || param.Lng == 0) && param.Station == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid param")
	}
	questions, err := h.Service.GenerateQuestions(c.Request().Context(), param.Station, param.Lat, param.Lng)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, questions)
}

func (h *Handlers) PostGourmetAnswer(c echo.Context) error {
	var param GourmetAnswerRequest
	if err := c.Bind(&param); err != nil {
		return err
	}
	// check param
	// WANT: マジックナンバーを消す
	if len(param.Answers) != service.QUESTION_NUM || len(param.ID) != 36 {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid param")
	}
	shop, err := h.Service.GenerateRecommend(c.Request().Context(), param.ID, param.Answers)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, shop)
}
