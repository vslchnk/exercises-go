package main

import (
	"crypto/sha256"
	"fmt"
)

func PopCount(x uint64) int {
	c := 0
	for i := uint64(0); i < 64; i++ {
		if x&(1<<i) != 0 {
			c++
		}
	}
	return c
}

func Count(a, b [32]byte) int {
	if len(a) != len(b) {
		return 0
	}

	c := 0
	for i := range a {
		c += PopCount(uint64(a[i] ^ b[i]))
	}

	return c
}

func main() {
	a := sha256.Sum256([]byte("Y"))
	b := sha256.Sum256([]byte("X"))

	fmt.Println(Count(a, b))
}
