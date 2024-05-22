package tgmsghandler

import (
	"backend/internal/ws"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// This function should only be called when the message is a command.
func HandleCommand(bot *tgbotapi.BotAPI, hub *ws.Hub, msg *tgbotapi.Message) {
	var err error = nil

	switch msg.Command() {
	case HelpCmdWord:
		err = HandleHelpCommand(bot, hub, msg)
	case AuthCmdWord:
		err = HandleAuthCommand(bot, hub, msg)
	case StartCmdWord:
		err = HandleStartCommand(bot, hub, msg)
	case AskCmdWord:
		err = HandleAskCommand(bot, hub, msg)
	case QueryCmdWord:
		err = HandleQueryCommand(bot, hub, msg)
	case RequestCmdWord:
		err = HandleRequestCommand(bot, hub, msg)
	}
	if err != nil {
		_, _ = sendTelegramMessage(bot, msg, err.Error())
	}

	// res := tgbotapi.NewMessage(msg.Chat.ID, response)
	// resMsg, err := bot.Send(res)
	// if err != nil {
	// 	return err
	// }

	// msgModel, err := saveTgMessageToDB(db, &resMsg, model.ByBot)
	// if err != nil {
	// 	return err
	// }

	// if err := broadcastMessage(hub, msgModel, model.ByBot); err != nil {
	// 	return err
	// }

	// switch cmd {
	// case AskCmdWord:
	// 	fallthrough
	// case QueryCmdWord:
	// 	fallthrough
	// case RequestCmdWord:
	// 	aiResponse, err := getAIResponse(db, msg.Chat.ID)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	aiReplyMsg, err := sendTelegramMessage(bot, msg, aiResponse)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	msgModel, err = saveTgMessageToDB(db, aiReplyMsg, model.ByBot)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	if err := broadcastMessage(hub, msgModel, model.ByBot); err != nil {
	// 		return err
	// 	}
	// }

	// return nil
}
