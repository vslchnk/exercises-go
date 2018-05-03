package main

import (
	"fmt"
	"unicode/utf8"
)

func reverse(b []byte) {
	var f, l rune
	var fn, ln int

	for i, j := 0, len(b); i < j-1; i, j = i+ln, j-fn {
		f, fn = utf8.DecodeRune(b[i:])
		l, ln = utf8.DecodeLastRune(b[:j])

		if ln > fn {
			copy(b[i+ln:], b[i+fn:j-ln])
		}

		copy(b[i:], []byte(string(l)))
		copy(b[j-fn:], []byte(string(f)))
	}
}

func main() {
	a := []byte("a\t\t\tb\t\ty")
	reverse(a)
	fmt.Printf("%s\n", a)
	b := []byte("ab世界y")
	reverse(b)
	fmt.Printf("%s\n", b)
}
