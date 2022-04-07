package controllers

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jaiprakash-singh/go-movies/pkg/config"
	"github.com/jaiprakash-singh/go-movies/pkg/models"
	"github.com/jaiprakash-singh/go-movies/pkg/utils"

	"github.com/shomali11/slacker"
)

var botCommands models.BotCommands

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Printf("Command Events:\nCommand: %q Timestamp: %q Parameters: %q Event: %v\n\n",
			event.Command, event.Timestamp, *event.Parameters, *event.Event)
	}
}

func InitialiseBotCommands() {
	botCommands = append(botCommands, &models.BotCommand{
		Usage:       "Show all",
		Description: "Show All Movies",
		Example:     "Show All Movies",
		Handler:     utils.BotShowAllMovies,
	})

	botCommands = append(botCommands, &models.BotCommand{
		Usage:       "Add movie with isbn <isbn> and title <title> directed by <firstname> <lastname>",
		Description: "Add movie",
		Example:     "Add movie with isbn 0061801240 and title Avatar directed by James Cameron",
		Handler:     utils.BotCreateBotMovie,
	})
}

func InitialiseBot() {
	os.Setenv("SLACK_MOVIE_BOT_TOKEN", config.SLACK_MOVIE_BOT_TOKEN)
	os.Setenv("SLACK_MOVIE_APP_TOKEN", config.SLACK_MOVIE_APP_TOKEN)
	bot := slacker.NewClient(os.Getenv("SLACK_MOVIE_BOT_TOKEN"), os.Getenv("SLACK_MOVIE_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())
	InitialiseBotCommands()

	for _, botCommand := range botCommands {
		bot.Command(botCommand.Usage, &slacker.CommandDefinition{
			Description: botCommand.Description,
			Example:     botCommand.Example,
			Handler:     botCommand.Handler,
		})
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
