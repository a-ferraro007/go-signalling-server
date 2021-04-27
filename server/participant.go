package server

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Participant struct {
	Host bool
	Conn *websocket.Conn
	Coop *Pool
	//other user information?
}

type Message struct {
	Type int `"json:type"`
	Body string `"json:type"`
	Client *websocket.Conn
}

func (c *Participant) Read(pool *Pool) {
	defer func() {
		pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()

		if err != nil {
			log.Println(err)
			return
		}
		message := Message{Type: messageType, Body: string(p), Client: c.Conn}
		pool.Broadcast <- message
		fmt.Printf("Message Received: %+v\n", message)
	}

}