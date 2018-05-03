package main

import "fmt"

func cutDoubles(s []string) []string {
	var i int
	var last string

	for _, v := range s {
		if v != last {
			s[i] = v
			last = v
			i++
		}
	}

	return s[:i]
}

func main() {
	s := []string{"1", "1", "2", "2", "2", "3"}
	fmt.Println(cutDoubles(s))
}
