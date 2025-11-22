package main

import (
	"fmt"
	"strings"
)

type Solution struct{}

func (s *Solution) makeGreat(st string) string {

	stack := []rune{}

	for _, char := range st {
		if len(stack) > 0 {
			top := string(stack[len(stack)-1])
			if string(char) == strings.ToUpper(top) || strings.ToUpper(string(char)) == top {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, char)
			}

		} else {
			stack = append(stack, char)
		}
	}
	result := ""
	for i := len(stack) - 1; i >= 0; i-- {
		result = result + string(stack[i])
		stack = stack[:len(stack)-1]
	}
	return result

}

func main() {
	solution := &Solution{}

	dataInput := "AaBbCcDdEeff"

	fmt.Println("Great Line:", solution.makeGreat(dataInput))
}
