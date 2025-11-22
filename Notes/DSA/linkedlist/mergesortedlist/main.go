package main

import "fmt"

type Solution struct{}

type ListNode struct {
	data int
	next *ListNode
}

/*
func (l *ListNode) insert(new *ListNode) {
	if l != nil {
		next := l.next
		l.next = new
		new.next = next
	} else {
		l.data = new.data

	}
}
*/

func (s *Solution) mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// Initialize a dummy node and a current pointer
	dummy := &ListNode{data: -1}
	current := dummy //use this to keep teh first place in the front so we can return it

	// Traverse through both lists until one is exhausted
	for l1 != nil && l2 != nil {
		// Compare nodes and append the smaller one to current
		if l1.data < l2.data {
			current.next = l1
			l1 = l1.next
		} else {
			current.next = l2
			l2 = l2.next
		}
		current = current.next
	}

	// Append the remaining nodes from the non-empty list
	if l1 != nil {
		current.next = l1
	} else {
		current.next = l2
	}

	// Return the merged sorted list
	return dummy.next
}

func main() {
	solution := &Solution{}

	list1 := &ListNode{data: 1}
	list1.next = &ListNode{data: 3}
	list1.next.next = &ListNode{data: 5}

	list2 := &ListNode{data: 2}
	list2.next = &ListNode{data: 4}
	list2.next.next = &ListNode{data: 6}

	result := solution.mergeTwoLists(list1, list2)
	//print(result)
	for result != nil {
		fmt.Println(result.data)
		result = result.next
	}

}
