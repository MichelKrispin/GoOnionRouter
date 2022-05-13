package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// All nodes
var nodes []node
var connectionsReceived []connection
var connectionsSend []connection

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"registered": nodes,
			"received":   connectionsReceived,
			"send":       connectionsSend,
		})
	})

	router.GET("/route", func(c *gin.Context) {
		routes, err := getRoute(nodes)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{
				"nodes": routes,
			})
		}
	})

	router.POST("/register", postRegister)
	router.POST("/receive", postReceived)
	router.POST("/send", postSend)

	router.Run("localhost:8888")
}
