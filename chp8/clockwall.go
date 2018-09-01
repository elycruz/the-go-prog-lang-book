package main

import (
	"io"
	"log"
	"net"
	"os"
	"strings"
)

/**
 * Calls all hosts passed in and prints their output to 'stdout'
 * @todo requires channels for it to work properly
 */

func main ()  {
	if len(os.Args) <= 1 {
		log.Print("Expected a list of locations.  Got none.")
		return
	}

	log.Print("Running for: ", strings.Join(os.Args[1:], ", "))

	for _, location := range os.Args[1:] {
		if !isValidLocation(location) {
			log.Printf("Location not valid.  Location: \"%s\"", location)
			continue
		}
		go makeNetCatConn(location)
	}
}

func isValidLocation (location string) bool {
	return len(location) > 0
}

func makeNetCatConn (location string) {
	conn, err := net.Dial("tcp", location)
	if err != nil { log.Fatal(err) }
	log.Printf("Dialing \"%s\"...", location)
	defer conn.Close()
	pipeReaderToWriter(conn, os.Stdout)
}

func pipeReaderToWriter(incoming io.Reader, outgoing io.Writer)  {
	if _, err := io.Copy(outgoing, incoming); err != nil {
		log.Fatal(err)
	}
}
