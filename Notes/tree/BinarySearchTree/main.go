package main

import (
	"fmt"
)

// Using this to show how to implementation of binary search tree

type TreeNode struct {
	data  int
	right *TreeNode
	left  *TreeNode
}
type BTS struct {
	root *TreeNode
}

func (bts *BTS) insert(value int) {
	newNode := &TreeNode{data: value} // create the new node with the value given in func
	if bts.root == nil {              // if there is not root yet for the bts then make this value the root
		bts.root = newNode
	} else { // if there is a root then we need to figure out how far we need to go down the tree. Remember right side is always greater than the parent and the left
		current := bts.root  //grab current which is the root
		var parent *TreeNode // make a parent variable because this is what we are going to grab each time teh for loop runs. That way when the current changes to nil we have the last parent value.
		for current != nil { // once this == nil that means the last current.next ran made it nil
			parent = current                 // We capture this so when the current.next moves current to nil we still have the parent node to use outside the for loop
			if newNode.data > current.data { // if the new node is greater the current node data move to the next node to righ
				current = current.right // Make current the node to right so we can run this loop again
			} else {
				current = current.left //if it is not greater than move the current down to left of the tree
			}
		}
		if newNode.data > parent.data { //Once current hits nil from the last current.next take the parent node and compare it to the new node.
			parent.right = newNode //if greater than make the parent right point to the new node meaning it is the next greatest
		} else {
			parent.left = newNode // if it is less make the parent left point at the left
		}
	}
}

// The delete funcs
func (bts *BTS) findMin(root *TreeNode) *TreeNode {
	if bts.root == nil {
		fmt.Println("No tree values")
	} else {
		current := root
		var parent *TreeNode
		for current != nil {
			parent = current
			current = current.left
		}
		return parent
	}
	return nil
}

// Use recursion to keep changing root until root equals int. Then check if right and left exists
func (bts *BTS) helpdeleteNode(root *TreeNode, value int) *TreeNode {
	if root == nil {
		return root
	}
	if value > root.data {
		root.right = bts.helpdeleteNode(root.right, value)
	} else if value < root.data {
		root.left = bts.helpdeleteNode(root.left, value)
	} else {
		if root.right == nil {
			temp := root.left
			root = nil
			return temp
		}
		if root.left == nil {
			temp := root.right
			root = nil
			return temp
		}
		minNode := bts.findMin(root.right)
		root.data = minNode.data
		root.right = bts.helpdeleteNode(minNode, minNode.data)
	}
	return root
}

func (bts *BTS) deleteMethod(data int) *TreeNode {
	return bts.helpdeleteNode(bts.root, data)
}

func (bts *BTS) search(data int) bool { //search and return true if value exists and false if it doesn't
	current := bts.root  //set this to the root of tree
	for current != nil { //loop until you get a nil. If you get a nil that means the node does not exists and you would then exits loop and return false
		if current.data == data { // First check in loop, see if it current node value is equal with the search value
			return true //return true if they are equal
		} else if data > current.data { // If that is not ture move to elseif and else statements. We need to know if we need to move left or right now in the tree
			current = current.right //move right if data is greater than current node
		} else {
			current = current.left //move left if data is less than current node
		}
	}
	return false //If loop exits with a nil that means value does not exist and we can return false
}

func (bts *BTS) findHeight(root *TreeNode) int {
	current := root
	if current == nil {
		return 0
	}
	left := 1 + bts.findHeight(current.left)
	right := 1 + bts.findHeight(current.right)
	if left > right {
		return left
	} else {
		return right
	}
}

func main() {
	bts := &BTS{}
	bts.insert(20)
	bts.insert(40)
	bts.insert(10)
	bts.insert(60)
	bts.insert(45)
	fmt.Println(bts.root.right.data)
	//bts.deleteMethod(40)
	fmt.Println(bts.root.right.data)
	fmt.Println(bts.search(60))
	fmt.Println(bts.findHeight(bts.root))
}
