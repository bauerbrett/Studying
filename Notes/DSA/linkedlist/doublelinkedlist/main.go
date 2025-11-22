package main

import "fmt"

// DLNode represents a node in the doubly linked list.
type DLNode struct {
	Val  int
	Prev *DLNode
	Next *DLNode
}

// Solution is a struct that manages the doubly linked list.
type Solution struct {
	head *DLNode
}

// Insert inserts a new node with the given Val at the front of the list.
func (s *Solution) Insert(Val int) {
	newNode := &DLNode{Val: Val}
	newNode.Next = s.head
	if s.head != nil { // if there was nothing to begin with do nothing, else make the new node the previous to the old head.
		s.head.Prev = newNode
	}
	s.head = newNode
}

// InsertAfter inserts a new node with the given Val after the provided prevNode.
func (s *Solution) InsertAfter(prevNode *DLNode, Val int) {
	if prevNode == nil {
		fmt.Println("The given previous node cannot be nil") // Check if the previous node exists
		return
	}
	newNode := &DLNode{Val: Val}
	newNode.Next = prevNode.Next // insert before the old nodes next, which means the new next should point to the old next
	prevNode.Next = newNode      //Since we are inserting after the prevnode, teh prevnode should point to newnode
	newNode.Prev = prevNode      //Since we are inserting after the prevnode, the new node previous should point to the node we inserted after
	if newNode.Next != nil {     //
		newNode.Next.Prev = newNode //So the node that gets pushed back after new node gets inserted needs its prev field to point to the new node inserted
	}
	//if it is nil that means it is the last item in the list, it is the rear.
}

// Delete removes the node with the given Val from the list.
func (s *Solution) Delete(key int) {
	temp := s.head
	for temp != nil {
		if temp.Val == key {
			if temp.Prev != nil {
				temp.Prev.Next = temp.Next //since you are deleting the element you need to put the next of the prev item as the next item after temp gets removed
			} else {
				s.head = temp.Next //if there is nothing in front of this temp just point the new head at teh item behind next
			}
			if temp.Next != nil { //
				temp.Next.Prev = temp.Prev // make the next items prev as the item behind temp
			}
			return
		}
		temp = temp.Next
	}
}

// Search searches for a node with the given Val and returns true if found, false otherwise.
func (s *Solution) Search(key int) bool {
	current := s.head
	for current != nil {
		if current.Val == key {
			return true // Return true if the key is found in the list
		}
		current = current.Next
	}
	return false // Return false if the key is not found in the list
}

func main() {
	dll := &Solution{}
	dll.Insert(1)
	dll.Insert(2)
	dll.Insert(3)

	fmt.Println("Search 2:", dll.Search(2)) // true
	dll.Delete(2)
	fmt.Println("Search 2:", dll.Search(2)) // false
}
