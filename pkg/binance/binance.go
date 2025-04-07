package binance

import (
	"context"
	"log"
	"github.com/adshao/go-binance/v2"
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
		log.Printf("Error fetching Binance price: %v", err)
		return 0, err
	}
	if len(prices) == 0 {
		return 0, nil
	}
	price := prices[0].Price
	log.Printf("Binance price for %s: %s", symbol, price)
	return binance.ConvertStringToFloat64(price)
}