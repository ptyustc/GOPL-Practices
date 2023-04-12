package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	var w io.Writer
	w = bytes.Buffer
	fmt.Println(w.(io.ReadWriter))
}
