package main

import "fmt"

type Solution struct{}

func (s *Solution) diagonalSum(matr [][]int) int {

	n := len(matr)

	totalSum := 0
	// So for this since this is a square i is going to work for the matr index and the arr index. Each time i goes up it will move to next arr and next value in arr.
	// And same for when you grab the right diagonal, for example when i == 2, it would be matr[2][3-1-2] which is the bottom left of the square.
	for i := 0; i < n; i++ {
		totalSum += matr[i][i]     // Add primary diagonal element
		totalSum += matr[i][n-1-i] // Add secondary diagonal element
	}

	if n%2 == 1 { // if odd
		fmt.Println(matr[n/2][n/2])
		totalSum -= matr[n/2][n/2]
	}

	return totalSum

}

func main() {

	solution := &Solution{}
	dataInputs := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("Diagonal Sum:", solution.diagonalSum(dataInputs))

}
