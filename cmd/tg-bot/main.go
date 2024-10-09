package main

import (
	"fmt"
	"log"
	"tg-bot/internal/bot"
	"tg-bot/internal/config"
)

func main() {
	cfg := config.MustLoad()
	log.Println("starting telegram bot...")

	baseURL := fmt.Sprintf("%s:%v/api/v1", cfg.BackendConfig.Host, cfg.BackendConfig.Port)

	tgBot := bot.New(baseURL, cfg.Token, cfg.Timeout)
	tgBot.MustRun()
}
