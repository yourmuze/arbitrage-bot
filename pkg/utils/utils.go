package utils

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Binance struct {
		APIKey string `json:"api_key"`
		Secret string `json:"secret"`
	} `json:"binance"`
	Ethereum struct {
		RPCURL string `json:"rpc_url"`
	} `json:"ethereum"`
	Symbols struct {
		Pair string `json:"pair"`
	} `json:"symbols"`
}

func LoadConfig() Config {
	file, err := os.Open("config/config.json")
	if err != nil {
		log.Fatalf("Failed to open config: %v", err)
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		log.Fatalf("Failed to decode config: %v", err)
	}
	return config
}