package main

import (
	"log"
	"net/http"

	"github.com/bertoort/websockets-intro/socket"
)

func main() {
	// Serve socket
	server := socket.NewServer("/reset")
	go server.Listen()

	// Serving static files
	http.Handle("/", http.FileServer(http.Dir("client")))
	log.Fatal(http.ListenAndServe(":3000", nil))
}
