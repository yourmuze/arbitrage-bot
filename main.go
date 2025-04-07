package main

import (
	"log"
	"github.com/yourusername/arbitrage-bot/pkg/arbitrage"
	"github.com/yourusername/arbitrage-bot/pkg/binance"
	"github.com/yourusername/arbitrage-bot/pkg/uniswap"
	"github.com/yourusername/arbitrage-bot/pkg/utils"
)

func main() {
	log.Println("Starting arbitrage bot...")
	config := utils.LoadConfig()

	binanceClient := binance.NewBinanceClient(config.Binance.APIKey, config.Binance.Secret)
	uniswapClient := uniswap.NewUniswapClient(config.Ethereum.RPCURL)
	bot := arbitrage.NewArbitrageBot(binanceClient, uniswapClient)

	bot.CheckOpportunity(config.Symbols.Pair)
}