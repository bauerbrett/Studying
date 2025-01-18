package main

import "fmt"

type Solution struct{}

func (s *Solution) largestAltitude(gain []int) int {
	maxAltitude := 0
	altitude := 0
	// TODO: Write your code here
	// Basically all of these points are different altitudes and each point is different altitudes. So if the first altitude is 2 and teh next is 4. The
	// altitude at at index 1 (point2) is 4. If if none of the other alitudes in the array after this take teh altitude over 4 than 4 is the highest altitude.
	for _, g := range gain { // Loop through the altitudes in array
		altitude += g               // This is teh net gain between the indexs in the array. So index[0] would be teh first altitude, and then index[1] would get added to [1].
		if altitude > maxAltitude { // if the altitude is ever bigger than the maxaltitude, make the maxaltitude that altidue.
			maxAltitude = altitude
		}
	}

	return maxAltitude
}

func main() {

	solution := &Solution{}

	testInputs := [][]int{
		{-5, 1, 5, 0, -7},
		{4, -3, 2, -1, -2},
		{2, 2, -3, -1, 2, 1, -5},
	}

	for _, x := range testInputs {
		difference := solution.largestAltitude(x)
		fmt.Println(difference)

	}

}
