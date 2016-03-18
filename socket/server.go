package socket

import (
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

// Server contains connection information
type Server struct {
	path        string
	connections map[int]*Connection
	addCh       chan *Connection
	delCh       chan *Connection
	sendCh      chan *ResetJSON
	doneCh      chan bool
}

// NewServer creates a new socket server.
func NewServer(path string) *Server {
	return &Server{
		path,
		make(map[int]*Connection),
		make(chan *Connection),
		make(chan *Connection),
		make(chan *ResetJSON),
		make(chan bool),
	}
}

func (s *Server) send(msg *ResetJSON) {
	for _, c := range s.connections {
		c.Write(msg)
	}
}

// Listen and serve.
// It serves connections and broadcast request.
func (s *Server) Listen() {

	log.Println("Listening server...")

	// websocket handler
	onConnected := func(ws *websocket.Conn) {
		defer func() {
			ws.Close()
		}()

		connection := NewConnection(ws, s)
		s.addCh <- connection
		connection.Listen()
	}
	http.Handle(s.path, websocket.Handler(onConnected))
	log.Println("Created handler")

	for {
		select {

		// Add new connection
		case c := <-s.addCh:
			log.Println("Added new connection")
			s.connections[c.id] = c
			// s.sendPastMessages(c)

		// del a connection
		case c := <-s.delCh:
			log.Println("Delete connection")
			delete(s.connections, c.id)

		// broadcast reset for all connections
		case msg := <-s.sendCh:
			log.Println("Clock Reset")
			s.send(msg)

		case <-s.doneCh:
			return
		}
	}
}
