package tgmsghandler

import (
	"backend/internal/api"
	"backend/internal/dataaccess/chat"
	"backend/internal/database"
	"backend/internal/model"
	"backend/internal/viewmodel"
	"backend/internal/ws"
	"backend/pkg/error/internalerror"
	"encoding/json"
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
	// Since all messages requires a non-null requestquery, and users may use /auth before starting a query,
	// We cannot save this message to the database but only broadcast it to the websocket hub.
	if err := broadcastDanglingMessage(hub, msg, model.ByGuest); err != nil {
		return err
	}

	db := database.GetDb()

	chat, err := chat.ReadByTgChatID(db, msg.Chat.ID)
	if err != nil {
		if internalerror.IsRecordNotFoundError(err) {
			_, err := sendTelegramMessage(bot, msg, NoChatFoundResponse)
			return err
		}
		return err
	}

	cred, err := extractCredentials(msg.Text)
	if err != nil {
		return err
	}

	if err := broadcastAuthRequest(hub, chat, cred); err != nil {
		return err
	}

	res, err := sendTelegramMessage(bot, msg, AuthRequestMadeResponse)
	if err != nil {
		return err
	}

	return broadcastDanglingMessage(hub, res, model.ByBot)
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

func broadcastAuthRequest(hub *ws.Hub, chat *model.Chat, cred string) error {
	tgAuthView := viewmodel.TgAuthView{
		ChatID:      chat.ID,
		Credentials: cred,
	}

	msgStruct := api.WebSocketMessage{
		Type: api.AuthType,
		Data: tgAuthView,
	}

	msgBytes, err := json.Marshal(msgStruct)
	if err != nil {
		return err
	}

	hub.Broadcast <- msgBytes
	return nil
}
