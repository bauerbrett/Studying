package main

import (
	"fmt"
)

type Node struct {
	val   int
	left  *Node
	right *Node
}

type Solution struct{}

/*
preOrder
1. get root
2. visit root
3. visit left traverse
4. visit right traverse
*/

func (s Solution) preOrder(root *Node) {

	if root == nil {
		return
	}
	fmt.Println(root.val)
	s.preOrder(root.left)
	s.preOrder(root.right)
}

func main() {
	// Create a sample BST:
	//         8
	//       /   \
	//      3     10
	//     / \      \
	//    1   6      14
	//       /
	//      4
	root := &Node{val: 8}
	root.left = &Node{val: 3}
	root.right = &Node{val: 10}
	root.left.left = &Node{val: 1}
	root.left.right = &Node{val: 6}
	root.left.right.left = &Node{val: 4}
	root.right.right = &Node{val: 14}

	s := &Solution{}

	s.preOrder(root)
}
