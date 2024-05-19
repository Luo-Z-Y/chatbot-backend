package tgmsghandler

import (
	"backend/internal/ws"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// This function should only be called when the message is a command.
func HandleCommand(bot *tgbotapi.BotAPI, hub *ws.Hub, msg *tgbotapi.Message) error {
	cmd := msg.Command()

	var response string
	var err error
	switch cmd {
	case HelpCmdWord:
		response, err = HandleHelpCommand(msg)
	case AuthCmdWord:
		response, err = HandleAuthCommand(msg, hub)
	case StartCmdWord:
		response, err = HandleStartCommand(msg)
	case AskCmdWord:
		response, err = HandleAskCommand(msg)
	case QueryCmdWord:
		response, err = HandleQueryCommand(msg)
	case RequestCmdWord:
		response, err = HandleRequestCommand(msg)
	}

	SendTextMessage(bot, msg, response)

	if err != nil {
		return err
	}

	switch cmd {
	case AskCmdWord:
		fallthrough
	case QueryCmdWord:
		fallthrough
	case RequestCmdWord:
		aiResponse := GetAIResponse(msg.Chat.ID)
		SendTextMessage(bot, msg, aiResponse)
	}

	return nil
}
