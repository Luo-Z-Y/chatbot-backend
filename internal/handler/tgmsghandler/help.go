package tgmsghandler

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	HelpIntro   = "I'm a bot that can help you with your tasks. Here are the available commands:"
	HelpCmdWord = "help"
	HelpCmdDesc = "Show help message"
)

var helpText = fmt.Sprintf(
	"%s\n\n%s - %s\n%s - %s\n%s - %s\n%s - %s\n%s - %s\n%s - %s",
	HelpIntro,
	HelpCmdWord, HelpCmdDesc,
	AuthCmdWord, AuthCmdDesc,
	StartCmdWord, StartCmdDesc,
	AskCmdWord, AskCmdDesc,
	QueryCmdWord, QueryCmdDesc,
	RequestCmdWord, RequestCmdDesc,
)

func HandleHelpCommand(msg *tgbotapi.Message) (string, error) {
	return helpText, nil
}
