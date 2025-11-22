package main

import (
	"container/heap"
	"fmt"
)

/*
Given a collection of sticks with different lengths. To combine any two sticks, there's a
cost involved, which is equal to the sum of their lengths.

Connect all the sticks into a single one with the minimum possible cost. Remember, once two sticks are combined,
they form a single stick whose length is the sum of the lengths of the two original sticks.

Examples
Input: [2, 4, 3]
Expected Output: 14
Justification: Combine sticks 2 and 3 for a cost of 5. Now, we have sticks [4,5]. Combine these at a cost of 9. Total cost = 5 + 9 = 14.
*/

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func minConnect(l []int) int {
	maxHeap := &IntHeap{}
	heap.Init(maxHeap)
	for _, num := range l {
		heap.Push(maxHeap, num)
	}
	total := 0
	for maxHeap.Len() > 1 {
		one := heap.Pop(maxHeap).(int)
		two := heap.Pop(maxHeap).(int)
		new := one + two
		total += new
		heap.Push(maxHeap, new)

	}
	return total

}

func main() {
	l := []int{5, 5, 5, 5}
	fmt.Println(minConnect(l))
}
