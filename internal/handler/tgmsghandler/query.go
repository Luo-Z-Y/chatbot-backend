package tgmsghandler

import (
	"backend/internal/dataaccess/booking"
	"backend/internal/dataaccess/chat"
	"backend/internal/database"
	"backend/internal/model"
	"backend/pkg/error/internalerror"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	QueryCmdWord = "query"
	QueryCmdDesc = "Make a simple query regarding the hotel"
)

func HandleQueryCommand(msg *tgbotapi.Message) (string, error) {
	tgChatID := msg.Chat.ID

	db := database.GetDb()

	chat, err := chat.ReadByTgChatID(db, tgChatID)
	if err != nil {
		if internalerror.IsRecordNotFoundError(err) {
			return NoChatFoundResponse, nil
		}
		return "", err
	}

	bk, err := booking.ReadByChatID(db, chat.ID)
	if err != nil && !internalerror.IsRecordNotFoundError(err) {
		return "", err
	}

	if err := createRequestQuery(db, model.TypeQuery, chat, bk); err != nil {
		return "", err
	}

	return QueryCreatedResponse, nil
}
