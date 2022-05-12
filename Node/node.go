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
	log.SetPrefix("[Node]: ")

	// Getting the port from the args or using default
	port := "8000"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	// Listening on the given port
	l, err := net.Listen("tcp", ":"+port)
	checkError(err)
	log.Println("Listenting on :" + port)

	// PROTOCOL
	// size of address | size of content | address | content
	// 	   4 bytes     |    4 bytes	     \   ...   \  ...
	for {
		c, err := l.Accept()
		checkError(err)
		defer c.Close()

		// If there is a connection parse the HTTP request input
		address, content := parseRequest(c)
		log.Println("Received:\n", address, "\nand\n", content, "\n----------------")

		// After receiving the data pass it on to the next server
		var response string
		// If the content starts with a GET, then this is the last hop
		if strings.HasPrefix(content, "GET") {
			// Find the host
			host := ""
			for _, line := range strings.Split(content, "\n") {
				if strings.HasPrefix(line, "Host: ") {
					host = strings.Split(line, " ")[1]
					host = host[:len(host)-1]
					host = "http://" + host
					break
				}
			}
			req, err := http.NewRequest("GET", host, nil)
			checkError(err)

			resp, err := http.DefaultClient.Do(req)
			b, err := httputil.DumpResponse(resp, true)
			checkError(err)
			response = string(b)
		} else { // Otherwise this is an intermediate hop
			nextConnection, err := net.Dial("tcp", address)
			checkError(err)
			log.Println("Connected to next server: ", address)

			nextConnection.Write([]byte(content))
			log.Println("Wrote request\n" + content)

			// Parse returning content again and put it into the response string
			_, response = parseRequest(c)
			log.Println("Received content next hop:\n", content, "\n----------------")
		}

		// Now wrap the response up again and send it back
		dummyAddress := "none"
		addressBytes := make([]byte, 4)
		binary.BigEndian.PutUint32(addressBytes[0:], uint32(len(dummyAddress)))

		contentBytes := make([]byte, 4)
		binary.BigEndian.PutUint32(contentBytes[0:], uint32(len(response)))

		response = string(addressBytes) + string(contentBytes) + dummyAddress + response

		// Pass the received response on
		w := bufio.NewWriter(c)
		/*
					w.WriteString(`HTTP/1.1 200 OK
			Content-Type: application/json; charset=utf-8
			Date: Wed, 11 May 2022 10:26:10 GMT
			Content-Length: 29

			{
				\"quote\": \"Some quote\"
			}`)
		*/
		w.WriteString(response)
		w.Flush()
		log.Println("Wrote response:\n", response, "\nClosing connection.\n----------------")
		c.Close()
	}
}