package tgmsghandler

import (
	"backend/internal/dataaccess/chat"
	"backend/internal/database"
	"backend/pkg/error/internalerror"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	AskCmdWord = "ask"
	AskCmdDesc = "Ask anything, and we will try to help you"
)

const (
	NoChatFoundResponse  = "Chat not found, please start a new chat with /" + StartCmdWord
	QueryCreatedResponse = "New query created, please wait for a response"
)

func HandleAskCommand(msg *tgbotapi.Message) (string, error) {
	tgChatID := msg.Chat.ID

	db := database.GetDb()

	chat, err := chat.ReadByTgChatID(db, tgChatID)
	if err != nil {
		if internalerror.IsRecordNotFoundError(err) {
			return NoChatFoundResponse, nil
		}
		return "", err
	}

	// --- Perform query to AI to determine the type of query --- //
	queryType := _tempRandomType()

	if err := createRequestQuery(db, queryType, chat, nil); err != nil {
		return "", err
	}

	return QueryCreatedResponse, nil
}
