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

type Route struct {
	Nodes []string
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
	var route Route
	json.Unmarshal([]byte(body), &route)
	route.Nodes = append(route.Nodes, "127.0.0.1:8080")

	/*
		// Create the HTTP request
		// First hop should be separate
		firstAddress := "127.0.0.1:8002"
		addresses := []string{
			"127.0.0.1:8000",
			"127.0.0.1:8001",
			"127.0.0.1:8080", // Last one is the server
		}
		request := buildRequest(serverRequest, addresses)
	*/
	firstAddress := route.Nodes[0]
	request := buildRequest(serverRequest, route.Nodes[1:])
	printRoute(firstAddress, route.Nodes[1:])

	// Send that request into the socket
	c, err := net.Dial("tcp", firstAddress)
	checkError(err)
	log.Println("Connected to " + firstAddress + " socket")

	c.Write([]byte(request))

	// Read the return
	_, response := parseRequest(c)
	log.Println("Received:\n", response)
}
