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

func (predicate PredicateList) forAll(number int) bool {
	for _, p := range predicate {
		if !p(number) {
			return false
		}
	}

	return true
}

func (predicate PredicateList) any(number int) bool {
	for _, p := range predicate {
		if p(number) {
			return true
		}
	}

	return false
}

func filter(numbers []int, isForAll bool, predicates ...Predicate) []int {
	var result []int
	predicateList := PredicateList(predicates)

	for _, num := range numbers {
		if isForAll && predicateList.forAll(num) {
			result = append(result, num)
		}
		if !isForAll && predicateList.any(num) {
			result = append(result, num)
		}
	}
	return result
}