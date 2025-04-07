package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
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
	Telegram struct {
		BotToken string `json:"bot_token"`
		ChatID   string `json:"chat_id"`
	} `json:"telegram"`
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

func LogToReadme(message string) {
	file, err := os.OpenFile("README.md", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Failed to open README: %v", err)
		return
	}
	defer file.Close()

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logEntry := fmt.Sprintf("%s - %s\n", timestamp, message)
	if _, err := file.WriteString(logEntry); err != nil {
		log.Printf("Failed to write to README: %v", err)
	}
}

func InitReadme() {
	if _, err := os.Stat("README.md"); os.IsNotExist(err) {
		file, err := os.Create("README.md")
		if err != nil {
			log.Fatalf("Failed to create README: %v", err)
		}
		defer file.Close()
		file.WriteString("# Arbitrage Bot Logs\n\n")
	}
}

func SendTelegramMessage(config Config, message string) error {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", config.Telegram.BotToken)
	data := url.Values{
		"chat_id": {config.Telegram.ChatID},
		"text":    {message},
	}

	resp, err := http.PostForm(apiURL, data)
	if err != nil {
		return fmt.Errorf("failed to send Telegram message: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Telegram API returned non-OK status: %d", resp.StatusCode)
	}
	return nil
}