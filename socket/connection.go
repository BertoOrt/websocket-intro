package socket

import (
	"io"
	"log"
	"time"

	"golang.org/x/net/websocket"
)

// Message sends true to restart the clock
type Message struct {
	GameOver bool      `json:"gameOver"`
	Time     time.Time `json:"time"`
}

const channelBuffSize = 100

var (
	id           int
	resetCounter int
)

// Connection contains client information.
type Connection struct {
	id     int
	ws     *websocket.Conn
	server *Server
	ch     chan *Message
	doneCh chan bool
}

// NewConnection creates new connection.
func NewConnection(ws *websocket.Conn, server *Server) *Connection {
	id++
	return &Connection{
		id,
		ws,
		server,
		make(chan *Message, channelBuffSize),
		make(chan bool),
	}
}

// Write sends the message to the connection channel
func (c *Connection) Write(msg *Message) {
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
			var msg Message
			err := websocket.JSON.Receive(c.ws, &msg)
			if err == io.EOF {
				c.doneCh <- true
			} else if msg.GameOver && !c.server.gameOver {
				log.Println("Game Over")
				c.server.gameOver = true
				c.doneCh <- true
			} else {
				resetCounter++
				log.Println("Reset count: ", resetCounter)
				time := time.Now()
				c.server.time = time
				msg.Time = c.server.time
				c.server.send(&msg)
			}
		}
	}
}
