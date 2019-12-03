package main

import (
	"log"
	"net"
	"os"
	"strings"

	"./cmd"
)

// SockAddr Address of Unix Socket
const SockAddr = "/tmp/locald.sock"

// newCommand handles new requests
func newCommand(c net.Conn) {
	buf := make([]byte, 1024)
	for {
		n, err := c.Read(buf[:])
		if err != nil {
			return
		}
		command := string(buf[0:n])
		command = strings.TrimSuffix(command, "\n")
		log.Printf("Command is %s", command)
		cmd.Execute(c, command)
	}
}

func main() {
	// Try to remove the socket on launch
	if err := os.RemoveAll(SockAddr); err != nil {
		log.Fatal(err)
	}

	// Listen to the socket
	l, err := net.Listen("unix", SockAddr)
	if err != nil {
		log.Fatal("Cannot to socket:", err)
	}
	defer l.Close()

	// Handle requests on the socket
	for {
		// Get a new connection
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("Error accepting connection", err)
		}

		// Send connection to Cobra in goroutine
		defer conn.Close()
		go newCommand(conn)
	}
}
