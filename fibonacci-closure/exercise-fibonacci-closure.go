package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	iter, first, second := 1, 0, 1

	return func() int {
		switch iter {
		case 1:
		case 2:
			first = 1
		default:
			temp := second
			second = first + second
			first = temp
		}
		iter++
		return first
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
