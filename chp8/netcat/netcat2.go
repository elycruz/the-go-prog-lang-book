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

func main ()  {
	port := flag.Uint64("p", 8000, "Port number")

	flag.Parse()

	conn, err := net.Dial("tcp", "localhost:" +
		strconv.FormatUint(*port, 10))

	if err != nil { log.Fatal(err) }
	defer conn.Close()
	go mustCopy2(os.Stdout, conn)
	mustCopy2(conn, os.Stdin)
}

func mustCopy2 (dest io.Writer, src io.Reader)  {
	if _, err := io.Copy(dest, src); err != nil {
		log.Fatal(err)
	}
}
