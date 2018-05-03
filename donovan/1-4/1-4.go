package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}
	for line, names := range counts {
		if len(strings.Split(names, " "))-1 > 1 {
			fmt.Printf("files: %s have the same string: %s\n", names, line)
		}
	}
}

func countLines(f *os.File, counts map[string]string, fileName string) {
	input := bufio.NewScanner(f)
	sep := " "

	for input.Scan() {
		counts[input.Text()] += fileName + sep
	}
}
