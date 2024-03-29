package main

import (
	"bufio"
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
	// size of address | size of content | Encrypted 256 AES key | address | content
	// 	   4 bytes     |    4 bytes	     \   512 byes            |   ...   \  ...
	for {
		c, err := l.Accept()
		checkError(err)

		// If there is a connection parse the HTTP request input
		address, content, key := parseRequest(c, port) // port)
		notifyReceive(nodesAddress, true)

		// After receiving the data pass it on to the next server
		var response string
		// If the content starts with a GET, then this is the last hop
		if strings.HasPrefix(content, "GET") {
			// This ignores completely the original HTTP request right now
			resp, err := http.Get("http://" + address)
			checkError(err)

			b, err := httputil.DumpResponse(resp, true) // Get response as string
			checkError(err)
			response = string(b)
		} else { // Otherwise this is an intermediate hop
			nextConnection, err := net.Dial("tcp", address) // Dial in to next node
			checkError(err)

			nextConnection.Write([]byte(content)) // Pass the content on

			// Parse returning content and put it into the response string
			_, response, _ = parseRequest(nextConnection, "")
		}

		// Wrap the response in and encrypt
		response = buildResponse(response, key)

		// Pass the received response on
		w := bufio.NewWriter(c)
		w.WriteString(response)
		w.Flush()

		notifySend(nodesAddress, true)
		log.Println("Passed response on. Closing connection.\n-----------------------")
		c.Close()
	}
}
