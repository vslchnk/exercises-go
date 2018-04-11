package main

import "fmt"

// pc[i] is the population count of i.
var pc [256]byte // 256 = 8^2 possible combinations of 0 and 1 for 8 slots

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	var sum byte
	for i := uint64(0); i < 8; i++ {
		sum += pc[byte(x>>(i*8))]
	}
	return int(sum)
}

func main() {
	for _, x := range []uint64{0, 1, 2, 3, 4, 8, 16, 32, 255} {
		fmt.Printf("\n%d popcount's is %d\n", x, PopCount(x))
	}
}
