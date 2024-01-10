package numberFiltering

// type IntSlice struct {
// 	data []int
// }

// func GetIntSlice(slice []int) IntSlice {
// 	return IntSlice {data: slice}
// }

// func (numbers IntSlice) filter(predicate func(int) bool) []int {
// 	var result []int
// 	for _, num := range numbers.data {
// 		if predicate(num) {
// 			result = append(result, num)
// 		}
// 	}
// 	return result
// }

// func (numbers IntSlice) filter(predicate func(int) bool) []int {
// 	var result []int
// 	for _, num := range numbers.data {
// 		if predicate(num) {
// 			result = append(result, num)
// 		}
// 	}
// 	return result
// }


func Filter(numbers []int, predicate func(int) bool) []int {
	var result []int
	for _, num := range numbers {
		if predicate(num) {
			result = append(result, num)
		}
	}
	return result
}



type Predicate func(n int) bool

type PredicateList []Predicate

// 1 2 3
// prime odd

func (predicate PredicateList) forAll(number int) bool {
	for _, p := range predicate {
		if !p(number) {
			return false
		}
	}

	return true
}


func filter(numbers []int, predicates ...Predicate) []int {
	var result []int
	for _, num := range numbers {
		allPredicateSatisfied := true
		for _, predicate := range predicates {
			if !predicate(num) {
				allPredicateSatisfied = false
				break
			}
		}
		if allPredicateSatisfied {
			result = append(result, num)
		}
	}
	return result
}
