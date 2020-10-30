package main

import (
	"flag"
	"fmt"

	"github.com/zedObaia/sakalli/sakalli"

	log "github.com/sirupsen/logrus"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	port := flag.String("port", "8080", "server port")
	flag.Parse()
	p := "8080"
	if *port != "" {
		p = *port
	}
	log.Info("Trying to run on port : ", p)
	gin.SetMode(gin.ReleaseMode)
	server := sakalli.NewServer()
	go server.Run()

	router := gin.Default()
	router.POST("/send/:token", sakalli.SendHandler(server))
	router.POST("/broadcast/", sakalli.BroadcastHandler(server))
	router.GET("/listen/:token", sakalli.WsHandler(server))
	router.Use(static.Serve("/static", static.LocalFile("./static", true)))
	err := router.Run(":" + p)
	if err != nil {
		msg := "Port %v is not free\n"
		log.Error(fmt.Sprintf(msg, p))
	}
}
