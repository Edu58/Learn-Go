package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Director struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

var movies []Movie

func main() {
	router := mux.NewRouter()

	movie1 := Movie{ID: "1", Isbn: "wereq4234234", Title: "The Book", Director: &Director{"Mans", "Runing"}}
	movies = append(movies, movie1)

	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at port 8000")
	http.ListenAndServe(":8000", router)
}
