package main

import (
	"container/list"
	"fmt"
)

//Bind the binary but this time with the better technique and queues.

/*
Key Idea: Generate Binary Numbers Using Patterns
Start with "1" as the first binary number.
Use the pattern:

For a binary number x, the next two binary numbers are:

x+0: Append 0 to the end of x. = 10
x+1: Append 1 to the end of x. = 10

This mimics how binary numbers grow in sequence:
From 1, we can generate 10 (append 0) and 11 (append 1).
From 10, we can generate 100 and 101.
From 11, we can generate 110 and 111.
*/

type Solution1 struct{}

type Queue struct {
	data *list.List
}

func newQueue() *Queue {
	return &Queue{
		data: list.New(),
	}
}

func (s *Solution1) getBinary(n int) []string {

	queue := newQueue()
	res := make([]string, n)
	queue.data.PushBack("1")

	for i := 0; i < n; i++ {
		front := queue.data.Front().Value
		res[i] = front.(string)
		queue.data.Remove(queue.data.Front())
		queue.data.PushBack(res[i] + "0")
		queue.data.PushBack(res[i] + "1")

	}
	return res
}

func main() {

	solution := &Solution1{}

	dataInput := 8

	fmt.Println("Binary Numbers:", solution.getBinary(dataInput))
}
