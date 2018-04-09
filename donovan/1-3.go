package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	t := time.Now()

	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	fmt.Println("1-st: ", time.Since(t).Nanoseconds())

	t = time.Now()

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	fmt.Println("2-nd: ", time.Since(t).Nanoseconds())

	t = time.Now()

	fmt.Println(strings.Join(os.Args[:], " "))

	fmt.Println("3-rd: ", time.Since(t).Nanoseconds())
}
