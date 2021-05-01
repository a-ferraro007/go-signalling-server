package server

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Participant struct {
	Conn *websocket.Conn
	Coop *Pool
	//other user information?
}

type Message struct {
	Message map[string]interface{}
	Client *Participant
}

func (c *Participant) Read(pool *Pool) {
	defer func() {
		pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		msg := Message{Client: c}
		err := c.Conn.ReadJSON(&msg.Message) //ReadJSON

		log.Println("LOG MESSAGE:")
		log.Println(msg.Message)

		if err != nil {
			log.Println(err)
			return
		}

		pool.Broadcast <- msg
		fmt.Printf("Message Received: %+v\n", msg)
	}

}