package netcat

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

func main ()  {
	port := flag.Uint64("p", 8000, "Port number")

	flag.Parse()

	conn, err := net.Dial("tcp", "localhost:" +
		strconv.FormatUint(*port, 10))

	if err != nil { log.Fatal(err) }

	defer conn.Close()

	mustCopy(os.Stdout, conn)
}

func mustCopy (dest io.Writer, src io.Reader)  {
	if _, err := io.Copy(dest, src); err != nil {
		log.Fatal(err)
	}
}
