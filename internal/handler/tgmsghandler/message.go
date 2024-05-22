package tgmsghandler

import (
	"backend/internal/database"
	"backend/internal/model"
	"backend/internal/ws"
	"backend/pkg/error/internalerror"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	NoChatQueryFoundResponse = fmt.Sprintf(
		"Chat or query not found, please start a chat with /%s followed by making a new query with /%s, /%s, or /%s",
		StartCmdWord,
		AskCmdWord,
		QueryCmdWord,
		RequestCmdWord,
	)
)

func HandleMessage(bot *tgbotapi.BotAPI, hub *ws.Hub, tgMsg *tgbotapi.Message) error {
	db := database.GetDb()

	msgModel, err := saveTgMessageToDB(db, tgMsg, model.ByGuest)
	if err != nil {
		if internalerror.IsRecordNotFoundError(err) {
			_, err := SendTelegramMessage(bot, tgMsg, NoChatQueryFoundResponse)
			return err
		}
		return err
	}

	if err := broadcastMessage(hub, msgModel); err != nil {
		return err
	}

	res, err := getAIResponse(db, tgMsg.Chat.ID)
	if err != nil {
		return err
	}

	_, err = SendTelegramMessage(bot, tgMsg, res)
	return err
}
