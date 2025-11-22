package main

import (
	"fmt"
	"math"
)

// Use two maps. One holds index (only first value, if exist dont add index to map), the other holds the count it appears.
// We then search for the rune with the smallest index that has a count of 1
func nonrepeat(s string) int {

	existIndex := map[rune]int{}
	count := map[rune]int{}
	//valid := map[string]int{}
	for i, value := range s {
		count[value]++
		if _, ok := existIndex[value]; !ok {
			existIndex[value] = i
		}
	}
	small := math.MaxInt32
	for key, val := range count {
		if val == 1 && existIndex[key] < small {
			small = existIndex[key]
		}
	}

	if small == math.MaxInt32 {
		return -1
	}
	return small
}

func main() {
	a := "apple"
	//b := "abcab"
	c := "abab"
	for _, val := range a {
		fmt.Println(string(val))
	}
	fmt.Println(nonrepeat(c))

}
