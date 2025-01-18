package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[len(arr)/2]
	var left, middle, right []int
	for _, x := range arr {
		switch {
		case x < pivot:
			left = append(left, x)
		case x == pivot:
			middle = append(middle, x)
		default:
			right = append(right, x)
		}
	}
	return append(append(quickSort(left), middle...), quickSort(right)...)
}

func main() {
	arr := []int{3, 6, 8, 10, 1, 2, 1}
	sortedArr := quickSort(arr)
	fmt.Println("Sorted array:", sortedArr)
}

//This implementation of quick sort has a time complexity of O(nlogn), on average.
// In the worst case, the time complexity is O(n^2), if the pivot is always the smallest or largest element in the array.

//The quick sort algorithm works by selecting a pivot element and partitioning the array into three subarrays: one containing
//elements less than the pivot, one containing elements equal to the pivot, and one containing elements greater than the pivot. It then recursively sorts the left and right subarrays.
