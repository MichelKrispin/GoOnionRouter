package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type connection struct {
	From string    `json:"from" binding:"required"`
	To   string    `json:"to" binding:"required"`
	Time time.Time `json:"time"`
}

type register struct {
	Address string `json:"address"`
}

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
	router.POST("/received", postReceived)
	router.POST("/send", postSend)

	router.Run("localhost:8888")
}
