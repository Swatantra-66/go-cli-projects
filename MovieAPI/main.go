package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	Title            string  `json:"title"`
	MovieId          string  `json:"movieId"`
	TimelineinMinute int     `json:"timeline"`
	Seasons          int     `json:"seasons"`
	Production       string  `json:"production"`
	Actors           *Actors `json:"actors"`
}

type Actors struct {
	MaleActors   []string `json:"maleActor"`
	FemaleActors []string `json:"femaleActor"`
}

func (m *Movie) isEmpty() bool {
	return m.Title == ""
}

var movie []Movie

// handlers
func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all movies")
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	err := json.NewEncoder(w).Encode(movie)
	if err != nil {
		http.Error(w, "404 not found", http.StatusNotFound)
	}
}

func GetOneMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one movie")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	for _, item := range movie {
		if item.MovieId == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("Sorry! No movie is associated with this id")
}

func AddMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Add one movie")
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	var movies Movie
	err := json.NewDecoder(r.Body).Decode(&movies)
	if err != nil {
		json.NewEncoder(w).Encode("Invalid JSON")
	}

	if movies.isEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
	}

	id := strconv.Itoa(rand.Intn(100))
	movies.MovieId = id
	movie = append(movie, movies)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(movies)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one movie")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	if r.Method != "PUT" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	for index, item := range movie {
		if item.MovieId == id {
			movie = append(movie[:index], movie[index+1:]...)

			var updated Movie
			err := json.NewDecoder(r.Body).Decode(&updated)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode("Invalid json payload")
				return
			}

			if updated.isEmpty() {
				json.NewEncoder(w).Encode("No data inside json")
				return
			}

			updated.MovieId = id
			movie = append(movie, updated)

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(updated)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("No movie is associated with this id")
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one movie")

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	for index, item := range movie {
		if item.MovieId == id {
			movie = append(movie[:index], movie[index+1:]...)
			json.NewEncoder(w).Encode("Movie deleted successfully")
			return
		}
	}
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(movie)
}

func main() {
	movie = append(movie, Movie{
		Title:            "Strangers Things",
		MovieId:          "1",
		TimelineinMinute: 45,
		Seasons:          5,
		Production:       "Duffer Brothers",
		Actors:           &Actors{MaleActors: []string{"David Harbour", "Finn Wolfhand", "Caleb McLaughlin", "Joe Keery", "Jamie Campbell"}, FemaleActors: []string{"Millie Brown", "Sadie Sink", "Maya Hawke", "Natalia Dyer", "Madelyn Cline"}}})
	movie = append(movie, Movie{
		Title:            "Strangers Things",
		MovieId:          "2",
		TimelineinMinute: 45,
		Seasons:          5,
		Production:       "Duffer Brothers",
		Actors:           &Actors{MaleActors: []string{"David Harbour", "Finn Wolfhand", "Caleb McLaughlin", "Joe Keery", "Jamie Campbell"}, FemaleActors: []string{"Millie Brown", "Sadie Sink", "Maya Hawke", "Natalia Dyer", "Madelyn Cline"}}})
	movie = append(movie, Movie{
		Title:            "Strangers Things",
		MovieId:          "3",
		TimelineinMinute: 45,
		Seasons:          5,
		Production:       "Duffer Brothers",
		Actors:           &Actors{MaleActors: []string{"David Harbour", "Finn Wolfhand", "Caleb McLaughlin", "Joe Keery", "Jamie Campbell"}, FemaleActors: []string{"Millie Brown", "Sadie Sink", "Maya Hawke", "Natalia Dyer", "Madelyn Cline"}}})
	movie = append(movie, Movie{
		Title:            "Strangers Things",
		MovieId:          "4",
		TimelineinMinute: 45,
		Seasons:          5,
		Production:       "Duffer Brothers",
		Actors:           &Actors{MaleActors: []string{"David Harbour", "Finn Wolfhand", "Caleb McLaughlin", "Joe Keery", "Jamie Campbell"}, FemaleActors: []string{"Millie Brown", "Sadie Sink", "Maya Hawke", "Natalia Dyer", "Madelyn Cline"}}})

	router := mux.NewRouter()

	//routing
	fileServer := http.FileServer(http.Dir("./static"))
	router.Handle("/", fileServer)

	router.HandleFunc("/movies", GetAllMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", GetOneMovie).Methods("GET")
	router.HandleFunc("/movies", AddMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", UpdateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", DeleteMovie).Methods("DELETE")

	fmt.Printf("Starting Movie API on http://localhost:5000/\n")
	log.Fatal(http.ListenAndServe(":5000", router))
}
