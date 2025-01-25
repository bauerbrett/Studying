package main

import "fmt"

// ListNode struct to represent a ListNode in the linked list
type ListNode struct {
	Val  int       // Data of the node
	Next *ListNode // Reference to the Next node in the list
}

// Solution struct to represent a Singly Linked List and perform various operations
type Solution struct {
	head *ListNode // head of the list, initially nil
}

// Inserts a new ListNode at the front of the list.
func (s *Solution) insert(Val int) {
	newNode := &ListNode{Val: Val} // Create a new node with the given Val
	newNode.Next = s.head          // Set the Next reference of the new node to the current head
	s.head = newNode               // Update the head of the list to the new node
}

// Inserts a new node after the given prevNode.
func (s *Solution) insertAfter(prevNode *ListNode, Val int) {
	// Check if the given prevNode is nil
	if prevNode == nil {
		fmt.Println("The given previous node cannot be nil")
		return
	}
	newNode := &ListNode{Val: Val} // Create a new node with the given Val
	newNode.Next = prevNode.Next   // Set the Next reference of the new node to the Next node of prevNode
	prevNode.Next = newNode        // Update the Next reference of prevNode to the new node
}

// Deletes the first occurrence of the specified key in the list
func (s *Solution) delete(key int) {
	temp := s.head
	var prev *ListNode // Initialize temp to head and prev to nil

	// Check if head node itself holds the key to be deleted
	if temp != nil && temp.Val == key {
		s.head = temp.Next // Update the head to the Next node
		return
	}

	// Search for the key to be deleted
	for temp != nil && temp.Val != key {
		prev = temp      // Update prev to the current node
		temp = temp.Next // Move to the Next node
	}

	// If key was not present in the list
	if temp == nil {
		return
	}

	// Unlink the node from the list
	prev.Next = temp.Next
}

// Searches for the key in the linked list and returns true if found, otherwise false
func (s *Solution) search(key int) bool {
	current := s.head // Initialize current to head

	for current != nil {
		if current.Val == key { // If the key is found, return true
			return true
		}
		current = current.Next // Move to the Next node
	}

	// Key not found, return false
	return false
}

func main() {
	list := &Solution{} // Create an empty list

	// Insert nodes into the list
	list.insert(1)
	list.insert(2)
	list.insertAfter(list.head.Next, 8)
	list.insert(3)

	// Search for a key in the list
	fmt.Println("Search 2:", list.search(2)) // true
	fmt.Println("Search 8:", list.search(8))

	// Delete a node from the list by key
	list.delete(2)

	// Check if the key is still present in the list
	fmt.Println("Search 2:", list.search(2)) // false
}
