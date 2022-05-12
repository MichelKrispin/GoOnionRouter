package main

import (
	"log"
	"net"
)

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func printRoute(firstHop string, hops []string) {
	result := "The route goes\n\t" + firstHop + " -> "
	l := len(hops) - 1
	for i, _ := range hops {
		result += hops[l-i]
		if i != l {
			result += " -> "
		}
	}
	log.Println(result)
}

const serverRequest = ("GET / HTTP/1.1\r\n" +
	"Host: 127.0.0.1:8080\r\n\r\n")

func main() {
	log.SetPrefix("[CLIENT] ")

	// Create the HTTP request
	// Address go in reverse order
	// First hop should no be in here
	firstAddress := "127.0.0.1:8002"
	addresses := []string{
		"127.0.0.1:8080", // First one is the server
		"127.0.0.1:8001",
		"127.0.0.1:8000",
	}
	request := buildRequest(serverRequest, addresses)

	printRoute(firstAddress, addresses)

	// Send that request into the socket
	c, err := net.Dial("tcp", firstAddress)
	checkError(err)
	log.Println("Connected to " + firstAddress + " socket")

	c.Write([]byte(request))

	// Read the return
	_, response := parseRequest(c)
	log.Println("Received:\n", response)
}
