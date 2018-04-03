package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0

	for zO := 1.1; math.Abs(z-zO) > 0.000001; {
		zO = z
		z -= (z*z - x) / (2 * z)
	}

	return z
}

func main() {
	fmt.Println("Self implement: ", Sqrt(2))
	fmt.Println("Math: ", math.Sqrt(2))
}
