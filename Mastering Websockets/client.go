package main

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

var (
	pongWait     = 10 * time.Second
	pingInterval = (pongWait * 9) / 10
)

type Client struct {
	manager *Manager
	conn    *websocket.Conn
	egress  chan []byte
}

type ClientList map[*Client]bool

func NewClient(manager *Manager, conn *websocket.Conn) *Client {
	return &Client{
		manager: manager,
		conn:    conn,
		// Prevent concurent writes to a websocket connection
		egress: make(chan []byte),
	}
}

func (c *Client) readMessages() {
	defer func() {
		c.manager.removeClient(c)
	}()

	if err := c.conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		log.Println(err)
		return
	}

	c.conn.SetReadLimit(512)
	c.conn.SetPongHandler(c.pongHandler)

	for {
		_, payload, err := c.conn.ReadMessage()

		if err != nil { 
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Error reading message: %v", err)
			}
			break
		}

		for client := range c.manager.clients {
			// Only broadcast to everyone in manager but self
			if client != c {
				log.Println("Broadcast sent")
				client.egress <- payload
			}

			log.Println("Broadcast skipped")
		}

		log.Printf("Message Received: %s", string(payload))
	}
}

func (c *Client) writeMessages() {
	defer func() {
		c.manager.removeClient(c)
	}()

	ticker := time.NewTicker(pingInterval)

	for {
		select {
		case message, ok := <-c.egress:
			if !ok {
				if err := c.conn.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.Printf("COnnection closed: %v", err)
				}
				return
			}

			if err := c.conn.WriteMessage(websocket.TextMessage, message); err != nil {
				log.Printf("Failed to send message: %v", err)
			}
			log.Println("Message sent")

		case <-ticker.C:
			log.Println("Sending Ping")

			// Send ping to client
			if err := c.conn.WriteMessage(websocket.PingMessage, []byte(``)); err != nil {
				log.Printf("Error sending ping: %v", err)
				return
			}
		}
	}
}

func (c *Client) pongHandler(pongMsg string) error {
	log.Println("pong")

	return c.conn.SetReadDeadline(time.Now().Add(pongWait))
}
