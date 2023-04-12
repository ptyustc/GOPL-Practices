package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	for _, arg := range os.Args[1:] {
		str := strings.Split(arg, "=")
		conn, err := net.Dial("tcp", str[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Dial err: %v", err)
			return
		}
		defer conn.Close()
		go mustCopy(os.Stdout, conn)
	}
	for true {
	}
}
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		fmt.Fprintf(os.Stderr, "Dial err: %v", err)
		return
	}
	// fmt.Println("1")
}
