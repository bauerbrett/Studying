package main

import "fmt"

type Solution struct{}

//See queue readme for more details on sliding window.
//This is a sliding window problem
// maxSub finds the maximum element for each subarray of size k
func (s *Solution) maxSub(arr []int, k int) []int {
	dq := []int{}
	n := len(arr)
	result := []int{}

	for i := 0; i < n; i++ {
		for len(dq) > 0 && dq[0] <= i-k+1 {
			dq = dq[1:]
		}
		for len(dq) > 0 && arr[i] >= arr[dq[0]] {
			dq = dq[1:]
		}
		dq = append(dq, i)

		if i >= k-1 {
			result = append(result, arr[dq[0]])
		}

	}
	return result

}

func main() {
	solution := &Solution{}
	arr := []int{9, 10, 9, -7, -4, -8, 2, -6}
	n := 3 // Size of the sliding window
	fmt.Println("Max of each window:", solution.maxSub(arr, n))
}
