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

const serverRequest = ("GET / HTTP/1.1\r\n" +
	"Host: 127.0.0.1:8080\r\n\r\n")

func main() {
	log.SetPrefix("[CLIENT] ")

	// Create the HTTP request
	// Address go in reverse order
	// First hop should no be in here
	addresses := []string{
		"127.0.0.1:8080", // First one is the server
		"127.0.0.1:8000",
		"127.0.0.1:8001",
	}

	request := buildRequest(serverRequest, addresses)

	// Send that request into the socket
	firstAddress := "127.0.0.1:8002"
	c, err := net.Dial("tcp", firstAddress)
	checkError(err)
	log.Println("Connected to " + firstAddress + " socket")

	c.Write([]byte(request))
	log.Println("Wrote request\n" + request)

	// Read the return
	_, response := parseRequest(c)
	/*
		buf := bufio.NewReader(c)
		response := ""
		byteCounter := 0
		foundLength := false
		contentLength := 0
		contentLengthStr := ""
		foundBody := false
		for {
			b, err := buf.ReadByte()
			if err != nil {
				if err.Error() != "EOF" {
					log.Fatalln(err)
				}
				break
			}
			if !foundLength && strings.Contains(response, "Content-Length: ") {
				if b == '\r' {
					foundLength = true
					contentLength, err = strconv.Atoi(contentLengthStr)
					checkError(err)
				} else {
					contentLengthStr += string(b)
				}
			} else if !foundBody && strings.Contains(response, "\r\n\r\n") {
				foundBody = true
			}
			response += string(b)
			if foundBody {
				if byteCounter >= contentLength-1 {
					c.Close()
					break
				}
				byteCounter++
			}
		}
	*/

	log.Println("Received:\n", response)
}
