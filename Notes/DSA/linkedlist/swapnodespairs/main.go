package main

import "fmt"

type Solution struct{}

type ListNode struct {
	data int
	next *ListNode
}

func (s *Solution) swapNodes(head *ListNode) *ListNode {

	dummy := &ListNode{} // we need this so we can return its next which is going to be the new head
	dummy.next = head    //point it initially at the first item
	previous := dummy    //make this previous start at the first nil item. This is the one taht will keep getting updated as we move foward in the list

	for head != nil && head.next != nil {
		firstNode := head       //get the nodes first so we dont lose them
		secondNode := head.next // same as above

		firstNode.next = secondNode.next // make sure the first node is point at the node behind the second node
		secondNode.next = firstNode      // make sure the the second node now points at the first node
		previous.next = secondNode       // make sure the prev tracker points at the second node now

		head = firstNode.next // move head to the node after the swapped first node
		previous = firstNode  // move previous up to where the initial first node is which is spot 2 now
	}

	return dummy.next

}

func main() {

	solution := &Solution{}

	list := &ListNode{data: 1}
	list.next = &ListNode{data: 2}
	list.next.next = &ListNode{data: 3}
	list.next.next.next = &ListNode{data: 4}

	result := solution.swapNodes(list)
	for result != nil {
		fmt.Println(result.data)
		result = result.next
	}

}
