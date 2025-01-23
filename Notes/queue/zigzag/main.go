package main

import (
	"fmt"
)

type Solution struct{}

func hasNext(arr []int) bool {
	return len(arr) >= 1
}

// Big array goes first in arr1 and small array goes second arr2
func next(arr1 []int, arr2 []int) []int {
	result := []int{}
	i := 0
	size1 := len(arr1)
	size2 := len(arr2)
	var bigger int
	if size1 > size2 {
		bigger = size1
	} else {
		bigger = size2
	}

	for j := 0; j < bigger; j++ {
		if hasNext(arr1) {
			result = append(result, arr1[i])
			arr1 = arr1[1:]
		}
		if hasNext(arr2) {
			result = append(result, arr2[i])
			arr2 = arr2[1:]
		}
	}
	return result

}

func (s *Solution) zigZag(arr1 []int, arr2 []int) []int {
	return next(arr1, arr2)
}

func main() {

	solution := &Solution{}

	arr1 := []int{1, 2, 3, 4}
	arr2 := []int{5, 6}

	fmt.Println("Zig Zag:", solution.zigZag(arr1, arr2))

}
