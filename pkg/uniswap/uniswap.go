package uniswap

import (
	"context"
	"math/big"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/yourmuze/arbitrage-bot/pkg/utils"
)

type UniswapClient struct {
	client *ethclient.Client
}

func NewUniswapClient(rpcURL string) *UniswapClient {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		utils.LogToReadme("Failed to connect to Ethereum RPC: " + err.Error())
		return nil
	}
	return &UniswapClient{client: client}
}

func (u *UniswapClient) GetPrice() float64 {
	// Адрес пула Uniswap V3 ETH/USDT (пример для Mainnet)
	poolAddress := common.HexToAddress("0x11b815efb8f581194ae79006d24e0d814b7697f6") // ETH/USDT pool
	// ABI для вызова slot0 (содержит sqrtPriceX96)
	slot0ABI := `[{"inputs":[],"name":"slot0","outputs":[{"name":"sqrtPriceX96","type":"uint160"}],"stateMutability":"view","type":"function"}]`

	// Создаем контракт
	contract, err := u.client.CallContract(context.Background(), ethclient.CallMsg{
		To:   &poolAddress,
		Data: common.FromHex("0x3850c7bd"), // Keccak-256 хэш "slot0()"
	}, nil)
	if err != nil {
		utils.LogToReadme("Failed to call Uniswap slot0: " + err.Error())
		return 0
	}

	// Парсим sqrtPriceX96 (первые 32 байта ответа)
	sqrtPriceX96 := new(big.Int).SetBytes(contract[:32])

	// Переводим sqrtPriceX96 в цену (для ETH/USDT)
	// Цена = (sqrtPriceX96^2 * 10^decimals) / 2^192
	price := new(big.Float).SetInt(new(big.Int).Mul(sqrtPriceX96, sqrtPriceX96))
	price.Quo(price, new(big.Float).SetInt(big.NewInt(2).Exp(big.NewInt(2), big.NewInt(192), nil)))
	price.Mul(price, big.NewFloat(1e6)) // USDT имеет 6 decimals

	priceFloat, _ := price.Float64()
	utils.LogToReadme("Uniswap ETH/USDT price: " + price.Text('f', 2))
	return priceFloat
}