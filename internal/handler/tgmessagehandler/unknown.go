package tgmessagehandler

import (
	"backend/internal/ws"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var unknownCommandResponse = fmt.Sprintf(
	"Sorry, I don't understand this command. Please use /%s to see the available commands.",
	HelpCmdWord,
)

func HandleUnknownCommand(bot *tgbotapi.BotAPI, _ *ws.Hub, msg *tgbotapi.Message) error {
	_, err := SendTelegramMessage(bot, msg, unknownCommandResponse)
	return err
}
