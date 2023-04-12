package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

var TZ, port string

func main() {
	flag.StringVar(&TZ, "TZ", "", "时区")
	flag.StringVar(&port, "port", "", "端口")
	flag.Parse()
	fmt.Println(TZ, " ", port)
	listener, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Listen err: %v", err)
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Accept err: %v", err)
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	var cstZone *time.Location
	switch TZ {
	case "US/Western":
		cstZone = time.FixedZone("CST", (-4)*3600)
	case "Asian/Tokyo":
		cstZone = time.FixedZone("CST", 9*3600)
	case "Asian/Beijing":
		cstZone = time.FixedZone("CST", 8*3600)
	}
	for {
		_, err := io.WriteString(conn, time.Now().In(cstZone).Format("15:04:05\n"))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Conn err: %v", err)
			return
		}
		// time.Sleep(1000 * time.Millisecond)
	}
}
