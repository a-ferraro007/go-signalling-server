package main

import (
	"log"
	"net/http"
	"pigeon-coop/server"
)



func main() {

	server.AllCoops.Init()

	http.HandleFunc("/create", server.CreateCoopRequestHandler)
	http.HandleFunc("/join", server.JoinCoopRequestHandler)
	http.HandleFunc("/get", server.GetCoopsRequestHandler)

	log.Println("Starting Server on :8000")
	http.ListenAndServe(":8000", nil)
}