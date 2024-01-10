package numberFiltering

// Given a list of integers, write a program to return only the even numbers from this list.
func evenNumbers(numbers []int) []int {
	return Filter(numbers, isEven)
}

// Given a list of integers, write a program to return only the odd numbers from this list.
func oddNumbers(numbers []int) []int {
	return Filter(numbers, isOdd)
}

// Given a list of integers, write a program to return only the odd numbers from this list.
func primeNumbers(numbers []int) []int {
	return Filter(numbers, isPrime)
}

// Given a list of integers, write a program to return only the odd prime numbers from this list.
// oddPrimeNumbers := numbersWrapper.filter(isOddPrime)
// oddPrimeNumbers := GetIntSlice(numbersWrapper.filter(isOdd)).filter(isPrime)
func oddPrimeNumbers(numbers []int) []int {
	return filter(numbers, oddPrimePredicates()...)
}

// Given a list of integers, write a program to return only the even and multiples of 5 from this list.
// evenMultipleOf5Numbers := numbersWrapper.filter(isEvenMultipleOf(5))
func evenMultipleOf(numbers []int) []int {
	return filter(numbers, evenMultipleOfPredicates()...)
}