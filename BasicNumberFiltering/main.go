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
	return filter(numbers, true, oddPrimePredicates()...)
}

// Given a list of integers, write a program to return only the even and multiples of 5 from this list.
// evenMultipleOf5Numbers := numbersWrapper.filter(isEvenMultipleOf(5))
func evenMultipleOf(numbers []int) []int {
	return filter(numbers, true, evenMultipleOfPredicates()...)
}

// Given a list of integers, write a program to return only the odd and multiples of 3 greater than 10 from this list.
func oddMultipleOf3Gt10(numbers []int) []int {
	return filter(numbers, true, oddMultipleOf3Gt10Predicates()...)
}

// Given a list of integers, and a set of conditions (odd, even, greater than 5, multiple of 3, prime, 
// and many more such custom conditions that may be dynamically defined by user),
// write a program to return only the integers from the given list that match ALL the conditions.
func filterAll(numbers []int, predicates ...Predicate) []int {
	return filter(numbers, true, predicates...)
}

// Given a list of integers, and a set of conditions (odd, even, greater than 5, multiple of 3, prime,
// and many more such custom conditions that may be dynamically defined by user),
// write a program to return only the integers from the given list that match ANY of the conditions.
func filterAny(numbers []int, predicates ...Predicate) []int {
	return filter(numbers, false, predicates...)
}