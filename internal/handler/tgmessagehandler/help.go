package tgmessagehandler

import (
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
	"%s\n\n/%s - %s\n/%s - %s\n/%s - %s\n/%s - %s\n/%s - %s",
	HelpIntro,
	HelpCmdWord, HelpCmdDesc,
	AuthCmdWord, AuthCmdDesc,
	StartCmdWord, StartCmdDesc,
	QueryCmdWord, QueryCmdDesc,
	RequestCmdWord, RequestCmdDesc,
)

func HandleHelpCommand(bot *tgbotapi.BotAPI, hub *ws.Hub, msg *tgbotapi.Message) error {
	_, err := SendTelegramMessage(bot, msg, helpMsg)
	return err
}
