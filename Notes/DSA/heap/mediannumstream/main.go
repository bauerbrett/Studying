package main

import (
	"container/heap"
	"fmt"
)

/*
Design a class to calculate the median of a stream of numbers. The class should have the following two methods:

insertNum(int num): stores the number in the class
findMedian(): returns the median of all numbers inserted in the class
If the count of numbers inserted in the class is even, the median will be the average of the two middle numbers.

Example 1:

1. insertNum(3)
2. insertNum(1)
3. findMedian() -> output: 2.0
4. insertNum(5)
5. findMedian() -> output: 3.0
6. insertNum(4)
7. findMedian() -> output: 3.5

Odd number of values: If the data set has an odd number of values, the median is the middle number in the ordered list.
Even number of values: If the data set has an even number of values, the median is the average (mean) of the two middle numbers in the ordered list.

so I did it this way but the more optomized way is to use two heaps lower and upper, lower is max heap and upper is minheap. This ensure the root nodes of each
are always the two middle number if even or the lower holds one extra if odd so we take the top node from lower if odd.

insert into lower first if possible but if it is bigger than lower[0] push into upper. Create balance logic to keep the heap balanced.
Pop from other side and push to the other heap.

*/

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(*h)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) insertNum(n int) {
	heap.Push(h, n)
}

func (h *IntHeap) findMedian() any {
	var median any
	newList := []int{}
	for h.Len() > 0 {
		num := heap.Pop(h)
		newList = append(newList, num.(int))
	}
	for _, num := range newList {
		heap.Push(h, num)
	}
	if len(newList)%2 == 0 {
		combined := (newList[len(newList)/2] + newList[(len(newList)/2)-1])
		median = float32(combined) / float32(2)
		return median
	} else {
		median = newList[len(newList)/2]
	}
	return median.(int)
}
func main() {
	h := &IntHeap{}
	heap.Init(h)
	h.insertNum(3)
	h.insertNum(1)
	fmt.Println(h.findMedian())
	h.insertNum(5)
	fmt.Println(h.findMedian())
	h.insertNum(4)
	fmt.Println(h.findMedian())
}
