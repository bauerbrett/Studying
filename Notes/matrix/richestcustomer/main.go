package main

import "fmt"

type Solution struct{}

func (s *Solution) calcRichest(accounts [][]int) int {
	richest := 0

	for _, value := range accounts { //Loop through all of the arrays
		amount := 0 // Before entering the loop for the individual array set an amount variable that will sum tje total of values in the account.
		for j := 0; j < len(value); j++ {
			amount += value[j] // In the loop add each value of the array to amount for this array
		}
		if amount > richest { //Check if this accounts total is higher than the exisiting highest, if it is replace it with this value
			richest = amount
		}

	}
	return richest

}

func main() {

	solution := &Solution{}

	dataInput := [][]int{
		{5, 2, 3},
		{0, 6, 7},
		{1, 2},
		{3, 4},
		{5, 6, 100},
	}

	richest := solution.calcRichest(dataInput)
	fmt.Println("The richest customer has:", richest)

}
