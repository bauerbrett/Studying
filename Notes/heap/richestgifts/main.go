package main

import (
	"container/heap"
	"fmt"
	"math"
)

/*
Dumb Ass Question:
Basically we need to find square root of biggest gift pile. So we can use a max heap because the biggest gifts will always be on top.
After taking gifts off we square root them, the square root number gets pushed back onto the heap. We do this k amount of times.
So if k is 2 we will pop and push 2 gift piles and keep their square roots. Because it is a heap, when we push back the square root,
the heap will reorder and max gift pile will be on top for next pop.

You're presented with several piles of gifts, with each pile containing a certain number of gifts. Every second, you'll engage in the following activity:

Pick the pile that contains the highest number of gifts. If multiple piles share this distinction, you can select any of them.
Compute the square root of the number of gifts in the selected pile, and then leave behind that many gifts (rounded down). Take all the other gifts from this pile.
You'll do this for "k" seconds. The objective is to find out how many gifts would still remain after these "k" seconds.
Examples
Input: gifts = [4, 9, 16], k = 2
Expected Output: 11
Justification:
Take from third pile (16 gifts): leave (  ) = 4 gifts, take 12. Remaining gifts = [4, 9, 4]
Take from second pile (9 gifts): leave (  ) = 3 gifts, take 6. Remaining gifts = [4, 3, 4]
*/

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h *IntHeap) Pop() interface{} {
	old := *h     // = the heap
	n := len(old) // get len of heap
	x := old[n-1] //
	*h = old[0 : n-1]
	return x
}
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int)) //just add the int to end of slice

}

func giftPile(gifts []int, k int) int {

	maxHeap := &IntHeap{}
	heap.Init(maxHeap)
	for _, gift := range gifts {
		heap.Push(maxHeap, -gift)
	}

	for i := 0; i < k; i++ {
		current := -heap.Pop(maxHeap).(int)

		heap.Push(maxHeap, -int((math.Sqrt(float64(current)))))
	}
	result := 0
	for maxHeap.Len() > 0 {
		result += -heap.Pop(maxHeap).(int)
	}
	return result

}
func main() {
	list := []int{4, 9, 16}
	k := 2
	fmt.Println(giftPile(list, k))
}
