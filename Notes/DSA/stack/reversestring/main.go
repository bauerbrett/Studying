package main

import "fmt"

type Solution struct{}

func (s *Solution) revString(st string) string {
	stack := []rune{}
	newStr := []rune{}
	for _, char := range st {
		stack = append(stack, char)
	}
	for i := 0; i < len(st); i++ {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		newStr = append(newStr, top)
	}
	return string(newStr)

}

func main() {
	solution := &Solution{}
	dataInput := "OpenAI"
	fmt.Println("Reverse string is:", solution.revString(dataInput))

}
