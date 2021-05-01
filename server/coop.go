package server

import (
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

//Hash map of connection pools
type CoopMap struct {
	Mutex sync.RWMutex
	Map map[string][]*Pool //Key type String, ValueType Pool Pointer
}


//Pointer Reciever function modifies CoopMap Struct
//Initializes CoopMap
func (c *CoopMap) Init(){
	log.Println(c)
	c.Map = make(map[string][]*Pool)
}

//func (c *CoopMap) getCoopById(id string) []Participant {
//	c.Mutex.RLock()
//	defer c.Mutex.RUnlock()

//	return c.Map[id]
//}

//Create Room generate id and push onto hash map
func (c *CoopMap) createCoop() string {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXUZ1234567890")

	b := make([]rune, 8)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	id := string(b)
	pool := NewPool()
	go pool.Start()
	c.Map[id] = append(c.Map[id], pool)
	return id
}

//insert into Coop and start reading messages
func (c *CoopMap) insertIntoCoop(id string, w *websocket.Conn) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	if c.Map[id] != nil {
		pool := c.Map[id][0] //Get the connection pool for the roomID
		participant := &Participant{w, pool} //New participant for this room

		pool.Register <- participant //Add Participant to the connection Pool
		go participant.Read(pool)
	}
}

//Delete Coop by ID
func (c *CoopMap) deleteCoop(id string ){
	c.Mutex.RLock()
	defer c.Mutex.Unlock()

	delete(c.Map, id)
}

