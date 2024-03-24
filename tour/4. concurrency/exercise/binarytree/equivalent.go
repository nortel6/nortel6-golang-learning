package main

import "fmt"

// import "golang.org/x/tour/tree"
//import "tree"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
// func Walk(t *tree.Tree, ch chan int)
func Walk(t *Tree, ch chan int) {
	// This declaration is needed
	// so it could call itself
	var walk func(t *Tree)
	walk = func(t *Tree) {
		if t == nil {
			return
		}

		walk(t.Left)
		ch <- t.Value
		walk(t.Right)
	}
	walk(t)
	// Should this be deferred?
	defer close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
// func Same(t1, t2 *tree.Tree) bool
func Same(t1, t2 *Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)

	for x := range c1 {
		select {
		case y := <-c2:
			if x != y {
				return false
			}
		}
	}

	return true
}

func main() {
	ch := make(chan int)
	go Walk(New(1), ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println(Same(New(1), New(1)))
	fmt.Println(Same(New(1), New(2)))
	fmt.Println(Same(New(2), New(2)))
}
