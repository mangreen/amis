package http

import (
	"amis/pkg/currency"
	"amis/pkg/currency/service"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo"
)

type currencyHandler struct {
	currencySvc currency.CurrencyService
}

type currencyReq struct {
	Coin  string `json:"coin" form:"coin" query:"coin" validate:"required"`
	Start string `json:"start" form:"start" query:"start" validate:"required"`
}

var (
	local, _ = time.LoadLocation("Asia/Taipei")
	begin    = time.Date(2019, time.January, 1, 0, 0, 0, 0, local)
)

func NewCurrencyHandler(e *echo.Echo, currencySvc currency.CurrencyService) *currencyHandler {
	handler := &currencyHandler{
		currencySvc: currencySvc,
	}

	e.GET("/api/currencies", handler.list)

	return handler
}

func (h *currencyHandler) list(c echo.Context) error {
	ctx := c.Request().Context()

	qs := new(currencyReq)

	if err := c.Bind(qs); err != nil {
		c.Logger().Error(err)
		return err
	}

	if err := c.Validate(qs); err != nil {
		c.Logger().Error(err)
		return err
	}

	if _, ok := service.CoinMap[qs.Coin]; !ok {
		err := errors.New("Coin is not support")
		c.Logger().Error(err)
		return err
	}

	startTime, err := time.ParseInLocation("20060102", qs.Start, local)
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	if begin.After(startTime) {
		err := errors.New("Start is out range")
		c.Logger().Error(err)
		return err
	}

	st := startTime.Unix()

	crcy, err := h.currencySvc.GetCurrency(ctx, strings.ToLower(qs.Coin), st)
	if err != nil {
		c.Logger().Error(err)
		return err
	}

	crcy.Time = qs.Start

	return c.JSON(http.StatusOK, crcy)
}
