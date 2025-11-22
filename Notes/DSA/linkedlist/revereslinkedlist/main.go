package main

import "fmt"

type Solution struct{}

type ListNode struct {
	Val  int
	Next *ListNode
}

//[nil,1,2,3,4,5,6,7]

// we need current-1 to be prev and we want 2 to be current
//Remade it here myself a second time to solve
func (s *Solution) reverseList2(head *ListNode) *ListNode {
	var prev *ListNode // we want this to just be a variable and bot actually make anything yet. We will put something in here with the first prev.
	current := head    // First item

	for current != nil { // In the above once we get to 7 there will be no next so current.next at that point will be nil, and when we assign next to current, current will == nil
		next := current.Next // track the next item so we can assign current in the next loop.
		current.Next = prev  // Flip it backward for the current in this loop. Make current point to the left at prev instead of right to the next item.

		prev = current // Now we can flip them. prev will == the current item so in the next item we can point the current.next to this one.
		current = next // Shift current to the next item we stored above.
	}

	return prev //Once current hits nil this prev will be the last item that is current so it will be the head with all items going left instead of right now.

}

// Solution represents the solution struct.

// reverseList reverses the order of a linked list.
func (s *Solution) reverseList(head *ListNode) *ListNode {
	var prev *ListNode // Initialize a pointer to the previous node
	current := head    // Initialize a pointer to the current node

	// Traverse the linked list
	//Note: So for this it is not about moving the nodes around. It is all about changing what next points to in the list
	for current != nil {
		next := current.Next // Store a pointer to the next node
		current.Next = prev  // Reverse the direction of the current node's pointer

		//The ones above moved the points, now this resets them to get ready to change what the next node points at
		prev = current // Move the prev pointer to the current node
		current = next // Move the current pointer to the next node
	}

	return prev // Return the new head of the reversed list
}

// printList prints the values of the linked list.
func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val) // Print the value of the current node
		head = head.Next            // Move to the next node
	}
	fmt.Println()
}

func main() {
	solution := Solution{}

	// Test case 1
	head1 := &ListNode{Val: 3}
	head1.Next = &ListNode{Val: 5}
	head1.Next.Next = &ListNode{Val: 2}
	printList(solution.reverseList(head1)) // Expected Output: 2 5 3

	// Test case 2
	head2 := &ListNode{Val: 7}
	printList(solution.reverseList(head2)) // Expected Output: 7

	// Test case 3
	head3 := &ListNode{Val: -1}
	head3.Next = &ListNode{Val: 0}
	head3.Next.Next = &ListNode{Val: 1}
	printList(solution.reverseList2(head3)) // Expected Output: 1 0 -1
}
