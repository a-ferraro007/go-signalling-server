package server

import (
	"fmt"
	"log"
)

type Pool struct {
	Register chan *Participant
	Unregister chan *Participant
	Clients map[*Participant]bool
	Broadcast chan Message
}

func NewPool() *Pool {
	return &Pool {
		Register: make(chan *Participant),
		Unregister: make(chan *Participant),
		Clients: make(map[*Participant]bool),
		Broadcast: make(chan Message),
	}
}

func (pool *Pool) Start() {
	for {
		select {
			case client := <-pool.Register:
				pool.Clients[client] = true
				fmt.Println("Size of the connection Pool: ", len(pool.Clients))
				for client, _ := range pool.Clients {
					log.Println(client)
				}
				break
			case client := <-pool.Unregister:
				delete(pool.Clients, client)
				fmt.Println("Size of the connection Pool: ", len(pool.Clients))
				for client, _ := range pool.Clients {
					fmt.Println(client)
				}
				break
			case message := <-pool.Broadcast:
				fmt.Println("Sending Message to All Connected Clients: ", len(pool.Clients)  - 1)
				fmt.Println(message.Message)
				for client, _ := range pool.Clients {
					if message.Client.Conn != client.Conn {
						if err := client.Conn.WriteJSON(message.Message); err != nil {
							fmt.Println(err)
							return
					}
				}
			}
		}
	}
}