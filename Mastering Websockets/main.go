package main

import (
	"log"
	"net/http"
)

func main() {
	setUpAPI()

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func setUpAPI() {
	manager := NewManager()
	http.Handle("/", http.FileServer(http.Dir("./frontend")))
	http.HandleFunc("/ws", manager.ServeWS)
}
