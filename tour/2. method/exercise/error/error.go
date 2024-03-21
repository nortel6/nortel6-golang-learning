package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

// Fun fact: Directly print e will recurse infinitely
// because that is counted as printing ErrNegativeSqrt,
// since ErrNegativeSqrt implements error interface,
// this function will be called again
// Thus, e must not be directly printed
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	z := 1.0
	for prev_z := 0.0; math.Abs(prev_z-z) > 0.000000000001; {
		prev_z = z
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)
	}
	return z, nil
}

func main() {

	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
