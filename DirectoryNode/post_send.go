package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func postSend(c *gin.Context) {
	var newConnection connection

	// Bind input to the new connection. Note that time is not required.
	if err := c.BindJSON(&newConnection); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if isRegistered(nodes, newConnection) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not registered."})
		return
	}

	newConnection.Time = time.Now()
	connectionsSend = append(connectionsSend, newConnection)
	c.IndentedJSON(http.StatusCreated, newConnection)
}
