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
	predicates := PredicateList {isOdd, isPrime}

	// when
	output := oddPrimeNumbers(numRange, predicates...)

	// then
	if !reflect.DeepEqual(output, want) {
		t.Errorf("wanted: %v but got: %v", want, output)
	}
}


func TestEvenMultipleOf5Numbers(t *testing.T) {
	// given
	numRange := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	want := []int{10}
	predicates := PredicateList {isEven, isMultipleOf5}

	// when
	output := evenMultipleOf(numRange, predicates...)

	// then
	if !reflect.DeepEqual(output, want) {
		t.Errorf("wanted: %v but got: %v", want, output)
	}
}

func TestOddMultipleOf3GreaterThan10Numbers(t *testing.T) {
	// given
	numRange := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	want := []int{15}
	predicates := PredicateList {isOdd, isMultipleOf3, isGreaterThan10}

	// when
	output := oddMultipleOf3Gt10(numRange, predicates...)

	// then
	if !reflect.DeepEqual(output, want) {
		t.Errorf("wanted: %v but got: %v", want, output)
	}
}

// TODO: Merge the TC1 and TC2 in a single test case single function, how do you do it? Blog about it once you learn.
func TestFilterAllTC1(t *testing.T) {
	// given
	numRange := []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	want := []int {9, 15}
	predicates := PredicateList {isOdd, isMultipleOf3, isGreaterThan5}

	// when
	output := filterAll(numRange, predicates...)

	// then
	if !reflect.DeepEqual(output, want) {
		t.Errorf("wanted: %v but got: %v", want, output)
	}
}

func TestFilterAllTC2(t *testing.T) {
	// given
	numRange := []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	want := []int {6, 12}
	predicates := PredicateList {isEven, isMultipleOf3, isLessThan15}

	// when
	output := filterAll(numRange, predicates...)

	// then
	if !reflect.DeepEqual(output, want) {
		t.Errorf("wanted: %v but got: %v", want, output)
	}
}


func TestFilterAnyTC1(t *testing.T) {
	// given
	numRange := []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	want := []int {2, 3, 5, 7, 10, 11, 13, 15, 16, 17, 18, 19, 20}
	predicates := PredicateList {isPrime, isMultipleOf5, isGreaterThan15}

	// when
	output := filterAny(numRange, predicates...)

	// then
	if !reflect.DeepEqual(output, want) {
		t.Errorf("wanted: %v but got: %v", want, output)
	}
}


func TestFilterAnyTC2(t *testing.T) {
	// given
	numRange := []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	want := []int {1, 2, 3, 4, 5, 6, 9, 12, 15, 18}
	predicates := PredicateList {isMultipleOf3, isLessThan6}

	// when
	output := filterAny(numRange, predicates...)

	// then
	if !reflect.DeepEqual(output, want) {
		t.Errorf("wanted: %v but got: %v", want, output)
	}
}