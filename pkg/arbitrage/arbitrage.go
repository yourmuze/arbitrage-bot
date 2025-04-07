package arbitrage

import (
	"log"
	"github.com/yourusername/arbitrage-bot/pkg/binance"
	"github.com/yourusername/arbitrage-bot/pkg/uniswap"
)

type ArbitrageBot struct {
	binance *binance.BinanceClient
	uniswap *uniswap.UniswapClient
}

func NewArbitrageBot(binance *binance.BinanceClient, uniswap *uniswap.UniswapClient) *ArbitrageBot {
	return &ArbitrageBot{binance: binance, uniswap: uniswap}
}

func (a *ArbitrageBot) CheckOpportunity(symbol string) {
	binancePrice, err := a.binance.GetPrice(symbol)
	if err != nil {
		return
	}
	uniswapPrice := a.uniswap.GetPrice()

	log.Printf("Binance: %f, Uniswap: %f", binancePrice, uniswapPrice)
	if binancePrice > uniswapPrice {
		log.Println("Opportunity: Buy on Uniswap, Sell on Binance")
	} else if uniswapPrice > binancePrice {
		log.Println("Opportunity: Buy on Binance, Sell on Uniswap")
	} else {
		log.Println("No arbitrage opportunity")
	}
}