package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	prev, curr := 0, 1
	return func() int {
		next := prev + curr
		prev = curr
		curr = next
		return curr - prev
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
