package main

import "fmt"

type Solution struct{}

func (s *Solution) findGreater(arr []int) []int {
	//newArr := []int{}
	stack := []int{}                //create stack
	for i := 0; i < len(arr); i++ { //lopp through each element in arr
		temp := []int{}                                       // creat temp stack
		for len(stack) > 0 && stack[len(stack)-1] >= arr[i] { //loop through and if stack > then 0 and is bigger than arr index pop it and put in temp stack
			temp = append(temp, stack[len(stack)-1]) //put top in in temp stack
			stack = stack[:len(stack)-1]             //remake stack witout the top
		}
		stack = append(stack, arr[i]) //put i element from arr into the stack

		stack = append(stack, temp...) // append the bigger elements from temp back into normal stack after teh smaller i element

	}
	fmt.Println(stack)
	return stack
}

func main() {
	solution := &Solution{}
	dataInput := []int{8, 3, 5, 6, 100, 43}
	fmt.Println("In order:", solution.findGreater(dataInput))
}
