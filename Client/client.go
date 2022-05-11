package main

import (
	"bufio"
	"log"
	"net"
	"strconv"
	"strings"
)

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

const serverRequest = ("GET / HTTP/1.1\r\n" +
	"Host: 127.0.0.1:8080\r\n\r\n")

func main() {
	log.SetPrefix("[Client]: ")

	// Create the HTTP request
	addresses := []string{"127.0.0.1:8000"}
	request := buildRequest(serverRequest, addresses)

	// Send that request into the socket
	c, err := net.Dial("tcp", addresses[0])
	checkError(err)
	log.Println("Connected to 127.0.0.1:8080 socket")

	c.Write([]byte(request))
	log.Println("Wrote request\n" + request)

	// time.Sleep(time.Second)
	// Read the return
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

	log.Println("Received:\n", response)
}
