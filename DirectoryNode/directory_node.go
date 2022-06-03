package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// All nodes
var nodes []node
var connectionsReceived []connection
var connectionsSend []connection

func main() {
	router := gin.Default()

	// Allow CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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

	router.GET("/update", func(c *gin.Context) {
		count := 0
		for {
			if len(connectionsReceived) >= 3 && len(connectionsSend) == 3 {
				break
			}
			if count > 100 {
				c.IndentedJSON(http.StatusBadRequest, gin.H{
					"error": "Timeout",
				})
				return
			}
			count++
			time.Sleep(5 * time.Millisecond)
		}
		c.IndentedJSON(http.StatusOK, gin.H{
			"receive": connectionsReceived,
			"send":    connectionsSend,
		})

	})

	router.POST("/register", postRegister)
	router.POST("/receive", postReceived)
	router.POST("/send", postSend)

	router.Run("localhost:8888")
}
