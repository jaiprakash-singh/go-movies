package utils

import (
	"github.com/jaiprakash-singh/go-movies/pkg/models"

	"github.com/shomali11/slacker"
)

func BotShowAllMovies(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
	response.Reply(models.GetMovies())
}

func BotCreateBotMovie(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
	isbn := request.Param("isbn")
	title := request.Param("title")
	firstname := request.Param("firstname")
	lastname := request.Param("lastname")

	movie := models.Movie{
		Isbn:  isbn,
		Title: title,
		Director: &models.Director{
			FirstName: firstname,
			LastName:  lastname,
		},
	}
	models.CreateMovie(movie)
	response.Reply(models.GetMovies())
}
