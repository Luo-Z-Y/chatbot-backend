package tgmsghandler

import (
	"backend/internal/dataaccess/chat"
	"backend/internal/database"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	AskCmdWord = "ask"
	AskCmdDesc = "Ask anything, and we will try to help you"
)

const (
	NoChatFoundResponse = "Chat not found, please start a new chat with /" + StartCmdWord
)

func HandleAskCommand(msg *tgbotapi.Message) (string, error) {
	tgChatID := msg.Chat.ID

	db := database.GetDb()

	chat, err := chat.ReadByTgChatID(db, tgChatID)
	if err != nil {
		return NoChatFoundResponse, err
	}

	// --- Perform query to AI to determine the type of query --- //
	queryType := _tempRandomType()

	if err := createRequestQueryTransaction(db, msg, chat, nil, queryType); err != nil {
		return "An error occurred while creating a new query", err
	}

	return "New query created", nil
}
