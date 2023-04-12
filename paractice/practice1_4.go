package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	counts := make(map[string]int)
	fileName := make(map[string]string)
	if len(files) == 0 {
		countLines(os.Stdin, counts, fileName)
	} else {
		for _, args := range files {
			f, err := os.Open(args)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2 err: %v\n", err)
				continue
			}
			countLines(f, counts, fileName)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%s %s\n", line, fileName[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, fileName map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		fileName[line] = f.Name()
		counts[line]++
	}
}
