package main

import "fmt"

type Solution struct{}

func (s *Solution) removeStar(st string) string {

	stack := []rune{}

	for _, char := range st { // loop through string
		if len(stack) > 0 { // if stack has item in it check if the char == *
			if char == '*' { // if it == *
				stack = stack[:len(stack)-1] // pop the top, this will remove the character to the left of *
			} else {
				stack = append(stack, char) //if it doesn't == *, then add it to stack
			}

		} else { // if stack has nothing in it just add the char into it
			stack = append(stack, char)
		}
	}
	return string(stack)

}

func main() {
	solution := &Solution{}
	dataInput := "abcd"
	fmt.Println("No Stars:", solution.removeStar(dataInput))
}
