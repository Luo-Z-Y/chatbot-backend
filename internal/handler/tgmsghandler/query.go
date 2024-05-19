package tgmsghandler

import (
	"backend/internal/dataaccess/booking"
	"backend/internal/dataaccess/chat"
	"backend/internal/database"
	"backend/internal/model"

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
		return NoChatFoundResponse, err
	}

	bk, _ := booking.ReadByChatID(db, chat.ID)

	if err := createRequestQueryTransaction(db, msg, chat, bk, model.TypeQuery); err != nil {
		return "An error occurred while creating a new query", err
	}

	return "New query created", nil
}
