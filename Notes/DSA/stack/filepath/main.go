package main

import "fmt"

type Solution struct{}

// Try and redo this problem useing string.split and string.join with a stack. You can choose what character to sepearte the string with split.
//And you can join with specific char also.

func (s *Solution) simplifyPath(st string) string {

	doubDig := false

	stack := []rune{}

	slashCount := 0
	numCount := 0

	for _, char := range st {

		if char != '/' && char != '.' {
			stack = append(stack, char)
		}
		if slashCount > 0 && char == '/' {
			continue
		} else if slashCount == 0 && char == '/' {
			stack = append(stack, char)
		}
		if char == '/' {
			slashCount += 1
		}
		if char != '/' && char != '.' && slashCount > 0 {
			slashCount -= 1
		}
		if char == '.' {
			numCount += 1
		}
		if char != '.' && numCount > 0 {
			numCount -= 1
		}
		if numCount == 2 {
			doubDig = true
		}
	}
	if stack[len(stack)-1] == '/' && len(stack) > 1 {
		stack = stack[:len(stack)-1]
	}
	if doubDig && len(stack) > 2 {
		stack = stack[:len(stack)-2]
	}

	fmt.Println(string(stack))

	return string(stack)

}

func main() {

	solution := &Solution{}

	dataInput := "/a//b////c/d//././/.."

	fmt.Print("Fixed Path:", solution.simplifyPath(dataInput))

}
