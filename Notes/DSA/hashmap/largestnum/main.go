package main

import (
	"fmt"
	"math"
)

func largestNum(arr []int) int {
	nums := map[int]int{}
	for _, val := range arr {
		nums[val]++
	}
	high := math.MinInt32
	for key, value := range nums {
		if value == 1 && key > high {
			high = key
		}
	}
	if high == math.MinInt32 {
		return -1
	}
	return high
}

func main() {
	//arr := []int{5, 7, 3, 7, 5, 8}
	arr1 := []int{1, 2, 3, 2, 1, 4, 4}
	fmt.Println(largestNum(arr1))

}
