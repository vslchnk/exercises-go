package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (ro rot13Reader) Read(stream []byte) (int, error) {
	n, err := ro.r.Read(stream)

	if err == nil {
		for i := range stream {
			if (stream[i] >= 'A' && stream[i] <= 'M') || (stream[i] >= 'a' && stream[i] <= 'm') {
				stream[i] += 13
			} else if (stream[i] >= 'N' && stream[i] <= 'Z') || (stream[i] >= 'n' && stream[i] <= 'z') {
				stream[i] -= 13
			}
		}
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
