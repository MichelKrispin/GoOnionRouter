package main

import (
	"io"
	"log"
	"net/http"
	"strings"
)

func getBody(resp *http.Response) string {
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}
	return string(body)
}

func notifyRegister(thisNodesPort string) {
	resp, err := http.Post(
		"http://127.0.0.1:8888/register",
		"application/json",
		strings.NewReader("{\"address\": \"127.0.0.1:"+thisNodesPort+"\"}"),
	)
	if err != nil {
		log.Println("Sending error:", err)
		return
	}

	if resp.StatusCode != 201 { // Code for created
		log.Println("[Failed] Register at directory node:", getBody(resp))
	} else {
		log.Println("Registered at directory node")
	}
}

func notifySend(thisNode string, receiver string) {
	resp, err := http.Post(
		"http://127.0.0.1:8888/send",
		"application/json",
		strings.NewReader("{\"from\": \""+thisNode+"\", \"to\": \""+receiver+"\"}"),
	)
	if err != nil {
		log.Println("Sending error:", err)
	}

	if resp.StatusCode != 201 { // Code for created
		log.Println("[Failed] Notify (send) to directory node:",
			getBody(resp),
			"\nThis node:", thisNode,
			", Receiver:", receiver,
		)
	} else {
		log.Println("Notified (send)    [\033[97;44m" +
			thisNode + "\033[0m -> " + receiver +
			"]")
	}
}

func notifyReceive(sender string, thisNode string) {
	resp, err := http.Post(
		"http://127.0.0.1:8888/receive",
		"application/json",
		strings.NewReader("{\"from\": \""+thisNode+"\", \"to\": \""+sender+"\"}"),
	)
	if err != nil {
		log.Println("Sending error:", err)
	}

	if resp.StatusCode != 201 { // Code for created
		log.Println("[Failed] Notify (receive) to directory node:",
			getBody(resp),
			"\nThis node:", thisNode,
			", Sender:", sender,
		)
	} else {
		log.Println("Notified (receive) [" +
			sender + " -> \033[97;44m" + thisNode +
			"\033[0m]")
	}
}