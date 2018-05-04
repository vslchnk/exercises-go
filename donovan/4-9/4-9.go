package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	count := make(map[string]int)

	f, err := os.Open("test.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer f.Close()

	input := bufio.NewScanner(f)

	input.Split(bufio.ScanWords)
	for input.Scan() {
		word := input.Text()
		count[word]++
	}

	if err := input.Err(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for s, n := range count {
		fmt.Printf("%s %d\n", s, n)
	}
}
