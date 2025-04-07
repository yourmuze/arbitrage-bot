package main

import (
	"os/exec"
	"github.com/yourmuze/arbitrage-bot/pkg/arbitrage"
	"github.com/yourmuze/arbitrage-bot/pkg/binance"
	"github.com/yourmuze/arbitrage-bot/pkg/uniswap"
	"github.com/yourmuze/arbitrage-bot/pkg/utils"
)

func pushToGit() {
	cmd := exec.Command("git", "add", "README.md")
	if err := cmd.Run(); err != nil {
		utils.LogToReadme("Failed to git add: " + err.Error())
		return
	}

	cmd = exec.Command("git", "commit", "-m", "Update logs")
	if err := cmd.Run(); err != nil {
		utils.LogToReadme("Failed to git commit: " + err.Error())
		return
	}

	cmd = exec.Command("git", "push", "origin", "master")
	if err := cmd.Run(); err != nil {
		utils.LogToReadme("Failed to git push: " + err.Error())
		return
	}
}

func main() {
	utils.InitReadme() // Инициализируем README
	config := utils.LoadConfig()

	binanceClient := binance.NewBinanceClient(config.Binance.APIKey, config.Binance.Secret)
	uniswapClient := uniswap.NewUniswapClient(config.Ethereum.RPCURL)
	bot := arbitrage.NewArbitrageBot(binanceClient, uniswapClient, config)

	bot.CheckOpportunity(config.Symbols.Pair)
	pushToGit() // Пушим логи
}