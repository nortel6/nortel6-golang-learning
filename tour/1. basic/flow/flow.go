package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func sqrt(x float64) string {
	// If syntax in Go
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	// You can declare variable within if block only
	// Wow this is... useful
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// Cannot use v here
	return lim
}

func main() {
	sum := 0
	// Go for loop, only for loop exists
	// Wacky syntax...
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// Equivalent to while loop
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// Infinite loop
	// for {
	// }

	fmt.Println(sqrt(2), sqrt(-4))

	// Interesting observation here
	// Second statement's output is printed first
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	// This is switch
	fmt.Print("Go runs on ")
	// Finally, no more breaks
	// Doesn't have to be constant and integers
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	// Another example
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	t := time.Now()
	// You dont need to specific a condition
	// to evaluate
	// Can be a clean way to write if-else chain
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	// Delay call
	defer fmt.Println("world")
	fmt.Println("hello")

	// Reverse count...
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
