package main

import (
	"flag"
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
	defer conn.Close()
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}
