package main

import "fmt"

func bubbleSort(arr []int) {
	n := len(arr)
	fmt.Println(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ { //5 on first run
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	bubbleSort(arr)
	fmt.Println("Sorted array is:", arr)
}

//This bubble sort algorithm has a time complexity of (O)n^2
//since the running time is quadratic in the size of the input. If the input size is n, it will take approximately n² steps to complete the bubble sort.
