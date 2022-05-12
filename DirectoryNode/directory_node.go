package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// All nodes
var nodes []string
var connectionsReceived []connection
var connectionsSend []connection

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"registered": nodes,
			"received":   connectionsReceived,
			"send":       connectionsSend,
		})
	})

	router.POST("/register", postRegister)
	router.POST("/receive", postReceived)
	router.POST("/send", postSend)

	router.Run("localhost:8888")
}
