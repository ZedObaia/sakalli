package sakalli

import (
	"github.com/sirupsen/logrus"
)

// Server : hub for ws connections
type Server struct {
	// Registered clients.
	clients map[string]map[*Client]bool

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
		//  map of ids that points to map of clients for easy removal and easy access
		clients: make(map[string]map[*Client]bool),
	}
}

// Run : Run websocket server
func (h *Server) Run() {
	for {
		select {
		case client := <-h.register:
			if _, ok := h.clients[client.id]; ok {
				h.clients[client.id][client] = true
			} else {
				h.clients[client.id] = make(map[*Client]bool)
				h.clients[client.id][client] = true
			}
			logrus.Info(h.clients)

		case client := <-h.unregister:
			if _, ok := h.clients[client.id][client]; ok {
				client.unRegister()
			}
		case message := <-h.broadcast:
			for _, id := range message.IDs {
				if _, ok := h.clients[id]; ok {
					logrus.Infoln("Broadcasting to ", len(h.clients[id]), "connected clients ID: ", id)
					for client := range h.clients[id] {

						select {
						case client.send <- message:
						default:
							client.unRegister()
						}

					}
				}
			}

		}
	}
}
