package main

import (
	"fmt"
	"math"
)

// No class
// But this is how you define a "method"
// In Go, it is a function with receiver argument
// Appears before function name
type Vertex struct {
	X, Y float64
}

// This is kinda wacky
// Method
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// This is a normal function
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// With pointer receiver, the modification of
// the state will persists
// Just like... a pointer
/*
 * This is more efficient, because we don't need
 * another copy of the data.
 * For a data type,
 * choose either pointer, or value receivers only
 * Don't mix match!
 */
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Apparenlt doesn't have to be a struct
// You can only declare method when the type
// is within the same package.
// In other words, you cannot add additional methods
// for built-in types directly
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := Vertex{3, 4}
	// Different ways of calling
	fmt.Println(v.Abs())
	fmt.Println(Abs(v))

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	// Equivalent of
	// (&v).Scale(10)
	// If that make sense to you
	v.Scale(10)
	fmt.Println(v.Abs())
	// In the case of a normal pointer function,
	// You need to explicitly pass in the address
	// of the value
	Scale(&v, 10)
	fmt.Println(v.Abs())

	p := &Vertex{4, 3}
	// Passing in pointer to normal receiver
	// Equivalent to (*p).Abs()
	// Only pass in the value
	// Basically, receiver controls what to receive
	// Very convenient...
	fmt.Println(p.Abs())
	fmt.Println(Abs(*p))
	// Of course, function cannot do the above
	// Compile Error!
	// fmt.Println(Abs(p))

}
