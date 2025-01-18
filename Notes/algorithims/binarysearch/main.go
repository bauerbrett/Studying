package main

import (
	"fmt"
)

func binarySearch(arr []int, x int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] < x {
			low = mid + 1
		} else if arr[mid] > x {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	arr := []int{2, 3, 4, 10, 40}
	x := 10
	fmt.Println(binarySearch(arr, x)) //returns 3

}

//Binary search is O(logn)
//This binary search algorithm has a time complexity of O(logn),
//since the running time increases logarithmically with the size of the input. If the input size is , it will take approximately  steps to complete the binary search.
//logarithmically time - complexity increases by one unit for each doubling of data.

//logarithm of 8 for base 2 = log2(8) = 3,
//Explanation: 23 = 8 Since 2 needs to be raised to a power of 3 to give 8, Thus logarithm of 8 base 2 is 3.

//logarithm of 81 for base 9 = log9(81) = 2,
//Explanation: 92 = 81 Since 9 needs to be raised to a power of 2 to give 81, Thus logarithm of 81 base 9 is 2.

//Understanding Binary Search and Logarithms
//Binary Search repeatedly divides the search space into halves, so the number of steps required to find an element in a sorted array of size
//𝑛 is determined by how many times 𝑛 can be halved until only 1 element is left.

// Mathematically, the number of divisions is expressed as log2(n) = y where:
// n is the size of the array.
// 2 is the base beacause we halve the array each time
// teh result y is the number of splits or comparisions need

//For array size of 100:
//log2(100) == 6.64
//This means the binary search algorithm would take at most 7 comparisons to locate the target or determine it's not present.
