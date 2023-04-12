package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "err: %v", err)
			os.Exit(1)
		}
		f, err := os.Create("par1_7.txt")
		out := bufio.NewWriter(f)
		_, err = io.Copy(out, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "err: %v", err)
			os.Exit(1)
		}
	}
}
