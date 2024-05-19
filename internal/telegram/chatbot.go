package telegram

import (
	"backend/internal/configs"
	"backend/internal/handler/tgmsghandler"
	"backend/internal/ws"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var botAPI *tgbotapi.BotAPI

func GetBotAPI() *tgbotapi.BotAPI {
	return botAPI
}

// Documentation: https://pkg.go.dev/github.com/go-telegram-bot-api/telegram-bot-api/v5#section-documentation
// Getting started: https://go-telegram-bot-api.dev/

func StartChatbot(cfg *configs.Config, hub *ws.Hub) {
	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		panic(err)
	} else {
		botAPI = bot
	}

	if cfg.GoEnv == "development" {
		bot.Debug = true
	}

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)
	handleUpdates(bot, hub, updates)
}

func handleUpdates(bot *tgbotapi.BotAPI, hub *ws.Hub, updates tgbotapi.UpdatesChannel) {
	// At most 1 field in an update will be set to a non-nil value
	// https://go-telegram-bot-api.dev/getting-started/important-notes
	for update := range updates {
		if update.Message != nil && update.Message.IsCommand() {
			go tgmsghandler.HandleCommand(bot, hub, update.Message)
			continue
		}

		if update.Message != nil {
			go tgmsghandler.HandleMessage(bot, hub, update.Message)
			continue
		}
	}
}
