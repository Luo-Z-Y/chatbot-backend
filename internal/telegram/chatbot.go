package telegram

import (
	"backend/internal/configs"
	"backend/internal/handler/tgeditedmsghandler"
	"backend/internal/handler/tgmsghandler"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Documentation: https://pkg.go.dev/github.com/go-telegram-bot-api/telegram-bot-api/v5#section-documentation
// Getting started: https://go-telegram-bot-api.dev/

func StartChatbot(cfg *configs.Config) {
	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		panic(err)
	}

	if cfg.GoEnv == "development" {
		bot.Debug = true
	}

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)
	handleUpdates(bot, updates)
}

func handleUpdates(bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) {
	// At most 1 field in an update will be set to a non-nil value
	// https://go-telegram-bot-api.dev/getting-started/important-notes
	for update := range updates {
		if update.Message != nil && update.Message.IsCommand() {
			go tgmsghandler.HandleCommand(bot, update.Message)
			continue
		}

		if update.Message != nil {
			go tgmsghandler.HandleMessage(bot, update.Message)
			continue
		}

		if update.EditedMessage != nil {
			go tgeditedmsghandler.HandleMessage(bot, update.EditedMessage)
			continue
		}
	}
}
