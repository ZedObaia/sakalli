package main

import (
	"github/zedObaia/sakalli/sakalli"

	"github.com/gin-gonic/gin"
)

func main() {
	server := sakalli.NewServer()
	go server.Run()

	router := gin.Default()
	router.POST("/send/:token", sakalli.SendHandler(server))
	router.GET("/ws/:token", sakalli.WsHandler(server))
	router.GET("/", func(c *gin.Context) {
		c.File("./public/home.html")
	})

	router.Run(":8080")
}
