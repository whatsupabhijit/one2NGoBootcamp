package numberFiltering

func isEven(number int) bool {
	return number % 2 == 0
}

func isOdd(number int) bool {
	return !isEven(number)
}

func isLessThan15(number int) bool{
	return number < 15
}

func isLessThan6(number int) bool{
	return number < 6
}

func isGreaterThan5(number int) bool{
	return number > 5
}

func isGreaterThan10(number int) bool{
	return number > 10
}

func isGreaterThan15(number int) bool{
	return number > 15
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

func isMultipleOf3(number int) bool {
	return number % 3 == 0
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

func oddMultipleOf3Gt10Predicates() []Predicate {
	predicates := []Predicate {
		isOdd,
		isMultipleOf3,
		isGreaterThan10,
	}
	return predicates
}