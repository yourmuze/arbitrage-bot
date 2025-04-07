package uniswap

import (
	"log"
)

type UniswapClient struct {
	rpcURL string
}

func NewUniswapClient(rpcURL string) *UniswapClient {
	return &UniswapClient{rpcURL: rpcURL}
}

func (u *UniswapClient) GetPrice() float64 {
	// Заглушка, позже добавим реальную логику
	log.Println("Uniswap price fetched (placeholder)")
	return 2000
}