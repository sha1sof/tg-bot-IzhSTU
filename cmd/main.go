package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"tg-bot/pkg/telegram"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(inputAPI())
	if err != nil {
		log.Fatal(err)
	}

	telegramBot := telegram.NewBot(bot)
	if err := telegramBot.Start(); err != nil {
		log.Fatal(err)
	}
}

func inputAPI() string {
	fmt.Println("Введите API ключ//Enter API key")
	var key string
	fmt.Scan(&key)
	return key
}
