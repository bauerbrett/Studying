package main

import (
	"fmt"
	"sort"
)

//Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

//Example 1:
//Input: nums= [1, 2, 3, 4]
////Output: false
//Explanation: There are no duplicates in the given array.

//Example 2:
//Input: nums= [1, 2, 3, 1]
//Output: true
//Explanation: '1' is repeating.

// Solution struct type
type Solution struct{}

// containsDuplicate checks for duplicates in a slice of integers
func (s *Solution) containsDuplicate(nums []int) bool {
	seen := make(map[int]int)
	if len(nums) == 0 || nums == nil {
		fmt.Println("No integers in array")
		return false
	}

	// Count occurrences of each number
	for _, num := range nums {
		seen[num]++
	}

	// Print how many times '1' appears
	fmt.Printf("In the array %v, the number 1 appears %d times.\n", nums, seen[1])

	// Check for duplicates
	for _, count := range seen {
		if count > 1 {
			return true
		}
	}
	return false
}

func (s *Solution) bruteForce(nums []int) bool { //Bruteforce, checks every single item in after i to see if it == i. This takes the most time O(n^2)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

func (s *Solution) sortingDupli(nums []int) bool { //sorting func, this will move dupes closer together which allows us to just add 1 on index to check
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] == nums[i+1] {
			return true
		}

	}
	return false
	// This func is O(nlogn) because the sort is the biggest time complexity here which is O(nlogn). The for loop is just one and is liner O(n) so this is dominated by the other.
}

func main() {
	solution := &Solution{}

	testInputs := [][]int{
		{1, 4, 5, 3, 1},
		{1, 2, 3, 4, 5, 1},
		{1, 3, 5, 8, 9, 9},
	}

	for _, input := range testInputs {
		output := solution.containsDuplicate(input)
		fmt.Println(output)
		/*
				if output {
					fmt.Println("Duplicates are in the array:", input)
				} else {
					fmt.Println("No duplicates in array:", input)
				}
			}
		*/

	}
}
