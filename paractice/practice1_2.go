package main

import (
	"fmt"
	"os"
)

func main() {
	for inx, arg := range os.Args[1:] {
		fmt.Println(inx, " ", arg)
	}
}
