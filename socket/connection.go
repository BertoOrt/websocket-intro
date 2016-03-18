package socket

import (
	"log"

	"golang.org/x/net/websocket"
)

// ResetJSON sends true to restart the clock
type ResetJSON struct {
	Reset bool `json:"reset"`
}

const channelBuffSize = 100

var id int

// Connection contains client information.
type Connection struct {
	id     int
	ws     *websocket.Conn
	server *Server
	ch     chan *ResetJSON
	doneCh chan bool
}

// NewConnection creates new connection.
func NewConnection(ws *websocket.Conn, server *Server) *Connection {
	id++
	return &Connection{
		id,
		ws,
		server,
		make(chan *ResetJSON, channelBuffSize),
		make(chan bool),
	}
}

// Conn connects socket
func (c *Connection) Conn() *websocket.Conn {
	return c.ws
}

func (c *Connection) Write(msg *ResetJSON) {
	select {
	case c.ch <- msg:
	default:
		c.server.delCh <- c
	}
}

// Listen Write and Read request via channel
func (c *Connection) Listen() {
	go c.listenWrite()
	c.listenRead()
}

// Listen write request via chanel
func (c *Connection) listenWrite() {
	for {
		select {

		// send reset to the connection
		case msg := <-c.ch:
			log.Println("Reset")
			websocket.JSON.Send(c.ws, msg)

		// receive done request
		case <-c.doneCh:
			c.server.delCh <- c
			c.doneCh <- true // for listenRead method
			return
		}
	}
}

// Listen read request via chanel
func (c *Connection) listenRead() {
	for {
		select {

		// receive done request
		case <-c.doneCh:
			c.server.delCh <- c
			c.doneCh <- true // for listenWrite method
			return

		// read data from websocket connection
		default:
			var msg ResetJSON
			websocket.JSON.Receive(c.ws, &msg)
			c.server.send(&msg)
		}
	}
}
