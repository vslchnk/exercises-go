package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	var buf bytes.Buffer
	t := n % 3

	if t > 0 {
		buf.WriteString(s[:t])
	}

	for i := t; i < n; i += 3 {
		if t > 0 {
			buf.WriteString(",")
		}

		buf.WriteString(s[i : i+3])
	}

	return buf.String()
}
