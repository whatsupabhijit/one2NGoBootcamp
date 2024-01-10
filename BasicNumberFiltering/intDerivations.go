package main

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

func isOddPrime(n int) bool {
	return isOdd(n) && isPrime(n)
}