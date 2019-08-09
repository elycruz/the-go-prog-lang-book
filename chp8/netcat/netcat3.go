package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

// Listens to 'stdin' and localhost (on given port)

func main() {
	port := flag.Uint64("p", 8000, "Port number")

	flag.Parse()

	conn, err := net.Dial("tcp", "localhost:"+
		strconv.FormatUint(*port, 10))

	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})
	done2 := make(chan struct{})

	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		mustCopy3(os.Stdout, conn)
		done <- struct{}{} // signal main go routine
	}()

	go func() {
		mustCopy3(conn, os.Stdin)
		done2 <- struct{}{}
	}()

	<-done // wait for background go routine to pipe to 'done' channel
	<-done2 // ""
}

func mustCopy3(dest io.Writer, src io.Reader) {
	if _, err := io.Copy(dest, src); err != nil {
		log.Fatal(err)
	}
}
