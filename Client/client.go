package main

import (
	"encoding/json"
	"io"
	"log"
	"net"
	"net/http"
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
}
