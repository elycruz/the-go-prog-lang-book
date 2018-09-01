package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
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
		go handleConn2(conn) // Handle one connection at a time
	}
}

func echo2(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func pipeReaderToWriter2(incoming io.Reader, outgoing io.Writer)  {
	if _, err := io.Copy(outgoing, incoming); err != nil {
		log.Fatal(err)
	}
}

func handleConn2(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo2(c, input.Text(), 1*time.Second)
	}
	// Ignoring potential errors
	c.Close()
}
