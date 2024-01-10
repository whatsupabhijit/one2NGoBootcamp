package numberFiltering

func isEven(number int) bool {
	return number % 2 == 0
}

func isOdd(number int) bool {
	return !isEven(number)
}

func isPrime(n int) bool {
	if n == 1 {
		return false
	}

	for i := 2; i * i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func isMultipleOf5(number int) bool {
	return number % 5 == 0
}

func oddPrimePredicates() []Predicate {
	predicates := []Predicate {
		isOdd,
		isPrime,
	}
	return predicates
}

func evenMultipleOfPredicates() []Predicate {
	predicates := []Predicate {
		isEven,
		isMultipleOf5,
	}
	return predicates
}