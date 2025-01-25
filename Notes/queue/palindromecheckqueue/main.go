package main

import (
	"fmt"
	"strings"
)

type Solution struct{}

func (s *Solution) checkPalindrome(st string) bool {

	queue := []rune{} //dequeue a double ended queue to remove from front and back

	formatStr := strings.ToLower(st)
	formatStr = strings.Replace(formatStr, " ", "", -1) // replace the space " " with not space "", and -1 means do it for all occurences of " "
	fmt.Println(formatStr)

	for _, value := range formatStr {
		queue = append(queue, value)
	}

	for len(queue) > 1 {

		front := queue[0]
		back := queue[len(queue)-1]
		queue = queue[1 : len(queue)-1]

		if front != back {
			return false

		}

	}
	return true

}

func main() {

	solution := &Solution{}

	dataInput := "A man a plan a canal Panama"

	if solution.checkPalindrome(dataInput) {
		fmt.Println("This is a valid palindrome:", dataInput)
	} else {
		fmt.Println("Not a valid palindrome!")
	}

}
