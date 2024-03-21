package main

import (
	"fmt"
	"strings"
)

// "golang.org/x/tour/wc"

func WordCount(s string) map[string]int {
	// I could just declare then make, but whatever
	// Using literal to initialize is easier
	count := map[string]int{}
	// First is always index (int)
	for _, str := range strings.Fields(s) {
		count[str] += 1
	}

	return count
}

func main() {
	// wc.Test(WordCount)
	fmt.Println(WordCount("hello world can you see this hello"))
}
