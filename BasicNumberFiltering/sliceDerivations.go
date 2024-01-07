package main

type IntSlice struct {
	data []int
}

func GetIntSlice(slice []int) IntSlice {
	return IntSlice {data: slice}
}

func (numbers IntSlice) filter(predicate func(int) bool) []int {
	var result []int
	for _, num := range numbers.data {
		if predicate(num) {
			result = append(result, num)
		}
	}
	return result
}
