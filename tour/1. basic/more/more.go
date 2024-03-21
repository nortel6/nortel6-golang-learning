package main

import (
	"fmt"
	"math"
	"strings"
)

// Struct, multiple fields
type Vertex struct {
	X, Y int
}

func main() {
	// Classic pointer
	// Go does not have pointer arithmetic though
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	// Access struct fields using .
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	// You can access struct fields with dot even as a pointer
	// Explicit dereference
	// Normally, it'd be like (*p).X
	p2 := &v
	p2.X = 1e9
	fmt.Println(v)

	v1 := Vertex{1, 2} // has type Vertex
	v2 := Vertex{X: 1} // Y:0 is implicit
	v3 := Vertex{}     // X:0 and Y:0
	p2 = &Vertex{1, 2} // has type *Vertex
	fmt.Println(v1, p2, v2, v3)

	// Array
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// Slices
	// It is references to array
	// Does not store any data
	// Just like Python's slices
	var s []int = primes[1:4]
	fmt.Println(s)

	// Modifying slices' elements = modifying the array
	s[2] = 71
	fmt.Println(primes)

	// You can create slice literal without declaring array
	// Array will be created, and the slice will refer to it
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	// 2D slices, pretty cool
	x := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(x)

	s = []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Extend its length.
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)

	// Out of bound error
	// s = s[:7]
	// printSlice(s)

	// Default slice is nil
	var ds []int
	fmt.Println(ds, len(ds), cap(ds))
	if ds == nil {
		fmt.Println("nil!")
	}

	// Using make to create a slice
	// len = 5
	a := make([]int, 5)
	printSlice2("a", a)

	// len = 0, cap = 5
	b := make([]int, 0, 5)
	printSlice2("b", b)

	c := b[:2]
	printSlice2("c", c)

	d := c[2:5]
	printSlice2("d", d)

	// Slices of slice
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)

	// var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	pow := make([]int, 10)

	// Don't need the value
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	// Need both
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	// Don't need the index
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

	// nil map when declared
	var m map[string]Vertex
	// This one initialized
	// var m = map[string]Vertex{}
	// Need to allocate first
	m = make(map[string]Vertex)

	m["Test"] = Vertex{1, 2}
	fmt.Println(m)

	// Literal (Initialization, I guess)
	var n = map[string]Vertex{
		// The Type is actually redundant here
		"Bell Labs": Vertex{
			40, -74,
		},
		"Google": {
			37, -122,
		},
	}
	fmt.Println(n)

	// Remove and assert
	delete(n, "Bell Labs")
	v, ok := n["Bell Labs"]
	fmt.Println(v, ok)

	// Function as a value
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	// Closure has their own fields (variables)
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func printSlice(s []int) {
	// Capacity is the number of elements in the array
	// (Counting starting from the first element in the slice)
	// Length is the number of elements in the current slice
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSlice2(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

// Equivalent to function... pointer I assume
// But instead, it is a value
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// This returns a function closure
// Each closure will have different sum
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
