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
	router.POST("/send/:token", sakalli.SendHandler(server))
	router.GET("/listen/:token", sakalli.WsHandler(server))
	router.GET("/", func(c *gin.Context) {
		c.File("./public/home.html")
	})
	router.Use(static.Serve("/static", static.LocalFile("./static", true)))
	router.Run(":8080")
}
