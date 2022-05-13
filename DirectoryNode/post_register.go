package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func postRegister(c *gin.Context) {
	var newNode node
	if err := c.BindJSON(&newNode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if nodesContain(nodes, newNode.Address) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Already registered."})
		return
	}

	nodes = append(nodes, newNode)
	fmt.Println("Registered nodes:", nodes)

	c.IndentedJSON(http.StatusCreated, newNode)
}
