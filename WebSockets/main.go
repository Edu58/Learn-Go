package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/ws", wsEndpoint)
	http.ListenAndServe("0.0.0.0:8084", nil)
}
