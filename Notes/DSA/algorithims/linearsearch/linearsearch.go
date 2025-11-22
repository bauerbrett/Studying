package linear

import "fmt"

func linearSearch(arr []int, x int) int {
	for i, n := range arr {
		if n == x {
			return i
		}
	}
	return -1
}

func linear() {
	arr := []int{3, 5, 7, 9, 11}
	x := 7
	fmt.Println(linearSearch(arr, x)) //going to return 2 because that is the place 7 is.
}

//This linear search algorithm has a time complexity of O(n),
//since the running time is directly proportional to the size of the input. If the input size is , the linear search will take n steps to complete.

//This linear search algorithm has a space complexity of O(1), since it only requires a
//constant amount of memory regardless of the size of the input. It creates a single variable (i) to store the loop index, but the memory required for this variable is constant.
