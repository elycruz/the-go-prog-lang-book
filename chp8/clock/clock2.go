package clock

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

func main () {
	port := flag.Uint64("p", 8000, "Port number")

	flag.Parse()

	listener, err := net.Listen("tcp", "localhost:" +
		strconv.FormatUint(*port, 10))
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

func handleConn2(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // E.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
