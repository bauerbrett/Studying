package main

//basically since the list is sorted check if there are dupes and once the node is bigger than the number we
// check move to the next one. No need to check all nodes since we know it is sorted and once the node val is bigger there can't be dupes.
import "fmt"

type ListNode struct {
	data int
	next *ListNode
}
type Solution struct{}

//[1, 1, 2]
func (s *Solution) removeDupes(head *ListNode) {
	//var check *ListNode
	current := head

	for current.next != nil { //loop until next element points at nil

		//fmt.Println(current.data, current.next.data)

		if current.data == current.next.data { // if they are equal

			current.next = current.next.next //skip the equal node. If it is at the end it would just point to nil and then the loop would finish

		} else { // If they are not equal just move current foward
			current = current.next
		}

	}

}

func (s *Solution) printNodes(head *ListNode) {
	if head == nil {
		fmt.Println("Head is nil. Cannont process!")

	}
	current := head
	for current != nil {
		fmt.Println(current.data)
		current = current.next

	}
}

func main() {

	solution := &Solution{}

	list := &ListNode{data: 1}

	list.next = &ListNode{data: 1}

	list.next.next = &ListNode{data: 1}
	solution.removeDupes(list)
	solution.printNodes(list)

}
