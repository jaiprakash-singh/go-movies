package models

import (
	"math/rand"
	"strconv"
)

type MoviesType []*Movie

func (m MoviesType) ToString() string {
	res := "List of movies:\n"
	for _, movie := range m {
		res += movie.ToString()
	}
	return res
}

func GetMovies() string {
	return Movies.ToString()
}

var Movies MoviesType

func (m MoviesType) Init() {
	Movies = append(Movies, &Movie{
		ID:    "1",
		Isbn:  "56342",
		Title: "Movie one",
		Director: &Director{
			FirstName: "John",
			LastName:  "Rambo",
		}})
	Movies = append(Movies, &Movie{
		ID:    "2",
		Isbn:  "57654",
		Title: "Movie two",
		Director: &Director{
			FirstName: "Steve",
			LastName:  "Rogers",
		}})
}

func CreateMovie(movie Movie) {
	movie.ID = strconv.Itoa(rand.Intn(10000))
	Movies = append(Movies, &movie)
}

func GetAllMovies(movie Movie) {
	movie.ID = strconv.Itoa(rand.Intn(10000))
	Movies = append(Movies, &movie)
}
