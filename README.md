# OnionRouter

A simple onion router written in Go as a student project during the course at [TUHH](https://www.tuhh.de/ide/homepage.html).

It uses the [Gin Web Framework](https://github.com/gin-gonic/gin) for the server and clients pages.

This is my first Go project so there are highly likely unidiomatic or inefficient code fragments.

## How to run it

1. Open up 6 terminals so each next step is done in different terminals.
   - One for the (dummy) server.
   - One for the client.
   - One for the directory node.
   - Three for the intermediate nodes.
2. Start the simple server that returns random quotes. (`cd Server` and `go run .`)
   - The server is now up and running on `localhost:8080`, where the result of a simple request can be seen.
3. Start the directory node. (`cd DirectoryNode` and `go run .`)
   - The directory node's informations can be viewed at `localhost:8888` where the registered node and the received and send connections from these nodes will be shown.
4. Start the three intermediate nodes in three different terminals. The port for each node has to be passed as an argument so after `cd Node` do
   - `go run . 8000`
   - `go run . 8001`
   - `go run . 8002`
   - After each command the directory node should output the received node connection with the corresponding public key. Refreshing the directory nodes dashboard (`localhost:8888`) should show the three registered connections.
5. Now the client can be started (`cd Client` and `go run .`). This will start a small webserver at `localhost:9999` which shows a simple website that is able to connect to the directory node, ask it for three nodes, builds up an onion request by creating three AES keys that will be encrypted with the corresponding node's public key and sends this request then to the first node. The website includes a small animation which shows how the sending, encryption and decryption works step by step. In the end the response is shown. Note, that this response is a dummy response on the website but the servers output in the terminal shows the HTTP result that the server sent back.

## Protocol

A custom, very simple protocol is used between the nodes and the client to wrap up a request.

```text
size of address | size of content | Encrypted 256 AES key | address | content
    4 bytes     |    4 bytes	    |     512 bytes         |   ...   |  ...
```
