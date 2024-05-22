package tgmsghandler

import (
	"backend/internal/database"
	"backend/internal/model"
	"backend/internal/ws"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// This function should only be called when the message is a command.
func HandleCommand(bot *tgbotapi.BotAPI, hub *ws.Hub, msg *tgbotapi.Message) error {
	db := database.GetDb()
	msgModel, err := saveTgMessageToDB(db, msg, model.ByGuest)
	if err != nil {
		return err
	}

	if err := broadcastMessage(hub, msgModel, model.ByGuest); err != nil {
		return err
	}

	var response string
	cmd := msg.Command()
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
	if err != nil {
		return err
	}

	res := tgbotapi.NewMessage(msg.Chat.ID, response)
	resMsg, err := bot.Send(res)
	if err != nil {
		return err
	}

	msgModel, err = saveTgMessageToDB(db, &resMsg, model.ByBot)
	if err != nil {
		return err
	}

	if err := broadcastMessage(hub, msgModel, model.ByBot); err != nil {
		return err
	}

	switch cmd {
	case AskCmdWord:
		fallthrough
	case QueryCmdWord:
		fallthrough
	case RequestCmdWord:
		aiResponse, err := getAIResponse(db, msg.Chat.ID)
		if err != nil {
			return err
		}

		aiReplyMsg, err := sendTelegramMessage(bot, msg, aiResponse)
		if err != nil {
			return err
		}

		msgModel, err = saveTgMessageToDB(db, aiReplyMsg, model.ByBot)
		if err != nil {
			return err
		}

		if err := broadcastMessage(hub, msgModel, model.ByBot); err != nil {
			return err
		}
	}

	return nil
}
