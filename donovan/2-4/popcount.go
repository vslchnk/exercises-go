package main

import "fmt"

func PopCount(x uint64) int {
	c := 0
	for i := uint64(0); i < 64; i++ {
		if x&(1<<i) != 0 {
			c++
		}
	}
	return c
}

func main() {
	for _, x := range []uint64{0, 1, 2, 3, 4, 8, 16, 32, 255} {
		fmt.Printf("\n%d popcount's is %d\n", x, PopCount(x))
	}
}
