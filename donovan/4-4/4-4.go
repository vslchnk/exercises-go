package main

import "fmt"

func rotateLeft(s []int, count int) {
	t := make([]int, len(s))
	copy(t, s)

	for i, j := 0, 0-count; i < len(s); i, j = i+1, i-count+1 {
		if j < 0 {
			j = len(s) - count + i
		}
		s[j] = t[i]
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a)
	rotateLeft(a, 2)
	fmt.Println(a)
}
