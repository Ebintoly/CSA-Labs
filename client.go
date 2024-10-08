package main

import (
	"bufio"
	"flag"

	// "flag"
	"fmt"
	"net"
	"os"
	"strings"
)

func read(conn net.Conn, aliveCheck chan bool) {
	//TODO In a continuous loop, read a message from the server and display it.
	for {
		reader := bufio.NewReader(conn)
		message, err := reader.ReadString('\n')
		if err != nil {
			aliveCheck <- false
		}
		message = strings.TrimSpace(message)
		fmt.Println(message)
	}
}

func write(conn net.Conn) {
	//TODO Continually get input from the user and send messages to the server.
	stdin := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("Enter Text: ")
		msg, _ := stdin.ReadString('\n')
		fmt.Fprintln(conn, msg)
	}
}

func main() {
	// Get the server address and port from the commandline arguments.
	addrPtr := flag.String("ip", "127.0.0.1:8030", "IP:port string to connect to")
	flag.Parse()
	conn, _ := net.Dial("tcp", *addrPtr)
	//TODO Try to connect to the server
	// conn, _ := net.Dial("tcp", "127.0.0.1:8030")
	//TODO Start asynchronously reading and displaying messages

	aliveCheck := make(chan bool, 1)
	go read(conn, aliveCheck)
	//TODO Start getting and sending user messages.
	go write(conn)

	for {
		select {

		case <-aliveCheck:
			return
		}
	}

}
