package main

import (
	"flag"
	"fmt"
	"github/zedObaia/sakalli/sakalli"

	log "github.com/sirupsen/logrus"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	port := flag.String("port", "8080", "server port")
	flag.Parse()
	log.Info("Trying to run on port : ", *port)
	gin.SetMode(gin.ReleaseMode)
	server := sakalli.NewServer()
	go server.Run()

	router := gin.Default()
	router.POST("/send/:token", sakalli.SendHandler(server))
	router.POST("/broadcast/", sakalli.BroadcastHandler(server))

	router.GET("/listen/:token", sakalli.WsHandler(server))
	router.GET("/", func(c *gin.Context) {
		c.File("./public/home.html")
	})
	router.Use(static.Serve("/static", static.LocalFile("./static", true)))
	err := router.Run(":" + *port)
	if err != nil {
		msg := "Port %v is not free\n"
		log.Error(fmt.Sprintf(msg, *port))
	}
}
