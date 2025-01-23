package main

import "fmt"

type Solution struct{}

func (s *Solution) revQueue(qu []int) []int {
	stack := []int{}
	newQueue := []int{}
	stack = append(stack, qu...)
	for i := len(stack) - 1; i >= 0; i-- {
		newQueue = append(newQueue, stack[i])
	}
	return newQueue
}

func main() {
	solution := &Solution{}
	dataInput := []int{1, 2, 3, 4, 5}
	fmt.Println("Updated Queue:", solution.revQueue(dataInput))

}
