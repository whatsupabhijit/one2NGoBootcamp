package main

import "fmt"

func main() {
	numbers := []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Given a list of integers, write a program to return only the even numbers from this list.
	evenNumbers := GetIntSlice(numbers).filter(isEven)

	fmt.Println("all Numbers: " , numbers)
	fmt.Println("even Numbers: " , evenNumbers)
 }