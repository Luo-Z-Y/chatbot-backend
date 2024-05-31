package bothandler

import (
	"backend/internal/telegram"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SendTelegramMessage(
	telegramChatID int64, // Note that this takes in the actual telegram's chat ID, not the DB's chat ID
	content string,
) (*tgbotapi.Message, error) {
	bot := telegram.GetBotAPI()

	response := tgbotapi.NewMessage(telegramChatID, content)

	msg, err := bot.Send(response)
	if err != nil {
		return nil, err
	}

	return &msg, nil
}
