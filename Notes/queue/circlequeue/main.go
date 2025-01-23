package main

import (
	"fmt"
	"math"
)

type CircularQueue struct {
	queue       []int
	front, rear int
	size        int
}

// Constructor to initialize the queue
func NewCircularQueue(size int) *CircularQueue {
	return &CircularQueue{
		queue: make([]int, size),
		front: -1,
		rear:  -1,
		size:  size,
	}
}

// Function to insert an element in the queue
func (cq *CircularQueue) Enqueue(element int) {
	if cq.front == (cq.rear+1)%(cq.size) {
		fmt.Println("Queue is Full")
	} else if cq.front == -1 { // Insert First Element
		cq.front, cq.rear = 0, 0
		cq.queue[cq.rear] = element
	} else if cq.rear == cq.size-1 && cq.front != 0 {
		cq.rear = 0
		cq.queue[cq.rear] = element
	} else {
		cq.rear = (cq.rear + 1) % cq.size
		cq.queue[cq.rear] = element
	}
}

// Function to delete an element from the queue
func (cq *CircularQueue) Dequeue() int {
	if cq.front == -1 {
		fmt.Println("Queue is Empty")
		return math.MinInt64
	}

	data := cq.queue[cq.front]
	if cq.front == cq.rear {
		cq.front, cq.rear = -1, -1
	} else {
		cq.front = (cq.front + 1) % cq.size
	}
	return data
}

// Function to display the elements of the queue
func (cq *CircularQueue) DisplayQueue() {
	if cq.front == -1 {
		fmt.Println("Queue is Empty")
		return
	}
	fmt.Print("Elements in the Circular Queue are: ")
	if cq.rear >= cq.front {
		for i := cq.front; i <= cq.rear; i++ {
			fmt.Print(cq.queue[i], " ")
		}
	} else {
		for i := cq.front; i < cq.size; i++ {
			fmt.Print(cq.queue[i], " ")
		}
		for i := 0; i <= cq.rear; i++ {
			fmt.Print(cq.queue[i], " ")
		}
	}
	fmt.Println()
}

// Main method to test the CircularQueue class
func main() {
	q := NewCircularQueue(5)

	// Inserting elements in the queue
	q.Enqueue(14)
	q.Enqueue(22)
	q.Enqueue(13)
	q.Enqueue(-6)

	// Display elements present in the queue
	q.DisplayQueue()

	// Deleting elements from the queue
	fmt.Println("Deleted value =", q.Dequeue())
	fmt.Println("Deleted value =", q.Dequeue())

	q.DisplayQueue()

	q.Enqueue(9)
	q.Enqueue(20)
	q.Enqueue(5)

	q.DisplayQueue()

	q.Enqueue(20)
}
