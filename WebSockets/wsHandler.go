package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Fatalln(err)
			return
		}

		log.Println(messageType)
		log.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Fatalln(err)
			return
		}
	}
}

func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Websocket Connected")

	go reader(ws)
}
