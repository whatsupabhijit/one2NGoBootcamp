package numberFiltering

func isEven(number int) bool {
	return number % 2 == 0
}

func isOdd(number int) bool {
	return !isEven(number)
}

// TODO: Write HOF so that you don't need to write multiple functions. WRite a blog about it
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
	for i := 2; i * i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return n > 1
}

func isMultipleOf3(number int) bool {
	return number % 3 == 0
}

func isMultipleOf5(number int) bool {
	return number % 5 == 0
}