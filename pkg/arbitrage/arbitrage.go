package arbitrage

import (
	"fmt"
	"github.com/yourmuze/arbitrage-bot/pkg/binance"
	"github.com/yourmuze/arbitrage-bot/pkg/uniswap"
	"github.com/yourmuze/arbitrage-bot/pkg/utils"
)

type ArbitrageBot struct {
	binance *binance.BinanceClient
	uniswap *uniswap.UniswapClient
	balance float64 // Тестовый баланс в USDT
	config  utils.Config
}

func NewArbitrageBot(binance *binance.BinanceClient, uniswap *uniswap.UniswapClient, config utils.Config) *ArbitrageBot {
	return &ArbitrageBot{
		binance: binance,
		uniswap: uniswap,
		balance: 5000.0, // Начальный баланс 5000 USDT
		config:  config,
	}
}

func (a *ArbitrageBot) CheckOpportunity(symbol string) {
	binancePrice, err := a.binance.GetPrice(symbol)
	if err != nil {
		return
	}
	uniswapPrice := a.uniswap.GetPrice()

	utils.LogToReadme(fmt.Sprintf("Balance: %.2f USDT", a.balance))
	utils.LogToReadme(fmt.Sprintf("Binance: %.2f, Uniswap: %.2f", binancePrice, uniswapPrice))

	// Симуляция сделки
	tradeAmount := 1.0 // 1 ETH для теста
	if binancePrice > uniswapPrice {
		profit := (binancePrice - uniswapPrice) * tradeAmount
		if profit > 0 && a.balance >= uniswapPrice*tradeAmount {
			a.balance += profit // Обновляем баланс
			msg := fmt.Sprintf("SUCCESS: Buy on Uniswap (%.2f), Sell on Binance (%.2f), Profit: %.2f USDT, New Balance: %.2f USDT",
				uniswapPrice, binancePrice, profit, a.balance)
			utils.LogToReadme(msg)
			utils.SendTelegramMessage(a.config, msg)
		}
	} else if uniswapPrice > binancePrice {
		profit := (uniswapPrice - binancePrice) * tradeAmount
		if profit > 0 && a.balance >= binancePrice*tradeAmount {
			a.balance += profit // Обновляем баланс
			msg := fmt.Sprintf("SUCCESS: Buy on Binance (%.2f), Sell on Uniswap (%.2f), Profit: %.2f USDT, New Balance: %.2f USDT",
				binancePrice, uniswapPrice, profit, a.balance)
			utils.LogToReadme(msg)
			utils.SendTelegramMessage(a.config, msg)
		}
	} else {
		utils.LogToReadme("No arbitrage opportunity")
	}
}