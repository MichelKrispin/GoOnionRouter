package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func postReceived(c *gin.Context) {
	var newConnection connection

	// Bind input to the new connection. Note that time is not required.
	if err := c.BindJSON(&newConnection); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !isRegistered(nodes, newConnection) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not registered."})
		return
	}

	connectionsReceived = append(connectionsReceived, newConnection)
	c.IndentedJSON(http.StatusCreated, newConnection)
}
