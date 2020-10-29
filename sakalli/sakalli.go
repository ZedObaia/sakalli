package sakalli

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Message : message sent to ws clients
type Message struct {
	Type interface{}            `json:"type"`
	Data map[string]interface{} `json:"data"`
	Page interface{}            `json:"page"`
}

// SendHandler : handels http requests to rely data to websocket clients
func SendHandler(server *Server) gin.HandlerFunc {

	accept := func(c *gin.Context) {

		data := make(map[string]interface{})
		err := c.BindJSON(&data)
		if err != nil {
			c.JSON(400, "Bad JSON")
		}

		server.broadcast <- Message{
			Page: data["page"].(interface{}),
			Type: data["type"].(interface{}),
			Data: data["data"].(map[string]interface{}),
		}

		c.JSON(200, gin.H{"body": data})

	}
	return accept
}

// WsHandler handle websocket connections
func WsHandler(server *Server) gin.HandlerFunc {
	notify := func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Error("Failed to upgrade connections", err)
			return
		}
		client := &Client{server: server, conn: conn, send: make(chan Message)}
		client.server.register <- client
		go client.writePump()
		go client.readPump()

	}

	return gin.HandlerFunc(notify)
}
