package main

import "fmt"

//This is a sliding window problem where a queue is used. More efficient would be to use a dequeue
/*
Example 1:

Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
Output: [3,3,5,5,6,7]s
Explanation:
Window position                Max
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
Example 2:

Input: nums = [1], k = 1
Output: [1]
*/

type Solution struct{}

func (s *Solution) maxSub(arr []int, n int) []int {

	//we need to return max of the sum of n in the array
	// so for example say we have a {1, 2, 3, 1, 4, 5, 2, 3, 6} and n = 3
	//as long as the arr/queue is == n we want to add amount of values that n =
	//so we start at 1 and the greatest int in the first 3 is 3, so 3 goes to result
	//we then dequeue 1 and move to 2. Then 2 and the next 2, the greatest is still 3. so 3 goes into the result
	// on and on until n < len(arr)
	//result [3, 3, ...etc]
	result := []int{} //result slice that will returned with the maxes
	// Ensure there are enough elements in the array for at least one window
	for len(arr) >= n {
		queue := []int{} //reset the queue when the window moves

		// Process the first 'n' elements of the current array. This is teh window
		for i := 0; i < n; i++ {
			// Maintain order in the deque
			for len(queue) > 0 && queue[len(queue)-1] < arr[i] {
				queue = queue[:len(queue)-1] // Pop smaller elements from the back
			}
			queue = append(queue, arr[i]) // Append the current element
		}

		// The largest element in the current window is at the front of the deque
		result = append(result, queue[0])

		// Slide the window by removing the first element of the array
		arr = arr[1:]
	}
	return result
}

func main() {

	solution := &Solution{}

	arr := []int{9, 10, 9, -7, -4, -8, 2, -6}
	n := 5

	fmt.Println("Max of all:", solution.maxSub(arr, n))

}
