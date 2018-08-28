package main

// Sequential clock server

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main () {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening on port:8000")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // E.g., connection aborted continue
			continue
		}
		fmt.Println("Received a connection.")
		handleConn(conn) // Handle one connection at a time
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // E.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
