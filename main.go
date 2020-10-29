package main

import (
	"github/zedObaia/sakalli/sakalli"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	server := sakalli.NewServer()
	go server.Run()

	router := gin.Default()
	router.POST("/send", sakalli.AcceptHandler(server))
	router.GET("/ws", sakalli.WsHandler(server))
	// static files
	router.Use(static.Serve("/", static.LocalFile("./public", true)))

	router.GET("/", func(c *gin.Context) {
		c.File("./public/home.html")
	})

	router.Run(":8080")
}
