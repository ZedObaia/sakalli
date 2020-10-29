package sakalli

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Message struct {
	Type interface{}            `json:"type"`
	Data map[string]interface{} `json:"data"`
	Page interface{}            `json:"page"`
}

func AcceptHandler(server *Server) gin.HandlerFunc {

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

		c.JSON(200, gin.H{"page": data})

	}
	return accept
}

func WsHandler(server *Server) gin.HandlerFunc {
	notify := func(c *gin.Context) {
		// ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		// if err != nil {
		// 	log.Error("Failed to upgrade connection")
		// }
		// defer ws.Close()

		// ch, ok := c.MustGet("channel").(chan Message)
		// if !ok {
		// 	// handle error here...
		// 	log.Error("Failed to get channel in ws")
		// }
		// for {
		// 	for msg := range ch {
		// 		err = ws.WriteJSON(msg)
		// 		if err != nil {
		// 			log.Warn("error write json: " + err.Error())
		// 			close(ch)
		// 			ws.Close()
		// 			break
		// 		} else {
		// 			log.Warn("Message sent ws")
		// 		}
		// 	}
		// }

		// ch, ok := c.MustGet("channel").(chan Message)
		// if !ok {
		// 	// handle error here...
		// 	log.Error("Failed to get channel in ws")
		// }
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Warn("Failed to upgrade ws")
			log.Error(err)
			return
		}
		client := &Client{server: server, conn: conn, send: make(chan Message)}
		client.server.register <- client
		// Allow collection of memory referenced by the caller by doing all work in
		// new goroutines.
		go client.writePump()
		// go client.readPump()

	}

	return gin.HandlerFunc(notify)
}
