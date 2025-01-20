package main

import (
	"container/list" // Import the container/list package for using the doubly-linked list.
	"fmt"
)

// MyStack is a generic stack data structure.
type MyStack[T any] struct {
	stack list.List // Create a list (doubly-linked list) to store stack elements.
}
type Solution struct{}

// IsEmpty checks if the stack is empty.
func (s *MyStack[T]) IsEmpty() bool {
	return s.stack.Len() == 0 // Check if the length of the list is 0 to determine if the stack is empty.
}

// Push adds an element to the top of the stack.
func (s *MyStack[T]) Push(data T) {
	s.stack.PushBack(data) // Add the element to the back of the list, representing the top of the stack.
}

// Pop removes and returns the top element from the stack.
func (s *MyStack[T]) Pop() (T, error) {
	var zeroValue T
	if s.IsEmpty() {
		return zeroValue, fmt.Errorf("EmptyStackException") // Return an error if the stack is empty.
	}
	back := s.stack.Back()     // Get a reference to the back (top) element of the list.
	s.stack.Remove(back)       // Remove the top element from the list.
	return back.Value.(T), nil // Return the value of the top element as type T.
}

// Peek returns the top element of the stack without removing it.
func (s *MyStack[T]) Peek() (T, error) {
	var zeroValue T
	if s.IsEmpty() {
		return zeroValue, fmt.Errorf("EmptyStackException") // Return an error if the stack is empty.
	}
	back := s.stack.Back()     // Get a reference to the back (top) element of the list.
	return back.Value.(T), nil // Return the value of the top element as type T without removing it.
}

func (s *Solution) balancedParanth(st string) bool {
	stack := &MyStack[string]{}

	for _, char := range st {
		switch char {
		case '[':
			stack.Push("[")
		case '(':
			stack.Push("(")
		case '{':
			stack.Push("{")
		case ']':
			l, _ := stack.Peek()
			if l == "[" {
				stack.Pop()
			}
		case '}':
			r, _ := stack.Peek()
			if r == "{" {
				stack.Pop()
			}
		case ')':
			a, _ := stack.Peek()
			if a == "(" {
				stack.Pop()
			}

		}

	}
	return stack.IsEmpty()
}

func main() {

	solution := &Solution{}

	dataInput := "(]"
	result := solution.balancedParanth(dataInput)

	fmt.Println(result)
}
