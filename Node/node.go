package main

import (
	"bufio"
	"encoding/binary"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
)

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	// Getting the port from the args or using default
	port := "8000"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}
	log.SetPrefix("[NODE:" + port + "] ")

	// Listening on the given port
	l, err := net.Listen("tcp", ":"+port)
	checkError(err)
	nodesAddress := "127.0.0.1:" + port
	notifyRegister(port)

	// PROTOCOL
	// size of address | size of content | address | content
	// 	   4 bytes     |    4 bytes	     \   ...   \  ...
	for {
		c, err := l.Accept()
		checkError(err)

		// If there is a connection parse the HTTP request input
		address, content := parseRequest(c)
		notifyReceive(c.RemoteAddr().String(), nodesAddress)

		// After receiving the data pass it on to the next server
		var response string
		// If the content starts with a GET, then this is the last hop
		if strings.HasPrefix(content, "GET") {
			// This ignores completely the original HTTP request right now
			resp, err := http.Get("http://" + address)
			checkError(err)
			notifySend(nodesAddress, address)

			b, err := httputil.DumpResponse(resp, true) // Get response as string
			checkError(err)
			response = string(b)
			notifyReceive(address, nodesAddress)
		} else { // Otherwise this is an intermediate hop
			nextConnection, err := net.Dial("tcp", address) // Dial in to next node
			checkError(err)

			nextConnection.Write([]byte(content)) // Pass the content on
			notifySend(nextConnection.LocalAddr().String(), address)

			// Parse returning content and put it into the response string
			_, response = parseRequest(nextConnection)
			notifyReceive(address, nextConnection.LocalAddr().String())
		}

		// Wrap the response up again if it isn't already wrapped up
		if strings.HasPrefix(response, "HTTP/") {
			dummyAddress := "none" // Will be ignored anyway
			addressBytes := make([]byte, 4)
			binary.BigEndian.PutUint32(addressBytes[0:], uint32(len(dummyAddress)))
			contentBytes := make([]byte, 4)
			binary.BigEndian.PutUint32(contentBytes[0:], uint32(len(response)))

			response = string(addressBytes) + string(contentBytes) + dummyAddress + response
		}

		// Pass the received response on
		w := bufio.NewWriter(c)
		w.WriteString(response)
		w.Flush()

		notifySend(nodesAddress, c.RemoteAddr().String())
		log.Println("Passed response back and closing connection.\n-----------------------")
		c.Close()
	}
}
