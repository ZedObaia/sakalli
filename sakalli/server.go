package sakalli

import log "github.com/sirupsen/logrus"

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

func NewServer() *Server {
	return &Server{
		broadcast:  make(chan Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Server) Run() {
	for {
		select {
		case client := <-h.register:
			log.Info("registered client ", &client, " for server ", &h)
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				log.Info("unregistered client ", &client, " from server ", &h)
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			log.Info("broadcasting from server ", &h)
			for client := range h.clients {
				log.Info("for client ", &client)
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
