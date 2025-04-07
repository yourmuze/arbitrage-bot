package binance

import (
	"context"
	"strconv"
	"github.com/adshao/go-binance/v2"
	"github.com/yourmuze/arbitrage-bot/pkg/utils"
)

type BinanceClient struct {
	client *binance.Client
}

func NewBinanceClient(apiKey, secret string) *BinanceClient {
	return &BinanceClient{
		client: binance.NewClient(apiKey, secret),
	}
}

func (b *BinanceClient) GetPrice(symbol string) (float64, error) {
	prices, err := b.client.NewListPricesService().Symbol(symbol).Do(context.Background())
	if err != nil {
		utils.LogToReadme("Error fetching Binance price: " + err.Error())
		return 0, err
	}
	if len(prices) == 0 {
		utils.LogToReadme("No price data returned for " + symbol)
		return 0, nil
	}
	price := prices[0].Price
	utils.LogToReadme("Binance price for " + symbol + ": " + price)
	priceFloat, err := strconv.ParseFloat(price, 64)
	if err != nil {
		utils.LogToReadme("Error converting price to float: " + err.Error())
		return 0, err
	}
	return priceFloat, nil
}