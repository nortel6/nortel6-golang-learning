package numberfiltering

import (
	"reflect"
	"runtime"
	"testing"
)

func TestEven(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	want := []int{2, 4, 6, 8, 10}
	got := FilterNumbers(data, IsEven)

	IsSameList(t, want, got)
}

func TestOdd(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	want := []int{1, 3, 5, 7, 9}
	got := FilterNumbers(data, IsOdd)

	IsSameList(t, want, got)
}

func TestPrime(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	want := []int{2, 3, 5, 7}
	got := FilterNumbers(data, IsPrime)

	IsSameList(t, want, got)
}

func TestOddPrime(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	want := []int{3, 5, 7}
	got := FilterNumbers(data, All(IsOdd, IsPrime))

	IsSameList(t, want, got)
}

func TestFilterEvenMultiplesOfFive(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	want := []int{10, 20}
	got := FilterNumbers(data, All(IsEven, MultipleOf(5)))

	IsSameList(t, want, got)
}

func TestFilterStorySix(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	want := []int{15}
	got := FilterNumbers(data, All(IsOdd, MultipleOf(3), GreaterThan(10)))

	IsSameList(t, want, got)
}

func TestFilterStorySeven(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	want := []int{9, 15}
	got := FilterNumbers(data, All(IsOdd, MultipleOf(3), GreaterThan(5)))
	IsSameList(t, want, got)

	want2 := []int{6, 12}
	got2 := FilterNumbers(data, All(IsEven, LesserThan(15), MultipleOf(3)))
	IsSameList(t, want2, got2)
}

func TestFilterStoryEight(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	want := []int{2, 3, 5, 7, 10, 11, 13, 15, 16, 17, 18, 19, 20}
	got := FilterNumbers(data, Any(IsPrime, GreaterThan(15), MultipleOf(5)))
	IsSameList(t, want, got)

	want2 := []int{1, 2, 3, 4, 5, 6, 9, 12, 15, 18}
	got2 := FilterNumbers(data, Any(LesserThan(6), MultipleOf(3)))
	IsSameList(t, want2, got2)
}

// Reusable test helper function
func IsSameList(t *testing.T, want []int, got []int) {
	// Test if length is the same
	n, m := len(want), len(got)
	if len(want) != len(got) {
		t.Errorf("'got' does not have the expected length;\n"+
			"want %v; got %v\n"+
			"len(want) %d; len(got) %d",
			want, got, n, m)
		return
	}

	// Test if values are the same
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("'got' does not match expected output;\n"+
				"value at index %d does not match;\n"+
				"want %v; got %v\n"+
				"want[%d] %d; got[%d] %d)\n",
				i, want, got, i, want[i], i, got[i])
			return
		}
	}
}

// https://stackoverflow.com/questions/7052693/how-to-get-the-name-of-a-function-in-go
// Passing in empty interface is just like passing in a general object
// No longer used, but I figured it is a cool thing to leave here
func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
