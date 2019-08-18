package main

import (
	"net/http"

	currencyHandler "amis/pkg/currency/delivery/http"
	currencyService "amis/pkg/currency/service"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/go-playground/validator.v9"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

type ErrResp struct {
	Message string `json:"message" xml:"message"`
}

func customHTTPErrorHandler(err error, c echo.Context) {
	if err == nil {
		return
	}

	c.Logger().Error(err)

	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	resp := &ErrResp{
		Message: err.Error(),
	}

	c.JSON(code, resp)
}

func main() {
	e := echo.New()

	e.Use(middleware.Recover())

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "ECHO | ${time_rfc3339_nano} | ${status} | ${latency_human} | ${method} ${uri} | remote_ip=${remote_ip} host=${host}\n",
	}))

	e.Validator = &CustomValidator{validator: validator.New()}
	e.HTTPErrorHandler = customHTTPErrorHandler

	currencySvc := currencyService.NewCurrencyService()
	currencyHandler.NewCurrencyHandler(e, currencySvc)

	e.GET("/*", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
