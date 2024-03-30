package numberfiltering

// Holy baby, I actually wrote something pretty good

/*
 * This design doesn't work well if we want
 * any condition to be true
 * Can only match ALL conditions

 * What can I do to make it more dynamic/flexible?
 * Pass in a function that has all the condition...
 * If I can build an "Any" function that generates
 * a function for comparison dynamically...
 */

// A little more dynamic
func FilterNumbers(nums []int, f func(int) bool) []int {
	ret := make([]int, 0)

	for _, num := range nums {
		if f(num) {
			ret = append(ret, num)
		}
	}

	return ret
}

// Honestly, these FilterX functions shouldn't even exist
// should be inline in the test cases only
// But whatever, I will just leave it here

func FilterEven(nums []int) []int {
	return FilterNumbers(nums, IsEven)
}

func FilterOdd(nums []int) []int {
	return FilterNumbers(nums, IsOdd)
}

func FilterPrime(nums []int) []int {
	return FilterNumbers(nums, IsPrime)
}

// What's the point of this? You cannot even have an even prime number
// except 2
func FilterOddPrime(nums []int) []int {
	return FilterNumbers(nums, All(IsOdd, IsPrime))
}

func FilterEvenMultiplesOfFive(nums []int) []int {
	return FilterNumbers(nums, All(IsEven, MultipleOf(5)))
}

// I shouldn't name it FilterOddMultiplesOfThreeGreaterThanTen
// I will just call it FilterStorySix lol
func FilterStorySix(nums []int) []int {
	// I can do this as well
	// return FilterNumbers(nums, All(IsOdd, MultipleOf(3), func(n int) bool {
	// 	return n > 1
	// }))
	return FilterNumbers(nums, All(IsOdd, MultipleOf(3), GreaterThan(10)))
}

// I only implement the first case here
func FilterStorySeven(nums []int) []int {
	return FilterNumbers(nums, All(IsOdd, MultipleOf(3), GreaterThan(5)))
}

// I only implement the first case here
func FilterStoryEight(nums []int) []int {
	return FilterNumbers(nums, Any(IsPrime, GreaterThan(15), MultipleOf(5)))
}

/*
 * I don't think this is what the expectations wanted,
 * but I will keep it.

 * This implementation basically uses closure to
 * "compress" all comparison (function) into a
 * single function.
 * Theoretically, you can add extends this very easily.

 * If this is not the expectations, I honestly don't know
 * How should it even look like.
 */

// Varadic functions
// Can take any amount of parameters, as long as the type matches

// int version
// func All(fs ...func(int) bool) func(int) bool {

// Generic version
func All[T any](fs ...func(T) bool) func(T) bool {
	// Does this count as closure?
	// Returns a function that depends on the fs parameter...
	// I think yes this is closure
	// So you can dynamically "build" a function... interesting
	// wow, never have I thought I can write something this beautiful
	return func(n T) bool {
		for _, f := range fs {
			// Not even one test should fail
			if !f(n) {
				return false
			}
		}
		return true
	}
}

// func Any(fs ...func(int) bool) func(int) bool {
// With Generic
func Any[T any](fs ...func(T) bool) func(T) bool {
	return func(n T) bool {
		for _, f := range fs {
			// As long as any test return true
			if f(n) {
				return true
			}
		}
		return false
	}
}

// Since IsMultipleOf(n, m) doesn't work
// We will need to use closure to generate a function
// that matches the parameters
// in order to be pass into FilterNumbers()
func MultipleOf(factor int) func(int) bool {
	return func(n int) bool {
		return n%factor == 0
	}
}

func GreaterThan(m int) func(int) bool {
	return func(n int) bool {
		return n > m
	}
}

func LesserThan(m int) func(int) bool {
	return func(n int) bool {
		return n < m
	}
}

func IsEven(n int) bool {
	return n%2 == 0
}

func IsOdd(n int) bool {
	return n%2 == 1
}

// https://en.wikipedia.org/wiki/Primality_test#Go
// Optimized Trial Divison Implementation from Wikipedia
// Tests until sqrt(num) is sufficient
// I don't really understand the optimization here...
func IsPrime(num int) bool {
	if num > 1 && num <= 3 {
		return true
	}
	if num <= 1 || num%2 == 0 || num%3 == 0 {
		return false
	}

	for i := 5; i*i <= num; i += 6 {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
	}
	return true
}
