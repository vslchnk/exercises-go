package main

import (
	"fmt"
)

func main() {
	a := "check"
	b := "kcehc"
	fmt.Println(isAnagram(a, b))

	a = "check"
	b = "kcech"
	fmt.Println(isAnagram(a, b))
}

func isAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	if a != Reverse(b) {
		return false
	}

	return true
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
