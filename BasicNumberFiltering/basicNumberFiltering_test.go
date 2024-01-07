package main

import (
	"reflect"
	"testing"
)

func TestEvenNumbers(t *testing.T) {
	// given
	numRange := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	want := []int{2, 4, 6, 8, 10}

	// when
	output := GetIntSlice(numRange).filter(isEven)

	// then
	if !reflect.DeepEqual(output, want) {
		t.Errorf("wanted: %v but got: %v", want, output)
	}
}