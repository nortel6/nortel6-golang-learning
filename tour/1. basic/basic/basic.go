package main

// Multi import with parentheses
import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

// You can even name the return before hand...
// Useful for short functions I guess
// Or as documentation
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// You can shorten the variable type as well
// Also multiple return looks like this
func swap(x, y string) (string, string) {
	return y, x
}

// Funny syntax, https://go.dev/blog/declaration-syntax
func add(x int, y int) int {
	return x + y
}

// Global (package level) variables
var c, python, java bool

// Cannot use short form declaration outside of function scope
// globaltest string := "test"
var globaltest string = "test"

// Apparently you can also use parentheses for this
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	// An untyped constant takes the type needed by its context.
	// Meaning Big can be taken as float, but not int (overflows)

	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	// Imported function's name depends on the import
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println("Test", math.Sqrt(4))
	// Exported variables always start with Capital
	// Only then are they accessible
	fmt.Println(math.Pi)
	// Function call works the same, also just realized
	// you can put semicolon, and go VSCode extension removes it
	// when you save the file.
	fmt.Println(add(42, 17))
	fmt.Println(swap("hello", "world"))
	fmt.Println(split(17))

	// Functionl evel variables
	var i int
	// If not initialized, default to "zero value"
	// depending on the type
	fmt.Println(i, c, python, java)

	// You can initialize and dynamically determine the type
	var x, y int = 1, 2
	// Also go is a keyword
	var golang, lisp, lox = true, false, "no!"
	fmt.Println(x, y, golang, lisp, lox)

	// Short variable declarations
	// Just replaces var syntax
	// The rest is the same
	// Although, this is not doable outside of function scope
	k := 3
	fmt.Println(k)

	// Also printf
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// Behavior is pretty expected
	// i = 42
	// var f float64 = math.Sqrt(float64(x*x + y*y))
	// var u uint = uint(f)
	i = 42
	// Argument must match variable type
	// f := math.Sqrt(x*x + y*y)
	f := math.Sqrt(float64(x*x + y*y))
	u := uint(f)
	fmt.Println(i, f, u)

	// Type inference
	// To be expected at this point
	a := 3.23
	fmt.Printf("type %T\n", a)

	// Constant
	// No short-form
	const World = "世界"
	fmt.Println("Hello", World)

	fmt.Println(needInt(Small))
	// Type not allowed, overflow detected
	// fmt.Println(needInt(Big))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
