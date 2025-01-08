package main

import (
	"log"
	"net/http"

	"github.com/Edu58/GoBookStore/src/routes"
	"github.com/gorilla/mux"
)

func main() {
	port := "8000"
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)
	http.Handle("/", router)
	err := http.ListenAndServe("0.0.0.0:"+port, router)
	if err != nil {
		log.Panicln("Could not start server", err)
	}

	log.Printf("Starting server on port %s", port)
}
