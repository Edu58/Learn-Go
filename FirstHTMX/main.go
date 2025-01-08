package main

import (
	"log"
	"net/http"

	"github.com/Edu58/Learn-Go/FirstHTMX/controllers"
)

func main() {
	mux := http.DefaultServeMux
	files := http.FileServer(http.Dir("./public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	
	
	mux.HandleFunc("/", controllers.Index)
	mux.HandleFunc("/movie", controllers.Create)

	log.Fatalln(http.ListenAndServe(":8000", mux))
}
