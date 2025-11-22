package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	data int
	next *ListNode
	prev *ListNode
}

//Couple ways I am think to solve this. First take ever node data element and push to dequeue and take both values off them at once and check if they are ==
// The other way is just create a string from the list and then reverse the list a create a string the otehr direction. If they are the same return true.

//new more efficient way is to find the tail of the list and the start of the list. move start foward with start.next and tail backwards with tail.prev
//if these items are not == return false, if they are continue checking until start and tail cross past each other

type Solution struct{}

func (s *Solution) validPalindrome(head *ListNode) bool {
	string1 := ""
	string2 := ""
	var temp *ListNode

	current := head
	fmt.Println(current.next.prev.data)
	for current != nil {
		string1 += strconv.Itoa(current.data)
		current = current.next
	}
	// [<1, 2 ,3,4,5]
	current1 := head
	for current1 != nil {
		temp = current1.prev          // makes the temp the prev of item
		current1.prev = current1.next // since we are flipping these around we change the prev of this element to point at the next
		current1.next = temp          // again since we are flipping around we have the next point to the temp which was the previous item

		current1 = current1.prev // current1.prev got reasigned to the next item in the list so me move current1 to that new element.
	}

	current1 = temp.prev //the reason we do this is because temp ends up the second to last in the reversed list. So we grab the prev item which gets us the last item on the list
	for current1 != nil {
		string2 += strconv.Itoa(current1.data)
		current1 = current1.next
	}
	if string1 == string2 {
		return true
	} else {
		return false
	}

}

func main() {

	solution := &Solution{}

	// Initialize the first node
	list := &ListNode{data: 1}

	// Add subsequent nodes with correct `prev` and `next` pointers
	list.next = &ListNode{data: 2, prev: list}
	list.next.next = &ListNode{data: 3, prev: list.next}
	list.next.next.next = &ListNode{data: 2, prev: list.next.next}
	list.next.next.next.next = &ListNode{data: 1, prev: list.next.next.next}

	fmt.Println("The list is a valid palindrome:", solution.validPalindrome(list))

}
