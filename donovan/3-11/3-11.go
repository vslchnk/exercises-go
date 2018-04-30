package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

//!+
// comma inserts commas in a decimal integer string.
func comma(s string) string {
	var buf bytes.Buffer

	if strings.HasPrefix(s, "-") {
		buf.WriteString("-")
		s = strings.TrimPrefix(s, "-")
	}

	if strings.HasPrefix(s, "+") {
		buf.WriteString("+")
		s = strings.TrimPrefix(s, "+")
	}

	var sp string
	if strings.Contains(s, ".") {
		split := strings.Split(s, ".")
		s, sp = split[0], split[1]
	}

	n := len(s)
	if n <= 3 {
		buf.WriteString(s)
		if len(sp) > 0 {
			buf.WriteString(".")
			buf.WriteString(sp)
		}

		return buf.String()
	}

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

	if len(sp) > 0 {
		buf.WriteString(".")
		buf.WriteString(sp)
	}

	return buf.String()
}
