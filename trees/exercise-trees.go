package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	var walker func(t *tree.Tree)

	walker = func(t *tree.Tree) {
		if t == nil {
			return
		}
		walker(t.Left)
		ch <- t.Value
		walker(t.Right)
	}

	walker(t)

	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	var v1, v2 int
	ok1, ok2 := true, false

	for ok1 {
		v1, ok1 = <-ch1
		v2, ok2 = <-ch2

		if v1 != v2 || ok1 != ok2 {
			return false
		}
	}

	return true
}

func main() {
	ch := make(chan int)

	go Walk(tree.New(1), ch)

	for v := range ch {
		fmt.Println(v)
	}

	ch = make(chan int)

	go Walk(tree.New(2), ch)

	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println("same: ", Same(tree.New(1), tree.New(1)))
	fmt.Println("different: ", Same(tree.New(1), tree.New(2)))
}
