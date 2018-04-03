package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)

	for _, value := range strings.Fields(s) {
		m[value] += 1
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
