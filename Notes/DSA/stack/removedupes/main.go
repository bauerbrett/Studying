package main

import "fmt"

type Solution struct{}

func (s *Solution) removeDupes(st string) string {
	stack := []rune{}

	for _, char := range st { // go through each char in string.
		if len(stack) > 0 { // if there is anything in stack already we need to compare the top to the new char
			top := stack[len(stack)-1] //grab top
			if char == top {           //compare top with new char
				stack = stack[:len(stack)-1] // if they are same pop the top
			} else {
				stack = append(stack, char) //else add teh new letter onto the stack
			}
		} else { //if nothing exist in stack add new char onto stack no matter what
			stack = append(stack, char)
		}

	}
	return string(stack)

}

func main() {

	solution := &Solution{}

	dataInput := "abba"

	fmt.Println("New line:", solution.removeDupes(dataInput))

}
