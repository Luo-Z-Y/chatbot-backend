package tgmsghandler

import (
	"backend/internal/model"
	"backend/internal/ws"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	HelpIntro   = "I'm a bot that can help you with your tasks. Here are the available commands:"
	HelpCmdWord = "help"
	HelpCmdDesc = "Show help message"
)

var helpMsg = fmt.Sprintf(
	"%s\n\n%s - %s\n%s - %s\n%s - %s\n%s - %s\n%s - %s\n%s - %s",
	HelpIntro,
	HelpCmdWord, HelpCmdDesc,
	AuthCmdWord, AuthCmdDesc,
	StartCmdWord, StartCmdDesc,
	AskCmdWord, AskCmdDesc,
	QueryCmdWord, QueryCmdDesc,
	RequestCmdWord, RequestCmdDesc,
)

func HandleHelpCommand(bot *tgbotapi.BotAPI, hub *ws.Hub, msg *tgbotapi.Message) error {
	// Since all messages requires a non-null requestquery, and users may use /help before starting a chat,
	// We cannot save this message to the database but only broadcast it to the websocket hub.
	if err := broadcastDanglingMessage(hub, msg, model.ByGuest); err != nil {
		return err
	}

	res, err := SendTelegramMessage(bot, msg, helpMsg)
	if err != nil {
		return err
	}

	return broadcastDanglingMessage(hub, res, model.ByBot)
}
