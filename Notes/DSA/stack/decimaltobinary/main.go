package main

import "fmt"

type Solution struct{}

func (s *Solution) decToBinary(num int) string {
	stack := []rune{}
	strRune := []rune{}
	for num > 0 { //Loop until num == 0
		remain := rune(num%2) + '0' //If you're working with binary digits (0 and 1), you need to explicitly convert the integers to their ASCII character equivalents ('0' and '1').
		//This can be done by adding '0' to the integer value to ensure you get the correct character.
		num /= 2                      // This gets the next number to divide by two until it == 0
		stack = append(stack, remain) //Put on the stack
	}
	for len(stack) > 0 { //While stack is not 0. THis is basically a while loop
		top := stack[len(stack)-1]     //Grab the top of stack
		stack = stack[:len(stack)-1]   // Pop the last element
		strRune = append(strRune, top) //Append to the rune slice
	}
	return string(strRune) //covert rune slice to string
}
func main() {
	solution := &Solution{}
	dataInput := 18
	fmt.Println("Output:", solution.decToBinary(dataInput))
}
