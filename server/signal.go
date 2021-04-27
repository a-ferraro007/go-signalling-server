package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

var AllCoops CoopMap //Typing As CoopMap
var id string

type resp struct {
	RoomID string `json:"roomID"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,

	//Check origin of the connection and allow any connection for now
	CheckOrigin: func(r *http.Request) bool { return true },
}

func Upgrader( w http.ResponseWriter, r  *http.Request) (*websocket.Conn, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
		log.Fatal("Upgrade error", err)
	}

	return conn, nil
}

//Create Coop and return CoopID
func CreateCoopRequestHandler(w http.ResponseWriter, r *http.Request)  {

	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))

	id = AllCoops.createCoop()
	log.Println(id)

	json.NewEncoder(w).Encode(resp{RoomID: id})
}

//Join Coop
func JoinCoopRequestHandler(w http.ResponseWriter, r *http.Request) {
	//ws, err := AllCoops.Upgrader(w, r)
	roomId, ok := r.URL.Query()["roomID"]

	if !ok {
		log.Println("Missing room Id")
		return
	}

	ws, _  := Upgrader(w, r)

	AllCoops.insertIntoCoop(strings.Join(roomId, " "), false, ws)

}

func GetCoopsRequestHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, AllCoops.Map)
}

