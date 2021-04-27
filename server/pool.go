package server

import (
	"fmt"
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
			case client := <-pool.Register: //pop? a client of the Register channel (chan *Client). set it true in the pool Clients Map.
				pool.Clients[client] = true
				fmt.Println("Size of the connection Pool: ", len(pool.Clients))
				for client, _ := range pool.Clients {
					fmt.Println(client)
					client.Conn.WriteJSON(Message{Type: 1, Body: "New User Joined ... "})
				}
				break
			case client := <-pool.Unregister: //pop? a client of the Register channel (chan *Client). delete from the pool.
				delete(pool.Clients, client)
				fmt.Println("Size of the connection Pool: ", len(pool.Clients))
				for client, _ := range pool.Clients {
					fmt.Println(client)
					client.Conn.WriteJSON(Message{Type: 1, Body: "User disconnected ... "})
				}
				break
			case message := <-pool.Broadcast:
				fmt.Println("Sending Message to All Connected Clients: ", len(pool.Clients)  - 1)
				for client, _ := range pool.Clients {
					fmt.Println(pool.Clients[client])
					if message.Client != client.Conn {
						if err := client.Conn.WriteJSON(message); err != nil {
							fmt.Println(err)
							return
					}
				}
			}
		}
	}
}