package main

import (
	"fmt"
	"math"
	"time"
)

type Abser interface {
	Abs() float64
}

type Vertex struct {
	X, Y float64
}

// This is kinda wacky
// Method
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	// Should always handle nil case
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var a Abser
	v := Vertex{3, 4}
	f := MyFloat(-math.Sqrt2)
	// Interface implemented
	a = f
	// If the method does not receive a pointer,
	// both of this will still work
	a = &v
	// a = v
	fmt.Println(a.Abs())

	// Implicit implementation, no "implements" keyword
	var b Abser = &Vertex{6, 8}
	fmt.Println(b.Abs())

	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	// This is only doable if it is struct
	// i = &F(math.Pi)
	i = F(math.Pi)
	describe(i)
	i.M()

	// Even if t is nil, i interface is not
	// t is the value (implementation) of i interface
	var t *T
	i = t
	describe(i)
	i.M()

	var i2 I
	describe(i2)
	// This will cause an error, because it is nil
	// (No implementation)
	// i2.M()

	// Holds nothing, can be anything
	// Generic?
	var empty interface{}
	describe_general(empty)

	empty = 42
	describe_general(empty)

	empty = "hello World!"
	describe_general(empty)

	// Get the string value...
	s := empty.(string)
	fmt.Println(s)

	// Test whether string exists in empty interface
	s, ok := empty.(string)
	fmt.Println(s, ok)

	// Test whether float is in empty interface
	x, ok := empty.(float64)
	fmt.Println(x, ok)

	// Not doable because float64 is not in i interface
	// Isn't this like hashmap? Fucking crazy shit man
	// f = i.(float64)
	// fmt.Println(f)

	do(21)
	do("Hello")
	do(true)

	y := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(y, z)

	// Error handling
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

// You can describe the underlying implementation
func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func describe_general(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func do(i interface{}) {
	// Type switches
	// Able to handle different type differently
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

type Person struct {
	Name string
	Age  int
}

// For implementing Stringer inferface
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}
