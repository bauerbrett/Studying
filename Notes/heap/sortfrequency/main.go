package main

import (
	"container/heap"
	"fmt"
)

/*
🧠 Problem Summary: Frequency Sort Using Max Heap (Go)
Goal: Given a string, return it sorted by character frequency (most frequent chars first).

Approach:
Count frequencies using a map[rune]int.
Create a custom maxHeap that holds freqEntry{char, freq} structs.
Implement heap.Interface methods: Len, Less, Swap, Push, Pop.
Use heap.Push() to insert all entries into the heap.
Pop from the heap, appending each character freq times to the result string.
Return the final string.

Heap Logic:
Less(i, j) returns true if freq[i] > freq[j] → max-heap.
Heap ensures the most frequent character is always on top.

Time Complexity:
Heap insertions: O(n log k) where k is unique characters.
Result string build: O(n)

Total: O(n log k)

Python difference:
Again python is much easier for a few reasons. \
Has a built in counter class to count the string char and put in a dictionary
Can put a tuple in the heap directly whereas in go we need to create a struct with a char and freq field and place the struct in the heap
Can use a * to put the char in result slice by the freq whereas go we need to loop through the freq amount and add.
*/

type freqEntry struct {
	char rune
	freq int
}
type IntHeap []freqEntry

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Swap(j, i int)      { h[j], h[i] = h[i], h[j] }
func (h IntHeap) Less(j, i int) bool { return h[j].freq > h[i].freq }

func (h *IntHeap) Push(n interface{}) {
	*h = append(*h, n.(freqEntry))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func sortFreq(s string) string {
	maxHeap := &IntHeap{}
	charCount := map[rune]int{}
	heap.Init(maxHeap)
	for _, v := range s {
		charCount[v]++
	}
	for k, v := range charCount {
		heap.Push(maxHeap, freqEntry{char: k, freq: v})
	}

	result := []rune{}
	for maxHeap.Len() > 0 {
		popped := heap.Pop(maxHeap).(freqEntry)
		//result += strings.Repeat(string(char), freq)
		//this would also work instead of for loop
		for i := 0; i < popped.freq; i++ {
			result = append(result, popped.char)
		}
		//fmt.Println(pop)
	}
	return string(result)

}

func main() {
	s := "trersesess"
	fmt.Println(sortFreq(s))

}
