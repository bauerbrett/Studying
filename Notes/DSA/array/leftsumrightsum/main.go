package main

import (
	"fmt"
	"math"
)

type Solution struct{}

func (s *Solution) leftrightsum(nums []int) []int {
	var differenceArray []int

	for i := 0; i < len(nums); i++ {
		var rightSum int
		for j := i + 1; j < len(nums); j++ { // this will loop through to teh end of the array and add the numbers that are after the index i.
			rightSum += nums[j]
		}
		var leftSum int
		for k := 0; k < i; k++ { // will loop through and add the numbers starting from o index i up to the index we are at in the firts for loop.
			leftSum += nums[k]
		}
		difference := int(math.Abs(float64(rightSum - leftSum))) // This is needed to keep it positive integers.
		differenceArray = append(differenceArray, difference)
	}
	return differenceArray
}

/*
func (s *Solution) findDifferenceArray(nums []int) []int { //This func speeds it up because it precomputes the total sum and subtracts from the indexes to find sum. This speeds it up to O(n)
	n := len(nums)
	differenceArray := make([]int, n)
	leftSum := 0
	rightSum := 0

	// Calculate the total sum of the array
	for _, num := range nums {
		rightSum += num
	}

	// Calculate the difference between left and right sums for each position
	for i := 0; i < n; i++ {
		rightSum -= nums[i]
		differenceArray[i] = int(math.Abs(float64(rightSum - leftSum)))
		leftSum += nums[i] // This is calculated after the difference is created because the right sum is already added up in the beginning so we just subtract the index
		// for leftsum we need to calculate this after because the right sum keeps subtraccting as we go up but the left some goes up as we continue.
		// Basically as we subtract from teh right some we add to the left sum.
	}

	return differenceArray
}
*/

func main() {

	solution := &Solution{}

	testInputs := [][]int{
		{2, 5, 1, 6, 1},
		{3, 5, 7, 8, 2, 1},
		{5, 6, 8, 9, 2, 4},
	}

	for _, x := range testInputs {
		difference := solution.leftrightsum(x)
		fmt.Println(difference)

	}

}
