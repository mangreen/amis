package currency

import (
	"context"
)

//go:generate mockgen -destination ./mock/mock_service.go -package mocsvc amis/pkg/currency/service CurrencyService
type CurrencyService interface {
	GetCurrency(ctx context.Context, coin string, start int64) (*Currency, error)
}

type Currency struct {
	Sources []string `json:"sources"`
	TWD     int      `json:"twd"`
	USD     float64  `json:"usd"`
	Time    string   `json:"time"`
}
