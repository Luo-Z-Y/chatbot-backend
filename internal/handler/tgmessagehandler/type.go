package tgmessagehandler

import (
	"backend/internal/ws"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler = func(bot *tgbotapi.BotAPI, hub *ws.Hub, msg *tgbotapi.Message) error
