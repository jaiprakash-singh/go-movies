package controllers

import (
	"encoding/json"

	"github.com/jaiprakash-singh/go-movies/pkg/models"

	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	//gin, fiber = router, socket
)

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(models.Movies)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	//fmt.Fprintf(w, "Request Parameter: %v", params)
	for index, item := range models.Movies {
		if item.ID == params["id"] {
			models.Movies = append(models.Movies[:index], models.Movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(models.Movies)
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for _, item := range models.Movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var movie models.Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000))
	models.Movies = append(models.Movies, &movie)
	json.NewEncoder(w).Encode(models.Movies)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	for index, item := range models.Movies {
		if item.ID == params["id"] {
			models.Movies = append(models.Movies[:index], models.Movies[index+1:]...)
			var movie models.Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			models.Movies = append(models.Movies, &movie)
			json.NewEncoder(w).Encode(movie)
		}
	}
}
