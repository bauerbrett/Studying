package main

import "fmt"

/*
Given a list of integers, determine the count of numbers for which there exists another number in the list that is greater by exactly one unit.

In other words, for each number x in the list, if x + 1 also exists in the list, then x is considered for the count.

Examples:

Example 1:

Input: [4, 3, 1, 5, 6]
Expected Output: 3
Justification: The numbers 4, 3, and 5 have 5, 4, and 6 respectively in the list, which are greater by exactly one unit.
*/
func countElements(l []int) int {
	set := map[int]bool{}
	for _, val := range l {
		set[val] = true
	}
	count := 0
	for val := range set {
		if _, ok := set[val+1]; ok {
			count++
		}
	}
	return count
}

func main() {
	//list := []int{4, 3, 1, 5, 6}
	list1 := []int{7, 8, 9, 10}

	fmt.Println(countElements(list1))

}
