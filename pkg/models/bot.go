package models

import (
	"github.com/shomali11/slacker"
)

type BotCommand struct {
	Usage       string
	Description string
	Example     string
	Handler     func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter)
}

type BotCommands []*BotCommand
