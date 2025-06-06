package main

import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

func getMovies(w http.ResponseWriter, r *http.Request) {

}

func getMovie(w http.ResponseWriter, r *http.Request) {

}

func createMovie(w http.ResponseWriter, r *http.Request) {

}

func updateMovie(w http.ResponseWriter, r *http.Request) {

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {

}

type Movie struct {
	ID      string   `json:"id"`
	Isbn    string   `json:"isbn"`
	Title   string   `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()
	port := 8000
	movies = append(movies, Movie{ID: "1", Isbn: "438743", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "438744", Title: "Movie Two", Director: &Director{Firstname: "Jane", Lastname: "Doe"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))

}
