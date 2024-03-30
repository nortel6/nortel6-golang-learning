package numberfiltering

import (
	"reflect"
	"runtime"
	"testing"
)

func TestEven(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	want := []int{2, 4, 6, 8, 10}

	IsSameList(t, data, want, FilterEven)
}

func TestOdd(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	want := []int{1, 3, 5, 7, 9}

	IsSameList(t, data, want, FilterOdd)
}

func TestPrime(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	want := []int{2, 3, 5, 7}

	IsSameList(t, data, want, FilterPrime)
}

func TestOddPrime(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	want := []int{3, 5, 7}

	IsSameList(t, data, want, FilterOddPrime)
}

func TestFilterEvenMultiplesOfFive(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	want := []int{10, 20}

	IsSameList(t, data, want, FilterEvenMultiplesOfFive)
}

func TestFilterStorySix(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	want := []int{15}

	IsSameList(t, data, want, FilterStorySix)
}

func TestFilterStorySeven(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	want := []int{9, 15}
	IsSameList(t, data, want, FilterStorySeven)

	want2 := []int{6, 12}
	secondTest := func(nums []int) []int {
		return FilterNumbers(data, All(IsEven, LesserThan(15), MultipleOf(3)))
	}
	IsSameList(t, data, want2, secondTest)
}

func TestFilterStoryEight(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	want := []int{2, 3, 5, 7, 10, 11, 13, 15, 16, 17, 18, 19, 20}
	IsSameList(t, data, want, FilterStoryEight)

	want2 := []int{1, 2, 3, 4, 5, 6, 9, 12, 15, 18}
	secondTest := func(nums []int) []int {
		return FilterNumbers(data, Any(LesserThan(6), MultipleOf(3)))
	}
	IsSameList(t, data, want2, secondTest)
}

// Reusable test helper function
func IsSameList(t *testing.T, data []int, want []int, f func([]int) []int) {
	got := f(data)
	// Test if length is the same
	n, m := len(want), len(got)
	if len(want) != len(got) {
		t.Errorf("%s(%v) does not have the expected length;\n"+
			"want %v; got %v\n"+
			"len(want) %d; len(got) %d",
			GetFunctionName(f), data, want, got, n, m)
		return
	}

	// Test if values are the same
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("%s(%v) does not match expected output;\n"+
				"value at index %d does not match;\n"+
				"want %v; got %v\n"+
				"want[%d] %d; got[%d] %d)\n",
				GetFunctionName(f), data, i, want, got, i, want[i], i, got[i])
			return
		}
	}
}

// https://stackoverflow.com/questions/7052693/how-to-get-the-name-of-a-function-in-go
// Passing in empty interface is just like passing in a general object
func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
