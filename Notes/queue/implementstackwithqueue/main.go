package main

import (
	"container/list"
	"fmt"
)

type MyStack struct {
	q *list.List
}

func Constructor() MyStack {
	return MyStack{
		q: list.New(),
	}
}

// THis is basically he main piece of the program.
// Since we need to use this as a queue we need to flip all the values around.
// So when we push someting onto the stack on teh top, we flip everything around so it goes to the front.
// This way when we do top, pop it will act as a queue and grab that last value.
func (m *MyStack) Push(x int) {
	m.q.PushBack(x)
	for i := 0; i < m.q.Len()-1; i++ {
		m.q.PushBack(m.q.Front().Value) // for each item in the len of the stack we send the front item to the back shifting everything to the front.
		fr := m.q.Front()               // get the front since the last step just created a copy
		m.q.Remove(fr)                  // delete the original value since we have the copy in the back now.
	}
}

func (m *MyStack) Pop() int {
	fr := m.q.Front()
	ret := fr.Value.(int)
	m.q.Remove(fr)
	return ret
}

func (m *MyStack) Top() int {
	fr := m.q.Front()
	return fr.Value.(int)
}

func (m *MyStack) Empty() bool {
	return m.q.Len() == 0
}

func main() {
	solution := Constructor()
	solution.Push(5)
	solution.Push(10)
	fmt.Println(
		solution.Top(),
		solution.Pop(),
		solution.Empty())

}
