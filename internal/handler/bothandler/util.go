package bothandler

import (
	"backend/internal/telegram"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// / This function is a duplicate of the one in tgmessagehandler/util.go because
// / the original function requires a Message and Bot parameter, which are not readily available by other functions
func SendTelegramMessage(
	telegramChatID int64, // Note that this takes in the actual telegram's chat ID, not the DB's chat ID
	messageID *int64,
	content string,
) (*tgbotapi.Message, error) {
	bot := telegram.GetBotAPI()

	response := tgbotapi.NewMessage(telegramChatID, content)
	if messageID != nil {
		response.ReplyToMessageID = int(*messageID)
	}

	msg, err := bot.Send(response)
	if err != nil {
		return nil, err
	}

	return &msg, nil
}
