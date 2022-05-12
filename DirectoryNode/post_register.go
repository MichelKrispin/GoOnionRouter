package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func postRegister(c *gin.Context) {
	var newAddress register
	if err := c.BindJSON(&newAddress); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if nodesContain(nodes, newAddress.Address) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Already registered."})
		return
	}

	nodes = append(nodes, newAddress.Address)
	fmt.Println("Registered nodes:", nodes)

	c.IndentedJSON(http.StatusCreated, newAddress)
}
