package main

import (
	"encoding/json"
	"io"
	"log"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

type node struct {
	Address   string `json:"address"`
	PublicKey string `json:"publickey"`
}

type route struct {
	Nodes []node `json:"nodes"`
}

func getBody(resp *http.Response) string {
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}
	return string(body)
}

const serverRequest = ("GET / HTTP/1.1\r\n" +
	"Host: 127.0.0.1:8080\r\n\r\n")

func main() {
	log.SetPrefix("[CLIENT] ")

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("assets", "./assets")

	router.GET("/", func(c *gin.Context) {
		// An information dashboard
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.GET("/connect", func(ginC *gin.Context) {
		// Connect using the route

		// Ask the directory node for a route
		resp, err := http.Get("http://127.0.0.1:8888/route")
		checkError(err)
		body := getBody(resp)
		var nodeRoute route
		json.Unmarshal([]byte(body), &nodeRoute)
		nodeRoute.Nodes = append(nodeRoute.Nodes, node{"127.0.0.1:8080", ""})

		request, keys := buildRequest(serverRequest, nodeRoute)
		printRoute(nodeRoute)

		// Send that request into the socket
		c, err := net.Dial("tcp", nodeRoute.Nodes[0].Address)
		checkError(err)
		log.Println("Connected to " + nodeRoute.Nodes[0].Address + " socket")

		c.Write([]byte(request))

		// Read the return
		_, response := parseRequest(c, keys)
		log.Println("Received:\n", response)

		// Return a success code
		ginC.IndentedJSON(http.StatusOK, gin.H{
			"success": true,
		})
	})

	router.Run("localhost:9999")
}
