package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	ID       string    `json:"id"`
	ISBN     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	id := r.PathValue("id")

	for _, item := range movies {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	id := r.PathValue("id")

	for index, item := range movies {

		if item.ID == id {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = id
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
		}
	}

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	id := r.PathValue("id")
	for index, item := range movies {

		if item.ID == id {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func main() {

	movies = append(movies, Movie{ID: "1", ISBN: "1234567", Title: "The Rock", Director: &Director{FirstName: "Joe", LastName: "Jones"}},
		Movie{ID: "2", ISBN: "098748765", Title: "The Musk", Director: &Director{FirstName: "Eli", LastName: "Job"}})

	router := http.NewServeMux()

	router.HandleFunc("GET /movies", getMovies)
	router.HandleFunc("GET /movies/{id}", getMovie)
	router.HandleFunc("POST /movies", createMovie)
	router.HandleFunc("PUT /movies/{id}", updateMovie)
	router.HandleFunc("DELETE /movies/{id}", deleteMovie)

	fmt.Println("server is running on port 4000")
	err := http.ListenAndServe(":4000", router)
	if err != nil {
		log.Fatal("error: ", err)
	}

}
