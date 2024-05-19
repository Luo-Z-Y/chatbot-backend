package tgmsghandler

import (
	"backend/internal/dataaccess/chat"
	"backend/internal/database"
	"backend/internal/ws"
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	AuthCmdWord = "auth"
	AuthCmdDesc = "Authenticate yourself to make a request. Please provide your credentials"
)

var (
	CredentialsNotFound = errors.New("Credentials not found")
)

func HandleAuthCommand(msg *tgbotapi.Message, hub *ws.Hub) (string, error) {
	tgChatID := msg.Chat.ID

	db := database.GetDb()

	_, err := chat.ReadByTgChatID(db, tgChatID)
	if err != nil {
		return NoChatFoundResponse, err
	}

	_, err = extractCredentials(msg.Text)
	if err != nil {
		return err.Error(), err
	}

	return "New query created", nil
}

// Commands are prefixed with a slash (/cmd args)
func extractCredentials(text string) (string, error) {
	for i, c := range text {
		if c == ' ' {
			if len(text) <= i+1 {
				return "", CredentialsNotFound
			}
			return text[i+1:], nil
		}
	}

	return "", CredentialsNotFound
}
