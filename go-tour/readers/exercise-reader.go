package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (r MyReader) Read(stream []byte) (int, error) {
	if stream == nil {
		return 0, nil
	}

	for i := range stream {
		stream[i] = 'A'
	}

	return len(stream), nil
}

func main() {
	reader.Validate(MyReader{})
}
