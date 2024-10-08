package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"strings"
)

type Message struct {
	sender  int
	message string
}

func handleError(err error) {
	// TODO: all
	// Deal with an error event.
	fmt.Println(err.Error())
}

func acceptConns(ln net.Listener, conns chan net.Conn) {
	// TODO: all
	for {
		conn, err := ln.Accept()
		if err != nil {
			handleError(err)
			continue
		}
		conns <- conn
	}
	// Continuously accept a network connection from the Listener
	// and add it to the channel for handling connections.
}

// func checkerro

func handleClient(client net.Conn, clientid int, msgs chan Message, clients map[int]net.Conn) {
	// TODO: all
	// So long as this connection is alive:
	// Read in new messages as delimited by '\n's
	// Tidy up each message and add it to the messages channel,
	// recording which client it came from.

	for {

		reader := bufio.NewReader(client)
		message, err := reader.ReadString('\n')
		message = strings.TrimSpace(message) //may not need

		if err != nil {
			handleError(err)
			delete(clients, clientid)
			fmt.Printf("connection closed for client: " + string(clientid))
			return
		}

		msg := Message{
			sender:  clientid,
			message: message,
		}

		msgs <- msg
	}

}

func main() {
	// Read in the network port we should listen on, from the commandline argument.
	// Default to port 8030
	portPtr := flag.String("port", ":8030", "port to listen on")
	flag.Parse()

	//TODO Create a Listener for TCP connections on the port given above.
	// ln, _ := net.Listen("tcp",":8030")
	ln, _ := net.Listen("tcp", *portPtr)
	//Create a channel for connections
	conns := make(chan net.Conn)
	//Create a channel for messages
	msgs := make(chan Message)
	//Create a mapping of IDs to connections
	clients := make(map[int]net.Conn)

	//Start accepting connections
	go acceptConns(ln, conns)
	clientsIDN := 0
	for {
		select {
		case conn := <-conns:

			//TODO Deal with a new connection
			// - assign a client ID
			// - add the client to the clients map
			// - start to asynchronously handle messages from this client
			clients[clientsIDN] = conn
			go handleClient(conn, clientsIDN, msgs, clients)

			clientsIDN++

		case msg := <-msgs:
			//TODO Deal with a new message
			// Send the message to all clients that aren't the sender
			for i := 0; i < clientsIDN; i++ {
				if i != msg.sender {
					fmt.Fprintln(clients[i], msg)

				}

			}

		}
	}
}

// panic: runtime error: invalid memory address or nil pointer dereference
// [signal SIGSEGV: segmentation violation code=0x1 addr=0x18 pc=0x4a5a2e]

// goroutine 1 [running]:
// fmt.Fprintln({0x0, 0x0}, {0xc000075f08, 0x1, 0x1})
//         /usr/local/go/src/fmt/print.go:305 +0x4e
// main.main()
