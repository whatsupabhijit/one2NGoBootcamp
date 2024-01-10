package numberFiltering

import (
	"reflect"
	"testing"
)

func TestEvenNumbers(t *testing.T) {
	// given
	numRange := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	want := []int{2, 4, 6, 8, 10}

	// when
	// output := GetIntSlice(numRange).filter(isEven)
	output := evenNumbers(numRange)

	// then
	if !reflect.DeepEqual(output, want) {
		t.Errorf("wanted: %v but got: %v", want, output)
	}
}


func TestOddNumbers(t *testing.T) {
	// given
	numRange := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	want := []int{1, 3, 5, 7, 9}

	// when
	// output := GetIntSlice(numRange).filter(isOdd)
	output := oddNumbers(numRange)

	// then
	if !reflect.DeepEqual(output, want) {
		t.Errorf("wanted: %v but got: %v", want, output)
	}
}

func TestPrimeNumbers(t *testing.T) {
	// given
	numRange := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	want := []int{2, 3, 5, 7}

	// when
	// output := GetIntSlice(numRange).filter(isPrime)
	output := primeNumbers(numRange)

	// then
	if !reflect.DeepEqual(output, want) {
		t.Errorf("wanted: %v but got: %v", want, output)
	}
}

func TestOddPrimeNumbers(t *testing.T) {
	// given
	numRange := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	want := []int{3, 5, 7}

	// when
	// output := GetIntSlice(numRange).filter(isOddPrime)
	// output := GetIntSlice(GetIntSlice(numRange).filter(isOdd)).filter(isPrime)
	output := oddPrimeNumbers(numRange)

	// then
	if !reflect.DeepEqual(output, want) {
		t.Errorf("wanted: %v but got: %v", want, output)
	}
}


func TestEvenMultipleOf5Numbers(t *testing.T) {
	// given
	numRange := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	want := []int{10}

	// when
	// output := GetIntSlice(numRange).filter(isOddPrime)
	// output := GetIntSlice(GetIntSlice(numRange).filter(isOdd)).filter(isPrime)
	output := evenMultipleOf(numRange)

	// then
	if !reflect.DeepEqual(output, want) {
		t.Errorf("wanted: %v but got: %v", want, output)
	}
}