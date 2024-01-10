package main

import "fmt"

func main() {
	numbers := []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	numbersWrapper := GetIntSlice(numbers)

	// Given a list of integers, write a program to return only the even numbers from this list.
	evenNumbers := numbersWrapper.filter(isEven)

	// Given a list of integers, write a program to return only the odd numbers from this list.
	oddNumbers := numbersWrapper.filter(isOdd)

	// Given a list of integers, write a program to return only the odd numbers from this list.
	primeNumbers := numbersWrapper.filter(isPrime)

	// Given a list of integers, write a program to return only the odd prime numbers from this list.
	oddPrimeNumbers := numbersWrapper.filter(isOddPrime)

	fmt.Println("all Numbers: " , numbers)
	fmt.Println("even Numbers: " , evenNumbers)
	fmt.Println("odd Numbers: " , oddNumbers)
	fmt.Println("prime Numbers: " , primeNumbers)
	fmt.Println("odd prime Numbers: " , oddPrimeNumbers)

}