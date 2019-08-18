package service

import (
	"amis/pkg/currency"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/imroc/req"
	"github.com/maicoin/max-exchange-api-go"
)

type currencyService struct {
	maxclient max.API
}

type geckoMarketChart struct {
	Prices       [][]float64 `json:"prices"`
	MarketCaps   [][]float64 `json:"market_caps"`
	TotalVolumes [][]float64 `json:"total_volumes"`
}

var (
	coingecko = "https://api.coingecko.com/api/v3/coins/%s/market_chart/range"

	CoinMap = map[string]string{
		"btc": "bitcoin",
		"eth": "ethereum",
	}
)

func NewCurrencyService() currency.CurrencyService {
	return &currencyService{}
}

func (svc *currencyService) GetCurrency(coin string, start int64) (*currency.Currency, error) {
	crcy := &currency.Currency{}

	maxTwdPrice, _ := svc.GetMax(coin+"twd", int32(start))
	maxUsdPrice, _ := svc.GetMax(coin+"usdt", int32(start))

	geckoTwdPrice, _ := svc.GetGecko(CoinMap[coin], "twd", start)
	geckoUsdPrice, _ := svc.GetGecko(CoinMap[coin], "usd", start)

	var twdPrice float64
	var usdPrice float64
	var count float64

	if maxTwdPrice != 0 && maxUsdPrice != 0 {
		twdPrice += maxTwdPrice
		usdPrice += maxUsdPrice
		crcy.Sources = append(crcy.Sources, "MAX SDK")
		count++
	}

	if geckoTwdPrice != 0 && geckoUsdPrice != 0 {
		twdPrice += geckoTwdPrice
		usdPrice += geckoUsdPrice
		crcy.Sources = append(crcy.Sources, "coingecko.com")
		count++
	}

	crcy.TWD = int(math.Round(twdPrice / count))

	usdPrice, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", usdPrice/count), 64)
	crcy.USD = usdPrice

	return crcy, nil
}

func (svc *currencyService) GetMax(market string, timestamp int32) (float64, error) {
	maxclient := max.NewClient()
	defer maxclient.Close()

	results, err := maxclient.K(context.Background(), market, max.Limit(1), max.Period(1), max.Timestamp(timestamp))
	if err != nil {
		return 0, err
	}

	if len(results) == 0 {
		return 0, errors.New("Records not found")
	}

	return results[0].Close, nil
}

func (svc *currencyService) GetGecko(id string, vs string, from int64) (float64, error) {
	to := time.Unix(from, 0).Add(time.Hour * time.Duration(1)).Unix()
	param := req.Param{
		"vs_currency": vs,
		"from":        from,
		"to":          to,
	}

	url := fmt.Sprintf(coingecko, id)
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := req.Get(url, client, param)
	if err != nil {
		return 0, err
	}

	byt := resp.Bytes()
	if err != nil {
		return 0, err
	}

	var gmc geckoMarketChart
	if err := json.Unmarshal(byt, &gmc); err != nil {
		return 0, err
	}

	if gmc.Prices == nil || len(gmc.Prices) == 0 || len(gmc.Prices[0]) < 1 {
		return 0, errors.New("Records not found")
	}

	return gmc.Prices[0][1], nil
}
