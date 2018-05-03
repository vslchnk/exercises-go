package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func cutDoubles(b []byte) []byte {
	var i int
	last := false

	for _, v := range b {
		if unicode.IsSpace(rune(v)) {
			if !last {
				space := make([]byte, 3)
				n := utf8.EncodeRune(space, ' ')
				copy(b[i:i+n], space)
				i += n
			}
			last = true
		} else {
			b[i] = v
			i++
			last = false
		}
	}

	return b[:i]
}

func main() {
	s := []byte("a\t\t\tb\t\ty")
	fmt.Printf("%s\n", cutDoubles(s))
}
