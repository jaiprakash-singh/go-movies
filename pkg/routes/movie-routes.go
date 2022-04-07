package routes

import (
	"github.com/jaiprakash-singh/go-movies/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterMovieRoutes = func(router *mux.Router) {
	router.HandleFunc("/movies", controllers.GetMovies).Methods("GET")
	router.HandleFunc("/movie/{id}", controllers.GetMovie).Methods("GET")
	router.HandleFunc("/movies", controllers.CreateMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", controllers.UpdateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", controllers.DeleteMovie).Methods("DELETE")
}
