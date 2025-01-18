package main

import "fmt"

func sumArrays(arr *[5]int) [5]int {
	var result [len(arr)]int

	if arr == nil || len(arr) == 0 {
		return [5]int{}
	}
	result[0] = arr[0] // Just set the first index as the value because then we can start the loop at 1 and there is not previous results to add to this since it is the first.
	for i := 1; i < len(arr); i++ {
		result[i] = result[i-1] + arr[i] //Just going to add the last result of result with the next array index the for loop is on. So if i = 1, it will result[1] = result[0] + arr[1]
	}
	return result
}

func main() {
	arr := [5]int{1, 4, 7, 10, 12}
	sum := sumArrays(&arr)
	fmt.Println(sum)
}
