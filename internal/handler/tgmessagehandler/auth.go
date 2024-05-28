package tgmessagehandler

import (
	"backend/internal/database"
	"backend/internal/ws"
	"errors"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	AuthCmdWord = "auth"
	AuthCmdDesc = "Authenticate yourself to make a request. Please provide your credentials"
)

const (
	AuthRequestMadeResponse = "Authentication request made. Pending response from staff."
)

var (
	CredentialsNotFound = errors.New("Credentials not found")
)

func HandleAuthCommand(bot *tgbotapi.BotAPI, hub *ws.Hub, msg *tgbotapi.Message) error {
	db := database.GetDb()

	chat, err := readChatByTgChatIDOrCreate(db, msg.Chat.ID)
	if err != nil {
		return err
	}

	cred, err := extractCredentials(msg.Text)
	if err != nil {
		return err
	}

	if err := broadcastAuthRequest(hub, chat, cred); err != nil {
		return err
	}

	_, err = SendTelegramMessage(bot, msg, AuthRequestMadeResponse)
	return err
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
