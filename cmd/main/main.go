package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jaiprakash-singh/go-movies/pkg/config"
	"github.com/jaiprakash-singh/go-movies/pkg/controllers"
	"github.com/jaiprakash-singh/go-movies/pkg/models"
	"github.com/jaiprakash-singh/go-movies/pkg/routes"

	"github.com/gorilla/mux"
)

func main() {
	models.Movies.Init()

	go controllers.InitialiseBot()

	r := mux.NewRouter()
	routes.RegisterMovieRoutes(r)
	fmt.Println("Starting the server at port", config.APP_PORT)
	log.Fatal(http.ListenAndServe(":"+config.APP_PORT, r))
}
