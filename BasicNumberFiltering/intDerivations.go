package main


func isEven(number int) bool {
	return number % 2 == 0
}

func isOdd(number int) bool {
	return !isEven(number)
}