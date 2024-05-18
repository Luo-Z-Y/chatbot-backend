package tgmsghandler

import (
	"backend/internal/handler/tgmsghandler/tgcmd"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleCommand(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) error {
	cmd := msg.Command()

	switch cmd {
	case "help":
		tgcmd.HandleHelpCommand(msg)
	case "new":
		tgcmd.HandleNewCommand(msg)
	case "query":
		tgcmd.HandleQueryCommand(msg)
	case "request":
		tgcmd.HandleRequestCommand(msg)
	}

	if _, err := bot.Send(tgbotapi.NewMessage(msg.Chat.ID, msg.Text)); err != nil {
		return err
	}

	return nil
}
