package main

import (
	"fmt"
)

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

/*
inorder
1. given root
2. traverse down to bottom left
3. visit process (value) the current node
4. traverse right
*/

type Solution struct{}

func (s Solution) inOrder(root *TreeNode) {
	if root == nil {
		return
	}
	s.inOrder(root.left)
	fmt.Println(root.val)
	s.inOrder(root.right)

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
	root := &TreeNode{val: 8}
	root.left = &TreeNode{val: 3}
	root.right = &TreeNode{val: 10}
	root.left.left = &TreeNode{val: 1}
	root.left.right = &TreeNode{val: 6}
	root.left.right.left = &TreeNode{val: 4}
	root.right.right = &TreeNode{val: 14}

	it := Solution{}
	fmt.Println("Inorder Traversal (Sorted Order):")
	it.inOrder(root) // Expected output: 1 3 4 6 8 10 14
	fmt.Println()
}
