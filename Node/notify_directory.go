package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
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

type Register struct {
	Address   string `json:"address"`
	PublicKey string `json:"publickey"`
}

func notifyRegister(thisNodesPort string) {
	publicKey := readStringFromFile("keys/public_" + thisNodesPort + ".pem")
	jsonRegister := Register{
		"127.0.0.1:" + thisNodesPort,
		publicKey,
	}
	jsonString, err := json.Marshal(jsonRegister)
	if err != nil {
		panic(nil)
	}

	resp, err := http.Post(
		"http://127.0.0.1:8888/register",
		"application/json",
		strings.NewReader(string(jsonString)),
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

func notifySend(thisNode string, success bool) {
	resp, err := http.Post(
		"http://127.0.0.1:8888/send",
		"application/json",
		strings.NewReader("{\"address\": \""+thisNode+"\", \"success\": "+strconv.FormatBool(success)+"}"),
	)
	if err != nil {
		log.Println("Sending error:", err)
	}

	if resp.StatusCode != 201 { // Code for created
		log.Println("[Failed] Notify (send) to directory node:",
			getBody(resp),
			"\nThis node:", thisNode,
			", Success:", success,
		)
	} else {
		log.Println("Notified (send)    [\033[97;44m" +
			thisNode + "\033[0m -> Success: " + strconv.FormatBool(success) +
			"]")
	}
}

func notifyReceive(thisNode string, success bool) {
	resp, err := http.Post(
		"http://127.0.0.1:8888/receive",
		"application/json",
		strings.NewReader("{\"address\": \""+thisNode+"\", \"success\": "+strconv.FormatBool(success)+"}"),
	)
	if err != nil {
		log.Println("Sending error:", err)
	}

	if resp.StatusCode != 201 { // Code for created
		log.Println("[Failed] Notify (receive) to directory node:",
			getBody(resp),
			"\nThis node:", thisNode,
			", Success:", success,
		)
	} else {
		log.Println("Notified (receive) [\033[97;44m" +
			thisNode + "\033[0m <- Success: " + strconv.FormatBool(success) + "]")
	}
}
