package main

import "fmt"

type Solution struct{}

func (s *Solution) highestNum(matr [][]int) []int {

	//Need to find the arr with the most 1s
	// Return the index of the arr in the matr and the count of 1s in the arr
	returnArr := []int{}
	highOnes := 0
	indexHigh := 0

	for i, value := range matr {
		sumOnes := 0
		for j := 0; j < len(value); j++ {
			if value[j] == 1 {
				sumOnes += value[j]
			}
		}
		if sumOnes > highOnes {
			highOnes = sumOnes
			indexHigh = i
		}
	}
	returnArr = append(returnArr, indexHigh, highOnes)

	return returnArr

}

func main() {

	solution := &Solution{}
	dataInputs := [][]int{
		{1, 0},
		{1, 1},
		{0, 1},
	}

	fmt.Println("Highest Num:", solution.highestNum(dataInputs))

}
