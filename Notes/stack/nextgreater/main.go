package main

import "fmt"

type Solution struct{}

//How to achieve Monotonic Increasing Stack?
//To achieve a monotonic increasing stack, you would typically push elements onto the stack while ensuring that the stack maintains a increasing order from bottom to top.
//When pushing a new element, you would pop elements from the stack that are smaller than the new element until the stack maintains the desired monotonic increasing property.

//To achieve a monotonic increasing stack, you can follow these step-by-step approaches:

//Initialize an empty stack.
//Iterate through the elements and for each element:
//while stack is not empty AND top of stack is more than the current element
//	pop element from the stack
//Push the current element onto the stack.
//At the end of the iteration, the stack will contain the monotonic increasing order of elements.

/*
How to achieve Monotonic Decreasing Stack?
To achieve a monotonic decreasing stack, you would typically push elements onto the stack while ensuring that the stack maintains a
decreasing order from bottom to top. When pushing a new element, you would pop elements from the stack that are greater than the new element until the
stack maintains the desired monotonic decreasing property.

To achieve a monotonic decreasing stack, you can follow these step-by-step approaches:

Start with an empty stack.
Iterate through the elements of the input array.
While stack is not empty AND top of stack is less than the current element:
	pop element from the stack
Push the current element onto the stack.
After iterating through all the elements, the stack will contain the elements in monotonic decreasing order.

*/

//This uses a monotonic decreasing stack
// nextGreater finds the next greater element for each item in the array.
func (s *Solution) nextGreater(arr []int) []int {
	stack := []int{}                  // Monotonic stack to keep track of elements
	greatArr := make([]int, len(arr)) // Result array

	// Traverse the array from right to left
	for i := len(arr) - 1; i >= 0; i-- {
		// Pop elements from the stack that are less than or equal to the current element
		for len(stack) > 0 && stack[len(stack)-1] <= arr[i] {
			stack = stack[:len(stack)-1]
		}

		// If the stack is empty, no greater element exists
		if len(stack) == 0 {
			greatArr[i] = -1
		} else {
			// The top of the stack is the next greater element
			greatArr[i] = stack[len(stack)-1]
			//[5, 0, 3, -1]
		}

		// Push the current element onto the stack
		stack = append(stack, arr[i])
	}

	return greatArr
}

func main() {
	solution := &Solution{}

	dataInput := []int{13, 100, 6, 12}
	fmt.Println("Next Greater Slice:", solution.nextGreater(dataInput)) // Output: [5 25 25 -1]
}
