package sakalli

import (
	"github.com/sirupsen/logrus"
)

// Server : hub for ws connections
type Server struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan Message

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

// NewServer returns new server instance
func NewServer() *Server {
	return &Server{
		broadcast:  make(chan Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

// Run : Run websocket server
func (h *Server) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			logrus.Infoln(len(h.clients), " clients connected")
			for client := range h.clients {
				if message.ID == client.id {

					select {
					case client.send <- message:
					default:
						close(client.send)
						delete(h.clients, client)
					}
				}

			}
		}
	}
}
