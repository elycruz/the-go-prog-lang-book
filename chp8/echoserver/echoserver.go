package echoserver

import (
"flag"
"fmt"
"io"
"log"
"net"
"strconv"
"time"
)

/**
 * Prints time with host name and port
 */

func main () {
	port := flag.Uint64("p", 8000, "Port number")

	flag.Parse()

	connUri := "localhost:" + strconv.FormatUint(*port, 10)
	listener, err := net.Listen("tcp", connUri)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Listening on port:%d", *port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // E.g., connection aborted continue
			continue
		}
		fmt.Println("Received a connection.")
		go handleConn(connUri, conn) // Handle one connection at a time
	}
}

func handleConn(cUri string, c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, cUri +" -> "+ time.Now().Format("15:04:05\n"))
		if err != nil {
			return // E.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
