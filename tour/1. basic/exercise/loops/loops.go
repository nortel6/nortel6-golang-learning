package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := x
	for prev_z := 0.0; math.Abs(prev_z-z) > 0.000000000001; {
		prev_z = z
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(49))
}
